package repo

import (
	cv1 "common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type TagRepo interface {
	Save(ctx context.Context, tx *gen.Client, tag *model.Tag) (*model.Tag, error)
	Saves(ctx context.Context, tx *gen.Client, tags []*model.Tag) ([]*model.Tag, error)

	GetById(ctx context.Context, tx *gen.Client, id int64) (*model.Tag, error)
	GetList(ctx context.Context, tx *gen.Client, req *TagGetReq) ([]*model.Tag, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *TagGetReq) ([]*model.Tag, *cv1.PageReply, error)
}

type TagGetReq struct {
}
