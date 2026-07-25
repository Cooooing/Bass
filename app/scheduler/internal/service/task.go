package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	schedulerv1 "common/proto/gen/scheduler/v1"
	"context"
	"scheduler/internal/biz/model"
	"scheduler/internal/biz/usecase"
	schedulerenum "scheduler/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SchedulerTaskService struct {
	schedulerv1.UnimplementedSchedulerTaskServiceServer
	taskUsecase *usecase.TaskUsecase
}

func NewSchedulerTaskService(
	taskUsecase *usecase.TaskUsecase,
) *SchedulerTaskService {
	return &SchedulerTaskService{
		taskUsecase: taskUsecase,
	}
}

func (s *SchedulerTaskService) RegisterGrpc(
	gs *grpc.Server,
) {
	schedulerv1.RegisterSchedulerTaskServiceServer(gs, s)
}

func (s *SchedulerTaskService) RegisterHttp(
	hs *http.Server,
) {
}

func (s *SchedulerTaskService) Upsert(
	ctx context.Context,
	req *schedulerv1.UpsertSchedulerTask_Req,
) (*schedulerv1.UpsertSchedulerTask_Resp, error) {
	if req.GetName() == "" || req.GetTitle() == "" || req.GetCronSpec() == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	alertEnabled := true
	if req.AlertEnabled != nil {
		alertEnabled = req.GetAlertEnabled()
	}
	upsertRow, err := s.taskUsecase.Upsert(ctx, &model.Task{
		ID:             req.GetId(),
		Name:           req.GetName(),
		Title:          req.GetTitle(),
		Description:    req.GetDescription(),
		Enabled:        req.GetEnabled(),
		CronSpec:       req.GetCronSpec(),
		Payload:        req.GetPayload(),
		TimeoutSeconds: req.GetTimeoutSeconds(),
		AllowOverlap:   req.GetAllowOverlap(),
		AlertEnabled:   alertEnabled,
	})
	if err != nil {
		return nil, err
	}
	row := upsertRow
	var createdAt *timestamppb.Timestamp
	if row.CreatedAt != nil {
		createdAt = timestamppb.New(*row.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if row.UpdatedAt != nil {
		updatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &schedulerv1.UpsertSchedulerTask_Resp{
		Row: &schedulerv1.UpsertSchedulerTask_Resp_SchedulerTask{
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			Id:             row.ID,
			Name:           row.Name,
			Title:          row.Title,
			Description:    row.Description,
			Enabled:        row.Enabled,
			CronSpec:       row.CronSpec,
			Payload:        row.Payload,
			TimeoutSeconds: row.TimeoutSeconds,
			AllowOverlap:   row.AllowOverlap,
			AlertEnabled:   row.AlertEnabled,
			Version:        row.Version,
		},
	}, nil
}

func (s *SchedulerTaskService) Get(
	ctx context.Context,
	req *schedulerv1.GetSchedulerTask_Req,
) (*schedulerv1.GetSchedulerTask_Resp, error) {
	row, err := s.taskUsecase.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	var createdAt *timestamppb.Timestamp
	if row.CreatedAt != nil {
		createdAt = timestamppb.New(*row.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if row.UpdatedAt != nil {
		updatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &schedulerv1.GetSchedulerTask_Resp{
		Row: &schedulerv1.GetSchedulerTask_Resp_SchedulerTask{
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			Id:             row.ID,
			Name:           row.Name,
			Title:          row.Title,
			Description:    row.Description,
			Enabled:        row.Enabled,
			CronSpec:       row.CronSpec,
			Payload:        row.Payload,
			TimeoutSeconds: row.TimeoutSeconds,
			AllowOverlap:   row.AllowOverlap,
			AlertEnabled:   row.AlertEnabled,
			Version:        row.Version,
		},
	}, nil
}

func (s *SchedulerTaskService) Page(
	ctx context.Context,
	req *schedulerv1.PageSchedulerTasks_Req,
) (*schedulerv1.PageSchedulerTasks_Resp, error) {
	query := &usecase.TaskPageReq{}
	if req.GetQuery() != nil {
		query.IDs = req.GetQuery().GetIds()
		if req.GetQuery().Name != nil {
			query.Name = new(req.GetQuery().GetName())
		}
		if req.GetQuery().Title != nil {
			query.Title = new(req.GetQuery().GetTitle())
		}
		if req.GetQuery().Enabled != nil {
			query.Enabled = new(req.GetQuery().GetEnabled())
		}
	}
	query.Page = req.GetPage()
	pageResp, err := s.taskUsecase.Page(ctx, query)
	if err != nil {
		return nil, err
	}
	rows := pageResp.Rows
	replyRows := make([]*schedulerv1.PageSchedulerTasks_Resp_SchedulerTask, 0, len(rows))
	for _, row := range rows {
		var createdAt *timestamppb.Timestamp
		if row.CreatedAt != nil {
			createdAt = timestamppb.New(*row.CreatedAt)
		}
		var updatedAt *timestamppb.Timestamp
		if row.UpdatedAt != nil {
			updatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, &schedulerv1.PageSchedulerTasks_Resp_SchedulerTask{
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			Id:             row.ID,
			Name:           row.Name,
			Title:          row.Title,
			Description:    row.Description,
			Enabled:        row.Enabled,
			CronSpec:       row.CronSpec,
			Payload:        row.Payload,
			TimeoutSeconds: row.TimeoutSeconds,
			AllowOverlap:   row.AllowOverlap,
			AlertEnabled:   row.AlertEnabled,
			Version:        row.Version,
		})
	}
	return &schedulerv1.PageSchedulerTasks_Resp{
		Page: pageResp.Page,
		Rows: replyRows,
	}, nil
}

func (s *SchedulerTaskService) ListAvailableTasks(
	ctx context.Context,
	req *schedulerv1.ListSchedulerAvailableTasks_Req,
) (*schedulerv1.ListSchedulerAvailableTasks_Resp, error) {
	keyword := ""
	if req.GetQuery() != nil {
		keyword = req.GetQuery().GetKeyword()
	}
	rows, err := s.taskUsecase.ListAvailableTasks(ctx, keyword)
	if err != nil {
		return nil, err
	}
	replyRows := make([]*schedulerv1.ListSchedulerAvailableTasks_Resp_SchedulerAvailableTask, 0, len(rows))
	for _, row := range rows {
		replyRows = append(replyRows, &schedulerv1.ListSchedulerAvailableTasks_Resp_SchedulerAvailableTask{
			Name:        row.Name,
			Title:       row.Title,
			Description: row.Description,
		})
	}
	return &schedulerv1.ListSchedulerAvailableTasks_Resp{
		Rows: replyRows,
	}, nil
}

func (s *SchedulerTaskService) PageExecutionRecords(
	ctx context.Context,
	req *schedulerv1.PageSchedulerTaskExecutionRecords_Req,
) (*schedulerv1.PageSchedulerTaskExecutionRecords_Resp, error) {
	query := &usecase.TaskExecutionRecordPageReq{}
	if req.GetQuery() != nil {
		query.IDs = req.GetQuery().GetIds()
		if req.GetQuery().TaskId != nil {
			query.TaskID = new(req.GetQuery().GetTaskId())
		}
		if req.GetQuery().Status != nil {
			statusValue, ok := schedulerenum.TaskExecutionStatusMap.ToEnum(req.GetQuery().GetStatus())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.Status = &statusValue
		}
		if req.GetQuery().TriggerType != nil {
			triggerType, ok := schedulerenum.TaskTriggerTypeMap.ToEnum(req.GetQuery().GetTriggerType())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.TriggerType = &triggerType
		}
	}
	query.Page = req.GetPage()
	pageResp, err := s.taskUsecase.PageExecutionRecords(ctx, query)
	if err != nil {
		return nil, err
	}
	rows := pageResp.Rows
	replyRows := make([]*schedulerv1.PageSchedulerTaskExecutionRecords_Resp_SchedulerTaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		var createdAt *timestamppb.Timestamp
		if row.CreatedAt != nil {
			createdAt = timestamppb.New(*row.CreatedAt)
		}
		var updatedAt *timestamppb.Timestamp
		if row.UpdatedAt != nil {
			updatedAt = timestamppb.New(*row.UpdatedAt)
		}
		var startedAt *timestamppb.Timestamp
		if row.StartedAt != nil {
			startedAt = timestamppb.New(*row.StartedAt)
		}
		var finishedAt *timestamppb.Timestamp
		if row.FinishedAt != nil {
			finishedAt = timestamppb.New(*row.FinishedAt)
		}
		statusValue := schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status)
		triggerType := schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType)
		replyRows = append(replyRows, &schedulerv1.PageSchedulerTaskExecutionRecords_Resp_SchedulerTaskExecutionRecord{
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Id:          row.ID,
			TaskId:      row.TaskID,
			ScheduledAt: timestamppb.New(row.ScheduledAt),
			StartedAt:   startedAt,
			FinishedAt:  finishedAt,
			DurationMs:  row.DurationMS,
			Status:      statusValue,
			TriggerType: triggerType,
			TaskVersion: row.TaskVersion,
			WorkerId:    row.WorkerID,
			Payload:     row.Payload,
			LastError:   row.LastError,
			TraceId:     row.TraceID,
		})
	}
	return &schedulerv1.PageSchedulerTaskExecutionRecords_Resp{
		Page: pageResp.Page,
		Rows: replyRows,
	}, nil
}

func (s *SchedulerTaskService) Trigger(
	ctx context.Context,
	req *schedulerv1.TriggerSchedulerTask_Req,
) (*schedulerv1.TriggerSchedulerTask_Resp, error) {
	row, err := s.taskUsecase.Trigger(ctx, &usecase.TaskTriggerReq{
		ID:      req.GetId(),
		Payload: req.GetPayload(),
	})
	if err != nil {
		return nil, err
	}
	var createdAt *timestamppb.Timestamp
	if row.CreatedAt != nil {
		createdAt = timestamppb.New(*row.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if row.UpdatedAt != nil {
		updatedAt = timestamppb.New(*row.UpdatedAt)
	}
	var startedAt *timestamppb.Timestamp
	if row.StartedAt != nil {
		startedAt = timestamppb.New(*row.StartedAt)
	}
	var finishedAt *timestamppb.Timestamp
	if row.FinishedAt != nil {
		finishedAt = timestamppb.New(*row.FinishedAt)
	}
	statusValue := schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status)
	triggerType := schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType)
	return &schedulerv1.TriggerSchedulerTask_Resp{
		Row: &schedulerv1.TriggerSchedulerTask_Resp_SchedulerTaskExecutionRecord{
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Id:          row.ID,
			TaskId:      row.TaskID,
			ScheduledAt: timestamppb.New(row.ScheduledAt),
			StartedAt:   startedAt,
			FinishedAt:  finishedAt,
			DurationMs:  row.DurationMS,
			Status:      statusValue,
			TriggerType: triggerType,
			TaskVersion: row.TaskVersion,
			WorkerId:    row.WorkerID,
			Payload:     row.Payload,
			LastError:   row.LastError,
			TraceId:     row.TraceID,
		},
	}, nil
}

func (s *SchedulerTaskService) CancelExecution(
	ctx context.Context,
	req *schedulerv1.CancelSchedulerTaskExecution_Req,
) (*schedulerv1.CancelSchedulerTaskExecution_Resp, error) {
	row, err := s.taskUsecase.CancelExecution(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	var createdAt *timestamppb.Timestamp
	if row.CreatedAt != nil {
		createdAt = timestamppb.New(*row.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if row.UpdatedAt != nil {
		updatedAt = timestamppb.New(*row.UpdatedAt)
	}
	var startedAt *timestamppb.Timestamp
	if row.StartedAt != nil {
		startedAt = timestamppb.New(*row.StartedAt)
	}
	var finishedAt *timestamppb.Timestamp
	if row.FinishedAt != nil {
		finishedAt = timestamppb.New(*row.FinishedAt)
	}
	statusValue := schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status)
	triggerType := schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType)
	return &schedulerv1.CancelSchedulerTaskExecution_Resp{
		Row: &schedulerv1.CancelSchedulerTaskExecution_Resp_SchedulerTaskExecutionRecord{
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Id:          row.ID,
			TaskId:      row.TaskID,
			ScheduledAt: timestamppb.New(row.ScheduledAt),
			StartedAt:   startedAt,
			FinishedAt:  finishedAt,
			DurationMs:  row.DurationMS,
			Status:      statusValue,
			TriggerType: triggerType,
			TaskVersion: row.TaskVersion,
			WorkerId:    row.WorkerID,
			Payload:     row.Payload,
			LastError:   row.LastError,
			TraceId:     row.TraceID,
		},
	}, nil
}

func (s *SchedulerTaskService) CheckExecutionRuntimes(
	ctx context.Context,
	req *schedulerv1.CheckSchedulerTaskExecutionRuntimes_Req,
) (*schedulerv1.CheckSchedulerTaskExecutionRuntimes_Resp, error) {
	rows, err := s.taskUsecase.CheckExecutionRuntimes(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	replyRows := make([]*schedulerv1.CheckSchedulerTaskExecutionRuntimes_Resp_SchedulerTaskExecutionRuntime, 0, len(rows))
	for _, row := range rows {
		stateValue := schedulerenum.TaskExecutionRuntimeStateMap.MustToProto(row.State)
		replyRows = append(replyRows, &schedulerv1.CheckSchedulerTaskExecutionRuntimes_Resp_SchedulerTaskExecutionRuntime{
			ExecutionRecordId: row.ExecutionRecordID,
			TaskId:            row.TaskID,
			State:             stateValue,
		})
	}
	return &schedulerv1.CheckSchedulerTaskExecutionRuntimes_Resp{
		Rows: replyRows,
	}, nil
}

func (s *SchedulerTaskService) MarkExecutionsUnknown(
	ctx context.Context,
	req *schedulerv1.MarkSchedulerTaskExecutionsUnknown_Req,
) (*schedulerv1.MarkSchedulerTaskExecutionsUnknown_Resp, error) {
	rows, err := s.taskUsecase.MarkExecutionsUnknown(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	replyRows := make([]*schedulerv1.MarkSchedulerTaskExecutionsUnknown_Resp_SchedulerTaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		var createdAt *timestamppb.Timestamp
		if row.CreatedAt != nil {
			createdAt = timestamppb.New(*row.CreatedAt)
		}
		var updatedAt *timestamppb.Timestamp
		if row.UpdatedAt != nil {
			updatedAt = timestamppb.New(*row.UpdatedAt)
		}
		var startedAt *timestamppb.Timestamp
		if row.StartedAt != nil {
			startedAt = timestamppb.New(*row.StartedAt)
		}
		var finishedAt *timestamppb.Timestamp
		if row.FinishedAt != nil {
			finishedAt = timestamppb.New(*row.FinishedAt)
		}
		statusValue := schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status)
		triggerType := schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType)
		replyRows = append(replyRows, &schedulerv1.MarkSchedulerTaskExecutionsUnknown_Resp_SchedulerTaskExecutionRecord{
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Id:          row.ID,
			TaskId:      row.TaskID,
			ScheduledAt: timestamppb.New(row.ScheduledAt),
			StartedAt:   startedAt,
			FinishedAt:  finishedAt,
			DurationMs:  row.DurationMS,
			Status:      statusValue,
			TriggerType: triggerType,
			TaskVersion: row.TaskVersion,
			WorkerId:    row.WorkerID,
			Payload:     row.Payload,
			LastError:   row.LastError,
			TraceId:     row.TraceID,
		})
	}
	return &schedulerv1.MarkSchedulerTaskExecutionsUnknown_Resp{
		Rows:         replyRows,
		UpdatedCount: uint32(len(replyRows)),
	}, nil
}
