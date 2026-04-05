package model

import (
	notifyv1 "common/gen/notify/v1"
	"time"
)

type Notification struct {
	UUID       string                          `json:"uuid"`
	Type       *notifyv1.NotificationType      `json:"type"`
	SenderId   int64                           `json:"sender_id"`
	SenderName string                          `json:"sender_name"`
	Channels   []*notifyv1.NotificationChannel `json:"channels"`
	Meta       Meta                            `json:"meta"`
	Status     notifyv1.NotificationStatus     `json:"status"`

	ReceiverIds   []int64                      `json:"-"` // 接收者ID
	Title         string                       `json:"-"` // 模板标题
	Content       string                       `json:"-"` // 模板内容
	ContentRender string                       `json:"-"` // 模板内容渲染结果(持久化)
	Channel       notifyv1.NotificationChannel `json:"-"` // 通知渠道
}

type Meta struct {
	AtUsernames        []string            `json:"at_usernames"`
	User               *UserMeta           `json:"user"`
	Users              []*UserMeta         `json:"users"`
	Article            *ArticleMeta        `json:"article"`
	Comment            *CommentMeta        `json:"comment"`
	RegisterVerifyCode *RegisterVerifyCode `json:"register_verify_code"`
}

type UserMeta struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

type ArticleMeta struct {
	ArticleId     int64  `json:"article_id"`
	Title         string `json:"title"`
	CreatedBy     int64  `json:"created_by"`
	CreatedByName string `json:"created_by_name"`
}

type CommentMeta struct {
	CommentId     int64  `json:"comment_id"`
	ArticleId     int64  `json:"article_id"`
	Content       string `json:"content"`
	ReplyId       *int64 `json:"reply_id"`
	CreatedBy     int64  `json:"created_by"`
	CreatedByName string `json:"created_by_name"`
}

type RegisterVerifyCode struct {
	Email         string        `json:"email"`
	Expire        time.Duration `json:"expire"`
	ExpireMinutes int           `json:"expire_minutes"`

	Phone string `json:"phone"`

	Code string `json:"code"`
}
