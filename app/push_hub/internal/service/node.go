package service

import (
	pushhubv1 "common/proto/gen/push_hub/v1"
	"context"
	"push_hub/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
)

// NodeService 实现 PushHubNodeServiceServer，委托给 usecase 层。
type NodeService struct {
	pushhubv1.UnimplementedPushHubNodeServiceServer
	nodeUc *usecase.NodeUsecase
}

func NewNodeService(nodeUc *usecase.NodeUsecase) *NodeService {
	return &NodeService{nodeUc: nodeUc}
}

func (s *NodeService) RegisterGrpc(gs *grpc.Server) {
	pushhubv1.RegisterPushHubNodeServiceServer(gs, s)
}

func (s *NodeService) RegisterNode(ctx context.Context, req *pushhubv1.RegisterNode_Request) (*pushhubv1.RegisterNode_Reply, error) {
	nodeID, err := s.nodeUc.RegisterNode(ctx, req.Address)
	if err != nil {
		return nil, err
	}
	return &pushhubv1.RegisterNode_Reply{NodeId: nodeID}, nil
}

func (s *NodeService) Heartbeat(ctx context.Context, req *pushhubv1.Heartbeat_Request) (*pushhubv1.Heartbeat_Reply, error) {
	if err := s.nodeUc.Heartbeat(ctx, req.NodeId, req.ConnectionCount); err != nil {
		return nil, err
	}
	return &pushhubv1.Heartbeat_Reply{}, nil
}

func (s *NodeService) ListNodes(ctx context.Context, req *pushhubv1.ListNodes_Request) (*pushhubv1.ListNodes_Reply, error) {
	nodes, err := s.nodeUc.ListNodes(ctx)
	if err != nil {
		return nil, err
	}
	return &pushhubv1.ListNodes_Reply{Nodes: nodes}, nil
}
