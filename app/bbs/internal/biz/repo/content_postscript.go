package repo

import (
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"
)

type ContentPostscriptClient interface {
	AddPostscript(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Reply, error)
}
