package domain

import (
	"im/internal/conf"
	"im/internal/data/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type ChatGroupDomain struct {
	conf *conf.Bootstrap
	log  *log.Helper
	db   *gen.Client
}

func NewChatGroupDomain(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
) (*ChatGroupDomain, error) {
	return &ChatGroupDomain{
		conf: conf,
		log:  log.NewHelper(logger),
		db:   db,
	}, nil
}
