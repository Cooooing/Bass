package repo

import (
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"
)

type PrivacySettingClient interface {
	GetCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.GetCurrentPrivacySetting_Request) (*bbsuserv1.GetCurrentPrivacySetting_Reply, error)
	UpdateCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.UpdateCurrentPrivacySetting_Request) (*bbsuserv1.UpdateCurrentPrivacySetting_Reply, error)
}
