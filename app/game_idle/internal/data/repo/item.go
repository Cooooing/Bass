package repo

import (
	"common/pkg/apperror"
	commonclient "common/pkg/client"
	cerrors "common/proto/gen/common/errors"
	"context"
	"encoding/json"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	itement "game_idle/internal/data/gen/item"
	"game_idle/internal/enum"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.ItemRepo = (*ItemRepo)(nil)

type ItemRepo struct {
	mutex          sync.RWMutex
	db             *gen.Client
	redisClient    *commonclient.RedisClient
	items          map[string]*model.Item
	loaded         bool
	itemsRedisKey  string
	loadedRedisKey string
	redisTTL       time.Duration
}

func NewItemRepo(db *gen.Client, redisClient *commonclient.RedisClient) (bizrepo.ItemRepo, error) {
	repo := &ItemRepo{
		db:             db,
		redisClient:    redisClient,
		items:          make(map[string]*model.Item),
		itemsRedisKey:  "game_idle:metadata:items",
		loadedRedisKey: "game_idle:metadata:items:loaded",
		redisTTL:       2 * time.Hour,
	}
	if err := repo.RefreshLocal(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *ItemRepo) Refresh(ctx context.Context) ([]*model.Item, error) {
	items, itemMap, err := r.loadFromDB(ctx)
	if err != nil {
		return nil, err
	}
	if err = r.saveRedis(ctx, itemMap); err != nil {
		return nil, err
	}
	r.replaceLocal(itemMap)
	return items, nil
}

func (r *ItemRepo) RefreshLocal(ctx context.Context) error {
	itemMap, err := r.loadFromRedis(ctx)
	if err == nil {
		r.replaceLocal(itemMap)
		return nil
	}
	_, itemMap, err = r.loadFromDB(ctx)
	if err != nil {
		return err
	}
	r.replaceLocal(itemMap)
	_ = r.saveRedis(ctx, itemMap)
	return nil
}

func (r *ItemRepo) loadFromDB(ctx context.Context) ([]*model.Item, map[string]*model.Item, error) {
	rows, err := r.db.Item.Query().
		Where(itement.DeletedAtIsNil()).
		Order(itement.BySort(), itement.ByID()).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	items := make([]*model.Item, 0, len(rows))
	itemMap := make(map[string]*model.Item, len(rows))
	for _, row := range rows {
		item := &model.Item{
			ID:          row.ID,
			Name:        row.Name,
			Type:        enum.ItemType(row.Type),
			Description: row.Description,
			Enabled:     row.Enabled,
			Sort:        row.Sort,
		}
		items = append(items, item)
		itemMap[item.ID] = item
	}
	return items, itemMap, nil
}

func (r *ItemRepo) loadFromRedis(ctx context.Context) (map[string]*model.Item, error) {
	loaded, err := r.redisClient.Client.Exists(ctx, r.loadedRedisKey).Result()
	if err != nil {
		return nil, err
	}
	if loaded == 0 {
		return nil, redis.Nil
	}
	values, err := r.redisClient.Client.HGetAll(ctx, r.itemsRedisKey).Result()
	if err != nil {
		return nil, err
	}
	items := make(map[string]*model.Item, len(values))
	for itemID, text := range values {
		item := &model.Item{}
		if err = json.Unmarshal([]byte(text), item); err != nil {
			return nil, err
		}
		items[itemID] = item
	}
	return items, nil
}

func (r *ItemRepo) saveRedis(ctx context.Context, itemMap map[string]*model.Item) error {
	values := make(map[string]any, len(itemMap))
	for itemID, item := range itemMap {
		data, err := json.Marshal(item)
		if err != nil {
			return err
		}
		values[itemID] = data
	}
	existing, err := r.redisClient.Client.HKeys(ctx, r.itemsRedisKey).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	stale := make([]string, 0)
	for _, itemID := range existing {
		if _, ok := itemMap[itemID]; !ok {
			stale = append(stale, itemID)
		}
	}
	_, err = r.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		if len(values) > 0 {
			pipe.HSet(ctx, r.itemsRedisKey, values)
		}
		if len(stale) > 0 {
			pipe.HDel(ctx, r.itemsRedisKey, stale...)
		}
		pipe.Expire(ctx, r.itemsRedisKey, r.redisTTL)
		pipe.Set(ctx, r.loadedRedisKey, "1", r.redisTTL)
		return nil
	})
	return err
}

func (r *ItemRepo) replaceLocal(itemMap map[string]*model.Item) {
	r.mutex.Lock()
	r.items = itemMap
	r.loaded = true
	r.mutex.Unlock()
}

func (r *ItemRepo) Get(ctx context.Context, itemID string) (*model.Item, error) {
	items, err := r.Map(ctx, []string{itemID})
	if err != nil {
		return nil, err
	}
	item, ok := items[itemID]
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ITEM_INVALID)
	}
	return item, nil
}

func (r *ItemRepo) Map(ctx context.Context, itemIDs []string) (map[string]*model.Item, error) {
	r.mutex.RLock()
	loaded := r.loaded
	r.mutex.RUnlock()
	if !loaded {
		if err := r.RefreshLocal(ctx); err != nil {
			return nil, err
		}
	}
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	items := make(map[string]*model.Item)
	if len(itemIDs) == 0 {
		for itemID, item := range r.items {
			items[itemID] = item
		}
		return items, nil
	}
	for _, itemID := range itemIDs {
		if item := r.items[itemID]; item != nil {
			items[itemID] = item
		}
	}
	return items, nil
}
