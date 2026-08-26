package repo

import (
	commonclient "common/pkg/client"
	"context"
	"encoding/json"
	"fmt"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.ActionQueueRepo = (*ActionQueueRepo)(nil)

type ActionQueueRepo struct {
	redisClient                  *commonclient.RedisClient
	queueRedisKeyFormat          string
	queueRedisKeyPattern         string
	queueRedisKeyCharacterPrefix string
}

func NewActionQueueRepo(redisClient *commonclient.RedisClient) bizrepo.ActionQueueRepo {
	return &ActionQueueRepo{
		redisClient:                  redisClient,
		queueRedisKeyFormat:          "game_idle:character_action_queue:%d",
		queueRedisKeyPattern:         "game_idle:character_action_queue:*",
		queueRedisKeyCharacterPrefix: "game_idle:character_action_queue:",
	}
}

func (r *ActionQueueRepo) ListCharacterIDs(ctx context.Context) ([]int64, error) {
	characterIDs := make([]int64, 0)
	iter := r.redisClient.Client.Scan(ctx, 0, r.queueRedisKeyPattern, 0).Iterator()
	for iter.Next(ctx) {
		characterID, err := strconv.ParseInt(
			strings.TrimPrefix(iter.Val(), r.queueRedisKeyCharacterPrefix),
			10,
			64,
		)
		if err != nil {
			return nil, err
		}
		characterIDs = append(characterIDs, characterID)
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return characterIDs, nil
}

func (r *ActionQueueRepo) Load(ctx context.Context, characterID int64) (*model.ActionQueue, error) {
	data, err := r.redisClient.Client.Get(ctx, r.queueRedisKey(characterID)).Bytes()
	if err == redis.Nil {
		return &model.ActionQueue{
			CharacterID: characterID,
			Items:       make([]*model.ActionQueueItem, 0),
		}, nil
	}
	if err != nil {
		return nil, err
	}
	queue := &model.ActionQueue{}
	if err := json.Unmarshal(data, queue); err != nil {
		return nil, err
	}
	return queue, nil
}

func (r *ActionQueueRepo) Save(ctx context.Context, queue *model.ActionQueue) error {
	data, err := json.Marshal(queue)
	if err != nil {
		return err
	}
	return r.redisClient.Client.Set(ctx, r.queueRedisKey(queue.CharacterID), data, 0).Err()
}

func (r *ActionQueueRepo) queueRedisKey(characterID int64) string {
	return fmt.Sprintf(r.queueRedisKeyFormat, characterID)
}
