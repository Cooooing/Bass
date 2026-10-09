package server

import (
	"context"

	"platform/internal/biz/usecase"
)

// AssetEventConsumerServer joins the durable NATS consumer to the application
// lifecycle. It has no HTTP or gRPC endpoint and is intentionally invisible
// to external callers.
type AssetEventConsumerServer struct {
	assetUsecase *usecase.AssetUsecase
}

func NewAssetEventConsumerServer(
	assetUsecase *usecase.AssetUsecase,
) *AssetEventConsumerServer {
	return &AssetEventConsumerServer{
		assetUsecase: assetUsecase,
	}
}

func (s *AssetEventConsumerServer) Start(ctx context.Context) error {
	return s.assetUsecase.StartConsumer(ctx)
}

func (s *AssetEventConsumerServer) Stop(ctx context.Context) error {
	return s.assetUsecase.StopConsumer(ctx)
}
