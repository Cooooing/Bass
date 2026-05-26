package repo

import (
	"common/api/gen/common"
	"common/pkg/constant"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationmeta"
	"notify/internal/data/gen/notificationrecord"
	notifyenum "notify/internal/enum"
	"time"

	utilent "common/pkg/util/ent"
)

var _ repo.NotificationRecordRepo = (*NotificationRecordRepo)(nil)

type NotificationRecordRepo struct {
	db *gen.Client
}

func NewNotificationRecordRepo(db *gen.Client) repo.NotificationRecordRepo {
	return &NotificationRecordRepo{
		db: db,
	}
}

func (r *NotificationRecordRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationRecordRepo) Save(ctx context.Context, u *model.NotificationRecord) (*model.NotificationRecord, error) {
	save, err := r.getClient(ctx).NotificationRecord.Create().
		SetNotificationID(u.NotificationID).
		SetReceiverID(u.ReceiverID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.NotificationRecord{
		ID:             save.ID,
		NotificationID: save.NotificationID,
		ReceiverID:     save.ReceiverID,
		ReadTime:       save.ReadTime,
		CreatedAt:      save.CreatedAt,
		UpdatedAt:      save.UpdatedAt,
	}, nil
}

func (r *NotificationRecordRepo) Saves(ctx context.Context, u []*model.NotificationRecord) ([]*model.NotificationRecord, error) {
	client := r.getClient(ctx)
	records := make([]*gen.NotificationRecordCreate, 0, len(u))
	for _, v := range u {
		records = append(records,
			client.NotificationRecord.Create().
				SetNotificationID(v.NotificationID).
				SetReceiverID(v.ReceiverID),
		)
	}
	saves, err := client.NotificationRecord.CreateBulk(records...).Save(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*model.NotificationRecord, len(saves))
	for i := range saves {
		res[i] = &model.NotificationRecord{
			ID:             saves[i].ID,
			NotificationID: saves[i].NotificationID,
			ReceiverID:     saves[i].ReceiverID,
			ReadTime:       saves[i].ReadTime,
			CreatedAt:      saves[i].CreatedAt,
			UpdatedAt:      saves[i].UpdatedAt,
		}
	}
	return res, nil
}

func (r *NotificationRecordRepo) Read(ctx context.Context, receiverId int64, startTime *time.Time, endTime *time.Time, notificationRecordIds []int64) (int, error) {
	update := r.getClient(ctx).NotificationRecord.Update().Where(
		notificationrecord.ReceiverIDEQ(receiverId),
		notificationrecord.ReadTimeIsNil(),
	)
	if startTime != nil {
		update = update.Where(notificationrecord.CreatedAtGTE(*startTime))
	}
	if endTime != nil {
		update = update.Where(notificationrecord.CreatedAtLTE(*endTime))
	}
	if len(notificationRecordIds) > 0 {
		update = update.Where(notificationrecord.IDIn(notificationRecordIds...))
	}
	count, err := update.
		SetReadTime(time.Now()).
		Save(ctx)
	return count, err
}

func (r *NotificationRecordRepo) UnreadCount(ctx context.Context, receiverId int64) (int, error) {
	return r.getClient(ctx).NotificationRecord.Query().
		Where(
			notificationrecord.ReceiverIDEQ(receiverId),
			notificationrecord.ReadTimeIsNil(),
		).
		Count(ctx)
}

func (r *NotificationRecordRepo) Get(ctx context.Context, req *repo.NotificationRecordGetReq) (*model.NotificationRecord, error) {
	query := r.getClient(ctx).NotificationRecord.Query()
	query = r.getQuery(query, req)
	n, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	record := &model.NotificationRecord{
		ID:             n.ID,
		NotificationID: n.NotificationID,
		ReceiverID:     n.ReceiverID,
		ReadTime:       n.ReadTime,
		CreatedAt:      n.CreatedAt,
		UpdatedAt:      n.UpdatedAt,
	}
	if req.WithMeta && n.Edges.NotificationMeta != nil {
		record.NotificationMeta = &model.NotificationMeta{
			ID:        n.Edges.NotificationMeta.ID,
			Title:     n.Edges.NotificationMeta.Title,
			Content:   n.Edges.NotificationMeta.Content,
			Status:    notifyenum.NotificationStatus(n.Edges.NotificationMeta.Status),
			CreatedAt: n.Edges.NotificationMeta.CreatedAt,
			UpdatedAt: n.Edges.NotificationMeta.UpdatedAt,
		}
	}
	return record, err
}

func (r *NotificationRecordRepo) GetList(ctx context.Context, req *repo.NotificationRecordGetReq) ([]*model.NotificationRecord, error) {
	var (
		records []*model.NotificationRecord
		err     error
	)
	query := r.getClient(ctx).NotificationRecord.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		record := &model.NotificationRecord{
			ID:             list[i].ID,
			NotificationID: list[i].NotificationID,
			ReceiverID:     list[i].ReceiverID,
			ReadTime:       list[i].ReadTime,
			CreatedAt:      list[i].CreatedAt,
			UpdatedAt:      list[i].UpdatedAt,
		}
		if req.WithMeta && list[i].Edges.NotificationMeta != nil {
			record.NotificationMeta = &model.NotificationMeta{
				ID:        list[i].Edges.NotificationMeta.ID,
				Title:     list[i].Edges.NotificationMeta.Title,
				Content:   list[i].Edges.NotificationMeta.Content,
				Status:    notifyenum.NotificationStatus(list[i].Edges.NotificationMeta.Status),
				CreatedAt: list[i].Edges.NotificationMeta.CreatedAt,
				UpdatedAt: list[i].Edges.NotificationMeta.UpdatedAt,
			}
		}
		records = append(records, record)
	}
	return records, nil
}

func (r *NotificationRecordRepo) GetPage(ctx context.Context, page *common.PageRequest, req *repo.NotificationRecordGetReq) ([]*model.NotificationRecord, *common.PageReply, error) {
	var (
		notificationRecords []*model.NotificationRecord
		err                 error
	)
	page = constant.PageValid(page)
	query := r.getClient(ctx).NotificationRecord.Query()
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
		record := &model.NotificationRecord{
			ID:             item.ID,
			NotificationID: item.NotificationID,
			ReceiverID:     item.ReceiverID,
			ReadTime:       item.ReadTime,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
		}
		if req.WithMeta && item.Edges.NotificationMeta != nil {
			record.NotificationMeta = &model.NotificationMeta{
				ID:        item.Edges.NotificationMeta.ID,
				Title:     item.Edges.NotificationMeta.Title,
				Content:   item.Edges.NotificationMeta.Content,
				Status:    notifyenum.NotificationStatus(item.Edges.NotificationMeta.Status),
				CreatedAt: item.Edges.NotificationMeta.CreatedAt,
				UpdatedAt: item.Edges.NotificationMeta.UpdatedAt,
			}
		}
		notificationRecords = append(notificationRecords, record)
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
	if req.ReceiverId != nil {
		query = query.Where(notificationrecord.ReceiverIDEQ(*req.ReceiverId))
	}
	if req.Status != nil {
		dbStatus, _ := notifyenum.NotificationStatusMap.ToEnum(*req.Status)
		query = query.Where(notificationrecord.HasNotificationMetaWith(notificationmeta.StatusEQ(notificationmeta.Status(dbStatus))))
	}
	if req.WithMeta {
		query = query.WithNotificationMeta()
	}
	return query
}
