package service

import (
	v1 "common/api/signal/v1"
	commonModel "common/pkg/model"
	"context"
	"fmt"
	"signal/internal/biz"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"signal/internal/data"
	"signal/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NodeService struct {
	v1.UnsafeSignalNodeServiceServer
	*BaseService
	nodeDomain *biz.NodeDomain
	nodeRepo   repo.NodeRepo
}

func NewNodeService(baseService *BaseService, nodeDomain *biz.NodeDomain, nodeRepo repo.NodeRepo) *NodeService {
	return &NodeService{
		BaseService: baseService,
		nodeDomain:  nodeDomain,
		nodeRepo:    nodeRepo,
	}
}

func (s *NodeService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterSignalNodeServiceServer(gs, s)
}

func (s *NodeService) RegisterHttp(hs *http.Server) {
	v1.RegisterSignalNodeServiceHTTPServer(hs, s)
}

func (s *NodeService) Save(ctx context.Context, req *v1.SignalNodeSaveRequest) (rsp *v1.SignalNodeSaveReply, err error) {
	if req.Node == nil {
		return nil, fmt.Errorf("node is nil")
	}
	save, err := s.nodeRepo.Save(ctx, s.db, &model.Node{Node: &gen.Node{
		OwnerID:     req.Node.OwnerId,
		Name:        req.Node.Name,
		Description: req.Node.Description,
		Secret:      s.nodeDomain.GenerateSecret(),
		CallbackURL: req.Node.CallbackUrl,
		Status:      req.Node.Status,
	}})
	if err != nil {
		return nil, err
	}
	return &v1.SignalNodeSaveReply{Node: save.ConvertToRpc()}, nil
}

func (s *NodeService) Update(ctx context.Context, req *v1.SignalNodeUpdateRequest) (*v1.SignalNodeUpdateReply, error) {
	if req.Node == nil {
		return nil, fmt.Errorf("node is nil")
	}
	update, err := s.nodeRepo.Update(ctx, s.db, &model.Node{Node: &gen.Node{
		ID:          req.Node.Id,
		OwnerID:     req.Node.OwnerId,
		Name:        req.Node.Name,
		Description: req.Node.Description,
		CallbackURL: req.Node.CallbackUrl,
		Status:      req.Node.Status,
	}})
	if err != nil {
		return nil, err
	}
	return &v1.SignalNodeUpdateReply{Node: update.ConvertToRpc()}, nil
}

func (s *NodeService) GetSecret(ctx context.Context, req *v1.SignalNodeGetSecretRequest) (*v1.SignalNodeGetSecretReply, error) {
	node, err := s.nodeRepo.GetOne(ctx, s.db, &repo.NodeGetReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.SignalNodeGetSecretReply{Secret: node.Secret}, nil
}

func (s *NodeService) UpdateSecret(ctx context.Context, req *v1.SignalNodeUpdateSecretRequest) (*v1.SignalNodeUpdateSecretReply, error) {
	err := s.nodeRepo.UpdateSecret(ctx, s.db, req.Id, s.nodeDomain.GenerateSecret())
	return &v1.SignalNodeUpdateSecretReply{}, err
}

func (s *NodeService) List(ctx context.Context, req *v1.SignalNodeListRequest) (*v1.SignalNodeListReply, error) {
	list, err := s.nodeRepo.GetList(ctx, s.db, &repo.NodeGetReq{})
	if err != nil {
		return nil, err
	}
	return &v1.SignalNodeListReply{Nodes: commonModel.ConvertToRpcList(list)}, err
}

func (s *NodeService) Negotiate(ctx context.Context, req *v1.SignalNodeNegotiateRequest) (*v1.SignalNodeNegotiateReply, error) {
	// 拉取当前可用的节点列表，作为协商候选
	list, err := s.nodeRepo.GetList(ctx, s.db, &repo.NodeGetReq{})
	if err != nil {
		return nil, err
	}
	// 将 internal Node 转换为 RPC Node
	var rpcNodes []*v1.Node
	for _, n := range list {
		rpcNodes = append(rpcNodes, n.ConvertToRpc())
	}
	// 简单的最佳节点草案：优先选取第一个节点（后续可基于分数排序）
	var best *v1.Node
	if len(rpcNodes) > 0 {
		best = rpcNodes[0]
		// 计算示例分数，保持与领域逻辑的一致性，实际实现可使用真实指标
		_ = s.nodeDomain.CalculateNodeScore(1.0, 10, 10, int64(len(rpcNodes)))
	}
	return &v1.SignalNodeNegotiateReply{Node: best, Nodes: rpcNodes}, nil
}

func (s *NodeService) Register(ctx context.Context, req *v1.SignalNodeRegisterRequest) (*v1.SignalNodeRegisterReply, error) {
	// TODO: 实现注册节点的领域逻辑
	return &v1.SignalNodeRegisterReply{}, nil
}

func (s *NodeService) Unregister(ctx context.Context, req *v1.SignalNodeUnregisterRequest) (*v1.SignalNodeUnregisterReply, error) {
	// TODO: 实现注销节点的领域逻辑
	return &v1.SignalNodeUnregisterReply{}, nil
}

func (s *NodeService) Online(ctx context.Context, req *v1.SignalNodeOnlineRequest) (*v1.SignalNodeOnlineReply, error) {
	// TODO: 实现节点上线的领域逻辑
	return &v1.SignalNodeOnlineReply{}, nil
}

func (s *NodeService) Offline(ctx context.Context, req *v1.SignalNodeOfflineRequest) (*v1.SignalNodeOfflineReply, error) {
	// TODO: 实现节点下线的领域逻辑
	return &v1.SignalNodeOfflineReply{}, nil
}

func (s *NodeService) OnlineList(ctx context.Context, req *v1.SignalNodeOnlineListRequest) (*v1.SignalNodeOnlineListReply, error) {
	// TODO: 实现在线节点列表的领域逻辑
	return &v1.SignalNodeOnlineListReply{}, nil
}
