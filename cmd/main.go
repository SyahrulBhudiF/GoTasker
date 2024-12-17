package main

import (
	"context"
	"github.com/SyahrulBhudiF/GoTasker/pkg"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

// Example of how to use the package
func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping Redis
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatal("Failed to connect to Redis:", err)
	}

	// Initialize Redis Connection
	pkg.Init(redisClient)

	// Register Task
	pkg.RegisterTask("send-email", func(ctx context.Context, payload string) error {
		logrus.Infof("Sending email with payload: %s", payload)
		select {
		case <-time.After(2 * time.Second): // Simulate email sending
			logrus.Info("Email sent successfully!")
			return nil
		case <-ctx.Done():
			logrus.Warn("Email task timeout!")
			return ctx.Err()
		}
	})

	// Add Task to Queue
	if err := pkg.AddTask("task-queue", "send-email"); err != nil {
		logrus.Fatal("Failed to add task:", err)
	}

	// Start Workers
	pkg.StartWorker("task-queue", 3, 5*time.Second)

	// Prevent main from exiting
	select {}
}
