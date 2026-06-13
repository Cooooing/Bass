package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
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

func (r *TotpClient) CheckEnableCodeTotp(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Request) (bool, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return false, err
	}
	reply, err := r.userClient.Totp.Validate(ctx, &userv1.ValidateTotp_Request{UserId: userID, Code: req.GetCode()})
	if err != nil {
		return false, err
	}
	return reply.GetVerified(), nil
}

func (r *TotpClient) ValidateTotp(ctx context.Context, req *bbsuserv1.DisableTotp_Request) (bool, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return false, err
	}
	reply, err := r.userClient.Totp.Validate(ctx, &userv1.ValidateTotp_Request{UserId: userID, Code: req.GetCode()})
	if err != nil {
		return false, err
	}
	return reply.GetVerified(), nil
}

func (r *TotpClient) BeginEnableTotp(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Request) (*bbsuserv1.BeginEnableTotp_Reply, error) {
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

func (r *TotpClient) ConfirmEnableTotp(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Request) (*bbsuserv1.ConfirmEnableTotp_Reply, error) {
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

func (r *TotpClient) DisableTotp(ctx context.Context, req *bbsuserv1.DisableTotp_Request) (*bbsuserv1.DisableTotp_Reply, error) {
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

func (r *TotpClient) GetCurrentTotp(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Request) (*bbsuserv1.GetCurrentTotp_Reply, error) {
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
