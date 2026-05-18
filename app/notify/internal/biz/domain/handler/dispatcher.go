package handler

import (
	"common/api/gen/common/enums"
	"context"
)

// EventHandler 事件处理器接口
type EventHandler interface {
	Handle(ctx context.Context, event *enums.Event) error
}

// Dispatcher 事件分发器
type Dispatcher struct {
	handlers       map[enums.EventType]EventHandler
	defaultHandler EventHandler
}

func NewDispatcher(
	follow *FollowHandler,
	articlePublish *ArticlePublishHandler,
	articleAction *ArticleActionHandler,
	comment *CommentHandler,
	commentAction *CommentActionHandler,
	defaultHandler *DefaultHandler,
) *Dispatcher {
	handlers := map[enums.EventType]EventHandler{
		enums.EventType_EVENT_TYPE_USER_FOLLOW_CREATED: follow,
		enums.EventType_EVENT_TYPE_USER_FOLLOW_DELETED: follow,
		enums.EventType_EVENT_TYPE_ARTICLE_PUBLISHED:   articlePublish,
		enums.EventType_EVENT_TYPE_ARTICLE_LIKED:       articleAction,
		enums.EventType_EVENT_TYPE_ARTICLE_THANKED:     articleAction,
		enums.EventType_EVENT_TYPE_ARTICLE_COLLECTED:   articleAction,
		enums.EventType_EVENT_TYPE_ARTICLE_WATCHED:     articleAction,
		enums.EventType_EVENT_TYPE_COMMENT_PUBLISHED:   comment,
		enums.EventType_EVENT_TYPE_COMMENT_LIKED:       commentAction,
	}
	return &Dispatcher{
		handlers:       handlers,
		defaultHandler: defaultHandler,
	}
}

func (d *Dispatcher) Dispatch(ctx context.Context, event *enums.Event) error {
	h, ok := d.handlers[event.Type]
	if !ok {
		h = d.defaultHandler
	}
	return h.Handle(ctx, event)
}
