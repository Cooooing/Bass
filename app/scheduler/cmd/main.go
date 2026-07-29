package main

import (
	commonClient "common/pkg/client"
	commonserver "common/pkg/server"
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"scheduler/internal/config"
	schedulerserver "scheduler/internal/server"

	"github.com/go-kratos/kratos/v3"
	"github.com/go-kratos/kratos/v3/transport"
)

var (
	Name          = "app"
	Version       = "v1.0.0"
	flagConf      = "configs/config.yaml"
	flagBootstrap = "configs/bootstrap.yaml"
)

func init() {
	flag.StringVar(&flagConf, "conf", "configs/config.yaml", "config path for config.yaml")
	flag.StringVar(&flagBootstrap, "bootstrap", "configs/bootstrap.yaml", "config path for bootstrap.yaml")
}

func newApp(
	c *config.Bootstrap,
	logger *slog.Logger,
	servers []transport.Server,
	cc *commonClient.ConsulClient,
	bootstrapServer *schedulerserver.SchedulerBootstrapServer,
	scheduledTaskConsumerServer *schedulerserver.ScheduledTaskConsumerServer,
	delayedTaskConsumerServer *schedulerserver.DelayedTaskConsumerServer,
) *kratos.App {
	hostname, _ := os.Hostname()
	id := fmt.Sprintf("%s.%s.%s", hostname, Name, Version)
	slog.Info("start server", "id", id)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
		kratos.Registrar(cc.Registrar()),
		kratos.BeforeStart(func(ctx context.Context) error {
			if err := bootstrapServer.Start(ctx); err != nil {
				return err
			}
			if err := scheduledTaskConsumerServer.Start(ctx); err != nil {
				_ = bootstrapServer.Stop(ctx)
				return err
			}
			if err := delayedTaskConsumerServer.Start(ctx); err != nil {
				_ = scheduledTaskConsumerServer.Stop(ctx)
				_ = bootstrapServer.Stop(ctx)
				return err
			}
			return nil
		}),
		kratos.BeforeStop(func(ctx context.Context) error {
			var err error
			if stopErr := delayedTaskConsumerServer.Stop(ctx); stopErr != nil {
				err = stopErr
			}
			if stopErr := scheduledTaskConsumerServer.Stop(ctx); stopErr != nil && err == nil {
				err = stopErr
			}
			if stopErr := bootstrapServer.Stop(ctx); stopErr != nil && err == nil {
				err = stopErr
			}
			return err
		}),
	)
}

func main() {
	flag.Parse()

	c, bc, confCleanup, err := config.LoadConfig(flagBootstrap, flagConf)
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

	if err := app.Run(); err != nil {
		panic(err)
	}
}
