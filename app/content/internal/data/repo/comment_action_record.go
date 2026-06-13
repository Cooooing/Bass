package repo

import (
	"common/proto/gen/common"
	"context"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/commentactionrecord"
	"content/internal/enum"

	"github.com/samber/lo"
)

var _ repo.CommentActionRecordRepo = (*CommentActionRecordRepo)(nil)

type CommentActionRecordRepo struct {
	db *gen.Client
}

func NewCommentActionRecordRepo(db *gen.Client) repo.CommentActionRecordRepo {
	return &CommentActionRecordRepo{db: db}
}

func (r *CommentActionRecordRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *CommentActionRecordRepo) Save(ctx context.Context, record *model.CommentActionRecord) (bool, error) {
	_, err := r.getClient(ctx).CommentActionRecord.Create().
		SetCommentID(record.CommentID).
		SetUserID(record.UserID).
		SetType(commentactionrecord.Type(record.Type)).
		Save(ctx)
	if gen.IsConstraintError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *CommentActionRecordRepo) Delete(ctx context.Context, commentId int64, userId int64, action enum.CommentAction) (int, error) {
	return r.getClient(ctx).CommentActionRecord.Delete().
		Where(commentactionrecord.CommentIDEQ(commentId)).
		Where(commentactionrecord.UserIDEQ(userId)).
		Where(commentactionrecord.TypeEQ(commentactionrecord.Type(action))).
		Exec(ctx)
}

func (r *CommentActionRecordRepo) Exist(ctx context.Context, req *repo.CommentActionRecordReq) (bool, error) {
	query := r.getClient(ctx).CommentActionRecord.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *CommentActionRecordRepo) Get(ctx context.Context, req *repo.CommentActionRecordReq) (*model.CommentActionRecord, error) {
	query := r.getClient(ctx).CommentActionRecord.Query()
	query = r.getQuery(query, req)
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.CommentActionRecord{
		ID:        row.ID,
		CommentID: row.CommentID,
		UserID:    row.UserID,
		Type:      enum.CommentAction(row.Type),
	}, nil
}

func (r *CommentActionRecordRepo) List(ctx context.Context, req *repo.CommentActionRecordReq) ([]*model.CommentActionRecord, error) {
	query := r.getClient(ctx).CommentActionRecord.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.CommentActionRecord, 0, len(list))
	for _, row := range list {
		result = append(result, &model.CommentActionRecord{
			ID:        row.ID,
			CommentID: row.CommentID,
			UserID:    row.UserID,
			Type:      enum.CommentAction(row.Type),
		})
	}
	return result, nil
}

func (r *CommentActionRecordRepo) Map(ctx context.Context, req *repo.CommentActionRecordReq) (map[int64]*model.CommentActionRecord, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(list, func(item *model.CommentActionRecord) (int64, *model.CommentActionRecord) {
		return item.ID, item
	}), nil
}

func (r *CommentActionRecordRepo) Count(ctx context.Context, req *repo.CommentActionRecordReq) (int, error) {
	query := r.getClient(ctx).CommentActionRecord.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *CommentActionRecordRepo) Page(ctx context.Context, page *common.PageRequest, req *repo.CommentActionRecordReq) ([]*model.CommentActionRecord, *common.PageReply, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).CommentActionRecord.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*model.CommentActionRecord, 0, len(list))
	for _, row := range list {
		result = append(result, &model.CommentActionRecord{
			ID:        row.ID,
			CommentID: row.CommentID,
			UserID:    row.UserID,
			Type:      enum.CommentAction(row.Type),
		})
	}
	return result, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *CommentActionRecordRepo) getQuery(query *gen.CommentActionRecordQuery, req *repo.CommentActionRecordReq) *gen.CommentActionRecordQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(commentactionrecord.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(commentactionrecord.IDIn(req.IDs...))
	}
	if req.CommentId != nil {
		query = query.Where(commentactionrecord.CommentIDEQ(*req.CommentId))
	}
	if len(req.CommentIds) > 0 {
		query = query.Where(commentactionrecord.CommentIDIn(req.CommentIds...))
	}
	if req.UserId != nil {
		query = query.Where(commentactionrecord.UserIDEQ(*req.UserId))
	}
	if len(req.UserIds) > 0 {
		query = query.Where(commentactionrecord.UserIDIn(req.UserIds...))
	}
	if req.Type != nil {
		query = query.Where(commentactionrecord.TypeEQ(commentactionrecord.Type(*req.Type)))
	}
	if len(req.Types) > 0 {
		dbTypes := lo.Map(req.Types, func(item enum.CommentAction, _ int) commentactionrecord.Type {
			return commentactionrecord.Type(item)
		})
		query = query.Where(commentactionrecord.TypeIn(dbTypes...))
	}
	return query
}
