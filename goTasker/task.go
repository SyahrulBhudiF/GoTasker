package goTasker

import (
	"context"
	"time"

	"github.com/SyahrulBhudiF/GoTasker/internal/queue"
	"github.com/SyahrulBhudiF/GoTasker/internal/registry"
	"github.com/SyahrulBhudiF/GoTasker/internal/scheduler"
	worker "github.com/SyahrulBhudiF/GoTasker/internal/worker"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var redisQueue *queue.RedisQueue

func Init(redisAddr *redis.Client) {
	var err error
	redisQueue, err = queue.NewRedisQueue(redisAddr, context.Background())
	if err != nil {
		logrus.Errorf("Failed to initialize RedisQueue: %v", err)
	}
	logrus.Infof("RedisQueue initialized")
}

func RegisterTask(name string, handler func(ctx context.Context, payload string) error) {
	registry.RegisterTask(name, handler)
}

func AddTask(queueName, task string) error {
	return redisQueue.AddTask(queueName, task)
}

func StartWorker(queueName string, workerCount int, timeout time.Duration) {
	worker.StartWorker(queueName, redisQueue, workerCount, timeout)
}

func InitScheduler() {
    scheduler.Init()
}

func ScheduleTask(second int, queueName string, taskName string) {
    scheduler.AddJob(second, func() {
        logrus.Infof("Scheduled task: %s", taskName)
        AddTask(queueName, taskName)
    })
}