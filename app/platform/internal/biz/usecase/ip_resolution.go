package usecase

import (
	"context"
	"fmt"
	"strings"

	commonModel "common/pkg/model"
	"platform/internal/config"

	"log/slog"

	"github.com/lionsoul2014/ip2region/binding/golang/service"
)

type IpResolutionUsecase struct {
	conf      *config.Bootstrap
	log       *slog.Logger
	ip2region *service.Ip2Region
}

func NewIpResolutionUsecase(conf *config.Bootstrap, logger *slog.Logger) (*IpResolutionUsecase, func(), error) {
	var ip2region *service.Ip2Region
	var cleanup func()
	if conf.Platform.IpData.Enable {
		v4Config, err := service.NewV4Config(service.VIndexCache, conf.Platform.IpData.Ipv4XdbPath, 20)
		if err != nil {
			return nil, nil, fmt.Errorf("create IPv4 config: %w", err)
		}
		v6Config, err := service.NewV6Config(service.VIndexCache, conf.Platform.IpData.Ipv6XdbPath, 20)
		if err != nil {
			return nil, nil, fmt.Errorf("create IPv6 config: %w", err)
		}
		ip2region, err = service.NewIp2Region(v4Config, v6Config)
		if err != nil {
			return nil, nil, fmt.Errorf("create ip2region service: %w", err)
		}
		cleanup = ip2region.Close
	}
	return &IpResolutionUsecase{conf: conf, log: logger, ip2region: ip2region}, cleanup, nil
}

type IpResolutionGetReq struct {
	IP string
}

type IpResolutionGetResponse struct {
	Info *commonModel.IpInfo
}

func (d *IpResolutionUsecase) Get(ctx context.Context, req *IpResolutionGetReq) (*IpResolutionGetResponse, error) {
	if req == nil {
		req = &IpResolutionGetReq{}
	}
	ip := req.IP
	def := "unknown"
	if d.ip2region == nil || ip == "" {
		return &IpResolutionGetResponse{
			Info: &commonModel.IpInfo{Ip: ip, Country: def, Province: def, City: def, ISP: def, CountryCode: def},
		}, nil
	}
	region, err := d.ip2region.SearchByStr(ip)
	if err != nil {
		return nil, fmt.Errorf("resolve IP: %w", err)
	}
	parts := strings.Split(region, "|")
	if len(parts) < 5 {
		return nil, fmt.Errorf("invalid IP region: %s", region)
	}
	clean := func(s string) string {
		if s == "0" {
			return def
		}
		return s
	}
	return &IpResolutionGetResponse{
		Info: &commonModel.IpInfo{Ip: ip, Country: clean(parts[0]), Province: clean(parts[1]), City: clean(parts[2]), ISP: clean(parts[3]), CountryCode: clean(parts[4])},
	}, nil
}
