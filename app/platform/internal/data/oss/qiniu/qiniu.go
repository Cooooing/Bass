package qiniu

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"context"
	"fmt"
	"platform/internal/biz/model"
	"platform/internal/config"
	"platform/internal/data/gen"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"log/slog"
)

type Qiniu struct {
	conf         *config.Bootstrap
	log          *slog.Logger
	db           *gen.Client
	consulClient *commonClient.ConsulClient
	redisClient  *commonClient.RedisClient
}

func NewQiniu(
	conf *config.Bootstrap,
	logger *slog.Logger,
	db *gen.Client,
	consulClient *commonClient.ConsulClient,
	redisClient *commonClient.RedisClient,
) *Qiniu {
	return &Qiniu{
		conf:         conf,
		log:          logger,
		db:           db,
		consulClient: consulClient,
		redisClient:  redisClient,
	}
}

func (q *Qiniu) Name() string {
	return constant.Qiniu.String()
}

func (q *Qiniu) Save(ctx context.Context, tx *gen.Client, o *model.ObjectStorage) (*model.ObjectStorage, error) {
	save, err := tx.ObjectStorage.Create().
		SetProvider(q.Name()).
		SetBucket(q.conf.Platform.Oss.Qiniu.Bucket).
		SetKey(o.Key).
		SetMimeType(o.MimeType).
		SetSize(o.Size).
		SetHash(o.Hash).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ObjectStorage{
		ID:           save.ID,
		Provider:     save.Provider,
		Bucket:       save.Bucket,
		Key:          save.Key,
		MimeType:     save.MimeType,
		Size:         save.Size,
		Hash:         save.Hash,
		UploadBy:     save.UploadBy,
		UploadByName: save.UploadByName,
		CreatedAt:    save.CreatedAt,
		UpdatedAt:    save.UpdatedAt,
	}, nil
}

func (q *Qiniu) UploadToken(ctx context.Context, key string, uploaderID int64, uploaderName string) (string, error) {
	mac := auth.New(q.conf.Platform.Oss.Qiniu.AccessKey, q.conf.Platform.Oss.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope:            fmt.Sprintf("%s:%s", q.conf.Platform.Oss.Qiniu.Bucket, key),
		CallbackURL:      q.conf.Platform.Oss.Qiniu.CallbackUrl,
		CallbackBody:     fmt.Sprintf(`{"key":"$(key)","hash":"$(etag)","size":"$(fsize)","bucket":"$(bucket)","name":"$(fname)","mime_type":"${mimeType}","upload_by":%d,"upload_by_name":"%s"}`, uploaderID, uploaderName),
		CallbackBodyType: "application/json",
		ReturnBody:       fmt.Sprintf(`{"key":"$(key)","hash":"$(etag)","size":"$(fsize)","bucket":"$(bucket)","name":"$(fname)","mime_type":"${mimeType}","upload_by":%d,"upload_by_name":"%s"}`, uploaderID, uploaderName),
		Expires:          uint64(q.conf.Platform.Oss.Qiniu.Timeout.Seconds),
		InsertOnly:       1,
		FsizeMin:         1024 * 1024 * q.conf.Platform.Oss.Qiniu.SizeMin,
		FsizeLimit:       1024 * 1024 * q.conf.Platform.Oss.Qiniu.SizeMax,
		FileType:         0,
	}
	return putPolicy.UploadToken(mac), nil
}

func (q *Qiniu) Status(ctx context.Context, key string, enable bool) error {
	mac := auth.New(q.conf.Platform.Oss.Qiniu.AccessKey, q.conf.Platform.Oss.Qiniu.SecretKey)
	bucketManager := storage.NewBucketManager(mac, nil)
	err := bucketManager.UpdateObjectStatus(q.conf.Platform.Oss.Qiniu.Bucket, key, enable)
	return err
}
