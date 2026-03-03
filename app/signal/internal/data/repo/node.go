package repo

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"context"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"signal/internal/data/base"
	"signal/internal/data/ent/gen"
	"signal/internal/data/ent/gen/node"
)

type NodeRepo struct {
	*base.BaseData
}

func NewNodeRepo(baseData *base.BaseData) repo.NodeRepo {
	return &NodeRepo{
		BaseData: baseData,
	}
}

func (r *NodeRepo) Save(ctx context.Context, tx *gen.Client, node *model.Node) (*model.Node, error) {
	save, err := tx.Node.Create().
		SetNillableOwnerID(node.OwnerID).
		SetKey(node.Key).
		SetName(node.Name).
		SetNillableDescription(node.Description).
		SetSecret(node.Secret).
		SetCallbackURL(node.CallbackURL).
		SetStatus(node.Status).
		Save(ctx)
	return &model.Node{Node: save}, err
}

func (r *NodeRepo) Update(ctx context.Context, tx *gen.Client, node *model.Node) (*model.Node, error) {
	update, err := tx.Node.UpdateOneID(node.ID).
		SetNillableOwnerID(node.OwnerID).
		SetKey(node.Key).
		SetName(node.Name).
		SetNillableDescription(node.Description).
		SetCallbackURL(node.CallbackURL).
		SetStatus(node.Status).
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

func (r *NodeRepo) GetMap(ctx context.Context, tx *gen.Client, req *repo.NodeGetReq) (map[string]*model.Node, error) {
	var err error
	nodes := make(map[string]*model.Node)
	query := tx.Node.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		nodes[item.Key] = &model.Node{Node: item}
	}
	return nodes, nil
}

func (r *NodeRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.NodeGetReq) ([]*model.Node, error) {
	var (
		nodes []*model.Node
		err   error
	)
	query := tx.Node.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		nodes = append(nodes, &model.Node{Node: item})
	}
	return nodes, nil
}

func (r *NodeRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.NodeGetReq) ([]*model.Node, *cv1.PageReply, error) {
	var (
		nodes []*model.Node
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
		nodes = append(nodes, &model.Node{Node: item})
	}
	return nodes, &cv1.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *NodeRepo) getQuery(query *gen.NodeQuery, req *repo.NodeGetReq) *gen.NodeQuery {
	if req.Id != nil {
		query = query.Where(node.IDEQ(*req.Id))
	}
	if req.Ids != nil {
		query = query.Where(node.IDIn(req.Ids...))
	}
	if req.Key != nil {
		query = query.Where(node.KeyEQ(*req.Key))
	}
	if req.Name != nil {
		query = query.Where(node.NameEQ(*req.Name))
	}
	if req.OwnerId != nil {
		query = query.Where(node.OwnerIDEQ(*req.OwnerId))
	}
	return query
}
