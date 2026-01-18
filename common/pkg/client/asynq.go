package client

import (
	"github.com/hibiken/asynq"
)

func NewAsynqServer(redisClient *RedisClient) *asynq.Server {
	return asynq.NewServerFromRedisClient(redisClient.Client, asynq.Config{})
}

func NewAsynqClient(redisClient *RedisClient) *asynq.Client {
	return asynq.NewClientFromRedisClient(redisClient.Client)
}
