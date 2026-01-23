package data

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"context"
	"encoding/json"
	"errors"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"signal/internal/data/ent/gen"
	"signal/internal/data/ent/gen/node"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type NodeRepo struct {
	*BaseRepo
}

func NewNodeRepo(baseRepo *BaseRepo) repo.NodeRepo {
	return &NodeRepo{
		BaseRepo: baseRepo,
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

func (r *NodeRepo) GetByKey(ctx context.Context, tx *gen.Client, key string) (*model.Node, error) {
	cacheKey := constant.GetKeySignalNode(key)

	data, err := r.redis.Client.HGet(ctx, cacheKey, constant.SignalNodeData).Result()
	if err == nil {
		var n model.Node
		if err := json.Unmarshal([]byte(data), &n); err != nil {
			return nil, err
		}
		return &n, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	n, err := tx.Node.Query().Where(node.KeyEQ(key)).Only(ctx)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(n)
	if err == nil {
		_ = r.redis.Client.HSet(ctx, cacheKey, map[string]interface{}{
			constant.SignalNodeData:               string(b),
			constant.SignalNodeCurrentConnections: 0,
			constant.SignalNodePingMs:             0,
			constant.SignalNodePowCostMs:          0,
			constant.SignalNodeLastPingTime:       time.Now().UnixMilli(),
		}).Err()
	}

	return &model.Node{Node: n}, nil
}

func (r *NodeRepo) Register(ctx context.Context, n *model.Node) error {
	// 初始化 HSet
	marshal, err := json.Marshal(n)
	if err != nil {
		return err
	}
	_, err = r.redis.Client.HSet(ctx, constant.GetKeySignalNode(n.Key), map[string]interface{}{
		constant.SignalNodeData:               string(marshal),
		constant.SignalNodeCurrentConnections: 0,
		constant.SignalNodePingMs:             0,
		constant.SignalNodePowCostMs:          0,
		constant.SignalNodeLastPingTime:       time.Now().UnixMilli(),
	}).Result()
	if err != nil {
		return err
	}

	// 初始化 ZSet 排名
	err = r.redis.Client.ZAdd(ctx, constant.SignalNodeRank, redis.Z{
		Score:  0,
		Member: n.Key,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *NodeRepo) Unregister(ctx context.Context, key string) error {
	pipe := r.redis.Client.TxPipeline()
	pipe.Del(ctx, constant.GetKeySignalNode(key))
	pipe.ZRem(ctx, constant.SignalNodeRank, key)
	_, err := pipe.Exec(ctx)
	return err
}

func (r *NodeRepo) UpdateConnections(ctx context.Context, key string, delta int64) error {
	redisKey := constant.GetKeySignalNode(key)
	pipe := r.redis.Client.TxPipeline()
	pipe.HIncrBy(ctx, redisKey, constant.SignalNodeCurrentConnections, delta)
	pipe.HSet(ctx, redisKey, constant.SignalNodeLastPingTime, time.Now().UnixMilli())
	_, err := pipe.Exec(ctx)
	return err
}

func (r *NodeRepo) UpdatePing(ctx context.Context, key string, pingMs int64) error {
	return r.redis.Client.HSet(ctx,
		constant.GetKeySignalNode(key),
		constant.SignalNodePingMs, pingMs,
		constant.SignalNodeLastPingTime, time.Now().UnixMilli(),
	).Err()
}

func (r *NodeRepo) UpdatePowCost(ctx context.Context, name string, powCostMs int64) error {
	return r.redis.Client.HSet(ctx,
		constant.GetKeySignalNode(name),
		constant.SignalNodePowCostMs, powCostMs,
		constant.SignalNodeLastPingTime, time.Now().UnixMilli(),
	).Err()
}

func (r *NodeRepo) UpdateScore(ctx context.Context, n *model.Node) error {
	result, err := r.redis.Client.HGetAll(ctx, constant.GetKeySignalNode(n.Key)).Result()
	if err != nil {
		return err
	}
	// 解析字段
	connections, _ := strconv.ParseInt(result[constant.SignalNodeCurrentConnections], 10, 64)
	pingMs, _ := strconv.ParseInt(result[constant.SignalNodePingMs], 10, 64)
	powCostMs, _ := strconv.ParseInt(result[constant.SignalNodePowCostMs], 10, 64)

	score := n.CalculateScore(connections, pingMs, powCostMs)

	err = r.redis.Client.ZAdd(ctx, constant.SignalNodeRank, redis.Z{
		Score:  score,
		Member: n.Key,
	}).Err()

	return err
}
