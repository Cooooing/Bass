package repo

import (
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	itement "game_idle/internal/data/gen/item"
	"game_idle/internal/enum"
	"sync"
)

var _ bizrepo.MetaItemRepo = (*MetaItemRepo)(nil)

type MetaItemRepo struct {
	mutex sync.RWMutex
	db    *gen.Client
	items map[string]*model.MetaItem
}

func NewMetaItemRepo(db *gen.Client) (bizrepo.MetaItemRepo, error) {
	repo := &MetaItemRepo{
		db: db,
	}
	if err := repo.Refresh(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *MetaItemRepo) Refresh(ctx context.Context) error {
	rows, err := r.db.Item.Query().
		Where(itement.DeletedAtIsNil()).
		Order(itement.BySort(), itement.ByID()).
		All(ctx)
	if err != nil {
		return err
	}
	items := make(map[string]*model.MetaItem, len(rows))
	for _, row := range rows {
		item := &model.MetaItem{
			ID:          row.ID,
			Name:        row.Name,
			Type:        enum.ItemType(row.Type),
			Description: row.Description,
			Enabled:     row.Enabled,
			Sort:        row.Sort,
		}
		items[item.ID] = item
	}
	r.mutex.Lock()
	r.items = items
	r.mutex.Unlock()
	return nil
}

func (r *MetaItemRepo) Get(ctx context.Context, itemID string) (*model.MetaItem, error) {
	r.mutex.RLock()
	item := r.items[itemID]
	r.mutex.RUnlock()
	return item, nil
}

func (r *MetaItemRepo) Map(ctx context.Context, itemIDs []string) (map[string]*model.MetaItem, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if len(itemIDs) == 0 {
		items := make(map[string]*model.MetaItem, len(r.items))
		for itemID, item := range r.items {
			items[itemID] = item
		}
		return items, nil
	}
	items := make(map[string]*model.MetaItem, len(itemIDs))
	for _, itemID := range itemIDs {
		if item := r.items[itemID]; item != nil {
			items[itemID] = item
		}
	}
	return items, nil
}
