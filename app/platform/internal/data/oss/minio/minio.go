package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"platform/internal/biz/repo"
	"platform/internal/config"
	"platform/internal/enum"
	"time"

	miniosdk "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var _ repo.ObjectStorageClient = (*Minio)(nil)

type Minio struct {
	client        *miniosdk.Client
	bucket        string
	uploadTimeout time.Duration
}

func NewMinio(
	conf *config.Bootstrap,
) (*Minio, error) {
	src := conf.GetPlatform().GetOss().GetMinio()
	if src.GetEndpoint() == "" || src.GetAccessKey() == "" || src.GetSecretKey() == "" || src.GetBucket() == "" {
		return nil, fmt.Errorf("minio config requires endpoint, access_key, secret_key and bucket")
	}
	client, err := miniosdk.New(
		src.GetEndpoint(),
		&miniosdk.Options{
			Creds:  credentials.NewStaticV4(src.GetAccessKey(), src.GetSecretKey(), ""),
			Secure: src.GetUseSsl(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("create minio client: %w", err)
	}
	uploadTimeout := 30 * time.Minute
	if src.GetTimeout() != nil && src.GetTimeout().AsDuration() > 0 {
		uploadTimeout = src.GetTimeout().AsDuration()
	}
	m := &Minio{
		client:        client,
		bucket:        src.GetBucket(),
		uploadTimeout: uploadTimeout,
	}
	err = m.CreateBucket(context.Background(), m.bucket)
	if err != nil {
		return nil, fmt.Errorf("create minio bucket: %w", err)
	}
	return m, nil
}

func (m *Minio) Name() string {
	return string(enum.ObjectStorageProviderMinio)
}

func (m *Minio) CreateBucket(
	ctx context.Context,
	bucket string,
) error {
	if bucket == "" {
		bucket = m.bucket
	}
	exists, err := m.client.BucketExists(ctx, bucket)
	if err != nil {
		return fmt.Errorf("check minio bucket: %w", err)
	}
	if exists {
		return nil
	}
	if err := m.client.MakeBucket(ctx, bucket, miniosdk.MakeBucketOptions{}); err != nil {
		return fmt.Errorf("create minio bucket: %w", err)
	}
	return nil
}

func (m *Minio) DeleteBucket(
	ctx context.Context,
	bucket string,
) error {
	if bucket == "" {
		bucket = m.bucket
	}
	if err := m.client.RemoveBucket(ctx, bucket); err != nil {
		return fmt.Errorf("delete minio bucket: %w", err)
	}
	return nil
}

func (m *Minio) Upload(
	ctx context.Context,
	req *repo.ObjectStorageUploadReq,
) (*repo.ObjectStorageUploadResp, error) {
	if req == nil {
		req = &repo.ObjectStorageUploadReq{}
	}
	info, err := m.client.PutObject(ctx, m.bucket, req.Key, bytes.NewReader(req.Content), int64(len(req.Content)), miniosdk.PutObjectOptions{
		ContentType: req.MimeType,
	})
	if err != nil {
		return nil, fmt.Errorf("upload minio object: %w", err)
	}
	return &repo.ObjectStorageUploadResp{
		Provider: enum.ObjectStorageProviderMinio,
		Bucket:   m.bucket,
		Key:      req.Key,
		MimeType: req.MimeType,
		Size:     info.Size,
		Hash:     info.ETag,
	}, nil
}

func (m *Minio) StreamUpload(
	ctx context.Context,
	req *repo.ObjectStorageStreamUploadReq,
) (*repo.ObjectStorageStreamUploadResp, error) {
	if req == nil {
		req = &repo.ObjectStorageStreamUploadReq{}
	}
	size := req.Size
	if size <= 0 {
		size = -1
	}
	info, err := m.client.PutObject(ctx, m.bucket, req.Key, req.Body, size, miniosdk.PutObjectOptions{
		ContentType: req.MimeType,
	})
	if err != nil {
		return nil, fmt.Errorf("stream upload minio object: %w", err)
	}
	return &repo.ObjectStorageStreamUploadResp{
		Provider: enum.ObjectStorageProviderMinio,
		Bucket:   m.bucket,
		Key:      req.Key,
		MimeType: req.MimeType,
		Size:     info.Size,
		Hash:     info.ETag,
	}, nil
}

func (m *Minio) Download(
	ctx context.Context,
	key string,
) (*repo.ObjectStorageDownloadResp, error) {
	object, err := m.client.GetObject(ctx, m.bucket, key, miniosdk.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("download minio object: %w", err)
	}
	defer object.Close()
	stat, err := object.Stat()
	if err != nil {
		return nil, fmt.Errorf("stat minio object: %w", err)
	}
	content, err := io.ReadAll(object)
	if err != nil {
		return nil, fmt.Errorf("read minio object: %w", err)
	}
	return &repo.ObjectStorageDownloadResp{
		Key:      key,
		MimeType: stat.ContentType,
		Size:     stat.Size,
		Content:  content,
	}, nil
}

func (m *Minio) StreamDownload(
	ctx context.Context,
	key string,
) (*repo.ObjectStorageStreamDownloadResp, error) {
	object, err := m.client.GetObject(ctx, m.bucket, key, miniosdk.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("stream download minio object: %w", err)
	}
	stat, err := object.Stat()
	if err != nil {
		_ = object.Close()
		return nil, fmt.Errorf("stat minio object: %w", err)
	}
	return &repo.ObjectStorageStreamDownloadResp{
		Key:      key,
		MimeType: stat.ContentType,
		Size:     stat.Size,
		Body:     object,
	}, nil
}

func (m *Minio) UploadToken(
	ctx context.Context,
	req *repo.ObjectStorageUploadTokenReq,
) (string, error) {
	if req == nil {
		req = &repo.ObjectStorageUploadTokenReq{}
	}
	ctx, cancel := context.WithTimeout(ctx, m.uploadTimeout)
	defer cancel()
	url, err := m.client.PresignedPutObject(ctx, m.bucket, req.Key, m.uploadTimeout)
	if err != nil {
		return "", fmt.Errorf("create minio upload token: %w", err)
	}
	return url.String(), nil
}

func (m *Minio) Status(
	ctx context.Context,
	req *repo.ObjectStorageStatusReq,
) error {
	return nil
}
