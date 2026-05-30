package repo

import (
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"
)

type ContentPostscriptRepo interface {
	AddPostscript(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Reply, error)
}
