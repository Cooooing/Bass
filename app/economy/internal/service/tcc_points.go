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

type TCCPointsService struct {
	v1.UnimplementedTCCPointsServiceServer

	usecase *usecase.TCCPointsUsecase
}

func NewTCCPointsService(usecase *usecase.TCCPointsUsecase) *TCCPointsService {
	return &TCCPointsService{usecase: usecase}
}

func (s *TCCPointsService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterTCCPointsServiceServer(server, s)
}

func (s *TCCPointsService) RegisterHTTP(*http.Server) {}

func (s *TCCPointsService) Try(ctx context.Context, req *v1.TCCPoints_Req) (*v1.TCCPoints_Resp, error) {
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

func (s *TCCPointsService) Confirm(ctx context.Context, req *v1.TCCPoints_Req) (*v1.TCCPoints_Resp, error) {
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

func (s *TCCPointsService) Cancel(ctx context.Context, req *v1.TCCPoints_Req) (*v1.TCCPoints_Resp, error) {
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

func (*TCCPointsService) requestFromProto(request *v1.TCCPoints_Req) (*usecase.TCCPointsReq, error) {
	recordType, ok := enum.EconomyRecordTypeMap.ToEnum(request.GetRecordType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	return &usecase.TCCPointsReq{
		UserID:     request.GetUserId(),
		Amount:     request.GetAmount(),
		RecordType: recordType,
		Remark:     request.Remark,
	}, nil
}

func (*TCCPointsService) responseToProto(response *usecase.TCCPointsResp) *v1.TCCPoints_Resp {
	mapper := economyProtoMapper{}
	return &v1.TCCPoints_Resp{
		Account: mapper.account(response.Account),
		Record:  mapper.record(response.Record),
	}
}
