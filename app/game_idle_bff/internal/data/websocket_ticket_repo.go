package data

import (
	commonclient "common/pkg/client"
	"context"
	"fmt"
	"game_idle_bff/internal/biz/repo"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ repo.WebSocketTicketRepo = (*WebSocketTicketRepo)(nil)

type WebSocketTicketRepo struct {
	redisClient    *commonclient.RedisClient
	redisKeyFormat string
}

func NewWebSocketTicketRepo(redisClient *commonclient.RedisClient) repo.WebSocketTicketRepo {
	return &WebSocketTicketRepo{
		redisClient:    redisClient,
		redisKeyFormat: "game_idle_bff:websocket:ticket:{character_id:%d}",
	}
}

func (r *WebSocketTicketRepo) Save(ctx context.Context, characterID int64, ticket string, ttl time.Duration) error {
	return r.redisClient.Client.Set(ctx, r.redisKey(characterID), ticket, ttl).Err()
}

func (r *WebSocketTicketRepo) Get(ctx context.Context, characterID int64) (string, time.Duration, error) {
	key := r.redisKey(characterID)
	value, err := r.redisClient.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", 0, nil
	}
	if err != nil {
		return "", 0, err
	}
	ttl, err := r.redisClient.Client.TTL(ctx, key).Result()
	if err != nil {
		return "", 0, err
	}
	return value, ttl, nil
}

func (r *WebSocketTicketRepo) redisKey(characterID int64) string {
	return fmt.Sprintf(r.redisKeyFormat, characterID)
}
