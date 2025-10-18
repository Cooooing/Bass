package client

import (
	"content/internal/conf"
	"context"
	"fmt"
	"sync"
	"time"

	etcdregistry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type EtcdCient struct {
	conf  *conf.Bootstrap
	log   *log.Helper
	conns sync.Map // map[string]*grpc.ClientConn
}

func NewEtcdClient(conf *conf.Bootstrap, logger *log.Helper) (*EtcdCient, func(), error) {
	etcdServer := &EtcdCient{
		conf: conf,
		log:  logger,
	}
	return etcdServer, etcdServer.CleanUp, nil
}

func (s *EtcdCient) NewConn() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   s.conf.Registry.Etcd.Endpoints,
		Username:    s.conf.Registry.Etcd.Username,
		Password:    s.conf.Registry.Etcd.Password,
		DialTimeout: s.conf.Registry.Etcd.Timeout.AsDuration(),
	})
}

func (s *EtcdCient) Registrar() registry.Registrar {
	cli, err := s.NewConn()
	if err != nil {
		log.Errorf("new etcd client error: %v", err)
	}
	return etcdregistry.New(cli)
}

func (s *EtcdCient) Discoverer(name string) *grpc.ClientConn {
	if v, ok := s.conns.Load(name); ok {
		if conn, ok := v.(*grpc.ClientConn); ok {
			return conn
		}
	}

	cli, err := s.NewConn()
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

	s.conns.Store(name, conn)
	return conn
}

func (s *EtcdCient) CleanUp() {
	s.conns.Range(func(key, value interface{}) bool {
		if conn, ok := value.(*grpc.ClientConn); ok {
			err := conn.Close()
			if err != nil {
				s.log.Errorf("close etcd conn error: %v", err)
			}
		}
		return true
	})
}
