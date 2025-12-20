package qiniu

import (
	"common/pkg/constant"
	"context"
	"fmt"
	"user/internal/biz/model"
	"user/internal/data/base"
	"user/internal/data/ent/gen"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Qiniu struct {
	*base.BaseRepo
}

func NewQiniu(baseRepo *base.BaseRepo) *Qiniu {
	return &Qiniu{
		BaseRepo: baseRepo,
	}
}

func (q *Qiniu) Name() string {
	return constant.Qiniu.String()
}

func (q *Qiniu) Save(ctx context.Context, tx *gen.Client, o *model.ObjectStorage) (*model.ObjectStorage, error) {
	save, err := tx.ObjectStorage.Create().
		SetProvider(q.Name()).
		SetBucket(q.Conf.Oss.Qiniu.Bucket).
		SetKey(o.Key).
		SetMimeType(o.MimeType).
		SetSize(o.Size).
		SetHash(o.Hash).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ObjectStorage{ObjectStorage: save}, nil
}

func (q *Qiniu) UploadToken(key string) string {
	mac := auth.New(q.Conf.Oss.Qiniu.AccessKey, q.Conf.Oss.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope:            fmt.Sprintf("%s:%s", q.Conf.Oss.Qiniu.Bucket, key),
		CallbackURL:      q.Conf.Oss.Qiniu.CallbackUrl,
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","size":"$(fsize)","bucket":"$(bucket)","name":"$(fname)","mime_type":"${mimeType}"}`,
		CallbackBodyType: "application/json",
		ReturnBody:       `{"key":"$(key)","hash":"$(etag)","size":"$(fsize)","bucket":"$(bucket)","name":"$(fname)","mime_type":"${mimeType}"}`,
		Expires:          uint64(q.Conf.Oss.Qiniu.Timeout.Seconds),
		InsertOnly:       1,
		FsizeMin:         1024 * 1024 * q.Conf.Oss.Qiniu.SizeMin,
		FsizeLimit:       1024 * 1024 * q.Conf.Oss.Qiniu.SizeMax,
		FileType:         0,
	}
	return putPolicy.UploadToken(mac)
}
