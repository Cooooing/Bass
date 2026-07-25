package usecase

import (
	"context"
	"testing"
	"time"

	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/config"
	"game_town/internal/enum"
)

type agentClientSpy struct {
	repo.AgentClient
	actCalled     bool
	planNpcCalled bool
	tickCalled    bool
}

func (c *agentClientSpy) Act(context.Context, *repo.ActReq) (*model.ActionResolution, error) {
	c.actCalled = true
	return &model.ActionResolution{
		Status:  "resolved",
		Summary: "玩家行动被模型裁决",
	}, nil
}

func (c *agentClientSpy) PlanNpc(context.Context, *repo.PlanNpcReq) (*model.NpcPlan, error) {
	c.planNpcCalled = true
	return &model.NpcPlan{
		Summary:        "NPC 生成了自主计划",
		Goal:           "继续调查",
		NextDecisionIn: 60,
	}, nil
}

func (c *agentClientSpy) Tick(context.Context, *repo.TickReq) (*model.ActionResolution, error) {
	c.tickCalled = true
	return &model.ActionResolution{
		Status:  "resolved",
		Summary: "世界根据近期事件继续变化",
	}, nil
}

type playerRepoCallStub struct {
	repo.PlayerRepo
	player *model.Player
}

func (r *playerRepoCallStub) Get(context.Context, *repo.PlayerQuery) (*model.Player, error) {
	return r.player, nil
}

type locationRepoCallStub struct {
	repo.LocationRepo
	location  *model.Location
	locations []*model.Location
}

func (r *locationRepoCallStub) Get(context.Context, *repo.LocationQuery) (*model.Location, error) {
	return r.location, nil
}

func (r *locationRepoCallStub) List(context.Context, *repo.LocationQuery) ([]*model.Location, error) {
	return r.locations, nil
}

type npcRepoCallStub struct {
	repo.NpcRepo
	npc  *model.Npc
	npcs []*model.Npc
}

func (r *npcRepoCallStub) Get(context.Context, *repo.NpcQuery) (*model.Npc, error) {
	return r.npc, nil
}

func (r *npcRepoCallStub) List(context.Context, *repo.NpcQuery) ([]*model.Npc, error) {
	return r.npcs, nil
}

type factionRepoCallStub struct {
	repo.FactionRepo
	factions []*model.Faction
}

func (r *factionRepoCallStub) List(context.Context, *repo.FactionQuery) ([]*model.Faction, error) {
	return r.factions, nil
}

type observationRepoCallStub struct {
	repo.ObservationRepo
}

func (r *observationRepoCallStub) List(context.Context, *repo.ObservationQuery) ([]*model.Observation, error) {
	return nil, nil
}

type eventRepoCallStub struct {
	repo.EventRepo
	events []*model.Event
}

func (r *eventRepoCallStub) List(context.Context, *repo.EventQuery) ([]*model.Event, error) {
	return r.events, nil
}

type npcMemoryRepoCallStub struct {
	repo.NpcMemoryRepo
}

func (r *npcMemoryRepoCallStub) List(context.Context, *repo.NpcMemoryQuery) ([]*model.NpcMemory, error) {
	return nil, nil
}

func TestCallAgentUsesModelForPlayerAction(t *testing.T) {
	client := &agentClientSpy{}
	runner := newCallAgentTestRunner(client)
	result := newCallAgentTestResult(enum.AgentJobTypePlayerActionInterpret)
	result.source.ActorPlayerID = newInt64(7)
	result.source.Content = "我调查最近的失踪事件"

	if err := runner.callAgent(context.Background(), result); err != nil {
		t.Fatalf("callAgent() error = %v", err)
	}
	if !client.actCalled {
		t.Fatalf("expected AgentClient.Act to be called")
	}
	if result.resolution == nil || result.resolution.Summary == "" {
		t.Fatalf("expected model action resolution, got %#v", result.resolution)
	}
}

func TestCallAgentUsesModelForNpcPlan(t *testing.T) {
	client := &agentClientSpy{}
	runner := newCallAgentTestRunner(client)
	result := newCallAgentTestResult(enum.AgentJobTypeNpcPlan)
	npcID := int64(3)
	result.source.NpcID = new(npcID)
	result.job.NpcID = new(npcID)

	if err := runner.callAgent(context.Background(), result); err != nil {
		t.Fatalf("callAgent() error = %v", err)
	}
	if !client.planNpcCalled {
		t.Fatalf("expected AgentClient.PlanNpc to be called")
	}
	if result.plan == nil || result.plan.Summary == "" {
		t.Fatalf("expected model npc plan, got %#v", result.plan)
	}
}

func TestCallAgentUsesModelForWorldTick(t *testing.T) {
	client := &agentClientSpy{}
	runner := newCallAgentTestRunner(client)
	result := newCallAgentTestResult(enum.AgentJobTypeWorldTick)

	if err := runner.callAgent(context.Background(), result); err != nil {
		t.Fatalf("callAgent() error = %v", err)
	}
	if !client.tickCalled {
		t.Fatalf("expected AgentClient.Tick to be called")
	}
	if result.resolution == nil || result.resolution.Summary == "" {
		t.Fatalf("expected model world tick resolution, got %#v", result.resolution)
	}
}

func newCallAgentTestRunner(client repo.AgentClient) *WorldAgentRunner {
	location := &model.Location{
		ID:         2,
		WorldID:    1,
		Name:       "中心据点",
		Accessible: true,
	}
	npc := &model.Npc{
		ID:                3,
		WorldID:           1,
		Name:              "巡游者",
		Goal:              "调查异常",
		CurrentLocationID: location.ID,
	}
	return &WorldAgentRunner{
		conf: &config.Bootstrap{
			GameTown: &config.GameTown{
				Agent: &config.Agent{
					RecentEventLimit: 20,
				},
				Memory: &config.Memory{
					EmbeddingEnabled: false,
				},
			},
		},
		agentClient: client,
		playerRepo: &playerRepoCallStub{
			player: &model.Player{
				ID:          7,
				DisplayName: "玩家",
			},
		},
		worldMemberRepo: &worldMemberRepoStub{
			member: &model.WorldMember{
				ID:                8,
				WorldID:           1,
				PlayerID:          7,
				CurrentLocationID: location.ID,
			},
		},
		locationRepo: &locationRepoCallStub{
			location:  location,
			locations: []*model.Location{location},
		},
		npcRepo: &npcRepoCallStub{
			npc:  npc,
			npcs: []*model.Npc{npc},
		},
		factionRepo: &factionRepoCallStub{
			factions: []*model.Faction{{ID: 4, WorldID: 1, Name: "议会"}},
		},
		observationRepo: &observationRepoCallStub{},
		eventRepo: &eventRepoCallStub{
			events: []*model.Event{{ID: 11, WorldID: 1, Sequence: 1, Summary: "近期事件"}},
		},
		npcMemoryRepo: &npcMemoryRepoCallStub{},
	}
}

func newCallAgentTestResult(jobType enum.AgentJobType) *agentResult {
	now := time.Now()
	return &agentResult{
		job: &model.AgentJob{
			ID:            1,
			WorldID:       1,
			Type:          jobType,
			SourceEventID: 11,
		},
		source: &model.Event{
			ID:         11,
			WorldID:    1,
			Type:       enum.EventTypePlayerActionSubmitted,
			Payload:    map[string]any{},
			OccurredAt: now,
			WorldTime:  now,
		},
		world: &model.World{
			ID:            1,
			Name:          "测试世界",
			AgentConfigID: 1,
		},
		state: &model.WorldState{
			WorldID:    1,
			WorldTime:  now,
			Summary:    "世界摘要",
			CurrentArc: "当前阶段",
		},
		config: &model.AgentConfig{
			ID: 1,
		},
	}
}

func newInt64(value int64) *int64 {
	return &value
}
