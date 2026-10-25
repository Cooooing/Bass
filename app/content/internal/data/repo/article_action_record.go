package repo

import (
	commonmodel "common/pkg/model"
	"context"

	utilent "common/pkg/util/ent"

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

func NewArticleActionRecordRepo(
	db *gen.Client,
) repo.ArticleActionRecordRepo {
	return &ArticleActionRecordRepo{
		db: db,
	}
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
		SetNillableClientRequestID(record.ClientRequestID).
		SetNillablePoints(record.Points).
		SetStatus(articleactionrecord.Status(record.Status)).
		Save(ctx)
	if gen.IsConstraintError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ArticleActionRecordRepo) TransitionStatus(ctx context.Context, clientRequestID string, from enum.ContentActionStatus, to enum.ContentActionStatus) (bool, error) {
	count, err := r.getClient(ctx).ArticleActionRecord.Update().
		Where(articleactionrecord.ClientRequestIDEQ(clientRequestID), articleactionrecord.StatusEQ(articleactionrecord.Status(from))).
		SetStatus(articleactionrecord.Status(to)).
		Save(ctx)
	return count == 1, err
}

func (r *ArticleActionRecordRepo) Delete(ctx context.Context, req *repo.ArticleActionRecordDeleteReq) (int, error) {
	articleID := req.ArticleID
	userID := req.UserID
	action := req.Action
	deleted, err := r.getClient(ctx).ArticleActionRecord.Delete().
		Where(articleactionrecord.ArticleIDEQ(articleID)).
		Where(articleactionrecord.UserIDEQ(userID)).
		Where(articleactionrecord.TypeEQ(articleactionrecord.Type(action))).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (r *ArticleActionRecordRepo) Exist(ctx context.Context, req *repo.ArticleActionRecordReq) (bool, error) {
	query := r.getClient(ctx).ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	exist, err := query.Exist(ctx)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (r *ArticleActionRecordRepo) Get(ctx context.Context, req *repo.ArticleActionRecordReq) (*model.ArticleActionRecord, error) {
	query := r.getClient(ctx).ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	c, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.ArticleActionRecord{
		ID:        c.ID,
		ArticleID: c.ArticleID,
		UserID:    c.UserID,
		Type:      enum.ArticleAction(c.Type), ClientRequestID: c.ClientRequestID, Points: c.Points, Status: enum.ContentActionStatus(c.Status),
	}, nil
}

func (r *ArticleActionRecordRepo) List(ctx context.Context, req *repo.ArticleActionRecordReq) ([]*model.ArticleActionRecord, error) {
	query := r.getClient(ctx).ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	rows := make([]*model.ArticleActionRecord, 0, len(list))
	for i := range list {
		rows = append(rows, &model.ArticleActionRecord{
			ID:        list[i].ID,
			ArticleID: list[i].ArticleID,
			UserID:    list[i].UserID,
			Type:      enum.ArticleAction(list[i].Type), ClientRequestID: list[i].ClientRequestID, Points: list[i].Points, Status: enum.ContentActionStatus(list[i].Status),
		})
	}
	return rows, nil
}

func (r *ArticleActionRecordRepo) Map(ctx context.Context, req *repo.ArticleActionRecordReq) (map[int64]*model.
	ArticleActionRecord, error) {
	listResp, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(listResp, func(item *model.ArticleActionRecord) (int64, *model.ArticleActionRecord) {
		return item.ID, item
	}), nil
}

func (r *ArticleActionRecordRepo) Count(ctx context.Context, req *repo.ArticleActionRecordReq) (int, error) {
	query := r.getClient(ctx).ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ArticleActionRecordRepo) Page(ctx context.Context, req *repo.ArticleActionRecordReq) (*repo.ArticleActionRecordPageResp, error) {
	page := commonmodel.NormalizePage(req.Page)
	query := r.getClient(ctx).ArticleActionRecord.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	total, err := countQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(page.Limit()).Offset(page.Offset()).All(ctx)
	if err != nil {
		return nil, err
	}
	rows := make([]*model.ArticleActionRecord, 0, len(list))
	for i := range list {
		rows = append(rows, &model.ArticleActionRecord{
			ID:        list[i].ID,
			ArticleID: list[i].ArticleID,
			UserID:    list[i].UserID,
			Type:      enum.ArticleAction(list[i].Type), ClientRequestID: list[i].ClientRequestID, Points: list[i].Points, Status: enum.ContentActionStatus(list[i].Status),
		})
	}
	return &repo.ArticleActionRecordPageResp{
		Rows: rows,
		Page: &commonmodel.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *ArticleActionRecordRepo) getQuery(query *gen.ArticleActionRecordQuery, req *repo.ArticleActionRecordReq) *gen.ArticleActionRecordQuery {
	if req.ID != nil {
		query = query.Where(articleactionrecord.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(articleactionrecord.IDIn(req.IDs...))
	}
	if req.ArticleID != nil {
		query = query.Where(articleactionrecord.ArticleIDEQ(*req.ArticleID))
	}
	if len(req.ArticleIDs) > 0 {
		query = query.Where(articleactionrecord.ArticleIDIn(req.ArticleIDs...))
	}
	if req.UserID != nil {
		query = query.Where(articleactionrecord.UserIDEQ(*req.UserID))
	}
	if len(req.UserIDs) > 0 {
		query = query.Where(articleactionrecord.UserIDIn(req.UserIDs...))
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
	if req.ClientRequestID != nil {
		query = query.Where(articleactionrecord.ClientRequestIDEQ(*req.ClientRequestID))
	}
	if req.Status != nil {
		query = query.Where(articleactionrecord.StatusEQ(articleactionrecord.Status(*req.Status)))
	}
	return query
}
