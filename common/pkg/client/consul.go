package client

import (
	"common/api/gen/common"
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	consulregistry "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	metadata2 "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/durationpb"

	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// ConsulClient 封装 Consul 客户端
type ConsulClient struct {
	conf *common.Consul
	log  *log.Helper

	reg       *consulregistry.Registry
	grpcConns sync.Map // map[string]*grpc.ClientConn
	httpConns sync.Map // map[string]*khttp.Client

	Client *consulapi.Client
}

// NewConsulClient 初始化 Consul 客户端
func NewConsulClient(logger *log.Helper, conf *common.Consul) (*ConsulClient, func(), error) {
	// 默认参数处理
	if conf.DialTimeout == nil {
		conf.DialTimeout = durationpb.New(5 * time.Second)
	}

	// 构造标准 Consul 配置
	cfg := consulapi.DefaultConfig()
	cfg.Address = conf.Address
	if conf.Token != "" {
		cfg.Token = conf.Token
	}
	if conf.Datacenter != "" {
		cfg.Datacenter = conf.Datacenter
	}

	// 自定义底层传输超时，防止 Consul 响应慢卡死进程
	cfg.HttpClient = &http.Client{
		Timeout: conf.DialTimeout.AsDuration(),
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: conf.DialTimeout.AsDuration(),
			}).DialContext,
		},
	}

	client, err := consulapi.NewClient(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("create consul client failed: %w", err)
	}

	// 创建 Kratos 注册器/发现器单例
	reg := consulregistry.New(client,
		consulregistry.WithHealthCheck(true),
		consulregistry.WithHeartbeat(true), // 开启心跳上报
	)

	c := &ConsulClient{
		conf:   conf,
		log:    logger,
		Client: client,
		reg:    reg,
	}

	// 连通性预检 (获取当前 DC 列表)
	_, err = client.Catalog().Datacenters()
	if err != nil {
		return nil, nil, fmt.Errorf("consul connection check failed: %w", err)
	}

	c.log.Infof("consul connected to [%v], datacenter: %v", conf.Address, cfg.Datacenter)
	return c, c.CleanUp, nil
}

// Registrar 返回单例注册器
func (c *ConsulClient) Registrar() registry.Registrar {
	return c.reg
}

// Discovery 返回单例发现器
func (c *ConsulClient) Discovery() registry.Discovery {
	return c.reg
}

// GetGrpcConn 获取并缓存 gRPC 连接 (线程安全)
func (c *ConsulClient) GetGrpcConn(service string) (*grpc.ClientConn, error) {
	if val, ok := c.grpcConns.Load(service); ok {
		return val.(*grpc.ClientConn), nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.conf.DialTimeout.AsDuration())
	defer cancel()

	conn, err := kgrpc.DialInsecure(
		ctx,
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", service)),
		kgrpc.WithDiscovery(c.reg),
		kgrpc.WithTimeout(c.conf.DialTimeout.AsDuration()),
		kgrpc.WithOptions(grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`)),
		kgrpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
			metadata2.Client(),
		),
		kgrpc.WithOptions(grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                60 * time.Second,
			Timeout:             20 * time.Second,
			PermitWithoutStream: true,
		})),
	)
	if err != nil {
		return nil, fmt.Errorf("dial service %s failed: %w", service, err)
	}

	actual, loaded := c.grpcConns.LoadOrStore(service, conn)
	if loaded {
		_ = conn.Close()
		return actual.(*grpc.ClientConn), nil
	}
	return conn, nil
}

// GetHTTPClient 获取并缓存 HTTP Client
func (c *ConsulClient) GetHTTPClient(service string) (*khttp.Client, error) {
	if val, ok := c.httpConns.Load(service); ok {
		return val.(*khttp.Client), nil
	}

	client, err := khttp.NewClient(
		context.Background(),
		khttp.WithEndpoint(fmt.Sprintf("discovery:///%s", service)),
		khttp.WithDiscovery(c.reg),
		khttp.WithMiddleware(
			metadata2.Client(),
			tracing.Client(),
		),
	)
	if err != nil {
		return nil, err
	}

	actual, loaded := c.httpConns.LoadOrStore(service, client)
	if loaded {
		return actual.(*khttp.Client), nil
	}
	return client, nil
}

// CleanUp 资源清理
func (c *ConsulClient) CleanUp() {
	c.grpcConns.Range(func(key, value any) bool {
		if conn, ok := value.(*grpc.ClientConn); ok {
			_ = conn.Close()
			c.log.Infof("closed grpc connection to %v", key)
		}
		return true
	})
	// 重置 Map
	c.grpcConns = sync.Map{}
	c.httpConns = sync.Map{}
	c.log.Info("consul client connections cleaned up")
}

func GetConsulServiceClient[T any](consul *ConsulClient, service string, newClient func(grpc.ClientConnInterface) T) (T, error) {
	conn, err := consul.GetGrpcConn(service)
	if err != nil {
		var zero T
		return zero, err
	}
	return newClient(conn), nil
}
