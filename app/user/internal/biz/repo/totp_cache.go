package repo

import (
	"context"
	"time"
)

type TotpSecretCache interface {
	Save(ctx context.Context, userID int64, secret string, ttl time.Duration) error
	Get(ctx context.Context, userID int64) (string, error)
}
