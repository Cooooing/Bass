package server

import (
	"common/api/gen/common"
	"common/pkg/constant"
	"fmt"

	consulconfig "github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	consulapi "github.com/hashicorp/consul/api"
)

func LoadConfig[T any](bootstrapPath string, path string) (*T, *common.Bootstrap, func(), error) {
	bc, err := loadBootstrap(bootstrapPath)
	cleanup := func() {} // 默认空函数
	if err != nil {
		return nil, nil, cleanup, err
	}

	if bc.Server.Mode == constant.Prod {
		c, err := loadConsulConfig[T](bc)
		return c, bc, cleanup, err
	}

	c, err := loadLocalConfig[T](bc, path)
	if err != nil {
		return nil, nil, cleanup, err
	}

	return c, bc, cleanup, nil
}

func loadBootstrap(path string) (*common.Bootstrap, error) {
	c := config.New(config.WithSource(env.NewSource(""), file.NewSource(path)))
	defer func(c config.Config) {
		_ = c.Close()
	}(c)

	if err := c.Load(); err != nil {
		return nil, err
	}

	var bc common.Bootstrap
	if err := c.Scan(&bc); err != nil {
		return nil, err
	}
	return &bc, nil
}

func loadLocalConfig[T any](bc *common.Bootstrap, path string) (*T, error) {
	c := config.New(config.WithSource(env.NewSource(""), file.NewSource(path)))
	defer func(c config.Config) {
		_ = c.Close()
	}(c)

	if err := c.Load(); err != nil {
		return nil, fmt.Errorf("load local config fail: %w", err)
	}
	localConf := new(T)
	if err := c.Scan(localConf); err != nil {
		return nil, fmt.Errorf("scan local config fail: %w", err)
	}

	return localConf, nil
}

func loadConsulConfig[T any](bc *common.Bootstrap) (*T, error) {
	// 初始化 Consul API 客户端配置
	cfg := consulapi.DefaultConfig()
	cfg.Address = bc.Config.Consul.Address
	cfg.Token = bc.Config.Consul.Token

	consulClient, err := consulapi.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("create consul api client fail: %w", err)
	}

	// 创建 Kratos 的 Consul 配置源
	consulSource, err := consulconfig.New(consulClient, consulconfig.WithPath(bc.Config.Consul.Path))
	if err != nil {
		return nil, fmt.Errorf("create consul source fail: %w", err)
	}

	c := config.New(config.WithSource(consulSource))
	if err := c.Load(); err != nil {
		return nil, fmt.Errorf("load config from consul fail: %w", err)
	}
	consulConf := new(T)
	if err := c.Scan(consulConf); err != nil {
		return nil, fmt.Errorf("scan consul config fail: %w", err)
	}
	defer func(c config.Config) {
		_ = c.Close()
	}(c)

	return consulConf, nil
}
