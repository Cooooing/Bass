package handler

import (
	commonModel "common/pkg/model"
	"common/pkg/util/handlerchain"

	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	NewFullHandler,
	NewRegisterVerifyCode,
)

func ProvideHandlers(
	registerVerifyCode *RegisterVerifyCode,
	fullHandler *FullHandler,
) map[string]handlerchain.Handler[*commonModel.Notification] {
	m := make(map[string]handlerchain.Handler[*commonModel.Notification])
	m[registerVerifyCode.Name()] = registerVerifyCode
	m[fullHandler.Name()] = fullHandler
	return m
}
