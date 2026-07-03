package main

import (
	commonClient "common/pkg/client"
	commonserver "common/pkg/server"
	"context"
	"flag"
	"fmt"
	"os"
	"user/internal/biz/usecase"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v3"
	ktransport "github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"log/slog"
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

func newApp(c *conf.Bootstrap, logger *slog.Logger, hs *http.Server, gs *grpc.Server, outboxPublisher *usecase.OutboxPublisher, outboxDeadLetterScanner *usecase.OutboxDeadLetterScanner, cc *commonClient.ConsulClient) *kratos.App {
	hostname, _ := os.Hostname()
	id := fmt.Sprintf("%s.%s.%s", hostname, Name, Version)
	slog.Info("start server", "id", id)

	servers := []ktransport.Server{hs, gs, outboxPublisher, outboxDeadLetterScanner}

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
		kratos.Registrar(cc.Registrar()),
	)
}

func main() {
	flag.Parse()

	c, bc, confCleanup, err := conf.LoadConfig(flagBootstrap, flagConf)
	if err != nil {
		panic(err)
	}
	defer confCleanup()
	Name = c.Server.Name
	Version = c.Server.Version

	ctx := context.Background()
	shutdownTracing, err := commonClient.SetupTracing(ctx, Name, Version, c.GetTrace())
	if err != nil {
		panic(err)
	}
	defer func() {
		err := shutdownTracing(ctx)
		if err != nil {
			panic(err)
		}
	}()
	logger := commonserver.NewLogger(c.GetServer(), bc.GetLog())
	app, cleanup, err := wireApp(c, c.GetServer(), logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 启动应用并等待停止信号。
	if err := app.Run(); err != nil {
		panic(err)
	}
}
