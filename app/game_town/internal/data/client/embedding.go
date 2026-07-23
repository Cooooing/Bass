package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"game_town/internal/biz/repo"
	"game_town/internal/config"
)

var _ repo.EmbeddingClient = (*EmbeddingClient)(nil)

type EmbeddingClient struct {
	enabled    bool
	baseURL    string
	model      string
	dimensions int
	timeout    time.Duration
	http       *http.Client
}

func NewEmbeddingClient(conf *config.Bootstrap) repo.EmbeddingClient {
	memory := conf.GetGameTown().GetMemory()
	timeout := 30 * time.Second
	if memory.GetTimeout() != nil && memory.GetTimeout().AsDuration() > 0 {
		timeout = memory.GetTimeout().AsDuration()
	}
	return &EmbeddingClient{
		enabled:    memory.GetEmbeddingEnabled(),
		baseURL:    strings.TrimRight(memory.GetBaseUrl(), "/"),
		model:      memory.GetModel(),
		dimensions: int(memory.GetDimensions()),
		timeout:    timeout,
		http:       &http.Client{Timeout: timeout},
	}
}

func (c *EmbeddingClient) Embed(ctx context.Context, input []string) ([][]float32, error) {
	if !c.enabled || len(input) == 0 {
		return nil, fmt.Errorf("embedding is disabled")
	}
	body, err := json.Marshal(map[string]any{"model": c.model, "input": input})
	if err != nil {
		return nil, err
	}
	requestCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(requestCtx, http.MethodPost, c.baseURL+"/api/embed", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("embedding http status %d: %s", resp.StatusCode, strings.TrimSpace(string(responseBody)))
	}
	var result struct {
		Embeddings [][]float32 `json:"embeddings"`
	}
	if err = json.Unmarshal(responseBody, &result); err != nil {
		return nil, err
	}
	if len(result.Embeddings) != len(input) {
		return nil, fmt.Errorf("embedding count mismatch: got %d want %d", len(result.Embeddings), len(input))
	}
	for _, vector := range result.Embeddings {
		if c.dimensions > 0 && len(vector) != c.dimensions {
			return nil, fmt.Errorf("embedding dimension mismatch: got %d want %d", len(vector), c.dimensions)
		}
	}
	return result.Embeddings, nil
}
