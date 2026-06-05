package enum

import "common/api/gen/common/enums"

// EventType 定义跨服务事件类型的内部持久化值。
type EventType string

const (
	EventTypeUserRegister          EventType = "user_register"
	EventTypeUserLogin             EventType = "user_login"
	EventTypeUserLogout            EventType = "user_logout"
	EventTypeUserFollow            EventType = "user_follow"
	EventTypeUserUnfollow          EventType = "user_unfollow"
	EventTypeUserBlock             EventType = "user_block"
	EventTypeUserUnblock           EventType = "user_unblock"
	EventTypeUserTfaEnable         EventType = "user_tfa_enable"
	EventTypeUserTfaDisable        EventType = "user_tfa_disable"
	EventTypeContentArticlePublish EventType = "content_article_publish"
	EventTypeContentArticleLike    EventType = "content_article_like"
	EventTypeContentArticleThank   EventType = "content_article_thank"
	EventTypeContentArticleCollect EventType = "content_article_collect"
	EventTypeContentArticleWatch   EventType = "content_article_watch"
	EventTypeContentCommentPublish EventType = "content_comment_publish"
	EventTypeContentCommentLike    EventType = "content_comment_like"
)

// EventTypeMap 将内部事件类型映射到 proto 事件类型。
var EventTypeMap = NewMapping[EventType, enums.EventType](map[EventType]Entry[EventType, enums.EventType]{
	EventTypeUserRegister:          {Proto: enums.EventType_EVENT_TYPE_USER_REGISTER},
	EventTypeUserLogin:             {Proto: enums.EventType_EVENT_TYPE_USER_LOGIN},
	EventTypeUserLogout:            {Proto: enums.EventType_EVENT_TYPE_USER_LOGOUT},
	EventTypeUserFollow:            {Proto: enums.EventType_EVENT_TYPE_USER_FOLLOW},
	EventTypeUserUnfollow:          {Proto: enums.EventType_EVENT_TYPE_USER_UNFOLLOW},
	EventTypeUserBlock:             {Proto: enums.EventType_EVENT_TYPE_USER_BLOCK},
	EventTypeUserUnblock:           {Proto: enums.EventType_EVENT_TYPE_USER_UNBLOCK},
	EventTypeUserTfaEnable:         {Proto: enums.EventType_EVENT_TYPE_USER_TFA_ENABLE},
	EventTypeUserTfaDisable:        {Proto: enums.EventType_EVENT_TYPE_USER_TFA_DISABLE},
	EventTypeContentArticlePublish: {Proto: enums.EventType_EVENT_TYPE_ARTICLE_PUBLISHED},
	EventTypeContentArticleLike:    {Proto: enums.EventType_EVENT_TYPE_ARTICLE_LIKED},
	EventTypeContentArticleThank:   {Proto: enums.EventType_EVENT_TYPE_ARTICLE_THANKED},
	EventTypeContentArticleCollect: {Proto: enums.EventType_EVENT_TYPE_ARTICLE_COLLECTED},
	EventTypeContentArticleWatch:   {Proto: enums.EventType_EVENT_TYPE_ARTICLE_WATCHED},
	EventTypeContentCommentPublish: {Proto: enums.EventType_EVENT_TYPE_COMMENT_PUBLISHED},
	EventTypeContentCommentLike:    {Proto: enums.EventType_EVENT_TYPE_COMMENT_LIKED},
})
