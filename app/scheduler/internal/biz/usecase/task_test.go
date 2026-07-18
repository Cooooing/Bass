package usecase

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"os"
	"scheduler/internal/biz/model"
	"scheduler/internal/biz/repo"
	taskimpl "scheduler/internal/biz/usecase/task"
	"scheduler/internal/config"
	schedulerenum "scheduler/internal/enum"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	utilent "common/pkg/util/ent"

	"google.golang.org/protobuf/types/known/durationpb"
)

type fakeExecutionRepo struct {
	mu                     sync.Mutex
	record                 *model.TaskExecutionRecord
	records                map[int64]*model.TaskExecutionRecord
	recordsByPeriod        map[string]*model.TaskExecutionRecord
	nextID                 int64
	getCalled              int
	createCalled           int
	created                bool
	scheduleConflict       bool
	createErr              error
	hasUnexpiredRunning    bool
	hasUnexpiredRunningErr error
	markFinishedRows       []*model.TaskExecutionRecord
}

func (r *fakeExecutionRepo) Get(_ context.Context, req *repo.TaskExecutionRecordGetReq) (*model.TaskExecutionRecord, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.getCalled++
	if req != nil && req.TaskID != nil && req.ScheduledAt != nil && r.recordsByPeriod != nil {
		if row, ok := r.recordsByPeriod[executionPeriodKey(*req.TaskID, *req.ScheduledAt)]; ok {
			return row, nil
		}
		return nil, errors.New("execution record not found")
	}
	if req != nil && req.ID != nil && r.records != nil {
		if row, ok := r.records[*req.ID]; ok {
			return row, nil
		}
		return nil, errors.New("execution record not found")
	}
	return r.record, nil
}

func (r *fakeExecutionRepo) List(_ context.Context, req *repo.TaskExecutionRecordGetReq) ([]*model.TaskExecutionRecord, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if req != nil && len(req.IDs) > 0 && r.records != nil {
		rows := make([]*model.TaskExecutionRecord, 0, len(req.IDs))
		for _, id := range req.IDs {
			if row := r.records[id]; row != nil {
				rows = append(rows, row)
			}
		}
		return rows, nil
	}
	rows := make([]*model.TaskExecutionRecord, 0, len(r.records))
	for _, row := range r.records {
		rows = append(rows, row)
	}
	return rows, nil
}

func (r *fakeExecutionRepo) Map(context.Context, *repo.TaskExecutionRecordGetReq) (map[int64]*model.TaskExecutionRecord, error) {
	return nil, nil
}

func (r *fakeExecutionRepo) Count(context.Context, *repo.TaskExecutionRecordGetReq) (int, error) {
	return 0, nil
}

func (r *fakeExecutionRepo) Page(context.Context, *repo.TaskExecutionRecordPageReq) (*repo.TaskExecutionRecordPageResp, error) {
	return &repo.TaskExecutionRecordPageResp{}, nil
}

func (r *fakeExecutionRepo) ExistsPeriod(_ context.Context, req *repo.TaskExecutionRecordExistsPeriodReq) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.recordsByPeriod == nil {
		return false, nil
	}
	_, ok := r.recordsByPeriod[executionPeriodKey(req.TaskID, req.ScheduledAt)]
	return ok, nil
}

func (r *fakeExecutionRepo) Create(_ context.Context, req *repo.TaskExecutionRecordCreateReq) (*repo.TaskExecutionRecordCreateResp, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.createCalled++
	if r.createErr != nil {
		return nil, r.createErr
	}
	if r.scheduleConflict {
		return &repo.TaskExecutionRecordCreateResp{Row: r.record, Conflict: true}, nil
	}
	if !r.created {
		return &repo.TaskExecutionRecordCreateResp{}, nil
	}
	if r.records == nil {
		r.records = map[int64]*model.TaskExecutionRecord{}
	}
	if r.recordsByPeriod == nil {
		r.recordsByPeriod = map[string]*model.TaskExecutionRecord{}
	}
	record := req.Record
	key := executionPeriodKey(record.TaskID, record.ScheduledAt)
	if current, ok := r.recordsByPeriod[key]; ok {
		return &repo.TaskExecutionRecordCreateResp{Row: current, Conflict: true}, nil
	}
	row := *record
	r.nextID++
	row.ID = r.nextID
	row.Status = req.Status
	now := time.Now()
	row.UpdatedAt = &now
	r.record = &row
	r.records[row.ID] = &row
	r.recordsByPeriod[key] = &row
	return &repo.TaskExecutionRecordCreateResp{Row: &row, Created: true}, nil
}

func (r *fakeExecutionRepo) HasUnexpiredRunning(context.Context, *repo.TaskExecutionRecordHasUnexpiredRunningReq) (bool, error) {
	return r.hasUnexpiredRunning, r.hasUnexpiredRunningErr
}

func (r *fakeExecutionRepo) MarkUnknown(_ context.Context, req *repo.TaskExecutionRecordMarkUnknownReq) ([]*model.TaskExecutionRecord, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	rows := make([]*model.TaskExecutionRecord, 0, len(req.IDs))
	for _, id := range req.IDs {
		row := r.records[id]
		if row == nil || row.Status != schedulerenum.TaskExecutionStatusRunning {
			continue
		}
		durationMS := int64(0)
		if row.StartedAt != nil {
			durationMS = req.FinishedAt.Sub(*row.StartedAt).Milliseconds()
		}
		row.Status = schedulerenum.TaskExecutionStatusUnknown
		row.FinishedAt = new(req.FinishedAt)
		row.DurationMS = new(durationMS)
		row.LastError = req.LastError
		rows = append(rows, row)
	}
	return rows, nil
}

func (r *fakeExecutionRepo) MarkFinished(_ context.Context, req *repo.TaskExecutionRecordMarkFinishedReq) (*repo.TaskExecutionRecordMarkFinishedResp, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.records != nil {
		current := r.records[req.ID]
		if current == nil || current.Status != schedulerenum.TaskExecutionStatusRunning {
			return &repo.TaskExecutionRecordMarkFinishedResp{Row: current}, nil
		}
		current.Status = req.Status
		current.FinishedAt = new(req.FinishedAt)
		current.DurationMS = new(req.DurationMS)
		current.LastError = req.LastError
		current.UpdatedAt = new(req.FinishedAt)
		r.markFinishedRows = append(r.markFinishedRows, current)
		return &repo.TaskExecutionRecordMarkFinishedResp{Row: current, Updated: true}, nil
	}
	row := &model.TaskExecutionRecord{
		ID:         req.ID,
		Status:     req.Status,
		FinishedAt: new(req.FinishedAt),
		DurationMS: new(req.DurationMS),
		LastError:  req.LastError,
	}
	r.markFinishedRows = append(r.markFinishedRows, row)
	return &repo.TaskExecutionRecordMarkFinishedResp{Row: row, Updated: true}, nil
}

type fakeTaskRepo struct {
	task    *model.Task
	listReq *repo.TaskGetReq
}

func (r *fakeTaskRepo) Get(_ context.Context, req *repo.TaskGetReq) (*model.Task, error) {
	if r.task == nil {
		return nil, errors.New("task not found")
	}
	if req != nil && req.ID != nil && r.task.ID != *req.ID {
		return nil, errors.New("task not found")
	}
	return r.task, nil
}

func (r *fakeTaskRepo) List(_ context.Context, req *repo.TaskGetReq) ([]*model.Task, error) {
	r.listReq = req
	if r.task == nil {
		return nil, nil
	}
	return []*model.Task{r.task}, nil
}

func (r *fakeTaskRepo) Map(context.Context, *repo.TaskGetReq) (map[int64]*model.Task, error) {
	return nil, nil
}

func (r *fakeTaskRepo) Count(context.Context, *repo.TaskGetReq) (int, error) {
	return 0, nil
}

func (r *fakeTaskRepo) Page(context.Context, *repo.TaskPageReq) (*repo.TaskPageResp, error) {
	return &repo.TaskPageResp{}, nil
}

func (r *fakeTaskRepo) Upsert(_ context.Context, row *model.Task) (*model.Task, error) {
	saved := *row
	if saved.ID == 0 {
		saved.ID = 1
	}
	if saved.Version == 0 {
		saved.Version = 1
	} else {
		saved.Version++
	}
	r.task = &saved
	return &saved, nil
}

func (r *fakeTaskRepo) Lock(context.Context, int64) error {
	return nil
}

type fakeTaskVersionRepo struct {
	mu      sync.Mutex
	created []*model.TaskVersion
	err     error
}

func (r *fakeTaskVersionRepo) Get(context.Context, *repo.TaskVersionGetReq) (*model.TaskVersion, error) {
	return nil, nil
}

func (r *fakeTaskVersionRepo) List(_ context.Context, req *repo.TaskVersionGetReq) ([]*model.TaskVersion, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	rows := make([]*model.TaskVersion, 0, len(r.created))
	for _, row := range r.created {
		if req != nil && len(req.TaskIDs) > 0 {
			ok := false
			for _, id := range req.TaskIDs {
				if row.TaskID == id {
					ok = true
					break
				}
			}
			if !ok {
				continue
			}
		}
		if req != nil && len(req.Versions) > 0 {
			ok := false
			for _, version := range req.Versions {
				if row.Version == version {
					ok = true
					break
				}
			}
			if !ok {
				continue
			}
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func (r *fakeTaskVersionRepo) Map(context.Context, *repo.TaskVersionGetReq) (map[int64]*model.TaskVersion, error) {
	return nil, nil
}

func (r *fakeTaskVersionRepo) Count(context.Context, *repo.TaskVersionGetReq) (int, error) {
	return 0, nil
}

func (r *fakeTaskVersionRepo) Page(context.Context, *repo.TaskVersionPageReq) (*repo.TaskVersionPageResp, error) {
	return &repo.TaskVersionPageResp{}, nil
}

func (r *fakeTaskVersionRepo) Create(_ context.Context, task *model.Task) (*model.TaskVersion, error) {
	if r.err != nil {
		return nil, r.err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	row := task
	version := &model.TaskVersion{
		ID:             int64(len(r.created) + 1),
		TaskID:         row.ID,
		Version:        row.Version,
		Name:           row.Name,
		Title:          row.Title,
		Description:    row.Description,
		Enabled:        row.Enabled,
		CronSpec:       row.CronSpec,
		Payload:        row.Payload,
		TimeoutSeconds: row.TimeoutSeconds,
		AllowOverlap:   row.AllowOverlap,
		AlertEnabled:   row.AlertEnabled,
	}
	r.created = append(r.created, version)
	return version, nil
}

type fakeTaskLockRepo struct {
	mu             sync.Mutex
	err            error
	schedules      map[string]string
	runningByTask  map[int64]string
	runningByID    map[int64]string
	acquireCalled  int
	registerCalled int
	refreshCalled  int
	releaseCalled  int
	releasedTokens []string
}

func (r *fakeTaskLockRepo) TryAcquireSchedule(_ context.Context, req *repo.TaskScheduleAcquireReq) (*repo.TaskScheduleAcquireResp, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.acquireCalled++
	if r.err != nil {
		return nil, r.err
	}
	if r.schedules == nil {
		r.schedules = map[string]string{}
	}
	key := executionPeriodKey(req.TaskID, req.ScheduledAt)
	if _, ok := r.schedules[key]; ok {
		return &repo.TaskScheduleAcquireResp{Decision: schedulerenum.TaskScheduleDecisionSkip}, nil
	}
	r.schedules[key] = "schedule-token-" + strconv.Itoa(r.acquireCalled)
	token := "running-token-" + strconv.Itoa(r.acquireCalled)
	if req.AllowOverlap {
		return &repo.TaskScheduleAcquireResp{Decision: schedulerenum.TaskScheduleDecisionRun, RunningToken: token}, nil
	}
	if r.runningByTask == nil {
		r.runningByTask = map[int64]string{}
	}
	if r.runningByTask[req.TaskID] != "" {
		return &repo.TaskScheduleAcquireResp{Decision: schedulerenum.TaskScheduleDecisionOverlap}, nil
	}
	r.runningByTask[req.TaskID] = token
	return &repo.TaskScheduleAcquireResp{
		Decision:     schedulerenum.TaskScheduleDecisionRun,
		RunningToken: token,
	}, nil
}

func (r *fakeTaskLockRepo) RegisterRunning(_ context.Context, req *repo.TaskRunningLockReq) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.registerCalled++
	if r.err != nil {
		return false, r.err
	}
	if r.runningByID == nil {
		r.runningByID = map[int64]string{}
	}
	if req.Exclusive && (r.runningByTask == nil || r.runningByTask[req.TaskID] != req.RunningToken) {
		return false, nil
	}
	r.runningByID[req.ExecutionRecordID] = req.RunningToken
	return true, nil
}

func (r *fakeTaskLockRepo) RefreshRunning(_ context.Context, req *repo.TaskRunningLockReq) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.refreshCalled++
	if r.err != nil {
		return false, r.err
	}
	if r.runningByID == nil || r.runningByID[req.ExecutionRecordID] != req.RunningToken {
		return false, nil
	}
	if req.Exclusive {
		return r.runningByTask != nil && r.runningByTask[req.TaskID] == req.RunningToken, nil
	}
	return true, nil
}

func (r *fakeTaskLockRepo) ReleaseRunning(_ context.Context, req *repo.TaskRunningLockReq) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.releaseCalled++
	r.releasedTokens = append(r.releasedTokens, req.RunningToken)
	if r.err != nil {
		return r.err
	}
	if req.RunningToken == "" {
		return nil
	}
	if r.runningByID != nil && r.runningByID[req.ExecutionRecordID] == req.RunningToken {
		delete(r.runningByID, req.ExecutionRecordID)
	}
	if req.Exclusive && r.runningByTask != nil && r.runningByTask[req.TaskID] == req.RunningToken {
		delete(r.runningByTask, req.TaskID)
	}
	return nil
}

func (r *fakeTaskLockRepo) MapRunning(_ context.Context, req *repo.TaskRunningMapReq) (map[int64]bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		return nil, r.err
	}
	result := make(map[int64]bool, len(req.ExecutionRecordIDs))
	for _, id := range req.ExecutionRecordIDs {
		result[id] = r.runningByID != nil && r.runningByID[id] != ""
	}
	return result, nil
}

type fakeTaskEventBus struct {
	mu             sync.Mutex
	cancelErr      error
	cancelMessages []*repo.TaskExecutionCanceledMessage
}

func (fakeTaskEventBus) PublishTaskChanged(context.Context, *repo.TaskChangedMessage) error {
	return nil
}

func (b *fakeTaskEventBus) PublishExecutionCanceled(_ context.Context, msg *repo.TaskExecutionCanceledMessage) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.cancelErr != nil {
		return b.cancelErr
	}
	b.cancelMessages = append(b.cancelMessages, msg)
	return nil
}

func (fakeTaskEventBus) SubscribeTaskChanged(context.Context) (<-chan repo.TaskChangedMessage, error) {
	return make(chan repo.TaskChangedMessage), nil
}

func (fakeTaskEventBus) SubscribeExecutionCanceled(context.Context) (<-chan repo.TaskExecutionCanceledMessage, error) {
	return make(chan repo.TaskExecutionCanceledMessage), nil
}

type fakeAlert struct {
	mu      sync.Mutex
	reason  string
	records int
}

func (a *fakeAlert) Alert(_ context.Context, req *repo.TaskAlertReq) error {
	reason := req.Reason
	a.mu.Lock()
	defer a.mu.Unlock()
	a.reason = reason
	a.records++
	return nil
}

type fakeTask struct {
	executed    chan string
	err         error
	block       chan struct{}
	nilOnCancel bool
	executions  atomic.Int64
}

func (t *fakeTask) Name() string { return "noop" }

func (t *fakeTask) Title() string { return "空任务" }

func (t *fakeTask) Description() string { return "用于测试的空任务" }

func (t *fakeTask) Execute(ctx context.Context, payload string) error {
	t.executions.Add(1)
	if t.executed != nil {
		t.executed <- payload
	}
	if t.block != nil {
		select {
		case <-ctx.Done():
			if t.nilOnCancel {
				return nil
			}
			return ctx.Err()
		case <-t.block:
		}
	}
	return t.err
}

func TestUpsertCreatesTaskVersionSnapshot(t *testing.T) {
	taskRepo := &fakeTaskRepo{}
	taskVersionRepo := &fakeTaskVersionRepo{}
	usecase := NewTaskUsecase(
		slog.New(slog.NewTextHandler(io.Discard, nil)),
		testConf(),
		testTx,
		taskRepo,
		taskVersionRepo,
		&fakeExecutionRepo{},
		&fakeTaskLockRepo{},
		map[string]taskimpl.Task{"noop": &fakeTask{}},
		&fakeTaskEventBus{},
		&fakeAlert{},
	)

	upsertResp, err := usecase.Upsert(context.Background(), testTask(true))
	if err != nil {
		t.Fatalf("Upsert returned error: %v", err)
	}
	saved := upsertResp
	if saved == nil {
		t.Fatal("expected saved task")
	}
	if len(taskVersionRepo.created) != 1 {
		t.Fatalf("expected one task version snapshot, got %d", len(taskVersionRepo.created))
	}
	version := taskVersionRepo.created[0]
	if version.TaskID != saved.ID || version.Version != saved.Version || version.Name != saved.Name {
		t.Fatalf("unexpected task version snapshot: %#v saved=%#v", version, saved)
	}
}

func TestCancelExecutionPublishesIntentWithoutChangingRecordStatus(t *testing.T) {
	startedAt := time.Now().Add(-time.Second)
	record := &model.TaskExecutionRecord{
		ID:        7,
		TaskID:    testTask(true).ID,
		StartedAt: new(startedAt),
		Status:    schedulerenum.TaskExecutionStatusRunning,
	}
	executionRepo := &fakeExecutionRepo{records: map[int64]*model.TaskExecutionRecord{record.ID: record}}
	eventBus := &fakeTaskEventBus{}
	usecase := NewTaskUsecase(
		slog.New(slog.NewTextHandler(io.Discard, nil)),
		testConf(),
		testTx,
		&fakeTaskRepo{task: testTask(true)},
		&fakeTaskVersionRepo{},
		executionRepo,
		&fakeTaskLockRepo{},
		map[string]taskimpl.Task{"noop": &fakeTask{}},
		eventBus,
		&fakeAlert{},
	)

	cancelResp, err := usecase.CancelExecution(context.Background(), record.ID)
	if err != nil {
		t.Fatalf("CancelExecution returned error: %v", err)
	}
	row := cancelResp
	if row.Status != schedulerenum.TaskExecutionStatusRunning || record.Status != schedulerenum.TaskExecutionStatusRunning {
		t.Fatalf("cancel should not finish execution record, row=%s record=%s", row.Status, record.Status)
	}
	if len(eventBus.cancelMessages) != 1 || eventBus.cancelMessages[0].ExecutionRecordID != record.ID {
		t.Fatalf("expected one cancel message for execution %d, got %#v", record.ID, eventBus.cancelMessages)
	}
	if len(executionRepo.markFinishedRows) != 0 {
		t.Fatalf("cancel should not mark finished, got %d rows", len(executionRepo.markFinishedRows))
	}
}

func TestCanceledContextMarksExecutionCanceledWhenTaskReturnsNil(t *testing.T) {
	taskImpl := &fakeTask{block: make(chan struct{}), nilOnCancel: true}
	executionRepo := &fakeExecutionRepo{created: true}
	usecase := newTestTaskUsecase(t, executionRepo, taskImpl)

	startResp, err := usecase.StartExecution(context.Background(), &TaskStartExecutionReq{Task: testTask(true), ScheduledAt: time.Now(), TriggerType: schedulerenum.TaskTriggerTypeManual})
	if err != nil || !startResp.Created {
		t.Fatalf("StartExecution returned record=%#v created=%v err=%v", startResp.Record, startResp.Created, err)
	}
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) && taskImpl.executions.Load() == 0 {
		time.Sleep(time.Millisecond)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := usecase.StopRunning(ctx); err != nil {
		t.Fatalf("StopRunning returned error: %v", err)
	}
	if len(executionRepo.markFinishedRows) != 1 {
		t.Fatalf("expected one finished row, got %d", len(executionRepo.markFinishedRows))
	}
	if executionRepo.markFinishedRows[0].Status != schedulerenum.TaskExecutionStatusCanceled {
		t.Fatalf("expected canceled status, got %s", executionRepo.markFinishedRows[0].Status)
	}
}

func TestStartExecutionScheduleConflictReturnsResult(t *testing.T) {
	existing := &model.TaskExecutionRecord{ID: 1}
	executionRepo := &fakeExecutionRepo{record: existing, scheduleConflict: true}
	usecase := newTestTaskUsecase(t, executionRepo, &fakeTask{})

	startResp, err := usecase.StartExecution(context.Background(), &TaskStartExecutionReq{Task: testTask(true), ScheduledAt: time.Now(), TriggerType: schedulerenum.TaskTriggerTypeSchedule})
	if err != nil {
		t.Fatalf("StartExecution returned error: %v", err)
	}
	record := startResp.Record
	created := startResp.Created
	scheduleConflict := startResp.Conflict
	if record == nil || record.ID != existing.ID {
		t.Fatalf("expected existing record on schedule key conflict, got %#v", record)
	}
	if created {
		t.Fatal("expected created=false on schedule conflict")
	}
	if !scheduleConflict {
		t.Fatal("expected scheduleConflict=true")
	}
	if executionRepo.createCalled != 1 {
		t.Fatalf("expected one create attempt, got %d", executionRepo.createCalled)
	}
}

func TestScheduleExecutionOverlapSkippedWhenRunningLockExists(t *testing.T) {
	alert := &fakeAlert{}
	executionRepo := &fakeExecutionRepo{created: true}
	task := testTask(false)
	taskLockRepo := &fakeTaskLockRepo{runningByTask: map[int64]string{task.ID: "existing-token"}}
	usecase := NewTaskUsecase(
		slog.New(slog.NewTextHandler(io.Discard, nil)),
		testConf(),
		testTx,
		&fakeTaskRepo{},
		&fakeTaskVersionRepo{},
		executionRepo,
		taskLockRepo,
		map[string]taskimpl.Task{"noop": &fakeTask{executed: make(chan string, 1)}},
		&fakeTaskEventBus{},
		alert,
	)

	scheduleResp, err := usecase.ScheduleExecution(context.Background(), &TaskScheduleExecutionReq{Task: task, ScheduledAt: time.Now(), TriggerType: schedulerenum.TaskTriggerTypeManual})
	if err != nil {
		t.Fatalf("StartExecution returned error: %v", err)
	}
	record := scheduleResp
	if record == nil || record.Status != schedulerenum.TaskExecutionStatusOverlapSkipped {
		t.Fatalf("expected overlap_skipped record, got %#v", record)
	}
	if alert.records != 1 || alert.reason != "overlap" {
		t.Fatalf("expected overlap alert, got count=%d reason=%q", alert.records, alert.reason)
	}
}

func TestStartExecutionExecutesAndMarksSuccess(t *testing.T) {
	executed := make(chan string, 1)
	alert := &fakeAlert{}
	executionRepo := &fakeExecutionRepo{created: true}
	usecase := newTestTaskUsecase(t, executionRepo, &fakeTask{executed: executed}, alert)
	task := testTask(true)

	scheduleResp, err := usecase.ScheduleExecution(context.Background(), &TaskScheduleExecutionReq{Task: task, ScheduledAt: time.Now(), TriggerType: schedulerenum.TaskTriggerTypeManual})
	if err != nil {
		t.Fatalf("StartExecution returned error: %v", err)
	}
	record := scheduleResp
	if record == nil || record.ID == 0 {
		t.Fatalf("expected running record, got %#v", record)
	}
	select {
	case payload := <-executed:
		if payload != task.Payload {
			t.Fatalf("expected payload %q, got %q", task.Payload, payload)
		}
	case <-time.After(time.Second):
		t.Fatal("expected task execution")
	}
	deadline := time.After(time.Second)
	for {
		executionRepo.mu.Lock()
		rows := append([]*model.TaskExecutionRecord(nil), executionRepo.markFinishedRows...)
		executionRepo.mu.Unlock()
		if len(rows) > 0 {
			if rows[0].Status != schedulerenum.TaskExecutionStatusSuccess {
				t.Fatalf("expected success, got %s", rows[0].Status)
			}
			if alert.records != 0 {
				t.Fatalf("expected no alert on success, got count=%d reason=%q", alert.records, alert.reason)
			}
			return
		}
		select {
		case <-deadline:
			t.Fatal("expected execution to mark success")
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func TestStartExecutionFailedAlertsOnce(t *testing.T) {
	alert := &fakeAlert{}
	executionRepo := &fakeExecutionRepo{created: true}
	usecase := newTestTaskUsecase(t, executionRepo, &fakeTask{err: errors.New("execute failed")}, alert)
	task := testTask(true)

	scheduleResp, err := usecase.ScheduleExecution(context.Background(), &TaskScheduleExecutionReq{Task: task, ScheduledAt: time.Now(), TriggerType: schedulerenum.TaskTriggerTypeManual})
	if err != nil {
		t.Fatalf("ScheduleExecution returned error: %v", err)
	}
	record := scheduleResp
	if record == nil {
		t.Fatal("expected running record")
	}
	deadline := time.After(time.Second)
	for {
		alert.mu.Lock()
		records := alert.records
		reason := alert.reason
		alert.mu.Unlock()
		if records > 0 {
			if records != 1 || reason != schedulerenum.TaskExecutionStatusFailed.String() {
				t.Fatalf("expected one failed alert, got count=%d reason=%q", records, reason)
			}
			return
		}
		select {
		case <-deadline:
			t.Fatal("expected failed alert")
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func newTestTaskUsecase(t *testing.T, executionRepo *fakeExecutionRepo, task taskimpl.Task, alerts ...*fakeAlert) *TaskUsecase {
	t.Helper()
	alert := &fakeAlert{}
	if len(alerts) > 0 && alerts[0] != nil {
		alert = alerts[0]
	}
	return NewTaskUsecase(
		slog.New(slog.NewTextHandler(io.Discard, nil)),
		testConf(),
		testTx,
		&fakeTaskRepo{task: testTask(true)},
		&fakeTaskVersionRepo{},
		executionRepo,
		&fakeTaskLockRepo{},
		map[string]taskimpl.Task{"noop": task},
		&fakeTaskEventBus{},
		alert,
	)
}

func testTx(ctx context.Context, fn func(ctx context.Context) error, _ ...utilent.TxOption) error {
	return fn(ctx)
}

func testConf() *config.Bootstrap {
	return &config.Bootstrap{
		Scheduler: &config.Scheduler{
			TaskTimeout: durationpb.New(30 * time.Second),
		},
	}
}

func testTask(allowOverlap bool) *model.Task {
	return &model.Task{
		ID:             1,
		Name:           "noop",
		Title:          "空任务",
		Description:    "用于测试的空任务",
		CronSpec:       "*/5 * * * * *",
		Payload:        `{"name":"test"}`,
		TimeoutSeconds: 10,
		AllowOverlap:   allowOverlap,
		AlertEnabled:   true,
		Version:        1,
	}
}

func TestScheduleExecutionHighConcurrencyCreatesOneRecord(t *testing.T) {
	competitors := envInt("SCHEDULER_RACE_COMPETITORS", 1000)
	taskImpl := &fakeTask{}
	executionRepo := &fakeExecutionRepo{created: true}
	usecase := newTestTaskUsecase(t, executionRepo, taskImpl)
	task := testTask(true)
	scheduledAt := time.Date(2026, 7, 4, 10, 0, 0, 123, time.UTC)
	var wg sync.WaitGroup
	wg.Add(competitors)
	for i := 0; i < competitors; i++ {
		go func() {
			defer wg.Done()
			_, _ = usecase.ScheduleExecution(context.Background(), &TaskScheduleExecutionReq{Task: task, ScheduledAt: scheduledAt, TriggerType: schedulerenum.TaskTriggerTypeSchedule})
		}()
	}
	wg.Wait()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) && taskImpl.executions.Load() == 0 {
		time.Sleep(time.Millisecond)
	}
	executionRepo.mu.Lock()
	recordCount := len(executionRepo.recordsByPeriod)
	executionRepo.mu.Unlock()
	if recordCount != 1 {
		t.Fatalf("expected one execution record, got %d", recordCount)
	}
	if taskImpl.executions.Load() != 1 {
		t.Fatalf("expected one task execution, got %d", taskImpl.executions.Load())
	}
}

func TestScheduleExecutionCreatesOneRecordPerPeriod(t *testing.T) {
	periods := envInt("SCHEDULER_PERIOD_COUNT", 20)
	competitors := envInt("SCHEDULER_PERIOD_COMPETITORS", 100)
	taskImpl := &fakeTask{}
	executionRepo := &fakeExecutionRepo{created: true}
	usecase := newTestTaskUsecase(t, executionRepo, taskImpl)
	task := testTask(true)
	base := time.Date(2026, 7, 4, 10, 0, 0, 0, time.UTC)
	var wg sync.WaitGroup
	for period := 0; period < periods; period++ {
		scheduledAt := base.Add(time.Duration(period) * time.Second)
		for i := 0; i < competitors; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, _ = usecase.ScheduleExecution(context.Background(), &TaskScheduleExecutionReq{Task: task, ScheduledAt: scheduledAt, TriggerType: schedulerenum.TaskTriggerTypeSchedule})
			}()
		}
	}
	wg.Wait()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) && int(taskImpl.executions.Load()) < periods {
		time.Sleep(time.Millisecond)
	}
	executionRepo.mu.Lock()
	recordCount := len(executionRepo.recordsByPeriod)
	executionRepo.mu.Unlock()
	if recordCount != periods {
		t.Fatalf("expected %d execution records, got %d", periods, recordCount)
	}
	if int(taskImpl.executions.Load()) != periods {
		t.Fatalf("expected %d task executions, got %d", periods, taskImpl.executions.Load())
	}
}

func TestScheduleExecutionRedisFailureFallsBackToDatabaseUnique(t *testing.T) {
	competitors := envInt("SCHEDULER_REDIS_FALLBACK_COMPETITORS", 1000)
	taskImpl := &fakeTask{}
	executionRepo := &fakeExecutionRepo{created: true}
	usecase := NewTaskUsecase(
		slog.New(slog.NewTextHandler(io.Discard, nil)),
		testConf(),
		testTx,
		&fakeTaskRepo{task: testTask(true)},
		&fakeTaskVersionRepo{},
		executionRepo,
		&fakeTaskLockRepo{err: errors.New("redis unavailable")},
		map[string]taskimpl.Task{"noop": taskImpl},
		&fakeTaskEventBus{},
		&fakeAlert{},
	)
	task := testTask(true)
	scheduledAt := time.Date(2026, 7, 4, 10, 0, 0, 123, time.UTC)
	var wg sync.WaitGroup
	wg.Add(competitors)
	for i := 0; i < competitors; i++ {
		go func() {
			defer wg.Done()
			_, _ = usecase.ScheduleExecution(context.Background(), &TaskScheduleExecutionReq{Task: task, ScheduledAt: scheduledAt, TriggerType: schedulerenum.TaskTriggerTypeSchedule})
		}()
	}
	wg.Wait()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) && taskImpl.executions.Load() == 0 {
		time.Sleep(time.Millisecond)
	}
	executionRepo.mu.Lock()
	recordCount := len(executionRepo.recordsByPeriod)
	createCalled := executionRepo.createCalled
	executionRepo.mu.Unlock()
	if recordCount != 1 {
		t.Fatalf("expected one execution record, got %d", recordCount)
	}
	if taskImpl.executions.Load() != 1 {
		t.Fatalf("expected one task execution, got %d", taskImpl.executions.Load())
	}
	if createCalled == 0 {
		t.Fatal("expected database fallback create attempts")
	}
}

func TestScheduleExecutionRedisFailureFallbackOverlapSkipped(t *testing.T) {
	taskImpl := &fakeTask{}
	executionRepo := &fakeExecutionRepo{created: true, hasUnexpiredRunning: true}
	usecase := NewTaskUsecase(
		slog.New(slog.NewTextHandler(io.Discard, nil)),
		testConf(),
		testTx,
		&fakeTaskRepo{task: testTask(false)},
		&fakeTaskVersionRepo{},
		executionRepo,
		&fakeTaskLockRepo{err: errors.New("redis unavailable")},
		map[string]taskimpl.Task{"noop": taskImpl},
		&fakeTaskEventBus{},
		&fakeAlert{},
	)
	task := testTask(false)
	scheduleResp, err := usecase.ScheduleExecution(context.Background(), &TaskScheduleExecutionReq{Task: task, ScheduledAt: time.Now(), TriggerType: schedulerenum.TaskTriggerTypeSchedule})
	if err != nil {
		t.Fatalf("ScheduleExecution returned error: %v", err)
	}
	record := scheduleResp
	if record == nil || record.Status != schedulerenum.TaskExecutionStatusOverlapSkipped {
		t.Fatalf("expected overlap_skipped, got %#v", record)
	}
	if taskImpl.executions.Load() != 0 {
		t.Fatalf("expected no task execution, got %d", taskImpl.executions.Load())
	}
}

func TestStaleRunningCanBeMarkedUnknownAndCannotBeOverwritten(t *testing.T) {
	executionRepo := &fakeExecutionRepo{created: true}
	taskVersionRepo := &fakeTaskVersionRepo{
		created: []*model.TaskVersion{{
			TaskID:         testTask(true).ID,
			Version:        testTask(true).Version,
			CronSpec:       testTask(true).CronSpec,
			TimeoutSeconds: testTask(true).TimeoutSeconds,
		}},
	}
	usecase := NewTaskUsecase(
		slog.New(slog.NewTextHandler(io.Discard, nil)),
		testConf(),
		testTx,
		&fakeTaskRepo{task: testTask(true)},
		taskVersionRepo,
		executionRepo,
		&fakeTaskLockRepo{},
		map[string]taskimpl.Task{"noop": &fakeTask{block: make(chan struct{})}},
		&fakeTaskEventBus{},
		&fakeAlert{},
	)
	task := testTask(true)
	startResp, err := usecase.StartExecution(context.Background(), &TaskStartExecutionReq{Task: task, ScheduledAt: time.Now(), TriggerType: schedulerenum.TaskTriggerTypeSchedule})
	if err != nil || !startResp.Created {
		t.Fatalf("StartExecution returned record=%#v created=%v err=%v", startResp.Record, startResp.Created, err)
	}
	record := startResp.Record
	oldStartedAt := time.Now().Add(-time.Hour)
	executionRepo.mu.Lock()
	executionRepo.records[record.ID].StartedAt = &oldStartedAt
	executionRepo.mu.Unlock()
	runtimeResp, err := usecase.CheckExecutionRuntimes(context.Background(), []int64{record.ID})
	if err != nil {
		t.Fatalf("CheckExecutionRuntimes returned error: %v", err)
	}
	runtimes := runtimeResp
	if len(runtimes) != 1 || runtimes[0].State != schedulerenum.TaskExecutionRuntimeStateStale {
		t.Fatalf("expected stale runtime, got %#v", runtimes)
	}
	unknownResp, err := usecase.MarkExecutionsUnknown(context.Background(), []int64{record.ID})
	if err != nil {
		t.Fatalf("MarkExecutionsUnknown returned error: %v", err)
	}
	unknownRows := unknownResp
	if len(unknownRows) != 1 || unknownRows[0].Status != schedulerenum.TaskExecutionStatusUnknown {
		t.Fatalf("expected unknown row, got %#v", unknownRows)
	}
	finishedAt := time.Now()
	markResp, err := executionRepo.MarkFinished(context.Background(), &repo.TaskExecutionRecordMarkFinishedReq{ID: record.ID, Status: schedulerenum.TaskExecutionStatusSuccess, FinishedAt: finishedAt, DurationMS: 1})
	row := markResp.Row
	updated := markResp.Updated
	if err != nil {
		t.Fatalf("MarkFinished returned error: %v", err)
	}
	if updated {
		t.Fatal("expected old execution not to overwrite unknown status")
	}
	if row == nil || row.Status != schedulerenum.TaskExecutionStatusUnknown {
		t.Fatalf("expected unknown status, got %#v", row)
	}
}

func executionPeriodKey(taskID int64, scheduledAt time.Time) string {
	return strconv.FormatInt(taskID, 10) + ":" + scheduledAt.UTC().Format(time.RFC3339Nano)
}

func envInt(name string, defaultValue int) int {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return defaultValue
	}
	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return defaultValue
	}
	return parsed
}
