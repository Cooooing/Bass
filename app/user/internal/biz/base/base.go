package base

import (
	"common/pkg/client"
	"common/pkg/util"
	"context"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

// TxRunner is a function that executes fn within a transaction context.
type TxRunner func(ctx context.Context, fn func(ctx context.Context) error) error

type BaseDomain struct {
	Conf      *conf.Bootstrap
	Log       *log.Helper
	TxRunner  TxRunner
	Consul    *client.ConsulClient
	Redis     *client.RedisClient
	EventPool *util.EventPool
}

func NewBaseDomain(
	conf *conf.Bootstrap,
	logger log.Logger,
	txRunner TxRunner,
	consul *client.ConsulClient,
	redis *client.RedisClient,
	eventPool *util.EventPool) *BaseDomain {
	return &BaseDomain{
		Conf:      conf,
		Log:       log.NewHelper(logger),
		TxRunner:  txRunner,
		Consul:    consul,
		Redis:     redis,
		EventPool: eventPool,
	}
}
