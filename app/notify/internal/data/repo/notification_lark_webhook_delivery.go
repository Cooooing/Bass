package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationlarkwebhookdelivery"
	"notify/internal/data/gen/predicate"
	notifyenum "notify/internal/enum"

	utilent "common/pkg/util/ent"
)

var _ bizrepo.NotificationLarkWebhookDeliveryRepo = (*NotificationLarkWebhookDeliveryRepo)(nil)

type NotificationLarkWebhookDeliveryRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewNotificationLarkWebhookDeliveryRepo(
	db *gen.Client,
) bizrepo.NotificationLarkWebhookDeliveryRepo {
	return &NotificationLarkWebhookDeliveryRepo{
		db: db,
	}
}

func (r *NotificationLarkWebhookDeliveryRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationLarkWebhookDeliveryRepo) SaveOrGet(ctx context.Context, delivery *model.NotificationLarkWebhookDelivery) (*model.NotificationLarkWebhookDelivery, error) {
	save, err := r.getClient(ctx).NotificationLarkWebhookDelivery.Create().
		SetEventID(delivery.EventID).
		SetEventType(notificationlarkwebhookdelivery.EventType(delivery.EventType)).
		SetWebhookID(delivery.WebhookID).
		SetRequestBody(delivery.RequestBody).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusProcessing)).
		SetAttemptCount(0).
		SetNillableHTTPStatus(delivery.HTTPStatus).
		SetNillableRespBody(delivery.RespBody).
		SetNillableSentAt(delivery.SentAt).
		Save(ctx)
	if err == nil {
		return r.larkWebhookDelivery(save), nil
	}
	if !gen.IsConstraintError(err) {
		return nil, err
	}
	exist, getErr := r.getClient(ctx).NotificationLarkWebhookDelivery.Query().
		Where(notificationlarkwebhookdelivery.EventIDEQ(delivery.EventID), notificationlarkwebhookdelivery.WebhookIDEQ(delivery.WebhookID)).
		Only(ctx)
	if getErr != nil {
		return nil, getErr
	}
	return r.larkWebhookDelivery(exist), nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Get(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryQuery) (*model.NotificationLarkWebhookDelivery, error) {
	list, err := r.List(ctx, req)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationLarkWebhookDeliveryRepo) List(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryQuery) ([]*model.NotificationLarkWebhookDelivery, error) {
	query := r.getClient(ctx).NotificationLarkWebhookDelivery.Query()
	query = r.getLarkWebhookQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationLarkWebhookDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, r.larkWebhookDelivery(item))
	}
	return result, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryQuery) (map[int64]*model.
	NotificationLarkWebhookDelivery, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationLarkWebhookDelivery, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Count(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryQuery) (int, error) {
	query := r.getClient(ctx).NotificationLarkWebhookDelivery.Query()
	query = r.getLarkWebhookQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Page(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryQuery) (*bizrepo.NotificationLarkWebhookDeliveryPageResp, error) {
	queryReq := req
	var pageReq *base.PageRequest
	if queryReq != nil {
		pageReq = queryReq.Page
	}
	page := r.normalizePage(pageReq)
	query := r.getClient(ctx).NotificationLarkWebhookDelivery.Query()
	query = r.getLarkWebhookQuery(query, queryReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationLarkWebhookDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, r.larkWebhookDelivery(item))
	}
	return &bizrepo.NotificationLarkWebhookDeliveryPageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Claim(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryClaimReq) (bool, error) {
	conditions := []predicate.NotificationLarkWebhookDelivery{
		notificationlarkwebhookdelivery.StatusEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusFailed)),
		notificationlarkwebhookdelivery.And(
			notificationlarkwebhookdelivery.StatusEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusProcessing)),
			notificationlarkwebhookdelivery.Or(
				notificationlarkwebhookdelivery.LastAttemptAtIsNil(),
				notificationlarkwebhookdelivery.LastAttemptAtLTE(req.Now.Add(-req.ProcessingTimeout)),
			),
		),
	}
	if req.RetryUnknown {
		conditions = append(conditions, notificationlarkwebhookdelivery.StatusEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusUnknown)))
	}
	count, err := r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(notificationlarkwebhookdelivery.IDEQ(req.ID), notificationlarkwebhookdelivery.Or(conditions...)).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusProcessing)).
		SetLastAttemptAt(req.Now).
		AddAttemptCount(1).
		Save(ctx)
	return count > 0, err
}

func (r *NotificationLarkWebhookDeliveryRepo) UpdateStatus(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryUpdateStatusReq) error {
	update := r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(notificationlarkwebhookdelivery.IDEQ(req.ID), notificationlarkwebhookdelivery.StatusNEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationlarkwebhookdelivery.Status(req.Status)).
		SetNillableHTTPStatus(req.HTTPStatus).
		SetNillableRespBody(req.RespBody)
	if req.SentAt != nil {
		update.SetSentAt(*req.SentAt)
	}
	return update.Exec(ctx)
}

func (r *NotificationLarkWebhookDeliveryRepo) getLarkWebhookQuery(query *gen.NotificationLarkWebhookDeliveryQuery, req *bizrepo.NotificationLarkWebhookDeliveryQuery) *gen.NotificationLarkWebhookDeliveryQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(notificationlarkwebhookdelivery.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(notificationlarkwebhookdelivery.IDIn(req.IDs...))
	}
	if req.EventID != nil {
		query = query.Where(notificationlarkwebhookdelivery.EventIDEQ(*req.EventID))
	}
	if len(req.EventIDs) > 0 {
		query = query.Where(notificationlarkwebhookdelivery.EventIDIn(req.EventIDs...))
	}
	if req.EventType != nil {
		query = query.Where(notificationlarkwebhookdelivery.EventTypeEQ(notificationlarkwebhookdelivery.EventType(*req.EventType)))
	}
	if req.WebhookID != nil {
		query = query.Where(notificationlarkwebhookdelivery.WebhookIDEQ(*req.WebhookID))
	}
	if req.Status != nil {
		query = query.Where(notificationlarkwebhookdelivery.StatusEQ(notificationlarkwebhookdelivery.Status(*req.Status)))
	}
	return query
}

func (r *NotificationLarkWebhookDeliveryRepo) larkWebhookDelivery(item *gen.NotificationLarkWebhookDelivery) *model.NotificationLarkWebhookDelivery {
	if item == nil {
		return nil
	}
	return &model.NotificationLarkWebhookDelivery{
		ID:            item.ID,
		EventID:       item.EventID,
		EventType:     commonenum.EventType(item.EventType),
		WebhookID:     item.WebhookID,
		RequestBody:   item.RequestBody,
		Status:        notifyenum.NotificationChannelStatus(item.Status),
		AttemptCount:  item.AttemptCount,
		LastAttemptAt: item.LastAttemptAt,
		HTTPStatus:    item.HTTPStatus,
		RespBody:      item.RespBody,
		SentAt:        item.SentAt,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}
}
