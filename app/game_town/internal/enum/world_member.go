package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

// WorldMemberRole 是玩家在世界内成员角色的内部持久化枚举。
type WorldMemberRole string

const (
	WorldMemberRoleOwner  WorldMemberRole = "owner"
	WorldMemberRoleMember WorldMemberRole = "member"
)

// WorldMemberRoleMap 维护内部持久化值与 proto 枚举之间的映射。
var WorldMemberRoleMap = commonenum.NewMapping[WorldMemberRole, v1.GameTownWorldMemberRole](map[WorldMemberRole]commonenum.Entry[WorldMemberRole, v1.GameTownWorldMemberRole]{
	WorldMemberRoleOwner:  {Proto: v1.GameTownWorldMemberRole_GAME_TOWN_WORLD_MEMBER_ROLE_OWNER},
	WorldMemberRoleMember: {Proto: v1.GameTownWorldMemberRole_GAME_TOWN_WORLD_MEMBER_ROLE_MEMBER},
})
