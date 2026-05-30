package data

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.TotpRepo = (*TotpRepo)(nil)

type TotpRepo struct {
	userClient *rpc.UserClient
}

func NewTotpRepo(userClient *rpc.UserClient) repo.TotpRepo {
	return &TotpRepo{userClient: userClient}
}

func (r *TotpRepo) BeginEnableTotp(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Request) (*bbsuserv1.BeginEnableTotp_Reply, error) {
	user, err := currentUser(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Totp.BeginEnable(ctx, &userv1.BeginEnableTotp_Request{UserId: user.ID, AccountName: user.Name})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.BeginEnableTotp_Reply{
		Url:    reply.GetUrl(),
		QrCode: reply.GetQrCode(),
	}, nil
}

func (r *TotpRepo) ConfirmEnableTotp(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Request) (*bbsuserv1.ConfirmEnableTotp_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.userClient.Totp.ConfirmEnable(ctx, &userv1.ConfirmEnableTotp_Request{UserId: userID, Code: req.GetCode()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ConfirmEnableTotp_Reply{}, nil
}

func (r *TotpRepo) DisableTotp(ctx context.Context, req *bbsuserv1.DisableTotp_Request) (*bbsuserv1.DisableTotp_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.userClient.Totp.Disable(ctx, &userv1.DisableTotp_Request{UserId: userID, Code: req.GetCode()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.DisableTotp_Reply{}, nil
}

func (r *TotpRepo) GetCurrentTotp(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Request) (*bbsuserv1.GetCurrentTotp_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Totp.Get(ctx, &userv1.GetTotp_Request{UserId: userID})
	if err != nil {
		return nil, err
	}
	totpSetting := reply.GetTotp()
	var out *bbsuserv1.Totp
	if totpSetting != nil {
		out = &bbsuserv1.Totp{
			UserId:     totpSetting.GetUserId(),
			Enable:     totpSetting.GetEnable(),
			EnableTime: formatProtoTime(totpSetting.GetEnableTime()),
		}
	}
	return &bbsuserv1.GetCurrentTotp_Reply{Totp: out}, nil
}
