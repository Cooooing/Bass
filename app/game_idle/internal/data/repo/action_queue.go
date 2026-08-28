package repo

import (
	commonclient "common/pkg/client"
	"context"
	"encoding/json"
	"fmt"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	characteractionqueueent "game_idle/internal/data/gen/characteractionqueue"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.ActionQueueRepo = (*ActionQueueRepo)(nil)

type ActionQueueRepo struct {
	db                         *gen.Client
	redisClient                *commonclient.RedisClient
	queueRedisKeyFormat        string
	queueRedisKeyPattern       string
	queueRedisKeyCharacterHead string
}

func NewActionQueueRepo(db *gen.Client, redisClient *commonclient.RedisClient) bizrepo.ActionQueueRepo {
	return &ActionQueueRepo{
		db:                         db,
		redisClient:                redisClient,
		queueRedisKeyFormat:        "game_idle:action_queue:{character_id:%d}",
		queueRedisKeyPattern:       "game_idle:action_queue:{character_id:*}",
		queueRedisKeyCharacterHead: "game_idle:action_queue:{character_id:",
	}
}

func (r *ActionQueueRepo) ListCharacterIDs(ctx context.Context) ([]int64, error) {
	characterIDs := make([]int64, 0)
	iter := r.redisClient.Client.Scan(ctx, 0, r.queueRedisKeyPattern, 0).Iterator()
	for iter.Next(ctx) {
		characterID, err := strconv.ParseInt(
			strings.TrimSuffix(strings.TrimPrefix(iter.Val(), r.queueRedisKeyCharacterHead), "}"),
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
		return r.loadFromDB(ctx, characterID)
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
	return r.saveRedis(ctx, queue)
}

func (r *ActionQueueRepo) Persist(ctx context.Context, characterID int64) error {
	queue, err := r.Load(ctx, characterID)
	if err != nil {
		return err
	}
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}
	if _, err = tx.CharacterActionQueue.Delete().
		Where(characteractionqueueent.CharacterIDEQ(characterID)).
		Exec(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}
	creates := make([]*gen.CharacterActionQueueCreate, 0, len(queue.Items))
	for index, item := range queue.Items {
		creates = append(creates, tx.CharacterActionQueue.Create().
			SetCharacterID(characterID).
			SetQueueItemID(item.ID).
			SetActionID(item.ActionID).
			SetTimes(item.Times).
			SetPosition(int32(index)).
			SetQueuedAt(item.CreatedAt))
	}
	if len(creates) > 0 {
		if err = tx.CharacterActionQueue.CreateBulk(creates...).Exec(ctx); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *ActionQueueRepo) Clear(ctx context.Context, characterID int64) error {
	return r.redisClient.Client.Del(ctx, r.queueRedisKey(characterID)).Err()
}

func (r *ActionQueueRepo) loadFromDB(ctx context.Context, characterID int64) (*model.ActionQueue, error) {
	rows, err := r.db.CharacterActionQueue.Query().
		Where(characteractionqueueent.CharacterIDEQ(characterID)).
		Order(characteractionqueueent.ByPosition(), characteractionqueueent.ByID()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	queue := &model.ActionQueue{
		CharacterID: characterID,
		Items:       make([]*model.ActionQueueItem, 0, len(rows)),
	}
	for _, row := range rows {
		queue.Items = append(queue.Items, &model.ActionQueueItem{
			ID:        row.QueueItemID,
			ActionID:  row.ActionID,
			Times:     row.Times,
			CreatedAt: row.QueuedAt,
		})
	}
	if err = r.saveRedis(ctx, queue); err != nil {
		return nil, err
	}
	return queue, nil
}

func (r *ActionQueueRepo) saveRedis(ctx context.Context, queue *model.ActionQueue) error {
	data, err := json.Marshal(queue)
	if err != nil {
		return err
	}
	return r.redisClient.Client.Set(ctx, r.queueRedisKey(queue.CharacterID), data, 0).Err()
}

func (r *ActionQueueRepo) queueRedisKey(characterID int64) string {
	return fmt.Sprintf(r.queueRedisKeyFormat, characterID)
}
