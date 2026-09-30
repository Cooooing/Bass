package repo

import (
	"common/pkg/client/rpc"
	platformv1 "common/proto/gen/platform/v1"
	"context"
	"im/internal/biz/repo"
)

var _ repo.AssetClient = (*AssetClient)(nil)

type AssetClient struct{ platform *rpc.PlatformClient }

func NewAssetClient(platform *rpc.PlatformClient) repo.AssetClient {
	return &AssetClient{platform: platform}
}
func (r *AssetClient) ValidateAvailable(ctx context.Context, assetID int64) error {
	_, err := r.platform.Asset.ValidateAvailable(ctx, &platformv1.ValidateAvailableAsset_Req{AssetId: assetID})
	return err
}
