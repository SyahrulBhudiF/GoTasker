package queue

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisQueue struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisQueue(client *redis.Client, ctx context.Context) *RedisQueue {
	return &RedisQueue{client: client, ctx: ctx}
}

func (r *RedisQueue) AddTask(queueName, task string) error {
	if err := r.client.LPush(r.ctx, queueName, task).Err(); err != nil {
		logrus.Errorf("Failed to push task: %v", err)
		return err
	}
	return nil
}

func (r *RedisQueue) PopTask(queueName string) (string, error) {
	result, err := r.client.RPop(r.ctx, queueName).Result()
	return result, err
}
