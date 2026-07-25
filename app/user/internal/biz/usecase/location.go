package usecase

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
)

type LocationUsecase struct {
	locationRepo repo.LocationRepo
}

func NewLocationUsecase(
	locationRepo repo.LocationRepo,
) *LocationUsecase {
	return &LocationUsecase{
		locationRepo: locationRepo,
	}
}

func (s *LocationUsecase) GetByUserID(
	ctx context.Context,
	userID int64,
) (*model.Location, error) {
	return s.locationRepo.Get(ctx, &repo.LocationGetReq{
		UserID: &userID,
	})
}

func (s *LocationUsecase) UpsertByUserID(
	ctx context.Context,
	location *model.Location,
) (*model.Location, error) {
	return s.locationRepo.UpsertByUserID(ctx, location)
}
