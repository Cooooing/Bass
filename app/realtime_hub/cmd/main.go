package main

import (
	commonclient "common/pkg/client"
	commonserver "common/pkg/server"
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"realtime_hub/internal/config"
	"time"

	"github.com/go-kratos/kratos/v3"
	"github.com/go-kratos/kratos/v3/transport"
)

var (
	Name          = "realtime_hub"
	Version       = "v1.0.0"
	flagConf      = "configs/config.yaml"
	flagBootstrap = "configs/bootstrap.yaml"
)

func init() {
	flag.StringVar(&flagConf, "conf", flagConf, "config path")
	flag.StringVar(&flagBootstrap, "bootstrap", flagBootstrap, "bootstrap path")
}
func newApp(config *config.Bootstrap, logger *slog.Logger, servers []transport.Server, consul *commonclient.ConsulClient) *kratos.App {
	host, _ := os.Hostname()
	return kratos.New(kratos.ID(fmt.Sprintf("%s.%s.%s", host, Name, Version)), kratos.Name(Name), kratos.Version(Version), kratos.Logger(logger), kratos.StopTimeout(time.Second), kratos.Server(servers...), kratos.Registrar(consul.Registrar()))
}
func main() {
	flag.Parse()
	config, bootstrap, cleanupConfig, err := config.NewConfig(flagBootstrap, flagConf)
	if err != nil {
		panic(err)
	}
	defer cleanupConfig()
	logger := commonserver.NewLogger(config.GetServer(), bootstrap.GetLog())
	app, cleanup, err := wireApp(config, config.GetServer(), logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}
	_ = context.Background()
}
