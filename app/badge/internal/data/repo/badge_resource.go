package repo

import (
	"badge/internal/biz/repo"
	"common/pkg/client/rpc"
	platformv1 "common/proto/gen/platform/v1"
	"context"
	"fmt"
	"io"
	"net/http"
)

// BadgeResourceRepo reads image and font bytes for Platform-confirmed assets.
type BadgeResourceRepo struct {
	client         *http.Client
	platformClient *rpc.PlatformClient
	maxBytes       int64
}

func NewBadgeResourceRepo(platformClient *rpc.PlatformClient) repo.BadgeResourceReader {
	return &BadgeResourceRepo{
		client:         &http.Client{},
		platformClient: platformClient,
		maxBytes:       3 * 1024 * 1024,
	}
}

func (r *BadgeResourceRepo) Read(ctx context.Context, assetID int64) ([]byte, error) {
	if assetID <= 0 || r.platformClient == nil {
		return nil, fmt.Errorf("badge resource asset is required")
	}
	resolved, err := r.platformClient.Asset.ResolvePublic(ctx, &platformv1.ResolvePublicAssets_Req{AssetIds: []int64{assetID}})
	if err != nil {
		return nil, err
	}
	resourceURL := resolved.GetUrls()[assetID]
	if resourceURL == "" {
		return nil, fmt.Errorf("badge resource asset is unavailable")
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, resourceURL, nil)
	if err != nil {
		return nil, err
	}
	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("badge resource request failed: %s", response.Status)
	}
	content, err := io.ReadAll(io.LimitReader(response.Body, r.maxBytes+1))
	if err != nil {
		return nil, err
	}
	if int64(len(content)) > r.maxBytes {
		return nil, fmt.Errorf("badge resource exceeds %d bytes", r.maxBytes)
	}
	return content, nil
}
