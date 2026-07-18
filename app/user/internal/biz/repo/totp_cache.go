package repo

import (
	"context"
	"time"
)

type TotpSecretCache interface {
	Save(ctx context.Context, req *TotpSecretCacheSaveReq) error
	Get(ctx context.Context, userID int64) (string, error)
}

type TotpSecretCacheSaveReq struct {
	UserID int64
	Secret string
	TTL    time.Duration
}
