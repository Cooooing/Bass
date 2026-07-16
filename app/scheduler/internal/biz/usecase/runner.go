package usecase

import (
	"common/pkg/constant"
	"context"
	"log/slog"
	"scheduler/internal/biz/model"
	"scheduler/internal/biz/repo"
	"scheduler/internal/config"
	schedulerenum "scheduler/internal/enum"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type SchedulerRunner struct {
	logger       *slog.Logger
	conf         *config.Bootstrap
	taskRepo     repo.TaskRepo
	taskUsecase  *TaskUsecase
	taskEventBus repo.TaskEventBus
	cron         *cron.Cron
	entryIDs     map[int64]cron.EntryID
	mu           sync.Mutex
	cancel       context.CancelFunc
}

func NewSchedulerRunner(logger *slog.Logger, conf *config.Bootstrap, taskRepo repo.TaskRepo, taskUsecase *TaskUsecase, taskEventBus repo.TaskEventBus) *SchedulerRunner {
	return &SchedulerRunner{
		logger:       logger,
		conf:         conf,
		taskRepo:     taskRepo,
		taskUsecase:  taskUsecase,
		taskEventBus: taskEventBus,
		cron:         cron.New(cron.WithParser(taskUsecase.cronParser)),
		entryIDs:     map[int64]cron.EntryID{},
	}
}

func (r *SchedulerRunner) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	r.cancel = cancel
	tasksResp, err := r.taskRepo.List(runCtx, &repo.TaskGetReq{Enabled: new(true)})
	if err != nil {
		return err
	}
	for _, item := range tasksResp.Rows {
		r.registerTask(runCtx, item)
	}
	changedResp, err := r.taskEventBus.SubscribeTaskChanged(runCtx, &repo.SubscribeTaskChangedReq{})
	if err != nil {
		return err
	}
	canceledResp, err := r.taskEventBus.SubscribeExecutionCanceled(runCtx, &repo.SubscribeExecutionCanceledReq{})
	if err != nil {
		return err
	}
	changedCh := changedResp.Messages
	canceledCh := canceledResp.Messages
	r.cron.Start()
	go func() {
		for {
			select {
			case <-runCtx.Done():
				return
			case msg, ok := <-changedCh:
				if !ok {
					return
				}
				taskResp, err := r.taskRepo.Get(runCtx, &repo.TaskGetReq{ID: &msg.TaskID})
				if err != nil {
					r.unregisterTask(runCtx, msg.TaskID)
					continue
				}
				r.registerTask(runCtx, taskResp.Row)
			case msg, ok := <-canceledCh:
				if !ok {
					return
				}
				r.taskUsecase.CancelExecutionLocally(runCtx, &TaskCancelExecutionLocallyReq{ID: msg.ExecutionRecordID})
			}
		}
	}()
	return nil
}

func (r *SchedulerRunner) Stop(ctx context.Context) error {
	if r.cancel != nil {
		r.cancel()
	}
	if r.cron != nil {
		stopCtx := r.cron.Stop()
		<-stopCtx.Done()
	}
	stopRunningCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	return r.taskUsecase.StopRunning(stopRunningCtx, &TaskStopRunningReq{})
}

func (r *SchedulerRunner) registerTask(ctx context.Context, task *model.Task) {
	schedule, err := r.taskUsecase.cronParser.Parse(task.CronSpec)
	if err != nil {
		r.logger.ErrorContext(ctx, "parse scheduler task cron failed", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name, constant.LogFieldErr, err)
		return
	}

	r.unregisterTask(ctx, task.ID)
	if !task.Enabled {
		r.logger.DebugContext(ctx, "scheduler task disabled", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name, "cron_spec", task.CronSpec)
		return
	}

	entryID := r.cron.Schedule(schedule, cron.FuncJob(func() {
		_, _ = r.taskUsecase.ScheduleExecution(ctx, &TaskScheduleExecutionReq{Task: task, ScheduledAt: time.Now().Truncate(time.Second), TriggerType: schedulerenum.TaskTriggerTypeSchedule})
	}))

	r.mu.Lock()
	r.entryIDs[task.ID] = entryID
	r.mu.Unlock()
	r.logger.DebugContext(ctx, "scheduler task registered", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name, "entry_id", int(entryID), "cron_spec", task.CronSpec)
}

func (r *SchedulerRunner) unregisterTask(ctx context.Context, taskID int64) {
	r.mu.Lock()
	entryID, ok := r.entryIDs[taskID]
	if ok {
		r.cron.Remove(entryID)
		delete(r.entryIDs, taskID)
		r.logger.DebugContext(ctx, "scheduler task unregistered", constant.LogFieldTaskID, taskID, "entry_id", int(entryID))
	}
	r.mu.Unlock()
}
