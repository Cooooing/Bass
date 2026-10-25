package consul

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

// Backup is a point-in-time Consul KV snapshot.
type Backup struct {
	CreatedAt  time.Time `json:"created_at"`
	Datacenter string    `json:"datacenter"`
	Entries    []KVEntry `json:"entries"`
}

// KVEntry contains the complete value for one KV key.
type KVEntry struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
}

// BackupAll reads every key under the target Consul KV root.
func BackupAll(ctx context.Context, address, datacenter, token string) (*Backup, error) {
	configuration := consulapi.DefaultConfig()
	configuration.Address = address
	configuration.Datacenter = datacenter
	configuration.Token = token
	client, err := consulapi.NewClient(configuration)
	if err != nil {
		return nil, fmt.Errorf("create Consul client: %w", err)
	}
	pairs, _, err := client.KV().List("", (&consulapi.QueryOptions{}).WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("list Consul KV: %w", err)
	}
	entries := make([]KVEntry, 0, len(pairs))
	for _, pair := range pairs {
		entries = append(entries, KVEntry{
			Key:   pair.Key,
			Value: append([]byte(nil), pair.Value...),
		})
	}
	sort.Slice(entries, func(left, right int) bool {
		return entries[left].Key < entries[right].Key
	})
	return &Backup{
		CreatedAt:  time.Now().UTC(),
		Datacenter: datacenter,
		Entries:    entries,
	}, nil
}

// BackupToDirectory reads all target KV entries and saves them under one timestamped directory.
func BackupToDirectory(ctx context.Context, address, datacenter, token, directory string) (string, int, error) {
	snapshot, err := BackupAll(ctx, address, datacenter, token)
	if err != nil {
		return "", 0, err
	}
	path, err := writeSnapshot(directory, snapshot)
	if err != nil {
		return "", 0, err
	}
	return path, len(snapshot.Entries), nil
}

func writeSnapshot(directory string, snapshot *Backup) (string, error) {
	if err := os.MkdirAll(directory, 0o700); err != nil {
		return "", fmt.Errorf("create backup directory: %w", err)
	}
	temporaryDirectory, err := os.MkdirTemp(directory, ".consul-*")
	if err != nil {
		return "", fmt.Errorf("create temporary backup directory: %w", err)
	}
	defer os.RemoveAll(temporaryDirectory)
	for _, entry := range snapshot.Entries {
		path, err := backupPath(temporaryDirectory, entry.Key)
		if err != nil {
			return "", err
		}
		if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
			return "", fmt.Errorf("create directory for %s: %w", entry.Key, err)
		}
		if err := os.WriteFile(path, entry.Value, 0o600); err != nil {
			return "", fmt.Errorf("write %s: %w", entry.Key, err)
		}
	}
	path := filepath.Join(directory, "consul-"+snapshot.CreatedAt.Format("20060102T150405Z"))
	if err := os.Rename(temporaryDirectory, path); err != nil {
		return "", fmt.Errorf("save backup directory: %w", err)
	}
	return path, nil
}

func backupPath(root, key string) (string, error) {
	parts := strings.Split(key, "/")
	for _, part := range parts {
		if part == "" || part == "." || part == ".." || strings.Contains(part, `\`) {
			return "", fmt.Errorf("unsafe Consul key: %q", key)
		}
	}
	return filepath.Join(append([]string{root}, parts...)...), nil
}
