package enum

import (
	commonenum "common/pkg/enum"
	realtimehubv1enum "common/proto/gen/realtime_hub/v1/enum"
)

// NodeGroup 表示 realtime node 的部署与接入边界。
type NodeGroup string

const (
	NodeGroupInternal NodeGroup = "internal" // 主站内部节点集群
)

var NodeGroupMap = commonenum.NewMapping[NodeGroup, realtimehubv1enum.NodeGroup](map[NodeGroup]commonenum.Entry[NodeGroup, realtimehubv1enum.NodeGroup]{
	NodeGroupInternal: {Proto: realtimehubv1enum.NodeGroup_NODE_GROUP_INTERNAL},
})
