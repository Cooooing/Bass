package service

import (
	pushhubv1 "common/proto/gen/push_hub/v1"
	"push_hub/internal/biz/usecase"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
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

func (s *PushEventService) PublishEvent(ctx context.Context, req *pushhubv1.PublishEvent_Request) (*pushhubv1.PublishEvent_Reply, error) {
	if err := s.pushEventUc.PublishEvent(ctx, int32(req.Type), req.UserId, req.Payload); err != nil {
		return nil, err
	}
	return &pushhubv1.PublishEvent_Reply{}, nil
}

func (s *PushEventService) BroadcastEvent(ctx context.Context, req *pushhubv1.BroadcastEvent_Request) (*pushhubv1.BroadcastEvent_Reply, error) {
	if err := s.pushEventUc.BroadcastEvent(ctx, int32(req.Type), req.Payload); err != nil {
		return nil, err
	}
	return &pushhubv1.BroadcastEvent_Reply{}, nil
}
