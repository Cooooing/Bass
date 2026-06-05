package usecase

import (
	"common/api/gen/common"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"time"
)

type StationMessageUsecase struct {
	stationMessageRepo repo.NotificationStationMessageRepo
}

func NewStationMessageUsecase(stationMessageRepo repo.NotificationStationMessageRepo) *StationMessageUsecase {
	return &StationMessageUsecase{stationMessageRepo: stationMessageRepo}
}

func (u *StationMessageUsecase) Page(ctx context.Context, page *common.PageRequest, query *repo.NotificationStationMessageQuery) ([]*model.NotificationStationMessage, *common.PageReply, error) {
	return u.stationMessageRepo.Page(ctx, page, query)
}

func (u *StationMessageUsecase) MarkRead(ctx context.Context, receiverID int64, ids []int64, startTime *time.Time, endTime *time.Time) (int, error) {
	return u.stationMessageRepo.MarkRead(ctx, receiverID, ids, startTime, endTime)
}

func (u *StationMessageUsecase) CountUnread(ctx context.Context, receiverID int64) (int, error) {
	return u.stationMessageRepo.CountUnread(ctx, receiverID)
}
