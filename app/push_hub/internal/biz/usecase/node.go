package usecase

import (
	"context"
	"fmt"
	pushhubv1 "common/proto/gen/push_hub/v1"
	"push_hub/internal/biz/repo"

	"github.com/google/uuid"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NodeUsecase 节点管理业务逻辑。
type NodeUsecase struct {
	log      *log.Helper
	registry repo.NodeRegistry
}

func NewNodeUsecase(logger log.Logger, registry repo.NodeRegistry) *NodeUsecase {
	return &NodeUsecase{
		log:      log.NewHelper(logger),
		registry: registry,
	}
}

// RegisterNode 注册新的推送节点，生成唯一节点 ID。
func (uc *NodeUsecase) RegisterNode(ctx context.Context, address string) (string, error) {
	nodeID := uuid.New().String()
	if err := uc.registry.RegisterNode(ctx, nodeID, address); err != nil {
		return "", fmt.Errorf("注册节点: %w", err)
	}
	uc.log.Infof("节点已注册: %s (%s)", nodeID, address)
	return nodeID, nil
}

// Heartbeat 更新节点心跳和当前连接数。
func (uc *NodeUsecase) Heartbeat(ctx context.Context, nodeID string, connectionCount int64) error {
	if err := uc.registry.UpdateHeartbeat(ctx, nodeID, connectionCount); err != nil {
		return fmt.Errorf("更新心跳: %w", err)
	}
	return nil
}

// ListNodes 列出所有节点信息。
func (uc *NodeUsecase) ListNodes(ctx context.Context) ([]*pushhubv1.NodeInfo, error) {
	nodes, err := uc.registry.ListNodes(ctx)
	if err != nil {
		return nil, fmt.Errorf("列出节点: %w", err)
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
