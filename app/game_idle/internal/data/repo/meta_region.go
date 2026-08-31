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

var _ bizrepo.MetaRegionRepo = (*MetaRegionRepo)(nil)

type MetaRegionRepo struct {
	mutex   sync.RWMutex
	db      *gen.Client
	regions map[string]*model.MetaRegion
}

func NewMetaRegionRepo(db *gen.Client) (bizrepo.MetaRegionRepo, error) {
	repo := &MetaRegionRepo{
		db: db,
	}
	if err := repo.Refresh(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *MetaRegionRepo) Refresh(ctx context.Context) error {
	rows, err := r.db.Region.Query().
		Where(regionent.DeletedAtIsNil()).
		Order(regionent.BySort(), regionent.ByID()).
		All(ctx)
	if err != nil {
		return err
	}
	regions := make(map[string]*model.MetaRegion, len(rows))
	for _, row := range rows {
		region := &model.MetaRegion{
			ID:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			ActionKind:  enum.ActionKind(row.ActionKind),
			Enabled:     row.Enabled,
			Sort:        row.Sort,
		}
		regions[region.ID] = region
	}
	r.mutex.Lock()
	r.regions = regions
	r.mutex.Unlock()
	return nil
}

func (r *MetaRegionRepo) Map(ctx context.Context, regionIDs []string) (map[string]*model.MetaRegion, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if len(regionIDs) == 0 {
		regions := make(map[string]*model.MetaRegion, len(r.regions))
		for regionID, region := range r.regions {
			regions[regionID] = region
		}
		return regions, nil
	}
	regions := make(map[string]*model.MetaRegion, len(regionIDs))
	for _, regionID := range regionIDs {
		if region := r.regions[regionID]; region != nil {
			regions[regionID] = region
		}
	}
	return regions, nil
}
