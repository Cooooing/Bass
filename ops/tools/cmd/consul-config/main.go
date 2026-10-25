package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"bass/ops-tools/internal/config"
	"bass/ops-tools/internal/consul"
)

const (
	envFilePath = "ops/tools/cmd/consul-config/.env"
	appRoot     = "app"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "consul-config:", err)
		os.Exit(1)
	}
}

func run() error {
	envPath, err := config.FindPath(envFilePath)
	if err != nil {
		return fmt.Errorf("find .env: %w", err)
	}
	if err := config.LoadEnvFile(envPath); err != nil {
		return err
	}
	address := requiredEnv("CONSUL_ADDRESS")
	datacenter := requiredEnv("CONSUL_DATACENTER")
	prefix := strings.Trim(requiredEnv("CONSUL_PREFIX"), "/")
	if address == "" || datacenter == "" || prefix == "" {
		return fmt.Errorf("CONSUL_ADDRESS, CONSUL_DATACENTER and CONSUL_PREFIX must be set")
	}
	root, err := config.FindPath(appRoot)
	if err != nil {
		return fmt.Errorf("find app configuration: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	keys, err := consul.Sync(ctx, address, datacenter, os.Getenv("CONSUL_TOKEN"), prefix, root)
	if err != nil {
		return err
	}
	for _, key := range keys {
		fmt.Println(key)
	}
	return nil
}

func requiredEnv(name string) string {
	return strings.TrimSpace(os.Getenv(name))
}
