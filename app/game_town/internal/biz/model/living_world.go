package model

import (
	"time"

	"game_town/internal/enum"
)

type Observation struct {
	ID         int64
	WorldID    int64
	EventID    int64
	NpcID      *int64
	PlayerID   *int64
	Source     enum.ObservationSource
	Certainty  enum.KnowledgeCertainty
	Summary    string
	Salience   float64
	ObservedAt time.Time
	WorldTime  time.Time
	CreatedAt  *time.Time
}

type Claim struct {
	ID            int64
	WorldID       int64
	OriginEventID *int64
	SubjectType   enum.EntityType
	SubjectID     int64
	Predicate     string
	Object        map[string]any
	Truth         enum.ClaimTruth
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

type ClaimDraft struct {
	SubjectType enum.EntityType `json:"subject_type"`
	SubjectID   int64           `json:"subject_id"`
	Predicate   string          `json:"predicate"`
	Object      map[string]any  `json:"object"`
	Truth       enum.ClaimTruth `json:"truth"`
}

type NpcBelief struct {
	ID             int64
	WorldID        int64
	NpcID          int64
	ClaimID        int64
	SourceNpcID    *int64
	SourcePlayerID *int64
	SourceEventID  *int64
	Stance         enum.BeliefStance
	Confidence     float64
	LearnedAt      time.Time
	UpdatedAt      *time.Time
}

type Relationship struct {
	ID         int64
	WorldID    int64
	SourceType enum.EntityType
	SourceID   int64
	TargetType enum.EntityType
	TargetID   int64
	Metrics    map[string]float64
	Tags       []string
	Version    int64
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

type Faction struct {
	ID          int64
	WorldID     int64
	Code        string
	Name        string
	Description string
	PublicGoal  string
	Status      enum.FactionStatus
	Attributes  map[string]any
	Version     int64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type FactionMembership struct {
	ID         int64
	WorldID    int64
	FactionID  int64
	MemberType enum.EntityType
	MemberID   int64
	Role       string
	Reputation map[string]float64
	Tags       []string
	JoinedAt   time.Time
	LeftAt     *time.Time
}

type NpcMemory struct {
	ID                  int64
	WorldID             int64
	NpcID               int64
	SourceEventID       *int64
	SourceObservationID *int64
	Kind                enum.MemoryKind
	Content             string
	Importance          float64
	OccurredWorldTime   time.Time
	LastRecalledAt      *time.Time
	EmbeddingModel      string
	EmbeddingStatus     enum.EmbeddingStatus
	EmbeddingError      string
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
}

type NpcView struct {
	Npc              *Npc
	Attitude         enum.RelationshipAttitude
	RelationshipTags []string
	PublicFactionIDs []int64
	KnownAt          time.Time
}

type FactionView struct {
	Faction           *Faction
	Attitude          enum.RelationshipAttitude
	RelationshipTags  []string
	KnownMemberNpcIDs []int64
	KnownLocationIDs  []int64
	PlayerRole        string
	ReputationTags    []string
	KnownAt           time.Time
}
