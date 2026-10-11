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
	usecaseRequest, err := s.requestFromProto(req)
	if err != nil {
		return nil, err
	}
	response, err := s.pointsUsecase.Try(ctx, usecaseRequest)
	if err != nil {
		return nil, err
	}
	return s.responseToProto(response), nil
}

func (s *PointsTccService) Comfirm(ctx context.Context, req *v1.PointsTccReq) (*v1.PointsTccResp, error) {
	usecaseRequest, err := s.requestFromProto(req)
	if err != nil {
		return nil, err
	}
	response, err := s.pointsUsecase.Comfirm(ctx, usecaseRequest)
	if err != nil {
		return nil, err
	}
	return s.responseToProto(response), nil
}

func (s *PointsTccService) Cancel(ctx context.Context, req *v1.PointsTccReq) (*v1.PointsTccResp, error) {
	usecaseRequest, err := s.requestFromProto(req)
	if err != nil {
		return nil, err
	}
	response, err := s.pointsUsecase.Cancel(ctx, usecaseRequest)
	if err != nil {
		return nil, err
	}
	return s.responseToProto(response), nil
}

func (*PointsTccService) requestFromProto(request *v1.PointsTccReq) (*usecase.PointsTccReq, error) {
	recordType, ok := enum.EconomyRecordTypeMap.ToEnum(request.GetRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	return &usecase.PointsTccReq{UserID: request.GetUserId(), Amount: request.GetAmount(), RecordType: recordType, Remark: request.Remark}, nil
}

func (*PointsTccService) responseToProto(response *usecase.PointsTccResp) *v1.PointsTccResp {
	mapper := economyProtoMapper{}
	return &v1.PointsTccResp{Account: mapper.account(response.Account), Record: mapper.record(response.Record)}
}
