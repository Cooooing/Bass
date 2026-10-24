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

type TCCPointsTransferService struct {
	v1.UnimplementedTCCPointsTransferServiceServer

	usecase *usecase.TCCPointsTransferUsecase
}

func NewTCCPointsTransferService(usecase *usecase.TCCPointsTransferUsecase) *TCCPointsTransferService {
	return &TCCPointsTransferService{usecase: usecase}
}

func (s *TCCPointsTransferService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterTCCPointsTransferServiceServer(server, s)
}

func (s *TCCPointsTransferService) RegisterHTTP(*http.Server) {}

func (s *TCCPointsTransferService) Try(ctx context.Context, req *v1.TCCPointsTransfer_Req) (*v1.TCCPointsTransfer_Resp, error) {
	request, err := s.requestFromProto(req)
	if err != nil {
		return nil, err
	}
	response, err := s.usecase.Try(ctx, request)
	if err != nil {
		return nil, err
	}
	return s.responseToProto(response), nil
}

func (s *TCCPointsTransferService) Confirm(ctx context.Context, req *v1.TCCPointsTransfer_Req) (*v1.TCCPointsTransfer_Resp, error) {
	request, err := s.requestFromProto(req)
	if err != nil {
		return nil, err
	}
	response, err := s.usecase.Confirm(ctx, request)
	if err != nil {
		return nil, err
	}
	return s.responseToProto(response), nil
}

func (s *TCCPointsTransferService) Cancel(ctx context.Context, req *v1.TCCPointsTransfer_Req) (*v1.TCCPointsTransfer_Resp, error) {
	request, err := s.requestFromProto(req)
	if err != nil {
		return nil, err
	}
	response, err := s.usecase.Cancel(ctx, request)
	if err != nil {
		return nil, err
	}
	return s.responseToProto(response), nil
}

func (*TCCPointsTransferService) requestFromProto(request *v1.TCCPointsTransfer_Req) (*usecase.TCCPointsTransferReq, error) {
	outRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(request.GetOutRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	inRecordType, ok := enum.EconomyRecordTypeMap.ToEnum(request.GetInRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	return &usecase.TCCPointsTransferReq{
		FromUserID:    request.GetFromUserId(),
		ToUserID:      request.GetToUserId(),
		Amount:        request.GetAmount(),
		OutRecordType: outRecordType,
		InRecordType:  inRecordType,
		Remark:        request.Remark,
	}, nil
}

func (*TCCPointsTransferService) responseToProto(response *usecase.TCCPointsTransferResp) *v1.TCCPointsTransfer_Resp {
	mapper := economyProtoMapper{}
	return &v1.TCCPointsTransfer_Resp{
		FromAccount: mapper.account(response.FromAccount),
		ToAccount:   mapper.account(response.ToAccount),
		OutRecord:   mapper.record(response.OutRecord),
		InRecord:    mapper.record(response.InRecord),
	}
}
