package server

import (
	"context"
	"platform/internal/biz/repo"
	"platform/internal/biz/usecase"
)

// AssetEventConsumerServer joins the durable NATS consumer to the application
// lifecycle. It has no HTTP or gRPC endpoint and is intentionally invisible
// to external callers.
type AssetEventConsumerServer struct {
	assetEventConsumer repo.AssetUploadEventConsumer
	assetUsecase       *usecase.AssetUsecase
}

func NewAssetEventConsumerServer(
	assetEventConsumer repo.AssetUploadEventConsumer,
	assetUsecase *usecase.AssetUsecase,
) *AssetEventConsumerServer {
	return &AssetEventConsumerServer{
		assetEventConsumer: assetEventConsumer,
		assetUsecase:       assetUsecase,
	}
}

func (s *AssetEventConsumerServer) Start(ctx context.Context) error {
	return s.assetEventConsumer.Consume(ctx, s.assetUsecase.RecordAssetUploadEvent)
}

func (s *AssetEventConsumerServer) Stop(ctx context.Context) error {
	return s.assetEventConsumer.Stop(ctx)
}
