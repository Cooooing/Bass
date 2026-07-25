package repo

import (
	"content/internal/biz/base"
	"context"

	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/contentmoderationrecord"
	"content/internal/enum"

	"github.com/samber/lo"
)

var _ repo.ContentModerationRecordRepo = (*ContentModerationRecordRepo)(nil)

type ContentModerationRecordRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewContentModerationRecordRepo(
	db *gen.Client,
) repo.ContentModerationRecordRepo {
	return &ContentModerationRecordRepo{
		db: db,
	}
}

func (r *ContentModerationRecordRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ContentModerationRecordRepo) Save(ctx context.Context, record *model.ContentModerationRecord) (*model.ContentModerationRecord, error) {
	save, err := r.getClient(ctx).ContentModerationRecord.Create().
		SetTarget(contentmoderationrecord.Target(record.Target)).
		SetTargetID(record.TargetID).
		SetAction(contentmoderationrecord.Action(record.Action)).
		SetNillableReasonCode(record.ReasonCode).
		SetNillableReason(record.Reason).
		SetOperatorID(record.OperatorID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ContentModerationRecord{
		ID:         save.ID,
		Target:     enum.ContentModerationTarget(save.Target),
		TargetID:   save.TargetID,
		Action:     enum.ContentModerationAction(save.Action),
		ReasonCode: save.ReasonCode,
		Reason:     save.Reason,
		OperatorID: save.OperatorID,
		CreatedAt:  save.CreatedAt,
		UpdatedAt:  save.UpdatedAt,
	}, nil
}

func (r *ContentModerationRecordRepo) Get(ctx context.Context, req *repo.ContentModerationRecordGetReq) (*model.ContentModerationRecord, error) {
	query := r.getClient(ctx).ContentModerationRecord.Query()
	query = r.getQuery(query, req)
	record, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ContentModerationRecord{
		ID:         record.ID,
		Target:     enum.ContentModerationTarget(record.Target),
		TargetID:   record.TargetID,
		Action:     enum.ContentModerationAction(record.Action),
		ReasonCode: record.ReasonCode,
		Reason:     record.Reason,
		OperatorID: record.OperatorID,
		CreatedAt:  record.CreatedAt,
		UpdatedAt:  record.UpdatedAt,
	}, nil
}

func (r *ContentModerationRecordRepo) List(ctx context.Context, req *repo.ContentModerationRecordGetReq) ([]*model.ContentModerationRecord, error) {
	query := r.getClient(ctx).ContentModerationRecord.Query()
	query = r.getQuery(query, req)
	records, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(records, func(record *gen.ContentModerationRecord, _ int) *model.ContentModerationRecord {
		return &model.ContentModerationRecord{
			ID:         record.ID,
			Target:     enum.ContentModerationTarget(record.Target),
			TargetID:   record.TargetID,
			Action:     enum.ContentModerationAction(record.Action),
			ReasonCode: record.ReasonCode,
			Reason:     record.Reason,
			OperatorID: record.OperatorID,
			CreatedAt:  record.CreatedAt,
			UpdatedAt:  record.UpdatedAt,
		}
	}), nil
}

func (r *ContentModerationRecordRepo) Map(ctx context.Context, req *repo.ContentModerationRecordGetReq) (map[int64]*model.ContentModerationRecord, error) {
	records, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(records, func(record *model.ContentModerationRecord) (int64, *model.ContentModerationRecord) {
		return record.ID, record
	}), nil
}

func (r *ContentModerationRecordRepo) Count(ctx context.Context, req *repo.ContentModerationRecordGetReq) (int, error) {
	query := r.getClient(ctx).ContentModerationRecord.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ContentModerationRecordRepo) Page(ctx context.Context, req *repo.ContentModerationRecordGetReq) (*repo.ContentModerationRecordPageResp, error) {
	page := r.normalizePage(req.Page)
	query := r.getClient(ctx).ContentModerationRecord.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	records, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	rows := lo.Map(records, func(record *gen.ContentModerationRecord, _ int) *model.ContentModerationRecord {
		return &model.ContentModerationRecord{
			ID:         record.ID,
			Target:     enum.ContentModerationTarget(record.Target),
			TargetID:   record.TargetID,
			Action:     enum.ContentModerationAction(record.Action),
			ReasonCode: record.ReasonCode,
			Reason:     record.Reason,
			OperatorID: record.OperatorID,
			CreatedAt:  record.CreatedAt,
			UpdatedAt:  record.UpdatedAt,
		}
	})
	return &repo.ContentModerationRecordPageResp{
		Rows: rows,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *ContentModerationRecordRepo) getQuery(query *gen.ContentModerationRecordQuery, req *repo.ContentModerationRecordGetReq) *gen.ContentModerationRecordQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(contentmoderationrecord.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(contentmoderationrecord.IDIn(req.IDs...))
	}
	if req.Target != nil {
		query = query.Where(contentmoderationrecord.TargetEQ(contentmoderationrecord.Target(*req.Target)))
	}
	if len(req.Targets) > 0 {
		query = query.Where(contentmoderationrecord.TargetIn(lo.Map(req.Targets, func(item enum.ContentModerationTarget, _ int) contentmoderationrecord.Target {
			return contentmoderationrecord.Target(item)
		})...))
	}
	if req.TargetID != nil {
		query = query.Where(contentmoderationrecord.TargetIDEQ(*req.TargetID))
	}
	if len(req.TargetIDs) > 0 {
		query = query.Where(contentmoderationrecord.TargetIDIn(req.TargetIDs...))
	}
	if req.Action != nil {
		query = query.Where(contentmoderationrecord.ActionEQ(contentmoderationrecord.Action(*req.Action)))
	}
	if len(req.Actions) > 0 {
		query = query.Where(contentmoderationrecord.ActionIn(lo.Map(req.Actions, func(item enum.ContentModerationAction, _ int) contentmoderationrecord.Action {
			return contentmoderationrecord.Action(item)
		})...))
	}
	if req.OperatorID != nil {
		query = query.Where(contentmoderationrecord.OperatorIDEQ(*req.OperatorID))
	}
	return query
}
