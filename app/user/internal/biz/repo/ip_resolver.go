package repo

import (
	commonModel "common/pkg/model"
	"context"
)

type IPClient interface {
	Resolve(ctx context.Context, ip string) (*commonModel.IpInfo, error)
}
