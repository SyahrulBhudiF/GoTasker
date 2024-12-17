package goTasker

import (
	"context"
	"github.com/SyahrulBhudiF/GoTasker/internal/queue"
	registry2 "github.com/SyahrulBhudiF/GoTasker/internal/registry"
	registry "github.com/SyahrulBhudiF/GoTasker/internal/worker"
	"github.com/redis/go-redis/v9"
	"time"
)

var redisQueue *queue.RedisQueue

func Init(redisAddr *redis.Client) {
	redisQueue = queue.NewRedisQueue(redisAddr, context.Background())
}

func RegisterTask(name string, handler func(ctx context.Context, payload string) error) {
	registry2.RegisterTask(name, handler)
}

func AddTask(queueName, task string) error {
	return redisQueue.AddTask(queueName, task)
}

func StartWorker(queueName string, workerCount int, timeout time.Duration) {
	registry.StartWorker(queueName, redisQueue, workerCount, timeout)
}
