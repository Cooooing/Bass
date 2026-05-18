package handler

import (
	"notify/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/log"
)

// BaseHandler 所有事件处理器的基类
type BaseHandler struct {
	log           *log.Helper
	notifyService *domain.NotifyService
}

func NewBaseHandler(logger log.Logger, notifyService *domain.NotifyService) BaseHandler {
	return BaseHandler{
		log:           log.NewHelper(logger),
		notifyService: notifyService,
	}
}
