package server

import (
	"common/pkg/constant"
	"common/proto/gen/common"
	"fmt"
	"reflect"

	"log/slog"

	consulconfig "github.com/go-kratos/kratos/contrib/config/consul/v3"
	"github.com/go-kratos/kratos/v3/config"
	"github.com/go-kratos/kratos/v3/config/env"
	"github.com/go-kratos/kratos/v3/config/file"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/protobuf/proto"
)

// LoadConfig 加载配置并保留底层 watcher，用于模块内注册热更新配置。
func LoadConfig[T proto.Message](bootstrapPath string, path string) (T, *common.Bootstrap, *HotConfigManager[T], func(), error) {
	var zero T
	cleanup := func() {}

	// bootstrap 始终从本地读取，业务配置来源再由 bootstrap 中的运行模式决定。
	bc, err := loadBootstrap(bootstrapPath)
	if err != nil {
		return zero, nil, nil, cleanup, err
	}

	var c config.Config
	if bc.Server.Mode != constant.Dev {
		cfg := consulapi.DefaultConfig()
		cfg.Address = bc.Config.Consul.Address
		cfg.Token = bc.Config.Consul.Token

		consulClient, err := consulapi.NewClient(cfg)
		if err != nil {
			return zero, nil, nil, cleanup, err
		}

		consulSource, err := consulconfig.New(consulClient, consulconfig.WithPath(bc.Config.Consul.Path))
		if err != nil {
			return zero, nil, nil, cleanup, err
		}

		c = config.New(config.WithSource(consulSource))
		if err = c.Load(); err != nil {
			_ = c.Close()
			return zero, nil, nil, cleanup, err
		}
	} else {
		c = config.New(config.WithSource(env.NewSource(""), file.NewSource(path)))
		if err = c.Load(); err != nil {
			_ = c.Close()
			return zero, nil, nil, cleanup, err
		}
	}

	typ := reflect.TypeFor[T]()
	if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
		_ = c.Close()
		return zero, nil, nil, cleanup, fmt.Errorf("config type must be pointer to proto message: %s", typ)
	}
	conf, ok := reflect.New(typ.Elem()).Interface().(T)
	if !ok {
		_ = c.Close()
		return zero, nil, nil, cleanup, fmt.Errorf("config type %s does not implement proto.Message", typ)
	}
	if err := c.Scan(conf); err != nil {
		_ = c.Close()
		return zero, nil, nil, cleanup, err
	}

	manager := &HotConfigManager[T]{
		cfg:     c,
		root:    conf,
		entries: make(map[string]*hotConfigEntry),
	}
	cleanup = func() {
		if closeErr := manager.Close(); closeErr != nil {
			slog.Error("close hot config manager fail", "error", closeErr)
		}
	}

	return conf, bc, manager, cleanup, nil
}

func loadBootstrap(path string) (*common.Bootstrap, error) {
	c := config.New(config.WithSource(env.NewSource(""), file.NewSource(path)))
	defer func(c config.Config) {
		_ = c.Close()
	}(c)

	if err := c.Load(); err != nil {
		return nil, err
	}

	var bc common.Bootstrap
	if err := c.Scan(&bc); err != nil {
		return nil, err
	}
	return &bc, nil
}
