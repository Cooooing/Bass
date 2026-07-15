package usecase

import (
	"context"
	"fmt"
	"time"

	pushhubv1 "common/proto/gen/push_hub/v1"
	"push_node/internal/biz/repo"
	"push_node/internal/config"

	"google.golang.org/grpc"
	"log/slog"
)

type NodeUsecase struct {
	conf       *config.Bootstrap
	log        *slog.Logger
	registry   repo.ConnectionRegistry
	nodeID     string
	hubClient  pushhubv1.PushHubNodeServiceClient
	hubConn    *grpc.ClientConn
	cancelLoop context.CancelFunc
}

func NewNodeUsecase(conf *config.Bootstrap, logger *slog.Logger, registry repo.ConnectionRegistry, nodeID string, hubConn *grpc.ClientConn) *NodeUsecase {
	return &NodeUsecase{
		conf:      conf,
		log:       logger,
		registry:  registry,
		nodeID:    nodeID,
		hubConn:   hubConn,
		hubClient: pushhubv1.NewPushHubNodeServiceClient(hubConn),
	}
}

func (uc *NodeUsecase) ConnectHub(ctx context.Context) error {
	uc.log.Info(fmt.Sprintf("start heartbeat loop: node_id=%s", uc.nodeID))
	loopCtx, cancel := context.WithCancel(ctx)
	uc.cancelLoop = cancel
	go uc.heartbeatLoop(loopCtx)
	return nil
}

func (uc *NodeUsecase) Stop() {
	if uc.cancelLoop != nil {
		uc.cancelLoop()
	}
	uc.log.Info(fmt.Sprintf("node stopped: node_id=%s", uc.nodeID))
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
	connectionCount := uc.registry.GetConnectionCount()
	_, err := uc.hubClient.Heartbeat(ctx, &pushhubv1.Heartbeat_Request{
		NodeId:          uc.nodeID,
		ConnectionCount: connectionCount,
	})
	if err != nil {
		uc.log.Warn(fmt.Sprintf("send heartbeat failed: err=%v", err))
		return
	}
	uc.log.Debug(fmt.Sprintf("heartbeat sent: node_id=%s connections=%d", uc.nodeID, connectionCount))
}

func RegisterWithHub(ctx context.Context, conn *grpc.ClientConn, conf *config.Bootstrap) (string, error) {
	client := pushhubv1.NewPushHubNodeServiceClient(conn)
	var nodeID string
	var lastErr error
	for i := 0; i < 3; i++ {
		resp, err := client.RegisterNode(ctx, &pushhubv1.RegisterNode_Request{
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
