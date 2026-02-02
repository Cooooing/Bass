package qiniu

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"fmt"
	"infra/internal/biz/model"
	"infra/internal/data/base"
	"infra/internal/data/ent/gen"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Qiniu struct {
	*base.BaseData
}

func NewQiniu(BaseData *base.BaseData) *Qiniu {
	return &Qiniu{
		BaseData: BaseData,
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

func (q *Qiniu) UploadToken(ctx context.Context, key string) (string, error) {
	infra, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return "", cv1.ErrorUnauthorized("infra not login")
	}
	mac := auth.New(q.Conf.Oss.Qiniu.AccessKey, q.Conf.Oss.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope:            fmt.Sprintf("%s:%s", q.Conf.Oss.Qiniu.Bucket, key),
		CallbackURL:      q.Conf.Oss.Qiniu.CallbackUrl,
		CallbackBody:     fmt.Sprintf(`{"key":"$(key)","hash":"$(etag)","size":"$(fsize)","bucket":"$(bucket)","name":"$(fname)","mime_type":"${mimeType}","upload_by":%d,"upload_by_name":"%s"}`, infra.ID, infra.Name),
		CallbackBodyType: "application/json",
		ReturnBody:       fmt.Sprintf(`{"key":"$(key)","hash":"$(etag)","size":"$(fsize)","bucket":"$(bucket)","name":"$(fname)","mime_type":"${mimeType}","upload_by":%d,"upload_by_name":"%s"}`, infra.ID, infra.Name),
		Expires:          uint64(q.Conf.Oss.Qiniu.Timeout.Seconds),
		InsertOnly:       1,
		FsizeMin:         1024 * 1024 * q.Conf.Oss.Qiniu.SizeMin,
		FsizeLimit:       1024 * 1024 * q.Conf.Oss.Qiniu.SizeMax,
		FileType:         0,
	}
	return putPolicy.UploadToken(mac), nil
}

func (q *Qiniu) Status(ctx context.Context, key string, enable bool) error {
	mac := auth.New(q.Conf.Oss.Qiniu.AccessKey, q.Conf.Oss.Qiniu.SecretKey)
	bucketManager := storage.NewBucketManager(mac, nil)
	err := bucketManager.UpdateObjectStatus(q.Conf.Oss.Qiniu.Bucket, key, enable)
	return err
}
