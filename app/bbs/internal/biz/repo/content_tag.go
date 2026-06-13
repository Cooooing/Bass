package repo

import (
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"
)

type ContentTagClient interface {
	CreateTag(ctx context.Context, req *bbscontentv1.CreateTag_Request) (*bbscontentv1.CreateTag_Reply, error)
	UpdateTag(ctx context.Context, req *bbscontentv1.UpdateTag_Request) (*bbscontentv1.UpdateTag_Reply, error)
	ListTags(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error)
}
