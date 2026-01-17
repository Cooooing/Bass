package data

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"context"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"signal/internal/data/ent/gen"
)

type NodeRepo struct {
	*BaseRepo
}

func NewNodeRepo(baseRepo *BaseRepo) *NodeRepo {
	return &NodeRepo{
		BaseRepo: baseRepo,
	}
}

func (r *NodeRepo) Save(ctx context.Context, tx *gen.Client, Node *model.Node) (*model.Node, error) {
	save, err := tx.Node.Create().
		SetID(Node.ID).
		SetNillableOwnerID(Node.OwnerID).
		SetName(Node.Name).
		SetNillableDescription(Node.Description).
		SetSecret(Node.Secret).
		SetCallbackURL(Node.CallbackURL).
		SetStatus(Node.Status).
		Save(ctx)
	return &model.Node{Node: save}, err
}

func (r *NodeRepo) Update(ctx context.Context, tx *gen.Client, Node *model.Node) (*model.Node, error) {
	update, err := tx.Node.UpdateOneID(Node.ID).
		SetNillableOwnerID(Node.OwnerID).
		SetName(Node.Name).
		SetNillableDescription(Node.Description).
		SetCallbackURL(Node.CallbackURL).
		SetStatus(Node.Status).
		Save(ctx)
	return &model.Node{Node: update}, err
}

func (r *NodeRepo) UpdateSecret(ctx context.Context, tx *gen.Client, id int64, secret string) error {
	_, err := tx.Node.UpdateOneID(id).
		SetSecret(secret).
		Save(ctx)
	return err
}

func (r *NodeRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.NodeGetReq) (*model.Node, error) {
	query := tx.Node.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cv1.ErrorBadRequest("Node is not found")
	}
	return &model.Node{Node: t}, err
}

func (r *NodeRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.NodeGetReq) ([]*model.Node, error) {
	var (
		Nodes []*model.Node
		err   error
	)
	query := tx.Node.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		Nodes = append(Nodes, &model.Node{Node: item})
	}
	return Nodes, nil
}

func (r *NodeRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.NodeGetReq) ([]*model.Node, *cv1.PageReply, error) {
	var (
		Nodes []*model.Node
		err   error
	)
	page = constant.PageValid(page)
	query := tx.Node.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	for _, item := range list {
		Nodes = append(Nodes, &model.Node{Node: item})
	}
	return Nodes, &cv1.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *NodeRepo) getQuery(query *gen.NodeQuery, req *repo.NodeGetReq) *gen.NodeQuery {
	return query
}

func (r *NodeRepo) Register(ctx context.Context, tx *gen.Client, id int64) error {
	// TODO implement me
	panic("implement me")
}
