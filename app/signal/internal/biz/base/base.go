package base

import (
	"common/pkg/util"
	"signal/internal/conf"
	"signal/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseDomain struct {
	Conf      *conf.Bootstrap
	Log       *log.Helper
	Db        *gen.Client
	EventPool *util.EventPool
}

func NewBaseDomain(conf *conf.Bootstrap, logger log.Logger, db *gen.Client, eventPool *util.EventPool) *BaseDomain {
	return &BaseDomain{
		Conf:      conf,
		Log:       log.NewHelper(logger),
		Db:        db,
		EventPool: eventPool,
	}
}
