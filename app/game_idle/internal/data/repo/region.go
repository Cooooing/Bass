package repo

import (
	commonclient "common/pkg/client"
	"context"
	"encoding/json"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	regionent "game_idle/internal/data/gen/region"
	"game_idle/internal/enum"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.RegionRepo = (*RegionRepo)(nil)

type RegionRepo struct {
	mutex           sync.RWMutex
	db              *gen.Client
	redisClient     *commonclient.RedisClient
	regions         map[string]*model.Region
	loaded          bool
	regionsRedisKey string
	loadedRedisKey  string
	redisTTL        time.Duration
}

func NewRegionRepo(db *gen.Client, redisClient *commonclient.RedisClient) (bizrepo.RegionRepo, error) {
	repo := &RegionRepo{
		db:              db,
		redisClient:     redisClient,
		regions:         make(map[string]*model.Region),
		regionsRedisKey: "game_idle:metadata:regions",
		loadedRedisKey:  "game_idle:metadata:regions:loaded",
		redisTTL:        2 * time.Hour,
	}
	if err := repo.RefreshLocal(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *RegionRepo) Refresh(ctx context.Context) ([]*model.Region, error) {
	regions, regionMap, err := r.loadFromDB(ctx)
	if err != nil {
		return nil, err
	}
	if err = r.saveRedis(ctx, regionMap); err != nil {
		return nil, err
	}
	r.replaceLocal(regionMap)
	return regions, nil
}

func (r *RegionRepo) RefreshLocal(ctx context.Context) error {
	regionMap, err := r.loadFromRedis(ctx)
	if err == nil {
		r.replaceLocal(regionMap)
		return nil
	}
	_, regionMap, err = r.loadFromDB(ctx)
	if err != nil {
		return err
	}
	r.replaceLocal(regionMap)
	_ = r.saveRedis(ctx, regionMap)
	return nil
}

func (r *RegionRepo) loadFromDB(ctx context.Context) ([]*model.Region, map[string]*model.Region, error) {
	rows, err := r.db.Region.Query().
		Where(regionent.DeletedAtIsNil()).
		Order(regionent.BySort(), regionent.ByID()).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	regions := make([]*model.Region, 0, len(rows))
	regionMap := make(map[string]*model.Region, len(rows))
	for _, row := range rows {
		region := &model.Region{
			ID:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			ActionKind:  enum.ActionKind(row.ActionKind),
			Enabled:     row.Enabled,
			Sort:        row.Sort,
		}
		regions = append(regions, region)
		regionMap[region.ID] = region
	}
	return regions, regionMap, nil
}

func (r *RegionRepo) loadFromRedis(ctx context.Context) (map[string]*model.Region, error) {
	loaded, err := r.redisClient.Client.Exists(ctx, r.loadedRedisKey).Result()
	if err != nil {
		return nil, err
	}
	if loaded == 0 {
		return nil, redis.Nil
	}
	values, err := r.redisClient.Client.HGetAll(ctx, r.regionsRedisKey).Result()
	if err != nil {
		return nil, err
	}
	regionMap := make(map[string]*model.Region, len(values))
	for regionID, text := range values {
		region := &model.Region{}
		if err = json.Unmarshal([]byte(text), region); err != nil {
			return nil, err
		}
		regionMap[regionID] = region
	}
	return regionMap, nil
}

func (r *RegionRepo) saveRedis(ctx context.Context, regionMap map[string]*model.Region) error {
	values := make(map[string]any, len(regionMap))
	for regionID, region := range regionMap {
		data, err := json.Marshal(region)
		if err != nil {
			return err
		}
		values[regionID] = data
	}
	existing, err := r.redisClient.Client.HKeys(ctx, r.regionsRedisKey).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	stale := make([]string, 0)
	for _, regionID := range existing {
		if _, ok := regionMap[regionID]; !ok {
			stale = append(stale, regionID)
		}
	}
	_, err = r.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		if len(values) > 0 {
			pipe.HSet(ctx, r.regionsRedisKey, values)
		}
		if len(stale) > 0 {
			pipe.HDel(ctx, r.regionsRedisKey, stale...)
		}
		pipe.Expire(ctx, r.regionsRedisKey, r.redisTTL)
		pipe.Set(ctx, r.loadedRedisKey, "1", r.redisTTL)
		return nil
	})
	return err
}

func (r *RegionRepo) replaceLocal(regionMap map[string]*model.Region) {
	r.mutex.Lock()
	r.regions = regionMap
	r.loaded = true
	r.mutex.Unlock()
}

func (r *RegionRepo) Map(ctx context.Context, regionIDs []string) (map[string]*model.Region, error) {
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
	regions := make(map[string]*model.Region)
	if len(regionIDs) == 0 {
		for regionID, region := range r.regions {
			regions[regionID] = region
		}
		return regions, nil
	}
	for _, regionID := range regionIDs {
		if region, ok := r.regions[regionID]; ok {
			regions[regionID] = region
		}
	}
	return regions, nil
}
