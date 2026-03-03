package main

import (
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/util"
	"context"
	"flag"
	"fmt"
	"os"

	"signal/internal/conf"
	"signal/internal/conf/bootstrap"
	"signal/internal/server"

	consulconfig "github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	consulapi "github.com/hashicorp/consul/api"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "app"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
	// flagConf is the config flag.
	flagConf = "configs"
	// flagConf is the config flag.
	flagBootstrap = "configs"
)

func init() {
	flag.StringVar(&flagConf, "conf", "configs/config.yaml", "config path for config.yaml")
	flag.StringVar(&flagBootstrap, "bootstrap", "configs/bootstrap.yaml", "config path for bottstrap.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, cc *client.ConsulClient, asynqServer *client.AsynqServer) *kratos.App {
	hostname, _ := os.Hostname()
	id := fmt.Sprintf("%s.%s.%s", hostname, Name, Version)
	log.Infof("start server %s", id)

	go asynqServer.Run()

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(cc.Registrar()),
	)
}

func main() {
	flag.Parse()

	c, bc, confCleanup, err := loadConfig()
	if err != nil {
		panic(err)
	}
	defer confCleanup()

	server.InitMetrics(Name)

	ctx := context.Background()
	shutdownTracing, err := util.SetupTracing(
		ctx,
		Name,
		Version,
		c.Trace.Endpoint,
		c.Trace.EnableOtel,
		c.Trace.Insecure,
		c.Trace.Sampler,
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := shutdownTracing(ctx)
		if err != nil {
			panic(err)
		}
	}()
	logger := util.NewLogger(Name, Version, c.Server.Mode, bc.Log.Level, bc.Log.File)
	app, cleanup, err := wireApp(c, logger, log.NewHelper(logger).WithContext(ctx))
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func loadConfig() (*conf.Bootstrap, *bootstrap.Bootstrap, func(), error) {
	bc, err := loadBootstrap()
	cleanup := func() {} // 默认空函数
	if err != nil {
		return nil, nil, cleanup, err
	}
	Name = bc.Server.Name
	Version = bc.Server.Version

	if bc.Server.Mode == constant.Prod {
		c, err := loadConsulConfig(bc)
		return c, bc, cleanup, err
	}

	c, err := loadLocalConfig(bc)
	if err != nil {
		return nil, nil, cleanup, err
	}

	return c, bc, cleanup, nil
}

func loadBootstrap() (*bootstrap.Bootstrap, error) {
	c := config.New(config.WithSource(env.NewSource(""), file.NewSource(flagBootstrap)))
	defer func(c config.Config) {
		_ = c.Close()
	}(c)

	if err := c.Load(); err != nil {
		return nil, err
	}

	var bc bootstrap.Bootstrap
	if err := c.Scan(&bc); err != nil {
		return nil, err
	}
	return &bc, nil
}

func loadLocalConfig(bc *bootstrap.Bootstrap) (*conf.Bootstrap, error) {
	c := config.New(config.WithSource(env.NewSource(""), file.NewSource(flagConf)))
	defer func(c config.Config) {
		err := c.Close()
		if err != nil {
			panic(err)
		}
	}(c)

	if err := c.Load(); err != nil {
		return nil, fmt.Errorf("load local config fail: %w", err)
	}
	var localConf conf.Bootstrap
	if err := c.Scan(&localConf); err != nil {
		return nil, fmt.Errorf("scan local config fail: %w", err)
	}
	localConf.Server.Name = bc.Server.Name
	localConf.Server.Version = bc.Server.Version
	localConf.Server.Mode = bc.Server.Mode

	return &localConf, nil
}

func loadConsulConfig(bc *bootstrap.Bootstrap) (*conf.Bootstrap, error) {
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
	consulConf := new(conf.Bootstrap)
	if err := c.Scan(consulConf); err != nil {
		return nil, fmt.Errorf("scan consul config fail: %w", err)
	}
	defer func(c config.Config) {
		_ = c.Close()
	}(c)

	consulConf.Server.Name = bc.Server.Name
	consulConf.Server.Version = bc.Server.Version
	consulConf.Server.Mode = bc.Server.Mode
	return consulConf, nil
}
