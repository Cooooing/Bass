package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"bass/ops-tools/internal/config"
	"bass/ops-tools/internal/consul"
)

const envFilePath = "ops/tools/cmd/consul-backup/.env"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "consul-backup:", err)
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
	backupDir := requiredEnv("BACKUP_DIR")
	if address == "" || datacenter == "" || backupDir == "" {
		return fmt.Errorf("CONSUL_ADDRESS, CONSUL_DATACENTER and BACKUP_DIR must be set")
	}
	resolvedBackupDir, err := resolveOutputDirectory(backupDir)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	path, count, err := consul.BackupToDirectory(ctx, address, datacenter, os.Getenv("CONSUL_TOKEN"), resolvedBackupDir)
	if err != nil {
		return err
	}
	fmt.Printf("Saved %d keys to %s\n", count, path)
	return nil
}

func requiredEnv(name string) string {
	return strings.TrimSpace(os.Getenv(name))
}

func resolveOutputDirectory(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	opsDirectory, err := config.FindPath("ops")
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(opsDirectory), path), nil
}
