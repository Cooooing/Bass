package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationemaildelivery"
	"notify/internal/data/gen/notificationlarkwebhookdelivery"
	"notify/internal/data/gen/notificationtencentsmsdelivery"
	"notify/internal/data/gen/predicate"
	notifyenum "notify/internal/enum"

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

func (r *NotificationEmailDeliveryRepo) SaveOrGet(ctx context.Context, req *bizrepo.NotificationEmailDeliverySaveOrGetReq) (*bizrepo.NotificationEmailDeliverySaveOrGetResponse, error) {
	delivery := req.Delivery
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
		return &bizrepo.NotificationEmailDeliverySaveOrGetResponse{Delivery: emailDeliveryModel(save)}, nil
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
	return &bizrepo.NotificationEmailDeliverySaveOrGetResponse{Delivery: emailDeliveryModel(exist)}, nil
}

func (r *NotificationEmailDeliveryRepo) Get(ctx context.Context, req *bizrepo.NotificationEmailDeliveryGetReq) (*bizrepo.NotificationEmailDeliveryGetResponse, error) {
	listResponse, err := r.List(ctx, &bizrepo.NotificationEmailDeliveryListReq{Query: emailGetQuery(req)})
	if err != nil || len(listResponse.Rows) == 0 {
		return &bizrepo.NotificationEmailDeliveryGetResponse{}, err
	}
	return &bizrepo.NotificationEmailDeliveryGetResponse{Item: listResponse.Rows[0]}, nil
}

func (r *NotificationEmailDeliveryRepo) List(ctx context.Context, req *bizrepo.NotificationEmailDeliveryListReq) (*bizrepo.NotificationEmailDeliveryListResponse, error) {
	query := r.getClient(ctx).NotificationEmailDelivery.Query()
	query = r.getEmailQuery(query, emailListQuery(req))
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationEmailDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, emailDeliveryModel(item))
	}
	return &bizrepo.NotificationEmailDeliveryListResponse{Rows: result}, nil
}

func (r *NotificationEmailDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationEmailDeliveryMapReq) (*bizrepo.NotificationEmailDeliveryMapResponse, error) {
	listResponse, err := r.List(ctx, &bizrepo.NotificationEmailDeliveryListReq{Query: emailMapQuery(req)})
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationEmailDelivery, len(listResponse.Rows))
	for _, item := range listResponse.Rows {
		result[item.ID] = item
	}
	return &bizrepo.NotificationEmailDeliveryMapResponse{Rows: result}, nil
}

func (r *NotificationEmailDeliveryRepo) Count(ctx context.Context, req *bizrepo.NotificationEmailDeliveryCountReq) (*bizrepo.NotificationEmailDeliveryCountResponse, error) {
	query := r.getClient(ctx).NotificationEmailDelivery.Query()
	query = r.getEmailQuery(query, emailCountQuery(req))
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationEmailDeliveryCountResponse{Count: count}, nil
}

func (r *NotificationEmailDeliveryRepo) Page(ctx context.Context, req *bizrepo.NotificationEmailDeliveryPageReq) (*bizrepo.NotificationEmailDeliveryPageResponse, error) {
	queryReq := emailPageQuery(req)
	var pageReq *base.PageRequest
	if queryReq != nil {
		pageReq = queryReq.Page
	}
	page := normalizePage(pageReq)
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
		result = append(result, emailDeliveryModel(item))
	}
	return &bizrepo.NotificationEmailDeliveryPageResponse{
		Rows: result,
		Page: &base.PageResponse{Total: int64(total), Page: page.Page, Size: page.Size},
	}, nil
}

func (r *NotificationEmailDeliveryRepo) Claim(ctx context.Context, req *bizrepo.NotificationEmailDeliveryClaimReq) (*bizrepo.NotificationEmailDeliveryClaimResponse, error) {
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
	return &bizrepo.NotificationEmailDeliveryClaimResponse{Claimed: count > 0}, err
}

func (r *NotificationEmailDeliveryRepo) MarkSucceeded(ctx context.Context, req *bizrepo.NotificationEmailDeliveryMarkSucceededReq) (*bizrepo.NotificationEmailDeliveryMarkSucceededResponse, error) {
	err := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(notificationemaildelivery.IDEQ(req.ID), notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded)).
		SetNillableProviderMessageID(req.ProviderMessageID).
		SetNillableProviderResponse(req.ProviderResponse).
		SetSentAt(req.SentAt).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationEmailDeliveryMarkSucceededResponse{}, nil
}

func (r *NotificationEmailDeliveryRepo) MarkFailed(ctx context.Context, req *bizrepo.NotificationEmailDeliveryMarkFailedReq) (*bizrepo.NotificationEmailDeliveryMarkFailedResponse, error) {
	err := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(notificationemaildelivery.IDEQ(req.ID), notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusFailed)).
		SetNillableProviderResponse(req.ProviderResponse).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationEmailDeliveryMarkFailedResponse{}, nil
}

func (r *NotificationEmailDeliveryRepo) MarkUnknown(ctx context.Context, req *bizrepo.NotificationEmailDeliveryMarkUnknownReq) (*bizrepo.NotificationEmailDeliveryMarkUnknownResponse, error) {
	err := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(notificationemaildelivery.IDEQ(req.ID), notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusUnknown)).
		SetNillableProviderResponse(req.ProviderResponse).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationEmailDeliveryMarkUnknownResponse{}, nil
}

func (r *NotificationEmailDeliveryRepo) MarkRateLimited(ctx context.Context, req *bizrepo.NotificationEmailDeliveryMarkRateLimitedReq) (*bizrepo.NotificationEmailDeliveryMarkRateLimitedResponse, error) {
	err := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(notificationemaildelivery.IDEQ(req.ID), notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusRateLimited)).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationEmailDeliveryMarkRateLimitedResponse{}, nil
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

func (r *NotificationTencentSMSDeliveryRepo) SaveOrGet(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliverySaveOrGetReq) (*bizrepo.NotificationTencentSMSDeliverySaveOrGetResponse, error) {
	delivery := req.Delivery
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
		return &bizrepo.NotificationTencentSMSDeliverySaveOrGetResponse{Delivery: tencentSMSDeliveryModel(save)}, nil
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
	return &bizrepo.NotificationTencentSMSDeliverySaveOrGetResponse{Delivery: tencentSMSDeliveryModel(exist)}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Get(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryGetReq) (*bizrepo.NotificationTencentSMSDeliveryGetResponse, error) {
	listResponse, err := r.List(ctx, &bizrepo.NotificationTencentSMSDeliveryListReq{Query: tencentSMSGetQuery(req)})
	if err != nil || len(listResponse.Rows) == 0 {
		return &bizrepo.NotificationTencentSMSDeliveryGetResponse{}, err
	}
	return &bizrepo.NotificationTencentSMSDeliveryGetResponse{Item: listResponse.Rows[0]}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) List(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryListReq) (*bizrepo.NotificationTencentSMSDeliveryListResponse, error) {
	query := r.getClient(ctx).NotificationTencentSMSDelivery.Query()
	query = r.getTencentSMSQuery(query, tencentSMSListQuery(req))
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationTencentSMSDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, tencentSMSDeliveryModel(item))
	}
	return &bizrepo.NotificationTencentSMSDeliveryListResponse{Rows: result}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryMapReq) (*bizrepo.NotificationTencentSMSDeliveryMapResponse, error) {
	listResponse, err := r.List(ctx, &bizrepo.NotificationTencentSMSDeliveryListReq{Query: tencentSMSMapQuery(req)})
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationTencentSMSDelivery, len(listResponse.Rows))
	for _, item := range listResponse.Rows {
		result[item.ID] = item
	}
	return &bizrepo.NotificationTencentSMSDeliveryMapResponse{Rows: result}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Count(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryCountReq) (*bizrepo.NotificationTencentSMSDeliveryCountResponse, error) {
	query := r.getClient(ctx).NotificationTencentSMSDelivery.Query()
	query = r.getTencentSMSQuery(query, tencentSMSCountQuery(req))
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationTencentSMSDeliveryCountResponse{Count: count}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Page(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryPageReq) (*bizrepo.NotificationTencentSMSDeliveryPageResponse, error) {
	queryReq := tencentSMSPageQuery(req)
	var pageReq *base.PageRequest
	if queryReq != nil {
		pageReq = queryReq.Page
	}
	page := normalizePage(pageReq)
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
		result = append(result, tencentSMSDeliveryModel(item))
	}
	return &bizrepo.NotificationTencentSMSDeliveryPageResponse{
		Rows: result,
		Page: &base.PageResponse{Total: int64(total), Page: page.Page, Size: page.Size},
	}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Claim(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryClaimReq) (*bizrepo.NotificationTencentSMSDeliveryClaimResponse, error) {
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
	return &bizrepo.NotificationTencentSMSDeliveryClaimResponse{Claimed: count > 0}, err
}

func (r *NotificationTencentSMSDeliveryRepo) MarkSucceeded(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryMarkSucceededReq) (*bizrepo.NotificationTencentSMSDeliveryMarkSucceededResponse, error) {
	err := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(notificationtencentsmsdelivery.IDEQ(req.ID), notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)).
		SetNillableProviderRequestID(req.ProviderRequestID).
		SetNillableProviderCode(req.ProviderCode).
		SetNillableProviderMessage(req.ProviderMessage).
		SetSentAt(req.SentAt).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationTencentSMSDeliveryMarkSucceededResponse{}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) MarkFailed(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryMarkFailedReq) (*bizrepo.NotificationTencentSMSDeliveryMarkFailedResponse, error) {
	err := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(notificationtencentsmsdelivery.IDEQ(req.ID), notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusFailed)).
		SetNillableProviderRequestID(req.ProviderRequestID).
		SetNillableProviderCode(req.ProviderCode).
		SetNillableProviderMessage(req.ProviderMessage).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationTencentSMSDeliveryMarkFailedResponse{}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) MarkUnknown(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryMarkUnknownReq) (*bizrepo.NotificationTencentSMSDeliveryMarkUnknownResponse, error) {
	err := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(notificationtencentsmsdelivery.IDEQ(req.ID), notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusUnknown)).
		SetNillableProviderRequestID(req.ProviderRequestID).
		SetNillableProviderCode(req.ProviderCode).
		SetNillableProviderMessage(req.ProviderMessage).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationTencentSMSDeliveryMarkUnknownResponse{}, nil
}

func (r *NotificationTencentSMSDeliveryRepo) MarkRateLimited(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryMarkRateLimitedReq) (*bizrepo.NotificationTencentSMSDeliveryMarkRateLimitedResponse, error) {
	err := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(notificationtencentsmsdelivery.IDEQ(req.ID), notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusRateLimited)).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationTencentSMSDeliveryMarkRateLimitedResponse{}, nil
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

func (r *NotificationLarkWebhookDeliveryRepo) SaveOrGet(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliverySaveOrGetReq) (*bizrepo.NotificationLarkWebhookDeliverySaveOrGetResponse, error) {
	delivery := req.Delivery
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
		return &bizrepo.NotificationLarkWebhookDeliverySaveOrGetResponse{Delivery: larkWebhookDeliveryModel(save)}, nil
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
	return &bizrepo.NotificationLarkWebhookDeliverySaveOrGetResponse{Delivery: larkWebhookDeliveryModel(exist)}, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Get(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryGetReq) (*bizrepo.NotificationLarkWebhookDeliveryGetResponse, error) {
	listResponse, err := r.List(ctx, &bizrepo.NotificationLarkWebhookDeliveryListReq{Query: larkWebhookGetQuery(req)})
	if err != nil || len(listResponse.Rows) == 0 {
		return &bizrepo.NotificationLarkWebhookDeliveryGetResponse{}, err
	}
	return &bizrepo.NotificationLarkWebhookDeliveryGetResponse{Item: listResponse.Rows[0]}, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) List(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryListReq) (*bizrepo.NotificationLarkWebhookDeliveryListResponse, error) {
	query := r.getClient(ctx).NotificationLarkWebhookDelivery.Query()
	query = r.getLarkWebhookQuery(query, larkWebhookListQuery(req))
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationLarkWebhookDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, larkWebhookDeliveryModel(item))
	}
	return &bizrepo.NotificationLarkWebhookDeliveryListResponse{Rows: result}, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryMapReq) (*bizrepo.NotificationLarkWebhookDeliveryMapResponse, error) {
	listResponse, err := r.List(ctx, &bizrepo.NotificationLarkWebhookDeliveryListReq{Query: larkWebhookMapQuery(req)})
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationLarkWebhookDelivery, len(listResponse.Rows))
	for _, item := range listResponse.Rows {
		result[item.ID] = item
	}
	return &bizrepo.NotificationLarkWebhookDeliveryMapResponse{Rows: result}, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Count(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryCountReq) (*bizrepo.NotificationLarkWebhookDeliveryCountResponse, error) {
	query := r.getClient(ctx).NotificationLarkWebhookDelivery.Query()
	query = r.getLarkWebhookQuery(query, larkWebhookCountQuery(req))
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationLarkWebhookDeliveryCountResponse{Count: count}, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Page(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryPageReq) (*bizrepo.NotificationLarkWebhookDeliveryPageResponse, error) {
	queryReq := larkWebhookPageQuery(req)
	var pageReq *base.PageRequest
	if queryReq != nil {
		pageReq = queryReq.Page
	}
	page := normalizePage(pageReq)
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
		result = append(result, larkWebhookDeliveryModel(item))
	}
	return &bizrepo.NotificationLarkWebhookDeliveryPageResponse{
		Rows: result,
		Page: &base.PageResponse{Total: int64(total), Page: page.Page, Size: page.Size},
	}, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Claim(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryClaimReq) (*bizrepo.NotificationLarkWebhookDeliveryClaimResponse, error) {
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
	return &bizrepo.NotificationLarkWebhookDeliveryClaimResponse{Claimed: count > 0}, err
}

func (r *NotificationLarkWebhookDeliveryRepo) MarkSucceeded(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryMarkSucceededReq) (*bizrepo.NotificationLarkWebhookDeliveryMarkSucceededResponse, error) {
	err := r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(notificationlarkwebhookdelivery.IDEQ(req.ID), notificationlarkwebhookdelivery.StatusNEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)).
		SetNillableHTTPStatus(req.HTTPStatus).
		SetNillableResponseBody(req.ResponseBody).
		SetSentAt(req.SentAt).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationLarkWebhookDeliveryMarkSucceededResponse{}, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) MarkFailed(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryMarkFailedReq) (*bizrepo.NotificationLarkWebhookDeliveryMarkFailedResponse, error) {
	err := r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(notificationlarkwebhookdelivery.IDEQ(req.ID), notificationlarkwebhookdelivery.StatusNEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusFailed)).
		SetNillableHTTPStatus(req.HTTPStatus).
		SetNillableResponseBody(req.ResponseBody).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationLarkWebhookDeliveryMarkFailedResponse{}, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) MarkUnknown(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryMarkUnknownReq) (*bizrepo.NotificationLarkWebhookDeliveryMarkUnknownResponse, error) {
	err := r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(notificationlarkwebhookdelivery.IDEQ(req.ID), notificationlarkwebhookdelivery.StatusNEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusUnknown)).
		SetNillableHTTPStatus(req.HTTPStatus).
		SetNillableResponseBody(req.ResponseBody).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NotificationLarkWebhookDeliveryMarkUnknownResponse{}, nil
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

func emailGetQuery(req *bizrepo.NotificationEmailDeliveryGetReq) *bizrepo.NotificationEmailDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func emailListQuery(req *bizrepo.NotificationEmailDeliveryListReq) *bizrepo.NotificationEmailDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func emailMapQuery(req *bizrepo.NotificationEmailDeliveryMapReq) *bizrepo.NotificationEmailDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func emailCountQuery(req *bizrepo.NotificationEmailDeliveryCountReq) *bizrepo.NotificationEmailDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func emailPageQuery(req *bizrepo.NotificationEmailDeliveryPageReq) *bizrepo.NotificationEmailDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}

func tencentSMSGetQuery(req *bizrepo.NotificationTencentSMSDeliveryGetReq) *bizrepo.NotificationTencentSMSDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func tencentSMSListQuery(req *bizrepo.NotificationTencentSMSDeliveryListReq) *bizrepo.NotificationTencentSMSDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func tencentSMSMapQuery(req *bizrepo.NotificationTencentSMSDeliveryMapReq) *bizrepo.NotificationTencentSMSDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func tencentSMSCountQuery(req *bizrepo.NotificationTencentSMSDeliveryCountReq) *bizrepo.NotificationTencentSMSDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func tencentSMSPageQuery(req *bizrepo.NotificationTencentSMSDeliveryPageReq) *bizrepo.NotificationTencentSMSDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}

func larkWebhookGetQuery(req *bizrepo.NotificationLarkWebhookDeliveryGetReq) *bizrepo.NotificationLarkWebhookDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func larkWebhookListQuery(req *bizrepo.NotificationLarkWebhookDeliveryListReq) *bizrepo.NotificationLarkWebhookDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func larkWebhookMapQuery(req *bizrepo.NotificationLarkWebhookDeliveryMapReq) *bizrepo.NotificationLarkWebhookDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func larkWebhookCountQuery(req *bizrepo.NotificationLarkWebhookDeliveryCountReq) *bizrepo.NotificationLarkWebhookDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}
func larkWebhookPageQuery(req *bizrepo.NotificationLarkWebhookDeliveryPageReq) *bizrepo.NotificationLarkWebhookDeliveryQuery {
	if req == nil {
		return nil
	}
	return req.Query
}

func emailDeliveryModel(item *gen.NotificationEmailDelivery) *model.NotificationEmailDelivery {
	if item == nil {
		return nil
	}
	return &model.NotificationEmailDelivery{
		ID: item.ID, EventID: item.EventID, EventType: commonenum.EventType(item.EventType), ReceiverID: item.ReceiverID,
		ToEmail: item.ToEmail, Subject: item.Subject, Body: item.Body, ContentType: item.ContentType,
		Status: notifyenum.NotificationChannelStatus(item.Status), AttemptCount: item.AttemptCount, LastAttemptAt: item.LastAttemptAt,
		ProviderMessageID: item.ProviderMessageID, ProviderResponse: item.ProviderResponse, SentAt: item.SentAt,
		CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt,
	}
}

func tencentSMSDeliveryModel(item *gen.NotificationTencentSMSDelivery) *model.NotificationTencentSMSDelivery {
	if item == nil {
		return nil
	}
	return &model.NotificationTencentSMSDelivery{
		ID: item.ID, EventID: item.EventID, EventType: commonenum.EventType(item.EventType), ReceiverID: item.ReceiverID,
		Phone: item.Phone, SMSSDKAppID: item.SmsSdkAppID, SignName: item.SignName, ProviderTemplateID: item.ProviderTemplateID,
		TemplateParams: item.TemplateParams, Status: notifyenum.NotificationChannelStatus(item.Status), AttemptCount: item.AttemptCount,
		LastAttemptAt: item.LastAttemptAt, ProviderRequestID: item.ProviderRequestID, ProviderCode: item.ProviderCode,
		ProviderMessage: item.ProviderMessage, SentAt: item.SentAt, CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt,
	}
}

func larkWebhookDeliveryModel(item *gen.NotificationLarkWebhookDelivery) *model.NotificationLarkWebhookDelivery {
	if item == nil {
		return nil
	}
	return &model.NotificationLarkWebhookDelivery{
		ID: item.ID, EventID: item.EventID, EventType: commonenum.EventType(item.EventType), WebhookID: item.WebhookID,
		RequestBody: item.RequestBody, Status: notifyenum.NotificationChannelStatus(item.Status), AttemptCount: item.AttemptCount,
		LastAttemptAt: item.LastAttemptAt, HTTPStatus: item.HTTPStatus, ResponseBody: item.ResponseBody, SentAt: item.SentAt,
		CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt,
	}
}
