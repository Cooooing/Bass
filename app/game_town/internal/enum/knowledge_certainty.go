package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type KnowledgeCertainty string

const (
	KnowledgeCertaintyConfirmed KnowledgeCertainty = "confirmed" // confirmed
	KnowledgeCertaintyLikely    KnowledgeCertainty = "likely"    // likely
	KnowledgeCertaintyRumor     KnowledgeCertainty = "rumor"     // rumor
	KnowledgeCertaintyUnknown   KnowledgeCertainty = "unknown"   // unknown
)

var KnowledgeCertaintyMap = commonenum.NewMapping[KnowledgeCertainty, v1.GameTownKnowledgeCertainty](map[KnowledgeCertainty]commonenum.Entry[KnowledgeCertainty, v1.GameTownKnowledgeCertainty]{
	KnowledgeCertaintyConfirmed: {Proto: v1.GameTownKnowledgeCertainty_GAME_TOWN_KNOWLEDGE_CERTAINTY_CONFIRMED}, KnowledgeCertaintyLikely: {Proto: v1.GameTownKnowledgeCertainty_GAME_TOWN_KNOWLEDGE_CERTAINTY_LIKELY}, KnowledgeCertaintyRumor: {Proto: v1.GameTownKnowledgeCertainty_GAME_TOWN_KNOWLEDGE_CERTAINTY_RUMOR}, KnowledgeCertaintyUnknown: {Proto: v1.GameTownKnowledgeCertainty_GAME_TOWN_KNOWLEDGE_CERTAINTY_UNKNOWN},
})

func (e KnowledgeCertainty) String() string { return string(e) }
