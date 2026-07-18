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
		return emailDeliveryModel(save), nil
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
	return emailDeliveryModel(exist), nil
}

func (r *NotificationEmailDeliveryRepo) Get(ctx context.Context, req *bizrepo.NotificationEmailDeliveryQuery) (*model.NotificationEmailDelivery, error) {
	list, err := r.List(ctx, emailGetQuery(req))
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationEmailDeliveryRepo) List(ctx context.Context, req *bizrepo.NotificationEmailDeliveryQuery) ([]*model.NotificationEmailDelivery, error) {
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
	return result, nil
}

func (r *NotificationEmailDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationEmailDeliveryQuery) (map[int64]*model.
	NotificationEmailDelivery, error) {
	list, err := r.List(ctx, emailMapQuery(req))
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
	query = r.getEmailQuery(query, emailCountQuery(req))
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationEmailDeliveryRepo) Page(ctx context.Context, req *bizrepo.NotificationEmailDeliveryQuery) (*bizrepo.NotificationEmailDeliveryPageResp, error) {
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
	return &bizrepo.NotificationEmailDeliveryPageResp{
		Rows: result,
		Page: &base.PageResp{Total: int64(total), Page: page.Page, Size: page.Size},
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

func (r *NotificationEmailDeliveryRepo) MarkSucceeded(ctx context.Context, req *bizrepo.NotificationEmailDeliveryMarkSucceededReq) error {
	err := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(notificationemaildelivery.IDEQ(req.ID), notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded)).
		SetNillableProviderMessageID(req.ProviderMessageID).
		SetNillableProviderResp(req.ProviderResp).
		SetSentAt(req.SentAt).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationEmailDeliveryRepo) MarkFailed(ctx context.Context, req *bizrepo.NotificationEmailDeliveryMarkFailedReq) error {
	err := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(notificationemaildelivery.IDEQ(req.ID), notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusFailed)).
		SetNillableProviderResp(req.ProviderResp).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationEmailDeliveryRepo) MarkUnknown(ctx context.Context, req *bizrepo.NotificationEmailDeliveryMarkUnknownReq) error {
	err := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(notificationemaildelivery.IDEQ(req.ID), notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusUnknown)).
		SetNillableProviderResp(req.ProviderResp).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationEmailDeliveryRepo) MarkRateLimited(ctx context.Context, id int64) error {
	err := r.getClient(ctx).NotificationEmailDelivery.Update().
		Where(notificationemaildelivery.IDEQ(id), notificationemaildelivery.StatusNEQ(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationemaildelivery.Status(notifyenum.NotificationChannelStatusRateLimited)).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
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
		return tencentSMSDeliveryModel(save), nil
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
	return tencentSMSDeliveryModel(exist), nil
}

func (r *NotificationTencentSMSDeliveryRepo) Get(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryQuery) (*model.NotificationTencentSMSDelivery, error) {
	list, err := r.List(ctx, tencentSMSGetQuery(req))
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationTencentSMSDeliveryRepo) List(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryQuery) ([]*model.NotificationTencentSMSDelivery, error) {
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
	return result, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryQuery) (map[int64]*model.
	NotificationTencentSMSDelivery, error) {
	list, err := r.List(ctx, tencentSMSMapQuery(req))
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
	query = r.getTencentSMSQuery(query, tencentSMSCountQuery(req))
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationTencentSMSDeliveryRepo) Page(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryQuery) (*bizrepo.NotificationTencentSMSDeliveryPageResp, error) {
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
	return &bizrepo.NotificationTencentSMSDeliveryPageResp{
		Rows: result,
		Page: &base.PageResp{Total: int64(total), Page: page.Page, Size: page.Size},
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

func (r *NotificationTencentSMSDeliveryRepo) MarkSucceeded(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryMarkSucceededReq) error {
	err := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(notificationtencentsmsdelivery.IDEQ(req.ID), notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)).
		SetNillableProviderRequestID(req.ProviderRequestID).
		SetNillableProviderCode(req.ProviderCode).
		SetNillableProviderMessage(req.ProviderMessage).
		SetSentAt(req.SentAt).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationTencentSMSDeliveryRepo) MarkFailed(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryMarkFailedReq) error {
	err := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(notificationtencentsmsdelivery.IDEQ(req.ID), notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusFailed)).
		SetNillableProviderRequestID(req.ProviderRequestID).
		SetNillableProviderCode(req.ProviderCode).
		SetNillableProviderMessage(req.ProviderMessage).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationTencentSMSDeliveryRepo) MarkUnknown(ctx context.Context, req *bizrepo.NotificationTencentSMSDeliveryMarkUnknownReq) error {
	err := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(notificationtencentsmsdelivery.IDEQ(req.ID), notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusUnknown)).
		SetNillableProviderRequestID(req.ProviderRequestID).
		SetNillableProviderCode(req.ProviderCode).
		SetNillableProviderMessage(req.ProviderMessage).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationTencentSMSDeliveryRepo) MarkRateLimited(ctx context.Context, id int64) error {
	err := r.getClient(ctx).NotificationTencentSMSDelivery.Update().
		Where(notificationtencentsmsdelivery.IDEQ(id), notificationtencentsmsdelivery.StatusNEQ(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationtencentsmsdelivery.Status(notifyenum.NotificationChannelStatusRateLimited)).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
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
		SetNillableRespBody(delivery.RespBody).
		SetNillableSentAt(delivery.SentAt).
		Save(ctx)
	if err == nil {
		return larkWebhookDeliveryModel(save), nil
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
	return larkWebhookDeliveryModel(exist), nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Get(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryQuery) (*model.NotificationLarkWebhookDelivery, error) {
	list, err := r.List(ctx, larkWebhookGetQuery(req))
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationLarkWebhookDeliveryRepo) List(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryQuery) ([]*model.NotificationLarkWebhookDelivery, error) {
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
	return result, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Map(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryQuery) (map[int64]*model.
	NotificationLarkWebhookDelivery, error) {
	list, err := r.List(ctx, larkWebhookMapQuery(req))
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
	query = r.getLarkWebhookQuery(query, larkWebhookCountQuery(req))
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationLarkWebhookDeliveryRepo) Page(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryQuery) (*bizrepo.NotificationLarkWebhookDeliveryPageResp, error) {
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
	return &bizrepo.NotificationLarkWebhookDeliveryPageResp{
		Rows: result,
		Page: &base.PageResp{Total: int64(total), Page: page.Page, Size: page.Size},
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

func (r *NotificationLarkWebhookDeliveryRepo) MarkSucceeded(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryMarkSucceededReq) error {
	err := r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(notificationlarkwebhookdelivery.IDEQ(req.ID), notificationlarkwebhookdelivery.StatusNEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded)).
		SetNillableHTTPStatus(req.HTTPStatus).
		SetNillableRespBody(req.RespBody).
		SetSentAt(req.SentAt).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationLarkWebhookDeliveryRepo) MarkFailed(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryMarkFailedReq) error {
	err := r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(notificationlarkwebhookdelivery.IDEQ(req.ID), notificationlarkwebhookdelivery.StatusNEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusFailed)).
		SetNillableHTTPStatus(req.HTTPStatus).
		SetNillableRespBody(req.RespBody).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationLarkWebhookDeliveryRepo) MarkUnknown(ctx context.Context, req *bizrepo.NotificationLarkWebhookDeliveryMarkUnknownReq) error {
	err := r.getClient(ctx).NotificationLarkWebhookDelivery.Update().
		Where(notificationlarkwebhookdelivery.IDEQ(req.ID), notificationlarkwebhookdelivery.StatusNEQ(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusSucceeded))).
		SetStatus(notificationlarkwebhookdelivery.Status(notifyenum.NotificationChannelStatusUnknown)).
		SetNillableHTTPStatus(req.HTTPStatus).
		SetNillableRespBody(req.RespBody).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
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

func emailGetQuery(query *bizrepo.NotificationEmailDeliveryQuery) *bizrepo.NotificationEmailDeliveryQuery {

	return query
}
func emailListQuery(query *bizrepo.NotificationEmailDeliveryQuery) *bizrepo.NotificationEmailDeliveryQuery {

	return query
}
func emailMapQuery(query *bizrepo.NotificationEmailDeliveryQuery) *bizrepo.NotificationEmailDeliveryQuery {

	return query
}
func emailCountQuery(query *bizrepo.NotificationEmailDeliveryQuery) *bizrepo.NotificationEmailDeliveryQuery {

	return query
}
func emailPageQuery(query *bizrepo.NotificationEmailDeliveryQuery) *bizrepo.NotificationEmailDeliveryQuery {

	return query
}

func tencentSMSGetQuery(query *bizrepo.NotificationTencentSMSDeliveryQuery) *bizrepo.NotificationTencentSMSDeliveryQuery {

	return query
}
func tencentSMSListQuery(query *bizrepo.NotificationTencentSMSDeliveryQuery) *bizrepo.NotificationTencentSMSDeliveryQuery {

	return query
}
func tencentSMSMapQuery(query *bizrepo.NotificationTencentSMSDeliveryQuery) *bizrepo.NotificationTencentSMSDeliveryQuery {

	return query
}
func tencentSMSCountQuery(query *bizrepo.NotificationTencentSMSDeliveryQuery) *bizrepo.NotificationTencentSMSDeliveryQuery {

	return query
}
func tencentSMSPageQuery(query *bizrepo.NotificationTencentSMSDeliveryQuery) *bizrepo.NotificationTencentSMSDeliveryQuery {

	return query
}

func larkWebhookGetQuery(query *bizrepo.NotificationLarkWebhookDeliveryQuery) *bizrepo.NotificationLarkWebhookDeliveryQuery {

	return query
}
func larkWebhookListQuery(query *bizrepo.NotificationLarkWebhookDeliveryQuery) *bizrepo.NotificationLarkWebhookDeliveryQuery {

	return query
}
func larkWebhookMapQuery(query *bizrepo.NotificationLarkWebhookDeliveryQuery) *bizrepo.NotificationLarkWebhookDeliveryQuery {

	return query
}
func larkWebhookCountQuery(query *bizrepo.NotificationLarkWebhookDeliveryQuery) *bizrepo.NotificationLarkWebhookDeliveryQuery {

	return query
}
func larkWebhookPageQuery(query *bizrepo.NotificationLarkWebhookDeliveryQuery) *bizrepo.NotificationLarkWebhookDeliveryQuery {

	return query
}

func emailDeliveryModel(item *gen.NotificationEmailDelivery) *model.NotificationEmailDelivery {
	if item == nil {
		return nil
	}
	return &model.NotificationEmailDelivery{
		ID: item.ID, EventID: item.EventID, EventType: commonenum.EventType(item.EventType), ReceiverID: item.ReceiverID,
		ToEmail: item.ToEmail, Subject: item.Subject, Body: item.Body, ContentType: item.ContentType,
		Status: notifyenum.NotificationChannelStatus(item.Status), AttemptCount: item.AttemptCount, LastAttemptAt: item.LastAttemptAt,
		ProviderMessageID: item.ProviderMessageID, ProviderResp: item.ProviderResp, SentAt: item.SentAt,
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
		LastAttemptAt: item.LastAttemptAt, HTTPStatus: item.HTTPStatus, RespBody: item.RespBody, SentAt: item.SentAt,
		CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt,
	}
}
