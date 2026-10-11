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

type PointsTransferTccService struct {
	v1.UnimplementedPointsTransferTccServiceServer
	pointsUsecase *usecase.PointsUsecase
}

func NewPointsTransferTccService(pointsUsecase *usecase.PointsUsecase) *PointsTransferTccService {
	return &PointsTransferTccService{pointsUsecase: pointsUsecase}
}

func (s *PointsTransferTccService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPointsTransferTccServiceServer(gs, s)
}

func (s *PointsTransferTccService) RegisterHTTP(hs *http.Server) {}

func (s *PointsTransferTccService) Try(ctx context.Context, req *v1.PointsTransferTccReq) (*v1.PointsTransferTccResp, error) {
	usecaseRequest, err := s.requestFromProto(req)
	if err != nil {
		return nil, err
	}
	response, err := s.pointsUsecase.TransferTry(ctx, usecaseRequest)
	if err != nil {
		return nil, err
	}
	return s.responseToProto(response), nil
}

func (s *PointsTransferTccService) Comfirm(ctx context.Context, req *v1.PointsTransferTccReq) (*v1.PointsTransferTccResp, error) {
	usecaseRequest, err := s.requestFromProto(req)
	if err != nil {
		return nil, err
	}
	response, err := s.pointsUsecase.TransferComfirm(ctx, usecaseRequest)
	if err != nil {
		return nil, err
	}
	return s.responseToProto(response), nil
}

func (s *PointsTransferTccService) Cancel(ctx context.Context, req *v1.PointsTransferTccReq) (*v1.PointsTransferTccResp, error) {
	usecaseRequest, err := s.requestFromProto(req)
	if err != nil {
		return nil, err
	}
	response, err := s.pointsUsecase.TransferCancel(ctx, usecaseRequest)
	if err != nil {
		return nil, err
	}
	return s.responseToProto(response), nil
}

func (*PointsTransferTccService) requestFromProto(request *v1.PointsTransferTccReq) (*usecase.PointsTransferTccReq, error) {
	outRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(request.GetOutRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	inRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(request.GetInRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	return &usecase.PointsTransferTccReq{FromUserID: request.GetFromUserId(), ToUserID: request.GetToUserId(), Amount: request.GetAmount(), OutRecordType: outRecordType, InRecordType: inRecordType, Remark: request.Remark}, nil
}

func (*PointsTransferTccService) responseToProto(response *usecase.PointsTransferTccResp) *v1.PointsTransferTccResp {
	mapper := economyProtoMapper{}
	return &v1.PointsTransferTccResp{FromAccount: mapper.account(response.FromAccount), ToAccount: mapper.account(response.ToAccount), OutRecord: mapper.record(response.OutRecord), InRecord: mapper.record(response.InRecord)}
}
