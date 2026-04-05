package model

import (
	v1 "common/gen/signal/v1"
	"common/pkg/util"
	"math"
	"signal/internal/data/ent/gen"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Node struct {
	*gen.Node

	Connections  int64      `json:"-"` // 当前节点的连接数
	PingMs       int64      `json:"-"` // 节点的 ping 耗时
	PowCostMs    int64      `json:"-"` // 节点的 pow 耗时（包含网络延迟）
	LastPingTime *time.Time `json:"-"` // 节点的最后 ping 时间
}

func (n *Node) CalculateScore() float64 {
	// 设置平滑常数
	const (
		alpha = 20.0
		beta  = 50.0
		scale = 1e9
	)

	// 处理分母项
	pPart := float64(n.PingMs) + alpha
	n.PowCostMs = n.PowCostMs - n.PingMs
	cPart := util.If(n.PowCostMs < 0, 0, float64(n.PowCostMs)) + beta

	// log10(N + 10) 保证 N=0 时结果为 1
	nPart := math.Log10(float64(n.Connections) + 10.0)

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
