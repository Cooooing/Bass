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

	Register(ctx context.Context, tx *gen.Client, id int64) error
}

type NodeGetReq struct {
	Id  int64
	Ids []int64
}
