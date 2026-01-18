package repo

import (
	cv1 "common/api/common/v1"
	"context"
	"signal/internal/biz/model"
	"signal/internal/data/ent/gen"
)

type NodeRepo interface {
	Save(ctx context.Context, tx *gen.Client, Node *model.Node) (*model.Node, error)

	Update(ctx context.Context, tx *gen.Client, Node *model.Node) (*model.Node, error)
	UpdateSecret(ctx context.Context, tx *gen.Client, id int64, secret string) error

	GetOne(ctx context.Context, tx *gen.Client, req *NodeGetReq) (*model.Node, error)
	GetList(ctx context.Context, tx *gen.Client, req *NodeGetReq) ([]*model.Node, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *NodeGetReq) ([]*model.Node, *cv1.PageReply, error)

	GetByName(ctx context.Context, tx *gen.Client, name string) (*model.Node, error)

	Register(ctx context.Context, n *model.Node) error
	Unregister(ctx context.Context, name string) error
	UpdateConnections(ctx context.Context, name string, delta int64) error
	UpdatePing(ctx context.Context, name string, pingMs int64) error
	UpdatePowCost(ctx context.Context, name string, powCostMs int64) error
	UpdateScore(ctx context.Context, n *model.Node) error
}

type NodeGetReq struct {
	Id      *int64
	Ids     []int64
	Name    *string
	OwnerId *int64
}
