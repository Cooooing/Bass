package service

import (
	"context"
	"platform/internal/biz/model"
	"platform/internal/biz/usecase"
	"platform/internal/enum"

	v1 "common/proto/gen/platform/v1"
	platformenum "common/proto/gen/platform/v1/enum"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	transporthttp "github.com/go-kratos/kratos/v3/transport/http"
)

// AssetService exposes completed assets only. MinIO event delivery reaches the
// usecase through a durable NATS consumer, not through this gRPC API.
type AssetService struct {
	v1.UnimplementedPlatformAssetServiceServer
	assetUsecase *usecase.AssetUsecase
}

func NewAssetService(assetUsecase *usecase.AssetUsecase) *AssetService {
	return &AssetService{assetUsecase: assetUsecase}
}

func (s *AssetService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterPlatformAssetServiceServer(server, s)
}

func (s *AssetService) RegisterHttp(*transporthttp.Server) {}

func (s *AssetService) PrepareDirectUpload(ctx context.Context, req *v1.PrepareDirectAssetUpload_Req) (*v1.PrepareDirectAssetUpload_Resp, error) {
	result, err := s.assetUsecase.PrepareDirectUpload(ctx, &usecase.PrepareDirectAssetUploadReq{
		Hash:       req.GetHash(),
		MimeType:   req.GetMimeType(),
		Size:       req.GetSize(),
		UploadByID: req.GetUploadByUserId(),
	})
	if err != nil {
		return nil, err
	}
	response := &v1.PrepareDirectAssetUpload_Resp{}
	if result.Asset != nil {
		response.AssetId = &result.Asset.ID
		return response, nil
	}
	response.UploadUrl = result.URL
	response.FormFields = result.Form
	return response, nil
}

func (s *AssetService) GetAvailableByHash(ctx context.Context, req *v1.GetAvailableAssetByHash_Req) (*v1.GetAvailableAssetByHash_Resp, error) {
	asset, err := s.assetUsecase.GetAvailableByHash(ctx, req.GetHash())
	if err != nil {
		return nil, err
	}
	return &v1.GetAvailableAssetByHash_Resp{Asset: s.asset(asset)}, nil
}

func (s *AssetService) ResolvePublic(ctx context.Context, req *v1.ResolvePublicAssets_Req) (*v1.ResolvePublicAssets_Resp, error) {
	urls, err := s.assetUsecase.ResolvePublic(ctx, req.GetAssetIds())
	if err != nil {
		return nil, err
	}
	return &v1.ResolvePublicAssets_Resp{Urls: urls}, nil
}

func (s *AssetService) ValidateAvailable(ctx context.Context, req *v1.ValidateAvailableAsset_Req) (*v1.ValidateAvailableAsset_Resp, error) {
	if err := s.assetUsecase.ValidateAvailable(ctx, req.GetAssetId()); err != nil {
		return nil, err
	}
	return &v1.ValidateAvailableAsset_Resp{}, nil
}

func (s *AssetService) Block(ctx context.Context, req *v1.BlockAsset_Req) (*v1.BlockAsset_Resp, error) {
	if err := s.assetUsecase.Block(ctx, req.GetAssetId(), req.GetReason(), req.GetOperatorUserId()); err != nil {
		return nil, err
	}
	return &v1.BlockAsset_Resp{}, nil
}

func (s *AssetService) asset(asset *model.Asset) *v1.Asset {
	if asset == nil {
		return nil
	}
	hash := ""
	if asset.Hash != nil {
		hash = *asset.Hash
	}
	return &v1.Asset{
		Id:       asset.ID,
		Hash:     hash,
		Url:      s.assetUsecase.PublicURL(asset),
		MimeType: asset.MimeType,
		Size:     asset.Size,
		Status:   assetStatus(asset.Status),
	}
}

func assetStatus(status enum.AssetStatus) platformenum.AssetStatus {
	switch status {
	case enum.AssetStatusBlocked:
		return platformenum.AssetStatus_ASSET_STATUS_BLOCKED
	default:
		return platformenum.AssetStatus_ASSET_STATUS_AVAILABLE
	}
}
