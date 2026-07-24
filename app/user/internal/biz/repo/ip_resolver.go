package repo

import (
	commonModel "common/pkg/model"
	"context"
)

type IPResolver interface {
	Resolve(ctx context.Context, ip string) (*commonModel.IpInfo, error)
}
