package main

import (
	"log/slog"
	"os"
	"time"

	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	"common/proto/gen/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/durationpb"
)

func newGameTownClient(addr string, consulAddr string, datacenter string, token string, timeout time.Duration) (*rpc.GameTownClient, func(), string, error) {
	if addr != "" {
		conn, err := grpc.NewClient(
			addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, nil, addr, err
		}
		return rpc.NewGameTownClient(conn), func() {
			_ = conn.Close()
		}, addr, nil
	}

	target := "consul://" + consulAddr + "/" + constant.GameTownServiceName.String()
	consul, cleanup, err := commonClient.NewConsulClient(
		slog.Default(),
		&common.Consul{
			Address:     consulAddr,
			Datacenter:  datacenter,
			Token:       token,
			DialTimeout: durationpb.New(timeout),
		},
		nil,
	)
	if err != nil {
		return nil, nil, target, err
	}
	client, err := rpc.ProvideGameTownClient(consul)
	if err != nil {
		cleanup()
		return nil, nil, target, err
	}
	return client, cleanup, target, nil
}

func envDefault(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
