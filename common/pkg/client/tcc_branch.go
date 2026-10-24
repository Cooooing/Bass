package client

import (
	"context"
	"fmt"

	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/protobuf/proto"
)

// TCCBranchCaller 通过服务发现为 DTM 注册可回调的 TCC 分支。
type TCCBranchCaller struct {
	consul *ConsulClient
}

func NewTCCBranchCaller(consul *ConsulClient) *TCCBranchCaller {
	return &TCCBranchCaller{consul: consul}
}

func (c *TCCBranchCaller) Call(ctx context.Context, tcc *dtmgrpc.TccGrpc, service string, request proto.Message, tryMethod, confirmMethod, cancelMethod string, response proto.Message) error {
	if c == nil || c.consul == nil || tcc == nil || service == "" || request == nil || response == nil || tryMethod == "" || confirmMethod == "" || cancelMethod == "" {
		return fmt.Errorf("TCC branch request is invalid")
	}
	endpoint, err := c.consul.GetHealthyGRPCEndpoint(service)
	if err != nil {
		return err
	}
	return tcc.CallBranch(request, endpoint+tryMethod, endpoint+confirmMethod, endpoint+cancelMethod, response)
}
