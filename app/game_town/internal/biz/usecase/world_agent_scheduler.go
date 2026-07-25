package usecase

import (
	"context"
	"fmt"
	"time"

	"common/pkg/constant"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

func (r *WorldAgentRunner) wakeScheduler() {
	select {
	case r.schedulerWake <- struct{}{}:
	default:
	}
}

func (r *WorldAgentRunner) schedule(
	ctx context.Context,
) {
	recoveryTicker := time.NewTicker(30 * time.Second)
	defer recoveryTicker.Stop()

	r.recoverAgentJobs(ctx)
	r.wakeScheduler()
	for {
		recover := false
		select {
		case <-ctx.Done():
			return
		case <-r.schedulerWake:
		case <-recoveryTicker.C:
			recover = true
		}
		if recover {
			r.recoverAgentJobs(ctx)
		}
		r.dispatch(ctx)
	}
}

func (r *WorldAgentRunner) recoverAgentJobs(
	ctx context.Context,
) {
	r.recoverGeneratingWorldJobs(ctx)
	r.recoverPendingCharacterJobs(ctx)
	r.recoverStaleRunningJobs(ctx)
}

func (r *WorldAgentRunner) dispatch(
	ctx context.Context,
) {
	status := enum.AgentJobStatusQueued
	now := time.Now()
	jobs, err := r.agentJobRepo.List(ctx, &repo.AgentJobQuery{
		Status:          new(status),
		AvailableBefore: new(now),
	})
	if err != nil {
		r.log.ErrorContext(ctx, "load queued agent jobs failed", constant.LogFieldErr, err)
		return
	}
	if len(jobs) > 0 {
		r.log.InfoContext(ctx, "dispatch queued agent jobs", "count", len(jobs))
	}
	jobs = r.orderDispatchJobs(jobs, now)
	for _, job := range jobs {
		world, ok := r.acquire(ctx, job)
		if !ok {
			continue
		}
		go r.execute(ctx, job, world)
	}
}

func (r *WorldAgentRunner) recoverGeneratingWorldJobs(
	ctx context.Context,
) {
	status := enum.WorldStatusGenerating
	worlds, err := r.worldRepo.List(ctx, &repo.WorldQuery{
		Status: new(status),
	})
	if err != nil {
		r.log.ErrorContext(ctx, "load generating worlds failed", constant.LogFieldErr, err)
		return
	}
	for _, world := range worlds {
		eventType := enum.EventTypeWorldGenerationRequested
		events, err := r.eventRepo.List(ctx, &repo.EventQuery{
			WorldID: new(world.ID),
			Type:    new(eventType),
		})
		if err != nil {
			r.log.ErrorContext(ctx, "load world generation events failed", "world_id", world.ID, constant.LogFieldErr, err)
			continue
		}
		if len(events) == 0 {
			continue
		}
		source := events[0]
		jobType := enum.AgentJobTypeWorldGenerate
		count, err := r.agentJobRepo.Count(ctx, &repo.AgentJobQuery{
			SourceEventID: new(source.ID),
			Type:          new(jobType),
		})
		if err != nil {
			r.log.ErrorContext(ctx, "count world generation jobs failed", "world_id", world.ID, constant.LogFieldEventID, source.ID, constant.LogFieldErr, err)
			continue
		}
		if count > 0 {
			continue
		}
		state, err := r.worldStateRepo.Get(ctx, &repo.WorldStateQuery{
			WorldID: new(world.ID),
		})
		if err != nil {
			r.log.ErrorContext(ctx, "load generating world state failed", "world_id", world.ID, constant.LogFieldErr, err)
			continue
		}
		_, err = r.agentJobRepo.Save(ctx, &model.AgentJob{
			WorldID:       world.ID,
			SourceEventID: source.ID,
			Type:          enum.AgentJobTypeWorldGenerate,
			Priority:      enum.AgentJobPriorityHigh,
			LaneKey:       "world",
			Status:        enum.AgentJobStatusQueued,
			WorldVersion:  state.Version,
			AvailableAt:   time.Now(),
		})
		if err != nil {
			r.log.ErrorContext(ctx, "recover world generation job failed", "world_id", world.ID, constant.LogFieldEventID, source.ID, constant.LogFieldErr, err)
			continue
		}
		r.log.WarnContext(ctx, "recovered missing world generation job", "world_id", world.ID, constant.LogFieldEventID, source.ID)
	}
}

func (r *WorldAgentRunner) recoverPendingCharacterJobs(
	ctx context.Context,
) {
	status := enum.WorldStatusActive
	worlds, err := r.worldRepo.List(ctx, &repo.WorldQuery{
		Status: new(status),
	})
	if err != nil {
		r.log.ErrorContext(ctx, "load active worlds failed", constant.LogFieldErr, err)
		return
	}
	for _, world := range worlds {
		r.recoverWorldCharacterJobs(ctx, world)
	}
}

func (r *WorldAgentRunner) recoverWorldCharacterJobs(
	ctx context.Context,
	world *model.World,
) {
	if world == nil {
		return
	}
	members, err := r.worldMemberRepo.List(ctx, &repo.WorldMemberQuery{
		WorldID: new(world.ID),
	})
	if err != nil {
		r.log.ErrorContext(ctx, "load world members failed", "world_id", world.ID, constant.LogFieldErr, err)
		return
	}
	for _, member := range members {
		if member == nil || member.CharacterReady {
			continue
		}
		if err = r.recoverMemberCharacterJob(ctx, world, member); err != nil {
			r.log.ErrorContext(ctx, "recover character job failed", "world_id", world.ID, "player_id", member.PlayerID, constant.LogFieldErr, err)
		}
	}
}

func (r *WorldAgentRunner) recoverMemberCharacterJob(
	ctx context.Context,
	world *model.World,
	member *model.WorldMember,
) error {
	eventType := enum.EventTypePlayerJoined
	events, err := r.eventRepo.List(ctx, &repo.EventQuery{
		WorldID:       new(world.ID),
		Type:          new(eventType),
		ActorPlayerID: new(member.PlayerID),
	})
	if err != nil {
		return err
	}
	if len(events) == 0 {
		return fmt.Errorf("player joined event not found")
	}
	source := events[len(events)-1]
	jobType := enum.AgentJobTypePlayerCharacterGenerate
	count, err := r.agentJobRepo.Count(ctx, &repo.AgentJobQuery{
		SourceEventID: new(source.ID),
		Type:          new(jobType),
	})
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	state, err := r.worldStateRepo.Get(ctx, &repo.WorldStateQuery{
		WorldID: new(world.ID),
	})
	if err != nil {
		return err
	}
	_, err = r.agentJobRepo.Save(ctx, &model.AgentJob{
		WorldID:       world.ID,
		SourceEventID: source.ID,
		Type:          enum.AgentJobTypePlayerCharacterGenerate,
		Priority:      enum.AgentJobPriorityHigh,
		LaneKey:       fmt.Sprintf("player:%d", member.PlayerID),
		Status:        enum.AgentJobStatusQueued,
		WorldVersion:  state.Version,
		AvailableAt:   time.Now(),
	})
	if err != nil {
		return err
	}
	r.log.WarnContext(ctx, "recovered missing character generation job", "world_id", world.ID, "player_id", member.PlayerID, constant.LogFieldEventID, source.ID)
	return nil
}

func (r *WorldAgentRunner) recoverStaleRunningJobs(
	ctx context.Context,
) {
	status := enum.AgentJobStatusRunning
	cutoff := time.Now().Add(-r.staleJobDuration())
	jobs, err := r.agentJobRepo.List(ctx, &repo.AgentJobQuery{
		Status:        new(status),
		StartedBefore: new(cutoff),
	})
	if err != nil {
		r.log.ErrorContext(ctx, "load stale running agent jobs failed", constant.LogFieldErr, err)
		return
	}
	for _, job := range jobs {
		_, err = r.agentJobRepo.Retry(ctx, &repo.AgentJobRetryReq{
			JobID:        job.ID,
			AttemptCount: job.AttemptCount,
			AvailableAt:  time.Now(),
			ErrorSummary: "recovered stale running job",
		})
		if err != nil {
			r.log.ErrorContext(ctx, "recover stale running agent job failed", constant.LogFieldTaskID, job.ID, constant.LogFieldErr, err)
			continue
		}
		r.log.WarnContext(ctx, "recovered stale running agent job", constant.LogFieldTaskID, job.ID, "world_id", job.WorldID, "type", job.Type)
	}
}

func fairAgentJobs(
	jobs []*model.AgentJob,
	cursor *int,
) []*model.AgentJob {
	ordered := make([]*model.AgentJob, 0, len(jobs))
	used := make(map[int64]bool, len(jobs))
	priorities := []enum.AgentJobPriority{
		enum.AgentJobPriorityHigh,
		enum.AgentJobPriorityLow,
	}
	baseCursor := 0
	if cursor != nil {
		baseCursor = *cursor
		*cursor = *cursor + 1
	}
	for _, priority := range priorities {
		worldOrder := make([]int64, 0)
		buckets := make(map[int64][]*model.AgentJob)
		for _, job := range jobs {
			if job == nil || job.Priority != priority {
				continue
			}
			if _, ok := buckets[job.WorldID]; !ok {
				worldOrder = append(worldOrder, job.WorldID)
			}
			buckets[job.WorldID] = append(buckets[job.WorldID], job)
		}
		if len(worldOrder) == 0 {
			continue
		}
		start := baseCursor % len(worldOrder)
		remaining := true
		for remaining {
			remaining = false
			for offset := 0; offset < len(worldOrder); offset++ {
				worldID := worldOrder[(start+offset)%len(worldOrder)]
				bucket := buckets[worldID]
				if len(bucket) == 0 {
					continue
				}
				remaining = true
				ordered = append(ordered, bucket[0])
				used[bucket[0].ID] = true
				buckets[worldID] = bucket[1:]
			}
		}
	}
	for _, job := range jobs {
		if job != nil && !used[job.ID] {
			ordered = append(ordered, job)
		}
	}
	return ordered
}

func (r *WorldAgentRunner) orderDispatchJobs(
	jobs []*model.AgentJob,
	now time.Time,
) []*model.AgentJob {
	overdueTicks, remaining := r.splitOverdueTickJobs(jobs, now)
	ordered := make([]*model.AgentJob, 0, len(jobs))
	ordered = append(ordered, fairAgentJobs(overdueTicks, &r.scheduleCursor)...)
	ordered = append(ordered, fairAgentJobs(remaining, &r.scheduleCursor)...)
	return ordered
}

func (r *WorldAgentRunner) promoteOverdueTickJobs(
	jobs []*model.AgentJob,
	now time.Time,
) []*model.AgentJob {
	promoted, remaining := r.splitOverdueTickJobs(jobs, now)
	if len(promoted) == 0 {
		return jobs
	}
	return append(promoted, remaining...)
}

func (r *WorldAgentRunner) splitOverdueTickJobs(
	jobs []*model.AgentJob,
	now time.Time,
) ([]*model.AgentJob, []*model.AgentJob) {
	threshold := r.tickInterval() / 2
	if threshold < 30*time.Second {
		threshold = 30 * time.Second
	}

	promoted := make([]*model.AgentJob, 0)
	remaining := make([]*model.AgentJob, 0, len(jobs))
	for _, job := range jobs {
		if job != nil && job.Type == enum.AgentJobTypeWorldTick && now.Sub(job.AvailableAt) >= threshold {
			promoted = append(promoted, job)
			continue
		}
		remaining = append(remaining, job)
	}
	return promoted, remaining
}

func (r *WorldAgentRunner) acquire(
	ctx context.Context,
	job *model.AgentJob,
) (*model.World, bool) {
	if job == nil {
		return nil, false
	}
	laneKey := fmt.Sprintf("%d:%s", job.WorldID, job.LaneKey)
	if job.Type == enum.AgentJobTypeMemoryEmbed {
		return r.acquireMemoryJob(ctx, job, laneKey)
	}

	r.mu.Lock()
	blocked := r.active >= r.globalLimit() ||
		r.activeWorld[job.WorldID] >= r.worldLimit() ||
		r.lanes[laneKey]
	r.mu.Unlock()
	if blocked {
		return nil, false
	}

	world, err := r.worldRepo.Get(ctx, &repo.WorldQuery{
		ID: new(job.WorldID),
	})
	if err != nil {
		r.log.ErrorContext(
			ctx,
			"load agent job world failed",
			constant.LogFieldTaskID,
			job.ID,
			"world_id",
			job.WorldID,
			constant.LogFieldErr,
			err,
		)
		return nil, false
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	if until := r.breakerUntil[world.AgentConfigID]; until.After(time.Now()) {
		return nil, false
	}
	if r.active >= r.globalLimit() ||
		r.activeWorld[job.WorldID] >= r.worldLimit() ||
		r.activeConfig[world.AgentConfigID] >= r.modelLimit() ||
		r.lanes[laneKey] {
		return nil, false
	}
	r.active++
	r.activeWorld[job.WorldID]++
	r.activeConfig[world.AgentConfigID]++
	r.lanes[laneKey] = true
	return world, true
}

func (r *WorldAgentRunner) acquireBackgroundJob(
	ctx context.Context,
	job *model.AgentJob,
	laneKey string,
) (*model.World, bool) {
	r.mu.Lock()
	blocked := r.lanes[laneKey]
	r.mu.Unlock()
	if blocked {
		return nil, false
	}

	world, err := r.worldRepo.Get(ctx, &repo.WorldQuery{
		ID: new(job.WorldID),
	})
	if err != nil {
		r.log.ErrorContext(
			ctx,
			"load background agent job world failed",
			constant.LogFieldTaskID,
			job.ID,
			"world_id",
			job.WorldID,
			constant.LogFieldErr,
			err,
		)
		return nil, false
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	if r.lanes[laneKey] {
		return nil, false
	}
	r.lanes[laneKey] = true
	return world, true
}

func (r *WorldAgentRunner) acquireMemoryJob(
	ctx context.Context,
	job *model.AgentJob,
	laneKey string,
) (*model.World, bool) {
	r.mu.Lock()
	blocked := r.activeMemory >= r.memoryLimit() || r.lanes[laneKey]
	r.mu.Unlock()
	if blocked {
		return nil, false
	}

	world, err := r.worldRepo.Get(ctx, &repo.WorldQuery{
		ID: new(job.WorldID),
	})
	if err != nil {
		r.log.ErrorContext(
			ctx,
			"load memory embedding job world failed",
			constant.LogFieldTaskID,
			job.ID,
			"world_id",
			job.WorldID,
			constant.LogFieldErr,
			err,
		)
		return nil, false
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	if r.activeMemory >= r.memoryLimit() || r.lanes[laneKey] {
		return nil, false
	}
	r.activeMemory++
	r.lanes[laneKey] = true
	return world, true
}

func (r *WorldAgentRunner) release(
	world *model.World,
	job *model.AgentJob,
) {
	r.mu.Lock()
	switch job.Type {
	case enum.AgentJobTypeMemoryEmbed:
		r.activeMemory--
	default:
		r.active--
		r.activeWorld[job.WorldID]--
		r.activeConfig[world.AgentConfigID]--
	}
	delete(r.lanes, fmt.Sprintf("%d:%s", job.WorldID, job.LaneKey))
	r.mu.Unlock()
	r.wakeScheduler()
}

func (r *WorldAgentRunner) execute(
	ctx context.Context,
	job *model.AgentJob,
	world *model.World,
) {
	defer r.release(world, job)

	running, err := r.agentJobRepo.MarkRunning(ctx, job.ID, time.Now())
	if err != nil {
		r.log.ErrorContext(
			ctx,
			"mark agent job running failed",
			constant.LogFieldTaskID,
			job.ID,
			constant.LogFieldErr,
			err,
		)
		return
	}
	result := &agentResult{
		job:   running,
		world: world,
		done:  make(chan struct{}),
	}
	r.log.InfoContext(ctx, "execute agent job", constant.LogFieldTaskID, running.ID, "world_id", running.WorldID, "type", running.Type, "attempt", running.AttemptCount)
	result.source, result.err = r.eventRepo.Get(ctx, &repo.EventQuery{
		ID: new(job.SourceEventID),
	})
	if result.err == nil {
		result.state, result.err = r.worldStateRepo.Get(ctx, &repo.WorldStateQuery{
			WorldID: new(job.WorldID),
		})
	}
	if result.err == nil {
		result.config, result.err = r.agentConfigRepo.Get(ctx, &repo.AgentConfigQuery{
			ID: new(world.AgentConfigID),
		})
	}
	if result.err == nil {
		result.err = r.callAgent(ctx, result)
	}
	if result.err != nil {
		r.log.WarnContext(ctx, "agent job call failed", constant.LogFieldTaskID, running.ID, "world_id", running.WorldID, "type", running.Type, constant.LogFieldErr, result.err)
	} else {
		r.log.InfoContext(ctx, "agent job call succeeded", constant.LogFieldTaskID, running.ID, "world_id", running.WorldID, "type", running.Type)
	}
	loop := r.ensureLoop(ctx, job.WorldID)
	select {
	case loop.results <- result:
	case <-ctx.Done():
		return
	}
	select {
	case <-result.done:
	case <-ctx.Done():
	}
}
