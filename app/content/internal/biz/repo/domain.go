package repo

import (
	"content/internal/biz/model"
	"context"
)

type Domain interface {
	Save(ctx context.Context, domain *model.Domain) (*model.Domain, error)
}
