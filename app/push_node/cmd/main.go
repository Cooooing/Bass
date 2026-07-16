package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	commonClient "common/pkg/client"
	commonserver "common/pkg/server"
	"push_node/internal/biz/usecase"
	"push_node/internal/config"

	"log/slog"

	"github.com/go-kratos/kratos/v3"
	ktransport "github.com/go-kratos/kratos/v3/transport"
	kratosgrpc "github.com/go-kratos/kratos/v3/transport/grpc"
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
	c *config.Bootstrap,
	logger *slog.Logger,
	hs *http.Server,
	gs *kratosgrpc.Server,
	consulClient *commonClient.ConsulClient,
	nodeUc *usecase.NodeUsecase,
) *kratos.App {
	hostname, _ := os.Hostname()
	id := fmt.Sprintf("%s.%s.%s", hostname, Name, Version)
	slog.Info("start server", "id", id)

	servers := []ktransport.Server{hs, gs}

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
		kratos.Registrar(consulClient.Registrar()),
		kratos.AfterStart(func(ctx context.Context) error {
			return nodeUc.ConnectHub(ctx, &usecase.ConnectHubReq{})
		}),
		kratos.BeforeStop(func(ctx context.Context) error {
			// 关闭心跳循环。
			return nodeUc.Stop(ctx, &usecase.StopReq{})
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
	commonServer := c.Server
	logger := commonserver.NewLogger(commonServer, bc.GetLog())

	// 向 push_hub 注册节点。
	hubConn, err := grpc.NewClient(c.PushNode.PushHubAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Sprintf("闁哄鏅濋崑鐐垫暜?push_hub 婵犮垺鍎肩划鍓ф喆? %v", err))
	}
	nodeID, err := usecase.RegisterWithHub(ctx, hubConn, c)
	if err != nil {
		_ = hubConn.Close()
		panic(fmt.Sprintf("濠电偛顦崝宀勫船娴犲鍤嶉柛灞剧矊娴狀垰顭块幆鎵翱閻? %v", err))
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
