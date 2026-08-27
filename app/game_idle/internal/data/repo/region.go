package repo

import (
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	regionent "game_idle/internal/data/gen/region"
	"game_idle/internal/enum"
	"sync"
)

var _ bizrepo.RegionRepo = (*RegionRepo)(nil)

type RegionRepo struct {
	mutex   sync.RWMutex
	db      *gen.Client
	regions map[string]*model.Region
	loaded  bool
}

func NewRegionRepo(db *gen.Client) (bizrepo.RegionRepo, error) {
	repo := &RegionRepo{
		db:      db,
		regions: make(map[string]*model.Region),
	}
	if _, err := repo.Refresh(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *RegionRepo) Refresh(ctx context.Context) ([]*model.Region, error) {
	rows, err := r.db.Region.Query().
		Where(regionent.DeletedAtIsNil()).
		Order(regionent.BySort(), regionent.ByID()).
		All(ctx)
	if err != nil {
		return nil, err
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
	r.mutex.Lock()
	r.regions = regionMap
	r.loaded = true
	r.mutex.Unlock()
	return regions, nil
}

func (r *RegionRepo) Map(ctx context.Context, regionIDs []string) (map[string]*model.Region, error) {
	r.mutex.RLock()
	loaded := r.loaded
	r.mutex.RUnlock()
	if !loaded {
		if _, err := r.Refresh(ctx); err != nil {
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
