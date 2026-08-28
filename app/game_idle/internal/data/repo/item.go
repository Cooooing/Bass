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

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.ItemRepo = (*ItemRepo)(nil)

type ItemRepo struct {
	db             *gen.Client
	redisClient    *commonclient.RedisClient
	itemsRedisKey  string
	loadedRedisKey string
}

func NewItemRepo(db *gen.Client, redisClient *commonclient.RedisClient) (bizrepo.ItemRepo, error) {
	repo := &ItemRepo{
		db:             db,
		redisClient:    redisClient,
		itemsRedisKey:  "game_idle:config:items",
		loadedRedisKey: "game_idle:config:items:loaded",
	}
	if _, err := repo.Refresh(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *ItemRepo) Refresh(ctx context.Context) ([]*model.Item, error) {
	rows, err := r.db.Item.Query().
		Where(itement.DeletedAtIsNil()).
		Order(itement.BySort(), itement.ByID()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*model.Item, 0, len(rows))
	values := make(map[string]any, len(rows))
	for _, row := range rows {
		item := &model.Item{
			ID:          row.ID,
			Name:        row.Name,
			Type:        enum.ItemType(row.Type),
			Description: row.Description,
			Enabled:     row.Enabled,
			Sort:        row.Sort,
		}
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		values[item.ID] = data
	}
	_, err = r.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, r.itemsRedisKey)
		if len(values) > 0 {
			pipe.HSet(ctx, r.itemsRedisKey, values)
		}
		pipe.Set(ctx, r.loadedRedisKey, "1", 0)
		return nil
	})
	return items, err
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
	loaded, err := r.redisClient.Client.Exists(ctx, r.loadedRedisKey).Result()
	if err != nil {
		return nil, err
	}
	if loaded == 0 {
		if _, err := r.Refresh(ctx); err != nil {
			return nil, err
		}
	}
	values := map[string]string{}
	if len(itemIDs) == 0 {
		values, err = r.redisClient.Client.HGetAll(ctx, r.itemsRedisKey).Result()
	} else {
		results, err := r.redisClient.Client.HMGet(ctx, r.itemsRedisKey, itemIDs...).Result()
		if err != nil {
			return nil, err
		}
		for index, result := range results {
			if text, ok := result.(string); ok {
				values[itemIDs[index]] = text
			}
		}
	}
	if err != nil {
		return nil, err
	}
	items := make(map[string]*model.Item, len(values))
	for itemID, text := range values {
		item := &model.Item{}
		if err := json.Unmarshal([]byte(text), item); err != nil {
			return nil, err
		}
		items[itemID] = item
	}
	return items, nil
}
