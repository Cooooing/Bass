package usecase

import (
	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"fmt"
	"log/slog"
	"os"
	"scheduler/internal/biz/model"
	"scheduler/internal/biz/repo"
	taskimpl "scheduler/internal/biz/usecase/task"
	"scheduler/internal/config"
	schedulerenum "scheduler/internal/enum"
	"strings"
	"time"
)

type DelayedTaskUsecase struct {
	logger   *slog.Logger
	conf     *config.Bootstrap
	repo     repo.DelayedTaskRepo
	tasks    map[string]taskimpl.Task
	workerID string
}

func NewDelayedTaskUsecase(
	logger *slog.Logger,
	conf *config.Bootstrap,
	delayedTaskRepo repo.DelayedTaskRepo,
	tasks map[string]taskimpl.Task,
) *DelayedTaskUsecase {
	workerID := "scheduler"
	if host, err := os.Hostname(); err == nil && host != "" {
		workerID = host
	}
	return &DelayedTaskUsecase{
		logger:   logger,
		conf:     conf,
		repo:     delayedTaskRepo,
		tasks:    tasks,
		workerID: workerID,
	}
}

type DelayedTaskRegisterReq struct {
	IdempotencyKey string
	TaskName       string
	Payload        string
	ExecuteAt      time.Time
	MaxAttempts    int32
	TimeoutSeconds int32
}

func (u *DelayedTaskUsecase) Register(ctx context.Context, req *DelayedTaskRegisterReq) (*model.DelayedTask, error) {
	payload := strings.TrimSpace(req.Payload)
	if payload == "" {
		payload = "{}"
	}
	if req.MaxAttempts <= 0 {
		req.MaxAttempts = 3
	}
	if req.TimeoutSeconds <= 0 {
		req.TimeoutSeconds = 30
	}
	return u.repo.Register(ctx, &model.DelayedTask{
		IdempotencyKey: strings.TrimSpace(req.IdempotencyKey),
		TaskName:       strings.TrimSpace(req.TaskName),
		Payload:        payload,
		ExecuteAt:      req.ExecuteAt,
		NextRunAt:      req.ExecuteAt,
		Status:         schedulerenum.DelayedTaskStatusPending,
		MaxAttempts:    req.MaxAttempts,
		TimeoutSeconds: req.TimeoutSeconds,
	})
}

func (u *DelayedTaskUsecase) Cancel(ctx context.Context, id int64, idempotencyKey string) error {
	if id != 0 {
		_, err := u.repo.Cancel(ctx, &repo.DelayedTaskGetReq{
			ID: new(id),
		})
		return err
	}
	key := strings.TrimSpace(idempotencyKey)
	_, err := u.repo.Cancel(ctx, &repo.DelayedTaskGetReq{
		IdempotencyKey: new(key),
	})
	return err
}

func (u *DelayedTaskUsecase) Get(ctx context.Context, id int64, idempotencyKey string) (*model.DelayedTask, error) {
	if id != 0 {
		return u.repo.Get(ctx, &repo.DelayedTaskGetReq{
			ID: new(id),
		})
	}
	key := strings.TrimSpace(idempotencyKey)
	return u.repo.Get(ctx, &repo.DelayedTaskGetReq{
		IdempotencyKey: new(key),
	})
}

type DelayedTaskPageReq struct {
	Page     *common.PageReq
	TaskName *string
	Status   *schedulerenum.DelayedTaskStatus
}
type DelayedTaskPageResp struct {
	Rows []*model.DelayedTask
	Page *common.PageResp
}

func (u *DelayedTaskUsecase) Page(ctx context.Context, req *DelayedTaskPageReq) (*DelayedTaskPageResp, error) {
	resp, err := u.repo.Page(ctx, &repo.DelayedTaskPageReq{
		Page: req.Page,
		DelayedTaskGetReq: repo.DelayedTaskGetReq{
			TaskName: req.TaskName,
			Status:   req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return &DelayedTaskPageResp{
		Rows: resp.Rows,
		Page: resp.Page,
	}, nil
}

func (u *DelayedTaskUsecase) Trigger(ctx context.Context, id int64) error {
	row, err := u.repo.Get(ctx, &repo.DelayedTaskGetReq{
		ID: new(id),
	})
	if err != nil {
		return err
	}
	if row == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return u.Execute(ctx, row)
}

func (u *DelayedTaskUsecase) RunDue(ctx context.Context, limit int) error {
	rows, err := u.repo.ListDue(ctx, time.Now(), limit)
	if err != nil {
		return err
	}
	for _, row := range rows {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		claimed, claimedRow, err := u.repo.MarkRunning(ctx, row.ID, u.workerID, time.Now().Add(time.Duration(row.TimeoutSeconds)*time.Second))
		if err != nil {
			u.logger.ErrorContext(ctx, "claim delayed task failed", "delayed_task_id", row.ID, "err", err)
			continue
		}
		if !claimed || claimedRow == nil {
			continue
		}
		go func(item *model.DelayedTask) {
			if err := u.Execute(context.WithoutCancel(ctx), item); err != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "execute delayed task failed", "delayed_task_id", item.ID, "err", err)
			}
		}(claimedRow)
	}
	return nil
}

func (u *DelayedTaskUsecase) Execute(ctx context.Context, row *model.DelayedTask) error {
	if row == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	task, ok := u.tasks[row.TaskName]
	if !ok {
		return u.markDelayedTaskFailed(ctx, row, fmt.Errorf("unknown delayed task: %s", row.TaskName))
	}
	timeout := time.Duration(row.TimeoutSeconds) * time.Second
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	callCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if err := task.Execute(callCtx, row.Payload); err != nil {
		return u.markDelayedTaskFailed(ctx, row, err)
	}
	_, err := u.repo.MarkSuccess(ctx, row.ID, time.Now())
	return err
}

func (u *DelayedTaskUsecase) markDelayedTaskFailed(ctx context.Context, row *model.DelayedTask, execErr error) error {
	attempt := row.Attempt + 1
	final := attempt >= row.MaxAttempts
	next := time.Now().Add(time.Duration(attempt) * time.Minute)
	lastError := execErr.Error()
	if len(lastError) > 2048 {
		lastError = lastError[:2048]
	}
	_, err := u.repo.MarkFailed(ctx, row.ID, attempt, final, next, lastError)
	if err != nil {
		return err
	}
	return execErr
}
