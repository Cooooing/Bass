package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"common/pkg/client"
	"common/pkg/constant"
	"push_hub/internal/biz/model"
	"push_hub/internal/biz/repo"
)

type PushEventUsecase struct {
	log      *slog.Logger
	registry repo.NodeRegistry
	natsPub  client.Publisher
}

func NewPushEventUsecase(
	logger *slog.Logger,
	registry repo.NodeRegistry,
	natsPub client.Publisher,
) *PushEventUsecase {
	return &PushEventUsecase{
		log:      logger,
		registry: registry,
		natsPub:  natsPub,
	}
}

type PublishEventReq struct {
	EventType int32
	UserID    int64
	Payload   string
}

func (uc *PushEventUsecase) PublishEvent(ctx context.Context, req *PublishEventReq) error {
	nodeIDs, err := uc.registry.GetUserNodes(ctx, req.UserID)
	if err != nil {
		return fmt.Errorf("get user nodes: %w", err)
	}

	if len(nodeIDs) == 0 {
		event := &model.PushEvent{
			Type:      req.EventType,
			UserID:    req.UserID,
			Payload:   req.Payload,
			CreatedAt: time.Now(),
		}
		if err := uc.registry.SaveOfflineEvent(ctx, &repo.SaveOfflineEventReq{
			UserID: req.UserID,
			Event:  event,
		}); err != nil {
			return fmt.Errorf("save offline event: %w", err)
		}
		uc.log.Debug(fmt.Sprintf("user has no online node, offline event saved: user_id=%d", req.UserID))
		return nil
	}

	msg := &model.PushEvent{
		Type:      req.EventType,
		UserID:    req.UserID,
		Payload:   req.Payload,
		CreatedAt: time.Now(),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal push event: %w", err)
	}

	for _, nodeID := range nodeIDs {
		subject := constant.GetPushNodeSubject(nodeID)
		if err := uc.natsPub.Publish(ctx, subject, &client.Message{
			Data: data,
		}); err != nil {
			uc.log.Warn(fmt.Sprintf("publish to node failed: node_id=%s err=%v", nodeID, err))
		}
	}
	return nil
}

type BroadcastEventReq struct {
	EventType int32
	Payload   string
}

func (uc *PushEventUsecase) BroadcastEvent(ctx context.Context, req *BroadcastEventReq) error {
	nodes, err := uc.registry.GetAllOnlineNodes(ctx)
	if err != nil {
		return fmt.Errorf("get online nodes: %w", err)
	}

	msg := &model.PushEvent{
		Type:      req.EventType,
		Payload:   req.Payload,
		CreatedAt: time.Now(),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal push event: %w", err)
	}

	for _, node := range nodes {
		subject := constant.GetPushNodeSubject(node.NodeID)
		if err := uc.natsPub.Publish(ctx, subject, &client.Message{
			Data: data,
		}); err != nil {
			uc.log.Warn(fmt.Sprintf("broadcast to node failed: node_id=%s err=%v", node.NodeID, err))
		}
	}
	return nil
}
