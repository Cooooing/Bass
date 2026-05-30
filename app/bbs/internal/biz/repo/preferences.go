package repo

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
)

type PreferencesRepo interface {
	GetCurrentPreferences(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Reply, error)
	UpdateCurrentPreferences(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Request) (*bbsuserv1.UpdateCurrentPreferences_Reply, error)
}
