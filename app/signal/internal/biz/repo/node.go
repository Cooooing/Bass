package repo

import (
	cv1 "common/api/common/v1"
	"context"
	"signal/internal/biz/model"
	"signal/internal/data/ent/gen"
)

type NodeRepo interface {
	Save(ctx context.Context, tx *gen.Client, node *model.Node) (*model.Node, error)

	Update(ctx context.Context, tx *gen.Client, node *model.Node) (*model.Node, error)
	UpdateSecret(ctx context.Context, tx *gen.Client, id int64, secret string) error

	GetOne(ctx context.Context, tx *gen.Client, req *NodeGetReq) (*model.Node, error)
	GetMap(ctx context.Context, tx *gen.Client, req *NodeGetReq) (map[string]*model.Node, error)
	GetList(ctx context.Context, tx *gen.Client, req *NodeGetReq) ([]*model.Node, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *NodeGetReq) ([]*model.Node, *cv1.PageReply, error)
}

type NodeGetReq struct {
	Id      *int64
	Ids     []int64
	Key     *string
	Keys    []string
	Name    *string
	OwnerId *int64
}
