package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

var EntityTypeMap = commonenum.NewMapping[EntityType, v1.GameTownEntityType](
	map[EntityType]commonenum.Entry[EntityType, v1.GameTownEntityType]{
		EntityTypePlayer:   {Proto: v1.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_PLAYER},
		EntityTypeNpc:      {Proto: v1.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_NPC},
		EntityTypeLocation: {Proto: v1.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_LOCATION},
		EntityTypeFaction:  {Proto: v1.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_FACTION},
	},
)

var KnowledgeCertaintyMap = commonenum.NewMapping[KnowledgeCertainty, v1.GameTownKnowledgeCertainty](
	map[KnowledgeCertainty]commonenum.Entry[KnowledgeCertainty, v1.GameTownKnowledgeCertainty]{
		KnowledgeCertaintyConfirmed: {Proto: v1.GameTownKnowledgeCertainty_GAME_TOWN_KNOWLEDGE_CERTAINTY_CONFIRMED},
		KnowledgeCertaintyLikely:    {Proto: v1.GameTownKnowledgeCertainty_GAME_TOWN_KNOWLEDGE_CERTAINTY_LIKELY},
		KnowledgeCertaintyRumor:     {Proto: v1.GameTownKnowledgeCertainty_GAME_TOWN_KNOWLEDGE_CERTAINTY_RUMOR},
		KnowledgeCertaintyUnknown:   {Proto: v1.GameTownKnowledgeCertainty_GAME_TOWN_KNOWLEDGE_CERTAINTY_UNKNOWN},
	},
)

var ObservationSourceMap = commonenum.NewMapping[ObservationSource, v1.GameTownObservationSource](
	map[ObservationSource]commonenum.Entry[ObservationSource, v1.GameTownObservationSource]{
		ObservationSourceExperienced: {Proto: v1.GameTownObservationSource_GAME_TOWN_OBSERVATION_SOURCE_EXPERIENCED},
		ObservationSourceWitnessed:   {Proto: v1.GameTownObservationSource_GAME_TOWN_OBSERVATION_SOURCE_WITNESSED},
		ObservationSourceTold:        {Proto: v1.GameTownObservationSource_GAME_TOWN_OBSERVATION_SOURCE_TOLD},
		ObservationSourceFaction:     {Proto: v1.GameTownObservationSource_GAME_TOWN_OBSERVATION_SOURCE_FACTION},
		ObservationSourcePublic:      {Proto: v1.GameTownObservationSource_GAME_TOWN_OBSERVATION_SOURCE_PUBLIC},
	},
)

var RelationshipAttitudeMap = commonenum.NewMapping[RelationshipAttitude, v1.GameTownRelationshipAttitude](
	map[RelationshipAttitude]commonenum.Entry[RelationshipAttitude, v1.GameTownRelationshipAttitude]{
		RelationshipAttitudeFriendly: {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_FRIENDLY},
		RelationshipAttitudeTrusting: {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_TRUSTING},
		RelationshipAttitudeNeutral:  {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_NEUTRAL},
		RelationshipAttitudeWary:     {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_WARY},
		RelationshipAttitudeAfraid:   {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_AFRAID},
		RelationshipAttitudeHostile:  {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_HOSTILE},
	},
)

var NpcLifeStatusMap = commonenum.NewMapping[NpcLifeStatus, v1.GameTownNpcLifeStatus](
	map[NpcLifeStatus]commonenum.Entry[NpcLifeStatus, v1.GameTownNpcLifeStatus]{
		NpcLifeStatusAlive:   {Proto: v1.GameTownNpcLifeStatus_GAME_TOWN_NPC_LIFE_STATUS_ALIVE},
		NpcLifeStatusMissing: {Proto: v1.GameTownNpcLifeStatus_GAME_TOWN_NPC_LIFE_STATUS_MISSING},
		NpcLifeStatusDead:    {Proto: v1.GameTownNpcLifeStatus_GAME_TOWN_NPC_LIFE_STATUS_DEAD},
	},
)

var LocationStatusMap = commonenum.NewMapping[LocationStatus, v1.GameTownLocationStatus](
	map[LocationStatus]commonenum.Entry[LocationStatus, v1.GameTownLocationStatus]{
		LocationStatusActive:    {Proto: v1.GameTownLocationStatus_GAME_TOWN_LOCATION_STATUS_ACTIVE},
		LocationStatusDamaged:   {Proto: v1.GameTownLocationStatus_GAME_TOWN_LOCATION_STATUS_DAMAGED},
		LocationStatusBlocked:   {Proto: v1.GameTownLocationStatus_GAME_TOWN_LOCATION_STATUS_BLOCKED},
		LocationStatusDestroyed: {Proto: v1.GameTownLocationStatus_GAME_TOWN_LOCATION_STATUS_DESTROYED},
		LocationStatusRuins:     {Proto: v1.GameTownLocationStatus_GAME_TOWN_LOCATION_STATUS_RUINS},
	},
)

var FactionStatusMap = commonenum.NewMapping[FactionStatus, v1.GameTownFactionStatus](
	map[FactionStatus]commonenum.Entry[FactionStatus, v1.GameTownFactionStatus]{
		FactionStatusActive:    {Proto: v1.GameTownFactionStatus_GAME_TOWN_FACTION_STATUS_ACTIVE},
		FactionStatusDeclining: {Proto: v1.GameTownFactionStatus_GAME_TOWN_FACTION_STATUS_DECLINING},
		FactionStatusCollapsed: {Proto: v1.GameTownFactionStatus_GAME_TOWN_FACTION_STATUS_COLLAPSED},
	},
)

var ClaimTruthMap = commonenum.NewMapping[ClaimTruth, v1.GameTownClaimTruth](
	map[ClaimTruth]commonenum.Entry[ClaimTruth, v1.GameTownClaimTruth]{
		ClaimTruthTrue:    {Proto: v1.GameTownClaimTruth_GAME_TOWN_CLAIM_TRUTH_TRUE},
		ClaimTruthFalse:   {Proto: v1.GameTownClaimTruth_GAME_TOWN_CLAIM_TRUTH_FALSE},
		ClaimTruthUnknown: {Proto: v1.GameTownClaimTruth_GAME_TOWN_CLAIM_TRUTH_UNKNOWN},
	},
)

var BeliefStanceMap = commonenum.NewMapping[BeliefStance, v1.GameTownBeliefStance](
	map[BeliefStance]commonenum.Entry[BeliefStance, v1.GameTownBeliefStance]{
		BeliefStanceBelieves: {Proto: v1.GameTownBeliefStance_GAME_TOWN_BELIEF_STANCE_BELIEVES},
		BeliefStanceDoubts:   {Proto: v1.GameTownBeliefStance_GAME_TOWN_BELIEF_STANCE_DOUBTS},
		BeliefStanceRejects:  {Proto: v1.GameTownBeliefStance_GAME_TOWN_BELIEF_STANCE_REJECTS},
	},
)

var MemoryKindMap = commonenum.NewMapping[MemoryKind, v1.GameTownMemoryKind](
	map[MemoryKind]commonenum.Entry[MemoryKind, v1.GameTownMemoryKind]{
		MemoryKindExperience:   {Proto: v1.GameTownMemoryKind_GAME_TOWN_MEMORY_KIND_EXPERIENCE},
		MemoryKindConversation: {Proto: v1.GameTownMemoryKind_GAME_TOWN_MEMORY_KIND_CONVERSATION},
		MemoryKindRelationship: {Proto: v1.GameTownMemoryKind_GAME_TOWN_MEMORY_KIND_RELATIONSHIP},
		MemoryKindReflection:   {Proto: v1.GameTownMemoryKind_GAME_TOWN_MEMORY_KIND_REFLECTION},
		MemoryKindRumor:        {Proto: v1.GameTownMemoryKind_GAME_TOWN_MEMORY_KIND_RUMOR},
	},
)

var EmbeddingStatusMap = commonenum.NewMapping[EmbeddingStatus, v1.GameTownEmbeddingStatus](
	map[EmbeddingStatus]commonenum.Entry[EmbeddingStatus, v1.GameTownEmbeddingStatus]{
		EmbeddingStatusPending: {Proto: v1.GameTownEmbeddingStatus_GAME_TOWN_EMBEDDING_STATUS_PENDING},
		EmbeddingStatusReady:   {Proto: v1.GameTownEmbeddingStatus_GAME_TOWN_EMBEDDING_STATUS_READY},
		EmbeddingStatusFailed:  {Proto: v1.GameTownEmbeddingStatus_GAME_TOWN_EMBEDDING_STATUS_FAILED},
	},
)
