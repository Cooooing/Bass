package handler

import (
	"bytes"
	notifyv1 "common/api/notify/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/cutil/handlerchain"
	commonModel "common/pkg/model"
	"context"
	"encoding/json"
	"html/template"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent"
	"notify/internal/data/ent/gen"
)

type FullHandler struct {
	*base.BaseDomain
	*handlerchain.BaseHandler[*commonModel.Notification]
	notificationMetaRepo   repo.NotificationMetaRepo
	notificationRecordRepo repo.NotificationRecordRepo
}

func NewFullHandler(base *base.BaseDomain, notificationMetaRepo repo.NotificationMetaRepo, notificationRecordRepo repo.NotificationRecordRepo) *FullHandler {
	return &FullHandler{
		BaseDomain:             base,
		BaseHandler:            &handlerchain.BaseHandler[*commonModel.Notification]{Name: "full_handler"},
		notificationMetaRepo:   notificationMetaRepo,
		notificationRecordRepo: notificationRecordRepo,
	}
}

func (h *FullHandler) Handle(ctx context.Context, data *commonModel.Notification) (*commonModel.Notification, error) {
	marshal, _ := json.Marshal(data)
	h.Log.Infof("handle notification: %s", marshal)
	switch *data.Type {
	case notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_PUBLISH:
		// Todo 查询关注用户

	case notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_LIKE:
		if data.Meta.Article != nil {
			data.ReceiverIds = append(data.ReceiverIds, data.Meta.Article.CreatedBy)
		}
	case notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_THANK:
		if data.Meta.Article != nil {
			data.ReceiverIds = append(data.ReceiverIds, data.Meta.Article.CreatedBy)
		}
	case notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_COLLECT:
		if data.Meta.Article != nil {
			data.ReceiverIds = append(data.ReceiverIds, data.Meta.Article.CreatedBy)
		}
	case notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_WATCH:
		if data.Meta.Article != nil {
			data.ReceiverIds = append(data.ReceiverIds, data.Meta.Article.CreatedBy)
		}
	case notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_REWARD:
		if data.Meta.Article != nil {
			data.ReceiverIds = append(data.ReceiverIds, data.Meta.Article.CreatedBy)
		}
	case notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_AT:
		if len(data.Meta.AtUsernames) > 0 {
			userServiceClient, err := client.GetServiceClient(h.Etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
			if err != nil {
				return nil, err
			}
			reply, err := userServiceClient.GetList(ctx, &userv1.GetListRequest{Query: &userv1.UserQueryParams{Names: data.Meta.AtUsernames}})
			if err != nil {
				return nil, err
			}
			for _, i := range reply.Users {
				data.ReceiverIds = append(data.ReceiverIds, i.Id)
			}
		}
	case notifyv1.NotificationType_NOTIFICATION_TYPE_COMMENT_PUBLISH:
		if data.Meta.Comment != nil {
			data.ReceiverIds = append(data.ReceiverIds, data.Meta.Comment.ArticleId)
			if data.Meta.Comment.ReplyId != nil {
				data.ReceiverIds = append(data.ReceiverIds, *data.Meta.Comment.ReplyId)
			}
		}
	case notifyv1.NotificationType_NOTIFICATION_TYPE_COMMENT_LIKE:
		if data.Meta.Comment != nil {
			data.ReceiverIds = append(data.ReceiverIds, data.Meta.Comment.CreatedBy)
		}
	case notifyv1.NotificationType_NOTIFICATION_TYPE_COMMENT_THANK:
		if data.Meta.Comment != nil {
			data.ReceiverIds = append(data.ReceiverIds, data.Meta.Comment.CreatedBy)
		}
	case notifyv1.NotificationType_NOTIFICATION_TYPE_COMMENT_COLLECT:
		if data.Meta.Comment != nil {
			data.ReceiverIds = append(data.ReceiverIds, data.Meta.Comment.CreatedBy)
		}
	case notifyv1.NotificationType_NOTIFICATION_TYPE_COMMENT_AT:
		if len(data.Meta.AtUsernames) > 0 {
			userServiceClient, err := client.GetServiceClient(h.Etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
			if err != nil {
				return nil, err
			}
			reply, err := userServiceClient.GetList(ctx, &userv1.GetListRequest{Query: &userv1.UserQueryParams{Names: data.Meta.AtUsernames}})
			if err != nil {
				return nil, err
			}
			for _, i := range reply.Users {
				data.ReceiverIds = append(data.ReceiverIds, i.Id)
			}
		}
	}

	// 按模板渲染
	tpl, err := template.New(data.UUID).Parse(data.Content)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, data.Meta); err != nil {
		return nil, err
	}
	data.ContentRender = buf.String()

	// 持久化
	err = ent.WithTx(ctx, h.Db, func(tx *gen.Client) error {
		meta, err := h.notificationMetaRepo.Save(ctx, tx, &model.NotificationMeta{NotificationMeta: &gen.NotificationMeta{
			UUID:             data.UUID,
			NotificationType: int32(*data.Type),
			SenderID:         data.SenderId,
			Meta:             data.Meta,
			Content:          data.ContentRender,
			Status:           int32(notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL),
		}})
		if err != nil {
			return err
		}
		if len(data.ReceiverIds) > 0 {
			records := make([]*model.NotificationRecord, 0, len(data.ReceiverIds))
			for _, receiverId := range data.ReceiverIds {
				records = append(records, &model.NotificationRecord{
					NotificationRecord: &gen.NotificationRecord{
						NotificationID: meta.ID,
						ReceiverID:     receiverId,
					},
				})
			}
			_, err = h.notificationRecordRepo.Saves(ctx, tx, records)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// 按渠道通知

	return h.BaseHandler.Next(ctx, data)
}

func (h *FullHandler) SetNext(next handlerchain.Handler[*commonModel.Notification]) {
	h.BaseHandler.SetNext(next)
}

func (h *FullHandler) Name() string {
	return h.BaseHandler.GetName()
}
