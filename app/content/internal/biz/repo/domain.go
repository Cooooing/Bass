package repo

import (
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type DomainRepo interface {
	Save(ctx context.Context, tx *gen.Client, domain *model.Domain) (*model.Domain, error)
	Update(ctx context.Context, tx *gen.Client, domain *model.Domain) (*model.Domain, error)

	GetById(ctx context.Context, tx *gen.Client, id int64) (*model.Domain, error)
}
