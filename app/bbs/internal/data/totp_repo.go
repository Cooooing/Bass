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

func NewTotpClient(
	userClient *rpc.UserClient,
) repo.TotpClient {
	return &TotpClient{
		userClient: userClient,
	}
}

func (r *TotpClient) CheckEnableCodeTotp(ctx context.Context, req *repo.CheckEnableCodeTotpReq) (bool, error) {
	reply, err := r.userClient.Totp.Validate(ctx, &userv1.ValidateTotp_Req{
		UserId: req.UserID,
		Code:   req.Code,
	})
	if err != nil {
		return false, err
	}
	return reply.GetVerified(), nil
}

func (r *TotpClient) ValidateTotp(ctx context.Context, req *repo.ValidateTotpReq) (bool, error) {
	reply, err := r.userClient.Totp.Validate(ctx, &userv1.ValidateTotp_Req{
		UserId: req.UserID,
		Code:   req.Code,
	})
	if err != nil {
		return false, err
	}
	return reply.GetVerified(), nil
}

func (r *TotpClient) BeginEnableTotp(ctx context.Context, req *repo.BeginEnableTotpReq) (*repo.BeginEnableTotpResp, error) {
	reply, err := r.userClient.Totp.BeginEnable(ctx, &userv1.BeginEnableTotp_Req{
		UserId:      req.UserID,
		AccountName: req.AccountName,
	})
	if err != nil {
		return nil, err
	}
	return &repo.BeginEnableTotpResp{
		URL:    reply.GetUrl(),
		QRCode: reply.GetQrCode(),
	}, nil
}

func (r *TotpClient) ConfirmEnableTotp(ctx context.Context, req *repo.ConfirmEnableTotpReq) error {
	_, err := r.userClient.Totp.ConfirmEnable(ctx, &userv1.ConfirmEnableTotp_Req{
		UserId: req.UserID,
		Code:   req.Code,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *TotpClient) DisableTotp(ctx context.Context, req *repo.DisableTotpReq) error {
	_, err := r.userClient.Totp.Disable(ctx, &userv1.DisableTotp_Req{
		UserId: req.UserID,
		Code:   req.Code,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *TotpClient) GetCurrentTotp(ctx context.Context, userID int64) (*repo.Totp, error) {
	reply, err := r.userClient.Totp.Get(ctx, &userv1.GetTotp_Req{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	totpSetting := reply.GetTotp()
	var out *repo.Totp
	if totpSetting != nil {
		out = &repo.Totp{
			UserID: totpSetting.GetUserId(),
			Enable: totpSetting.GetEnable(),
		}
		if totpSetting.GetEnableTime() != nil {
			out.EnableTime = new(totpSetting.GetEnableTime().AsTime())
		}
	}
	return out, nil
}
