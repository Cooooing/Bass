package repo

import (
	"common/api/gen/common"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationstationmessage"
	notifyenum "notify/internal/enum"
	"time"

	utilent "common/pkg/util/ent"
)

var _ bizrepo.NotificationStationMessageRepo = (*NotificationStationMessageRepo)(nil)

type NotificationStationMessageRepo struct {
	db *gen.Client
}

func NewNotificationStationMessageRepo(db *gen.Client) bizrepo.NotificationStationMessageRepo {
	return &NotificationStationMessageRepo{db: db}
}

func (r *NotificationStationMessageRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationStationMessageRepo) Save(ctx context.Context, message *model.NotificationStationMessage) (*model.NotificationStationMessage, error) {
	save, err := r.getClient(ctx).NotificationStationMessage.Create().
		SetEventID(message.EventID).
		SetEventType(notificationstationmessage.EventType(message.EventType)).
		SetReceiverID(message.ReceiverID).
		SetTitle(message.Title).
		SetContent(message.Content).
		SetStatus(notificationstationmessage.Status(notifyenum.NotificationChannelStatusSucceeded)).
		SetNillableReadAt(message.ReadAt).
		Save(ctx)
	if err == nil {
		return &model.NotificationStationMessage{
			ID:         save.ID,
			EventID:    save.EventID,
			EventType:  commonenum.EventType(save.EventType),
			ReceiverID: save.ReceiverID,
			Title:      save.Title,
			Content:    save.Content,
			Status:     notifyenum.NotificationChannelStatus(save.Status),
			ReadAt:     save.ReadAt,
			CreatedAt:  save.CreatedAt,
			UpdatedAt:  save.UpdatedAt,
		}, nil
	}
	if !gen.IsConstraintError(err) {
		return nil, err
	}
	exist, getErr := r.getClient(ctx).NotificationStationMessage.Query().
		Where(
			notificationstationmessage.EventIDEQ(message.EventID),
			notificationstationmessage.ReceiverIDEQ(message.ReceiverID),
		).
		Only(ctx)
	if getErr != nil {
		return nil, getErr
	}
	return &model.NotificationStationMessage{
		ID:         exist.ID,
		EventID:    exist.EventID,
		EventType:  commonenum.EventType(exist.EventType),
		ReceiverID: exist.ReceiverID,
		Title:      exist.Title,
		Content:    exist.Content,
		Status:     notifyenum.NotificationChannelStatus(exist.Status),
		ReadAt:     exist.ReadAt,
		CreatedAt:  exist.CreatedAt,
		UpdatedAt:  exist.UpdatedAt,
	}, nil
}

func (r *NotificationStationMessageRepo) Page(ctx context.Context, page *common.PageRequest, req *bizrepo.NotificationStationMessageQuery) ([]*model.NotificationStationMessage, *common.PageReply, error) {
	page = constant.PageValid(page)
	query := r.getClient(ctx).NotificationStationMessage.Query()
	if req != nil {
		if len(req.IDs) > 0 {
			query = query.Where(notificationstationmessage.IDIn(req.IDs...))
		}
		if req.ReceiverID != nil {
			query = query.Where(notificationstationmessage.ReceiverIDEQ(*req.ReceiverID))
		}
		if req.EventType != nil {
			eventType, ok := commonenum.EventTypeMap.ToEnum(*req.EventType)
			if ok {
				query = query.Where(notificationstationmessage.EventTypeEQ(notificationstationmessage.EventType(eventType)))
			}
		}
		if req.Unread != nil {
			if *req.Unread {
				query = query.Where(notificationstationmessage.ReadAtIsNil())
			} else {
				query = query.Where(notificationstationmessage.ReadAtNotNil())
			}
		}
	}

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.
		Order(gen.Desc(notificationstationmessage.FieldCreatedAt), gen.Desc(notificationstationmessage.FieldID)).
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}

	items := make([]*model.NotificationStationMessage, 0, len(list))
	for _, item := range list {
		items = append(items, &model.NotificationStationMessage{
			ID:         item.ID,
			EventID:    item.EventID,
			EventType:  commonenum.EventType(item.EventType),
			ReceiverID: item.ReceiverID,
			Title:      item.Title,
			Content:    item.Content,
			Status:     notifyenum.NotificationChannelStatus(item.Status),
			ReadAt:     item.ReadAt,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		})
	}
	return items, &common.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *NotificationStationMessageRepo) MarkRead(ctx context.Context, receiverID int64, ids []int64, startTime *time.Time, endTime *time.Time) (int, error) {
	update := r.getClient(ctx).NotificationStationMessage.Update().
		Where(
			notificationstationmessage.ReceiverIDEQ(receiverID),
			notificationstationmessage.ReadAtIsNil(),
		)
	if len(ids) > 0 {
		update = update.Where(notificationstationmessage.IDIn(ids...))
	}
	if startTime != nil {
		update = update.Where(notificationstationmessage.CreatedAtGTE(*startTime))
	}
	if endTime != nil {
		update = update.Where(notificationstationmessage.CreatedAtLTE(*endTime))
	}
	return update.SetReadAt(time.Now()).Save(ctx)
}

func (r *NotificationStationMessageRepo) CountUnread(ctx context.Context, receiverID int64) (int, error) {
	return r.getClient(ctx).NotificationStationMessage.Query().
		Where(
			notificationstationmessage.ReceiverIDEQ(receiverID),
			notificationstationmessage.ReadAtIsNil(),
		).
		Count(ctx)
}
