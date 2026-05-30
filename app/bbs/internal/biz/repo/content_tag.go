package repo

import (
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"
)

type ContentTagRepo interface {
	ListTags(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error)
}
