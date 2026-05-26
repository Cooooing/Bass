package repo

import (
	commonv1 "common/api/gen/common"
	"common/api/gen/common/enums"
	commonClient "common/pkg/client"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	"context"
	"encoding/json"
	"fmt"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationtemplate"
	notifyenum "notify/internal/enum"
	"time"

	utilent "common/pkg/util/ent"
)

var _ repo.NotificationTemplateRepo = (*NotificationTemplateRepo)(nil)

type NotificationTemplateRepo struct {
	db          *gen.Client
	redisClient *commonClient.RedisClient
}

func NewNotificationTemplateRepo(
	db *gen.Client,
	redisClient *commonClient.RedisClient,
) repo.NotificationTemplateRepo {
	r := &NotificationTemplateRepo{
		db:          db,
		redisClient: redisClient,
	}
	return r
}

func (r *NotificationTemplateRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationTemplateRepo) Save(ctx context.Context, u *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	client := r.getClient(ctx)
	if u.Enable {
		_, err := client.NotificationTemplate.Update().
			SetEnable(false).
			Where(notificationtemplate.EventTypeEQ(notificationtemplate.EventType(u.EventType)), notificationtemplate.ChannelEQ(notificationtemplate.Channel(u.Channel))).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	save, err := client.NotificationTemplate.Create().
		SetEventType(notificationtemplate.EventType(u.EventType)).
		SetChannel(notificationtemplate.Channel(u.Channel)).
		SetLanguage(notificationtemplate.Language(u.Language)).
		SetTitle(u.Title).
		SetContent(u.Content).
		SetEnable(u.Enable).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	r.deleteTemplateCaches(ctx, u.EventType)
	return &model.NotificationTemplate{
		ID:        save.ID,
		EventType: commonenum.EventType(save.EventType),
		Channel:   notifyenum.NotificationChannel(save.Channel),
		Language:  notifyenum.Language(save.Language),
		Title:     save.Title,
		Content:   save.Content,
		Enable:    save.Enable,
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}, nil
}

func (r *NotificationTemplateRepo) Update(ctx context.Context, u *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	client := r.getClient(ctx)
	if u.Enable {
		_, err := client.NotificationTemplate.Update().
			SetEnable(false).
			Where(notificationtemplate.EventTypeEQ(notificationtemplate.EventType(u.EventType)), notificationtemplate.ChannelEQ(notificationtemplate.Channel(u.Channel))).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	update, err := client.NotificationTemplate.Update().
		Where(
			notificationtemplate.EventTypeEQ(notificationtemplate.EventType(u.EventType)),
			notificationtemplate.ChannelEQ(notificationtemplate.Channel(u.Channel)),
			notificationtemplate.LanguageEQ(notificationtemplate.Language(u.Language)),
		).
		SetTitle(u.Title).
		SetContent(u.Content).
		SetEnable(u.Enable).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	if update == 0 {
		return nil, nil
	}
	r.deleteTemplateCaches(ctx, u.EventType)
	protoEventType, ok := commonenum.EventTypeMap.ToProto(u.EventType)
	if !ok {
		return nil, nil
	}
	protoChannel, ok := notifyenum.NotificationChannelMap.ToProto(u.Channel)
	if !ok {
		return nil, nil
	}
	return r.Get(ctx, &repo.NotificationTemplateGetReq{
		EventType: new(protoEventType),
		Channel:   new(protoChannel),
		Language:  new(string(u.Language)),
	})
}

func (r *NotificationTemplateRepo) GetTemplates(ctx context.Context, eventType enums.EventType, language string) ([]*model.NotificationTemplate, error) {
	cacheKey := fmt.Sprintf("notify:tpl:%s:%s", eventType, language)

	// 1. 查缓存
	if r.redisClient != nil && r.redisClient.Client != nil {
		cached, err := r.redisClient.Client.Get(ctx, cacheKey).Result()
		if err == nil {
			var list []*model.NotificationTemplate
			if json.Unmarshal([]byte(cached), &list) == nil {
				return list, nil
			}
		}
	}

	// 2. 查 DB
	list, err := r.GetList(ctx, &repo.NotificationTemplateGetReq{
		EventType: &eventType,
		Language:  &language,
		Enable:    new(true),
	})
	if err != nil {
		return nil, err
	}

	// 3. 回填缓存
	if len(list) > 0 && r.redisClient != nil && r.redisClient.Client != nil {
		if b, err := json.Marshal(list); err == nil {
			r.redisClient.Client.Set(ctx, cacheKey, b, time.Hour)
		}
	}

	return list, nil
}

func (r *NotificationTemplateRepo) deleteTemplateCaches(ctx context.Context, eventType commonenum.EventType) {
	if r.redisClient == nil || r.redisClient.Client == nil {
		return
	}
	protoEventType, ok := commonenum.EventTypeMap.ToProto(eventType)
	if !ok {
		return
	}
	for _, language := range notifyenum.LanguageMap.EnumValues() {
		cacheKey := fmt.Sprintf("notify:tpl:%s:%s", protoEventType, language)
		_ = r.redisClient.Client.Del(ctx, cacheKey).Err()
	}
}

func (r *NotificationTemplateRepo) Get(ctx context.Context, req *repo.NotificationTemplateGetReq) (*model.NotificationTemplate, error) {
	query := r.getClient(ctx).NotificationTemplate.Query()
	query = r.getQuery(query, req)
	n, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.NotificationTemplate{
		ID:        n.ID,
		EventType: commonenum.EventType(n.EventType),
		Channel:   notifyenum.NotificationChannel(n.Channel),
		Language:  notifyenum.Language(n.Language),
		Title:     n.Title,
		Content:   n.Content,
		Enable:    n.Enable,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}, nil
}

func (r *NotificationTemplateRepo) GetMap(ctx context.Context, req *repo.NotificationTemplateGetReq) (map[string]*model.NotificationTemplate, error) {
	var err error
	records := make(map[string]*model.NotificationTemplate)
	query := r.getClient(ctx).NotificationTemplate.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		v := &model.NotificationTemplate{
			ID:        list[i].ID,
			EventType: commonenum.EventType(list[i].EventType),
			Channel:   notifyenum.NotificationChannel(list[i].Channel),
			Language:  notifyenum.Language(list[i].Language),
			Title:     list[i].Title,
			Content:   list[i].Content,
			Enable:    list[i].Enable,
			CreatedAt: list[i].CreatedAt,
			UpdatedAt: list[i].UpdatedAt,
		}
		records[v.GetKey()] = v
	}
	return records, nil
}

func (r *NotificationTemplateRepo) GetList(ctx context.Context, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, error) {
	var err error
	records := make([]*model.NotificationTemplate, 0)
	query := r.getClient(ctx).NotificationTemplate.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		records = append(records, &model.NotificationTemplate{
			ID:        list[i].ID,
			EventType: commonenum.EventType(list[i].EventType),
			Channel:   notifyenum.NotificationChannel(list[i].Channel),
			Language:  notifyenum.Language(list[i].Language),
			Title:     list[i].Title,
			Content:   list[i].Content,
			Enable:    list[i].Enable,
			CreatedAt: list[i].CreatedAt,
			UpdatedAt: list[i].UpdatedAt,
		})
	}
	return records, nil
}

func (r *NotificationTemplateRepo) GetPage(ctx context.Context, page *commonv1.PageRequest, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, *commonv1.PageReply, error) {
	var err error
	notificationTemplates := make([]*model.NotificationTemplate, 0)
	page = constant.PageValid(page)
	query := r.getClient(ctx).NotificationTemplate.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	for _, item := range list {
		notificationTemplates = append(notificationTemplates, &model.NotificationTemplate{
			ID:        item.ID,
			EventType: commonenum.EventType(item.EventType),
			Channel:   notifyenum.NotificationChannel(item.Channel),
			Language:  notifyenum.Language(item.Language),
			Title:     item.Title,
			Content:   item.Content,
			Enable:    item.Enable,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	return notificationTemplates, &commonv1.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *NotificationTemplateRepo) getQuery(query *gen.NotificationTemplateQuery, req *repo.NotificationTemplateGetReq) *gen.NotificationTemplateQuery {
	if req.NotificationTemplateId != nil {
		query = query.Where(notificationtemplate.IDEQ(*req.NotificationTemplateId))
	}
	if len(req.NotificationTemplateIds) > 0 {
		query = query.Where(notificationtemplate.IDIn(req.NotificationTemplateIds...))
	}
	if req.EventType != nil {
		dbEventType, _ := commonenum.EventTypeMap.ToEnum(*req.EventType)
		query = query.Where(notificationtemplate.EventTypeEQ(notificationtemplate.EventType(dbEventType)))
	}
	if req.Channel != nil {
		dbChannel, _ := notifyenum.NotificationChannelMap.ToEnum(*req.Channel)
		query = query.Where(notificationtemplate.ChannelEQ(notificationtemplate.Channel(dbChannel)))
	}
	if len(req.Channels) > 0 {
		var channels []notificationtemplate.Channel
		for _, channel := range req.Channels {
			if channel == nil {
				continue
			}
			dbChannel, _ := notifyenum.NotificationChannelMap.ToEnum(*channel)
			channels = append(channels, notificationtemplate.Channel(dbChannel))
		}
		query = query.Where(notificationtemplate.ChannelIn(channels...))
	}
	if req.Language != nil {
		query = query.Where(notificationtemplate.LanguageEQ(notificationtemplate.Language(*req.Language)))
	}
	if req.Enable != nil {
		query = query.Where(notificationtemplate.EnableEQ(*req.Enable))
	}
	return query
}
