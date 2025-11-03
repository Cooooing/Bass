package repo

import (
	v1 "common/api/content/v1"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type TagRepo interface {
	Save(ctx context.Context, tx *gen.Client, tag *model.Tag) (*model.Tag, error)

	GetById(ctx context.Context, tx *gen.Client, id int64) (*model.Tag, error)
	GetList(ctx context.Context, tx *gen.Client, req *v1.GetTagRequest) (*v1.GetTagReply, error)
}
