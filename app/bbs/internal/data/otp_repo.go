package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
)

var _ repo.OtpClient = (*OtpClient)(nil)

type OtpClient struct {
	userClient *rpc.UserClient
}

func NewOtpClient(
	userClient *rpc.UserClient,
) repo.OtpClient {
	return &OtpClient{
		userClient: userClient,
	}
}

func (r *OtpClient) BeginEnableTotp(ctx context.Context, req *repo.BeginEnableTotpReq) (*repo.BeginEnableTotpResp, error) {
	reply, err := r.userClient.Otp.BeginEnableTotp(ctx, &userv1.BeginEnableTotp_Req{
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

func (r *OtpClient) ConfirmEnableTotp(ctx context.Context, req *repo.ConfirmEnableTotpReq) error {
	_, err := r.userClient.Otp.ConfirmEnableTotp(ctx, &userv1.ConfirmEnableTotp_Req{
		UserId: req.UserID,
		Code:   req.Code,
	})
	return err
}

func (r *OtpClient) DisableTotp(ctx context.Context, req *repo.DisableTotpReq) error {
	_, err := r.userClient.Otp.DisableTotp(ctx, &userv1.DisableTotp_Req{
		UserId: req.UserID,
		Code:   req.Code,
	})
	return err
}

func (r *OtpClient) GetCurrentTotp(ctx context.Context, userID int64) (*repo.Totp, error) {
	reply, err := r.userClient.Otp.GetTotp(ctx, &userv1.GetTotp_Req{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	totpSetting := reply.GetTotp()
	if totpSetting == nil {
		return nil, nil
	}
	out := &repo.Totp{
		UserID: totpSetting.GetUserId(),
		Enable: totpSetting.GetEnable(),
	}
	if totpSetting.GetEnableTime() != nil {
		out.EnableTime = new(totpSetting.GetEnableTime().AsTime())
	}
	return out, nil
}

func (r *OtpClient) SendEmailOtp(ctx context.Context, req *repo.SendEmailOtpReq) (*repo.OtpCodeResp, error) {
	reply, err := r.userClient.Otp.SendEmailOtp(ctx, &userv1.SendEmailOtp_Req{
		Email:  req.Email,
		UserId: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &repo.OtpCodeResp{
		Code: reply.GetCode(),
	}, nil
}

func (r *OtpClient) SendPhoneOtp(ctx context.Context, req *repo.SendPhoneOtpReq) (*repo.OtpCodeResp, error) {
	reply, err := r.userClient.Otp.SendPhoneOtp(ctx, &userv1.SendPhoneOtp_Req{
		Phone:  req.Phone,
		UserId: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &repo.OtpCodeResp{
		Code: reply.GetCode(),
	}, nil
}
