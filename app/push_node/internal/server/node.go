package server

import (
	"context"
	"fmt"

	pushhubv1 "common/proto/gen/push_hub/v1"
	"push_node/internal/biz/usecase"
	"push_node/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NodeServer struct {
	nodeUsecase *usecase.NodeUsecase
	sseUsecase  *usecase.SSEUsecase
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

func NewPushHubNodeClient(conn *grpc.ClientConn) pushhubv1.PushHubNodeServiceClient {
	return pushhubv1.NewPushHubNodeServiceClient(conn)
}

func NewNodeServer(
	nodeUsecase *usecase.NodeUsecase,
	sseUsecase *usecase.SSEUsecase,
) *NodeServer {
	return &NodeServer{
		nodeUsecase: nodeUsecase,
		sseUsecase:  sseUsecase,
	}
}

func (s *NodeServer) Start(ctx context.Context) error {
	if err := s.nodeUsecase.ConnectHub(ctx); err != nil {
		return err
	}
	if err := s.sseUsecase.StartConsuming(ctx, s.nodeUsecase.NodeID()); err != nil {
		_ = s.nodeUsecase.Stop(ctx)
		return err
	}
	return nil
}

func (s *NodeServer) Stop(ctx context.Context) error {
	var err error
	if stopErr := s.sseUsecase.StopConsuming(ctx); stopErr != nil {
		err = stopErr
	}
	if stopErr := s.nodeUsecase.Stop(ctx); stopErr != nil && err == nil {
		err = stopErr
	}
	return err
}
