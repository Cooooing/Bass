// Package config loads command-local environment files and resolves service configuration.
package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	kratosconfig "github.com/go-kratos/kratos/v3/config"
	"github.com/go-kratos/kratos/v3/config/env"
	"github.com/go-kratos/kratos/v3/config/file"
	"go.yaml.in/yaml/v3"
)

// LoadEnvFile loads values that are not already supplied by the process environment.
func LoadEnvFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, found := strings.Cut(strings.TrimPrefix(line, "export "), "=")
		if !found || strings.TrimSpace(key) == "" {
			return fmt.Errorf("invalid .env line %d", lineNumber)
		}
		key = strings.TrimSpace(key)
		if _, setByProcess := os.LookupEnv(key); !setByProcess {
			if err := os.Setenv(key, strings.Trim(strings.TrimSpace(value), "\"'")); err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}

// LoadResolvedYAML uses the same Kratos environment and file sources as services.
// Kratos resolves placeholders while the YAML node tree preserves source ordering.
func LoadResolvedYAML(path string) ([]byte, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	loader := kratosconfig.New(kratosconfig.WithSource(env.NewSource(""), file.NewSource(path)))
	defer loader.Close()
	if err := loader.Load(); err != nil {
		return nil, err
	}
	var document yaml.Node
	if err := yaml.Unmarshal(contents, &document); err != nil {
		return nil, err
	}
	if err := resolveYAMLNode(&document, nil, loader); err != nil {
		return nil, err
	}
	resolved, err := yaml.Marshal(&document)
	if err != nil {
		return nil, err
	}
	return resolved, nil
}

func resolveYAMLNode(node *yaml.Node, path []string, loader kratosconfig.Config) error {
	switch node.Kind {
	case yaml.DocumentNode:
		for _, child := range node.Content {
			if err := resolveYAMLNode(child, path, loader); err != nil {
				return err
			}
		}
	case yaml.MappingNode:
		for index := 0; index < len(node.Content); index += 2 {
			key := node.Content[index].Value
			if err := resolveYAMLNode(node.Content[index+1], append(path, key), loader); err != nil {
				return err
			}
		}
	case yaml.SequenceNode:
		for index, child := range node.Content {
			if err := resolveYAMLNode(child, append(path, fmt.Sprintf("%d", index)), loader); err != nil {
				return err
			}
		}
	case yaml.ScalarNode:
		if !strings.Contains(node.Value, "${") {
			return nil
		}
		value, err := loader.Value(strings.Join(path, ".")).String()
		if err != nil {
			return fmt.Errorf("resolve %s: %w", strings.Join(path, "."), err)
		}
		node.Value = value
	}
	return nil
}

// FindPath resolves a repository-relative path by walking upward from the current directory.
func FindPath(path string) (string, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		candidate := filepath.Join(workingDirectory, path)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		parent := filepath.Dir(workingDirectory)
		if parent == workingDirectory {
			return "", fmt.Errorf("path does not exist: %s", path)
		}
		workingDirectory = parent
	}
}
