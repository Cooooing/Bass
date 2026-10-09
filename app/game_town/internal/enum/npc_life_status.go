package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type NpcLifeStatus string

const (
	NpcLifeStatusAlive   NpcLifeStatus = "alive"   // alive
	NpcLifeStatusMissing NpcLifeStatus = "missing" // missing
	NpcLifeStatusDead    NpcLifeStatus = "dead"    // dead
)

var NpcLifeStatusMap = commonenum.NewMapping[NpcLifeStatus, v1.GameTownNpcLifeStatus](map[NpcLifeStatus]commonenum.Entry[NpcLifeStatus, v1.GameTownNpcLifeStatus]{NpcLifeStatusAlive: {Proto: v1.GameTownNpcLifeStatus_GAME_TOWN_NPC_LIFE_STATUS_ALIVE}, NpcLifeStatusMissing: {Proto: v1.GameTownNpcLifeStatus_GAME_TOWN_NPC_LIFE_STATUS_MISSING}, NpcLifeStatusDead: {Proto: v1.GameTownNpcLifeStatus_GAME_TOWN_NPC_LIFE_STATUS_DEAD}})

func (e NpcLifeStatus) String() string { return string(e) }
