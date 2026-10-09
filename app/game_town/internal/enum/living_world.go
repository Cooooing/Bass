package enum

type EntityType string

const (
	EntityTypePlayer   EntityType = "player"   // player
	EntityTypeNpc      EntityType = "npc"      // npc
	EntityTypeLocation EntityType = "location" // location
	EntityTypeFaction  EntityType = "faction"  // faction
)

type KnowledgeCertainty string

const (
	KnowledgeCertaintyConfirmed KnowledgeCertainty = "confirmed" // confirmed
	KnowledgeCertaintyLikely    KnowledgeCertainty = "likely"    // likely
	KnowledgeCertaintyRumor     KnowledgeCertainty = "rumor"     // rumor
	KnowledgeCertaintyUnknown   KnowledgeCertainty = "unknown"   // unknown
)

type ObservationSource string

const (
	ObservationSourceExperienced ObservationSource = "experienced" // experienced
	ObservationSourceWitnessed   ObservationSource = "witnessed"   // witnessed
	ObservationSourceTold        ObservationSource = "told"        // told
	ObservationSourceFaction     ObservationSource = "faction"     // faction
	ObservationSourcePublic      ObservationSource = "public"      // public
)

type RelationshipAttitude string

const (
	RelationshipAttitudeFriendly RelationshipAttitude = "friendly" // friendly
	RelationshipAttitudeTrusting RelationshipAttitude = "trusting" // trusting
	RelationshipAttitudeNeutral  RelationshipAttitude = "neutral"  // neutral
	RelationshipAttitudeWary     RelationshipAttitude = "wary"     // wary
	RelationshipAttitudeAfraid   RelationshipAttitude = "afraid"   // afraid
	RelationshipAttitudeHostile  RelationshipAttitude = "hostile"  // hostile
)

type NpcLifeStatus string

const (
	NpcLifeStatusAlive   NpcLifeStatus = "alive"   // alive
	NpcLifeStatusMissing NpcLifeStatus = "missing" // missing
	NpcLifeStatusDead    NpcLifeStatus = "dead"    // dead
)

type LocationStatus string

const (
	LocationStatusActive    LocationStatus = "active"    // active
	LocationStatusDamaged   LocationStatus = "damaged"   // damaged
	LocationStatusBlocked   LocationStatus = "blocked"   // blocked
	LocationStatusDestroyed LocationStatus = "destroyed" // destroyed
	LocationStatusRuins     LocationStatus = "ruins"     // ruins
)

type FactionStatus string

const (
	FactionStatusActive    FactionStatus = "active"    // active
	FactionStatusDeclining FactionStatus = "declining" // declining
	FactionStatusCollapsed FactionStatus = "collapsed" // collapsed
)

type ClaimTruth string

const (
	ClaimTruthTrue    ClaimTruth = "true"    // true
	ClaimTruthFalse   ClaimTruth = "false"   // false
	ClaimTruthUnknown ClaimTruth = "unknown" // unknown
)

type BeliefStance string

const (
	BeliefStanceBelieves BeliefStance = "believes" // believes
	BeliefStanceDoubts   BeliefStance = "doubts"   // doubts
	BeliefStanceRejects  BeliefStance = "rejects"  // rejects
)

type MemoryKind string

const (
	MemoryKindExperience   MemoryKind = "experience"   // experience
	MemoryKindConversation MemoryKind = "conversation" // conversation
	MemoryKindRelationship MemoryKind = "relationship" // relationship
	MemoryKindReflection   MemoryKind = "reflection"   // reflection
	MemoryKindRumor        MemoryKind = "rumor"        // rumor
)

type EmbeddingStatus string

const (
	EmbeddingStatusPending EmbeddingStatus = "pending" // pending
	EmbeddingStatusReady   EmbeddingStatus = "ready"   // ready
	EmbeddingStatusFailed  EmbeddingStatus = "failed"  // failed
)

func (e EntityType) String() string {
	return string(e)
}

func (e KnowledgeCertainty) String() string {
	return string(e)
}

func (e ObservationSource) String() string {
	return string(e)
}

func (e RelationshipAttitude) String() string {
	return string(e)
}

func (e NpcLifeStatus) String() string {
	return string(e)
}

func (e LocationStatus) String() string {
	return string(e)
}

func (e FactionStatus) String() string {
	return string(e)
}

func (e ClaimTruth) String() string {
	return string(e)
}

func (e BeliefStance) String() string {
	return string(e)
}

func (e MemoryKind) String() string {
	return string(e)
}

func (e EmbeddingStatus) String() string {
	return string(e)
}
