package repo

import (
	commonv1 "common/api/gen/common"
	"common/api/gen/common/enums"
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"context"
	"encoding/json"
	"fmt"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/conf"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationtemplate"
	notifyenum "notify/internal/enum"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.NotificationTemplateRepo = (*NotificationTemplateRepo)(nil)

type NotificationTemplateRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
}

func NewNotificationTemplateRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
) repo.NotificationTemplateRepo {
	r := &NotificationTemplateRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
	}
	return r
}

func (r *NotificationTemplateRepo) Save(ctx context.Context, tx *gen.Client, u *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	if u.Enable {
		_, err := tx.NotificationTemplate.Update().
			SetEnable(false).
			Where(notificationtemplate.EventTypeEQ(u.EventType), notificationtemplate.ChannelEQ(u.Channel)).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	save, err := tx.NotificationTemplate.Create().
		SetEventType(u.EventType).
		SetChannel(u.Channel).
		SetTitle(u.Title).
		SetContent(u.Content).
		SetEnable(u.Enable).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.NotificationTemplate{NotificationTemplate: save}, nil
}

func (r *NotificationTemplateRepo) Update(ctx context.Context, tx *gen.Client, u *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	if u.Enable {
		_, err := tx.NotificationTemplate.Update().
			SetEnable(false).
			Where(notificationtemplate.EventTypeEQ(u.EventType), notificationtemplate.ChannelEQ(u.Channel)).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	update, err := tx.NotificationTemplate.UpdateOne(u.NotificationTemplate).Save(ctx)
	return &model.NotificationTemplate{NotificationTemplate: update}, err
}

func (r *NotificationTemplateRepo) GetTemplates(ctx context.Context, eventType enums.EventType, language string) ([]*model.NotificationTemplate, error) {
	cacheKey := fmt.Sprintf("notify:tpl:%s:%s", eventType, language)

	// 1. 查缓存
	cached, err := r.redis.Client.Get(ctx, cacheKey).Result()
	if err == nil {
		var list []*model.NotificationTemplate
		if json.Unmarshal([]byte(cached), &list) == nil {
			return list, nil
		}
	}

	// 2. 查 DB
	list, err := r.GetList(ctx, r.db, &repo.NotificationTemplateGetReq{
		EventType: &eventType,
		Language:  &language,
		Enable:    new(true),
	})
	if err != nil {
		return nil, err
	}

	// 3. 回填缓存
	if len(list) > 0 {
		if b, err := json.Marshal(list); err == nil {
			r.redis.Client.Set(ctx, cacheKey, b, time.Hour)
		}
	}

	return list, nil
}

func (r *NotificationTemplateRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.NotificationTemplateGetReq) (*model.NotificationTemplate, error) {
	query := tx.NotificationTemplate.Query()
	query = r.getQuery(query, req)
	n, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	return &model.NotificationTemplate{NotificationTemplate: n}, err
}

func (r *NotificationTemplateRepo) GetMap(ctx context.Context, tx *gen.Client, req *repo.NotificationTemplateGetReq) (map[string]*model.NotificationTemplate, error) {
	var err error
	records := make(map[string]*model.NotificationTemplate)
	query := tx.NotificationTemplate.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		v := &model.NotificationTemplate{NotificationTemplate: list[i]}
		records[v.GetKey()] = v
	}
	return records, nil
}

func (r *NotificationTemplateRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, error) {
	var err error
	records := make([]*model.NotificationTemplate, 0)
	query := tx.NotificationTemplate.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		records = append(records, &model.NotificationTemplate{NotificationTemplate: list[i]})
	}
	return records, nil
}

func (r *NotificationTemplateRepo) GetPage(ctx context.Context, tx *gen.Client, page *commonv1.PageRequest, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, *commonv1.PageReply, error) {
	var err error
	notificationTemplates := make([]*model.NotificationTemplate, 0)
	page = constant.PageValid(page)
	query := tx.NotificationTemplate.Query()
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
		notificationTemplates = append(notificationTemplates, &model.NotificationTemplate{NotificationTemplate: item})
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
		dbEventType, _ := notifyenum.EventTypeMap.ToEnum(*req.EventType)
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
