package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	platformv1 "common/proto/gen/platform/v1"
	"context"
)

var _ repo.AssetClient = (*AssetClient)(nil)

type AssetClient struct {
	platformClient *rpc.PlatformClient
}

func NewAssetClient(platformClient *rpc.PlatformClient) repo.AssetClient {
	return &AssetClient{platformClient: platformClient}
}

func (r *AssetClient) Get(ctx context.Context, assetID int64) (*repo.Asset, error) {
	if assetID <= 0 || r.platformClient == nil {
		return nil, nil
	}
	reply, err := r.platformClient.Oss.ResolveAssets(ctx, &platformv1.ResolveAssetsOss_Req{AssetIds: []int64{assetID}})
	if err != nil {
		return nil, err
	}
	asset := reply.GetAssets()[assetID]
	if asset == nil {
		return nil, nil
	}
	return &repo.Asset{ID: assetID, URL: asset.GetUrl()}, nil
}

func (r *AssetClient) Map(ctx context.Context, req *repo.AssetGetReq) (map[int64]*repo.Asset, error) {
	if req == nil || len(req.IDs) == 0 || r.platformClient == nil {
		return map[int64]*repo.Asset{}, nil
	}
	reply, err := r.platformClient.Oss.ResolveAssets(ctx, &platformv1.ResolveAssetsOss_Req{AssetIds: req.IDs})
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*repo.Asset, len(reply.GetAssets()))
	for assetID, asset := range reply.GetAssets() {
		if asset == nil {
			continue
		}
		out[assetID] = &repo.Asset{ID: assetID, URL: asset.GetUrl()}
	}
	return out, nil
}
