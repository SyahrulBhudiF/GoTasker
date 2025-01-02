package queue

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisQueue struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisQueue(client *redis.Client, ctx context.Context) (*RedisQueue, error) {
	if client == nil {
		return nil, errors.New("redis client is nil")
	}
	if ctx == nil {
		return nil, errors.New("context is nil")
	}
	return &RedisQueue{client: client, ctx: ctx}, nil
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
