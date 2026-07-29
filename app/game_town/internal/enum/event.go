package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type EventType string

const (
	EventTypeWorldGenerationRequested    EventType = "world_generation_requested"
	EventTypeWorldReady                  EventType = "world_ready"
	EventTypeWorldGenerationFailed       EventType = "world_generation_failed"
	EventTypePlayerJoined                EventType = "player_joined"
	EventTypePlayerMoved                 EventType = "player_moved"
	EventTypePlayerTalked                EventType = "player_talked"
	EventTypeNpcThinking                 EventType = "npc_thinking"
	EventTypeNpcReplied                  EventType = "npc_replied"
	EventTypePlayerActed                 EventType = "player_acted"
	EventTypeActionResolved              EventType = "action_resolved"
	EventTypeWorldTickRequested          EventType = "world_tick_requested"
	EventTypeWorldEvolved                EventType = "world_evolved"
	EventTypeAgentJobFailed              EventType = "agent_job_failed"
	EventTypePlayerActionSubmitted       EventType = "player_action_submitted"
	EventTypeActionRejected              EventType = "action_rejected"
	EventTypeActionClarificationRequired EventType = "action_clarification_required"
	EventTypeNpcPlanned                  EventType = "npc_planned"
	EventTypeNpcMoved                    EventType = "npc_moved"
	EventTypeNpcStateChanged             EventType = "npc_state_changed"
	EventTypeNpcDied                     EventType = "npc_died"
	EventTypeLocationChanged             EventType = "location_changed"
	EventTypeFactionChanged              EventType = "faction_changed"
	EventTypeClaimShared                 EventType = "claim_shared"
	EventTypeNpcPlanRequested            EventType = "npc_plan_requested"
	EventTypePlayerCharacterReady        EventType = "player_character_ready"
	EventTypePlayerCharacterFailed       EventType = "player_character_failed"
)

var EventTypeMap = commonenum.NewMapping[EventType, v1.GameTownEventType](
	map[EventType]commonenum.Entry[EventType, v1.GameTownEventType]{
		EventTypeWorldGenerationRequested:    {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_GENERATION_REQUESTED},
		EventTypeWorldReady:                  {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_READY},
		EventTypeWorldGenerationFailed:       {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_GENERATION_FAILED},
		EventTypePlayerJoined:                {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_JOINED},
		EventTypePlayerMoved:                 {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_MOVED},
		EventTypePlayerTalked:                {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_TALKED},
		EventTypeNpcThinking:                 {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_THINKING},
		EventTypeNpcReplied:                  {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_REPLIED},
		EventTypePlayerActed:                 {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_ACTED},
		EventTypeActionResolved:              {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_ACTION_RESOLVED},
		EventTypeWorldTickRequested:          {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_TICK_REQUESTED},
		EventTypeWorldEvolved:                {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_WORLD_EVOLVED},
		EventTypeAgentJobFailed:              {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_AGENT_JOB_FAILED},
		EventTypePlayerActionSubmitted:       {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_ACTION_SUBMITTED},
		EventTypeActionRejected:              {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_ACTION_REJECTED},
		EventTypeActionClarificationRequired: {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_ACTION_CLARIFICATION_REQUIRED},
		EventTypeNpcPlanned:                  {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_PLANNED},
		EventTypeNpcMoved:                    {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_MOVED},
		EventTypeNpcStateChanged:             {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_STATE_CHANGED},
		EventTypeNpcDied:                     {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_DIED},
		EventTypeLocationChanged:             {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_LOCATION_CHANGED},
		EventTypeFactionChanged:              {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_FACTION_CHANGED},
		EventTypeClaimShared:                 {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_CLAIM_SHARED},
		EventTypeNpcPlanRequested:            {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_NPC_PLAN_REQUESTED},
		EventTypePlayerCharacterReady:        {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_CHARACTER_READY},
		EventTypePlayerCharacterFailed:       {Proto: v1.GameTownEventType_GAME_TOWN_EVENT_TYPE_PLAYER_CHARACTER_FAILED},
	},
)

func (e EventType) String() string {
	return string(e)
}
