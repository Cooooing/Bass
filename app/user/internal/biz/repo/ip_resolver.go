package repo

import (
	"context"
	"user/internal/biz/model"
)

type IPClient interface {
	Resolve(ctx context.Context, ip string) (*model.IpInfo, error)
}
