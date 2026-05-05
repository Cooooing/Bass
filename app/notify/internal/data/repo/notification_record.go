package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/notify/v1"
	"common/pkg/constant"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	database "notify/internal/data/base"
	"notify/internal/data/ent/gen"
	"notify/internal/data/ent/gen/notificationmeta"
	"notify/internal/data/ent/gen/notificationrecord"
	"time"
)

type NotificationRecordRepo struct {
	*database.BaseData
}

func NewNotificationRecordRepo(repo *database.BaseData) repo.NotificationRecordRepo {
	return &NotificationRecordRepo{
		BaseData: repo,
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

func (r *NotificationRecordRepo) Read(ctx context.Context, tx *gen.Client, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error) {
	update := tx.NotificationRecord.Update().Where(notificationrecord.ReceiverIDEQ(receiverId))
	if startTime != nil {
		update = update.Where(notificationrecord.ReadTimeGTE(*startTime))
	}
	if endTime != nil {
		update = update.Where(notificationrecord.ReadTimeLTE(*endTime))
	}
	if len(notificationRecordIds) > 0 {
		update = update.Where(notificationrecord.IDIn(notificationRecordIds...))
	}
	count, err := update.
		SetReadTime(time.Now()).
		Save(ctx)
	return count, err
}

func (r *NotificationRecordRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.NotificationRecordGetReq) (*model.NotificationRecord, error) {
	query := tx.NotificationRecord.Query()
	query = r.getQuery(query, req)
	n, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	return &model.NotificationRecord{NotificationRecord: n, WithMeta: req.WithMeta}, err
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
		records = append(records, &model.NotificationRecord{NotificationRecord: list[i], WithMeta: req.WithMeta})
	}
	return records, nil
}

func (r *NotificationRecordRepo) GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *repo.NotificationRecordGetReq) ([]*model.NotificationRecord, *common.PageReply, error) {
	var (
		notificationRecords []*model.NotificationRecord
		err                 error
	)
	page = constant.PageValid(page)
	query := tx.NotificationRecord.Query()
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
		notificationRecords = append(notificationRecords, &model.NotificationRecord{NotificationRecord: item, WithMeta: req.WithMeta})
	}
	return notificationRecords, &common.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *NotificationRecordRepo) getQuery(query *gen.NotificationRecordQuery, req *repo.NotificationRecordGetReq) *gen.NotificationRecordQuery {
	if req.NotificationRecordId != nil {
		query = query.Where(notificationrecord.IDEQ(*req.NotificationRecordId))
	}
	if len(req.NotificationRecordIds) > 0 {
		query = query.Where(notificationrecord.IDIn(req.NotificationRecordIds...))
	}
	if req.NotificationMetaId != nil {
		query = query.Where(notificationrecord.NotificationIDEQ(*req.NotificationMetaId))
	}
	if len(req.NotificationMetaIds) > 0 {
		query = query.Where(notificationrecord.NotificationIDIn(req.NotificationMetaIds...))
	}
	if req.SenderId != nil {
		query = query.Where(notificationrecord.HasNotificationMetaWith(notificationmeta.SenderIDEQ(*req.SenderId)))
	}
	if req.ReceiverId != nil {
		query = query.Where(notificationrecord.ReceiverIDEQ(*req.ReceiverId))
	}
	if req.Status != nil {
		query = query.Where(notificationrecord.HasNotificationMetaWith(notificationmeta.StatusEQ(v1.NotificationStatus_name[int32(*req.Status)])))
	}
	if req.NotificationType != nil {
		query = query.Where(notificationrecord.HasNotificationMetaWith(notificationmeta.NotificationTypeEQ(v1.NotificationType_name[int32(*req.NotificationType)])))
	}
	if req.WithMeta {
		query = query.WithNotificationMeta()
	}
	return query
}
