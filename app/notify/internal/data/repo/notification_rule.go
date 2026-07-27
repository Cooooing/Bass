package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationrule"
	notifyenum "notify/internal/enum"

	utilent "common/pkg/util/ent"
	entsql "entgo.io/ent/dialect/sql"
)

var _ bizrepo.NotificationRuleRepo = (*NotificationRuleRepo)(nil)

type NotificationRuleRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewNotificationRuleRepo(
	db *gen.Client,
) bizrepo.NotificationRuleRepo {
	return &NotificationRuleRepo{
		db: db,
	}
}

func (r *NotificationRuleRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationRuleRepo) Upsert(ctx context.Context, rule *model.NotificationRule) (*model.NotificationRule, error) {
	err := r.getClient(ctx).NotificationRule.Create().
		SetEventType(notificationrule.EventType(rule.EventType)).
		SetChannel(notificationrule.Channel(rule.Channel)).
		SetLanguage(notificationrule.Language(rule.Language)).
		SetEnabled(rule.Enabled).
		OnConflict(
			entsql.ConflictColumns(notificationrule.FieldEventType, notificationrule.FieldChannel, notificationrule.FieldLanguage),
			entsql.ConflictWhere(entsql.IsNull(notificationrule.FieldDeletedAt)),
		).
		UpdateEnabled().
		UpdateUpdatedAt().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return r.Get(ctx, &bizrepo.NotificationRuleQuery{
		EventType: &rule.EventType,
		Channel:   &rule.Channel,
		Language:  &rule.Language,
	})
}

func (r *NotificationRuleRepo) BulkUpsert(ctx context.Context, rules []*model.NotificationRule) error {
	if len(rules) == 0 {
		return nil
	}
	creates := make([]*gen.NotificationRuleCreate, 0, len(rules))
	for _, rule := range rules {
		if rule == nil {
			continue
		}
		creates = append(creates, r.getClient(ctx).NotificationRule.Create().
			SetEventType(notificationrule.EventType(rule.EventType)).
			SetChannel(notificationrule.Channel(rule.Channel)).
			SetLanguage(notificationrule.Language(rule.Language)).
			SetEnabled(rule.Enabled))
	}
	if len(creates) == 0 {
		return nil
	}
	return r.getClient(ctx).NotificationRule.CreateBulk(creates...).
		OnConflict(
			entsql.ConflictColumns(notificationrule.FieldEventType, notificationrule.FieldChannel, notificationrule.FieldLanguage),
			entsql.ConflictWhere(entsql.IsNull(notificationrule.FieldDeletedAt)),
		).
		UpdateEnabled().
		UpdateUpdatedAt().
		Exec(ctx)
}

func (r *NotificationRuleRepo) Get(ctx context.Context, req *bizrepo.NotificationRuleQuery) (*model.NotificationRule, error) {
	list, err := r.List(ctx, req)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationRuleRepo) List(ctx context.Context, req *bizrepo.NotificationRuleQuery) ([]*model.NotificationRule, error) {
	query := r.getClient(ctx).NotificationRule.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	rules := make([]*model.NotificationRule, 0, len(list))
	for _, item := range list {
		rules = append(rules, &model.NotificationRule{
			ID:        item.ID,
			EventType: commonenum.EventType(item.EventType),
			Channel:   notifyenum.NotificationChannel(item.Channel),
			Language:  notifyenum.Language(item.Language),
			Enabled:   item.Enabled,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	return rules, nil
}

func (r *NotificationRuleRepo) Map(ctx context.Context, req *bizrepo.NotificationRuleQuery) (map[int64]*model.NotificationRule, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationRule, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *NotificationRuleRepo) Count(ctx context.Context, req *bizrepo.NotificationRuleQuery) (int, error) {
	query := r.getClient(ctx).NotificationRule.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationRuleRepo) Page(ctx context.Context, req *bizrepo.NotificationRuleQuery) (*bizrepo.NotificationRulePageResp, error) {
	queryReq := req
	var pageReq *base.PageRequest
	if queryReq != nil {
		pageReq = queryReq.Page
	}
	page := r.normalizePage(pageReq)
	query := r.getClient(ctx).NotificationRule.Query()
	query = r.getQuery(query, queryReq)
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

	rules := make([]*model.NotificationRule, 0, len(list))
	for _, item := range list {
		rules = append(rules, &model.NotificationRule{
			ID:        item.ID,
			EventType: commonenum.EventType(item.EventType),
			Channel:   notifyenum.NotificationChannel(item.Channel),
			Language:  notifyenum.Language(item.Language),
			Enabled:   item.Enabled,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	return &bizrepo.NotificationRulePageResp{
		Rows: rules,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *NotificationRuleRepo) getQuery(query *gen.NotificationRuleQuery, req *bizrepo.NotificationRuleQuery) *gen.NotificationRuleQuery {
	query = query.Where(notificationrule.DeletedAtIsNil())
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(notificationrule.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(notificationrule.IDIn(req.IDs...))
	}
	if req.EventType != nil {
		query = query.Where(notificationrule.EventTypeEQ(notificationrule.EventType(*req.EventType)))
	}
	if len(req.EventTypes) > 0 {
		values := make([]notificationrule.EventType, 0, len(req.EventTypes))
		for _, item := range req.EventTypes {
			values = append(values, notificationrule.EventType(item))
		}
		query = query.Where(notificationrule.EventTypeIn(values...))
	}
	if req.Channel != nil {
		query = query.Where(notificationrule.ChannelEQ(notificationrule.Channel(*req.Channel)))
	}
	if len(req.Channels) > 0 {
		values := make([]notificationrule.Channel, 0, len(req.Channels))
		for _, item := range req.Channels {
			values = append(values, notificationrule.Channel(item))
		}
		query = query.Where(notificationrule.ChannelIn(values...))
	}
	if req.Language != nil {
		query = query.Where(notificationrule.LanguageEQ(notificationrule.Language(*req.Language)))
	}
	if len(req.Languages) > 0 {
		values := make([]notificationrule.Language, 0, len(req.Languages))
		for _, item := range req.Languages {
			values = append(values, notificationrule.Language(item))
		}
		query = query.Where(notificationrule.LanguageIn(values...))
	}
	if req.Enabled != nil {
		query = query.Where(notificationrule.EnabledEQ(*req.Enabled))
	}
	return query
}
