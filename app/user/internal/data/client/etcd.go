package client

import (
	"context"
	"fmt"
	"sync"
	"time"
	"user/internal/conf"

	etcdregistry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type EtcdClient struct {
	conf  *conf.Bootstrap
	log   *log.Helper
	conns sync.Map // map[string]*grpc.ClientConn
}

func NewEtcdClient(conf *conf.Bootstrap, log *log.Helper) (*EtcdClient, func(), error) {
	etcdServer := &EtcdClient{
		conf: conf,
		log:  log,
	}
	conn, err := etcdServer.NewConn()
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = conn.Get(ctx, "ping_key_not_exist")
	if err != nil {
		log.Errorf(fmt.Errorf("etcd ping failed:%w", err).Error())
	} else {
		log.Infof("etcd: connected to %+v", conf.Registry.Etcd.Endpoints)
	}
	defer func(conn *clientv3.Client) {
		err := conn.Close()
		if err != nil {
			log.Errorf("close etcd conn error: %v", err)
		}
	}(conn)

	return etcdServer, etcdServer.CleanUp, nil
}

func (c *EtcdClient) NewConn() (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.conf.Registry.Etcd.Endpoints,
		Username:    c.conf.Registry.Etcd.Username,
		Password:    c.conf.Registry.Etcd.Password,
		DialTimeout: c.conf.Registry.Etcd.Timeout.AsDuration(),
	})
	if err != nil {
		log.Errorf("new etcd client error: %v", err)
	}
	return client, err
}

func (c *EtcdClient) Registrar() registry.Registrar {
	cli, err := c.NewConn()
	if err != nil {
		log.Errorf("new etcd client error: %v", err)
	}
	return etcdregistry.New(cli)
}

func (c *EtcdClient) Discoverer(name string) *grpc.ClientConn {
	if v, ok := c.conns.Load(name); ok {
		if conn, ok := v.(*grpc.ClientConn); ok {
			return conn
		}
	}

	cli, err := c.NewConn()
	if err != nil {
		log.Errorf("new etcd client error: %v", err)
	}

	dis := etcdregistry.New(cli)
	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", name)),
		kgrpc.WithDiscovery(dis),
		kgrpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	c.conns.Store(name, conn)
	return conn
}

func (c *EtcdClient) CleanUp() {
	c.conns.Range(func(key, value interface{}) bool {
		if conn, ok := value.(*grpc.ClientConn); ok {
			err := conn.Close()
			if err != nil {
				c.log.Errorf("close etcd conn error: %v", err)
			}
		}
		return true
	})
	c.log.Infof("etcd connections closed")
}
