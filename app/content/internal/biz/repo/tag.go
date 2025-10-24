package repo

import (
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type TagRepo interface {
	Save(ctx context.Context, client *gen.Client, tag *model.Tag) (*model.Tag, error)
}
