package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type ObservationSource string

const (
	ObservationSourceExperienced ObservationSource = "experienced" // experienced
	ObservationSourceWitnessed   ObservationSource = "witnessed"   // witnessed
	ObservationSourceTold        ObservationSource = "told"        // told
	ObservationSourceFaction     ObservationSource = "faction"     // faction
	ObservationSourcePublic      ObservationSource = "public"      // public
)

var ObservationSourceMap = commonenum.NewMapping[ObservationSource, v1.GameTownObservationSource](map[ObservationSource]commonenum.Entry[ObservationSource, v1.GameTownObservationSource]{
	ObservationSourceExperienced: {Proto: v1.GameTownObservationSource_GAME_TOWN_OBSERVATION_SOURCE_EXPERIENCED}, ObservationSourceWitnessed: {Proto: v1.GameTownObservationSource_GAME_TOWN_OBSERVATION_SOURCE_WITNESSED}, ObservationSourceTold: {Proto: v1.GameTownObservationSource_GAME_TOWN_OBSERVATION_SOURCE_TOLD}, ObservationSourceFaction: {Proto: v1.GameTownObservationSource_GAME_TOWN_OBSERVATION_SOURCE_FACTION}, ObservationSourcePublic: {Proto: v1.GameTownObservationSource_GAME_TOWN_OBSERVATION_SOURCE_PUBLIC},
})

func (e ObservationSource) String() string { return string(e) }
