package repo

import (
	"context"
	"time"
)

type TotpSecretCache interface {
	Save(ctx context.Context, req *TotpSecretCacheSaveReq) (*TotpSecretCacheSaveResponse, error)
	Get(ctx context.Context, req *TotpSecretCacheGetReq) (*TotpSecretCacheGetResponse, error)
}

type TotpSecretCacheSaveReq struct {
	UserID int64
	Secret string
	TTL    time.Duration
}

type TotpSecretCacheSaveResponse struct{}

type TotpSecretCacheGetReq struct {
	UserID int64
}

type TotpSecretCacheGetResponse struct {
	Secret string
}
