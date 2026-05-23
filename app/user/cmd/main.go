package main

import (
	commonClient "common/pkg/client"
	"common/pkg/util"
	commonServer "common/pkg/util/server"
	"context"
	"flag"
	"fmt"
	"os"
	"user/internal/conf"
	"user/internal/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// 构建时可通过 -ldflags "-X main.Version=x.y.z" 注入版本。
var (
	// Name 是编译产物名称。
	Name = "app"
	// Version 是编译产物版本。
	Version = "v1.0.0"
	// flagConf 是业务配置文件路径参数。
	flagConf = "configs/config.yaml"
	// flagBootstrap 是启动配置文件路径参数。
	flagBootstrap = "configs/bootstrap.yaml"
)

func init() {
	flag.StringVar(&flagConf, "conf", "configs/config.yaml", "config path for config.yaml")
	flag.StringVar(&flagBootstrap, "bootstrap", "configs/bootstrap.yaml", "config path for bootstrap.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, cc *commonClient.ConsulClient) *kratos.App {
	hostname, _ := os.Hostname()
	id := fmt.Sprintf("%s.%s.%s", hostname, Name, Version)
	log.Infof("start server %s", id)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(gs),
		kratos.Registrar(cc.Registrar()),
	)
}

func main() {
	flag.Parse()

	c, bc, confCleanup, err := commonServer.LoadConfig[conf.Bootstrap](flagBootstrap, flagConf)
	if err != nil {
		panic(err)
	}
	defer confCleanup()
	Name = c.Server.Name
	Version = c.Server.Version

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
	log.SetLogger(logger)
	app, cleanup, err := wireApp(c, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 启动应用并等待停止信号。
	if err := app.Run(); err != nil {
		panic(err)
	}
}
