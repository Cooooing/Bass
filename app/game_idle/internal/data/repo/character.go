package repo

import (
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	characterent "game_idle/internal/data/gen/character"
	"game_idle/internal/enum"
	"sync"
	"time"
)

var _ bizrepo.CharacterRepo = (*CharacterRepo)(nil)

type CharacterRepo struct {
	db         *gen.Client
	mutex      sync.RWMutex
	characters map[int64]*model.Character
}

func NewCharacterRepo(db *gen.Client) bizrepo.CharacterRepo {
	return &CharacterRepo{
		db:         db,
		characters: make(map[int64]*model.Character),
	}
}

func (r *CharacterRepo) Save(ctx context.Context, character *model.Character) (*model.Character, error) {
	status := character.Status
	if status == "" {
		status = enum.CharacterStatusActive
	}
	row, err := r.db.Character.Create().
		SetUserID(character.UserID).
		SetSlot(character.Slot).
		SetName(character.Name).
		SetNameKey(character.NameKey).
		SetActionQueueCapacity(character.ActionQueueCapacity).
		SetMaxOfflineSeconds(int64(character.MaxOfflineDuration / time.Second)).
		SetStatus(characterent.Status(status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	character = &model.Character{
		ID:                  row.ID,
		UserID:              row.UserID,
		Slot:                row.Slot,
		Name:                row.Name,
		NameKey:             row.NameKey,
		ActionQueueCapacity: row.ActionQueueCapacity,
		MaxOfflineDuration:  time.Duration(row.MaxOfflineSeconds) * time.Second,
		Status:              enum.CharacterStatus(row.Status),
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
		LastOfflineAt:       row.LastOfflineAt,
		DeletedAt:           row.DeletedAt,
	}
	r.mutex.Lock()
	cachedCharacter := *character
	r.characters[character.ID] = &cachedCharacter
	r.mutex.Unlock()
	return character, nil
}

func (r *CharacterRepo) Get(ctx context.Context, characterID int64) (*model.Character, error) {
	r.mutex.RLock()
	if cachedCharacter := r.characters[characterID]; cachedCharacter != nil {
		character := *cachedCharacter
		r.mutex.RUnlock()
		return &character, nil
	}
	r.mutex.RUnlock()
	row, err := r.db.Character.Query().
		Where(characterent.IDEQ(characterID), characterent.DeletedAtIsNil()).
		First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	character := &model.Character{
		ID:                  row.ID,
		UserID:              row.UserID,
		Slot:                row.Slot,
		Name:                row.Name,
		NameKey:             row.NameKey,
		ActionQueueCapacity: row.ActionQueueCapacity,
		MaxOfflineDuration:  time.Duration(row.MaxOfflineSeconds) * time.Second,
		Status:              enum.CharacterStatus(row.Status),
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
		LastOfflineAt:       row.LastOfflineAt,
		DeletedAt:           row.DeletedAt,
	}
	r.mutex.Lock()
	cachedCharacter := *character
	r.characters[character.ID] = &cachedCharacter
	r.mutex.Unlock()
	return character, nil
}

func (r *CharacterRepo) GetName(ctx context.Context, characterID int64) (string, error) {
	r.mutex.RLock()
	character := r.characters[characterID]
	r.mutex.RUnlock()
	if character != nil {
		return character.Name, nil
	}
	character, err := r.Get(ctx, characterID)
	if err != nil {
		return "", err
	}
	if character == nil {
		return "", nil
	}
	return character.Name, nil
}

func (r *CharacterRepo) List(ctx context.Context, req *bizrepo.ListCharacterReq) ([]*model.Character, error) {
	query := r.db.Character.Query().
		Where(characterent.DeletedAtIsNil())
	if req.UserID != nil {
		query = query.Where(characterent.UserIDEQ(*req.UserID))
	}
	if req.CharacterID != nil && *req.CharacterID > 0 {
		query = query.Where(characterent.IDEQ(*req.CharacterID))
	}
	if req.NameKey != nil && *req.NameKey != "" {
		query = query.Where(characterent.NameKeyEQ(*req.NameKey))
	}
	rows, err := query.
		Order(characterent.BySlot()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	characters := make([]*model.Character, 0, len(rows))
	for _, row := range rows {
		character := &model.Character{
			ID:                  row.ID,
			UserID:              row.UserID,
			Slot:                row.Slot,
			Name:                row.Name,
			NameKey:             row.NameKey,
			ActionQueueCapacity: row.ActionQueueCapacity,
			MaxOfflineDuration:  time.Duration(row.MaxOfflineSeconds) * time.Second,
			Status:              enum.CharacterStatus(row.Status),
			CreatedAt:           row.CreatedAt,
			UpdatedAt:           row.UpdatedAt,
			LastOfflineAt:       row.LastOfflineAt,
			DeletedAt:           row.DeletedAt,
		}
		characters = append(characters, character)
		r.mutex.Lock()
		cachedCharacter := *character
		r.characters[character.ID] = &cachedCharacter
		r.mutex.Unlock()
	}
	return characters, nil
}

func (r *CharacterRepo) UpdateLastOfflineAt(ctx context.Context, characterID int64, at time.Time) (bool, error) {
	affected, err := r.db.Character.Update().
		Where(characterent.IDEQ(characterID), characterent.DeletedAtIsNil()).
		SetLastOfflineAt(at).
		Save(ctx)
	if err != nil {
		return false, err
	}
	if affected == 0 {
		return false, nil
	}
	r.mutex.Lock()
	if character := r.characters[characterID]; character != nil {
		offlineAt := at
		character.LastOfflineAt = &offlineAt
	}
	r.mutex.Unlock()
	return true, nil
}
