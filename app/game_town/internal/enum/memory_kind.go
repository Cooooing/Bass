package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type MemoryKind string

const (
	MemoryKindExperience   MemoryKind = "experience"   // experience
	MemoryKindConversation MemoryKind = "conversation" // conversation
	MemoryKindRelationship MemoryKind = "relationship" // relationship
	MemoryKindReflection   MemoryKind = "reflection"   // reflection
	MemoryKindRumor        MemoryKind = "rumor"        // rumor
)

var MemoryKindMap = commonenum.NewMapping[MemoryKind, v1.GameTownMemoryKind](map[MemoryKind]commonenum.Entry[MemoryKind, v1.GameTownMemoryKind]{MemoryKindExperience: {Proto: v1.GameTownMemoryKind_GAME_TOWN_MEMORY_KIND_EXPERIENCE}, MemoryKindConversation: {Proto: v1.GameTownMemoryKind_GAME_TOWN_MEMORY_KIND_CONVERSATION}, MemoryKindRelationship: {Proto: v1.GameTownMemoryKind_GAME_TOWN_MEMORY_KIND_RELATIONSHIP}, MemoryKindReflection: {Proto: v1.GameTownMemoryKind_GAME_TOWN_MEMORY_KIND_REFLECTION}, MemoryKindRumor: {Proto: v1.GameTownMemoryKind_GAME_TOWN_MEMORY_KIND_RUMOR}})

func (e MemoryKind) String() string { return string(e) }
