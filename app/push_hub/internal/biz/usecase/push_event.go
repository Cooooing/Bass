package usecase

import (
	"common/pkg/util"
	"context"
	"encoding/json"
	"fmt"
	"push_hub/internal/biz/model"
	"push_hub/internal/biz/repo"
	"time"

	"common/pkg/client"
	"log/slog"
)

// PushEventUsecase 推送事件业务逻辑。
type PushEventUsecase struct {
	log      *util.LogHelper
	registry repo.NodeRegistry
	natsPub  client.Publisher
}

func NewPushEventUsecase(logger *slog.Logger, registry repo.NodeRegistry, natsPub client.Publisher) *PushEventUsecase {
	return &PushEventUsecase{
		log:      util.NewLogHelper(logger),
		registry: registry,
		natsPub:  natsPub,
	}
}

// PublishEvent 单播推送事件到指定用户。
func (uc *PushEventUsecase) PublishEvent(ctx context.Context, eventType int32, userID int64, payload string) error {
	nodeIDs, err := uc.registry.GetUserNodes(ctx, userID)
	if err != nil {
		return fmt.Errorf("获取用户节点: %w", err)
	}

	// 用户无在线节点，存储离线事件
	if len(nodeIDs) == 0 {
		event := &model.PushEvent{
			Type:      eventType,
			UserID:    userID,
			Payload:   payload,
			CreatedAt: time.Now(),
		}
		if err := uc.registry.SaveOfflineEvent(ctx, userID, event); err != nil {
			return fmt.Errorf("保存离线事件: %w", err)
		}
		uc.log.Debugf("用户 %d 无在线节点，事件已存为离线", userID)
		return nil
	}

	// 构建推送消息
	msg := &model.PushEvent{
		Type:      eventType,
		UserID:    userID,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("序列化事件: %w", err)
	}

	// 发送到每个关联节点
	for _, nodeID := range nodeIDs {
		subject := fmt.Sprintf("push.node.%s", nodeID)
		if err := uc.natsPub.Publish(ctx, subject, &client.Message{Data: data}); err != nil {
			uc.log.Warnf("发布到节点 %s 失败: %v", nodeID, err)
		}
	}

	return nil
}

// BroadcastEvent 广播推送事件到所有在线节点。
func (uc *PushEventUsecase) BroadcastEvent(ctx context.Context, eventType int32, payload string) error {
	nodes, err := uc.registry.GetAllOnlineNodes(ctx)
	if err != nil {
		return fmt.Errorf("获取在线节点: %w", err)
	}

	msg := &model.PushEvent{
		Type:      eventType,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("序列化事件: %w", err)
	}

	for _, node := range nodes {
		subject := fmt.Sprintf("push.node.%s", node.NodeID)
		if err := uc.natsPub.Publish(ctx, subject, &client.Message{Data: data}); err != nil {
			uc.log.Warnf("广播到节点 %s 失败: %v", node.NodeID, err)
		}
	}

	return nil
}
