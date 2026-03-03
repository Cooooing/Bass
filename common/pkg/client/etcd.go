package client

import (
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"

	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	etcdregistry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/durationpb"
)

// EtcdClient 封装了 Etcd 资源与微服务连接池
type EtcdClient struct {
	conf *model.EtcdConf
	log  *log.Helper
	cli  *clientv3.Client
	reg  *etcdregistry.Registry // 单例注册/发现中心

	grpcConns sync.Map // key=serviceName, value=*grpc.ClientConn
	httpConns sync.Map // key=serviceName, value=*khttp.Client
}

// NewEtcdClient 初始化生产级 Etcd 客户端
func NewEtcdClient(logger *log.Helper, conf *model.EtcdConf) (*EtcdClient, func(), error) {
	// 配置默认值补全
	if conf.AutoSyncInterval == nil {
		conf.AutoSyncInterval = durationpb.New(30 * time.Second)
	}
	if conf.DialKeepAliveTime == nil {
		conf.DialKeepAliveTime = durationpb.New(20 * time.Second)
	}
	if conf.DialKeepAliveTimeout == nil {
		conf.DialKeepAliveTimeout = durationpb.New(10 * time.Second)
	}
	if conf.DialTimeout == nil {
		conf.DialTimeout = durationpb.New(5 * time.Second)
	}

	// 构造 Etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            conf.Endpoints,
		Username:             conf.Username,
		Password:             conf.Password,
		DialTimeout:          conf.DialTimeout.AsDuration(),
		AutoSyncInterval:     conf.AutoSyncInterval.AsDuration(),
		DialKeepAliveTime:    conf.DialKeepAliveTime.AsDuration(),
		DialKeepAliveTimeout: conf.DialKeepAliveTimeout.AsDuration(),
		PermitWithoutStream:  conf.PermitWithoutStream,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create etcd client: %w", err)
	}

	// 创建单例 Registry
	reg := etcdregistry.New(cli)

	c := &EtcdClient{
		conf: conf,
		log:  logger,
		cli:  cli,
		reg:  reg,
	}

	// 启动预检
	// 通过 MemberList 确认连接和 Auth 配置是否正确
	ctx, cancel := context.WithTimeout(context.Background(), conf.DialTimeout.AsDuration())
	defer cancel()
	if _, err := cli.MemberList(ctx); err != nil {
		_ = cli.Close()
		return nil, nil, fmt.Errorf("etcd connection check failed: %w", err)
	}

	c.log.Infof("etcd connected to [%v]", conf.Endpoints)
	return c, c.CleanUp, nil
}

// Registrar 获取服务注册器单例
func (c *EtcdClient) Registrar() registry.Registrar {
	return c.reg
}

// Discovery 获取服务发现器单例
func (c *EtcdClient) Discovery() registry.Discovery {
	return c.reg
}

// GetGrpcConn 获取或创建缓存的 gRPC 连接
func (c *EtcdClient) GetGrpcConn(service string) (*grpc.ClientConn, error) {
	if val, ok := c.grpcConns.Load(service); ok {
		return val.(*grpc.ClientConn), nil
	}

	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", service)),
		kgrpc.WithDiscovery(c.reg),
		kgrpc.WithTimeout(c.conf.DialTimeout.AsDuration()),
		// 设置 gRPC 默认负载均衡策略
		kgrpc.WithOptions(grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`)),
		kgrpc.WithMiddleware(
			tracing.Client(),
			c.tokenClientMiddleware(),
		),
		kgrpc.WithOptions(grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                c.conf.DialKeepAliveTime.AsDuration(),
			Timeout:             c.conf.DialKeepAliveTimeout.AsDuration(),
			PermitWithoutStream: c.conf.PermitWithoutStream,
		})),
	)
	if err != nil {
		return nil, fmt.Errorf("dial grpc service %s failed: %w", service, err)
	}

	actual, loaded := c.grpcConns.LoadOrStore(service, conn)
	if loaded {
		_ = conn.Close()
		return actual.(*grpc.ClientConn), nil
	}
	return conn, nil
}

// GetHTTPClient 获取或创建缓存的 HTTP 客户端
func (c *EtcdClient) GetHTTPClient(service string) (*khttp.Client, error) {
	if val, ok := c.httpConns.Load(service); ok {
		return val.(*khttp.Client), nil
	}

	client, err := khttp.NewClient(
		context.Background(),
		khttp.WithEndpoint(fmt.Sprintf("discovery:///%s", service)),
		khttp.WithDiscovery(c.reg),
		khttp.WithTimeout(c.conf.DialTimeout.AsDuration()),
		khttp.WithMiddleware(
			tracing.Client(),
			func(next middleware.Handler) middleware.Handler {
				return func(ctx context.Context, req interface{}) (interface{}, error) {
					if token, ok := util.GetContextValue[string](ctx, constant.CtxToken); ok && token != "" {
						if r, ok := req.(khttp.Request); ok {
							// 注入 Token 到 HTTP Header
							r.Header.Set(constant.HeaderAuthentication, "Bearer "+token)
						}
					}
					return next(ctx, req)
				}
			},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("create http client for %s failed: %w", service, err)
	}

	actual, loaded := c.httpConns.LoadOrStore(service, client)
	if loaded {
		// Kratos HTTP Client 不直接暴露 Close，旧的会被 GC 回收，
		// 但 Discovery 里的 Watcher 是基于单例 reg 的，不会泄露
		return actual.(*khttp.Client), nil
	}
	return client, nil
}

// tokenClientMiddleware gRPC 客户端 Token 注入中间件
func (c *EtcdClient) tokenClientMiddleware() middleware.Middleware {
	return func(next middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if token, ok := util.GetContextValue[string](ctx, constant.CtxToken); ok && token != "" {
				// gRPC Metadata Key 必须为小写
				ctx = metadata.AppendToOutgoingContext(ctx,
					strings.ToLower(constant.HeaderAuthentication),
					"Bearer "+token,
				)
			}
			return next(ctx, req)
		}
	}
}

// CleanUp 资源清理
func (c *EtcdClient) CleanUp() {
	// 关闭所有 gRPC 连接
	c.grpcConns.Range(func(key, value any) bool {
		if conn, ok := value.(*grpc.ClientConn); ok {
			_ = conn.Close()
		}
		return true
	})

	// 关闭 Etcd 客户端
	if c.cli != nil {
		if err := c.cli.Close(); err != nil {
			c.log.Errorf("failed to close etcd client: %v", err)
		}
	}
	// 重置 Map
	c.grpcConns = sync.Map{}
	c.httpConns = sync.Map{}
	c.log.Info("etcd client and all service connections cleaned up")
}

// GetEtcdServiceClient 泛型客户端获取方法
func GetEtcdServiceClient[T any](etcd *EtcdClient, service string, newClient func(grpc.ClientConnInterface) T) (T, error) {
	conn, err := etcd.GetGrpcConn(service)
	if err != nil {
		var zero T
		return zero, err
	}
	return newClient(conn), nil
}
