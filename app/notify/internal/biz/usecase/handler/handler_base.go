package handler

import (
	"notify/internal/biz/usecase"

	"github.com/go-kratos/kratos/v2/log"
)

// BaseHandler 所有事件处理器的基类
type BaseHandler struct {
	log           *log.Helper
	notifyService *usecase.NotifyUsecase
}

func NewBaseHandler(logger log.Logger, notifyService *usecase.NotifyUsecase) BaseHandler {
	return BaseHandler{
		log:           log.NewHelper(logger),
		notifyService: notifyService,
	}
}
