package usecase

import (
	"im/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

type ChatGroupUsecase struct {
	conf *conf.Bootstrap
	log  *log.Helper
}

func NewChatGroupUsecase(
	conf *conf.Bootstrap,
	logger log.Logger,
) (*ChatGroupUsecase, error) {
	return &ChatGroupUsecase{
		conf: conf,
		log:  log.NewHelper(logger),
	}, nil
}
