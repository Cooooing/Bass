package data

import (
	"connector/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewBaseRepo,
)

type BaseRepo struct {
	conf *conf.Bootstrap
	log  *log.Helper
}

func NewBaseRepo(conf *conf.Bootstrap, log *log.Helper) *BaseRepo {
	return &BaseRepo{
		conf: conf,
		log:  log,
	}
}
