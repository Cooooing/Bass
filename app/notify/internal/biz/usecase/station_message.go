package usecase

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"time"
)

type StationMessageUsecase struct {
	stationMessageRepo repo.NotificationStationMessageRepo
}

func NewStationMessageUsecase(
	stationMessageRepo repo.NotificationStationMessageRepo,
) *StationMessageUsecase {
	return &StationMessageUsecase{
		stationMessageRepo: stationMessageRepo,
	}
}

type StationMessagePageReq struct {
	Page       *base.PageRequest
	IDs        []int64
	ReceiverID *int64
	EventType  *commonenum.EventType
	Unread     *bool
}

type StationMessagePageResp struct {
	Rows []*model.NotificationStationMessage
	Page *base.PageResp
}

func (u *StationMessageUsecase) Page(ctx context.Context, req *StationMessagePageReq) (*StationMessagePageResp, error) {
	if req == nil {
		req = &StationMessagePageReq{}
	}
	pageResp, err := u.stationMessageRepo.Page(ctx, &repo.NotificationStationMessageQuery{
		Page:       req.Page,
		IDs:        req.IDs,
		ReceiverID: req.ReceiverID,
		EventType:  req.EventType,
		Unread:     req.Unread,
	})
	if err != nil {
		return nil, err
	}
	return &StationMessagePageResp{
		Rows: pageResp.Rows,
		Page: pageResp.Page,
	}, nil
}

type StationMessageMarkReadReq struct {
	ReceiverID int64
	IDs        []int64
	StartTime  *time.Time
	EndTime    *time.Time
}

func (u *StationMessageUsecase) MarkRead(ctx context.Context, req *StationMessageMarkReadReq) (int, error) {
	if req == nil {
		req = &StationMessageMarkReadReq{}
	}
	count, err := u.stationMessageRepo.MarkRead(ctx, &repo.NotificationStationMessageMarkReadReq{
		ReceiverID: req.ReceiverID,
		IDs:        req.IDs,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *StationMessageUsecase) CountUnread(ctx context.Context, receiverID int64) (int, error) {

	count, err := u.stationMessageRepo.CountUnread(ctx, receiverID)
	if err != nil {
		return 0, err
	}
	return count, nil
}
