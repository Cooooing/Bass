package biz

import (
	"common/pkg/cutil/base/str"
	"context"
	"signal/internal/biz/base"

	"github.com/sony/sonyflake/v2"
)

type NodeDomain struct {
	*base.BaseDomain
	sf *sonyflake.Sonyflake
}

func NewNodeDomain(baseDomain *base.BaseDomain) (*NodeDomain, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &NodeDomain{
		BaseDomain: baseDomain,
		sf:         sf,
	}, nil
}

// GenerateSecret 生成一个 32 位随机字符串
func (d *NodeDomain) GenerateSecret() string {
	return str.RandStr(d.sf, 32, true, true, true, false)
}

func (d *NodeDomain) Register(ctx context.Context) error {
	return nil
}

func (d *NodeDomain) Unregister(ctx context.Context) error {
	return nil
}
