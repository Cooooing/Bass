package repo

import "context"

type AssetClient interface {
	Get(ctx context.Context, assetID int64) (*Asset, error)
	Map(ctx context.Context, req *AssetGetReq) (map[int64]*Asset, error)
}

type AssetGetReq struct {
	IDs []int64
}

type Asset struct {
	ID  int64
	URL string
}
