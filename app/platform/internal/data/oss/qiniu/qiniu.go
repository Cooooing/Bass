package qiniu

import (
	"context"
	"fmt"
	"platform/internal/biz/repo"
	"platform/internal/config"
)

// Qiniu is deliberately not implemented for the content-addressed direct
// upload flow. The interface remains so a provider can be added without
// leaking provider choices into business services.
type Qiniu struct{}

func NewQiniu(_ *config.Bootstrap) *Qiniu {
	return &Qiniu{}
}

func (q *Qiniu) Name() string { return "qiniu" }

func (q *Qiniu) Bucket() string { return "" }

func (q *Qiniu) PrepareDirectUpload(context.Context, *repo.DirectAssetUploadReq) (*repo.DirectAssetUploadResp, error) {
	return nil, fmt.Errorf("Qiniu asset direct upload is not implemented")
}

func (q *Qiniu) Stat(context.Context, string) (*repo.AssetObject, error) {
	return nil, fmt.Errorf("Qiniu asset storage is not implemented")
}

func (q *Qiniu) PublicURL(string) string { return "" }

func (q *Qiniu) Get(context.Context, string) (*repo.StoredObject, error) {
	return nil, fmt.Errorf("Qiniu asset storage is not implemented")
}

func (q *Qiniu) Put(context.Context, *repo.PutStoredObjectReq) (*repo.StoredObject, error) {
	return nil, fmt.Errorf("Qiniu asset storage is not implemented")
}
