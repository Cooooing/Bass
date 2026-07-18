package service

import (
	"context"

	pushhubv1 "common/proto/gen/push_hub/v1"
	"push_hub/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *NodeService) RegisterHttp(hs *http.Server) {}

func (s *NodeService) RegisterNode(ctx context.Context, req *pushhubv1.RegisterNode_Req) (*pushhubv1.RegisterNode_Resp, error) {
	nodeID, err := s.nodeUc.RegisterNode(ctx, req.Address)
	if err != nil {
		return nil, err
	}
	return &pushhubv1.RegisterNode_Resp{NodeId: nodeID}, nil
}

func (s *NodeService) Heartbeat(ctx context.Context, req *pushhubv1.Heartbeat_Req) (*pushhubv1.Heartbeat_Resp, error) {
	if err := s.nodeUc.Heartbeat(ctx, &usecase.HeartbeatReq{
		NodeID:          req.NodeId,
		ConnectionCount: req.ConnectionCount,
	}); err != nil {
		return nil, err
	}
	return &pushhubv1.Heartbeat_Resp{}, nil
}

func (s *NodeService) ListNodes(ctx context.Context, req *pushhubv1.ListNodes_Req) (*pushhubv1.ListNodes_Resp, error) {
	rows, err := s.nodeUc.ListNodes(ctx)
	if err != nil {
		return nil, err
	}
	replyRows := make([]*pushhubv1.ListNodes_Resp_NodeInfo, 0, len(rows))
	for _, node := range rows {
		replyRows = append(replyRows, &pushhubv1.ListNodes_Resp_NodeInfo{
			NodeId:          node.NodeID,
			Address:         node.Address,
			ConnectionCount: node.ConnectionCount,
			Status:          pushhubv1.NodeStatus(node.Status),
			RegisteredAt:    timestamppb.New(node.RegisteredAt),
			LastHeartbeatAt: timestamppb.New(node.LastHeartbeatAt),
		})
	}
	return &pushhubv1.ListNodes_Resp{Rows: replyRows}, nil
}
