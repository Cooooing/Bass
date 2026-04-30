package handler

import (
	commonModel "common/pkg/model"
	"common/pkg/util/handlerchain"
	"context"
)

type Filter struct {
	*handlerchain.BaseHandler[*commonModel.Notification]
}

func NewFilter() *Filter {
	return &Filter{BaseHandler: new(handlerchain.NewBaseHandler[*commonModel.Notification]("filter"))}
}

func (h *Filter) Handle(ctx context.Context, data *commonModel.Notification) (*commonModel.Notification, error) {

	return h.BaseHandler.Next(ctx, data)
}

func (h *Filter) SetNext(next handlerchain.Handler[*commonModel.Notification]) {
	h.BaseHandler.SetNext(next)
}

func (h *Filter) Name() string {
	return h.BaseHandler.Name()
}
