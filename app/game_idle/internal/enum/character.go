package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

// CharacterStatus 表示挂机游戏角色状态。
type CharacterStatus string

const (
	CharacterStatusActive   CharacterStatus = "active"   // 正常
	CharacterStatusDisabled CharacterStatus = "disabled" // 禁用
)

var CharacterStatusMap = commonenum.NewMapping[CharacterStatus, v1.CharacterStatus](
	map[CharacterStatus]commonenum.Entry[CharacterStatus, v1.CharacterStatus]{
		CharacterStatusActive: {
			Proto: v1.CharacterStatus_CHARACTER_STATUS_ACTIVE,
		},
		CharacterStatusDisabled: {
			Proto: v1.CharacterStatus_CHARACTER_STATUS_DISABLED,
		},
	},
)

func (e CharacterStatus) String() string {
	return string(e)
}

func CharacterStatusValues() []string {
	return CharacterStatusMap.EnumValues()
}
