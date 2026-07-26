package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/user/v1"
	"context"
	"time"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type OutboxService struct {
	v1.UnimplementedOutboxServiceServer
	outboxUsecase *usecase.OutboxUsecase
}

func NewOutboxService(
	outboxUsecase *usecase.OutboxUsecase,
) *OutboxService {
	return &OutboxService{
		outboxUsecase: outboxUsecase,
	}
}

func (s *OutboxService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterOutboxServiceServer(gs, s)
}

func (s *OutboxService) RegisterHttp(hs *http.Server) {
}

func (s *OutboxService) Publish(ctx context.Context, req *v1.PublishOutboxEvent_Req) (*v1.PublishOutboxEvent_Resp, error) {
	if req == nil || req.GetId() == 0 || req.GetMaxRetry() < 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	publishTimeout := time.Duration(0)
	if req.GetPublishTimeout() != nil {
		publishTimeout = req.GetPublishTimeout().AsDuration()
	}
	if publishTimeout < 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := s.outboxUsecase.Publish(ctx, &usecase.PublishOutboxEventReq{
		ID:             req.GetId(),
		PublishTimeout: publishTimeout,
		MaxRetry:       req.GetMaxRetry(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.PublishOutboxEvent_Resp{
		Published: resp.Published,
		Skipped:   resp.Skipped,
	}, nil
}

func (s *OutboxService) PublishBatch(ctx context.Context, req *v1.PublishOutboxEvents_Req) (*v1.PublishOutboxEvents_Resp, error) {
	limit := int32(0)
	publishTimeout := time.Duration(0)
	maxRetry := int32(0)
	if req != nil {
		if req.GetPublishTimeout() != nil {
			publishTimeout = req.GetPublishTimeout().AsDuration()
		}
		if req.GetLimit() < 0 || req.GetMaxRetry() < 0 || publishTimeout < 0 {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		limit = req.GetLimit()
		maxRetry = req.GetMaxRetry()
	}
	resp, err := s.outboxUsecase.PublishBatch(ctx, &usecase.PublishOutboxEventsReq{
		Limit:          int(limit),
		PublishTimeout: publishTimeout,
		MaxRetry:       maxRetry,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PublishOutboxEvents_Resp{
		Claimed:   resp.Claimed,
		Published: resp.Published,
		Failed:    resp.Failed,
		Skipped:   resp.Skipped,
	}, nil
}
