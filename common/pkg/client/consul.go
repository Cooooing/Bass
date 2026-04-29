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

// ConsulClient manages Consul-based service registration and discovery.
type ConsulClient struct {
	conf   *common.Consul
	logger *log.Helper

	// Client is the raw Consul API client, exposed for advanced usage.
	Client *consulapi.Client
	reg    *consulregistry.Registry

	// Long-lived connection pools.
	// gRPC reconnection is handled internally via resolver + keepalive.
	grpcConns sync.Map // service -> *grpc.ClientConn
	httpConns sync.Map // service -> *khttp.Client

	// Tracks registered service IDs for automatic deregistration on shutdown.
	mu       sync.Mutex
	services []string
}

// NewConsulClient creates a ConsulClient and returns a cleanup function.
func NewConsulClient(logger log.Logger, conf *common.Consul) (*ConsulClient, func(), error) {
	if conf.DialTimeout == nil {
		conf.DialTimeout = durationpb.New(5 * time.Second)
	}

	// Build Consul API config
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
	)

	// Connectivity pre-check
	if _, err = apiClient.Catalog().Datacenters(); err != nil {
		return nil, nil, fmt.Errorf("consul unreachable: %w", err)
	}

	c := &ConsulClient{
		conf:   conf,
		logger: log.NewHelper(logger),
		Client: apiClient,
		reg:    reg,
	}

	c.logger.Infof("consul connected: %s (dc=%s)", conf.Address, cfg.Datacenter)
	return c, c.Close, nil
}

// Registrar returns the Kratos Registrar backed by Consul.
func (c *ConsulClient) Registrar() registry.Registrar { return c.reg }

// Discovery returns the Kratos Discovery backed by Consul.
func (c *ConsulClient) Discovery() registry.Discovery { return c.reg }

// RegisterService registers a service with Consul.
// The service will be automatically deregistered when Close is called.
func (c *ConsulClient) RegisterService(svc *consulapi.AgentServiceRegistration) error {
	if err := c.Client.Agent().ServiceRegister(svc); err != nil {
		return fmt.Errorf("register %s: %w", svc.ID, err)
	}
	c.mu.Lock()
	c.services = append(c.services, svc.ID)
	c.mu.Unlock()
	c.logger.Infof("registered: %s", svc.ID)
	return nil
}

// GetGrpcConn returns a cached gRPC connection for the named service.
// Connections are long-lived: gRPC handles reconnection, the Consul resolver
// handles endpoint updates, and keepalive probes detect dead connections.
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
				PermitWithoutStream: true,
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

	c.logger.Infof("grpc conn: %s", service)
	return conn, nil
}

// GetHTTPClient returns a cached HTTP client for the named service.
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

	c.logger.Infof("http client: %s", service)
	return cl, nil
}

// Close deregisters all services and closes all connections.
func (c *ConsulClient) Close() {
	// Deregister services
	c.mu.Lock()
	ids := c.services
	c.services = nil
	c.mu.Unlock()

	for _, id := range ids {
		if err := c.Client.Agent().ServiceDeregister(id); err != nil {
			c.logger.Warnf("deregister %s: %v", id, err)
		} else {
			c.logger.Infof("deregistered: %s", id)
		}
	}

	// Close gRPC connections
	c.grpcConns.Range(func(_, v any) bool {
		_ = v.(*grpc.ClientConn).Close()
		return true
	})

	// Close HTTP clients
	c.httpConns.Range(func(_, v any) bool {
		_ = v.(*khttp.Client).Close()
		return true
	})

	c.logger.Info("consul client closed")
}
