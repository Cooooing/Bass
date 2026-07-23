package repo

import "context"

type EmbeddingClient interface {
	Embed(context.Context, []string) ([][]float32, error)
}
