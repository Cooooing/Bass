package repo

import (
	"context"
)

type AssetClient interface {
	Get(ctx context.Context, assetID int64) (*Asset, error)
	Map(ctx context.Context, req *AssetGetReq) (map[int64]*Asset, error)
	PrepareDirectUpload(ctx context.Context, req *PrepareDirectAssetUploadReq) (*PrepareDirectAssetUploadResp, error)
	GetAvailableByHash(ctx context.Context, hash string) (*Asset, error)
	ValidateAvailable(ctx context.Context, assetID int64) error
}

type AssetGetReq struct {
	IDs []int64
}

type Asset struct {
	ID       int64
	URL      string
	MimeType string
	Size     int64
}

type PrepareDirectAssetUploadReq struct {
	Hash       string
	MimeType   string
	Size       int64
	UploadByID int64
}

type PrepareDirectAssetUploadResp struct {
	AssetID    *int64
	UploadURL  string
	FormFields map[string]string
}
