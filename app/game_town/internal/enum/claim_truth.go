package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type ClaimTruth string

const (
	ClaimTruthTrue    ClaimTruth = "true"    // true
	ClaimTruthFalse   ClaimTruth = "false"   // false
	ClaimTruthUnknown ClaimTruth = "unknown" // unknown
)

var ClaimTruthMap = commonenum.NewMapping[ClaimTruth, v1.GameTownClaimTruth](map[ClaimTruth]commonenum.Entry[ClaimTruth, v1.GameTownClaimTruth]{ClaimTruthTrue: {Proto: v1.GameTownClaimTruth_GAME_TOWN_CLAIM_TRUTH_TRUE}, ClaimTruthFalse: {Proto: v1.GameTownClaimTruth_GAME_TOWN_CLAIM_TRUTH_FALSE}, ClaimTruthUnknown: {Proto: v1.GameTownClaimTruth_GAME_TOWN_CLAIM_TRUTH_UNKNOWN}})

func (e ClaimTruth) String() string { return string(e) }
