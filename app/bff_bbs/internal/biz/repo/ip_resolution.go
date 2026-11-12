package repo

import (
	"bff_bbs/internal/biz/model"
	"context"
)

// IPResolutionClient resolves an IP address through the platform IP data service.
type IPResolutionClient interface {
	Resolve(ctx context.Context, ip string) (*model.IpInfo, error)
}
