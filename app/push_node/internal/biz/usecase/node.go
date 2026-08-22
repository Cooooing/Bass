package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	pushhubv1 "common/proto/gen/push_hub/v1"
	"push_node/internal/biz/repo"
	"push_node/internal/config"
)

type NodeUsecase struct {
	conf       *config.Bootstrap
	log        *slog.Logger
	registry   repo.ConnectionRegistry
	nodeID     string
	hubClient  pushhubv1.PushHubNodeServiceClient
	cancelLoop context.CancelFunc
}

func NewNodeUsecase(
	conf *config.Bootstrap,
	logger *slog.Logger,
	registry repo.ConnectionRegistry,
	hubClient pushhubv1.PushHubNodeServiceClient,
) *NodeUsecase {
	return &NodeUsecase{
		conf:      conf,
		log:       logger,
		registry:  registry,
		hubClient: hubClient,
	}
}

func (uc *NodeUsecase) ConnectHub(ctx context.Context) error {
	nodeID, err := RegisterWithHub(ctx, uc.hubClient, uc.conf)
	if err != nil {
		return err
	}
	uc.nodeID = nodeID
	uc.log.Info(fmt.Sprintf("start heartbeat loop: node_id=%s", uc.nodeID))
	loopCtx, cancel := context.WithCancel(ctx)
	uc.cancelLoop = cancel
	go uc.heartbeatLoop(loopCtx)
	return nil
}

func (uc *NodeUsecase) Stop(ctx context.Context) error {
	_ = ctx
	if uc.cancelLoop != nil {
		uc.cancelLoop()
	}
	uc.log.Info(fmt.Sprintf("node stopped: node_id=%s", uc.nodeID))
	return nil
}

func (uc *NodeUsecase) heartbeatLoop(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	uc.sendHeartbeat(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			uc.sendHeartbeat(ctx)
		}
	}
}

func (uc *NodeUsecase) sendHeartbeat(ctx context.Context) {
	connectionCount, err := uc.registry.GetConnectionCount(ctx)
	if err != nil {
		uc.log.Warn(fmt.Sprintf("get connection count failed: err=%v", err))
		return
	}
	_, err = uc.hubClient.Heartbeat(ctx, &pushhubv1.Heartbeat_Req{
		NodeId:          uc.nodeID,
		ConnectionCount: connectionCount,
	})
	if err != nil {
		uc.log.Warn(fmt.Sprintf("send heartbeat failed: err=%v", err))
		return
	}
	uc.log.Debug(fmt.Sprintf("heartbeat sent: node_id=%s connections=%d", uc.nodeID, connectionCount))
}

func RegisterWithHub(ctx context.Context, client pushhubv1.PushHubNodeServiceClient, conf *config.Bootstrap) (string, error) {
	var nodeID string
	var lastErr error
	for i := 0; i < 3; i++ {
		resp, err := client.RegisterNode(ctx, &pushhubv1.RegisterNode_Req{
			Address: fmt.Sprintf("%s:%d", conf.Http.Host, conf.Http.Port),
		})
		if err != nil {
			lastErr = err
			time.Sleep(time.Second)
			continue
		}
		nodeID = resp.NodeId
		lastErr = nil
		break
	}
	if lastErr != nil {
		return "", fmt.Errorf("register node failed after 3 retries: %w", lastErr)
	}
	return nodeID, nil
}
