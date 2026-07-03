package client

import (
	"common/pkg/constant"
	"common/proto/gen/common"
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"

	consulregistry "github.com/go-kratos/kratos/contrib/registry/consul/v3"
	metadata2 "github.com/go-kratos/kratos/v3/middleware/metadata"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/registry"
	kgrpc "github.com/go-kratos/kratos/v3/transport/grpc"
	khttp "github.com/go-kratos/kratos/v3/transport/http"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/durationpb"
)

type ConsulClient struct {
	conf     *common.Consul
	logger   *slog.Logger
	observer *Observer

	Client *consulapi.Client
	reg    *consulregistry.Registry

	grpcConns sync.Map
	httpConns sync.Map

	mu       sync.Mutex
	services []string
}

func NewConsulClient(logger *slog.Logger, conf *common.Consul, observer *Observer) (*ConsulClient, func(), error) {
	if observer == nil {
		observer = NewObservability(logger, nil)
	}
	if conf.DialTimeout == nil {
		conf.DialTimeout = durationpb.New(5 * time.Second)
	}
	if conf.Datacenter == "" {
		conf.Datacenter = "dc1"
	}

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
			DialContext: (&net.Dialer{Timeout: conf.DialTimeout.AsDuration()}).DialContext,
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

	if _, err = apiClient.Catalog().Datacenters(); err != nil {
		return nil, nil, fmt.Errorf("consul unreachable: %w", err)
	}

	c := &ConsulClient{
		conf:     conf,
		logger:   logger,
		observer: observer,
		Client:   apiClient,
		reg:      reg,
	}

	c.logger.Info("consul connected", constant.LogFieldKind, constant.LogKindConsul, constant.LogFieldAddress, conf.Address, constant.LogFieldDatacenter, cfg.Datacenter)
	return c, c.Close, nil
}

func (c *ConsulClient) Registrar() registry.Registrar { return c.reg }

func (c *ConsulClient) Discovery() registry.Discovery { return c.reg }

func (c *ConsulClient) RegisterService(svc *consulapi.AgentServiceRegistration) error {
	if err := c.Client.Agent().ServiceRegister(svc); err != nil {
		return fmt.Errorf("register %s: %w", svc.ID, err)
	}
	c.mu.Lock()
	c.services = append(c.services, svc.ID)
	c.mu.Unlock()
	c.logger.Info("consul service registered", constant.LogFieldKind, constant.LogKindConsul, "service_id", svc.ID)
	return nil
}

func (c *ConsulClient) GetGrpcConn(service string) (*grpc.ClientConn, error) {
	if v, ok := c.grpcConns.Load(service); ok {
		return v.(*grpc.ClientConn), nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.conf.DialTimeout.AsDuration())
	defer cancel()

	conn, err := kgrpc.NewClient(ctx,
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", service)),
		kgrpc.WithDiscovery(c.reg),
		kgrpc.WithTimeout(c.conf.DialTimeout.AsDuration()),
		kgrpc.WithOptions(
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                30 * time.Second,
				Timeout:             10 * time.Second,
				PermitWithoutStream: false,
			}),
		),
		kgrpc.WithMiddleware(
			recovery.Recovery(),
			c.observer.ClientMiddleware(service),
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

	c.logger.Info("grpc connection created", constant.LogFieldKind, constant.LogKindClient, constant.LogFieldTarget, service)
	return conn, nil
}

func (c *ConsulClient) GetHTTPClient(service string) (*khttp.Client, error) {
	if v, ok := c.httpConns.Load(service); ok {
		return v.(*khttp.Client), nil
	}

	cl, err := khttp.NewClient(context.Background(),
		khttp.WithEndpoint(fmt.Sprintf("discovery:///%s", service)),
		khttp.WithDiscovery(c.reg),
		khttp.WithMiddleware(
			c.observer.ClientMiddleware(service),
			metadata2.Client(),
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

	c.logger.Info("http client created", constant.LogFieldKind, constant.LogKindClient, constant.LogFieldTarget, service)
	return cl, nil
}

func (c *ConsulClient) Close() {
	c.mu.Lock()
	ids := c.services
	c.services = nil
	c.mu.Unlock()

	for _, id := range ids {
		if err := c.Client.Agent().ServiceDeregister(id); err != nil {
			c.logger.Error("consul service deregister failed", constant.LogFieldKind, constant.LogKindConsul, "service_id", id, constant.LogFieldErr, err)
		} else {
			c.logger.Info("consul service deregistered", constant.LogFieldKind, constant.LogKindConsul, "service_id", id)
		}
	}

	c.grpcConns.Range(func(_, v any) bool {
		_ = v.(*grpc.ClientConn).Close()
		return true
	})

	c.httpConns.Range(func(_, v any) bool {
		_ = v.(*khttp.Client).Close()
		return true
	})

	c.logger.Info("consul client closed", constant.LogFieldKind, constant.LogKindConsul)
}
