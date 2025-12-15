package data

import (
	cv1 "common/api/common/v1"
	v1 "common/api/notify/v1"
	"common/pkg/constant"
	"common/pkg/cutil/base"
	"common/pkg/cutil/collections/dict"
	"context"
	"encoding/json"
	"fmt"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent/gen"
	"notify/internal/data/ent/gen/notificationtemplate"
	"time"
)

type NotificationTemplateRepo struct {
	*BaseRepo
}

func NewNotificationTemplateRepo(repo *BaseRepo) repo.NotificationTemplateRepo {
	return &NotificationTemplateRepo{
		BaseRepo: repo,
	}
}

func (r *NotificationTemplateRepo) Save(ctx context.Context, tx *gen.Client, u *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	if u.Enable {
		_, err := tx.NotificationTemplate.Update().
			SetEnable(false).
			Where(notificationtemplate.NotificationTypeEQ(u.NotificationType), notificationtemplate.ChannelEQ(u.Channel)).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	save, err := tx.NotificationTemplate.Create().
		SetNotificationType(u.NotificationType).
		SetChannel(u.Channel).
		SetContent(u.Content).
		SetProcessors(u.Processors).
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
			Where(notificationtemplate.NotificationTypeEQ(u.NotificationType), notificationtemplate.ChannelEQ(u.Channel)).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	update, err := tx.NotificationTemplate.UpdateOne(u.NotificationTemplate).Save(ctx)
	return &model.NotificationTemplate{NotificationTemplate: update}, err
}

func (r *NotificationTemplateRepo) SaveCache(ctx context.Context, records dict.Map[string, *model.NotificationTemplate]) error {
	redisData := make(map[string]string)
	for key, template := range records.Map() {
		data, err := json.Marshal(template)
		if err != nil {
			return err
		}
		redisData[key] = string(data)
	}
	err := r.redis.Client.HSet(ctx, constant.GetKeyNotificationTemplateMap(), redisData).Err()
	if err != nil {
		return err
	}
	err = r.redis.Client.HExpire(ctx, constant.GetKeyNotificationTemplateMap(), 1*time.Hour, records.Keys()...).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationTemplateRepo) GetCache(ctx context.Context, notificationType *v1.NotificationType, channels []*v1.NotificationChannel) (dict.Map[string, *model.NotificationTemplate], error) {
	keys := make([]string, 0, len(channels))
	for i := range channels {
		key := constant.GetKeyNotificationTemplate(notificationType, channels[i])
		keys = append(keys, key)
	}
	results, err := r.redis.Client.HMGet(ctx, constant.GetKeyNotificationTemplateMap(), keys...).Result()
	if err != nil {
		return nil, err
	}
	// 判断是否全部 nil
	allNil := true
	for _, v := range results {
		if v != nil {
			allNil = false
			break
		}
	}
	if allNil {
		// 如果缓存中没有，则从数据库获取
		templateMap, dbErr := r.GetMap(ctx, r.db, &repo.NotificationTemplateGetReq{
			NotificationType: notificationType,
			Channels:         channels,
			Enable:           base.Ptr(true),
		})
		if dbErr != nil {
			return nil, dbErr
		}
		// 将数据库中查询到的数据存入缓存
		cacheErr := r.SaveCache(ctx, templateMap)
		if cacheErr != nil {
			return nil, cacheErr
		}
		return templateMap, nil
	}

	// 如果缓存中有数据，反序列化并返回
	templateMap := dict.New[string, *model.NotificationTemplate](len(results))
	for i, result := range results {
		if result == nil {
			continue
		}
		template := &model.NotificationTemplate{}
		err = json.Unmarshal([]byte(result.(string)), template)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal template for key %s: %w", keys[i], err)
		}
		templateMap.Set(keys[i], template)
	}

	return templateMap, nil
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

func (r *NotificationTemplateRepo) GetMap(ctx context.Context, tx *gen.Client, req *repo.NotificationTemplateGetReq) (dict.Map[string, *model.NotificationTemplate], error) {
	var err error
	records := dict.New[string, *model.NotificationTemplate](0)
	query := tx.NotificationTemplate.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		v := &model.NotificationTemplate{NotificationTemplate: list[i]}
		records.Set(v.GetKey(), v)
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

func (r *NotificationTemplateRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, *cv1.PageReply, error) {
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
	return notificationTemplates, &cv1.PageReply{
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
	if req.NotificationType != nil {
		query = query.Where(notificationtemplate.NotificationTypeEQ(int32(*req.NotificationType)))
	}
	if req.Channel != nil {
		query = query.Where(notificationtemplate.ChannelEQ(int32(*req.Channel)))
	}
	if len(req.Channels) > 0 {
		var channelsInt32 []int32
		for _, channel := range req.Channels {
			if channel == nil {
				continue
			}
			channelsInt32 = append(channelsInt32, int32(*channel))
		}
		query = query.Where(notificationtemplate.ChannelIn(channelsInt32...))
	}
	if req.Enable != nil {
		query = query.Where(notificationtemplate.EnableEQ(*req.Enable))
	}
	return query
}
