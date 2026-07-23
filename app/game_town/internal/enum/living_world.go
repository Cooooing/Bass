package enum

type EntityType string

const (
	EntityTypePlayer   EntityType = "player"
	EntityTypeNpc      EntityType = "npc"
	EntityTypeLocation EntityType = "location"
	EntityTypeFaction  EntityType = "faction"
)

type KnowledgeCertainty string

const (
	KnowledgeCertaintyConfirmed KnowledgeCertainty = "confirmed"
	KnowledgeCertaintyLikely    KnowledgeCertainty = "likely"
	KnowledgeCertaintyRumor     KnowledgeCertainty = "rumor"
	KnowledgeCertaintyUnknown   KnowledgeCertainty = "unknown"
)

type ObservationSource string

const (
	ObservationSourceExperienced ObservationSource = "experienced"
	ObservationSourceWitnessed   ObservationSource = "witnessed"
	ObservationSourceTold        ObservationSource = "told"
	ObservationSourceFaction     ObservationSource = "faction"
	ObservationSourcePublic      ObservationSource = "public"
)

type RelationshipAttitude string

const (
	RelationshipAttitudeFriendly RelationshipAttitude = "friendly"
	RelationshipAttitudeTrusting RelationshipAttitude = "trusting"
	RelationshipAttitudeNeutral  RelationshipAttitude = "neutral"
	RelationshipAttitudeWary     RelationshipAttitude = "wary"
	RelationshipAttitudeAfraid   RelationshipAttitude = "afraid"
	RelationshipAttitudeHostile  RelationshipAttitude = "hostile"
)

type NpcLifeStatus string

const (
	NpcLifeStatusAlive   NpcLifeStatus = "alive"
	NpcLifeStatusMissing NpcLifeStatus = "missing"
	NpcLifeStatusDead    NpcLifeStatus = "dead"
)

type LocationStatus string

const (
	LocationStatusActive    LocationStatus = "active"
	LocationStatusDamaged   LocationStatus = "damaged"
	LocationStatusBlocked   LocationStatus = "blocked"
	LocationStatusDestroyed LocationStatus = "destroyed"
	LocationStatusRuins     LocationStatus = "ruins"
)

type FactionStatus string

const (
	FactionStatusActive    FactionStatus = "active"
	FactionStatusDeclining FactionStatus = "declining"
	FactionStatusCollapsed FactionStatus = "collapsed"
)

type ClaimTruth string

const (
	ClaimTruthTrue    ClaimTruth = "true"
	ClaimTruthFalse   ClaimTruth = "false"
	ClaimTruthUnknown ClaimTruth = "unknown"
)

type BeliefStance string

const (
	BeliefStanceBelieves BeliefStance = "believes"
	BeliefStanceDoubts   BeliefStance = "doubts"
	BeliefStanceRejects  BeliefStance = "rejects"
)

type MemoryKind string

const (
	MemoryKindExperience   MemoryKind = "experience"
	MemoryKindConversation MemoryKind = "conversation"
	MemoryKindRelationship MemoryKind = "relationship"
	MemoryKindReflection   MemoryKind = "reflection"
	MemoryKindRumor        MemoryKind = "rumor"
)

type EmbeddingStatus string

const (
	EmbeddingStatusPending EmbeddingStatus = "pending"
	EmbeddingStatusReady   EmbeddingStatus = "ready"
	EmbeddingStatusFailed  EmbeddingStatus = "failed"
)
