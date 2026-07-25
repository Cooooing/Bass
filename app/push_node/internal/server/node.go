package server

import (
	"context"
	"fmt"
	"log/slog"

	"push_node/internal/biz/usecase"
	"push_node/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NodeServer struct {
	nodeUsecase *usecase.NodeUsecase
}

func NewPushHubConn(
	conf *config.Bootstrap,
) (*grpc.ClientConn, func(), error) {
	conn, err := grpc.NewClient(conf.PushNode.PushHubAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("create push_hub grpc client failed: %w", err)
	}
	cleanup := func() {
		_ = conn.Close()
	}
	return conn, cleanup, nil
}

func ProvideNodeID(conf *config.Bootstrap, conn *grpc.ClientConn) (string, error) {
	nodeID, err := usecase.RegisterWithHub(context.Background(), conn, conf)
	if err != nil {
		return "", err
	}
	slog.Info("node registered", "node_id", nodeID)
	return nodeID, nil
}

func NewNodeServer(
	nodeUsecase *usecase.NodeUsecase,
) *NodeServer {
	return &NodeServer{
		nodeUsecase: nodeUsecase,
	}
}

func (s *NodeServer) Start(ctx context.Context) error {
	return s.nodeUsecase.ConnectHub(ctx)
}

func (s *NodeServer) Stop(ctx context.Context) error {
	return s.nodeUsecase.Stop(ctx)
}
