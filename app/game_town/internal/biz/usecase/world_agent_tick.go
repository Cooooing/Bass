package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"common/pkg/constant"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

func (r *WorldAgentRunner) scheduleTicks(ctx context.Context) {
	r.restoreTickTimers(ctx)
	recoveryTicker := time.NewTicker(time.Minute)
	defer recoveryTicker.Stop()
	defer r.stopTickTimers()

	for {
		select {
		case <-ctx.Done():
			return
		case <-recoveryTicker.C:
			r.restoreTickTimers(ctx)
		}
	}
}

func (r *WorldAgentRunner) restoreTickTimers(ctx context.Context) {
	status := enum.WorldStatusActive
	worlds, err := r.worldRepo.List(ctx, &repo.WorldQuery{
		Status: new(status),
	})
	if err != nil {
		r.log.ErrorContext(ctx, "load active world tick timers failed", constant.LogFieldErr, err)
		return
	}
	for _, world := range worlds {
		state, err := r.worldStateRepo.Get(ctx, &repo.WorldStateQuery{
			WorldID: new(world.ID),
		})
		if err != nil {
			r.log.ErrorContext(ctx, "load active world tick state failed", "world_id", world.ID, constant.LogFieldErr, err)
			continue
		}
		r.armWorldTimer(ctx, state.WorldID, state.NextDueAt)
	}
}

func (r *WorldAgentRunner) refreshWorldTimer(ctx context.Context, worldID int64) {
	state, err := r.worldStateRepo.Get(ctx, &repo.WorldStateQuery{
		WorldID: new(worldID),
	})
	if err != nil {
		r.log.ErrorContext(ctx, "refresh world tick timer failed", "world_id", worldID, constant.LogFieldErr, err)
		return
	}
	r.armWorldTimer(ctx, worldID, state.NextDueAt)
}

func (r *WorldAgentRunner) armWorldTimer(ctx context.Context, worldID int64, dueAt *time.Time) {
	if dueAt == nil {
		return
	}
	delay := time.Until(*dueAt)
	if delay < time.Second {
		delay = time.Second
	}

	r.tickMu.Lock()
	if timer := r.tickTimers[worldID]; timer != nil {
		timer.Stop()
	}
	r.tickTimers[worldID] = time.AfterFunc(delay, func() {
		r.triggerWorldTick(ctx, worldID)
	})
	r.tickMu.Unlock()
}

func (r *WorldAgentRunner) stopTickTimers() {
	r.tickMu.Lock()
	defer r.tickMu.Unlock()
	for worldID, timer := range r.tickTimers {
		timer.Stop()
		delete(r.tickTimers, worldID)
	}
}

func (r *WorldAgentRunner) triggerWorldTick(ctx context.Context, worldID int64) {
	if ctx.Err() != nil {
		return
	}
	state, err := r.worldStateRepo.Get(ctx, &repo.WorldStateQuery{
		WorldID: new(worldID),
	})
	if err != nil {
		r.log.ErrorContext(ctx, "load world tick state failed", "world_id", worldID, constant.LogFieldErr, err)
		return
	}
	if state.NextDueAt != nil && time.Now().Before(*state.NextDueAt) {
		r.armWorldTimer(ctx, worldID, state.NextDueAt)
		return
	}
	if err = r.scheduleWorldTick(ctx, state, time.Now()); err != nil {
		r.log.ErrorContext(ctx, "schedule world tick failed", "world_id", worldID, constant.LogFieldErr, err)
	}
	r.refreshWorldTimer(ctx, worldID)
}

func (r *WorldAgentRunner) scheduleWorldTick(ctx context.Context, state *model.WorldState, now time.Time) error {
	pending, err := r.hasPendingTickWork(ctx, state)
	if err != nil {
		return err
	}
	if pending {
		next := now.Add(r.tickInterval())
		return r.worldStateRepo.UpdateNextDue(ctx, state.WorldID, new(next))
	}

	var events []*model.Event
	err = r.tx(ctx, func(ctx context.Context) error {
		worldTime := r.nextWorldTime(state, now)
		nextTickAt := now.Add(r.tickInterval())
		updated, err := r.worldStateRepo.AdvanceTime(ctx, &repo.WorldStateAdvanceTimeReq{
			WorldID:    state.WorldID,
			Version:    state.Version,
			WorldTime:  worldTime,
			TimeAnchor: now,
			NextDueAt:  new(nextTickAt),
		})
		if err != nil {
			return err
		}
		if err = r.worldStateRepo.UpdateNextTick(ctx, state.WorldID, nextTickAt); err != nil {
			return err
		}
		dueNpcs, err := r.npcRepo.List(ctx, &repo.NpcQuery{
			WorldID:            new(state.WorldID),
			NextDecisionBefore: new(updated.WorldTime),
		})
		if err != nil {
			return err
		}
		for _, npc := range dueNpcs {
			pending, err := r.hasPendingNpcPlanWork(ctx, state.WorldID, npc.ID, state.AgentCursor)
			if err != nil {
				return err
			}
			if pending {
				continue
			}

			event, err := r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
				WorldID:    state.WorldID,
				Type:       enum.EventTypeNpcPlanRequested,
				NpcID:      new(npc.ID),
				LocationID: new(npc.CurrentLocationID),
				Summary:    "请求 NPC 自主计划",
			})
			if err != nil {
				return err
			}
			events = append(events, event)
		}

		event, err := r.eventUsecase.AppendInTx(ctx, &AppendEventReq{
			WorldID: state.WorldID,
			Type:    enum.EventTypeWorldTickRequested,
			Summary: "请求世界演进",
			Payload: map[string]any{"public": true},
		})
		if err != nil {
			return err
		}
		events = append(events, event)
		return nil
	})
	if err != nil {
		return err
	}

	for _, event := range events {
		r.eventUsecase.Publish(event)
	}
	return nil
}

func (r *WorldAgentRunner) hasPendingNpcPlanWork(ctx context.Context, worldID int64, npcID int64, cursor uint64) (bool, error) {
	statuses := []enum.AgentJobStatus{
		enum.AgentJobStatusQueued,
		enum.AgentJobStatusRunning,
	}
	jobType := enum.AgentJobTypeNpcPlan
	laneKey := fmt.Sprintf("npc:%d", npcID)
	jobCount, jobErr := r.agentJobRepo.Count(ctx, &repo.AgentJobQuery{
		WorldID:  new(worldID),
		Type:     new(jobType),
		Statuses: statuses,
		LaneKey:  new(laneKey),
	})
	eventType := enum.EventTypeNpcPlanRequested
	pendingCount, pendingErr := r.eventRepo.Count(ctx, &repo.EventQuery{
		WorldID:       new(worldID),
		Type:          new(eventType),
		NpcID:         new(npcID),
		AfterSequence: new(cursor),
	})
	if jobErr != nil || pendingErr != nil {
		return false, errors.Join(jobErr, pendingErr)
	}
	return jobCount > 0 || pendingCount > 0, nil
}

func (r *WorldAgentRunner) hasPendingTickWork(ctx context.Context, state *model.WorldState) (bool, error) {
	statuses := []enum.AgentJobStatus{
		enum.AgentJobStatusQueued,
		enum.AgentJobStatusRunning,
	}
	jobType := enum.AgentJobTypeWorldTick
	tickEvent := enum.EventTypeWorldTickRequested
	jobCount, jobErr := r.agentJobRepo.Count(ctx, &repo.AgentJobQuery{
		WorldID:  new(state.WorldID),
		Type:     new(jobType),
		Statuses: statuses,
	})
	pendingCount, pendingErr := r.eventRepo.Count(ctx, &repo.EventQuery{
		WorldID:       new(state.WorldID),
		Type:          new(tickEvent),
		AfterSequence: new(state.AgentCursor),
	})
	if jobErr != nil || pendingErr != nil {
		return false, errors.Join(jobErr, pendingErr)
	}
	return jobCount+pendingCount > 0, nil
}

func (r *WorldAgentRunner) nextWorldTime(state *model.WorldState, now time.Time) time.Time {
	scale := state.TimeScale
	if scale == 0 {
		scale = 24
	}
	elapsed := now.Sub(state.TimeAnchor)
	return state.WorldTime.Add(time.Duration(float64(elapsed) * float64(scale)))
}

func (r *WorldAgentRunner) recentEvents(ctx context.Context, worldID int64) ([]*model.Event, error) {
	limit := int(r.conf.GetGameTown().GetAgent().GetRecentEventLimit())
	if limit <= 0 {
		limit = 20
	}
	return r.eventRepo.List(ctx, &repo.EventQuery{
		WorldID:     new(worldID),
		RecentLimit: limit,
	})
}

func (r *WorldAgentRunner) recordFailure(config *model.AgentConfig) {
	if config == nil {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.breakerFailures[config.ID]++
	if r.breakerFailures[config.ID] < 3 {
		return
	}

	duration := 15 * time.Second
	r.breakerUntil[config.ID] = time.Now().Add(duration)
	r.breakerFailures[config.ID] = 0
	time.AfterFunc(duration, r.wakeScheduler)
}

func (r *WorldAgentRunner) recordSuccess(configID int64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.breakerFailures, configID)
	delete(r.breakerUntil, configID)
}

func (r *WorldAgentRunner) globalLimit() int {
	value := int(r.conf.GetGameTown().GetAgent().GetGlobalConcurrency())
	if value <= 0 {
		return 4
	}
	return value
}

func (r *WorldAgentRunner) modelLimit() int {
	value := int(r.conf.GetGameTown().GetAgent().GetModelConcurrency())
	if value <= 0 {
		return 2
	}
	return value
}

func (r *WorldAgentRunner) worldLimit() int {
	value := int(r.conf.GetGameTown().GetAgent().GetWorldConcurrency())
	if value <= 0 {
		return 2
	}
	return value
}

func (r *WorldAgentRunner) memoryLimit() int {
	return 1
}

func (r *WorldAgentRunner) embeddingEnabled() bool {
	return r.conf.GetGameTown().GetMemory().GetEmbeddingEnabled()
}

func (r *WorldAgentRunner) maxRetry() int {
	value := int(r.conf.GetGameTown().GetAgent().GetMaxRetry())
	if value <= 0 {
		return 2
	}
	return value
}

func (r *WorldAgentRunner) retryDelay() time.Duration {
	delay := r.conf.GetGameTown().GetAgent().GetRetryBaseDelay()
	if delay != nil && delay.AsDuration() >= 5*time.Second {
		return delay.AsDuration()
	}
	return 5 * time.Second
}

func (r *WorldAgentRunner) staleJobDuration() time.Duration {
	timeout := time.Duration(0)
	configs, err := r.agentConfigRepo.List(context.Background(), nil)
	if err == nil {
		for _, config := range configs {
			value := time.Duration(config.TimeoutSeconds) * time.Second
			if value > timeout {
				timeout = value
			}
		}
	}
	if timeout <= 0 {
		timeout = time.Minute
	}
	return timeout + time.Minute
}

func (r *WorldAgentRunner) tickInterval() time.Duration {
	interval := r.conf.GetGameTown().GetAgent().GetTickInterval()
	if interval != nil && interval.AsDuration() > 0 {
		return interval.AsDuration()
	}
	return 5 * time.Minute
}

func (r *WorldAgentRunner) tickScanInterval() time.Duration {
	interval := r.tickInterval() / 4
	if interval < time.Second {
		return time.Second
	}
	if interval > 15*time.Second {
		return 15 * time.Second
	}
	return interval
}

func (r *WorldAgentRunner) uint32Value(values map[string]any, key string) uint32 {
	value, ok := values[key].(float64)
	if !ok {
		return 0
	}
	return uint32(value)
}

func (r *WorldAgentRunner) truncateError(err error) string {
	if err == nil {
		return ""
	}
	value := err.Error()
	if len(value) > 512 {
		return value[:512]
	}
	return value
}
