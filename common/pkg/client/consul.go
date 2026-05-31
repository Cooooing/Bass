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

// ConsulClient 管理基于 Consul 的服务注册与发现。
type ConsulClient struct {
	conf *common.Consul
	log  *log.Helper

	// Client 是原始 Consul API 客户端，供高级场景使用。
	Client *consulapi.Client
	reg    *consulregistry.Registry

	// 长生命周期连接池。
	// gRPC 重连由 resolver 和 keepalive 处理。
	grpcConns sync.Map // 服务名 -> *grpc.ClientConn
	httpConns sync.Map // 服务名 -> *khttp.Client

	// 记录已注册服务 ID，关闭时自动注销。
	mu       sync.Mutex
	services []string
}

// NewConsulClient 创建 ConsulClient 并返回清理函数。
func NewConsulClient(logger log.Logger, conf *common.Consul) (*ConsulClient, func(), error) {
	// 默认值
	if conf.DialTimeout == nil {
		conf.DialTimeout = durationpb.New(5 * time.Second)
	}
	if conf.Datacenter == "" {
		conf.Datacenter = "dc1"
	}

	// 构建 Consul API 配置。
	cfg := consulapi.DefaultConfig()
	cfg.Address = conf.Address
	if conf.Token != "" {
		cfg.Token = conf.Token
	}
	if conf.Datacenter != "" {
		cfg.Datacenter = conf.Datacenter
	}
	cfg.HttpClient = &http.Client{
		Timeout: conf.DialTimeout.AsDuration(),
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: conf.DialTimeout.AsDuration(),
			}).DialContext,
		},
	}

	apiClient, err := consulapi.NewClient(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("create consul client: %w", err)
	}

	reg := consulregistry.New(apiClient,
		consulregistry.WithHealthCheck(false),
		consulregistry.WithHeartbeat(true),
		consulregistry.WithHealthCheckInterval(10),
		consulregistry.WithDeregisterCriticalServiceAfter(30),
	)

	// 连接预检查。
	if _, err = apiClient.Catalog().Datacenters(); err != nil {
		return nil, nil, fmt.Errorf("consul unreachable: %w", err)
	}

	c := &ConsulClient{
		conf:   conf,
		log:    log.NewHelper(logger),
		Client: apiClient,
		reg:    reg,
	}

	c.log.Infof("consul connected: %s (dc=%s)", conf.Address, cfg.Datacenter)
	return c, c.Close, nil
}

// Registrar 返回基于 Consul 的 Kratos 注册器。
func (c *ConsulClient) Registrar() registry.Registrar { return c.reg }

// Discovery 返回基于 Consul 的 Kratos 服务发现。
func (c *ConsulClient) Discovery() registry.Discovery { return c.reg }

// RegisterService 将服务注册到 Consul。
// 调用 Close 时会自动注销该服务。
func (c *ConsulClient) RegisterService(svc *consulapi.AgentServiceRegistration) error {
	if err := c.Client.Agent().ServiceRegister(svc); err != nil {
		return fmt.Errorf("register %s: %w", svc.ID, err)
	}
	c.mu.Lock()
	c.services = append(c.services, svc.ID)
	c.mu.Unlock()
	c.log.Infof("registered: %s", svc.ID)
	return nil
}

// GetGrpcConn 返回指定服务的缓存 gRPC 连接。
// 连接为长生命周期：gRPC 处理重连，Consul resolver 处理端点更新，keepalive 探测失效连接。
func (c *ConsulClient) GetGrpcConn(service string) (*grpc.ClientConn, error) {
	if v, ok := c.grpcConns.Load(service); ok {
		return v.(*grpc.ClientConn), nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.conf.DialTimeout.AsDuration())
	defer cancel()

	conn, err := kgrpc.DialInsecure(ctx,
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", service)),
		kgrpc.WithDiscovery(c.reg),
		kgrpc.WithTimeout(c.conf.DialTimeout.AsDuration()),
		kgrpc.WithOptions(
			grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                30 * time.Second,
				Timeout:             10 * time.Second,
				PermitWithoutStream: false,
			}),
		),
		kgrpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
			metadata2.Client(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("dial %s: %w", service, err)
	}

	actual, loaded := c.grpcConns.LoadOrStore(service, conn)
	if loaded {
		_ = conn.Close()
		return actual.(*grpc.ClientConn), nil
	}

	c.log.Infof("grpc conn: %s", service)
	return conn, nil
}

// GetHTTPClient 返回指定服务的缓存 HTTP 客户端。
func (c *ConsulClient) GetHTTPClient(service string) (*khttp.Client, error) {
	if v, ok := c.httpConns.Load(service); ok {
		return v.(*khttp.Client), nil
	}

	cl, err := khttp.NewClient(context.Background(),
		khttp.WithEndpoint(fmt.Sprintf("discovery:///%s", service)),
		khttp.WithDiscovery(c.reg),
		khttp.WithMiddleware(
			metadata2.Client(),
			tracing.Client(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("http client %s: %w", service, err)
	}

	actual, loaded := c.httpConns.LoadOrStore(service, cl)
	if loaded {
		_ = cl.Close()
		return actual.(*khttp.Client), nil
	}

	c.log.Infof("http client: %s", service)
	return cl, nil
}

// Close 注销所有服务并关闭所有连接。
func (c *ConsulClient) Close() {
	// 注销服务。
	c.mu.Lock()
	ids := c.services
	c.services = nil
	c.mu.Unlock()

	for _, id := range ids {
		if err := c.Client.Agent().ServiceDeregister(id); err != nil {
			c.log.Warnf("deregister %s: %v", id, err)
		} else {
			c.log.Infof("deregistered: %s", id)
		}
	}

	// 关闭 gRPC 连接。
	c.grpcConns.Range(func(_, v any) bool {
		_ = v.(*grpc.ClientConn).Close()
		return true
	})

	// 关闭 HTTP 客户端。
	c.httpConns.Range(func(_, v any) bool {
		_ = v.(*khttp.Client).Close()
		return true
	})

	c.log.Info("consul client closed")
}
