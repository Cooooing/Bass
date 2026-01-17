package biz

import (
	"common/pkg/cutil/base/str"
	"context"
	"math"

	"github.com/sony/sonyflake/v2"
)

type NodeDomain struct {
	*BaseDomain
	sf *sonyflake.Sonyflake
}

func NewNodeDomain(baseDomain *BaseDomain) (*NodeDomain, error) {
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

// CalculateNodeScore 计算节点分数
func (d *NodeDomain) CalculateNodeScore(weight float64, ping int64, powCost int64, connections int64) int64 {
	// 设置平滑常数
	const alpha = 20.0
	const beta = 50.0
	const scale = 1e9 // 10^9

	// 处理分母项
	pPart := float64(ping) + alpha
	cPart := float64(powCost) + beta

	// log10(N + 10) 保证 N=0 时结果为 1
	nPart := math.Log10(float64(connections) + 10.0)

	// 计算综合得分
	score := (weight * scale) / (pPart * cPart * nPart)

	return int64(score)
}

func (d *NodeDomain) Register(ctx context.Context) error {
	return nil
}

func (d *NodeDomain) Unregister(ctx context.Context) error {
	return nil
}
