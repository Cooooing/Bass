package repo

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
)

type LocationRepo interface {
	GetCurrentLocation(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Request) (*bbsuserv1.GetCurrentLocation_Reply, error)
	UpsertCurrentLocation(ctx context.Context, req *bbsuserv1.UpsertCurrentLocation_Request) (*bbsuserv1.UpsertCurrentLocation_Reply, error)
}
