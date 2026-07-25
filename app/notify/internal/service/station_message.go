package service

import (
	commonenum "common/pkg/enum"
	"common/proto/gen/common"
	v1 "common/proto/gen/notify/v1"
	"context"
	"notify/internal/biz/base"
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

func NewStationMessageService(
	stationMessageUsecase *usecase.StationMessageUsecase,
) *StationMessageService {
	return &StationMessageService{
		stationMessageUsecase: stationMessageUsecase,
	}
}

func (s *StationMessageService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyStationMessageServiceServer(gs, s)
}

func (s *StationMessageService) RegisterHttp(hs *http.Server) {
}

func (s *StationMessageService) List(ctx context.Context, req *v1.ListStationMessages_Req) (*v1.ListStationMessages_Resp, error) {
	if req == nil {
		req = &v1.ListStationMessages_Req{}
	}
	query := &usecase.StationMessagePageReq{
		ReceiverID: new(req.GetUserId()),
	}
	if req != nil && req.Query != nil {
		query.IDs = req.Query.GetIds()
		query.EventType = req.Query.EventType
		query.Unread = req.Query.Unread
	}
	var pageReq *base.PageRequest
	if req != nil {
		pageReq = &base.PageRequest{
			Page: int64(req.GetPage().GetPage()),
			Size: int64(req.GetPage().GetSize()),
		}
	}
	query.Page = pageReq
	pageResp, err := s.stationMessageUsecase.Page(ctx, query)
	if err != nil {
		return nil, err
	}
	rows := pageResp.Rows
	replyRows := make([]*v1.ListStationMessages_Resp_NotificationStationMessage, 0, len(rows))
	for _, row := range rows {
		if row == nil {
			continue
		}
		item := &v1.ListStationMessages_Resp_NotificationStationMessage{
			Id:         row.ID,
			EventId:    row.EventID,
			ReceiverId: row.ReceiverID,
			Title:      row.Title,
			Content:    row.Content,
		}
		item.EventType = commonenum.EventTypeMap.MustToProto(row.EventType)
		item.Status = notifyenum.NotificationChannelStatusMap.MustToProto(row.Status)
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
	return &v1.ListStationMessages_Resp{
		Page: &common.PageResp{
			Page:  uint32(pageResp.Page.Page),
			Size:  uint32(pageResp.Page.Size),
			Total: uint32(pageResp.Page.Total),
		},
		Rows: replyRows,
	}, nil
}

func (s *StationMessageService) MarkRead(ctx context.Context, req *v1.MarkReadStationMessage_Req) (*v1.MarkReadStationMessage_Resp, error) {
	if req == nil {
		req = &v1.MarkReadStationMessage_Req{}
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
	markReadResp, err := s.stationMessageUsecase.MarkRead(ctx, &usecase.StationMessageMarkReadReq{
		ReceiverID: req.GetUserId(),
		IDs:        ids,
		StartTime:  startTime,
		EndTime:    endTime,
	})
	if err != nil {
		return nil, err
	}
	return &v1.MarkReadStationMessage_Resp{
		Count: int32(markReadResp),
	}, nil
}

func (s *StationMessageService) CountUnread(ctx context.Context, req *v1.CountUnreadStationMessages_Req) (*v1.CountUnreadStationMessages_Resp, error) {
	if req == nil {
		req = &v1.CountUnreadStationMessages_Req{}
	}
	countResp, err := s.stationMessageUsecase.CountUnread(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &v1.CountUnreadStationMessages_Resp{
		Count: int64(countResp),
	}, nil
}
