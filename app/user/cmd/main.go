package main

import (
	"common/pkg"
	"common/pkg/client"
	"flag"
	"fmt"
	"os"
	"time"
	"user/internal/conf"
	"user/internal/conf/bootstrap"
	"user/internal/server"

	"github.com/go-kratos/kratos/contrib/config/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "user"
	// Version is the version of the compiled software.
	Version string = "v1.0.0"
	// flagConf is the config flag.
	flagConf string = "configs"
	// flagConf is the config flag.
	flagBootstrap string = "configs"
)

func init() {
	flag.StringVar(&flagConf, "conf", "configs/config.yaml", "config path for etcd bootstrap")
	flag.StringVar(&flagBootstrap, "bootstrap", "configs/bootstrap.yaml", "config path for bootstrap.yaml")
}

func newApp(logger log.Logger, log *log.Helper, gs *grpc.Server, hs *http.Server, es *client.EtcdClient) *kratos.App {
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
			gs,
			hs,
		),
		kratos.Registrar(es.Registrar()),
	)
}

func main() {
	flag.Parse()

	c, bc, err := loadConfig()
	if err != nil {
		panic(err)
	}

	server.InitMetrics(Name)

	logger := pkg.Logger(c.Server.Mode, bc.Log.Level, bc.Log.File)
	app, cleanup, err := wireApp(c, logger, log.NewHelper(logger))
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func loadConfig() (*conf.Bootstrap, *bootstrap.Bootstrap, error) {
	bc, err := loadBootstrap()
	if err != nil {
		return nil, nil, err
	}
	var c *conf.Bootstrap
	if bc.Mode == "dev" || bc.Mode == "prod" {
		c, err := loadLocalConfig(bc)
		return c, bc, err
	} else {
		c, cli, err := loadEtcdConfig(bc)
		if err != nil {
			panic(err)
		}
		defer func(cli *clientv3.Client) {
			err := cli.Close()
			if err != nil {
				panic(err)
			}
		}(cli)
		return c, bc, nil
	}

	return c, bc, nil
}

func loadBootstrap() (*bootstrap.Bootstrap, error) {
	c := config.New(config.WithSource(file.NewSource(flagBootstrap)))
	defer func(c config.Config) {
		err := c.Close()
		if err != nil {
			panic(err)
		}
	}(c)

	if err := c.Load(); err != nil {
		return nil, fmt.Errorf("load bootstrap.yaml fail: %w", err)
	}

	var bc bootstrap.Bootstrap
	if err := c.Scan(&bc); err != nil {
		return nil, fmt.Errorf("scan bootstrap.yaml fail: %w", err)
	}

	Name = bc.Name
	Version = bc.Version

	return &bc, nil
}

func loadLocalConfig(bc *bootstrap.Bootstrap) (*conf.Bootstrap, error) {
	c := config.New(config.WithSource(file.NewSource(flagConf)))
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
	localConf.Server.Name = bc.Name
	localConf.Server.Version = bc.Version
	localConf.Server.Mode = bc.Mode

	return &localConf, nil
}

func loadEtcdConfig(bc *bootstrap.Bootstrap) (*conf.Bootstrap, *clientv3.Client, error) {
	timeout := time.Second * 5 // 默认
	if bc.Registry.Etcd.Timeout != nil {
		timeout = bc.Registry.Etcd.Timeout.AsDuration()
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   bc.Registry.Etcd.Endpoints,
		Username:    bc.Registry.Etcd.Username,
		Password:    bc.Registry.Etcd.Password,
		DialTimeout: timeout,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("connect etcd fail: %w", err)
	}

	etcdSource, err := etcd.New(cli, etcd.WithPath(bc.Config.Etcd))
	if err != nil {
		return nil, nil, fmt.Errorf("create etcd source fail: %w", err)
	}
	c := config.New(config.WithSource(etcdSource))
	if err := c.Load(); err != nil {
		return nil, nil, fmt.Errorf("load etcd config fail: %w", err)
	}
	var etcdConf conf.Bootstrap
	if err := c.Scan(&etcdConf); err != nil {
		return nil, nil, fmt.Errorf("scan etcd config fail: %w", err)
	}

	err = c.Watch("watch", func(s string, value config.Value) {
		if err := c.Scan(&etcdConf); err != nil {
			log.Errorf(fmt.Errorf("scan etcd config fail: %w", err).Error())
		}
	})
	if err != nil {
		return nil, nil, fmt.Errorf("watch etcd config fail: %w", err)
	}

	etcdConf.Server.Name = bc.Name
	etcdConf.Server.Version = bc.Version
	etcdConf.Server.Mode = bc.Mode

	return &etcdConf, cli, nil
}
