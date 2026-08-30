package repo

import (
	commonclient "common/pkg/client"
	"context"
	"fmt"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	characteractionqueueent "game_idle/internal/data/gen/characteractionqueue"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.ActionQueueRepo = (*ActionQueueRepo)(nil)

type ActionQueueRepo struct {
	db          *gen.Client
	redisClient *commonclient.RedisClient
	keys        actionQueueRedisKeys
}

type actionQueueRedisKeys struct {
	characters           string
	queueItemIDsFormat   string
	queueItemsFormat     string
	queueItemFieldJoiner string
	loadedField          string
}

func NewActionQueueRepo(db *gen.Client, redisClient *commonclient.RedisClient) bizrepo.ActionQueueRepo {
	return &ActionQueueRepo{
		db:          db,
		redisClient: redisClient,
		keys: actionQueueRedisKeys{
			characters:           "game_idle:action_queue:characters",
			queueItemIDsFormat:   "game_idle:character:{character_id:%d}:action_queue:item_ids",
			queueItemsFormat:     "game_idle:character:{character_id:%d}:action_queue:items",
			queueItemFieldJoiner: ":",
			loadedField:          "__loaded",
		},
	}
}

func (r *ActionQueueRepo) ListCharacterIDs(ctx context.Context) ([]int64, error) {
	seen := map[int64]struct{}{}
	characterIDValues, err := r.redisClient.Client.SMembers(ctx, r.keys.characters).Result()
	if err != nil {
		return nil, err
	}
	for _, value := range characterIDValues {
		characterID, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		seen[characterID] = struct{}{}
	}
	var rows []struct {
		CharacterID int64 `json:"character_id"`
	}
	if err := r.db.CharacterActionQueue.Query().
		Unique(true).
		Select(characteractionqueueent.FieldCharacterID).
		Scan(ctx, &rows); err != nil {
		return nil, err
	}
	characterIDs := make([]int64, 0, len(seen)+len(rows))
	for characterID := range seen {
		characterIDs = append(characterIDs, characterID)
	}
	for _, row := range rows {
		if _, ok := seen[row.CharacterID]; ok {
			continue
		}
		characterIDs = append(characterIDs, row.CharacterID)
	}
	return characterIDs, nil
}

func (r *ActionQueueRepo) Load(ctx context.Context, characterID int64) (*model.ActionQueue, error) {
	itemIDs, err := r.redisClient.Client.LRange(ctx, r.queueItemIDsRedisKey(characterID), 0, -1).Result()
	if err != nil {
		return nil, err
	}
	if len(itemIDs) == 0 {
		loaded, err := r.isRedisLoaded(ctx, characterID)
		if err != nil || loaded {
			return r.emptyQueue(characterID), err
		}
		return r.loadFromDB(ctx, characterID)
	}
	return r.loadFromRedis(ctx, characterID, itemIDs)
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
	_, err := r.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, r.queueItemIDsRedisKey(characterID), r.queueItemsRedisKey(characterID))
		pipe.SRem(ctx, r.keys.characters, characterID)
		return nil
	})
	return err
}

func (r *ActionQueueRepo) loadFromDB(ctx context.Context, characterID int64) (*model.ActionQueue, error) {
	rows, err := r.db.CharacterActionQueue.Query().
		Where(characteractionqueueent.CharacterIDEQ(characterID)).
		Order(characteractionqueueent.ByPosition(), characteractionqueueent.ByID()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	queue := r.emptyQueue(characterID)
	queue.Items = make([]*model.ActionQueueItem, 0, len(rows))
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
	_, err := r.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, r.queueItemIDsRedisKey(queue.CharacterID), r.queueItemsRedisKey(queue.CharacterID))
		pipe.SAdd(ctx, r.keys.characters, queue.CharacterID)
		pipe.HSet(ctx, r.queueItemsRedisKey(queue.CharacterID), r.keys.loadedField, "1")
		for _, item := range queue.Items {
			pipe.RPush(ctx, r.queueItemIDsRedisKey(queue.CharacterID), item.ID)
			pipe.HSet(ctx, r.queueItemsRedisKey(queue.CharacterID), r.queueItemFields(item))
		}
		return nil
	})
	return err
}

func (r *ActionQueueRepo) isRedisLoaded(ctx context.Context, characterID int64) (bool, error) {
	return r.redisClient.Client.HExists(ctx, r.queueItemsRedisKey(characterID), r.keys.loadedField).Result()
}

func (r *ActionQueueRepo) emptyQueue(characterID int64) *model.ActionQueue {
	return &model.ActionQueue{
		CharacterID: characterID,
		Items:       make([]*model.ActionQueueItem, 0),
	}
}

func (r *ActionQueueRepo) queueItemFields(item *model.ActionQueueItem) map[string]any {
	return map[string]any{
		r.queueItemField(item.ID, "action_id"):           item.ActionID,
		r.queueItemField(item.ID, "times"):               item.Times,
		r.queueItemField(item.ID, "queued_at_unix_nano"): item.CreatedAt.UnixNano(),
	}
}

func (r *ActionQueueRepo) loadFromRedis(
	ctx context.Context,
	characterID int64,
	itemIDs []string,
) (*model.ActionQueue, error) {
	fields := make([]string, 0, len(itemIDs)*3)
	for _, itemID := range itemIDs {
		fields = append(
			fields,
			r.queueItemField(itemID, "action_id"),
			r.queueItemField(itemID, "times"),
			r.queueItemField(itemID, "queued_at_unix_nano"),
		)
	}
	values, err := r.redisClient.Client.HMGet(ctx, r.queueItemsRedisKey(characterID), fields...).Result()
	if err != nil {
		return nil, err
	}
	queue := r.emptyQueue(characterID)
	queue.Items = make([]*model.ActionQueueItem, 0, len(itemIDs))
	for index, itemID := range itemIDs {
		valueIndex := index * 3
		if values[valueIndex] == nil || values[valueIndex+1] == nil || values[valueIndex+2] == nil {
			continue
		}
		times, err := strconv.ParseInt(fmt.Sprint(values[valueIndex+1]), 10, 64)
		if err != nil {
			return nil, err
		}
		queuedAtUnixNano, err := strconv.ParseInt(fmt.Sprint(values[valueIndex+2]), 10, 64)
		if err != nil {
			return nil, err
		}
		queue.Items = append(queue.Items, &model.ActionQueueItem{
			ID:        itemID,
			ActionID:  fmt.Sprint(values[valueIndex]),
			Times:     times,
			CreatedAt: time.Unix(0, queuedAtUnixNano),
		})
	}
	if len(queue.Items) != len(itemIDs) {
		if err = r.saveRedis(ctx, queue); err != nil {
			return nil, err
		}
	}
	return queue, nil
}

func (r *ActionQueueRepo) queueItemIDsRedisKey(characterID int64) string {
	return fmt.Sprintf(r.keys.queueItemIDsFormat, characterID)
}

func (r *ActionQueueRepo) queueItemsRedisKey(characterID int64) string {
	return fmt.Sprintf(r.keys.queueItemsFormat, characterID)
}

func (r *ActionQueueRepo) queueItemField(queueItemID string, field string) string {
	return queueItemID + r.keys.queueItemFieldJoiner + field
}
