package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v3/transport/http"
)

type TotpService struct {
	bbsuserv1.UnimplementedTotpServiceServer
	totpUsecase *usecase.TotpUsecase
}

func NewTotpService(totpUsecase *usecase.TotpUsecase) *TotpService {
	return &TotpService{totpUsecase: totpUsecase}
}
func (s *TotpService) RegisterHttp(hs *http.Server) { bbsuserv1.RegisterTotpServiceHTTPServer(hs, s) }
func (s *TotpService) BeginEnable(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Request) (*bbsuserv1.BeginEnableTotp_Response, error) {
	user, err := currentUser(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.totpUsecase.BeginEnableTotp(ctx, &usecase.BeginEnableTotpReq{UserID: user.ID, AccountName: user.Name})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.BeginEnableTotp_Response{Url: response.URL, QrCode: response.QRCode}, nil
}
func (s *TotpService) ConfirmEnable(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Request) (*bbsuserv1.ConfirmEnableTotp_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.totpUsecase.ConfirmEnableTotp(ctx, &usecase.ConfirmEnableTotpReq{UserID: userID, Code: req.GetCode()})
	return &bbsuserv1.ConfirmEnableTotp_Response{}, err
}
func (s *TotpService) Disable(ctx context.Context, req *bbsuserv1.DisableTotp_Request) (*bbsuserv1.DisableTotp_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.totpUsecase.DisableTotp(ctx, &usecase.DisableTotpReq{UserID: userID, Code: req.GetCode()})
	return &bbsuserv1.DisableTotp_Response{}, err
}
func (s *TotpService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Request) (*bbsuserv1.GetCurrentTotp_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.totpUsecase.GetCurrentTotp(ctx, &usecase.GetCurrentTotpReq{UserID: userID})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentTotp_Response{Totp: response.Totp}, nil
}
