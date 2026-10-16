package client

import (
	"context"
	"fmt"
	"strings"

	utilent "common/pkg/util/ent"
	"common/proto/gen/common"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/dtm-labs/client/dtmgrpc/dtmgimp"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DTMClient struct {
	conf *common.DTM
}

func NewDTMClient(conf *common.DTM) *DTMClient {
	return &DTMClient{conf: conf}
}

// InTCC 创建并执行全局 TCC 事务。
func (c *DTMClient) InTCC(ctx context.Context, fn dtmgrpc.TccGlobalFunc) (string, error) {
	endpoint := c.grpcEndpoint()
	if endpoint == "" || fn == nil {
		return "", fmt.Errorf("dtm tcc request invalid")
	}
	gid, err := c.newGID(ctx, endpoint)
	if err != nil {
		return "", err
	}
	ctx, cancel := c.withRequestTimeout(ctx)
	defer cancel()
	err = dtmgrpc.TccGlobalTransaction2(endpoint, gid, func(tcc *dtmgrpc.TccGrpc) {
		tcc.Context = ctx
		tcc.WaitResult = true
		if c.conf != nil && c.conf.GetTimeout() != nil && c.conf.GetTimeout().AsDuration() > 0 {
			tcc.WithGlobalTransRequestTimeout(int64(c.conf.GetTimeout().AsDuration().Seconds()))
		}
		if c.conf != nil && c.conf.GetTccTimeoutToFail() != nil && c.conf.GetTccTimeoutToFail().AsDuration() > 0 {
			tcc.TimeoutToFail = int64(c.conf.GetTccTimeoutToFail().AsDuration().Seconds())
		}
		if c.conf != nil && c.conf.GetRetryInterval() != nil && c.conf.GetRetryInterval().AsDuration() > 0 {
			tcc.RetryInterval = int64(c.conf.GetRetryInterval().AsDuration().Seconds())
		}
	}, fn)
	return gid, err
}

// BarrierCall 在当前 DTM 分支的屏障事务中执行业务操作。
func (c *DTMClient) BarrierCall(ctx context.Context, store *BarrierStore, busiCall dtmcli.BarrierBusiFunc) error {
	if store == nil || store.DB == nil {
		return fmt.Errorf("barrier store is required")
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return err
	}
	barrier.DBType = store.DBType
	barrier.BarrierTableName = utilent.DTMBarrierTableName
	return barrier.CallWithDB(store.DB, busiCall)
}

func (c *DTMClient) grpcEndpoint() string {
	if c.conf == nil {
		return ""
	}
	return strings.TrimRight(c.conf.GetGrpcEndpoint(), "/")
}

func (c *DTMClient) newGID(ctx context.Context, endpoint string) (string, error) {
	ctx, cancel := c.withRequestTimeout(ctx)
	defer cancel()
	response, err := dtmgimp.MustGetDtmClient(endpoint).NewGid(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	if response.GetGid() == "" {
		return "", fmt.Errorf("dtm new gid empty")
	}
	return response.GetGid(), nil
}

func (c *DTMClient) withRequestTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}
	if c.conf == nil || c.conf.GetTimeout() == nil || c.conf.GetTimeout().AsDuration() <= 0 {
		return ctx, func() {}
	}
	return context.WithTimeout(ctx, c.conf.GetTimeout().AsDuration())
}
