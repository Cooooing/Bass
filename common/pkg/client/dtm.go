package client

import (
	"context"
	"fmt"
	"strings"

	"common/proto/gen/common"

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

func (c *DTMClient) GRPCEndpoint() string {
	if c.conf == nil {
		return ""
	}
	return strings.TrimRight(c.conf.GetGrpcEndpoint(), "/")
}

func (c *DTMClient) NewGID(ctx context.Context) (string, error) {
	endpoint := c.GRPCEndpoint()
	if endpoint == "" {
		return "", fmt.Errorf("dtm grpc endpoint empty")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if c.conf != nil && c.conf.GetTimeout() != nil && c.conf.GetTimeout().AsDuration() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.conf.GetTimeout().AsDuration())
		defer cancel()
	}
	resp, err := dtmgimp.MustGetDtmClient(endpoint).NewGid(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	if resp.GetGid() == "" {
		return "", fmt.Errorf("dtm new gid empty")
	}
	return resp.GetGid(), nil
}

func (c *DTMClient) InTCC(ctx context.Context, gid string, fn dtmgrpc.TccGlobalFunc) error {
	endpoint := c.GRPCEndpoint()
	if endpoint == "" {
		return fmt.Errorf("dtm grpc endpoint empty")
	}
	if gid == "" || fn == nil {
		return fmt.Errorf("dtm tcc request invalid")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return dtmgrpc.TccGlobalTransaction2(endpoint, gid, func(tcc *dtmgrpc.TccGrpc) {
		if c.conf != nil && c.conf.GetTimeout() != nil && c.conf.GetTimeout().AsDuration() > 0 {
			tcc.WithGlobalTransRequestTimeout(int64(c.conf.GetTimeout().AsDuration().Seconds()))
		}
		if c.conf != nil && c.conf.GetMaxRetryTimes() > 0 {
			tcc.WithRetryLimit(int64(c.conf.GetMaxRetryTimes()))
		}
		if c.conf != nil && c.conf.GetRetryInterval() != nil && c.conf.GetRetryInterval().AsDuration() > 0 {
			tcc.RetryInterval = int64(c.conf.GetRetryInterval().AsDuration().Seconds())
		}
	}, fn)
}
