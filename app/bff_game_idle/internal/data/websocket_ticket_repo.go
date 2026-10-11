package data

import (
	commonclient "common/pkg/client"
	"context"
	"errors"
	"fmt"
	"time"

	"bff_game_idle/internal/biz/repo"

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
		redisKeyFormat: "bff_game_idle:websocket:ticket:{character_id:%d}",
	}
}

func (r *WebSocketTicketRepo) Save(ctx context.Context, characterID int64, ticket string, ttl time.Duration) error {
	return r.redisClient.Client.Set(ctx, fmt.Sprintf(r.redisKeyFormat, characterID), ticket, ttl).Err()
}

func (r *WebSocketTicketRepo) Consume(ctx context.Context, characterID int64, ticket string) (bool, error) {
	value, err := r.redisClient.Client.GetDel(ctx, fmt.Sprintf(r.redisKeyFormat, characterID)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, err
	}
	return value == ticket, nil
}
