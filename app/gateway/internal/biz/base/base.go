package base

import (
	"gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseDomain struct {
	Conf *conf.Bootstrap
	Log  *log.Helper
}

func NewBaseDomain(conf *conf.Bootstrap, log *log.Helper) *BaseDomain {
	return &BaseDomain{
		Conf: conf,
		Log:  log,
	}
}
