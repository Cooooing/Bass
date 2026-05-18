package service

import (
	v1 "common/api/gen/signal/v1"
	commonModel "common/pkg/model"
	"context"
	"fmt"
	"signal/internal/biz/domain"
	"signal/internal/biz/model"
	"signal/internal/biz/repo"
	"signal/internal/data/gen"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NodeService struct {
	v1.UnsafeSignalNodeServiceServer
	db         *gen.Client
	nodeDomain *domain.NodeDomain
	nodeRepo   repo.NodeRepo
}

func NewNodeService(db *gen.Client, nodeDomain *domain.NodeDomain,
	nodeRepo repo.NodeRepo) *NodeService {
	return &NodeService{
		db:         db,
		nodeDomain: nodeDomain,
		nodeRepo:   nodeRepo,
	}
}

func (s *NodeService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterSignalNodeServiceServer(gs, s)
}

func (s *NodeService) RegisterHttp(hs *http.Server) {
	v1.RegisterSignalNodeServiceHTTPServer(hs, s)
}

func (s *NodeService) Save(ctx context.Context, req *v1.SaveSignalNode_Request) (rsp *v1.SaveSignalNode_Reply, err error) {
	if req.Node == nil {
		return nil, fmt.Errorf("node is nil")
	}
	save, err := s.nodeRepo.Save(ctx, s.db, &model.Node{Node: &gen.Node{
		OwnerID:     req.Node.OwnerId,
		Name:        req.Node.Name,
		Key:         req.Node.Key,
		Description: req.Node.Description,
		Secret:      s.nodeDomain.GenerateSecret(),
		CallbackURL: req.Node.CallbackUrl,
		Status:      int32(req.Node.Status),
	}})
	if err != nil {
		return nil, err
	}
	return &v1.SaveSignalNode_Reply{Node: save.ConvertToRpc()}, nil
}

func (s *NodeService) Update(ctx context.Context, req *v1.UpdateSignalNode_Request) (*v1.UpdateSignalNode_Reply, error) {
	if req.Node == nil {
		return nil, fmt.Errorf("node is nil")
	}
	update, err := s.nodeRepo.Update(ctx, s.db, &model.Node{Node: &gen.Node{
		ID:          req.Node.Id,
		OwnerID:     req.Node.OwnerId,
		Key:         req.Node.Key,
		Name:        req.Node.Name,
		Description: req.Node.Description,
		CallbackURL: req.Node.CallbackUrl,
		Status:      int32(req.Node.Status),
	}})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateSignalNode_Reply{Node: update.ConvertToRpc()}, nil
}

func (s *NodeService) GetSecret(ctx context.Context, req *v1.GetSecretSignalNode_Request) (*v1.GetSecretSignalNode_Reply, error) {
	node, err := s.nodeRepo.GetOne(ctx, s.db, &repo.NodeGetReq{Id: new(req.Id)})
	if err != nil {
		return nil, err
	}
	return &v1.GetSecretSignalNode_Reply{Secret: node.Secret}, nil
}

func (s *NodeService) UpdateSecret(ctx context.Context, req *v1.UpdateSecretSignalNode_Request) (*v1.UpdateSecretSignalNode_Reply, error) {
	err := s.nodeRepo.UpdateSecret(ctx, s.db, req.Id, s.nodeDomain.GenerateSecret())
	return &v1.UpdateSecretSignalNode_Reply{}, err
}

func (s *NodeService) List(ctx context.Context, req *v1.ListSignalNode_Request) (*v1.ListSignalNode_Reply, error) {
	list, err := s.nodeRepo.GetList(ctx, s.db, &repo.NodeGetReq{})
	if err != nil {
		return nil, err
	}
	return &v1.ListSignalNode_Reply{Nodes: commonModel.ConvertToRpcList(list)}, err
}

func (s *NodeService) Negotiate(ctx context.Context, req *v1.NegotiateSignalNode_Request) (*v1.NegotiateSignalNode_Reply, error) {
	nodes, err := s.nodeDomain.Negotiate(ctx)
	if err != nil {
		return nil, err
	}
	ticket, err := s.nodeDomain.Ticket(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.NegotiateSignalNode_Reply{
		Ticket: ticket,
		Nodes:  commonModel.ConvertToRpcList(nodes),
	}, nil
}

func (s *NodeService) Register(ctx context.Context, req *v1.RegisterSignalNode_Request) (*v1.RegisterSignalNode_Reply, error) {
	err := s.nodeDomain.Register(ctx)
	return &v1.RegisterSignalNode_Reply{}, err
}

func (s *NodeService) Unregister(ctx context.Context, req *v1.UnregisterSignalNode_Request) (*v1.UnregisterSignalNode_Reply, error) {
	err := s.nodeDomain.Unregister(ctx, "")
	return &v1.UnregisterSignalNode_Reply{}, err
}

func (s *NodeService) Ticket(ctx context.Context, req *v1.TicketSignalNode_Request) (*v1.TicketSignalNode_Reply, error) {
	ticket, err := s.nodeDomain.Ticket(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.TicketSignalNode_Reply{Ticket: ticket}, nil
}

func (s *NodeService) Online(ctx context.Context, req *v1.OnlineSignalNode_Request) (*v1.OnlineSignalNode_Reply, error) {
	sessionId, err := s.nodeDomain.Online(ctx, req.Ticket)
	if err != nil {
		return nil, err
	}
	return &v1.OnlineSignalNode_Reply{
		SessionId: sessionId,
	}, nil
}

func (s *NodeService) Offline(ctx context.Context, req *v1.OfflineSignalNode_Request) (*v1.OfflineSignalNode_Reply, error) {

	return &v1.OfflineSignalNode_Reply{}, nil
}

func (s *NodeService) OnlineList(ctx context.Context, req *v1.OnlineListSignalNode_Request) (*v1.OnlineListSignalNode_Reply, error) {
	// TODO: 实现在线节点列表的领域逻辑
	return &v1.OnlineListSignalNode_Reply{}, nil
}
