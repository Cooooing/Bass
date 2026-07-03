package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"common/pkg/client"
	"push_hub/internal/biz/model"
	"push_hub/internal/biz/repo"

	"log/slog"
)

type PushEventUsecase struct {
	log      *slog.Logger
	registry repo.NodeRegistry
	natsPub  client.Publisher
}

func NewPushEventUsecase(logger *slog.Logger, registry repo.NodeRegistry, natsPub client.Publisher) *PushEventUsecase {
	return &PushEventUsecase{log: logger, registry: registry, natsPub: natsPub}
}

func (uc *PushEventUsecase) PublishEvent(ctx context.Context, eventType int32, userID int64, payload string) error {
	nodeIDs, err := uc.registry.GetUserNodes(ctx, userID)
	if err != nil {
		return fmt.Errorf("get user nodes: %w", err)
	}

	if len(nodeIDs) == 0 {
		event := &model.PushEvent{Type: eventType, UserID: userID, Payload: payload, CreatedAt: time.Now()}
		if err := uc.registry.SaveOfflineEvent(ctx, userID, event); err != nil {
			return fmt.Errorf("save offline event: %w", err)
		}
		uc.log.Debug(fmt.Sprintf("user has no online node, offline event saved: user_id=%d", userID))
		return nil
	}

	msg := &model.PushEvent{Type: eventType, UserID: userID, Payload: payload, CreatedAt: time.Now()}
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal push event: %w", err)
	}

	for _, nodeID := range nodeIDs {
		subject := fmt.Sprintf("push.node.%s", nodeID)
		if err := uc.natsPub.Publish(ctx, subject, &client.Message{Data: data}); err != nil {
			uc.log.Warn(fmt.Sprintf("publish to node failed: node_id=%s err=%v", nodeID, err))
		}
	}
	return nil
}

func (uc *PushEventUsecase) BroadcastEvent(ctx context.Context, eventType int32, payload string) error {
	nodes, err := uc.registry.GetAllOnlineNodes(ctx)
	if err != nil {
		return fmt.Errorf("get online nodes: %w", err)
	}

	msg := &model.PushEvent{Type: eventType, Payload: payload, CreatedAt: time.Now()}
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal push event: %w", err)
	}

	for _, node := range nodes {
		subject := fmt.Sprintf("push.node.%s", node.NodeID)
		if err := uc.natsPub.Publish(ctx, subject, &client.Message{Data: data}); err != nil {
			uc.log.Warn(fmt.Sprintf("broadcast to node failed: node_id=%s err=%v", node.NodeID, err))
		}
	}
	return nil
}
