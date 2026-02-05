package client

import (
	"common/pkg/constant"
	"common/pkg/model"
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	consulregistry "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/durationpb"

	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ConsulClient 封装 Consul 客户端
type ConsulClient struct {
	conf *model.ConsulConf
	log  *log.Helper

	client *consulapi.Client
	reg    *consulregistry.Registry

	grpcConns sync.Map // map[string]*grpc.ClientConn
	httpConns sync.Map // map[string]*khttp.Client
}

// NewConsulClient 初始化 Consul 客户端
func NewConsulClient(logger *log.Helper, conf *model.ConsulConf) (*ConsulClient, func(), error) {
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
		client: client,
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

	// 使用全局 Context 进行 Dial
	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", service)),
		kgrpc.WithDiscovery(c.reg),
		kgrpc.WithTimeout(5*time.Second),
		kgrpc.WithMiddleware(
			tracing.Client(),
			func(next middleware.Handler) middleware.Handler {
				return func(ctx context.Context, req interface{}) (interface{}, error) {
					if token, ok := ctx.Value(constant.CtxToken).(string); ok && token != "" {
						ctx = metadata.AppendToOutgoingContext(ctx,
							strings.ToLower(constant.HeaderAuthentication),
							"Bearer "+token,
						)
					}
					return next(ctx, req)
				}
			},
		),
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
			tracing.Client(),
			func(next middleware.Handler) middleware.Handler {
				return func(ctx context.Context, req interface{}) (interface{}, error) {
					if token, ok := ctx.Value(constant.CtxToken).(string); ok && token != "" {
						if r, ok := req.(khttp.Request); ok {
							r.Header.Set(constant.HeaderAuthentication, "Bearer "+token)
						}
					}
					return next(ctx, req)
				}
			},
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
