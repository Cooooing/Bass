package usecase

import (
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"fmt"
	"strings"

	"platform/internal/conf"

	"github.com/lionsoul2014/ip2region/binding/golang/service"
	"log/slog"
)

// IpResolutionUsecase IP 解析用例，封装 ip2region 库。
type IpResolutionUsecase struct {
	conf      *conf.Bootstrap
	log       *util.LogHelper
	ip2region *service.Ip2Region
}

// NewIpResolutionUsecase 创建 IP 解析用例，返回用例实例和清理函数。
func NewIpResolutionUsecase(
	conf *conf.Bootstrap,
	logger *slog.Logger,
) (*IpResolutionUsecase, func(), error) {
	var ip2region *service.Ip2Region
	var cleanup func()
	if conf.Server.IpData.Enable {
		v4Config, err := service.NewV4Config(service.VIndexCache, conf.Server.IpData.Ipv4XdbPath,
			20)
		if err != nil {
			return nil, nil, fmt.Errorf("创建 IPv4 配置失败: %s", err)
		}
		v6Config, err := service.NewV6Config(service.VIndexCache, conf.Server.IpData.Ipv6XdbPath,
			20)
		if err != nil {
			return nil, nil, fmt.Errorf("创建 IPv6 配置失败: %s", err)
		}
		ip2region, err = service.NewIp2Region(v4Config, v6Config)
		if err != nil {
			return nil, nil, fmt.Errorf("创建 ip2region 服务失败: %s", err)
		}
		cleanup = ip2region.Close
	}
	return &IpResolutionUsecase{
		conf:      conf,
		log:       util.NewLogHelper(logger),
		ip2region: ip2region,
	}, cleanup, nil
}

// Get 根据 IP 地址查询地理信息。
func (d *IpResolutionUsecase) Get(ctx context.Context, ip string) (*commonModel.IpInfo, error) {
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
	if err != nil {
		return nil, fmt.Errorf("IP 解析失败: %s", err)
	}

	parts := strings.Split(region, "|")
	if len(parts) < 5 {
		return nil, fmt.Errorf("IP 解析结果格式异常: %s", region)
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
