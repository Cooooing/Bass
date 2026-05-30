package repo

import (
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"
)

type ContentDomainRepo interface {
	ListDomains(ctx context.Context, req *bbscontentv1.ListDomains_Request) (*bbscontentv1.ListDomains_Reply, error)
}
