// Package consul synchronizes local service configuration to Consul KV.
package consul

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"

	consulapi "github.com/hashicorp/consul/api"

	"bass/ops-tools/internal/config"
)

// Sync loads every app module config.yaml, expands its environment references, and writes it to Consul.
func Sync(ctx context.Context, address, datacenter, token, prefix, appRoot string) ([]string, error) {
	configs, err := loadConfigs(appRoot, prefix)
	if err != nil {
		return nil, err
	}
	configuration := consulapi.DefaultConfig()
	configuration.Address = address
	configuration.Datacenter = datacenter
	configuration.Token = token
	client, err := consulapi.NewClient(configuration)
	if err != nil {
		return nil, fmt.Errorf("create Consul client: %w", err)
	}

	keys := sortedKeys(configs)
	for _, key := range keys {
		if _, err := client.KV().Put(&consulapi.KVPair{Key: key, Value: configs[key]}, (&consulapi.WriteOptions{}).WithContext(ctx)); err != nil {
			return nil, fmt.Errorf("upload %s: %w", key, err)
		}
	}
	return keys, nil
}

func loadConfigs(appRoot, prefix string) (map[string][]byte, error) {
	paths, err := filepath.Glob(filepath.Join(appRoot, "*", "configs", "config.yaml"))
	if err != nil {
		return nil, err
	}
	if len(paths) == 0 {
		return nil, fmt.Errorf("no config.yaml files found below %s", appRoot)
	}
	configs := make(map[string][]byte, len(paths))
	for _, path := range paths {
		content, err := config.LoadResolvedYAML(path)
		if err != nil {
			return nil, fmt.Errorf("resolve %s: %w", path, err)
		}
		service := filepath.Base(filepath.Dir(filepath.Dir(path)))
		key := filepath.ToSlash(filepath.Join(prefix, service, "config.yaml"))
		configs[key] = content
	}
	return configs, nil
}

func sortedKeys(configs map[string][]byte) []string {
	keys := make([]string, 0, len(configs))
	for key := range configs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
