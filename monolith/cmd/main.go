package main

import (
	commonserver "common/pkg/server"
	"flag"
	"fmt"
	"log/slog"
	"monolith/internal/config"
	"os"

	"github.com/go-kratos/kratos/v3"
	"github.com/go-kratos/kratos/v3/transport"
)

var (
	Name    = "monolith"
	Version = "v1.0.0"

	flagConf      = "configs/config.yaml"
	flagBootstrap = "configs/bootstrap.yaml"
)

func init() {
	flag.StringVar(&flagConf, "conf", flagConf, "config path for config.yaml")
	flag.StringVar(&flagBootstrap, "bootstrap", flagBootstrap, "config path for bootstrap.yaml")
}

func newApp(c *config.Bootstrap, logger *slog.Logger, servers []transport.Server) *kratos.App {
	hostname, _ := os.Hostname()
	id := fmt.Sprintf("%s.%s.%s", hostname, Name, Version)
	slog.Info("start monolith", "id", id)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
	)
}

func main() {
	flag.Parse()

	c, bc, confCleanup, err := config.LoadConfig(flagBootstrap, flagConf)
	if err != nil {
		panic(err)
	}
	defer confCleanup()

	if c.GetServer().GetName() != "" {
		Name = c.GetServer().GetName()
	}
	if c.GetServer().GetVersion() != "" {
		Version = c.GetServer().GetVersion()
	}

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
