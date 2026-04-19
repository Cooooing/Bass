package base

import (
	"common/pkg/client/rpc"
	"common/pkg/util"
	"connector/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseDomain struct {
	Conf             *conf.Bootstrap
	Log              *log.Helper
	EventPool        *util.EventPool
	SignalNodeClient *rpc.SignalNodeClient
}

func NewBaseDomain(conf *conf.Bootstrap, logger log.Logger, eventPool *util.EventPool, signalNodeClient *rpc.SignalNodeClient) *BaseDomain {
	return &BaseDomain{
		Conf:             conf,
		Log:              log.NewHelper(logger),
		EventPool:        eventPool,
		SignalNodeClient: signalNodeClient,
	}
}
