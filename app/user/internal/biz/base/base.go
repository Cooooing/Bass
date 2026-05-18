package base

import "context"

// TxRunner is a function that executes fn within a transaction context.
type TxRunner func(ctx context.Context, fn func(ctx context.Context) error) error
