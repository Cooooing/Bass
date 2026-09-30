package repo

import "context"

type AssetClient interface {
	Get(ctx context.Context, assetID int64) (*Asset, error)
	Map(ctx context.Context, req *AssetGetReq) (map[int64]*Asset, error)
	Upload(ctx context.Context, req *AssetUploadReq) (*Asset, error)
	Validate(ctx context.Context, assetID, userID int64) error
}

type AssetGetReq struct {
	IDs []int64
}

type Asset struct {
	ID  int64
	URL string
}

type AssetUploadReq struct {
	UserID   int64
	FileName string
	MimeType string
	Content  []byte
}
