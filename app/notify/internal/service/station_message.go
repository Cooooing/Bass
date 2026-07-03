package service

import (
	commonenum "common/pkg/enum"
	"common/proto/gen/common"
	v1 "common/proto/gen/notify/v1"
	"context"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StationMessageService struct {
	v1.UnimplementedNotifyStationMessageServiceServer
	stationMessageUsecase *usecase.StationMessageUsecase
}

func NewStationMessageService(stationMessageUsecase *usecase.StationMessageUsecase) *StationMessageService {
	return &StationMessageService{stationMessageUsecase: stationMessageUsecase}
}

func (s *StationMessageService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyStationMessageServiceServer(gs, s)
}

func (s *StationMessageService) RegisterHttp(hs *http.Server) {}

func (s *StationMessageService) List(ctx context.Context, req *v1.ListStationMessages_Request) (*v1.ListStationMessages_Reply, error) {
	if req == nil {
		req = &v1.ListStationMessages_Request{}
	}
	query := &repo.NotificationStationMessageQuery{ReceiverID: new(req.GetUserId())}
	if req != nil && req.Query != nil {
		query.IDs = req.Query.GetIds()
		query.EventType = req.Query.EventType
		query.Unread = req.Query.Unread
	}
	var pageReq *common.PageRequest
	if req != nil {
		pageReq = req.GetPage()
	}
	rows, page, err := s.stationMessageUsecase.Page(ctx, pageReq, query)
	if err != nil {
		return nil, err
	}
	replyRows := make([]*v1.NotificationStationMessage, 0, len(rows))
	for _, row := range rows {
		if row == nil {
			continue
		}
		item := &v1.NotificationStationMessage{
			Id:         row.ID,
			EventId:    row.EventID,
			ReceiverId: row.ReceiverID,
			Title:      row.Title,
			Content:    row.Content,
		}
		if protoEventType, ok := commonenum.EventTypeMap.ToProto(row.EventType); ok {
			item.EventType = protoEventType
		}
		if protoStatus, ok := notifyenum.NotificationChannelStatusMap.ToProto(row.Status); ok {
			item.Status = protoStatus
		}
		if row.CreatedAt != nil {
			item.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		if row.ReadAt != nil {
			item.ReadAt = timestamppb.New(*row.ReadAt)
		}
		replyRows = append(replyRows, item)
	}
	return &v1.ListStationMessages_Reply{Page: page, Rows: replyRows}, nil
}

func (s *StationMessageService) MarkRead(ctx context.Context, req *v1.MarkReadStationMessage_Request) (*v1.MarkReadStationMessage_Reply, error) {
	if req == nil {
		req = &v1.MarkReadStationMessage_Request{}
	}
	var startTime *time.Time
	var endTime *time.Time
	if req != nil && req.CreatedTimeRange != nil {
		if req.CreatedTimeRange.Start != nil {
			startTime = new(req.CreatedTimeRange.Start.AsTime())
		}
		if req.CreatedTimeRange.End != nil {
			endTime = new(req.CreatedTimeRange.End.AsTime())
		}
	}
	var ids []int64
	if req != nil {
		ids = req.GetIds()
	}
	count, err := s.stationMessageUsecase.MarkRead(ctx, req.GetUserId(), ids, startTime, endTime)
	if err != nil {
		return nil, err
	}
	return &v1.MarkReadStationMessage_Reply{Count: int32(count)}, nil
}

func (s *StationMessageService) CountUnread(ctx context.Context, req *v1.CountUnreadStationMessages_Request) (*v1.CountUnreadStationMessages_Reply, error) {
	if req == nil {
		req = &v1.CountUnreadStationMessages_Request{}
	}
	count, err := s.stationMessageUsecase.CountUnread(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &v1.CountUnreadStationMessages_Reply{Count: int64(count)}, nil
}
