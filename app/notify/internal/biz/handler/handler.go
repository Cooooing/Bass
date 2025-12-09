package handler

import (
	"common/pkg/cutil/collections/dict"
	"common/pkg/cutil/handlerchain"
	commonModel "common/pkg/model"

	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	NewFullHandler,
	NewFilter,
)

func ProvideHandlers(
	filter *Filter,
	fullHandler *FullHandler,
) dict.Map[string, handlerchain.Handler[*commonModel.Notification]] {
	m := dict.New[string, handlerchain.Handler[*commonModel.Notification]](0)
	m.Set(filter.Name(), filter)
	m.Set(fullHandler.Name(), fullHandler)
	return m
}
