package repo

import (
	"context"
	"time"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/agentjob"
	"game_town/internal/enum"

	"github.com/samber/lo"
)

var _ bizrepo.AgentJobRepo = (*AgentJobRepo)(nil)

type AgentJobRepo struct {
	db *gen.Client
}

func NewAgentJobRepo(
	db *gen.Client,
) bizrepo.AgentJobRepo {
	return &AgentJobRepo{
		db: db,
	}
}

func (r *AgentJobRepo) getClient(
	ctx context.Context,
) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *AgentJobRepo) Save(
	ctx context.Context,
	row *model.AgentJob,
) (*model.AgentJob, error) {
	saved, err := r.getClient(ctx).AgentJob.Create().
		SetWorldID(row.WorldID).
		SetSourceEventID(row.SourceEventID).
		SetType(agentjob.Type(row.Type)).
		SetPriority(agentjob.Priority(row.Priority)).
		SetLaneKey(row.LaneKey).
		SetStatus(agentjob.Status(row.Status)).
		SetWorldVersion(row.WorldVersion).
		SetNillableNpcID(row.NpcID).
		SetAttemptCount(row.AttemptCount).
		SetAvailableAt(row.AvailableAt).
		SetErrorSummary(row.ErrorSummary).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.AgentJob{
		ID:            saved.ID,
		WorldID:       saved.WorldID,
		SourceEventID: saved.SourceEventID,
		Type:          enum.AgentJobType(saved.Type),
		Priority:      enum.AgentJobPriority(saved.Priority),
		LaneKey:       saved.LaneKey,
		Status:        enum.AgentJobStatus(saved.Status),
		WorldVersion:  saved.WorldVersion,
		NpcID:         saved.NpcID,
		AttemptCount:  saved.AttemptCount,
		AvailableAt:   saved.AvailableAt,
		StartedAt:     saved.StartedAt,
		FinishedAt:    saved.FinishedAt,
		ErrorSummary:  saved.ErrorSummary,
		CreatedAt:     saved.CreatedAt,
		UpdatedAt:     saved.UpdatedAt,
	}, nil
}

func agentJobQuery(
	q *gen.AgentJobQuery,
	req *bizrepo.AgentJobQuery,
) *gen.AgentJobQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(agentjob.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		q = q.Where(agentjob.IDIn(req.IDs...))
	}
	if req.WorldID != nil {
		q = q.Where(agentjob.WorldID(*req.WorldID))
	}
	if req.SourceEventID != nil {
		q = q.Where(agentjob.SourceEventID(*req.SourceEventID))
	}
	if req.Type != nil {
		q = q.Where(agentjob.TypeEQ(agentjob.Type(*req.Type)))
	}
	if req.Priority != nil {
		q = q.Where(agentjob.PriorityEQ(agentjob.Priority(*req.Priority)))
	}
	if req.Status != nil {
		q = q.Where(agentjob.StatusEQ(agentjob.Status(*req.Status)))
	}
	if len(req.Statuses) > 0 {
		statuses := lo.Map(req.Statuses, func(status enum.AgentJobStatus, _ int) agentjob.Status {
			return agentjob.Status(status)
		})
		q = q.Where(agentjob.StatusIn(statuses...))
	}
	if req.LaneKey != nil {
		q = q.Where(agentjob.LaneKey(*req.LaneKey))
	}
	if req.AvailableBefore != nil {
		q = q.Where(agentjob.AvailableAtLTE(*req.AvailableBefore))
	}
	if req.StartedBefore != nil {
		q = q.Where(agentjob.StartedAtNotNil(), agentjob.StartedAtLT(*req.StartedBefore))
	}
	return q
}

func (r *AgentJobRepo) Get(
	ctx context.Context,
	req *bizrepo.AgentJobQuery,
) (*model.AgentJob, error) {
	row, err := agentJobQuery(r.getClient(ctx).AgentJob.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.AgentJob{
		ID:            row.ID,
		WorldID:       row.WorldID,
		SourceEventID: row.SourceEventID,
		Type:          enum.AgentJobType(row.Type),
		Priority:      enum.AgentJobPriority(row.Priority),
		LaneKey:       row.LaneKey,
		Status:        enum.AgentJobStatus(row.Status),
		WorldVersion:  row.WorldVersion,
		NpcID:         row.NpcID,
		AttemptCount:  row.AttemptCount,
		AvailableAt:   row.AvailableAt,
		StartedAt:     row.StartedAt,
		FinishedAt:    row.FinishedAt,
		ErrorSummary:  row.ErrorSummary,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}, nil
}

func (r *AgentJobRepo) List(
	ctx context.Context,
	req *bizrepo.AgentJobQuery,
) ([]*model.AgentJob, error) {
	rows, err := agentJobQuery(r.getClient(ctx).AgentJob.Query(), req).Order(agentjob.ByPriority(), agentjob.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.AgentJob, _ int) *model.AgentJob {
		return &model.AgentJob{
			ID:            row.ID,
			WorldID:       row.WorldID,
			SourceEventID: row.SourceEventID,
			Type:          enum.AgentJobType(row.Type),
			Priority:      enum.AgentJobPriority(row.Priority),
			LaneKey:       row.LaneKey,
			Status:        enum.AgentJobStatus(row.Status),
			WorldVersion:  row.WorldVersion,
			NpcID:         row.NpcID,
			AttemptCount:  row.AttemptCount,
			AvailableAt:   row.AvailableAt,
			StartedAt:     row.StartedAt,
			FinishedAt:    row.FinishedAt,
			ErrorSummary:  row.ErrorSummary,
			CreatedAt:     row.CreatedAt,
			UpdatedAt:     row.UpdatedAt,
		}
	})
	return out, nil
}

func (r *AgentJobRepo) Map(
	ctx context.Context,
	req *bizrepo.AgentJobQuery,
) (map[int64]*model.AgentJob, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.AgentJob, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *AgentJobRepo) Count(
	ctx context.Context,
	req *bizrepo.AgentJobQuery,
) (int, error) {
	return agentJobQuery(r.getClient(ctx).AgentJob.Query(), req).Count(ctx)
}

func (r *AgentJobRepo) Page(
	ctx context.Context,
	req *bizrepo.AgentJobPageReq,
) (*bizrepo.AgentJobPageResp, error) {
	p := page(req.Page)
	q := agentJobQuery(r.getClient(ctx).AgentJob.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(agentjob.ByPriority(), agentjob.ByID()).Offset(pageOffset(p)).Limit(pageLimit(p)).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.AgentJob, _ int) *model.AgentJob {
		return &model.AgentJob{
			ID:            row.ID,
			WorldID:       row.WorldID,
			SourceEventID: row.SourceEventID,
			Type:          enum.AgentJobType(row.Type),
			Priority:      enum.AgentJobPriority(row.Priority),
			LaneKey:       row.LaneKey,
			Status:        enum.AgentJobStatus(row.Status),
			WorldVersion:  row.WorldVersion,
			NpcID:         row.NpcID,
			AttemptCount:  row.AttemptCount,
			AvailableAt:   row.AvailableAt,
			StartedAt:     row.StartedAt,
			FinishedAt:    row.FinishedAt,
			ErrorSummary:  row.ErrorSummary,
			CreatedAt:     row.CreatedAt,
			UpdatedAt:     row.UpdatedAt,
		}
	})
	return &bizrepo.AgentJobPageResp{
		Rows: out,
		Page: basePage(total, p),
	}, nil
}

func (r *AgentJobRepo) MarkRunning(
	ctx context.Context,
	id int64,
	now time.Time,
) (*model.AgentJob, error) {
	row, err := r.getClient(ctx).AgentJob.UpdateOneID(id).SetStatus(agentjob.StatusRunning).SetStartedAt(now).AddAttemptCount(1).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.AgentJob{
		ID:            row.ID,
		WorldID:       row.WorldID,
		SourceEventID: row.SourceEventID,
		Type:          enum.AgentJobType(row.Type),
		Priority:      enum.AgentJobPriority(row.Priority),
		LaneKey:       row.LaneKey,
		Status:        enum.AgentJobStatus(row.Status),
		WorldVersion:  row.WorldVersion,
		NpcID:         row.NpcID,
		AttemptCount:  row.AttemptCount,
		AvailableAt:   row.AvailableAt,
		StartedAt:     row.StartedAt,
		FinishedAt:    row.FinishedAt,
		ErrorSummary:  row.ErrorSummary,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}, nil
}

func (r *AgentJobRepo) MarkSucceeded(
	ctx context.Context,
	id int64,
	now time.Time,
) (*model.AgentJob, error) {
	row, err := r.getClient(ctx).AgentJob.UpdateOneID(id).SetStatus(agentjob.StatusSucceeded).SetFinishedAt(now).SetErrorSummary("").Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.AgentJob{
		ID:            row.ID,
		WorldID:       row.WorldID,
		SourceEventID: row.SourceEventID,
		Type:          enum.AgentJobType(row.Type),
		Priority:      enum.AgentJobPriority(row.Priority),
		LaneKey:       row.LaneKey,
		Status:        enum.AgentJobStatus(row.Status),
		WorldVersion:  row.WorldVersion,
		NpcID:         row.NpcID,
		AttemptCount:  row.AttemptCount,
		AvailableAt:   row.AvailableAt,
		StartedAt:     row.StartedAt,
		FinishedAt:    row.FinishedAt,
		ErrorSummary:  row.ErrorSummary,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}, nil
}

func (r *AgentJobRepo) Retry(
	ctx context.Context,
	req *bizrepo.AgentJobRetryReq,
) (*model.AgentJob, error) {
	row, err := r.getClient(ctx).AgentJob.UpdateOneID(req.JobID).
		SetStatus(agentjob.StatusQueued).
		SetAttemptCount(req.AttemptCount).
		SetAvailableAt(req.AvailableAt).
		SetErrorSummary(req.ErrorSummary).
		ClearStartedAt().
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.AgentJob{
		ID:            row.ID,
		WorldID:       row.WorldID,
		SourceEventID: row.SourceEventID,
		Type:          enum.AgentJobType(row.Type),
		Priority:      enum.AgentJobPriority(row.Priority),
		LaneKey:       row.LaneKey,
		Status:        enum.AgentJobStatus(row.Status),
		WorldVersion:  row.WorldVersion,
		NpcID:         row.NpcID,
		AttemptCount:  row.AttemptCount,
		AvailableAt:   row.AvailableAt,
		StartedAt:     row.StartedAt,
		FinishedAt:    row.FinishedAt,
		ErrorSummary:  row.ErrorSummary,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}, nil
}

func (r *AgentJobRepo) MarkFailed(
	ctx context.Context,
	req *bizrepo.AgentJobMarkFailedReq,
) (*model.AgentJob, error) {
	row, err := r.getClient(ctx).AgentJob.UpdateOneID(req.JobID).
		SetStatus(agentjob.StatusFailed).
		SetFinishedAt(req.FinishedAt).
		SetErrorSummary(req.ErrorSummary).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.AgentJob{
		ID:            row.ID,
		WorldID:       row.WorldID,
		SourceEventID: row.SourceEventID,
		Type:          enum.AgentJobType(row.Type),
		Priority:      enum.AgentJobPriority(row.Priority),
		LaneKey:       row.LaneKey,
		Status:        enum.AgentJobStatus(row.Status),
		WorldVersion:  row.WorldVersion,
		NpcID:         row.NpcID,
		AttemptCount:  row.AttemptCount,
		AvailableAt:   row.AvailableAt,
		StartedAt:     row.StartedAt,
		FinishedAt:    row.FinishedAt,
		ErrorSummary:  row.ErrorSummary,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}, nil
}
