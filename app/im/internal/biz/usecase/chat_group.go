package usecase

import (
	"im/internal/conf"
	"im/internal/data/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type ChatGroupUsecase struct {
	conf *conf.Bootstrap
	log  *log.Helper
	db   *gen.Client
}

func NewChatGroupUsecase(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
) (*ChatGroupUsecase, error) {
	return &ChatGroupUsecase{
		conf: conf,
		log:  log.NewHelper(logger),
		db:   db,
	}, nil
}
