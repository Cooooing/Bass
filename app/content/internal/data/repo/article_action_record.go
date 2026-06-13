package repo

import (
	cerrors "common/proto/gen/common/errors"
	"context"

	"common/pkg/apperror"
	"common/pkg/server"
	utilent "common/pkg/util/ent"
	"common/proto/gen/common"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/articleactionrecord"
	"content/internal/enum"

	"github.com/samber/lo"
)

var _ repo.ArticleActionRecordRepo = (*ArticleActionRecordRepo)(nil)

type ArticleActionRecordRepo struct {
	db *gen.Client
}

func NewArticleActionRecordRepo(db *gen.Client) repo.ArticleActionRecordRepo {
	return &ArticleActionRecordRepo{db: db}
}

func (r *ArticleActionRecordRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ArticleActionRecordRepo) Save(ctx context.Context, record *model.ArticleActionRecord) (bool, error) {
	_, err := r.getClient(ctx).ArticleActionRecord.Create().
		SetArticleID(record.ArticleID).
		SetUserID(record.UserID).
		SetType(articleactionrecord.Type(record.Type)).
		Save(ctx)
	if gen.IsConstraintError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ArticleActionRecordRepo) Delete(ctx context.Context, articleId int64, userId int64, action enum.ArticleAction) (int, error) {
	return r.getClient(ctx).ArticleActionRecord.Delete().
		Where(articleactionrecord.ArticleIDEQ(articleId)).
		Where(articleactionrecord.UserIDEQ(userId)).
		Where(articleactionrecord.TypeEQ(articleactionrecord.Type(action))).
		Exec(ctx)
}

func (r *ArticleActionRecordRepo) Exist(ctx context.Context, req *repo.ArticleActionRecordReq) (bool, error) {
	query := r.getClient(ctx).ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *ArticleActionRecordRepo) Get(ctx context.Context, req *repo.ArticleActionRecordReq) (*model.ArticleActionRecord, error) {
	query := r.getClient(ctx).ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	c, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_ACTION_RECORD_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.ArticleActionRecord{
		ID:        c.ID,
		ArticleID: c.ArticleID,
		UserID:    c.UserID,
		Type:      enum.ArticleAction(c.Type),
	}, nil
}

func (r *ArticleActionRecordRepo) List(ctx context.Context, req *repo.ArticleActionRecordReq) ([]*model.ArticleActionRecord, error) {
	query := r.getClient(ctx).ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	articleActionRecords := make([]*model.ArticleActionRecord, 0, len(list))
	for i := range list {
		articleActionRecords = append(articleActionRecords, &model.ArticleActionRecord{
			ID:        list[i].ID,
			ArticleID: list[i].ArticleID,
			UserID:    list[i].UserID,
			Type:      enum.ArticleAction(list[i].Type),
		})
	}
	return articleActionRecords, nil
}

func (r *ArticleActionRecordRepo) Map(ctx context.Context, req *repo.ArticleActionRecordReq) (map[int64]*model.ArticleActionRecord, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(list, func(item *model.ArticleActionRecord) (int64, *model.ArticleActionRecord) {
		return item.ID, item
	}), nil
}

func (r *ArticleActionRecordRepo) Count(ctx context.Context, req *repo.ArticleActionRecordReq) (int, error) {
	query := r.getClient(ctx).ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *ArticleActionRecordRepo) Page(ctx context.Context, page *common.PageRequest, req *repo.ArticleActionRecordReq) ([]*model.ArticleActionRecord, *common.PageReply, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	total, err := countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	articleActionRecords := make([]*model.ArticleActionRecord, 0, len(list))
	for i := range list {
		articleActionRecords = append(articleActionRecords, &model.ArticleActionRecord{
			ID:        list[i].ID,
			ArticleID: list[i].ArticleID,
			UserID:    list[i].UserID,
			Type:      enum.ArticleAction(list[i].Type),
		})
	}
	return articleActionRecords, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *ArticleActionRecordRepo) getQuery(query *gen.ArticleActionRecordQuery, req *repo.ArticleActionRecordReq) *gen.ArticleActionRecordQuery {
	if req.ID != nil {
		query = query.Where(articleactionrecord.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(articleactionrecord.IDIn(req.IDs...))
	}
	if req.ArticleId != nil {
		query = query.Where(articleactionrecord.ArticleIDEQ(*req.ArticleId))
	}
	if len(req.ArticleIds) > 0 {
		query = query.Where(articleactionrecord.ArticleIDIn(req.ArticleIds...))
	}
	if req.UserId != nil {
		query = query.Where(articleactionrecord.UserIDEQ(*req.UserId))
	}
	if len(req.UserIds) > 0 {
		query = query.Where(articleactionrecord.UserIDIn(req.UserIds...))
	}
	if req.Type != nil {
		query = query.Where(articleactionrecord.TypeEQ(articleactionrecord.Type(*req.Type)))
	}
	if len(req.Types) > 0 {
		dbTypes := lo.Map(req.Types, func(item enum.ArticleAction, _ int) articleactionrecord.Type {
			return articleactionrecord.Type(item)
		})
		query = query.Where(articleactionrecord.TypeIn(dbTypes...))
	}
	return query
}
