package repo

import (
	commonenum "common/pkg/enum"
	"common/proto/gen/common"
	"context"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationemaildelivery"
	"notify/internal/data/gen/notificationlarkwebhookdelivery"
	"notify/internal/data/gen/notificationtencentsmsdelivery"
	"notify/internal/data/gen/predicate"
	notifyenum "notify/internal/enum"
	"time"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ bizrepo.NotificationEmailDeliveryRepo = (*NotificationEmailDeliveryRepo)(nil)
var _ bizrepo.NotificationTencentSMSDeliveryRepo = (*NotificationTencentSMSDeliveryRepo)(nil)
var _ bizrepo.NotificationLarkWebhookDeliveryRepo = (*NotificationLarkWebhookDeliveryRepo)(nil)

type NotificationEmailDeliveryRepo struct {
	db *gen.Client
}

func NewNotificationEmailDeliveryRepo(db *gen.Client) bizrepo.NotificationEmailDeliveryRepo {
	return &NotificationEmailDeliveryRepo{db: db}
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
		SetNillableProviderResponse(delivery.ProviderResponse).
		SetNillableSentAt(delivery.SentAt).
		Save(ctx)
	if err == nil {
		return &model.NotificationEmailDelivery{
			ID:                save.ID,
			EventID:           save.EventID,
			EventType:         commonenum.EventType(save.EventType),
			ReceiverID:        save.ReceiverID,
			ToEmail:           save.ToEmail,
			Subject:           save.Subject,
			Body:              save.Body,
			ContentType:       save.ContentType,
			Status:            notifyenum.NotificationChannelStatus(save.Status),
			AttemptCount:      save.AttemptCount,
			LastAttemptAt:     save.LastAttemptAt,
			ProviderMessageID: save.ProviderMessageID,
			ProviderResponse:  save.ProviderResponse,
			SentAt:            save.SentAt,
			CreatedAt:         save.CreatedAt,
			UpdatedAt:         save.UpdatedAt,
		}, nil
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
	return &model.NotificationEmailDelivery{
		ID:                exist.ID,
		EventID:           exist.EventID,
		EventType:         commonenum.EventType(exist.EventType),
		ReceiverID:        exist.ReceiverID,
		ToEmail:           exist.ToEmail,
		Subject:           exist.Subject,
		Body:              exist.Body,
		ContentType:       exist.ContentType,
		Status:            notifyenum.NotificationChannelStatus(exist.Status),
		AttemptCount:      exist.AttemptCount,
		LastAttemptAt:     exist.LastAttemptAt,
		ProviderMessageID: exist.ProviderMessageID,
		ProviderResponse:  exist.ProviderResponse,
		SentAt:            exist.SentAt,
		CreatedAt:         exist.CreatedAt,
		UpdatedAt:         exist.UpdatedAt,
	}, nil
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
		result = append(result, &model.NotificationEmailDelivery{
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
			ProviderResponse:  item.ProviderResponse,
			SentAt:            item.SentAt,
			CreatedAt:         item.CreatedAt,
			UpdatedAt:         item.UpdatedAt,
		})
	}
	return result, nil
}

func (r *NotificationEmailDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationEmailDeliveryQuery) (map[int64]*model.NotificationEmailDelivery, error) {
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
	return query.Count(ctx)
}

func (r *NotificationEmailDeliveryRepo) Page(ctx context.Context, page *common.PageRequest, req *bizrepo.NotificationEmailDeliveryQuery) ([]*model.NotificationEmailDelivery, *common.PageReply, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).NotificationEmailDelivery.Query()
	query = r.getEmailQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*model.NotificationEmailDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationEmailDelivery{
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
			ProviderResponse:  item.ProviderResponse,
			SentAt:            item.SentAt,
			CreatedAt:         item.CreatedAt,
			UpdatedAt:         item.UpdatedAt,
		})
	}
	return result, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *NotificationEmailDeliveryRepo) Claim(ctx context.Context, id int64, now time.Time, processingTimeout time.Duration, retryUnknown bool) (bool, error) {
	conditions := []predicate.NotificationEmailDelivery{
		notificationemaildelivery.StatusEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusFailed)),
		notificationemaildelivery.And(
			notificationemaildelivery.StatusEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusProcessing)),
			notificationemaildelivery.Or(
				notificationemaildelivery.LastAttemptAtIsNil(),
				notificationemaildelivery.LastAttemptAtLTE(now.Add(-processingTimeout)),
			),
		),
	}
	if retryUnknown {
		conditions = append(conditions, notificationemaildelivery.StatusEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusUnknown)))
	}
	count, err := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(
			notificationemaildelivery.IDEQ(id),
			notificationemaildelivery.Or(conditions...),
		).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusProcessing)).
		SetLastAttemptAt(now).
		AddAttemptCount(1).
		Save(ctx)
	return count > 0, err
}

func (r *NotificationEmailDeliveryRepo) MarkSucceeded(ctx context.Context, id int64, providerMessageID *string, providerResponse *string, sentAt time.Time) error {
	return r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(
			notificationemaildelivery.IDEQ(id),
			notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded)).
		SetNillableProviderMessageID(providerMessageID).
		SetNillableProviderResponse(providerResponse).
		SetSentAt(sentAt).
		Exec(ctx)
}

func (r *NotificationEmailDeliveryRepo) MarkFailed(ctx context.Context, id int64, providerResponse *string) error {
	return r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(
			notificationemaildelivery.IDEQ(id),
			notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusFailed)).
		SetNillableProviderResponse(providerResponse).
		Exec(ctx)
}

func (r *NotificationEmailDeliveryRepo) MarkUnknown(ctx context.Context, id int64, providerResponse *string) error {
	return r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(
			notificationemaildelivery.IDEQ(id),
			notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusUnknown)).
		SetNillableProviderResponse(providerResponse).
		Exec(ctx)
}

func (r *NotificationEmailDeliveryRepo) MarkRateLimited(ctx context.Context, id int64) error {
	return r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(
			notificationemaildelivery.IDEQ(id),
			notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusRateLimited)).
		Exec(ctx)
}

type NotificationTencentSMSDeliveryRepo struct {
	db *gen.Client
}

func NewNotificationTencentSMSDeliveryRepo(db *gen.Client) bizrepo.NotificationTencentSMSDeliveryRepo {
	return &NotificationTencentSMSDeliveryRepo{db: db}
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
		return &model.NotificationTencentSMSDelivery{
			ID:                 save.ID,
			EventID:            save.EventID,
			EventType:          commonenum.EventType(save.EventType),
			ReceiverID:         save.ReceiverID,
			Phone:              save.Phone,
			SMSSDKAppID:        save.SmsSdkAppID,
			SignName:           save.SignName,
			ProviderTemplateID: save.ProviderTemplateID,
			TemplateParams:     save.TemplateParams,
			Status:             notifyenum.NotificationChannelStatus(save.Status),
			AttemptCount:       save.AttemptCount,
			LastAttemptAt:      save.LastAttemptAt,
			ProviderRequestID:  save.ProviderRequestID,
			ProviderCode:       save.ProviderCode,
			ProviderMessage:    save.ProviderMessage,
			SentAt:             save.SentAt,
			CreatedAt:          save.CreatedAt,
			UpdatedAt:          save.UpdatedAt,
		}, nil
	}
	if !gen.IsConstraintError(err) {
		return nil, err
	}
	exist, getErr := r.getClient(ctx).NotificationTencentSMSDelivery.Query().
		Where(
			notificationtencentsmsdelivery.EventIDEQ(delivery.EventID),
			notificationtencentsmsdelivery.PhoneEQ(delivery.Phone),
		).
		Only(ctx)
	if getErr != nil {
		return nil, getErr
	}
	return &model.NotificationTencentSMSDelivery{
		ID:                 exist.ID,
		EventID:            exist.EventID,
		EventType:          commonenum.EventType(exist.EventType),
		ReceiverID:         exist.ReceiverID,
		Phone:              exist.Phone,
		SMSSDKAppID:        exist.SmsSdkAppID,
		SignName:           exist.SignName,
		ProviderTemplateID: exist.ProviderTemplateID,
		TemplateParams:     exist.TemplateParams,
		Status:             notifyenum.NotificationChannelStatus(exist.Status),
		AttemptCount:       exist.AttemptCount,
		LastAttemptAt:      exist.LastAttemptAt,
		ProviderRequestID:  exist.ProviderRequestID,
		ProviderCode:       exist.ProviderCode,
		ProviderMessage:    exist.ProviderMessage,
		SentAt:             exist.SentAt,
		CreatedAt:          exist.CreatedAt,
		UpdatedAt:          exist.UpdatedAt,
	}, nil
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
		result = append(result, &model.NotificationTencentSMSDelivery{
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
		})
	}
	return result, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryQuery) (map[int64]*model.NotificationTencentSMSDelivery, error) {
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
	return query.Count(ctx)
}

func (r *NotificationTencentSMSDeliveryRepo) Page(ctx context.Context, page *common.PageRequest, req *bizrepo.NotificationTencentSMSDeliveryQuery) ([]*model.NotificationTencentSMSDelivery, *common.PageReply, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).NotificationTencentSMSDelivery.Query()
	query = r.getTencentSMSQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*model.NotificationTencentSMSDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationTencentSMSDelivery{
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
		})
	}
	return result, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Claim(ctx context.Context, id int64, now time.Time, processingTimeout time.Duration, retryUnknown bool) (bool, error) {
	conditions := []predicate.NotificationTencentSMSDelivery{
		notificationtencentsmsdelivery.StatusEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusFailed)),
		notificationtencentsmsdelivery.And(
			notificationtencentsmsdelivery.StatusEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusProcessing)),
			notificationtencentsmsdelivery.Or(
				notificationtencentsmsdelivery.LastAttemptAtIsNil(),
				notificationtencentsmsdelivery.LastAttemptAtLTE(now.Add(-processingTimeout)),
			),
		),
	}
	if retryUnknown {
		conditions = append(conditions, notificationtencentsmsdelivery.StatusEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusUnknown)))
	}
	count, err := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(
			notificationtencentsmsdelivery.IDEQ(id),
			notificationtencentsmsdelivery.Or(conditions...),
		).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusProcessing)).
		SetLastAttemptAt(now).
		AddAttemptCount(1).
		Save(ctx)
	return count > 0, err
}

func (r *NotificationTencentSMSDeliveryRepo) MarkSucceeded(ctx context.Context, id int64, providerRequestID *string, providerCode *string, providerMessage *string, sentAt time.Time) error {
	return r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(
			notificationtencentsmsdelivery.IDEQ(id),
			notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)).
		SetNillableProviderRequestID(providerRequestID).
		SetNillableProviderCode(providerCode).
		SetNillableProviderMessage(providerMessage).
		SetSentAt(sentAt).
		Exec(ctx)
}

func (r *NotificationTencentSMSDeliveryRepo) MarkFailed(ctx context.Context, id int64, providerRequestID *string, providerCode *string, providerMessage *string) error {
	return r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(
			notificationtencentsmsdelivery.IDEQ(id),
			notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusFailed)).
		SetNillableProviderRequestID(providerRequestID).
		SetNillableProviderCode(providerCode).
		SetNillableProviderMessage(providerMessage).
		Exec(ctx)
}

func (r *NotificationTencentSMSDeliveryRepo) MarkUnknown(ctx context.Context, id int64, providerRequestID *string, providerCode *string, providerMessage *string) error {
	return r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(
			notificationtencentsmsdelivery.IDEQ(id),
			notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusUnknown)).
		SetNillableProviderRequestID(providerRequestID).
		SetNillableProviderCode(providerCode).
		SetNillableProviderMessage(providerMessage).
		Exec(ctx)
}

func (r *NotificationTencentSMSDeliveryRepo) MarkRateLimited(ctx context.Context, id int64) error {
	return r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(
			notificationtencentsmsdelivery.IDEQ(id),
			notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusRateLimited)).
		Exec(ctx)
}

type NotificationLarkWebhookDeliveryRepo struct {
	db *gen.Client
}

func NewNotificationLarkWebhookDeliveryRepo(db *gen.Client) bizrepo.NotificationLarkWebhookDeliveryRepo {
	return &NotificationLarkWebhookDeliveryRepo{db: db}
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
		SetNillableResponseBody(delivery.ResponseBody).
		SetNillableSentAt(delivery.SentAt).
		Save(ctx)
	if err == nil {
		return &model.NotificationLarkWebhookDelivery{
			ID:            save.ID,
			EventID:       save.EventID,
			EventType:     commonenum.EventType(save.EventType),
			WebhookID:     save.WebhookID,
			RequestBody:   save.RequestBody,
			Status:        notifyenum.NotificationChannelStatus(save.Status),
			AttemptCount:  save.AttemptCount,
			LastAttemptAt: save.LastAttemptAt,
			HTTPStatus:    save.HTTPStatus,
			ResponseBody:  save.ResponseBody,
			SentAt:        save.SentAt,
			CreatedAt:     save.CreatedAt,
			UpdatedAt:     save.UpdatedAt,
		}, nil
	}
	if !gen.IsConstraintError(err) {
		return nil, err
	}
	exist, getErr := r.getClient(ctx).NotificationLarkWebhookDelivery.Query().
		Where(
			notificationlarkwebhookdelivery.EventIDEQ(delivery.EventID),
			notificationlarkwebhookdelivery.WebhookIDEQ(delivery.WebhookID),
		).
		Only(ctx)
	if getErr != nil {
		return nil, getErr
	}
	return &model.NotificationLarkWebhookDelivery{
		ID:            exist.ID,
		EventID:       exist.EventID,
		EventType:     commonenum.EventType(exist.EventType),
		WebhookID:     exist.WebhookID,
		RequestBody:   exist.RequestBody,
		Status:        notifyenum.NotificationChannelStatus(exist.Status),
		AttemptCount:  exist.AttemptCount,
		LastAttemptAt: exist.LastAttemptAt,
		HTTPStatus:    exist.HTTPStatus,
		ResponseBody:  exist.ResponseBody,
		SentAt:        exist.SentAt,
		CreatedAt:     exist.CreatedAt,
		UpdatedAt:     exist.UpdatedAt,
	}, nil
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
		result = append(result, &model.NotificationLarkWebhookDelivery{
			ID:            item.ID,
			EventID:       item.EventID,
			EventType:     commonenum.EventType(item.EventType),
			WebhookID:     item.WebhookID,
			RequestBody:   item.RequestBody,
			Status:        notifyenum.NotificationChannelStatus(item.Status),
			AttemptCount:  item.AttemptCount,
			LastAttemptAt: item.LastAttemptAt,
			HTTPStatus:    item.HTTPStatus,
			ResponseBody:  item.ResponseBody,
			SentAt:        item.SentAt,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
		})
	}
	return result, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryQuery) (map[int64]*model.NotificationLarkWebhookDelivery, error) {
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
	return query.Count(ctx)
}

func (r *NotificationLarkWebhookDeliveryRepo) Page(ctx context.Context, page *common.PageRequest, req *bizrepo.NotificationLarkWebhookDeliveryQuery) ([]*model.NotificationLarkWebhookDelivery, *common.PageReply, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).NotificationLarkWebhookDelivery.Query()
	query = r.getLarkWebhookQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*model.NotificationLarkWebhookDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationLarkWebhookDelivery{
			ID:            item.ID,
			EventID:       item.EventID,
			EventType:     commonenum.EventType(item.EventType),
			WebhookID:     item.WebhookID,
			RequestBody:   item.RequestBody,
			Status:        notifyenum.NotificationChannelStatus(item.Status),
			AttemptCount:  item.AttemptCount,
			LastAttemptAt: item.LastAttemptAt,
			HTTPStatus:    item.HTTPStatus,
			ResponseBody:  item.ResponseBody,
			SentAt:        item.SentAt,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
		})
	}
	return result, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Claim(ctx context.Context, id int64, now time.Time, processingTimeout time.Duration, retryUnknown bool) (bool, error) {
	conditions := []predicate.NotificationLarkWebhookDelivery{
		notificationlarkwebhookdelivery.StatusEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusFailed)),
		notificationlarkwebhookdelivery.And(
			notificationlarkwebhookdelivery.StatusEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusProcessing)),
			notificationlarkwebhookdelivery.Or(
				notificationlarkwebhookdelivery.LastAttemptAtIsNil(),
				notificationlarkwebhookdelivery.LastAttemptAtLTE(now.Add(-processingTimeout)),
			),
		),
	}
	if retryUnknown {
		conditions = append(conditions, notificationlarkwebhookdelivery.StatusEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusUnknown)))
	}
	count, err := r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(
			notificationlarkwebhookdelivery.IDEQ(id),
			notificationlarkwebhookdelivery.Or(conditions...),
		).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusProcessing)).
		SetLastAttemptAt(now).
		AddAttemptCount(1).
		Save(ctx)
	return count > 0, err
}

func (r *NotificationLarkWebhookDeliveryRepo) MarkSucceeded(ctx context.Context, id int64, httpStatus *int, responseBody *string, sentAt time.Time) error {
	return r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(
			notificationlarkwebhookdelivery.IDEQ(id),
			notificationlarkwebhookdelivery.StatusNEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)).
		SetNillableHTTPStatus(httpStatus).
		SetNillableResponseBody(responseBody).
		SetSentAt(sentAt).
		Exec(ctx)
}

func (r *NotificationLarkWebhookDeliveryRepo) MarkFailed(ctx context.Context, id int64, httpStatus *int, responseBody *string) error {
	return r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(
			notificationlarkwebhookdelivery.IDEQ(id),
			notificationlarkwebhookdelivery.StatusNEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusFailed)).
		SetNillableHTTPStatus(httpStatus).
		SetNillableResponseBody(responseBody).
		Exec(ctx)
}

func (r *NotificationLarkWebhookDeliveryRepo) MarkUnknown(ctx context.Context, id int64, httpStatus *int, responseBody *string) error {
	return r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(
			notificationlarkwebhookdelivery.IDEQ(id),
			notificationlarkwebhookdelivery.StatusNEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)),
		).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusUnknown)).
		SetNillableHTTPStatus(httpStatus).
		SetNillableResponseBody(responseBody).
		Exec(ctx)
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
