package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationtencentsmsdelivery"
	"notify/internal/data/gen/predicate"
	notifyenum "notify/internal/enum"

	utilent "common/pkg/util/ent"
)

var _ bizrepo.NotificationTencentSMSDeliveryRepo = (*NotificationTencentSMSDeliveryRepo)(nil)

type NotificationTencentSMSDeliveryRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewNotificationTencentSMSDeliveryRepo(
	db *gen.Client,
) bizrepo.NotificationTencentSMSDeliveryRepo {
	return &NotificationTencentSMSDeliveryRepo{
		db: db,
	}
}

func (r *NotificationTencentSMSDeliveryRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationTencentSMSDeliveryRepo) SaveOrGet(ctx context.Context, delivery *model.NotificationTencentSMSDelivery) (*model.NotificationTencentSMSDelivery, error) {
	save, err := r.getClient(ctx).NotificationTencentSMSDelivery.Create().
		SetEventID(delivery.EventID).
		SetEventType(notificationtencentsmsdelivery.EventType(delivery.EventType)).
		SetNillableReceiverID(delivery.ReceiverID).
		SetPhone(delivery.Phone).
		SetSmsSdkAppID(delivery.SMSSDKAppID).
		SetSignName(delivery.SignName).
		SetProviderTemplateID(delivery.ProviderTemplateID).
		SetTemplateParams(delivery.TemplateParams).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusProcessing)).
		SetAttemptCount(0).
		SetNillableProviderRequestID(delivery.ProviderRequestID).
		SetNillableProviderCode(delivery.ProviderCode).
		SetNillableProviderMessage(delivery.ProviderMessage).
		SetNillableSentAt(delivery.SentAt).
		Save(ctx)
	if err == nil {
		return r.tencentSMSDelivery(save), nil
	}
	if !gen.IsConstraintError(err) {
		return nil, err
	}
	exist, getErr := r.getClient(ctx).NotificationTencentSMSDelivery.Query().
		Where(notificationtencentsmsdelivery.EventIDEQ(delivery.EventID), notificationtencentsmsdelivery.PhoneEQ(delivery.Phone)).
		Only(ctx)
	if getErr != nil {
		return nil, getErr
	}
	return r.tencentSMSDelivery(exist), nil
}

func (r *NotificationTencentSMSDeliveryRepo) Get(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryQuery) (*model.NotificationTencentSMSDelivery, error) {
	list, err := r.List(ctx, req)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationTencentSMSDeliveryRepo) List(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryQuery) ([]*model.NotificationTencentSMSDelivery, error) {
	query := r.getClient(ctx).NotificationTencentSMSDelivery.Query()
	query = r.getTencentSMSQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationTencentSMSDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, r.tencentSMSDelivery(item))
	}
	return result, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryQuery) (map[int64]*model.
	NotificationTencentSMSDelivery, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationTencentSMSDelivery, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Count(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryQuery) (int, error) {
	query := r.getClient(ctx).NotificationTencentSMSDelivery.Query()
	query = r.getTencentSMSQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Page(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryQuery) (*bizrepo.NotificationTencentSMSDeliveryPageResp, error) {
	queryReq := req
	var pageReq *base.PageRequest
	if queryReq != nil {
		pageReq = queryReq.Page
	}
	page := r.normalizePage(pageReq)
	query := r.getClient(ctx).NotificationTencentSMSDelivery.Query()
	query = r.getTencentSMSQuery(query, queryReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationTencentSMSDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, r.tencentSMSDelivery(item))
	}
	return &bizrepo.NotificationTencentSMSDeliveryPageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Claim(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryClaimReq) (bool, error) {
	conditions := []predicate.NotificationTencentSMSDelivery{
		notificationtencentsmsdelivery.StatusEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusFailed)),
		notificationtencentsmsdelivery.And(
			notificationtencentsmsdelivery.StatusEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusProcessing)),
			notificationtencentsmsdelivery.Or(
				notificationtencentsmsdelivery.LastAttemptAtIsNil(),
				notificationtencentsmsdelivery.LastAttemptAtLTE(req.Now.Add(-req.ProcessingTimeout)),
			),
		),
	}
	if req.RetryUnknown {
		conditions = append(conditions, notificationtencentsmsdelivery.StatusEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusUnknown)))
	}
	count, err := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(notificationtencentsmsdelivery.IDEQ(req.ID), notificationtencentsmsdelivery.Or(conditions...)).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusProcessing)).
		SetLastAttemptAt(req.Now).
		AddAttemptCount(1).
		Save(ctx)
	return count > 0, err
}

func (r *NotificationTencentSMSDeliveryRepo) UpdateStatus(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryUpdateStatusReq) error {
	update := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(notificationtencentsmsdelivery.IDEQ(req.ID), notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationtencentsmsdelivery.Status(req.Status)).
		SetNillableProviderRequestID(req.ProviderRequestID).
		SetNillableProviderCode(req.ProviderCode).
		SetNillableProviderMessage(req.ProviderMessage)
	if req.SentAt != nil {
		update.SetSentAt(*req.SentAt)
	}
	return update.Exec(ctx)
}

func (r *NotificationTencentSMSDeliveryRepo) getTencentSMSQuery(query *gen.NotificationTencentSMSDeliveryQuery, req *bizrepo.NotificationTencentSMSDeliveryQuery) *gen.NotificationTencentSMSDeliveryQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(notificationtencentsmsdelivery.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(notificationtencentsmsdelivery.IDIn(req.IDs...))
	}
	if req.EventID != nil {
		query = query.Where(notificationtencentsmsdelivery.EventIDEQ(*req.EventID))
	}
	if len(req.EventIDs) > 0 {
		query = query.Where(notificationtencentsmsdelivery.EventIDIn(req.EventIDs...))
	}
	if req.EventType != nil {
		query = query.Where(notificationtencentsmsdelivery.EventTypeEQ(notificationtencentsmsdelivery.EventType(*req.EventType)))
	}
	if req.ReceiverID != nil {
		query = query.Where(notificationtencentsmsdelivery.ReceiverIDEQ(*req.ReceiverID))
	}
	if req.Phone != nil {
		query = query.Where(notificationtencentsmsdelivery.PhoneEQ(*req.Phone))
	}
	if req.Status != nil {
		query = query.Where(notificationtencentsmsdelivery.StatusEQ(notificationtencentsmsdelivery.Status(*req.Status)))
	}
	return query
}

func (r *NotificationTencentSMSDeliveryRepo) tencentSMSDelivery(item *gen.NotificationTencentSMSDelivery) *model.NotificationTencentSMSDelivery {
	if item == nil {
		return nil
	}
	return &model.NotificationTencentSMSDelivery{
		ID:                 item.ID,
		EventID:            item.EventID,
		EventType:          commonenum.EventType(item.EventType),
		ReceiverID:         item.ReceiverID,
		Phone:              item.Phone,
		SMSSDKAppID:        item.SmsSdkAppID,
		SignName:           item.SignName,
		ProviderTemplateID: item.ProviderTemplateID,
		TemplateParams:     item.TemplateParams,
		Status:             notifyenum.NotificationChannelStatus(item.Status),
		AttemptCount:       item.AttemptCount,
		LastAttemptAt:      item.LastAttemptAt,
		ProviderRequestID:  item.ProviderRequestID,
		ProviderCode:       item.ProviderCode,
		ProviderMessage:    item.ProviderMessage,
		SentAt:             item.SentAt,
		CreatedAt:          item.CreatedAt,
		UpdatedAt:          item.UpdatedAt,
	}
}
