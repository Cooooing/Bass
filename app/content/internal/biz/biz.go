package biz

import (
	"common/pkg/client"
	"content/internal/conf"
	"content/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	NewBaseDomain,

	NewArticleDomain,
	NewCommentDomain,
	NewDomainDomain,
)

type BaseDomain struct {
	conf     *conf.Bootstrap
	log      *log.Helper
	db       *gen.Client
	rabbitmq *client.RabbitMQClient
}

func NewBaseDomain(conf *conf.Bootstrap, log *log.Helper, db *gen.Client, rabbitmq *client.RabbitMQClient) *BaseDomain {
	return &BaseDomain{
		conf:     conf,
		log:      log,
		db:       db,
		rabbitmq: rabbitmq,
	}
}
