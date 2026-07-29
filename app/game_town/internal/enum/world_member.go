package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type WorldMemberRole string

const (
	WorldMemberRoleOwner  WorldMemberRole = "owner"
	WorldMemberRoleMember WorldMemberRole = "member"
)

var WorldMemberRoleMap = commonenum.NewMapping[WorldMemberRole, v1.GameTownWorldMemberRole](
	map[WorldMemberRole]commonenum.Entry[WorldMemberRole, v1.GameTownWorldMemberRole]{
		WorldMemberRoleOwner: {
			Proto: v1.GameTownWorldMemberRole_GAME_TOWN_WORLD_MEMBER_ROLE_OWNER,
		},
		WorldMemberRoleMember: {
			Proto: v1.GameTownWorldMemberRole_GAME_TOWN_WORLD_MEMBER_ROLE_MEMBER,
		},
	},
)

func (e WorldMemberRole) String() string {
	return string(e)
}
