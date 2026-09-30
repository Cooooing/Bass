package repo

import (
	commonmodel "common/pkg/model"
	"context"
)

// IPResolutionClient resolves an IP address through the platform IP data service.
type IPResolutionClient interface {
	Resolve(ctx context.Context, ip string) (*commonmodel.IpInfo, error)
}
