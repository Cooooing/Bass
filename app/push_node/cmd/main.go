package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	commonClient "common/pkg/client"
	commonserver "common/pkg/server"
	"common/proto/gen/common"
	"push_node/internal/biz/usecase"
	"push_node/internal/conf"

	"log/slog"

	"github.com/go-kratos/kratos/v3"
	ktransport "github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	c *conf.Bootstrap,
	logger *slog.Logger,
	hs *http.Server,
	consulClient *commonClient.ConsulClient,
	nodeUc *usecase.NodeUsecase,
) *kratos.App {
	hostname, _ := os.Hostname()
	id := fmt.Sprintf("%s.%s.%s", hostname, Name, Version)
	slog.Info("start server", "id", id)

	servers := []ktransport.Server{hs}

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
		kratos.Registrar(consulClient.Registrar()),
		kratos.AfterStart(func(ctx context.Context) error {
			return nodeUc.ConnectHub(ctx)
		}),
		kratos.BeforeStop(func(ctx context.Context) error {
			// 鎼存梻鏁ら崑婊勵剾閸撳秵鏌囧鈧?push_hub
			nodeUc.Stop()
			return nil
		}),
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
	commonServer := &common.Server{Name: c.Server.Name, Version: c.Server.Version, Mode: c.Server.Mode}
	logger := commonserver.NewLogger(commonServer, bc.GetLog())

	// 闂冭埖顔?1閿涙碍鏁為崘灞藉煂 push_hub 閼惧嘲褰囬懞鍌滃仯 ID
	hubConn, err := grpc.NewClient(c.Server.PushHubAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Sprintf("鏉╃偞甯?push_hub 婢惰精瑙? %v", err))
	}
	nodeID, err := usecase.RegisterWithHub(ctx, hubConn, c)
	if err != nil {
		_ = hubConn.Close()
		panic(fmt.Sprintf("濞夈劌鍞介懞鍌滃仯婢惰精瑙? %v", err))
	}
	slog.Info("node registered", "node_id", nodeID)

	app, cleanup, err := wireApp(c, commonServer, logger, hubConn, nodeID)
	if err != nil {
		_ = hubConn.Close()
		panic(err)
	}
	defer func() {
		cleanup()
		_ = hubConn.Close()
	}()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
