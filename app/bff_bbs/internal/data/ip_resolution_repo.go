package data

import (
	"common/pkg/client/rpc"
	commonmodel "common/pkg/model"
	platformv1 "common/proto/gen/platform/v1"
	"context"

	"bff_bbs/internal/biz/repo"
)

var _ repo.IPResolutionClient = (*IPResolutionClient)(nil)

type IPResolutionClient struct {
	platformClient *rpc.PlatformClient
}

func NewIPResolutionClient(platformClient *rpc.PlatformClient) repo.IPResolutionClient {
	return &IPResolutionClient{platformClient: platformClient}
}

func (r *IPResolutionClient) Resolve(ctx context.Context, ip string) (*commonmodel.IpInfo, error) {
	resp, err := r.platformClient.IPResolution.ResolveIP(ctx, &platformv1.ResolveIP_Req{Ip: ip})
	if err != nil {
		return nil, err
	}
	return &commonmodel.IpInfo{
		Ip:          ip,
		Country:     resp.GetCountry(),
		Province:    resp.GetProvince(),
		City:        resp.GetCity(),
		ISP:         resp.GetIsp(),
		CountryCode: resp.GetCountryCode(),
	}, nil
}
