package usecase

import (
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/base"
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

type StationMessagePageReq struct {
	Page       *base.PageRequest
	IDs        []int64
	ReceiverID *int64
	EventType  *enums.EventType
	Unread     *bool
}

type StationMessagePageResponse struct {
	Rows []*model.NotificationStationMessage
	Page *base.PageResponse
}

func (u *StationMessageUsecase) Page(ctx context.Context, req *StationMessagePageReq) (*StationMessagePageResponse, error) {
	if req == nil {
		req = &StationMessagePageReq{}
	}
	pageResponse, err := u.stationMessageRepo.Page(ctx, &repo.NotificationStationMessagePageReq{Query: &repo.NotificationStationMessageQuery{
		Page:       req.Page,
		IDs:        req.IDs,
		ReceiverID: req.ReceiverID,
		EventType:  req.EventType,
		Unread:     req.Unread,
	}})
	if err != nil {
		return nil, err
	}
	return &StationMessagePageResponse{Rows: pageResponse.Rows, Page: pageResponse.Page}, nil
}

type StationMessageMarkReadReq struct {
	ReceiverID int64
	IDs        []int64
	StartTime  *time.Time
	EndTime    *time.Time
}

type StationMessageMarkReadResponse struct {
	Count int
}

func (u *StationMessageUsecase) MarkRead(ctx context.Context, req *StationMessageMarkReadReq) (*StationMessageMarkReadResponse, error) {
	if req == nil {
		req = &StationMessageMarkReadReq{}
	}
	response, err := u.stationMessageRepo.MarkRead(ctx, &repo.NotificationStationMessageMarkReadReq{
		ReceiverID: req.ReceiverID,
		IDs:        req.IDs,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
	})
	if err != nil {
		return nil, err
	}
	return &StationMessageMarkReadResponse{Count: response.Count}, nil
}

type StationMessageCountUnreadReq struct {
	ReceiverID int64
}

type StationMessageCountUnreadResponse struct {
	Count int
}

func (u *StationMessageUsecase) CountUnread(ctx context.Context, req *StationMessageCountUnreadReq) (*StationMessageCountUnreadResponse, error) {
	if req == nil {
		req = &StationMessageCountUnreadReq{}
	}
	response, err := u.stationMessageRepo.CountUnread(ctx, &repo.NotificationStationMessageCountUnreadReq{ReceiverID: req.ReceiverID})
	if err != nil {
		return nil, err
	}
	return &StationMessageCountUnreadResponse{Count: response.Count}, nil
}
