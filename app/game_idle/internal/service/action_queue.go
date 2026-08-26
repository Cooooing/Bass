package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ActionQueueService struct {
	v1.UnimplementedGameIdleActionQueueServiceServer
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewActionQueueService(actionQueueUsecase *usecase.ActionQueueUsecase) *ActionQueueService {
	return &ActionQueueService{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (s *ActionQueueService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterGameIdleActionQueueServiceServer(server, s)
}

func (s *ActionQueueService) RegisterHttp(*http.Server) {
}

func (s *ActionQueueService) List(ctx context.Context, req *v1.ListGameIdleActionQueue_Request) (*v1.ListGameIdleActionQueue_Resp, error) {
	if req.GetCharacterId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.actionQueueUsecase.List(ctx, req.GetCharacterId())
	if err != nil {
		return nil, err
	}
	queue := &v1.ActionQueue{
		CharacterId: row.Queue.CharacterID,
		Items:       make([]*v1.ActionQueueItem, 0, len(row.Queue.Items)),
	}
	for _, item := range row.Queue.Items {
		queue.Items = append(queue.Items, &v1.ActionQueueItem{
			ActionId:  item.ActionID,
			Times:     item.Times,
			CreatedAt: timestamppb.New(item.CreatedAt),
		})
	}
	return &v1.ListGameIdleActionQueue_Resp{Queue: queue}, nil
}

func (s *ActionQueueService) Add(ctx context.Context, req *v1.AddGameIdleAction_Request) (*v1.AddGameIdleAction_Resp, error) {
	if req.GetCharacterId() <= 0 || req.GetActionId() == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.actionQueueUsecase.Add(ctx, &usecase.AddActionReq{
		CharacterID: req.GetCharacterId(),
		ActionID:    req.GetActionId(),
		Times:       req.GetTimes(),
		Position:    req.Position,
	}); err != nil {
		return nil, err
	}
	return &v1.AddGameIdleAction_Resp{}, nil
}

func (s *ActionQueueService) Move(ctx context.Context, req *v1.MoveGameIdleAction_Request) (*v1.MoveGameIdleAction_Resp, error) {
	if req.GetCharacterId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.actionQueueUsecase.Move(ctx, &usecase.MoveActionReq{
		CharacterID:     req.GetCharacterId(),
		CurrentPosition: req.GetCurrentPosition(),
		TargetPosition:  req.GetTargetPosition(),
	}); err != nil {
		return nil, err
	}
	return &v1.MoveGameIdleAction_Resp{}, nil
}

func (s *ActionQueueService) Remove(ctx context.Context, req *v1.RemoveGameIdleAction_Request) (*v1.RemoveGameIdleAction_Resp, error) {
	if req.GetCharacterId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.actionQueueUsecase.Remove(ctx, &usecase.RemoveActionReq{
		CharacterID: req.GetCharacterId(),
		Position:    req.GetPosition(),
	}); err != nil {
		return nil, err
	}
	return &v1.RemoveGameIdleAction_Resp{}, nil
}

func (s *ActionQueueService) Clear(ctx context.Context, req *v1.ClearGameIdleActionQueue_Request) (*v1.ClearGameIdleActionQueue_Resp, error) {
	if req.GetCharacterId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.actionQueueUsecase.Clear(ctx, req.GetCharacterId()); err != nil {
		return nil, err
	}
	return &v1.ClearGameIdleActionQueue_Resp{}, nil
}
