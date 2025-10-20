package repo

import (
	"content/internal/biz/model"
	"context"
)

type Tag interface {
	Save(ctx context.Context, tag *model.Tag) (*model.Tag, error)
}
