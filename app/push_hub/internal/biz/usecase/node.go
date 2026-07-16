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

func NewNodeUsecase(logger *slog.Logger, registry repo.NodeRegistry) *NodeUsecase {
	return &NodeUsecase{log: logger, registry: registry}
}

type RegisterNodeReq struct {
	Address string
}

type RegisterNodeResponse struct {
	NodeID string
}

func (uc *NodeUsecase) RegisterNode(ctx context.Context, req *RegisterNodeReq) (*RegisterNodeResponse, error) {
	nodeID := uuid.New().String()
	if _, err := uc.registry.RegisterNode(ctx, &repo.RegisterNodeReq{NodeID: nodeID, Address: req.Address}); err != nil {
		return nil, fmt.Errorf("register node: %w", err)
	}
	uc.log.Info(fmt.Sprintf("node registered: node_id=%s address=%s", nodeID, req.Address))
	return &RegisterNodeResponse{NodeID: nodeID}, nil
}

type HeartbeatReq struct {
	NodeID          string
	ConnectionCount int64
}

func (uc *NodeUsecase) Heartbeat(ctx context.Context, req *HeartbeatReq) error {
	if _, err := uc.registry.UpdateHeartbeat(ctx, &repo.UpdateHeartbeatReq{NodeID: req.NodeID, ConnectionCount: req.ConnectionCount}); err != nil {
		return fmt.Errorf("update heartbeat: %w", err)
	}
	return nil
}

type ListNodesReq struct{}

type ListNodesResponse struct {
	Rows []*model.NodeInfo
}

func (uc *NodeUsecase) ListNodes(ctx context.Context, req *ListNodesReq) (*ListNodesResponse, error) {
	nodesResp, err := uc.registry.ListNodes(ctx, &repo.ListNodesReq{})
	if err != nil {
		return nil, fmt.Errorf("list nodes: %w", err)
	}
	return &ListNodesResponse{Rows: nodesResp.Rows}, nil
}
