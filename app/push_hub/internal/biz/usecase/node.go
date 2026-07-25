package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"push_hub/internal/biz/model"
	"push_hub/internal/biz/repo"

	"github.com/google/uuid"
)

type NodeUsecase struct {
	log      *slog.Logger
	registry repo.NodeRegistry
}

func NewNodeUsecase(
	logger *slog.Logger,
	registry repo.NodeRegistry,
) *NodeUsecase {
	return &NodeUsecase{
		log:      logger,
		registry: registry,
	}
}

func (uc *NodeUsecase) RegisterNode(
	ctx context.Context,
	address string,
) (string, error) {
	nodeID := uuid.New().String()
	if err := uc.registry.RegisterNode(ctx, &repo.RegisterNodeReq{
		NodeID:  nodeID,
		Address: address,
	}); err != nil {
		return "", fmt.Errorf("register node: %w", err)
	}
	uc.log.Info(fmt.Sprintf("node registered: node_id=%s address=%s", nodeID, address))
	return nodeID, nil
}

type HeartbeatReq struct {
	NodeID          string
	ConnectionCount int64
}

func (uc *NodeUsecase) Heartbeat(
	ctx context.Context,
	req *HeartbeatReq,
) error {
	if err := uc.registry.UpdateHeartbeat(ctx, &repo.UpdateHeartbeatReq{
		NodeID:          req.NodeID,
		ConnectionCount: req.ConnectionCount,
	}); err != nil {
		return fmt.Errorf("update heartbeat: %w", err)
	}
	return nil
}

func (uc *NodeUsecase) ListNodes(
	ctx context.Context,
) ([]*model.NodeInfo, error) {
	rows, err := uc.registry.ListNodes(ctx)
	if err != nil {
		return nil, fmt.Errorf("list nodes: %w", err)
	}
	return rows, nil
}
