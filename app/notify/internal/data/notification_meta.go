package data

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent/gen"
	"notify/internal/data/ent/gen/notificationmeta"
)

type NotificationMetaRepo struct {
	*BaseRepo
}

func NewNotificationMetaRepo(repo *BaseRepo) repo.NotificationMetaRepo {
	return &NotificationMetaRepo{
		BaseRepo: repo,
	}
}

func (r *NotificationMetaRepo) Save(ctx context.Context, tx *gen.Client, u *model.NotificationMeta) (*model.NotificationMeta, error) {
	save, err := tx.NotificationMeta.Create().
		SetUUID(u.UUID).
		SetNotificationType(u.NotificationType).
		SetSenderID(u.SenderID).
		SetMeta(u.Meta).
		SetContent(u.Content).
		SetStatus(u.Status).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.NotificationMeta{NotificationMeta: save}, nil
}

func (r *NotificationMetaRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.NotificationMetaGetReq) (*model.NotificationMeta, error) {
	query := tx.NotificationMeta.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cv1.ErrorBadRequest("notification meta is not found")
	}
	return &model.NotificationMeta{NotificationMeta: t}, err
}

func (r *NotificationMetaRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, error) {
	var (
		records []*model.NotificationMeta
		err     error
	)
	query := tx.NotificationMeta.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		records = append(records, &model.NotificationMeta{NotificationMeta: item})
	}
	return records, nil
}

func (r *NotificationMetaRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, *cv1.PageReply, error) {
	var (
		notificationMetas []*model.NotificationMeta
		err               error
	)
	page = constant.PageValid(page)
	query := tx.NotificationMeta.Query()
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
		notificationMetas = append(notificationMetas, &model.NotificationMeta{NotificationMeta: item})
	}
	return notificationMetas, &cv1.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *NotificationMetaRepo) getQuery(query *gen.NotificationMetaQuery, req *repo.NotificationMetaGetReq) *gen.NotificationMetaQuery {
	if req.NotificationMetaId != nil {
		query = query.Where(notificationmeta.IDEQ(*req.NotificationMetaId))
	}
	if len(req.NotificationMeraIds) > 0 {
		query = query.Where(notificationmeta.IDIn(req.NotificationMeraIds...))
	}
	if req.UUID != nil {
		query = query.Where(notificationmeta.UUIDEQ(*req.UUID))
	}
	if len(req.UUIDs) > 0 {
		query = query.Where(notificationmeta.UUIDIn(req.UUIDs...))
	}
	if req.NotificationMetaType != nil {
		query = query.Where(notificationmeta.NotificationTypeEQ(int32(*req.NotificationMetaType)))
	}
	if req.SenderId != nil {
		query = query.Where(notificationmeta.SenderIDEQ(*req.SenderId))
	}
	if len(req.SenderIds) > 0 {
		query = query.Where(notificationmeta.SenderIDIn(req.SenderIds...))
	}
	if req.Status != nil {
		query = query.Where(notificationmeta.StatusEQ(int32(*req.Status)))
	}
	return query
}
