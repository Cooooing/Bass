package usecase

import (
	"context"

	"bff_bbs/internal/biz/model"
	"bff_bbs/internal/biz/repo"
)

type HolidayUsecase struct {
	holidayRepo repo.HolidayClient
}

func NewHolidayUsecase(holidayRepo repo.HolidayClient) *HolidayUsecase {
	return &HolidayUsecase{holidayRepo: holidayRepo}
}

func (u *HolidayUsecase) GetCurrent(ctx context.Context) (*model.HolidayInfo, error) {
	return u.holidayRepo.GetCurrent(ctx)
}
