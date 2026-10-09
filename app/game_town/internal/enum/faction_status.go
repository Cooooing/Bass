package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type FactionStatus string

const (
	FactionStatusActive    FactionStatus = "active"    // active
	FactionStatusDeclining FactionStatus = "declining" // declining
	FactionStatusCollapsed FactionStatus = "collapsed" // collapsed
)

var FactionStatusMap = commonenum.NewMapping[FactionStatus, v1.GameTownFactionStatus](map[FactionStatus]commonenum.Entry[FactionStatus, v1.GameTownFactionStatus]{FactionStatusActive: {Proto: v1.GameTownFactionStatus_GAME_TOWN_FACTION_STATUS_ACTIVE}, FactionStatusDeclining: {Proto: v1.GameTownFactionStatus_GAME_TOWN_FACTION_STATUS_DECLINING}, FactionStatusCollapsed: {Proto: v1.GameTownFactionStatus_GAME_TOWN_FACTION_STATUS_COLLAPSED}})

func (e FactionStatus) String() string { return string(e) }
