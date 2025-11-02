package client

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"common/pkg/model"

	etcdregistry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

// ConnPool 是 gRPC 客户端连接池
type ConnPool struct {
	conns []*grpc.ClientConn
	next  uint32
	once  sync.Once // 每个池自己控制只初始化一次
	err   error
}

// Get 返回一个轮询的 gRPC 连接
func (p *ConnPool) Get() *grpc.ClientConn {
	if len(p.conns) == 0 {
		return nil
	}
	// 使用原子操作保证并发安全轮询
	n := atomic.AddUint32(&p.next, 1)
	return p.conns[int(n%uint32(len(p.conns)))]
}

// EtcdClient 封装 Etcd 和 gRPC 连接池
type EtcdClient struct {
	conf  *model.EtcdConf
	log   *log.Helper
	pools sync.Map         // key=serviceName, value=*ConnPool
	cli   *clientv3.Client // 持久化 Etcd 客户端
}

// NewEtcdClient 创建 EtcdClient 实例
func NewEtcdClient(log *log.Helper, conf *model.EtcdConf) (*EtcdClient, func(), error) {
	// 初始化 Etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.Endpoints,
		Username:    conf.Username,
		Password:    conf.Password,
		DialTimeout: conf.Timeout.AsDuration(),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("create etcd client failed: %w", err)
	}

	c := &EtcdClient{
		conf: conf,
		log:  log,
		cli:  cli,
	}

	// 测试 Etcd 是否可用
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if _, err := cli.Get(ctx, "ping_key_not_exist"); err != nil {
		c.log.Errorf("etcd ping failed: %v", err)
	} else {
		c.log.Infof("etcd connected successfully: %+v", conf.Endpoints)
	}

	return c, c.CleanUp, nil
}

// Registrar 返回 Etcd 服务注册器
func (c *EtcdClient) Registrar() registry.Registrar {
	return etcdregistry.New(c.cli)
}

// newGrpcConn 创建 gRPC 连接
func (c *EtcdClient) newGrpcConn(service string) (*grpc.ClientConn, error) {
	dis := etcdregistry.New(c.cli)
	return kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", service)),
		kgrpc.WithDiscovery(dis),
		kgrpc.WithTimeout(c.conf.Timeout.AsDuration()),
	)
}

// getConnFromPool 获取 gRPC 连接池中的连接，如果池不存在则初始化
func (c *EtcdClient) getConnFromPool(service string, poolSize int) (*grpc.ClientConn, error) {
	val, _ := c.pools.LoadOrStore(service, &ConnPool{})
	pool := val.(*ConnPool)

	pool.once.Do(func() {
		for i := 0; i < poolSize; i++ {
			conn, err := c.newGrpcConn(service)
			if err != nil {
				c.log.Errorf("failed to create grpc conn for %s: %v", service, err)
				pool.err = err
				continue
			}
			pool.conns = append(pool.conns, conn)
		}
		if len(pool.conns) == 0 {
			pool.err = fmt.Errorf("no grpc connections available for service %s", service)
		} else {
			c.log.Infof("created connection pool for service %s (size=%d)", service, len(pool.conns))
		}
	})

	if pool.err != nil {
		return nil, pool.err
	}

	return pool.Get(), nil
}

// CleanUp 关闭所有 gRPC 连接和 Etcd 客户端
func (c *EtcdClient) CleanUp() {
	// 遍历关闭 gRPC 连接池
	c.pools.Range(func(key, value any) bool {
		pool := value.(*ConnPool)
		for _, conn := range pool.conns {
			if err := conn.Close(); err != nil {
				c.log.Warnf("close grpc conn failed (%v): %v", key, err)
			}
		}
		return true
	})
	// 清空连接池
	c.pools = sync.Map{}
	// 关闭 Etcd 客户端
	if c.cli != nil {
		_ = c.cli.Close()
	}
	c.log.Info("all grpc connections and etcd client closed")
}

// GetServiceClient 泛型客户端工厂，从池中获取连接并返回客户端
func GetServiceClient[T any](ctx context.Context, etcd *EtcdClient, service string, newClient func(grpc.ClientConnInterface) T) (T, error) {
	const poolSize = 3 // 连接池大小，可改为配置
	conn, err := etcd.getConnFromPool(service, poolSize)
	if err != nil {
		var zero T
		return zero, fmt.Errorf("get client conn failed: %w", err)
	}
	return newClient(conn), nil
}
