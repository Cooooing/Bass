package repo

import (
	commonv1 "common/api/gen/common"
	"common/api/gen/common/enums"
	v1 "common/api/gen/notify/v1"
	"common/pkg/constant"
	"common/pkg/enum"
	"context"
	"encoding/json"
	"fmt"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	database "notify/internal/data/base"
	"notify/internal/data/ent/gen"
	"notify/internal/data/ent/gen/notificationtemplate"
	"time"
)

type NotificationTemplateRepo struct {
	*database.BaseData
}

func NewNotificationTemplateRepo(repo *database.BaseData) (repo.NotificationTemplateRepo, error) {
	r := &NotificationTemplateRepo{
		BaseData: repo,
	}
	err := r.init()
	return r, err
}

func (r *NotificationTemplateRepo) init() error {
	ctx := context.Background()

	templates := []*model.NotificationTemplate{
		{
			NotificationTemplate: &gen.NotificationTemplate{
				EventType: enum.EventTypeArticlePublished,
				Channel:   enum.NotificationChannelEmail,
				Title:     "欢迎注册",
				Content:   "你好 {{.username}}，欢迎注册！",
				Enable:    true,
			},
		},
	}

	existList, err := r.Db.NotificationTemplate.Query().Where(notificationtemplate.Enable(true)).All(ctx)
	if err != nil {
		return err
	}
	existSet := make(map[string]struct{}, len(existList))
	for _, e := range existList {
		existSet[fmt.Sprintf("%s_%s", e.EventType, e.Channel)] = struct{}{}
	}

	for _, tpl := range templates {
		key := fmt.Sprintf("%s_%s", tpl.EventType, tpl.Channel)
		if _, ok := existSet[key]; ok {
			continue
		}
		if _, err := r.Save(ctx, r.Db, tpl); err != nil {
			return err
		}
	}
	return nil
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

func (r *NotificationTemplateRepo) GetTemplates(ctx context.Context, eventType enums.EventType) ([]*model.NotificationTemplate, error) {
	cacheKey := fmt.Sprintf("notify:tpl:%s", eventType)

	// 1. 查缓存
	cached, err := r.Redis.Client.Get(ctx, cacheKey).Result()
	if err == nil {
		var list []*model.NotificationTemplate
		if json.Unmarshal([]byte(cached), &list) == nil {
			return list, nil
		}
	}

	// 2. 查 DB
	list, err := r.GetList(ctx, r.Db, &repo.NotificationTemplateGetReq{
		EventType: &eventType,
		Enable:    new(true),
	})
	if err != nil {
		return nil, err
	}

	// 3. 回填缓存
	if len(list) > 0 {
		if b, err := json.Marshal(list); err == nil {
			r.Redis.Client.Set(ctx, cacheKey, b, time.Hour)
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
		query = query.Where(notificationtemplate.EventTypeEQ(enum.EventType(req.EventType.String())))
	}
	if req.Channel != nil {
		query = query.Where(notificationtemplate.ChannelEQ(enum.NotificationChannel(v1.NotificationChannel_name[int32(*req.Channel)])))
	}
	if len(req.Channels) > 0 {
		var channels []enum.NotificationChannel
		for _, channel := range req.Channels {
			if channel == nil {
				continue
			}
			channels = append(channels, enum.NotificationChannel(v1.NotificationChannel_name[int32(*channel)]))
		}
		query = query.Where(notificationtemplate.ChannelIn(channels...))
	}
	if req.Enable != nil {
		query = query.Where(notificationtemplate.EnableEQ(*req.Enable))
	}
	return query
}
