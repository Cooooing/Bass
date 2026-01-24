package repo

import (
	cv1 "common/api/common/v1"
	"common/pkg/cutil/collections/dict"
	"context"
	"signal/internal/biz/model"
	"signal/internal/data/ent/gen"
)

type NodeRepo interface {
	Save(ctx context.Context, tx *gen.Client, node *model.Node) (*model.Node, error)

	Update(ctx context.Context, tx *gen.Client, node *model.Node) (*model.Node, error)
	UpdateSecret(ctx context.Context, tx *gen.Client, id int64, secret string) error

	GetOne(ctx context.Context, tx *gen.Client, req *NodeGetReq) (*model.Node, error)
	GetMap(ctx context.Context, tx *gen.Client, req *NodeGetReq) (dict.Map[string, *model.Node], error)
	GetList(ctx context.Context, tx *gen.Client, req *NodeGetReq) ([]*model.Node, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *NodeGetReq) ([]*model.Node, *cv1.PageReply, error)

	GetByKey(ctx context.Context, tx *gen.Client, key string) (*model.Node, error)

	Register(ctx context.Context, n *model.Node) error
	Unregister(ctx context.Context, key string) error
	IsOnline(ctx context.Context, key string) (bool, error)
	UpdateConnections(ctx context.Context, key string, delta int64) error
	UpdatePing(ctx context.Context, key string, pingMs int64) error
	UpdatePowCost(ctx context.Context, key string, powCostMs int64) error
	UpdateScore(ctx context.Context, n *model.Node) error
	GetOnlineNodeKeys(ctx context.Context) ([]string, error)
}

type NodeGetReq struct {
	Id      *int64
	Ids     []int64
	Key     *string
	Keys    []string
	Name    *string
	OwnerId *int64
}
