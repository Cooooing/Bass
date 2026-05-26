package repo

import (
	"context"

	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/articleactionrecord"
	"content/internal/enum"
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

func (r *ArticleActionRecordRepo) Save(ctx context.Context, record *model.ArticleActionRecord) (*model.ArticleActionRecord, error) {
	save, err := r.getClient(ctx).ArticleActionRecord.Create().
		SetArticleID(record.ArticleID).
		SetUserID(record.UserID).
		SetType(articleactionrecord.Type(record.Type)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ArticleActionRecord{
		ID:        save.ID,
		ArticleID: save.ArticleID,
		UserID:    save.UserID,
		Type:      enum.ArticleAction(save.Type),
	}, nil
}

func (r *ArticleActionRecordRepo) Delete(ctx context.Context, articleId int64, userId int64, action v1.ArticleAction) (int, error) {
	dbType, _ := enum.ArticleActionMap.ToEnum(action)
	return r.getClient(ctx).ArticleActionRecord.Delete().
		Where(articleactionrecord.ArticleIDEQ(articleId)).
		Where(articleactionrecord.UserIDEQ(userId)).
		Where(articleactionrecord.TypeEQ(articleactionrecord.Type(dbType))).
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
		return nil, cerrors.ErrorBadRequest("article action record is not found")
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

func (r *ArticleActionRecordRepo) GetList(ctx context.Context, req *repo.ArticleActionRecordReq) ([]*model.ArticleActionRecord, error) {
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

func (r *ArticleActionRecordRepo) GetPage(ctx context.Context, page *common.PageRequest, req *repo.ArticleActionRecordReq) ([]*model.ArticleActionRecord, *common.PageReply, error) {
	page = constant.PageValid(page)
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
	if req.ArticleId != nil {
		query = query.Where(articleactionrecord.ArticleIDEQ(*req.ArticleId))
	}
	if req.UserId != nil {
		query = query.Where(articleactionrecord.UserIDEQ(*req.UserId))
	}
	if req.Type != nil {
		dbType, _ := enum.ArticleActionMap.ToEnum(*req.Type)
		query = query.Where(articleactionrecord.TypeEQ(articleactionrecord.Type(dbType)))
	}
	return query
}
