package enum

import (
	"common/api/gen/common/enums"
	"common/pkg/enum"
)

type EventType string

const (
	EventTypeUserFollowCreated EventType = "user_follow_created"
	EventTypeUserFollowDeleted EventType = "user_follow_deleted"
	EventTypeArticlePublished  EventType = "article_published"
	EventTypeArticleLiked      EventType = "article_liked"
	EventTypeArticleThanked    EventType = "article_thanked"
	EventTypeArticleCollected  EventType = "article_collected"
	EventTypeArticleWatched    EventType = "article_watched"
	EventTypeCommentPublished  EventType = "comment_published"
	EventTypeCommentLiked      EventType = "comment_liked"
)

var EventTypeMap = enum.NewMapping[EventType, enums.EventType](map[EventType]enum.Entry[EventType, enums.EventType]{
	EventTypeUserFollowCreated: {Proto: enums.EventType_EVENT_TYPE_USER_FOLLOW_CREATED},
	EventTypeUserFollowDeleted: {Proto: enums.EventType_EVENT_TYPE_USER_FOLLOW_DELETED},
	EventTypeArticlePublished:  {Proto: enums.EventType_EVENT_TYPE_ARTICLE_PUBLISHED},
	EventTypeArticleLiked:      {Proto: enums.EventType_EVENT_TYPE_ARTICLE_LIKED},
	EventTypeArticleThanked:    {Proto: enums.EventType_EVENT_TYPE_ARTICLE_THANKED},
	EventTypeArticleCollected:  {Proto: enums.EventType_EVENT_TYPE_ARTICLE_COLLECTED},
	EventTypeArticleWatched:    {Proto: enums.EventType_EVENT_TYPE_ARTICLE_WATCHED},
	EventTypeCommentPublished:  {Proto: enums.EventType_EVENT_TYPE_COMMENT_PUBLISHED},
	EventTypeCommentLiked:      {Proto: enums.EventType_EVENT_TYPE_COMMENT_LIKED},
})
