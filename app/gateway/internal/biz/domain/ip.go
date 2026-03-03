package domain

import (
	commonModel "common/pkg/model"
	"context"
	"fmt"
	"strings"

	domainbase "gateway/internal/biz/base"

	"github.com/lionsoul2014/ip2region/binding/golang/service"
)

type IpDomain struct {
	*domainbase.BaseDomain
	ip2region *service.Ip2Region
}

func NewIpDomain(baseDomain *domainbase.BaseDomain) (*IpDomain, func(), error) {
	var ip2region *service.Ip2Region
	var cleanup func()
	if baseDomain.Conf.Server.IpData.Enable {
		v4Config, err := service.NewV4Config(service.VIndexCache, baseDomain.Conf.Server.IpData.Ipv4XdbPath, 20)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create v4 config: %s", err)
		}
		v6Config, err := service.NewV6Config(service.VIndexCache, baseDomain.Conf.Server.IpData.Ipv6XdbPath, 20)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create v6 config: %s", err)
		}
		ip2region, err = service.NewIp2Region(v4Config, v6Config)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create ip2region service: %s", err)
		}
		cleanup = ip2region.Close
	}
	return &IpDomain{
		BaseDomain: baseDomain,
		ip2region:  ip2region,
	}, cleanup, nil
}

func (d *IpDomain) GetInfo(ctx context.Context, ip string) (*commonModel.IpInfo, error) {
	def := "未知"
	if d.ip2region == nil || ip == "" {
		return &commonModel.IpInfo{
			Ip:          ip,
			Country:     def,
			Province:    def,
			City:        def,
			ISP:         def,
			CountryCode: def,
		}, nil
	}
	region, err := d.ip2region.SearchByStr(ip)

	parts := strings.Split(region, "|")
	if len(parts) < 5 {
		return nil, fmt.Errorf("failed to parse ip region: %s", err)
	}

	// 清洗数据函数
	clean := func(s string) string {
		if s == "0" {
			return def
		}
		return s
	}

	return &commonModel.IpInfo{
		Ip:          ip,
		Country:     clean(parts[0]),
		Province:    clean(parts[1]),
		City:        clean(parts[2]),
		ISP:         clean(parts[3]),
		CountryCode: clean(parts[4]),
	}, nil
}
