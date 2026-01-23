package model

import (
	v1 "common/api/signal/v1"
	"math"
	"signal/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Node struct {
	*gen.Node
}

func (n *Node) CalculateScore(connections int64, pingMs int64, powCostMs int64) float64 {
	// 设置平滑常数
	const (
		alpha = 20.0
		beta  = 50.0
		scale = 1e9
	)

	// 处理分母项
	pPart := float64(pingMs) + alpha
	cPart := float64(powCostMs) + beta

	// log10(N + 10) 保证 N=0 时结果为 1
	nPart := math.Log10(float64(connections) + 10.0)

	// 计算综合得分
	return (n.Weight * scale) / (pPart * cPart * nPart)
}

// ConvertToRpc 转换为RPC返回格式
func (n *Node) ConvertToRpc() *v1.Node {
	group := &v1.Node{
		CreatedAt:   timestamppb.New(*n.CreatedAt),
		UpdatedAt:   timestamppb.New(*n.UpdatedAt),
		Id:          n.ID,
		Key:         n.Key,
		Name:        n.Name,
		Description: n.Description,
		CallbackUrl: n.CallbackURL,
		Status:      n.Status,
	}
	return group
}
