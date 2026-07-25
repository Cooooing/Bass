package qiniu

import (
	"bytes"
	"context"
	"fmt"
	"platform/internal/biz/repo"
	"platform/internal/config"
	"platform/internal/enum"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

var _ repo.ObjectStorageClient = (*Qiniu)(nil)

type Qiniu struct {
	conf *config.Bootstrap
}

func NewQiniu(
	conf *config.Bootstrap,
) *Qiniu {
	return &Qiniu{
		conf: conf,
	}
}

func (q *Qiniu) Name() string {
	return string(enum.ObjectStorageProviderQiniu)
}

func (q *Qiniu) CreateBucket(ctx context.Context, bucket string) error {
	if bucket == "" {
		bucket = q.conf.Platform.Oss.Qiniu.Bucket
	}
	mac := auth.New(q.conf.Platform.Oss.Qiniu.AccessKey, q.conf.Platform.Oss.Qiniu.SecretKey)
	bucketManager := storage.NewBucketManager(mac, nil)
	if err := bucketManager.CreateBucket(bucket, storage.RIDHuadong); err != nil {
		return fmt.Errorf("create qiniu bucket: %w", err)
	}
	return nil
}

func (q *Qiniu) DeleteBucket(ctx context.Context, bucket string) error {
	if bucket == "" {
		bucket = q.conf.Platform.Oss.Qiniu.Bucket
	}
	mac := auth.New(q.conf.Platform.Oss.Qiniu.AccessKey, q.conf.Platform.Oss.Qiniu.SecretKey)
	bucketManager := storage.NewBucketManager(mac, nil)
	if err := bucketManager.DropBucket(bucket); err != nil {
		return fmt.Errorf("delete qiniu bucket: %w", err)
	}
	return nil
}

func (q *Qiniu) Upload(ctx context.Context, req *repo.ObjectStorageUploadReq) (*repo.ObjectStorageUploadResp, error) {
	if req == nil {
		req = &repo.ObjectStorageUploadReq{}
	}
	mac := auth.New(q.conf.Platform.Oss.Qiniu.AccessKey, q.conf.Platform.Oss.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", q.conf.Platform.Oss.Qiniu.Bucket, req.Key),
	}
	if q.conf.Platform.Oss.Qiniu.Timeout != nil {
		putPolicy.Expires = uint64(q.conf.Platform.Oss.Qiniu.Timeout.Seconds)
	}
	putRet := &storage.PutRet{}
	uploader := storage.NewFormUploader(nil)
	if err := uploader.Put(ctx, putRet, putPolicy.UploadToken(mac), req.Key, bytes.NewReader(req.Content), int64(len(req.Content)), &storage.PutExtra{
		MimeType: req.MimeType,
	}); err != nil {
		return nil, err
	}
	return &repo.ObjectStorageUploadResp{
		Provider: enum.ObjectStorageProviderQiniu,
		Bucket:   q.conf.Platform.Oss.Qiniu.Bucket,
		Key:      putRet.Key,
		MimeType: req.MimeType,
		Size:     int64(len(req.Content)),
		Hash:     putRet.Hash,
	}, nil
}

func (q *Qiniu) StreamUpload(ctx context.Context, req *repo.ObjectStorageStreamUploadReq) (*repo.ObjectStorageStreamUploadResp, error) {
	if req == nil {
		req = &repo.ObjectStorageStreamUploadReq{}
	}
	mac := auth.New(q.conf.Platform.Oss.Qiniu.AccessKey, q.conf.Platform.Oss.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", q.conf.Platform.Oss.Qiniu.Bucket, req.Key),
	}
	if q.conf.Platform.Oss.Qiniu.Timeout != nil {
		putPolicy.Expires = uint64(q.conf.Platform.Oss.Qiniu.Timeout.Seconds)
	}
	putRet := &storage.PutRet{}
	uploader := storage.NewResumeUploader(nil)
	if err := uploader.PutWithoutSize(ctx, putRet, putPolicy.UploadToken(mac), req.Key, req.Body, nil); err != nil {
		return nil, err
	}
	return &repo.ObjectStorageStreamUploadResp{
		Provider: enum.ObjectStorageProviderQiniu,
		Bucket:   q.conf.Platform.Oss.Qiniu.Bucket,
		Key:      putRet.Key,
		MimeType: req.MimeType,
		Size:     req.Size,
		Hash:     putRet.Hash,
	}, nil
}

func (q *Qiniu) Download(ctx context.Context, key string) (*repo.ObjectStorageDownloadResp, error) {
	return nil, fmt.Errorf("qiniu download requires download domain config")
}

func (q *Qiniu) StreamDownload(ctx context.Context, key string) (*repo.ObjectStorageStreamDownloadResp, error) {
	return nil, fmt.Errorf("qiniu stream download requires download domain config")
}

func (q *Qiniu) UploadToken(ctx context.Context, req *repo.ObjectStorageUploadTokenReq) (string, error) {
	mac := auth.New(q.conf.Platform.Oss.Qiniu.AccessKey, q.conf.Platform.Oss.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope:            fmt.Sprintf("%s:%s", q.conf.Platform.Oss.Qiniu.Bucket, req.Key),
		CallbackURL:      q.conf.Platform.Oss.Qiniu.CallbackUrl,
		CallbackBody:     fmt.Sprintf(`{"key":"$(key)","hash":"$(etag)","size":"$(fsize)","bucket":"$(bucket)","name":"$(fname)","mime_type":"${mimeType}","upload_by":%d}`, req.UploaderID),
		CallbackBodyType: "application/json",
		ReturnBody:       fmt.Sprintf(`{"key":"$(key)","hash":"$(etag)","size":"$(fsize)","bucket":"$(bucket)","name":"$(fname)","mime_type":"${mimeType}","upload_by":%d}`, req.UploaderID),
		Expires:          uint64(q.conf.Platform.Oss.Qiniu.Timeout.Seconds),
		InsertOnly:       1,
		FsizeMin:         1024 * 1024 * q.conf.Platform.Oss.Qiniu.SizeMin,
		FsizeLimit:       1024 * 1024 * q.conf.Platform.Oss.Qiniu.SizeMax,
		FileType:         0,
	}
	return putPolicy.UploadToken(mac), nil
}

func (q *Qiniu) Status(ctx context.Context, req *repo.ObjectStorageStatusReq) error {
	mac := auth.New(q.conf.Platform.Oss.Qiniu.AccessKey, q.conf.Platform.Oss.Qiniu.SecretKey)
	bucketManager := storage.NewBucketManager(mac, nil)
	err := bucketManager.UpdateObjectStatus(q.conf.Platform.Oss.Qiniu.Bucket, req.Key, req.Enable)
	if err != nil {
		return err
	}
	return nil
}
