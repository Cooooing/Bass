package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	characterent "game_idle/internal/data/gen/character"
	"game_idle/internal/enum"
	"strings"
	"sync"
	"time"
)

var _ bizrepo.CharacterRepo = (*CharacterRepo)(nil)

type CharacterRepo struct {
	db         *gen.Client
	mutex      sync.RWMutex
	characters map[int64]*model.Character
	names      map[int64]string
}

func NewCharacterRepo(db *gen.Client) bizrepo.CharacterRepo {
	return &CharacterRepo{
		db:         db,
		characters: make(map[int64]*model.Character),
		names:      make(map[int64]string),
	}
}

func (r *CharacterRepo) Save(ctx context.Context, character *model.Character) (*model.Character, error) {
	if character == nil || character.UserID <= 0 || character.Name == "" || character.NameKey == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	if character.Slot <= 0 || character.ActionQueueCapacity <= 0 || character.MaxOfflineDuration <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
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
	if gen.IsConstraintError(err) && strings.Contains(err.Error(), "game_idle_characters_user_slot_active_unique") {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_LIMIT_EXCEEDED)
	}
	if gen.IsConstraintError(err) && strings.Contains(err.Error(), "game_idle_characters_name_key_active_unique") {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NAME_TAKEN)
	}
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
		DeletedAt:           row.DeletedAt,
	}
	r.cache(character)
	return character, nil
}

func (r *CharacterRepo) Get(ctx context.Context, characterID int64) (*model.Character, error) {
	r.mutex.RLock()
	character := r.characters[characterID]
	r.mutex.RUnlock()
	if character != nil {
		return character, nil
	}
	row, err := r.db.Character.Query().
		Where(characterent.IDEQ(characterID), characterent.DeletedAtIsNil()).
		First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NOT_FOUND)
	}
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
		DeletedAt:           row.DeletedAt,
	}
	r.cache(character)
	return character, nil
}

func (r *CharacterRepo) GetName(ctx context.Context, characterID int64) (string, error) {
	r.mutex.RLock()
	name := r.names[characterID]
	r.mutex.RUnlock()
	if name != "" {
		return name, nil
	}
	character, err := r.Get(ctx, characterID)
	if err != nil {
		return "", err
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
			DeletedAt:           row.DeletedAt,
		}
		characters = append(characters, character)
		r.cache(character)
	}
	return characters, nil
}

func (r *CharacterRepo) cache(character *model.Character) {
	r.mutex.Lock()
	r.characters[character.ID] = character
	r.names[character.ID] = character.Name
	r.mutex.Unlock()
}
