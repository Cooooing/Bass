package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"fmt"
	"notify/internal/biz/base"
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

func NewNotificationStationMessageRepo(
	db *gen.Client,
) bizrepo.NotificationStationMessageRepo {
	return &NotificationStationMessageRepo{
		db: db,
	}
}

func (r *NotificationStationMessageRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationStationMessageRepo) Save(ctx context.Context, message *model.NotificationStationMessage) (*model.NotificationStationMessage, error) {
	if message == nil {
		return nil, fmt.Errorf("notification station message save request is nil")
	}
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
		return stationMessageModel(save), nil
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
	return stationMessageModel(exist), nil
}

func (r *NotificationStationMessageRepo) Get(ctx context.Context, req *bizrepo.NotificationStationMessageQuery) (*model.NotificationStationMessage, error) {
	query := r.getClient(ctx).NotificationStationMessage.Query()
	query = r.getQuery(query, stationMessageGetQuery(req))
	item, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return stationMessageModel(item), nil
}

func (r *NotificationStationMessageRepo) List(ctx context.Context, req *bizrepo.NotificationStationMessageQuery) ([]*model.NotificationStationMessage, error) {
	query := r.getClient(ctx).NotificationStationMessage.Query()
	query = r.getQuery(query, stationMessageListQuery(req))
	list, err := query.
		Order(gen.Desc(notificationstationmessage.FieldCreatedAt), gen.Desc(notificationstationmessage.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*model.NotificationStationMessage, 0, len(list))
	for _, item := range list {
		items = append(items, stationMessageModel(item))
	}
	return items, nil
}

func (r *NotificationStationMessageRepo) Map(ctx context.Context, req *bizrepo.NotificationStationMessageQuery) (map[int64]*model.
	NotificationStationMessage, error) {
	list, err := r.List(ctx, stationMessageMapQuery(req))
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationStationMessage, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *NotificationStationMessageRepo) Count(ctx context.Context, req *bizrepo.NotificationStationMessageQuery) (int, error) {
	query := r.getClient(ctx).NotificationStationMessage.Query()
	query = r.getQuery(query, stationMessageCountQuery(req))
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationStationMessageRepo) Page(ctx context.Context, req *bizrepo.NotificationStationMessageQuery) (*bizrepo.NotificationStationMessagePageResp, error) {
	queryReq := stationMessagePageQuery(req)
	var pageReq *base.PageRequest
	if queryReq != nil {
		pageReq = queryReq.Page
	}
	page := normalizePage(pageReq)
	query := r.getClient(ctx).NotificationStationMessage.Query()
	query = r.getQuery(query, queryReq)

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.
		Order(gen.Desc(notificationstationmessage.FieldCreatedAt), gen.Desc(notificationstationmessage.FieldID)).
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*model.NotificationStationMessage, 0, len(list))
	for _, item := range list {
		items = append(items, stationMessageModel(item))
	}
	return &bizrepo.NotificationStationMessagePageResp{
		Rows: items,
		Page: &base.PageResp{
			Total: int64(count),
			Size:  page.Size,
			Page:  page.Page,
		},
	}, nil
}

func (r *NotificationStationMessageRepo) MarkRead(ctx context.Context, req *bizrepo.NotificationStationMessageMarkReadReq) (int, error) {
	if req == nil {
		return 0, fmt.Errorf("notification station message mark read request is nil")
	}
	update := r.getClient(ctx).NotificationStationMessage.Update().
		Where(
			notificationstationmessage.ReceiverIDEQ(req.ReceiverID),
			notificationstationmessage.ReadAtIsNil(),
		)
	if len(req.IDs) > 0 {
		update = update.Where(notificationstationmessage.IDIn(req.IDs...))
	}
	if req.StartTime != nil {
		update = update.Where(notificationstationmessage.CreatedAtGTE(*req.StartTime))
	}
	if req.EndTime != nil {
		update = update.Where(notificationstationmessage.CreatedAtLTE(*req.EndTime))
	}
	count, err := update.SetReadAt(time.Now()).Save(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationStationMessageRepo) CountUnread(ctx context.Context, receiverID int64) (int, error) {

	count, err := r.getClient(ctx).NotificationStationMessage.Query().
		Where(
			notificationstationmessage.ReceiverIDEQ(receiverID),
			notificationstationmessage.ReadAtIsNil(),
		).
		Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationStationMessageRepo) getQuery(query *gen.NotificationStationMessageQuery, req *bizrepo.NotificationStationMessageQuery) *gen.NotificationStationMessageQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(notificationstationmessage.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(notificationstationmessage.IDIn(req.IDs...))
	}
	if len(req.EventIDs) > 0 {
		query = query.Where(notificationstationmessage.EventIDIn(req.EventIDs...))
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
	return query
}

func stationMessageGetQuery(query *bizrepo.NotificationStationMessageQuery) *bizrepo.NotificationStationMessageQuery {

	return query
}

func stationMessageListQuery(query *bizrepo.NotificationStationMessageQuery) *bizrepo.NotificationStationMessageQuery {

	return query
}

func stationMessageMapQuery(query *bizrepo.NotificationStationMessageQuery) *bizrepo.NotificationStationMessageQuery {

	return query
}

func stationMessageCountQuery(query *bizrepo.NotificationStationMessageQuery) *bizrepo.NotificationStationMessageQuery {

	return query
}

func stationMessagePageQuery(query *bizrepo.NotificationStationMessageQuery) *bizrepo.NotificationStationMessageQuery {

	return query
}

func stationMessageModel(item *gen.NotificationStationMessage) *model.NotificationStationMessage {
	if item == nil {
		return nil
	}
	return &model.NotificationStationMessage{
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
	}
}
