package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
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

var RelationshipAttitudeMap = commonenum.NewMapping[RelationshipAttitude, v1.GameTownRelationshipAttitude](map[RelationshipAttitude]commonenum.Entry[RelationshipAttitude, v1.GameTownRelationshipAttitude]{
	RelationshipAttitudeFriendly: {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_FRIENDLY}, RelationshipAttitudeTrusting: {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_TRUSTING}, RelationshipAttitudeNeutral: {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_NEUTRAL}, RelationshipAttitudeWary: {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_WARY}, RelationshipAttitudeAfraid: {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_AFRAID}, RelationshipAttitudeHostile: {Proto: v1.GameTownRelationshipAttitude_GAME_TOWN_RELATIONSHIP_ATTITUDE_HOSTILE},
})

func (e RelationshipAttitude) String() string { return string(e) }
