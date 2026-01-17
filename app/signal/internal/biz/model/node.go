package model

import (
	v1 "common/api/signal/v1"
	"signal/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Node struct {
	*gen.Node
}

// ConvertToRpc 转换为RPC返回格式
func (a *Node) ConvertToRpc() *v1.Node {
	group := &v1.Node{
		CreatedAt:   timestamppb.New(*a.CreatedAt),
		UpdatedAt:   timestamppb.New(*a.UpdatedAt),
		Id:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		CallbackUrl: a.CallbackURL,
		Status:      a.Status,
	}
	return group
}
