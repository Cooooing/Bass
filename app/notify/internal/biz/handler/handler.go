package handler

import (
	"common/pkg/cutil/collections/dict"
	"common/pkg/cutil/handlerchain"
	commonModel "common/pkg/model"

	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	NewFullHandler,
	NewRegisterVerifyCode,
)

func ProvideHandlers(
	registerVerifyCode *RegisterVerifyCode,
	fullHandler *FullHandler,
) dict.Map[string, handlerchain.Handler[*commonModel.Notification]] {
	m := dict.New[string, handlerchain.Handler[*commonModel.Notification]](0)
	m.Set(registerVerifyCode.Name(), registerVerifyCode)
	m.Set(fullHandler.Name(), fullHandler)
	return m
}
