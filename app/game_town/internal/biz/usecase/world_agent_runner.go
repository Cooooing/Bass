package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"common/pkg/constant"
	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/config"
	"game_town/internal/enum"

	"github.com/samber/lo"
)

type worldLoop struct {
	worldID int64
	wake    chan struct{}
	results chan *agentResult
}

const worldEventConsumeBatchSize = 64

type agentResult struct {
	job        *model.AgentJob
	source     *model.Event
	world      *model.World
	state      *model.WorldState
	config     *model.AgentConfig
	player     *model.Player
	member     *model.WorldMember
	location   *model.Location
	npc        *model.Npc
	draft      *model.WorldDraft
	reply      *model.NpcReply
	resolution *model.ActionResolution
	plan       *model.NpcPlan
	character  *model.CharacterDraft
	memories   []*model.NpcMemory
	memory     *model.NpcMemory
	embedding  []float32
	npcs       []*model.Npc
	locations  []*model.Location
	factions   []*model.Faction
	sideEvents []*model.Event
	err        error
	done       chan struct{}
}

// WorldAgentRunner 负责恢复并驱动每个世界的独立 Agent Loop。
type WorldAgentRunner struct {
	log                   *slog.Logger
	conf                  *config.Bootstrap
	tx                    base.Tx
	worldRepo             repo.WorldRepo
	playerRepo            repo.PlayerRepo
	worldMemberRepo       repo.WorldMemberRepo
	locationRepo          repo.LocationRepo
	npcRepo               repo.NpcRepo
	worldStateRepo        repo.WorldStateRepo
	observationRepo       repo.ObservationRepo
	npcMemoryRepo         repo.NpcMemoryRepo
	claimRepo             repo.ClaimRepo
	npcBeliefRepo         repo.NpcBeliefRepo
	relationshipRepo      repo.RelationshipRepo
	factionRepo           repo.FactionRepo
	factionMembershipRepo repo.FactionMembershipRepo
	worldRuleRepo         repo.WorldRuleRepo
	embeddingClient       repo.EmbeddingClient
	agentConfigRepo       repo.AgentConfigRepo
	eventRepo             repo.EventRepo
	agentJobRepo          repo.AgentJobRepo
	agentClient           repo.AgentClient
	eventNotifier         repo.EventNotifier
	eventUsecase          *EventUsecase

	mu              sync.Mutex
	loops           map[int64]*worldLoop
	cancel          context.CancelFunc
	active          int
	activeMemory    int
	activeWorld     map[int64]int
	activeConfig    map[int64]int
	lanes           map[string]bool
	breakerFailures map[int64]int
	breakerUntil    map[int64]time.Time
	scheduleCursor  int
	schedulerWake   chan struct{}
	tickMu          sync.Mutex
	tickTimers      map[int64]*time.Timer
}

func NewWorldAgentRunner(
	log *slog.Logger,
	conf *config.Bootstrap,
	tx base.Tx,
	worldRepo repo.WorldRepo,
	playerRepo repo.PlayerRepo,
	worldMemberRepo repo.WorldMemberRepo,
	locationRepo repo.LocationRepo,
	npcRepo repo.NpcRepo,
	worldStateRepo repo.WorldStateRepo,
	agentConfigRepo repo.AgentConfigRepo,
	observationRepo repo.ObservationRepo,
	npcMemoryRepo repo.NpcMemoryRepo,
	claimRepo repo.ClaimRepo,
	npcBeliefRepo repo.NpcBeliefRepo,
	relationshipRepo repo.RelationshipRepo,
	factionRepo repo.FactionRepo,
	factionMembershipRepo repo.FactionMembershipRepo,
	worldRuleRepo repo.WorldRuleRepo,
	embeddingClient repo.EmbeddingClient,
	eventRepo repo.EventRepo,
	agentJobRepo repo.AgentJobRepo,
	agentClient repo.AgentClient,
	eventNotifier repo.EventNotifier,
	eventUsecase *EventUsecase,
) *WorldAgentRunner {
	return &WorldAgentRunner{
		log:                   log,
		conf:                  conf,
		tx:                    tx,
		worldRepo:             worldRepo,
		playerRepo:            playerRepo,
		worldMemberRepo:       worldMemberRepo,
		locationRepo:          locationRepo,
		npcRepo:               npcRepo,
		worldStateRepo:        worldStateRepo,
		agentConfigRepo:       agentConfigRepo,
		observationRepo:       observationRepo,
		npcMemoryRepo:         npcMemoryRepo,
		claimRepo:             claimRepo,
		npcBeliefRepo:         npcBeliefRepo,
		relationshipRepo:      relationshipRepo,
		factionRepo:           factionRepo,
		factionMembershipRepo: factionMembershipRepo,
		worldRuleRepo:         worldRuleRepo,
		embeddingClient:       embeddingClient,
		eventRepo:             eventRepo,
		agentJobRepo:          agentJobRepo,
		agentClient:           agentClient,
		eventNotifier:         eventNotifier,
		eventUsecase:          eventUsecase,
		loops:                 make(map[int64]*worldLoop),
		activeWorld:           make(map[int64]int),
		activeConfig:          make(map[int64]int),
		lanes:                 make(map[string]bool),
		breakerFailures:       make(map[int64]int),
		breakerUntil:          make(map[int64]time.Time),
		schedulerWake:         make(chan struct{}, 1),
		tickTimers:            make(map[int64]*time.Timer),
	}
}

func (r *WorldAgentRunner) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	r.cancel = cancel

	running := enum.AgentJobStatusRunning
	jobs, err := r.agentJobRepo.List(runCtx, &repo.AgentJobQuery{Status: new(running)})
	if err != nil {
		cancel()
		return err
	}
	for _, job := range jobs {
		if _, err := r.agentJobRepo.Retry(runCtx, &repo.AgentJobRetryReq{
			JobID:        job.ID,
			AttemptCount: job.AttemptCount,
			AvailableAt:  time.Now(),
			ErrorSummary: "recovered after restart",
		}); err != nil {
			cancel()
			return err
		}
	}
	worlds, err := r.worldRepo.List(runCtx, nil)
	if err != nil {
		cancel()
		return err
	}
	for _, world := range worlds {
		if world.Status != enum.WorldStatusArchived {
			r.ensureLoop(runCtx, world.ID)
		}
	}
	all, unsubscribe := r.eventNotifier.SubscribeAll()
	go func() {
		defer unsubscribe()
		for {
			select {
			case <-runCtx.Done():
				return
			case worldID := <-all:
				loop := r.ensureLoop(runCtx, worldID)
				select {
				case loop.wake <- struct{}{}:
				default:
				}
			}
		}
	}()
	go r.recoverWorldLoops(runCtx)
	go r.schedule(runCtx)
	go r.scheduleTicks(runCtx)
	return nil
}

func (r *WorldAgentRunner) Stop(_ context.Context) error {
	if r.cancel != nil {
		r.cancel()
	}
	return nil
}

func (r *WorldAgentRunner) ensureLoop(ctx context.Context, worldID int64) *worldLoop {
	r.mu.Lock()
	defer r.mu.Unlock()

	if loop := r.loops[worldID]; loop != nil {
		return loop
	}
	loop := &worldLoop{
		worldID: worldID,
		wake:    make(chan struct{}, 1),
		results: make(chan *agentResult, 16),
	}
	loop.wake <- struct{}{}
	r.loops[worldID] = loop
	go r.runWorld(ctx, loop)
	return loop
}

func (r *WorldAgentRunner) runWorld(ctx context.Context, loop *worldLoop) {
	for {
		select {
		case <-ctx.Done():
			return
		case result := <-loop.results:
			r.applyResult(ctx, result)
			if result.done != nil {
				close(result.done)
			}
			continue
		default:
		}

		select {
		case <-ctx.Done():
			return
		case <-loop.wake:
			if r.consume(ctx, loop.worldID) {
				select {
				case loop.wake <- struct{}{}:
				default:
				}
			}
		case result := <-loop.results:
			r.applyResult(ctx, result)
			if result.done != nil {
				close(result.done)
			}
		}
	}
}

func (r *WorldAgentRunner) recoverWorldLoops(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			r.mu.Lock()
			loops := lo.Values(r.loops)
			r.mu.Unlock()
			for _, loop := range loops {
				select {
				case loop.wake <- struct{}{}:
				default:
				}
			}
		}
	}
}

func (r *WorldAgentRunner) consume(ctx context.Context, worldID int64) bool {
	state, err := r.worldStateRepo.Get(ctx, &repo.WorldStateQuery{WorldID: new(worldID)})
	if err != nil {
		r.log.ErrorContext(ctx, "load world state failed", "world_id", worldID, constant.LogFieldErr, err)
		return false
	}
	events, err := r.eventRepo.List(ctx, &repo.EventQuery{
		WorldID:       new(worldID),
		AfterSequence: new(state.AgentCursor),
		Limit:         worldEventConsumeBatchSize,
	})
	if err != nil {
		r.log.ErrorContext(ctx, "load world events failed", "world_id", worldID, constant.LogFieldErr, err)
		return false
	}
	for _, event := range events {
		r.log.InfoContext(ctx, "consume world event", "world_id", worldID, constant.LogFieldEventID, event.ID, "sequence", event.Sequence, "type", event.Type)
		var jobType enum.AgentJobType
		var priority enum.AgentJobPriority
		laneKey := ""
		switch event.Type {
		case enum.EventTypeWorldGenerationRequested:
			jobType = enum.AgentJobTypeWorldGenerate
			priority = enum.AgentJobPriorityHigh
			laneKey = "world"
		case enum.EventTypePlayerJoined:
			jobType = enum.AgentJobTypePlayerCharacterGenerate
			priority = enum.AgentJobPriorityHigh
			if event.ActorPlayerID != nil {
				laneKey = fmt.Sprintf("player:%d", *event.ActorPlayerID)
			}
		case enum.EventTypePlayerTalked:
			jobType = enum.AgentJobTypeNpcTalk
			priority = enum.AgentJobPriorityHigh
			if event.NpcID != nil {
				laneKey = fmt.Sprintf("npc:%d", *event.NpcID)
			}
		case enum.EventTypePlayerActed, enum.EventTypePlayerActionSubmitted:
			jobType = enum.AgentJobTypePlayerActionInterpret
			priority = enum.AgentJobPriorityHigh
			laneKey = "world"
			if event.NpcID != nil {
				laneKey = fmt.Sprintf("npc:%d", *event.NpcID)
			}
		case enum.EventTypeWorldTickRequested:
			jobType = enum.AgentJobTypeWorldTick
			priority = enum.AgentJobPriorityLow
			laneKey = "world"
		case enum.EventTypeNpcPlanRequested:
			jobType = enum.AgentJobTypeNpcPlan
			priority = enum.AgentJobPriorityLow
			laneKey = fmt.Sprintf("npc:%d", *event.NpcID)
		}
		var thinking *model.Event
		var fastEvent *model.Event
		jobCreated := false
		err = r.tx(ctx, func(ctx context.Context) error {
			var handled bool
			fastEvent, handled, err = r.resolveFastPlayerActionInTx(ctx, event)
			if err != nil {
				return err
			}
			if laneKey != "" {
				if handled {
					return r.worldStateRepo.AdvanceCursor(ctx, worldID, event.Sequence)
				}
				count, err := r.agentJobRepo.Count(ctx, &repo.AgentJobQuery{
					SourceEventID: new(event.ID),
				})
				if err != nil {
					return err
				}
				if count == 0 {
					if _, err = r.agentJobRepo.Save(ctx, &model.AgentJob{
						WorldID:       worldID,
						SourceEventID: event.ID,
						Type:          jobType,
						Priority:      priority,
						LaneKey:       laneKey,
						Status:        enum.AgentJobStatusQueued,
						WorldVersion:  state.Version,
						NpcID:         event.NpcID,
						AvailableAt:   time.Now(),
					}); err != nil {
						return err
					}
					jobCreated = true
					if event.Type == enum.EventTypePlayerTalked || event.Type == enum.EventTypePlayerActionSubmitted && event.NpcID != nil {
						thinking, err = r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
							WorldID:          worldID,
							Type:             enum.EventTypeNpcThinking,
							ActorPlayerID:    event.ActorPlayerID,
							NpcID:            event.NpcID,
							LocationID:       event.LocationID,
							CausationEventID: new(event.ID),
							Summary:          "NPC 正在思考",
						})
						if err != nil {
							return err
						}
					}
				}
			}
			return r.worldStateRepo.AdvanceCursor(ctx, worldID, event.Sequence)
		})
		if err != nil {
			r.log.ErrorContext(
				ctx,
				"consume world event failed",
				"world_id",
				worldID,
				constant.LogFieldEventID,
				event.ID,
				constant.LogFieldErr,
				err,
			)
			return false
		}
		if jobCreated {
			r.log.InfoContext(ctx, "agent job queued from event", "world_id", worldID, constant.LogFieldEventID, event.ID, "job_type", jobType, "lane_key", laneKey)
			r.wakeScheduler()
		}
		if thinking != nil {
			r.eventUsecase.Publish(thinking)
		}
		if fastEvent != nil {
			r.eventUsecase.Publish(fastEvent)
		}
	}
	return len(events) == worldEventConsumeBatchSize
}

func (r *WorldAgentRunner) resolveFastPlayerActionInTx(ctx context.Context, event *model.Event) (*model.Event, bool, error) {
	if event.Type != enum.EventTypePlayerActionSubmitted || event.ActorPlayerID == nil || event.NpcID != nil {
		return nil, false, nil
	}
	member, err := r.worldMemberRepo.Get(ctx, &repo.WorldMemberQuery{
		WorldID:  new(event.WorldID),
		PlayerID: event.ActorPlayerID,
	})
	if err != nil {
		return nil, true, err
	}

	targets := actionTargets(event.Payload)
	var locationID int64
	for _, target := range targets {
		if target.Type == enum.EntityTypeLocation {
			locationID = target.ID
			break
		}
	}
	if locationID <= 0 {
		return nil, false, nil
	}

	location, err := r.locationRepo.Get(ctx, &repo.LocationQuery{ID: new(locationID)})
	if err != nil {
		rejected, appendErr := r.appendFastActionEvent(ctx, event, enum.EventTypeActionRejected, "目标地点不存在", "你暂时无法前往这个地点。", member.CurrentLocationID)
		return rejected, true, appendErr
	}
	if location.WorldID != event.WorldID || !location.Accessible || location.Status != enum.LocationStatusActive {
		rejected, appendErr := r.appendFastActionEvent(ctx, event, enum.EventTypeActionRejected, "目标地点不可达", "道路暂时不可通行，世界会继续演化。", member.CurrentLocationID)
		return rejected, true, appendErr
	}
	if member.CurrentLocationID != location.ID {
		if _, err = r.worldMemberRepo.Move(ctx, member.ID, location.ID); err != nil {
			return nil, true, err
		}
	}
	summary := "玩家移动到 " + location.Name
	content := "你前往 " + location.Name + "。这里的局势会继续随世界事件变化。"
	resolved, err := r.appendFastActionEvent(ctx, event, enum.EventTypeActionResolved, summary, content, location.ID)
	return resolved, true, err
}

func (r *WorldAgentRunner) appendFastActionEvent(ctx context.Context, source *model.Event, eventType enum.EventType, summary string, content string, locationID int64) (*model.Event, error) {
	return r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
		WorldID:          source.WorldID,
		Type:             eventType,
		ActorPlayerID:    source.ActorPlayerID,
		LocationID:       new(locationID),
		CausationEventID: new(source.ID),
		Summary:          summary,
		Content:          content,
		Payload: map[string]any{
			"fast_resolved": true,
			"source":        "world_loop",
		},
	})
}
