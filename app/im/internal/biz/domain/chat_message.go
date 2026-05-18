package domain

import (
	"im/internal/conf"
	"im/internal/data/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type ChatMessageDomain struct {
	conf *conf.Bootstrap
	log  *log.Helper
	db   *gen.Client
}

func NewChatMessageDomain(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
) (*ChatMessageDomain, error) {
	return &ChatMessageDomain{
		conf: conf,
		log:  log.NewHelper(logger),
		db:   db,
	}, nil
}
