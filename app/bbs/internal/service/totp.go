package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type TotpService struct {
	bbsuserv1.UnimplementedTotpServiceServer
	totpUsecase *usecase.TotpUsecase
}

func NewTotpService(
	totpUsecase *usecase.TotpUsecase,
) *TotpService {
	return &TotpService{
		totpUsecase: totpUsecase,
	}
}

func (s *TotpService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterTotpServiceHTTPServer(hs, s)
}

func (s *TotpService) RegisterGrpc(gs *grpc.Server) {
}

func (s *TotpService) BeginEnable(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Req) (*bbsuserv1.BeginEnableTotp_Resp, error) {
	user, err := currentUser(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.totpUsecase.BeginEnableTotp(ctx, &usecase.BeginEnableTotpReq{
		UserID:      user.ID,
		AccountName: user.Name,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.BeginEnableTotp_Resp{
		Url:    resp.URL,
		QrCode: resp.QRCode,
	}, nil
}

func (s *TotpService) ConfirmEnable(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Req) (*bbsuserv1.ConfirmEnableTotp_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.totpUsecase.ConfirmEnableTotp(ctx, &usecase.ConfirmEnableTotpReq{
		UserID: userID,
		Code:   req.GetCode(),
	})
	return &bbsuserv1.ConfirmEnableTotp_Resp{}, err
}

func (s *TotpService) Disable(ctx context.Context, req *bbsuserv1.DisableTotp_Req) (*bbsuserv1.DisableTotp_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.totpUsecase.DisableTotp(ctx, &usecase.DisableTotpReq{
		UserID: userID,
		Code:   req.GetCode(),
	})
	return &bbsuserv1.DisableTotp_Resp{}, err
}

func (s *TotpService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Req) (*bbsuserv1.GetCurrentTotp_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	totp, err := s.totpUsecase.GetCurrentTotp(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentTotp_Resp{
		Totp: totp,
	}, nil
}
