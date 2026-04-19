package main

import (
	commonClient "common/pkg/client"
	"common/pkg/util"
	commonServer "common/pkg/util/server"
	"context"
	"flag"
	"fmt"
	"gateway/internal/conf"
	"gateway/internal/server"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
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

func newApp(logger log.Logger, hs *http.Server, cc *commonClient.ConsulClient) *kratos.App {
	hostname, _ := os.Hostname()
	id := fmt.Sprintf("%s.%s.%s", hostname, Name, Version)
	log.Infof("start server %s", id)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
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
	app, cleanup, err := wireApp(c, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
