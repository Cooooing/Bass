package repo

import (
	commonenum "common/pkg/enum"
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
