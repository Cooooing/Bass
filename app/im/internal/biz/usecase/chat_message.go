package usecase

import (
	"im/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

type ChatMessageUsecase struct {
	conf *conf.Bootstrap
	log  *log.Helper
}

func NewChatMessageUsecase(
	conf *conf.Bootstrap,
	logger log.Logger,
) (*ChatMessageUsecase, error) {
	return &ChatMessageUsecase{
		conf: conf,
		log:  log.NewHelper(logger),
	}, nil
}
