package repo

import (
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/config"
	"game_idle/internal/data/gen"
	characteritement "game_idle/internal/data/gen/characteritem"
)

var _ bizrepo.CharacterBackpackRepo = (*CharacterBackpackRepo)(nil)

type CharacterBackpackRepo struct {
	db               *gen.Client
	cache            *CharacterStateCache
	persistThreshold int64
}

func NewCharacterBackpackRepo(
	conf *config.Bootstrap,
	db *gen.Client,
	cache *CharacterStateCache,
) (bizrepo.CharacterBackpackRepo, error) {
	persistThreshold := int64(100)
	if conf.GetGameIdle().GetBackpack().GetPersistThreshold() > 0 {
		persistThreshold = int64(conf.GetGameIdle().GetBackpack().GetPersistThreshold())
	}
	return &CharacterBackpackRepo{
		db:               db,
		cache:            cache,
		persistThreshold: persistThreshold,
	}, nil
}

func (r *CharacterBackpackRepo) LoadItems(ctx context.Context, characterID int64) error {
	character := r.cache.character(characterID)
	character.mutex.RLock()
	backpack := character.backpack
	character.mutex.RUnlock()
	if backpack != nil {
		return nil
	}

	rows, err := r.db.CharacterItem.Query().
		Where(characteritement.CharacterIDEQ(characterID)).
		All(ctx)
	if err != nil {
		return err
	}
	items := make(map[string]*model.CharacterItem, len(rows))
	for _, row := range rows {
		items[row.ItemID] = &model.CharacterItem{
			ID:            row.ID,
			CharacterID:   row.CharacterID,
			ItemID:        row.ItemID,
			Quantity:      row.Quantity,
			TotalObtained: row.TotalObtained,
			TotalConsumed: row.TotalConsumed,
			CreatedAt:     row.CreatedAt,
			UpdatedAt:     row.UpdatedAt,
		}
	}
	character.mutex.Lock()
	if character.backpack == nil {
		character.backpack = &cachedBackpack{
			items:          items,
			operationCount: 0,
		}
	}
	character.mutex.Unlock()
	return nil
}

func (r *CharacterBackpackRepo) MapItems(ctx context.Context, req *bizrepo.BackpackMapReq) (map[string]*model.CharacterItem, error) {
	if err := r.LoadItems(ctx, req.CharacterID); err != nil {
		return nil, err
	}
	character := r.cache.character(req.CharacterID)
	character.mutex.RLock()
	defer character.mutex.RUnlock()
	source := character.backpack.items
	items := make(map[string]*model.CharacterItem)
	if len(req.ItemIDs) == 0 {
		for itemID, item := range source {
			if item.Quantity > 0 {
				copyItem := *item
				items[itemID] = &copyItem
			}
		}
		return items, nil
	}
	for _, itemID := range req.ItemIDs {
		item := source[itemID]
		if item != nil && item.Quantity > 0 {
			copyItem := *item
			items[itemID] = &copyItem
		}
	}
	return items, nil
}

func (r *CharacterBackpackRepo) PersistItems(ctx context.Context, characterID int64) error {
	if err := r.LoadItems(ctx, characterID); err != nil {
		return err
	}
	character := r.cache.character(characterID)
	character.mutex.RLock()
	source := character.backpack.items
	items := make([]*model.CharacterItem, 0, len(source))
	for _, item := range source {
		copyItem := *item
		items = append(items, &copyItem)
	}
	character.mutex.RUnlock()

	creates := make([]*gen.CharacterItemCreate, 0, len(items))
	for _, item := range items {
		creates = append(creates, r.db.CharacterItem.Create().
			SetCharacterID(characterID).
			SetItemID(item.ItemID).
			SetQuantity(item.Quantity).
			SetTotalObtained(item.TotalObtained).
			SetTotalConsumed(item.TotalConsumed))
	}
	if len(creates) > 0 {
		if err := r.db.CharacterItem.CreateBulk(creates...).
			OnConflictColumns(characteritement.FieldCharacterID, characteritement.FieldItemID).
			UpdateQuantity().
			UpdateTotalObtained().
			UpdateTotalConsumed().
			UpdateUpdatedAt().
			Exec(ctx); err != nil {
			return err
		}
	}
	character.mutex.Lock()
	character.backpack.operationCount = 0
	character.mutex.Unlock()
	return nil
}

func (r *CharacterBackpackRepo) CheckItems(ctx context.Context, req *bizrepo.BackpackCheckReq) (bool, error) {
	if err := r.LoadItems(ctx, req.CharacterID); err != nil {
		return false, err
	}
	character := r.cache.character(req.CharacterID)
	character.mutex.RLock()
	defer character.mutex.RUnlock()
	items := character.backpack.items
	for itemID, quantity := range req.Items {
		item := items[itemID]
		if item == nil || item.Quantity < quantity {
			return false, nil
		}
	}
	return true, nil
}
