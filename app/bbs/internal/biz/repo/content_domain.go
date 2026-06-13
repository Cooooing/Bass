package repo

import (
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"
)

type ContentDomainClient interface {
	ListDomains(ctx context.Context, req *bbscontentv1.ListDomains_Request) (*bbscontentv1.ListDomains_Reply, error)
}
