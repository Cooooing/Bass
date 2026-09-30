package repo

import "context"

// AssetClient keeps IM independent of platform transport details.
type AssetClient interface {
	ValidateAvailable(ctx context.Context, assetID int64) error
}
