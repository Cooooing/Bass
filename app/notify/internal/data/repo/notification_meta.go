package repo

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	"common/pkg/constant"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationmeta"
	notifyenum "notify/internal/enum"

	utilent "common/pkg/util/ent"
)

var _ repo.NotificationMetaRepo = (*NotificationMetaRepo)(nil)

type NotificationMetaRepo struct {
	db *gen.Client
}

func NewNotificationMetaRepo(db *gen.Client) repo.NotificationMetaRepo {
	return &NotificationMetaRepo{
		db: db,
	}
}

func (r *NotificationMetaRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationMetaRepo) Save(ctx context.Context, u *model.NotificationMeta) (*model.NotificationMeta, error) {
	save, err := r.getClient(ctx).NotificationMeta.Create().
		SetTitle(u.Title).
		SetContent(u.Content).
		SetStatus(notificationmeta.Status(u.Status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.NotificationMeta{
		ID:        save.ID,
		Title:     save.Title,
		Content:   save.Content,
		Status:    notifyenum.NotificationStatus(save.Status),
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}, nil
}

func (r *NotificationMetaRepo) Get(ctx context.Context, req *repo.NotificationMetaGetReq) (*model.NotificationMeta, error) {
	query := r.getClient(ctx).NotificationMeta.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cerrors.ErrorBadRequest("notification meta is not found")
	}
	return &model.NotificationMeta{
		ID:        t.ID,
		Title:     t.Title,
		Content:   t.Content,
		Status:    notifyenum.NotificationStatus(t.Status),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, err
}

func (r *NotificationMetaRepo) GetList(ctx context.Context, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, error) {
	var (
		records []*model.NotificationMeta
		err     error
	)
	query := r.getClient(ctx).NotificationMeta.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		records = append(records, &model.NotificationMeta{
			ID:        item.ID,
			Title:     item.Title,
			Content:   item.Content,
			Status:    notifyenum.NotificationStatus(item.Status),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	return records, nil
}

func (r *NotificationMetaRepo) GetPage(ctx context.Context, page *common.PageRequest, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, *common.PageReply, error) {
	var (
		notificationMetas []*model.NotificationMeta
		err               error
	)
	page = constant.PageValid(page)
	query := r.getClient(ctx).NotificationMeta.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	for _, item := range list {
		notificationMetas = append(notificationMetas, &model.NotificationMeta{
			ID:        item.ID,
			Title:     item.Title,
			Content:   item.Content,
			Status:    notifyenum.NotificationStatus(item.Status),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	return notificationMetas, &common.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *NotificationMetaRepo) getQuery(query *gen.NotificationMetaQuery, req *repo.NotificationMetaGetReq) *gen.NotificationMetaQuery {
	if req.NotificationMetaId != nil {
		query = query.Where(notificationmeta.IDEQ(*req.NotificationMetaId))
	}
	if len(req.NotificationMetaIds) > 0 {
		query = query.Where(notificationmeta.IDIn(req.NotificationMetaIds...))
	}
	if req.Status != nil {
		dbStatus, _ := notifyenum.NotificationStatusMap.ToEnum(*req.Status)
		query = query.Where(notificationmeta.StatusEQ(notificationmeta.Status(dbStatus)))
	}
	return query
}
