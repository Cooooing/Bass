package repo

import (
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type DomainRepo interface {
	Save(ctx context.Context, client *gen.Client, domain *model.Domain) (*model.Domain, error)
	Update(ctx context.Context, client *gen.Client, domain *model.Domain) (*model.Domain, error)
}
