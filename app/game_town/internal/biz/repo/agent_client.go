package repo

import (
	"context"

	"game_town/internal/biz/model"
)

type AgentClient interface {
	GenerateWorld(context.Context, *GenerateWorldReq) (*model.WorldDraft, error)
	GenerateCharacter(context.Context, *GenerateCharacterReq) (*model.CharacterDraft, error)
	Talk(context.Context, *TalkReq) (*model.NpcReply, error)
	Act(context.Context, *ActReq) (*model.ActionResolution, error)
	PlanNpc(context.Context, *PlanNpcReq) (*model.NpcPlan, error)
	Tick(context.Context, *TickReq) (*model.ActionResolution, error)
}

type GenerateWorldReq struct {
	Config                  *model.AgentConfig
	World                   *model.World
	NpcCount, LocationCount uint32
}

type GenerateCharacterReq struct {
	Config       *model.AgentConfig
	World        *model.World
	State        *model.WorldState
	Player       *model.Player
	Member       *model.WorldMember
	Location     *model.Location
	RecentEvents []*model.Event
	Preference   string
}

type TalkReq struct {
	Config       *model.AgentConfig
	World        *model.World
	State        *model.WorldState
	Player       *model.Player
	Member       *model.WorldMember
	Location     *model.Location
	Npc          *model.Npc
	RecentEvents []*model.Event
	Memories     []*model.NpcMemory
	Content      string
}

type ActReq struct {
	Config       *model.AgentConfig
	World        *model.World
	State        *model.WorldState
	Player       *model.Player
	Member       *model.WorldMember
	Location     *model.Location
	Npc          *model.Npc
	RecentEvents []*model.Event
	Memories     []*model.NpcMemory
	Content      string
	Targets      []model.EntityRef
}

type PlanNpcReq struct {
	Config       *model.AgentConfig
	World        *model.World
	State        *model.WorldState
	Location     *model.Location
	Npc          *model.Npc
	RecentEvents []*model.Event
	Memories     []*model.NpcMemory
}

type TickReq struct {
	Config       *model.AgentConfig
	World        *model.World
	State        *model.WorldState
	RecentEvents []*model.Event
	Npcs         []*model.Npc
	Locations    []*model.Location
	Factions     []*model.Faction
}
