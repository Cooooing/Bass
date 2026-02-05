package service

import (
	v1 "common/api/signal/v1"
	"common/pkg/cutil/base"
	commonModel "common/pkg/model"
	"context"
	"fmt"
	"signal/internal/biz/domain"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"signal/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NodeService struct {
	v1.UnsafeSignalNodeServiceServer
	*BaseService
	nodeDomain *domain.NodeDomain
	nodeRepo   repo.NodeRepo
}

func NewNodeService(baseService *BaseService, nodeDomain *domain.NodeDomain, nodeRepo repo.NodeRepo) *NodeService {
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
	save, err := s.nodeRepo.Save(ctx, s.Db, &model.Node{Node: &gen.Node{
		OwnerID:     req.Node.OwnerId,
		Name:        req.Node.Name,
		Key:         req.Node.Key,
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
	update, err := s.nodeRepo.Update(ctx, s.Db, &model.Node{Node: &gen.Node{
		ID:          req.Node.Id,
		OwnerID:     req.Node.OwnerId,
		Key:         req.Node.Key,
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
	node, err := s.nodeRepo.GetOne(ctx, s.Db, &repo.NodeGetReq{Id: base.Ptr(req.Id)})
	if err != nil {
		return nil, err
	}
	return &v1.SignalNodeGetSecretReply{Secret: node.Secret}, nil
}

func (s *NodeService) UpdateSecret(ctx context.Context, req *v1.SignalNodeUpdateSecretRequest) (*v1.SignalNodeUpdateSecretReply, error) {
	err := s.nodeRepo.UpdateSecret(ctx, s.Db, req.Id, s.nodeDomain.GenerateSecret())
	return &v1.SignalNodeUpdateSecretReply{}, err
}

func (s *NodeService) List(ctx context.Context, req *v1.SignalNodeListRequest) (*v1.SignalNodeListReply, error) {
	list, err := s.nodeRepo.GetList(ctx, s.Db, &repo.NodeGetReq{})
	if err != nil {
		return nil, err
	}
	return &v1.SignalNodeListReply{Nodes: commonModel.ConvertToRpcList(list)}, err
}

func (s *NodeService) Negotiate(ctx context.Context, req *v1.SignalNodeNegotiateRequest) (*v1.SignalNodeNegotiateReply, error) {
	nodes, err := s.nodeDomain.Negotiate(ctx)
	if err != nil {
		return nil, err
	}
	ticket, err := s.nodeDomain.Ticket(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.SignalNodeNegotiateReply{
		Ticket: ticket,
		Nodes:  commonModel.ConvertToRpcList(nodes),
	}, nil
}

func (s *NodeService) Register(ctx context.Context, req *v1.SignalNodeRegisterRequest) (*v1.SignalNodeRegisterReply, error) {
	err := s.nodeDomain.Register(ctx)
	return &v1.SignalNodeRegisterReply{}, err
}

func (s *NodeService) Unregister(ctx context.Context, req *v1.SignalNodeUnregisterRequest) (*v1.SignalNodeUnregisterReply, error) {
	err := s.nodeDomain.Unregister(ctx, "")
	return &v1.SignalNodeUnregisterReply{}, err
}

func (s *NodeService) Ticket(ctx context.Context, req *v1.SignalNodeTicketRequest) (*v1.SignalNodeTicketReply, error) {
	ticket, err := s.nodeDomain.Ticket(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.SignalNodeTicketReply{Ticket: ticket}, nil
}

func (s *NodeService) Online(ctx context.Context, req *v1.SignalNodeOnlineRequest) (*v1.SignalNodeOnlineReply, error) {
	sessionId, err := s.nodeDomain.Online(ctx, req.Ticket)
	if err != nil {
		return nil, err
	}
	return &v1.SignalNodeOnlineReply{
		SessionId: sessionId,
	}, nil
}

func (s *NodeService) Offline(ctx context.Context, req *v1.SignalNodeOfflineRequest) (*v1.SignalNodeOfflineReply, error) {

	return &v1.SignalNodeOfflineReply{}, nil
}

func (s *NodeService) OnlineList(ctx context.Context, req *v1.SignalNodeOnlineListRequest) (*v1.SignalNodeOnlineListReply, error) {
	// TODO: 实现在线节点列表的领域逻辑
	return &v1.SignalNodeOnlineListReply{}, nil
}
