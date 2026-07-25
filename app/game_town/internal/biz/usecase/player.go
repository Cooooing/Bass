package usecase

import (
	"context"
	"strings"

	"common/pkg/apperror"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

type PlayerUsecase struct {
	playerRepo repo.PlayerRepo
}

func NewPlayerUsecase(
	playerRepo repo.PlayerRepo,
) *PlayerUsecase {
	return &PlayerUsecase{
		playerRepo: playerRepo,
	}
}

type RegisterPlayerReq struct {
	Name        string
	DisplayName string
}

func (u *PlayerUsecase) Register(ctx context.Context, req *RegisterPlayerReq) (*model.Player, error) {
	name := strings.ToLower(strings.TrimSpace(req.Name))
	displayName := strings.TrimSpace(req.DisplayName)
	if name == "" || len(name) > 64 {
		return nil, apperror.CommonInvalidArgument()
	}
	if displayName == "" {
		displayName = name
	}
	return u.playerRepo.Save(ctx, &model.Player{
		Name:        name,
		DisplayName: displayName,
		Status:      enum.PlayerStatusActive,
	})
}

func (u *PlayerUsecase) Get(ctx context.Context, playerID int64) (*model.Player, error) {
	return u.playerRepo.Get(ctx, &repo.PlayerQuery{
		ID: new(playerID),
	})
}
