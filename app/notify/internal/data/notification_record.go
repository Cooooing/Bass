package data

import (
	cv1 "common/api/common/v1"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent/gen"
	"notify/internal/data/ent/gen/notificationmeta"
	"notify/internal/data/ent/gen/notificationrecord"
)

type NotificationRecordRepo struct {
	*BaseRepo
}

func NewNotificationRecordRepo(repo *BaseRepo) repo.NotificationRecordRepo {
	return &NotificationRecordRepo{
		BaseRepo: repo,
	}
}

func (r *NotificationRecordRepo) Save(ctx context.Context, tx *gen.Client, u *model.NotificationRecord) (*model.NotificationRecord, error) {
	save, err := tx.NotificationRecord.Create().
		SetNotificationID(u.NotificationID).
		SetReceiverID(u.ReceiverID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.NotificationRecord{NotificationRecord: save}, nil
}

func (r *NotificationRecordRepo) Saves(ctx context.Context, tx *gen.Client, u []*model.NotificationRecord) ([]*model.NotificationRecord, error) {
	records := make([]*gen.NotificationRecordCreate, 0, len(u))
	for _, v := range u {
		records = append(records,
			tx.NotificationRecord.Create().
				SetNotificationID(v.NotificationID).
				SetReceiverID(v.ReceiverID),
		)
	}
	saves, err := tx.NotificationRecord.CreateBulk(records...).Save(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*model.NotificationRecord, len(saves))
	for i := range saves {
		res[i] = &model.NotificationRecord{NotificationRecord: saves[i]}
	}
	return res, nil
}

func (r *NotificationRecordRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.NotificationRecordGetReq) (*model.NotificationRecord, error) {
	query := tx.NotificationRecord.Query()
	query = r.getQuery(query, req)
	n, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	return &model.NotificationRecord{NotificationRecord: n}, err
}

func (r *NotificationRecordRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.NotificationRecordGetReq) ([]*model.NotificationRecord, error) {
	var (
		records []*model.NotificationRecord
		err     error
	)
	query := tx.NotificationRecord.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		records = append(records, &model.NotificationRecord{NotificationRecord: list[i]})
	}
	return records, nil
}

func (r *NotificationRecordRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.NotificationRecordGetReq) ([]*model.NotificationRecord, *cv1.PageReply, error) {
	// TODO implement me
	panic("implement me")
}

func (r *NotificationRecordRepo) getQuery(query *gen.NotificationRecordQuery, req *repo.NotificationRecordGetReq) *gen.NotificationRecordQuery {
	if req.NotificationMetaId != nil {
		query = query.Where(notificationrecord.NotificationIDEQ(*req.NotificationMetaId))
	}
	if req.NotificationMetaIds != nil {
		query = query.Where(notificationrecord.NotificationIDIn(req.NotificationMetaIds...))
	}
	if req.NotificationRecordId != nil {
		query = query.Where(notificationrecord.IDEQ(*req.NotificationRecordId))
	}
	if req.NotificationRecordIds != nil {
		query = query.Where(notificationrecord.IDIn(req.NotificationRecordIds...))
	}
	if req.ReceiverId != nil {
		query = query.Where(notificationrecord.ReceiverIDEQ(*req.ReceiverId))
	}
	if req.WithMeta {
		query = query.WithNotificationMeta(func(query *gen.NotificationMetaQuery) {
			if req.NotificationType != nil {
				query = query.Where(notificationmeta.NotificationTypeEQ(int32(*req.NotificationType)))
			}
			if req.Status != nil {
				query = query.Where(notificationmeta.StatusEQ(int32(*req.Status)))
			}
			if req.SenderId != nil {
				query = query.Where(notificationmeta.SenderIDEQ(*req.SenderId))
			}
		})
	}
	return query
}
