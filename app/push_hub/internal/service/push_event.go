package service

import (
	"context"

	pushhubv1 "common/proto/gen/push_hub/v1"
	"push_hub/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// PushEventService 实现 PushHubEventServiceServer，委托给 usecase 层。
type PushEventService struct {
	pushhubv1.UnimplementedPushHubEventServiceServer
	pushEventUc *usecase.PushEventUsecase
}

func NewPushEventService(pushEventUc *usecase.PushEventUsecase) *PushEventService {
	return &PushEventService{pushEventUc: pushEventUc}
}

func (s *PushEventService) RegisterGrpc(gs *grpc.Server) {
	pushhubv1.RegisterPushHubEventServiceServer(gs, s)
}

func (s *PushEventService) RegisterHttp(hs *http.Server) {}

func (s *PushEventService) PublishEvent(ctx context.Context, req *pushhubv1.PublishEvent_Request) (*pushhubv1.PublishEvent_Response, error) {
	if err := s.pushEventUc.PublishEvent(ctx, &usecase.PublishEventReq{
		EventType: int32(req.Type),
		UserID:    req.UserId,
		Payload:   req.Payload,
	}); err != nil {
		return nil, err
	}
	return &pushhubv1.PublishEvent_Response{}, nil
}

func (s *PushEventService) BroadcastEvent(ctx context.Context, req *pushhubv1.BroadcastEvent_Request) (*pushhubv1.BroadcastEvent_Response, error) {
	if err := s.pushEventUc.BroadcastEvent(ctx, &usecase.BroadcastEventReq{
		EventType: int32(req.Type),
		Payload:   req.Payload,
	}); err != nil {
		return nil, err
	}
	return &pushhubv1.BroadcastEvent_Response{}, nil
}
