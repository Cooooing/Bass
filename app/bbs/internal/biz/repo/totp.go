package repo

import (
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"
)

type TotpClient interface {
	CheckEnableCodeTotp(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Request) (bool, error)
	ValidateTotp(ctx context.Context, req *bbsuserv1.DisableTotp_Request) (bool, error)
	BeginEnableTotp(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Request) (*bbsuserv1.BeginEnableTotp_Reply, error)
	ConfirmEnableTotp(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Request) (*bbsuserv1.ConfirmEnableTotp_Reply, error)
	DisableTotp(ctx context.Context, req *bbsuserv1.DisableTotp_Request) (*bbsuserv1.DisableTotp_Reply, error)
	GetCurrentTotp(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Request) (*bbsuserv1.GetCurrentTotp_Reply, error)
}
