package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type PlayerStatus string

const (
	PlayerStatusActive   PlayerStatus = "active"
	PlayerStatusDisabled PlayerStatus = "disabled"
)

var PlayerStatusMap = commonenum.NewMapping[PlayerStatus, v1.GameTownPlayerStatus](
	map[PlayerStatus]commonenum.Entry[PlayerStatus, v1.GameTownPlayerStatus]{
		PlayerStatusActive: {
			Proto: v1.GameTownPlayerStatus_GAME_TOWN_PLAYER_STATUS_ACTIVE,
		},
		PlayerStatusDisabled: {
			Proto: v1.GameTownPlayerStatus_GAME_TOWN_PLAYER_STATUS_DISABLED,
		},
	},
)

func (e PlayerStatus) String() string {
	return string(e)
}
