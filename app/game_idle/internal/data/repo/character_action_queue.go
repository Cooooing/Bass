package repo

import (
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	characteractionqueueent "game_idle/internal/data/gen/characteractionqueue"
)

var _ bizrepo.CharacterActionQueueRepo = (*CharacterActionQueueRepo)(nil)

type CharacterActionQueueRepo struct {
	db    *gen.Client
	cache *CharacterStateCache
}

func NewCharacterActionQueueRepo(db *gen.Client, cache *CharacterStateCache) bizrepo.CharacterActionQueueRepo {
	return &CharacterActionQueueRepo{
		db:    db,
		cache: cache,
	}
}

func (r *CharacterActionQueueRepo) ListCharacterIDs(ctx context.Context) ([]int64, error) {
	seen := map[int64]struct{}{}
	for characterID, character := range r.cache.charactersSnapshot() {
		character.mutex.RLock()
		if character.queue != nil {
			seen[characterID] = struct{}{}
		}
		character.mutex.RUnlock()
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

func (r *CharacterActionQueueRepo) Load(ctx context.Context, characterID int64) (*model.CharacterActionQueue, error) {
	character := r.cache.character(characterID)
	character.mutex.RLock()
	queueCache := character.queue
	character.mutex.RUnlock()
	if queueCache != nil {
		items := make([]*model.CharacterActionQueueItem, len(queueCache.queue.Items))
		for index, item := range queueCache.queue.Items {
			copyItem := *item
			items[index] = &copyItem
		}
		return &model.CharacterActionQueue{
			CharacterID: queueCache.queue.CharacterID,
			Items:       items,
		}, nil
	}
	rows, err := r.db.CharacterActionQueue.Query().
		Where(characteractionqueueent.CharacterIDEQ(characterID)).
		Order(characteractionqueueent.ByPosition(), characteractionqueueent.ByID()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	queue := &model.CharacterActionQueue{
		CharacterID: characterID,
		Items:       make([]*model.CharacterActionQueueItem, 0, len(rows)),
	}
	for _, row := range rows {
		queue.Items = append(queue.Items, &model.CharacterActionQueueItem{
			ID:        row.QueueItemID,
			ActionID:  row.ActionID,
			Times:     row.Times,
			CreatedAt: row.QueuedAt,
		})
	}
	if err = r.Save(ctx, queue); err != nil {
		return nil, err
	}
	return queue, nil
}

func (r *CharacterActionQueueRepo) Save(ctx context.Context, queue *model.CharacterActionQueue) error {
	character := r.cache.character(queue.CharacterID)
	items := make([]*model.CharacterActionQueueItem, len(queue.Items))
	for index, item := range queue.Items {
		copyItem := *item
		items[index] = &copyItem
	}
	character.mutex.Lock()
	character.queue = &cachedActionQueue{
		queue: &model.CharacterActionQueue{
			CharacterID: queue.CharacterID,
			Items:       items,
		},
	}
	character.mutex.Unlock()
	return nil
}

func (r *CharacterActionQueueRepo) Persist(ctx context.Context, characterID int64) error {
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
