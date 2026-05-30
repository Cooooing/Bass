package repo

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
)

type TotpRepo interface {
	BeginEnableTotp(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Request) (*bbsuserv1.BeginEnableTotp_Reply, error)
	ConfirmEnableTotp(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Request) (*bbsuserv1.ConfirmEnableTotp_Reply, error)
	DisableTotp(ctx context.Context, req *bbsuserv1.DisableTotp_Request) (*bbsuserv1.DisableTotp_Reply, error)
	GetCurrentTotp(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Request) (*bbsuserv1.GetCurrentTotp_Reply, error)
}
