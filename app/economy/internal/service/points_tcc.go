package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/economy/v1"
	"context"

	"economy/internal/biz/usecase"
	"economy/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type PointsTccService struct {
	v1.UnimplementedPointsTccServiceServer

	pointsUsecase *usecase.PointsUsecase
}

func NewPointsTccService(pointsUsecase *usecase.PointsUsecase) *PointsTccService {
	return &PointsTccService{pointsUsecase: pointsUsecase}
}

func (s *PointsTccService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPointsTccServiceServer(gs, s)
}

func (s *PointsTccService) RegisterHTTP(hs *http.Server) {
}

func (s *PointsTccService) Try(ctx context.Context, req *v1.PointsTccReq) (*v1.PointsTccResp, error) {
	operation := s.pointsUsecase.Try
	return s.execute(ctx, req, operation)
}

func (s *PointsTccService) Comfirm(ctx context.Context, req *v1.PointsTccReq) (*v1.PointsTccResp, error) {
	operation := s.pointsUsecase.Comfirm
	return s.execute(ctx, req, operation)
}

func (s *PointsTccService) Cancel(ctx context.Context, req *v1.PointsTccReq) (*v1.PointsTccResp, error) {
	operation := s.pointsUsecase.Cancel
	return s.execute(ctx, req, operation)
}

func (s *PointsTccService) execute(ctx context.Context, request *v1.PointsTccReq, operation func(context.Context, *usecase.PointsTccReq) (*usecase.PointsTccResp, error)) (*v1.PointsTccResp, error) {
	recordType, ok := enum.EconomyRecordTypeMap.ToEnum(request.GetRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	response, err := operation(ctx, &usecase.PointsTccReq{UserID: request.GetUserId(), Amount: request.GetAmount(), RecordType: recordType, Remark: request.Remark})
	if err != nil {
		return nil, err
	}
	mapper := economyProtoMapper{}
	return &v1.PointsTccResp{Account: mapper.account(response.Account), Record: mapper.record(response.Record)}, nil
}
