package service

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type TfaService struct {
	bbsuserv1.UnimplementedTfaServiceServer
	userClient *rpc.UserClient
}

func NewTfaService(userClient *rpc.UserClient) *TfaService {
	return &TfaService{userClient: userClient}
}

func (s *TfaService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterTfaServiceServer(gs, s)
}

func (s *TfaService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterTfaServiceHTTPServer(hs, s)
}

func (s *TfaService) Validate(ctx context.Context, req *bbsuserv1.ValidateTfa_Request) (*bbsuserv1.ValidateTfa_Reply, error) {
	reply, err := s.userClient.Tfa.Validate(forwardAuth(ctx), &userv1.ValidateTfa_Request{Code: req.GetCode()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ValidateTfa_Reply{Verified: reply.GetVerified()}, nil
}

func (s *TfaService) BeginEnable(ctx context.Context, req *bbsuserv1.BeginEnableTfa_Request) (*bbsuserv1.BeginEnableTfa_Reply, error) {
	reply, err := s.userClient.Tfa.BeginEnable(forwardAuth(ctx), &userv1.BeginEnableTfa_Request{})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.BeginEnableTfa_Reply{
		Data:        reply.GetData(),
		ContentType: reply.GetContentType(),
	}, nil
}

func (s *TfaService) ConfirmEnable(ctx context.Context, req *bbsuserv1.ConfirmEnableTfa_Request) (*bbsuserv1.ConfirmEnableTfa_Reply, error) {
	_, err := s.userClient.Tfa.ConfirmEnable(forwardAuth(ctx), &userv1.ConfirmEnableTfa_Request{Code: req.GetCode()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ConfirmEnableTfa_Reply{}, nil
}

func (s *TfaService) Disable(ctx context.Context, req *bbsuserv1.DisableTfa_Request) (*bbsuserv1.DisableTfa_Reply, error) {
	_, err := s.userClient.Tfa.Disable(forwardAuth(ctx), &userv1.DisableTfa_Request{Code: req.GetCode()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.DisableTfa_Reply{}, nil
}

func (s *TfaService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentTfa_Request) (*bbsuserv1.GetCurrentTfa_Reply, error) {
	reply, err := s.userClient.Tfa.GetCurrent(forwardAuth(ctx), &userv1.GetCurrentTfa_Request{})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentTfa_Reply{Tfa: toBFFTfa(reply.GetTfa())}, nil
}

func toBFFTfa(in *userv1.Tfa) *bbsuserv1.Tfa {
	if in == nil {
		return nil
	}
	return &bbsuserv1.Tfa{
		UserId:     in.GetUserId(),
		Enable:     in.GetEnable(),
		EnableTime: formatProtoTime(in.GetEnableTime()),
	}
}
