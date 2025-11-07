package biz

import (
	"common/pkg/util"
	"gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	NewBaseDomain,

	util.NewTokenRepo,
)

type BaseDomain struct {
	conf *conf.Bootstrap
	log  *log.Helper
}

func NewBaseDomain(conf *conf.Bootstrap, log *log.Helper) *BaseDomain {
	return &BaseDomain{
		conf: conf,
		log:  log,
	}
}
