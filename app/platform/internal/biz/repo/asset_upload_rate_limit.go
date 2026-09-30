package repo

import (
	"context"
	"time"
)

// AssetUploadRateLimitState is the outcome of consuming a direct-upload
// allowance. RetryAfter is populated only when the request was rejected.
type AssetUploadRateLimitState struct {
	Allowed    bool
	RetryAfter time.Duration
}

// AssetUploadRateLimitCache atomically consumes one direct-upload allowance
// for an authenticated uploader. It stores only short-lived operational state;
// upload and asset audit facts remain in PostgreSQL and MinIO metadata.
type AssetUploadRateLimitCache interface {
	Allow(ctx context.Context, uploadByID int64, window time.Duration, maxCount int64) (*AssetUploadRateLimitState, error)
}
