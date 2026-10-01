package base

import (
	"context"
)

type Tx func(ctx context.Context, fn func(ctx context.Context) error) error
