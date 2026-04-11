package repo

import (
	cv1 "common/api/gen/common/v1"
	v1 "common/api/gen/content/v1"
	"common/pkg/constant"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	basedata "content/internal/data/base"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/articleactionrecord"
	"context"
)

type ArticleActionRecordRepo struct {
	*basedata.BaseData
	client *gen.Client
}

func NewArticleActionRecordRepo(BaseData *basedata.BaseData, client *gen.Client) repo.ArticleActionRecordRepo {
	return &ArticleActionRecordRepo{
		BaseData: BaseData,
		client:   client,
	}
}

func (r *ArticleActionRecordRepo) Save(ctx context.Context, tx *gen.Client, record *model.ArticleActionRecord) (*model.ArticleActionRecord, error) {
	save, err := tx.ArticleActionRecord.Create().
		SetArticleID(record.ArticleID).
		SetUserID(record.UserID).
		SetType(record.Type).
		Save(ctx)
	return &model.ArticleActionRecord{ArticleActionRecord: save}, err
}

func (r *ArticleActionRecordRepo) Delete(ctx context.Context, tx *gen.Client, articleId int64, userId int64, action v1.ArticleAction) error {
	_, err := tx.ArticleActionRecord.Delete().
		Where(articleactionrecord.ArticleIDEQ(articleId)).
		Where(articleactionrecord.UserIDEQ(userId)).
		Where(articleactionrecord.TypeEQ(int32(action))).
		Exec(ctx)
	return err
}

func (r *ArticleActionRecordRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.ArticleActionRecordReq) (*model.ArticleActionRecord, error) {
	query := tx.ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	c, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cv1.ErrorBadRequest("article action record is not found")
	}
	return &model.ArticleActionRecord{ArticleActionRecord: c}, err
}

func (r *ArticleActionRecordRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.ArticleActionRecordReq) ([]*model.ArticleActionRecord, error) {
	var (
		articleActionRecords []*model.ArticleActionRecord
		err                  error
	)
	query := tx.ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		articleActionRecords = append(articleActionRecords, &model.ArticleActionRecord{ArticleActionRecord: list[i]})
	}
	return articleActionRecords, nil
}

func (r *ArticleActionRecordRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.ArticleActionRecordReq) ([]*model.ArticleActionRecord, *cv1.PageReply, error) {
	var (
		articleActionRecords []*model.ArticleActionRecord
		err                  error
		total                int
	)
	page = constant.PageValid(page)
	query := tx.ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	total, err = countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	for i := range list {
		articleActionRecords = append(articleActionRecords, &model.ArticleActionRecord{ArticleActionRecord: list[i]})
	}
	return articleActionRecords, &cv1.PageReply{
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
		query = query.Where(articleactionrecord.TypeEQ(int32(*req.Type)))
	}
	return query
}
