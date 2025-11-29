package data

import (
	cv1 "common/api/common/v1"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent/gen"
)

type NotificationMetaRepo struct {
	*BaseRepo
}

func NewNotificationMetaRepo(repo *BaseRepo) repo.NotificationMetaRepo {
	return &NotificationMetaRepo{
		BaseRepo: repo,
	}
}

func (n *NotificationMetaRepo) Save(ctx context.Context, tx *gen.Client, u *model.NotificationMeta) (*model.NotificationMeta, error) {
	save, err := tx.NotificationMeta.Create().
		SetUUID(u.UUID).
		SetNotificationType(u.NotificationType).
		SetSenderID(u.SenderID).
		SetMeta(nil).
		SetContent("").
		SetStatus(0).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.NotificationMeta{NotificationMeta: save}, nil
}

func (r *NotificationMetaRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.NotificationMetaGetReq) (*model.NotificationMeta, error) {
	// TODO implement me
	panic("implement me")
}

func (r *NotificationMetaRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, error) {
	// TODO implement me
	panic("implement me")
}

func (r *NotificationMetaRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, *cv1.PageReply, error) {
	// TODO implement me
	panic("implement me")
}
