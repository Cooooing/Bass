package usecase

import (
	"context"
	"fmt"
	"time"

	pushhubv1 "common/proto/gen/push_hub/v1"
	"push_node/internal/biz/repo"
	"push_node/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
)

// NodeUsecase 推送节点生命周期管理业务逻辑。
type NodeUsecase struct {
	conf       *conf.Bootstrap
	log        *log.Helper
	registry   repo.ConnectionRegistry
	nodeID     string
	hubClient  pushhubv1.PushHubNodeServiceClient
	hubConn    *grpc.ClientConn
	cancelLoop context.CancelFunc
}

// NewNodeUsecase 创建 NodeUsecase 实例。
func NewNodeUsecase(
	conf *conf.Bootstrap,
	logger log.Logger,
	registry repo.ConnectionRegistry,
	nodeID string,
	hubConn *grpc.ClientConn,
) *NodeUsecase {
	return &NodeUsecase{
		conf:      conf,
		log:       log.NewHelper(logger),
		registry:  registry,
		nodeID:    nodeID,
		hubConn:   hubConn,
		hubClient: pushhubv1.NewPushHubNodeServiceClient(hubConn),
	}
}

// ConnectHub 启动心跳循环（在应用 AfterStart 时调用）。
func (uc *NodeUsecase) ConnectHub(ctx context.Context) error {
	uc.log.Infof("开始心跳循环: node=%s", uc.nodeID)

	loopCtx, cancel := context.WithCancel(ctx)
	uc.cancelLoop = cancel
	go uc.heartbeatLoop(loopCtx)

	return nil
}

// Stop 停止心跳循环。
func (uc *NodeUsecase) Stop() {
	if uc.cancelLoop != nil {
		uc.cancelLoop()
	}
	uc.log.Infof("节点已停止: %s", uc.nodeID)
}

// heartbeatLoop 定期向 push_hub 发送心跳。
func (uc *NodeUsecase) heartbeatLoop(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	// 启动时立即发送一次心跳
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

// sendHeartbeat 向 push_hub 发送当前心跳状态。
func (uc *NodeUsecase) sendHeartbeat(ctx context.Context) {
	_, err := uc.hubClient.Heartbeat(ctx, &pushhubv1.Heartbeat_Request{
		NodeId:          uc.nodeID,
		ConnectionCount: uc.registry.GetConnectionCount(),
	})
	if err != nil {
		uc.log.Warnf("心跳发送失败: %v", err)
	} else {
		uc.log.Debugf("心跳已发送: node=%s connections=%d", uc.nodeID, uc.registry.GetConnectionCount())
	}
}

// RegisterWithHub 向 push_hub 注册节点，返回节点 ID。
// 在 Wire 初始化之前调用，用于获取节点标识。
func RegisterWithHub(ctx context.Context, conn *grpc.ClientConn, conf *conf.Bootstrap) (string, error) {
	client := pushhubv1.NewPushHubNodeServiceClient(conn)

	var nodeID string
	var lastErr error
	for i := 0; i < 3; i++ {
		resp, err := client.RegisterNode(ctx, &pushhubv1.RegisterNode_Request{
			Address: fmt.Sprintf("%s:%d", conf.Server.Http.Host, conf.Server.Http.Port),
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
		return "", fmt.Errorf("注册节点失败（已重试 3 次）: %w", lastErr)
	}
	return nodeID, nil
}
