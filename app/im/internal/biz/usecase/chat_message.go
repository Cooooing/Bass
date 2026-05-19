package usecase

import (
	"im/internal/conf"
	"im/internal/data/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type ChatMessageUsecase struct {
	conf *conf.Bootstrap
	log  *log.Helper
	db   *gen.Client
}

func NewChatMessageUsecase(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
) (*ChatMessageUsecase, error) {
	return &ChatMessageUsecase{
		conf: conf,
		log:  log.NewHelper(logger),
		db:   db,
	}, nil
}
