package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
)

var _ repo.TotpClient = (*TotpClient)(nil)

type TotpClient struct {
	userClient *rpc.UserClient
}

func NewTotpClient(userClient *rpc.UserClient) repo.TotpClient {
	return &TotpClient{userClient: userClient}
}

func (r *TotpClient) CheckEnableCodeTotp(ctx context.Context, req *repo.CheckEnableCodeTotpReq) (*repo.CheckEnableCodeTotpResponse, error) {
	reply, err := r.userClient.Totp.Validate(ctx, &userv1.ValidateTotp_Request{UserId: req.UserID, Code: req.Code})
	if err != nil {
		return nil, err
	}
	return &repo.CheckEnableCodeTotpResponse{Verified: reply.GetVerified()}, nil
}

func (r *TotpClient) ValidateTotp(ctx context.Context, req *repo.ValidateTotpReq) (*repo.ValidateTotpResponse, error) {
	reply, err := r.userClient.Totp.Validate(ctx, &userv1.ValidateTotp_Request{UserId: req.UserID, Code: req.Code})
	if err != nil {
		return nil, err
	}
	return &repo.ValidateTotpResponse{Verified: reply.GetVerified()}, nil
}

func (r *TotpClient) BeginEnableTotp(ctx context.Context, req *repo.BeginEnableTotpReq) (*repo.BeginEnableTotpResponse, error) {
	reply, err := r.userClient.Totp.BeginEnable(ctx, &userv1.BeginEnableTotp_Request{UserId: req.UserID, AccountName: req.AccountName})
	if err != nil {
		return nil, err
	}
	return &repo.BeginEnableTotpResponse{
		URL:    reply.GetUrl(),
		QRCode: reply.GetQrCode(),
	}, nil
}

func (r *TotpClient) ConfirmEnableTotp(ctx context.Context, req *repo.ConfirmEnableTotpReq) (*repo.ConfirmEnableTotpResponse, error) {
	_, err := r.userClient.Totp.ConfirmEnable(ctx, &userv1.ConfirmEnableTotp_Request{UserId: req.UserID, Code: req.Code})
	if err != nil {
		return nil, err
	}
	return &repo.ConfirmEnableTotpResponse{}, nil
}

func (r *TotpClient) DisableTotp(ctx context.Context, req *repo.DisableTotpReq) (*repo.DisableTotpResponse, error) {
	_, err := r.userClient.Totp.Disable(ctx, &userv1.DisableTotp_Request{UserId: req.UserID, Code: req.Code})
	if err != nil {
		return nil, err
	}
	return &repo.DisableTotpResponse{}, nil
}

func (r *TotpClient) GetCurrentTotp(ctx context.Context, req *repo.GetCurrentTotpReq) (*repo.GetCurrentTotpResponse, error) {
	reply, err := r.userClient.Totp.Get(ctx, &userv1.GetTotp_Request{UserId: req.UserID})
	if err != nil {
		return nil, err
	}
	totpSetting := reply.GetTotp()
	var out *repo.Totp
	if totpSetting != nil {
		out = &repo.Totp{
			UserID:     totpSetting.GetUserId(),
			Enable:     totpSetting.GetEnable(),
			EnableTime: formatProtoTime(totpSetting.GetEnableTime()),
		}
	}
	return &repo.GetCurrentTotpResponse{Totp: out}, nil
}
