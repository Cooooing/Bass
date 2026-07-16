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

func NewNotificationStationMessageRepo(db *gen.Client) bizrepo.NotificationStationMessageRepo {
	return &NotificationStationMessageRepo{db: db}
}

func (r *NotificationStationMessageRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationStationMessageRepo) Save(ctx context.Context, req *bizrepo.NotificationStationMessageSaveReq) (*bizrepo.NotificationStationMessageSaveResponse, error) {
	if req == nil || req.Message == nil {
		return nil, fmt.Errorf("notification station message save request is nil")
	}
	message := req.Message
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
		return &bizrepo.NotificationStationMessageSaveResponse{Message: stationMessageModel(save)}, nil
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
	return &bizrepo.NotificationStationMessageSaveResponse{Message: stationMessageModel(exist)}, nil
}

func (r *NotificationStationMessageRepo) Get(ctx context.Context, req *bizrepo.NotificationStationMessageGetReq) (*bizrepo.NotificationStationMessageGetResponse, error) {
	query := r.getClient(ctx).NotificationStationMessage.Query()
	query = r.getQuery(query, stationMessageGetQuery(req))
	item, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return &bizrepo.NotificationStationMessageGetResponse{}, nil
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationStationMessageGetResponse{Item: stationMessageModel(item)}, nil
}

func (r *NotificationStationMessageRepo) List(ctx context.Context, req *bizrepo.NotificationStationMessageListReq) (*bizrepo.NotificationStationMessageListResponse, error) {
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
	return &bizrepo.NotificationStationMessageListResponse{Rows: items}, nil
}

func (r *NotificationStationMessageRepo) Map(ctx context.Context, req *bizrepo.NotificationStationMessageMapReq) (*bizrepo.NotificationStationMessageMapResponse, error) {
	listResponse, err := r.List(ctx, &bizrepo.NotificationStationMessageListReq{Query: stationMessageMapQuery(req)})
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationStationMessage, len(listResponse.Rows))
	for _, item := range listResponse.Rows {
		result[item.ID] = item
	}
	return &bizrepo.NotificationStationMessageMapResponse{Rows: result}, nil
}

func (r *NotificationStationMessageRepo) Count(ctx context.Context, req *bizrepo.NotificationStationMessageCountReq) (*bizrepo.NotificationStationMessageCountResponse, error) {
	query := r.getClient(ctx).NotificationStationMessage.Query()
	query = r.getQuery(query, stationMessageCountQuery(req))
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationStationMessageCountResponse{Count: count}, nil
}

func (r *NotificationStationMessageRepo) Page(ctx context.Context, req *bizrepo.NotificationStationMessagePageReq) (*bizrepo.NotificationStationMessagePageResponse, error) {
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
	return &bizrepo.NotificationStationMessagePageResponse{
		Rows: items,
		Page: &base.PageResponse{
			Total: int64(count),
			Size:  page.Size,
			Page:  page.Page,
		},
	}, nil
}

func (r *NotificationStationMessageRepo) MarkRead(ctx context.Context, req *bizrepo.NotificationStationMessageMarkReadReq) (*bizrepo.NotificationStationMessageMarkReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("notification station message mark read request is nil")
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
		return nil, err
	}
	return &bizrepo.NotificationStationMessageMarkReadResponse{Count: count}, nil
}

func (r *NotificationStationMessageRepo) CountUnread(ctx context.Context, req *bizrepo.NotificationStationMessageCountUnreadReq) (*bizrepo.NotificationStationMessageCountUnreadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("notification station message count unread request is nil")
	}
	count, err := r.getClient(ctx).NotificationStationMessage.Query().
		Where(
			notificationstationmessage.ReceiverIDEQ(req.ReceiverID),
			notificationstationmessage.ReadAtIsNil(),
		).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationStationMessageCountUnreadResponse{Count: count}, nil
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

func stationMessageGetQuery(req *bizrepo.NotificationStationMessageGetReq) *bizrepo.NotificationStationMessageQuery {
	if req == nil {
		return nil
	}
	return req.Query
}

func stationMessageListQuery(req *bizrepo.NotificationStationMessageListReq) *bizrepo.NotificationStationMessageQuery {
	if req == nil {
		return nil
	}
	return req.Query
}

func stationMessageMapQuery(req *bizrepo.NotificationStationMessageMapReq) *bizrepo.NotificationStationMessageQuery {
	if req == nil {
		return nil
	}
	return req.Query
}

func stationMessageCountQuery(req *bizrepo.NotificationStationMessageCountReq) *bizrepo.NotificationStationMessageQuery {
	if req == nil {
		return nil
	}
	return req.Query
}

func stationMessagePageQuery(req *bizrepo.NotificationStationMessagePageReq) *bizrepo.NotificationStationMessageQuery {
	if req == nil {
		return nil
	}
	return req.Query
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
