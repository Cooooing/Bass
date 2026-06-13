package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	commonClient "common/pkg/client"
	"common/pkg/util"
	"push_node/internal/biz/usecase"
	"push_node/internal/conf"
	"push_node/internal/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func newApp(
	logger log.Logger,
	hs *http.Server,
	consulClient *commonClient.ConsulClient,
	nodeUc *usecase.NodeUsecase,
) *kratos.App {
	hostname, _ := os.Hostname()
	id := fmt.Sprintf("%s.%s.%s", hostname, Name, Version)
	log.Infof("start server %s", id)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(hs),
		kratos.Registrar(consulClient.Registrar()),
		kratos.AfterStart(func(ctx context.Context) error {
			// 应用启动后连接 push_hub 并开始心跳
			return nodeUc.ConnectHub(ctx)
		}),
		kratos.BeforeStop(func(ctx context.Context) error {
			// 应用停止前断开 push_hub
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

	// 阶段 1：注册到 push_hub 获取节点 ID
	hubConn, err := grpc.NewClient(c.Server.PushHubAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Sprintf("连接 push_hub 失败: %v", err))
	}
	nodeID, err := usecase.RegisterWithHub(ctx, hubConn, c)
	if err != nil {
		_ = hubConn.Close()
		panic(fmt.Sprintf("注册节点失败: %v", err))
	}
	log.Infof("节点注册成功: node_id=%s", nodeID)

	// 阶段 2：Wire 初始化所有依赖
	app, cleanup, err := wireApp(c, logger, hubConn, nodeID)
	if err != nil {
		_ = hubConn.Close()
		panic(err)
	}
	defer func() {
		cleanup()
		_ = hubConn.Close()
	}()

	// 启动应用并等待停止信号。
	if err := app.Run(); err != nil {
		panic(err)
	}
}
