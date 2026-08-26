package repo

import (
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	characterent "game_idle/internal/data/gen/character"
	"game_idle/internal/enum"
	"strings"
	"time"
)

var _ bizrepo.CharacterRepo = (*CharacterRepo)(nil)

type CharacterRepo struct {
	db *gen.Client
}

func NewCharacterRepo(db *gen.Client) bizrepo.CharacterRepo {
	return &CharacterRepo{
		db: db,
	}
}

func (r *CharacterRepo) Save(ctx context.Context, character *model.Character) (*model.Character, error) {
	if character == nil || character.UserID <= 0 || character.Name == "" || character.NameKey == "" {
		return nil, model.ErrCharacterInvalid
	}
	if character.Slot <= 0 || character.ActionQueueCapacity <= 0 || character.MaxOfflineDuration <= 0 {
		return nil, model.ErrCharacterInvalid
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
		return nil, model.ErrCharacterLimitExceeded
	}
	if gen.IsConstraintError(err) && strings.Contains(err.Error(), "game_idle_characters_name_key_active_unique") {
		return nil, model.ErrCharacterNameDuplicate
	}
	if err != nil {
		return nil, err
	}
	return &model.Character{
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
	}, nil
}

func (r *CharacterRepo) Get(ctx context.Context, characterID int64) (*model.Character, error) {
	row, err := r.db.Character.Query().
		Where(characterent.IDEQ(characterID), characterent.DeletedAtIsNil()).
		First(ctx)
	if gen.IsNotFound(err) {
		return nil, model.ErrCharacterNotFound
	}
	if err != nil {
		return nil, err
	}
	return &model.Character{
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
	}, nil
}

func (r *CharacterRepo) ListByUserID(ctx context.Context, userID int64) ([]*model.Character, error) {
	if userID <= 0 {
		return nil, model.ErrCharacterInvalid
	}
	rows, err := r.db.Character.Query().
		Where(characterent.UserIDEQ(userID), characterent.DeletedAtIsNil()).
		Order(characterent.BySlot()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	characters := make([]*model.Character, 0, len(rows))
	for _, row := range rows {
		characters = append(characters, &model.Character{
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
		})
	}
	return characters, nil
}
