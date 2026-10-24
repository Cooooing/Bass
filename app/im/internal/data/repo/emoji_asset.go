package repo

import (
	"context"

	"common/pkg/client/rpc"
	platformv1 "common/proto/gen/platform/v1"

	bizrepo "im/internal/biz/repo"
)

// EmojiAssetClient 通过 Platform 验证图片资产的全局可用性。
type EmojiAssetClient struct {
	asset platformv1.PlatformAssetServiceClient
}

func NewEmojiAssetClient(client *rpc.PlatformClient) bizrepo.EmojiAssetClient {
	return &EmojiAssetClient{asset: client.Asset}
}

func (c *EmojiAssetClient) ValidateAvailable(ctx context.Context, assetID int64) error {
	_, err := c.asset.ValidateAvailable(ctx, &platformv1.ValidateAvailableAsset_Req{AssetId: assetID})
	return err
}
