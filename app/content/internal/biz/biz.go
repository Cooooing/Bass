package biz

import (
	"content/internal/conf"
	"content/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	NewBaseDomain,

	NewArticleDomain,
)

type BaseDomain struct {
	conf *conf.Bootstrap
	log  *log.Helper
	db   *gen.Client
}

func NewBaseDomain(conf *conf.Bootstrap, log *log.Helper, db *gen.Client) *BaseDomain {
	return &BaseDomain{
		conf: conf,
		log:  log,
		db:   db,
	}
}
