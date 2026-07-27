package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationemaildelivery"
	"notify/internal/data/gen/predicate"
	notifyenum "notify/internal/enum"

	utilent "common/pkg/util/ent"
)

var _ bizrepo.NotificationEmailDeliveryRepo = (*NotificationEmailDeliveryRepo)(nil)

type NotificationEmailDeliveryRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewNotificationEmailDeliveryRepo(
	db *gen.Client,
) bizrepo.NotificationEmailDeliveryRepo {
	return &NotificationEmailDeliveryRepo{
		db: db,
	}
}

func (r *NotificationEmailDeliveryRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationEmailDeliveryRepo) SaveOrGet(ctx context.Context, delivery *model.NotificationEmailDelivery) (*model.NotificationEmailDelivery, error) {
	save, err := r.getClient(ctx).NotificationEmailDelivery.Create().
		SetEventID(delivery.EventID).
		SetEventType(notificationemaildelivery.EventType(delivery.EventType)).
		SetNillableReceiverID(delivery.ReceiverID).
		SetToEmail(delivery.ToEmail).
		SetSubject(delivery.Subject).
		SetBody(delivery.Body).
		SetContentType(delivery.ContentType).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusProcessing)).
		SetAttemptCount(0).
		SetNillableProviderMessageID(delivery.ProviderMessageID).
		SetNillableProviderResp(delivery.ProviderResp).
		SetNillableSentAt(delivery.SentAt).
		Save(ctx)
	if err == nil {
		return r.emailDelivery(save), nil
	}
	if !gen.IsConstraintError(err) {
		return nil, err
	}
	exist, getErr := r.getClient(ctx).NotificationEmailDelivery.Query().
		Where(
			notificationemaildelivery.EventIDEQ(delivery.EventID),
			notificationemaildelivery.ToEmailEQ(delivery.ToEmail),
		).
		Only(ctx)
	if getErr != nil {
		return nil, getErr
	}
	return r.emailDelivery(exist), nil
}

func (r *NotificationEmailDeliveryRepo) Get(ctx context.Context, req *bizrepo.NotificationEmailDeliveryQuery) (*model.NotificationEmailDelivery, error) {
	list, err := r.List(ctx, req)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationEmailDeliveryRepo) List(ctx context.Context, req *bizrepo.NotificationEmailDeliveryQuery) ([]*model.NotificationEmailDelivery, error) {
	query := r.getClient(ctx).NotificationEmailDelivery.Query()
	query = r.getEmailQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationEmailDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, r.emailDelivery(item))
	}
	return result, nil
}

func (r *NotificationEmailDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationEmailDeliveryQuery) (map[int64]*model.
	NotificationEmailDelivery, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationEmailDelivery, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *NotificationEmailDeliveryRepo) Count(ctx context.Context, req *bizrepo.NotificationEmailDeliveryQuery) (int, error) {
	query := r.getClient(ctx).NotificationEmailDelivery.Query()
	query = r.getEmailQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationEmailDeliveryRepo) Page(ctx context.Context, req *bizrepo.NotificationEmailDeliveryQuery) (*bizrepo.NotificationEmailDeliveryPageResp, error) {
	queryReq := req
	var pageReq *base.PageRequest
	if queryReq != nil {
		pageReq = queryReq.Page
	}
	page := r.normalizePage(pageReq)
	query := r.getClient(ctx).NotificationEmailDelivery.Query()
	query = r.getEmailQuery(query, queryReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationEmailDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, r.emailDelivery(item))
	}
	return &bizrepo.NotificationEmailDeliveryPageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *NotificationEmailDeliveryRepo) Claim(ctx context.Context, req *bizrepo.NotificationEmailDeliveryClaimReq) (bool, error) {
	conditions := []predicate.NotificationEmailDelivery{
		notificationemaildelivery.StatusEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusFailed)),
		notificationemaildelivery.And(
			notificationemaildelivery.StatusEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusProcessing)),
			notificationemaildelivery.Or(
				notificationemaildelivery.LastAttemptAtIsNil(),
				notificationemaildelivery.LastAttemptAtLTE(req.Now.Add(-req.ProcessingTimeout)),
			),
		),
	}
	if req.RetryUnknown {
		conditions = append(conditions, notificationemaildelivery.StatusEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusUnknown)))
	}
	count, err := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(notificationemaildelivery.IDEQ(req.ID), notificationemaildelivery.Or(conditions...)).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusProcessing)).
		SetLastAttemptAt(req.Now).
		AddAttemptCount(1).
		Save(ctx)
	return count > 0, err
}

func (r *NotificationEmailDeliveryRepo) UpdateStatus(ctx context.Context, req *bizrepo.NotificationEmailDeliveryUpdateStatusReq) error {
	update := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(notificationemaildelivery.IDEQ(req.ID), notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationemaildelivery.Status(req.Status)).
		SetNillableProviderMessageID(req.ProviderMessageID).
		SetNillableProviderResp(req.ProviderResp)
	if req.SentAt != nil {
		update.SetSentAt(*req.SentAt)
	}
	return update.Exec(ctx)
}

func (r *NotificationEmailDeliveryRepo) getEmailQuery(query *gen.NotificationEmailDeliveryQuery, req *bizrepo.NotificationEmailDeliveryQuery) *gen.NotificationEmailDeliveryQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(notificationemaildelivery.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(notificationemaildelivery.IDIn(req.IDs...))
	}
	if req.EventID != nil {
		query = query.Where(notificationemaildelivery.EventIDEQ(*req.EventID))
	}
	if len(req.EventIDs) > 0 {
		query = query.Where(notificationemaildelivery.EventIDIn(req.EventIDs...))
	}
	if req.EventType != nil {
		query = query.Where(notificationemaildelivery.EventTypeEQ(notificationemaildelivery.EventType(*req.EventType)))
	}
	if req.ReceiverID != nil {
		query = query.Where(notificationemaildelivery.ReceiverIDEQ(*req.ReceiverID))
	}
	if req.ToEmail != nil {
		query = query.Where(notificationemaildelivery.ToEmailEQ(*req.ToEmail))
	}
	if req.Status != nil {
		query = query.Where(notificationemaildelivery.StatusEQ(notificationemaildelivery.Status(*req.Status)))
	}
	return query
}

func (r *NotificationEmailDeliveryRepo) emailDelivery(item *gen.NotificationEmailDelivery) *model.NotificationEmailDelivery {
	if item == nil {
		return nil
	}
	return &model.NotificationEmailDelivery{
		ID:                item.ID,
		EventID:           item.EventID,
		EventType:         commonenum.EventType(item.EventType),
		ReceiverID:        item.ReceiverID,
		ToEmail:           item.ToEmail,
		Subject:           item.Subject,
		Body:              item.Body,
		ContentType:       item.ContentType,
		Status:            notifyenum.NotificationChannelStatus(item.Status),
		AttemptCount:      item.AttemptCount,
		LastAttemptAt:     item.LastAttemptAt,
		ProviderMessageID: item.ProviderMessageID,
		ProviderResp:      item.ProviderResp,
		SentAt:            item.SentAt,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}
}
