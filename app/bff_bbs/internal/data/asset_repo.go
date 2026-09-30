package data

import (
	"bff_bbs/internal/biz/repo"
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
	reply, err := r.platformClient.Asset.ResolvePublic(ctx, &platformv1.ResolvePublicAssets_Req{AssetIds: []int64{assetID}})
	if err != nil {
		return nil, err
	}
	url := reply.GetUrls()[assetID]
	if url == "" {
		return nil, nil
	}
	return &repo.Asset{ID: assetID, URL: url}, nil
}

func (r *AssetClient) Map(ctx context.Context, req *repo.AssetGetReq) (map[int64]*repo.Asset, error) {
	if req == nil || len(req.IDs) == 0 || r.platformClient == nil {
		return map[int64]*repo.Asset{}, nil
	}
	reply, err := r.platformClient.Asset.ResolvePublic(ctx, &platformv1.ResolvePublicAssets_Req{AssetIds: req.IDs})
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*repo.Asset, len(reply.GetUrls()))
	for assetID, url := range reply.GetUrls() {
		if url == "" {
			continue
		}
		out[assetID] = &repo.Asset{ID: assetID, URL: url}
	}
	return out, nil
}

func (r *AssetClient) PrepareDirectUpload(ctx context.Context, req *repo.PrepareDirectAssetUploadReq) (*repo.PrepareDirectAssetUploadResp, error) {
	reply, err := r.platformClient.Asset.PrepareDirectUpload(ctx, &platformv1.PrepareDirectAssetUpload_Req{
		Hash:           req.Hash,
		MimeType:       req.MimeType,
		Size:           req.Size,
		UploadByUserId: req.UploadByID,
	})
	if err != nil {
		return nil, err
	}
	return &repo.PrepareDirectAssetUploadResp{AssetID: reply.AssetId, UploadURL: reply.GetUploadUrl(), FormFields: reply.GetFormFields()}, nil
}

func (r *AssetClient) GetAvailableByHash(ctx context.Context, hash string) (*repo.Asset, error) {
	reply, err := r.platformClient.Asset.GetAvailableByHash(ctx, &platformv1.GetAvailableAssetByHash_Req{Hash: hash})
	if err != nil {
		return nil, err
	}
	if reply.GetAsset() == nil {
		return nil, nil
	}
	return &repo.Asset{
		ID:       reply.GetAsset().GetId(),
		URL:      reply.GetAsset().GetUrl(),
		MimeType: reply.GetAsset().GetMimeType(),
		Size:     reply.GetAsset().GetSize(),
	}, nil
}

func (r *AssetClient) ValidateAvailable(ctx context.Context, assetID int64) error {
	_, err := r.platformClient.Asset.ValidateAvailable(ctx, &platformv1.ValidateAvailableAsset_Req{AssetId: assetID})
	return err
}
