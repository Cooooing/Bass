package qiniu

import (
	cerrors "common/api/gen/common/errors"
	commonClient "common/pkg/client"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"fmt"
	"notify/internal/biz/model"
	"notify/internal/conf"
	"notify/internal/data/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Qiniu struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
}

func NewQiniu(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
) *Qiniu {
	return &Qiniu{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
	}
}

func (q *Qiniu) Name() string {
	return constant.Qiniu.String()
}

func (q *Qiniu) Save(ctx context.Context, tx *gen.Client, o *model.ObjectStorage) (*model.ObjectStorage, error) {
	save, err := tx.ObjectStorage.Create().
		SetProvider(q.Name()).
		SetBucket(q.conf.Server.Oss.Qiniu.Bucket).
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
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return "", cerrors.ErrorUnauthorized("user not login")
	}
	mac := auth.New(q.conf.Server.Oss.Qiniu.AccessKey, q.conf.Server.Oss.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope:            fmt.Sprintf("%s:%s", q.conf.Server.Oss.Qiniu.Bucket, key),
		CallbackURL:      q.conf.Server.Oss.Qiniu.CallbackUrl,
		CallbackBody:     fmt.Sprintf(`{"key":"$(key)","hash":"$(etag)","size":"$(fsize)","bucket":"$(bucket)","name":"$(fname)","mime_type":"${mimeType}","upload_by":%d,"upload_by_name":"%s"}`, user.ID, user.Name),
		CallbackBodyType: "application/json",
		ReturnBody:       fmt.Sprintf(`{"key":"$(key)","hash":"$(etag)","size":"$(fsize)","bucket":"$(bucket)","name":"$(fname)","mime_type":"${mimeType}","upload_by":%d,"upload_by_name":"%s"}`, user.ID, user.Name),
		Expires:          uint64(q.conf.Server.Oss.Qiniu.Timeout.Seconds),
		InsertOnly:       1,
		FsizeMin:         1024 * 1024 * q.conf.Server.Oss.Qiniu.SizeMin,
		FsizeLimit:       1024 * 1024 * q.conf.Server.Oss.Qiniu.SizeMax,
		FileType:         0,
	}
	return putPolicy.UploadToken(mac), nil
}

func (q *Qiniu) Status(ctx context.Context, key string, enable bool) error {
	mac := auth.New(q.conf.Server.Oss.Qiniu.AccessKey, q.conf.Server.Oss.Qiniu.SecretKey)
	bucketManager := storage.NewBucketManager(mac, nil)
	err := bucketManager.UpdateObjectStatus(q.conf.Server.Oss.Qiniu.Bucket, key, enable)
	return err
}
