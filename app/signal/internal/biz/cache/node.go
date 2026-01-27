package cache

import (
	"context"
	"signal/internal/biz/model"
)

type NodeCache interface {
	SetNode(ctx context.Context, n *model.Node) error
	DelNode(ctx context.Context, key string)
	UpdateNodeConnections(ctx context.Context, key string, delta int64) error
	UpdateNodePing(ctx context.Context, key string, pingMs int64) error
	UpdateNodePowCost(ctx context.Context, key string, powCostMs int64) error
	GetNode(ctx context.Context, key string) (*model.Node, error)

	SetNodeRank(ctx context.Context, key string, score float64) error
	DelNodeRank(ctx context.Context, key string)
	ExistsNodeRank(ctx context.Context, key string) (bool, error)
	GetOnlineNodeKeys(ctx context.Context) ([]string, error)
	CalculateScore(ctx context.Context, key string) (float64, error)
}
