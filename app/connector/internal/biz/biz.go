package biz

import (
	"common/pkg/util"
	"connector/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	NewBaseDomain,
	util.NewEventPool,
	NewSessionDomain,
)

type BaseDomain struct {
	conf      *conf.Bootstrap
	log       *log.Helper
	eventPool *util.EventPool
}

func NewBaseDomain(conf *conf.Bootstrap, log *log.Helper, eventPool *util.EventPool) *BaseDomain {
	return &BaseDomain{
		conf:      conf,
		log:       log,
		eventPool: eventPool,
	}
}
