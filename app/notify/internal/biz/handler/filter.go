package handler

import (
	"common/pkg/cutil/handlerchain"
	commonModel "common/pkg/model"
	"context"
)

type Filter struct {
	*handlerchain.BaseHandler[*commonModel.Notification]
}

func NewFilter() *Filter {
	return &Filter{BaseHandler: &handlerchain.BaseHandler[*commonModel.Notification]{Name: "filter"}}
}

func (h *Filter) Handle(ctx context.Context, data *commonModel.Notification) (*commonModel.Notification, error) {

	return h.BaseHandler.Next(ctx, data)
}

func (h *Filter) SetNext(next handlerchain.Handler[*commonModel.Notification]) {
	h.BaseHandler.SetNext(next)
}

func (h *Filter) Name() string {
	return h.BaseHandler.GetName()
}
