package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"user/internal/biz/repo"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TfaService struct {
	v1.UnimplementedTfaServiceServer
	tfaUsecase *usecase.TfaUsecase
	tfaRepo    repo.TfaRepo
}

func NewTfaService(tfaUsecase *usecase.TfaUsecase, tfaRepo repo.TfaRepo) *TfaService {
	return &TfaService{
		tfaUsecase: tfaUsecase,
		tfaRepo:    tfaRepo,
	}
}

func (s *TfaService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterTfaServiceServer(gs, s)
}

func (s *TfaService) RegisterHttp(hs *http.Server) {}

func (s *TfaService) Validate(ctx context.Context, req *v1.ValidateTfa_Request) (*v1.ValidateTfa_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	verified := s.tfaUsecase.Validate(ctx, current.TwofaSecret, req.Code)
	return &v1.ValidateTfa_Reply{Verified: verified}, nil
}

func (s *TfaService) BeginEnable(ctx context.Context, req *v1.BeginEnableTfa_Request) (*v1.BeginEnableTfa_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	if current.TwofaEnable {
		return nil, cerrors.ErrorBadRequest("2FA already enabled")
	}
	buf, err := s.tfaUsecase.Enable(ctx, current.Name)
	if err != nil {
		return nil, err
	}
	return &v1.BeginEnableTfa_Reply{Data: buf, ContentType: "image/png"}, nil
}

func (s *TfaService) ConfirmEnable(ctx context.Context, req *v1.ConfirmEnableTfa_Request) (*v1.ConfirmEnableTfa_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err := s.tfaUsecase.Confirm(ctx, current.Name, req.Code)
	return &v1.ConfirmEnableTfa_Reply{}, err
}

func (s *TfaService) Disable(ctx context.Context, req *v1.DisableTfa_Request) (*v1.DisableTfa_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	if !current.TwofaEnable {
		return nil, cerrors.ErrorBadRequest("2FA already disabled")
	}
	err := s.tfaUsecase.Disable(ctx, current.Name, current.TwofaSecret, req.Code)
	return &v1.DisableTfa_Reply{}, err
}

func (s *TfaService) GetCurrent(ctx context.Context, req *v1.GetCurrentTfa_Request) (*v1.GetCurrentTfa_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	tfa, _ := s.tfaRepo.FindByUserID(ctx, current.ID)
	reply := &v1.Tfa{UserId: current.ID}
	if tfa != nil {
		reply.Enable = tfa.Enable
		if tfa.EnableTime != nil {
			reply.EnableTime = timestamppb.New(*tfa.EnableTime)
		}
	}
	return &v1.GetCurrentTfa_Reply{Tfa: reply}, nil
}
