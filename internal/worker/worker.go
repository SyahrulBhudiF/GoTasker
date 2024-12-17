package registry

import (
	"context"
	"github.com/SyahrulBhudiF/GoTasker/internal/queue"
	registry2 "github.com/SyahrulBhudiF/GoTasker/internal/registry"
	"time"

	"github.com/sirupsen/logrus"
)

func StartWorker(queueName string, redisQueue *queue.RedisQueue, workerCount int, timeout time.Duration) {
	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			logrus.Infof("Worker %d started", workerID)
			for {
				task, err := redisQueue.PopTask(queueName)
				if err != nil {
					logrus.Errorf("Worker %d failed to pop task: %v", workerID, err)
					continue
				}

				logrus.Infof("Worker %d processing task: %s", workerID, task)
				handler := registry2.GetTaskHandler(task)
				if handler == nil {
					logrus.Errorf("Worker %d no handler found for task: %s", workerID, task)
					continue
				}

				ctx, cancel := context.WithTimeout(context.Background(), timeout)
				if err := handler(ctx, "payload-example"); err != nil {
					logrus.Errorf("Task %s failed: %v", task, err)
				} else {
					logrus.Infof("Task %s completed successfully", task)
				}
				cancel()
			}
		}(i + 1)
	}
}
