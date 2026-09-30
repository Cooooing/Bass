package repo

import (
	"context"
	"platform/internal/biz/model"
)

type AssetRepo interface {
	CreateOrGet(ctx context.Context, row *model.Asset) (*model.Asset, error)
	Get(ctx context.Context, id int64) (*model.Asset, error)
	GetByHash(ctx context.Context, hash string) (*model.Asset, error)
	Map(ctx context.Context, ids []int64) (map[int64]*model.Asset, error)
	Block(ctx context.Context, id int64, reason string, operatorID int64) error
}
