package usecase

import (
	"context"
	"fmt"

	pushhubv1 "common/proto/gen/push_hub/v1"
	"push_hub/internal/biz/repo"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

type NodeUsecase struct {
	log      *slog.Logger
	registry repo.NodeRegistry
}

func NewNodeUsecase(logger *slog.Logger, registry repo.NodeRegistry) *NodeUsecase {
	return &NodeUsecase{log: logger, registry: registry}
}

func (uc *NodeUsecase) RegisterNode(ctx context.Context, address string) (string, error) {
	nodeID := uuid.New().String()
	if err := uc.registry.RegisterNode(ctx, nodeID, address); err != nil {
		return "", fmt.Errorf("register node: %w", err)
	}
	uc.log.Info(fmt.Sprintf("node registered: node_id=%s address=%s", nodeID, address))
	return nodeID, nil
}

func (uc *NodeUsecase) Heartbeat(ctx context.Context, nodeID string, connectionCount int64) error {
	if err := uc.registry.UpdateHeartbeat(ctx, nodeID, connectionCount); err != nil {
		return fmt.Errorf("update heartbeat: %w", err)
	}
	return nil
}

func (uc *NodeUsecase) ListNodes(ctx context.Context) ([]*pushhubv1.NodeInfo, error) {
	nodes, err := uc.registry.ListNodes(ctx)
	if err != nil {
		return nil, fmt.Errorf("list nodes: %w", err)
	}

	result := make([]*pushhubv1.NodeInfo, 0, len(nodes))
	for _, n := range nodes {
		result = append(result, &pushhubv1.NodeInfo{
			NodeId:          n.NodeID,
			Address:         n.Address,
			ConnectionCount: n.ConnectionCount,
			Status:          pushhubv1.NodeStatus(n.Status),
			RegisteredAt:    timestamppb.New(n.RegisteredAt),
			LastHeartbeatAt: timestamppb.New(n.LastHeartbeatAt),
		})
	}
	return result, nil
}
