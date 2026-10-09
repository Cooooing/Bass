package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type BeliefStance string

const (
	BeliefStanceBelieves BeliefStance = "believes" // believes
	BeliefStanceDoubts   BeliefStance = "doubts"   // doubts
	BeliefStanceRejects  BeliefStance = "rejects"  // rejects
)

var BeliefStanceMap = commonenum.NewMapping[BeliefStance, v1.GameTownBeliefStance](map[BeliefStance]commonenum.Entry[BeliefStance, v1.GameTownBeliefStance]{BeliefStanceBelieves: {Proto: v1.GameTownBeliefStance_GAME_TOWN_BELIEF_STANCE_BELIEVES}, BeliefStanceDoubts: {Proto: v1.GameTownBeliefStance_GAME_TOWN_BELIEF_STANCE_DOUBTS}, BeliefStanceRejects: {Proto: v1.GameTownBeliefStance_GAME_TOWN_BELIEF_STANCE_REJECTS}})

func (e BeliefStance) String() string { return string(e) }
