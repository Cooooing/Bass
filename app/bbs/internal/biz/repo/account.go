package repo

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"common/api/gen/common"
	"context"
)

type AccountRepo interface {
	GetCurrentAccount(ctx context.Context, req *bbsuserv1.GetCurrentAccount_Request) (*bbsuserv1.GetCurrentAccount_Reply, error)
	GetProfileAccount(ctx context.Context, req *bbsuserv1.GetProfileAccount_Request) (*bbsuserv1.GetProfileAccount_Reply, error)
	UpdateProfileAccount(ctx context.Context, req *bbsuserv1.UpdateProfileAccount_Request) (*bbsuserv1.UpdateProfileAccount_Reply, error)
	AvatarAccount(ctx context.Context, req *bbsuserv1.AvatarAccount_Request) (*common.ImageReply, error)
}
