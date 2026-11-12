package repo

import (
	"common/pkg/client/rpc"
	platformv1 "common/proto/gen/platform/v1"
	"context"
	"user/internal/biz/model"

	bizrepo "user/internal/biz/repo"
)

var _ bizrepo.IPClient = (*IPClient)(nil)

type IPClient struct{ platformClient *rpc.PlatformClient }

func NewIPClient(
	platformClient *rpc.PlatformClient,
) bizrepo.IPClient {
	return &IPClient{
		platformClient: platformClient,
	}
}

func (r *IPClient) Resolve(ctx context.Context, ip string) (*model.IpInfo, error) {
	reply, err := r.platformClient.IPResolution.ResolveIP(ctx, &platformv1.ResolveIP_Req{
		Ip: ip,
	})
	if err != nil {
		return nil, err
	}
	return &model.IpInfo{
		Ip:          ip,
		Country:     reply.GetCountry(),
		Province:    reply.GetProvince(),
		City:        reply.GetCity(),
		ISP:         reply.GetIsp(),
		CountryCode: reply.GetCountryCode(),
	}, nil
}
