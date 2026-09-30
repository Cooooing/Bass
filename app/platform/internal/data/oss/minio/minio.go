package minio

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"platform/internal/biz/repo"
	"platform/internal/config"
	"platform/internal/enum"
	"strconv"
	"strings"
	"time"

	miniosdk "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var _ repo.StorageClient = (*Minio)(nil)

type Minio struct {
	client        *miniosdk.Client
	browserClient *miniosdk.Client
	bucket        string
	browserScheme string
	browserHost   string
	uploadTimeout time.Duration
}

func NewMinio(conf *config.Bootstrap) (*Minio, error) {
	src := conf.GetPlatform().GetOss().GetMinio()
	if src.GetEndpoint() == "" || src.GetAccessKey() == "" || src.GetSecretKey() == "" || src.GetBucket() == "" {
		return nil, fmt.Errorf("minio config requires endpoint, access_key, secret_key and bucket")
	}
	client, err := miniosdk.New(src.GetEndpoint(), &miniosdk.Options{
		Creds:  credentials.NewStaticV4(src.GetAccessKey(), src.GetSecretKey(), ""),
		Secure: src.GetUseSsl(),
	})
	if err != nil {
		return nil, fmt.Errorf("create MinIO client: %w", err)
	}
	browserEndpoint := src.GetBrowserEndpoint()
	if browserEndpoint == "" {
		// A local or single-network deployment exposes the same endpoint to the
		// service and browser. Only deployments with separate public routing need
		// to override browser_endpoint.
		browserEndpoint = src.GetEndpoint()
	}
	browserUseSSL := src.GetBrowserUseSsl()
	if browserEndpoint == src.GetEndpoint() && !browserUseSSL {
		browserUseSSL = src.GetUseSsl()
	}
	browserClient, err := miniosdk.New(browserEndpoint, &miniosdk.Options{
		Creds:  credentials.NewStaticV4(src.GetAccessKey(), src.GetSecretKey(), ""),
		Secure: browserUseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("create browser MinIO client: %w", err)
	}
	timeout := 30 * time.Minute
	if src.GetTimeout() != nil && src.GetTimeout().AsDuration() > 0 {
		timeout = src.GetTimeout().AsDuration()
	}
	m := &Minio{
		client:        client,
		browserClient: browserClient,
		bucket:        src.GetBucket(),
		browserHost:   browserEndpoint,
		browserScheme: map[bool]string{true: "https", false: "http"}[browserUseSSL],
		uploadTimeout: timeout,
	}
	return m, nil
}

func (m *Minio) Name() string {
	return enum.AssetProviderMinio.String()
}

func (m *Minio) Bucket() string {
	return m.bucket
}

func (m *Minio) PrepareDirectUpload(ctx context.Context, req *repo.DirectAssetUploadReq) (*repo.DirectAssetUploadResp, error) {
	if req == nil || req.Hash == "" || req.MimeType == "" || req.Size <= 0 || req.UploadByID <= 0 {
		return nil, fmt.Errorf("direct asset upload requires hash, MIME type, size and uploader")
	}
	policy := miniosdk.NewPostPolicy()
	if err := policy.SetExpires(time.Now().UTC().Add(m.uploadTimeout)); err != nil {
		return nil, fmt.Errorf("set upload policy expiry: %w", err)
	}
	if err := policy.SetBucket(m.bucket); err != nil {
		return nil, fmt.Errorf("set upload policy bucket: %w", err)
	}
	objectKey := AssetObjectKey(req.Hash)
	if err := policy.SetKey(objectKey); err != nil {
		return nil, fmt.Errorf("set upload policy object key: %w", err)
	}
	if err := policy.SetContentType(req.MimeType); err != nil {
		return nil, fmt.Errorf("set upload policy MIME type: %w", err)
	}
	if err := policy.SetContentLengthRange(req.Size, req.Size); err != nil {
		return nil, fmt.Errorf("set upload policy size: %w", err)
	}
	// The policy fixes uploader metadata before it reaches the browser. MinIO
	// returns it from StatObject, so the event consumer can record provenance
	// without an upload-session table or a business callback URL.
	if err := policy.SetUserMetadata("upload-by", strconv.FormatInt(req.UploadByID, 10)); err != nil {
		return nil, fmt.Errorf("set upload policy uploader metadata: %w", err)
	}
	digest, err := hexSHA256ToBase64(req.Hash)
	if err != nil {
		return nil, err
	}
	if err := policy.SetChecksum(miniosdk.NewChecksumString(miniosdk.ChecksumSHA256, digest)); err != nil {
		return nil, fmt.Errorf("set upload policy checksum: %w", err)
	}
	u, formFields, err := m.browserClient.PresignedPostPolicy(ctx, policy)
	if err != nil {
		return nil, fmt.Errorf("create MinIO POST policy: %w", err)
	}
	return &repo.DirectAssetUploadResp{URL: u.String(), FormFields: formFields}, nil
}

func (m *Minio) Stat(ctx context.Context, objectKey string) (*repo.AssetObject, error) {
	info, err := m.client.StatObject(ctx, m.bucket, objectKey, miniosdk.StatObjectOptions{Checksum: true})
	if err != nil {
		return nil, fmt.Errorf("stat MinIO object: %w", err)
	}
	// MinIO's HEAD response preserves the provider's canonical header casing
	// (for example Upload-By). Treat metadata names as HTTP headers instead of
	// relying on the lowercase key that was used in the POST policy.
	uploadByID, _ := strconv.ParseInt(userMetadataValue(info.UserMetadata, "upload-by"), 10, 64)
	return &repo.AssetObject{
		Provider:       enum.AssetProviderMinio,
		Bucket:         m.bucket,
		ObjectKey:      objectKey,
		MimeType:       info.ContentType,
		Size:           info.Size,
		ProviderETag:   info.ETag,
		ChecksumSHA256: info.ChecksumSHA256,
		UploadByID:     uploadByID,
	}, nil
}

func userMetadataValue(metadata map[string]string, name string) string {
	for key, value := range metadata {
		if strings.EqualFold(key, name) {
			return value
		}
	}
	return ""
}

func (m *Minio) PublicURL(objectKey string) string {
	if objectKey == "" || m.browserHost == "" {
		return ""
	}
	return (&url.URL{Scheme: m.browserScheme, Host: m.browserHost, Path: "/" + m.bucket + "/" + objectKey}).String()
}

func (m *Minio) Get(ctx context.Context, objectKey string) (*repo.StoredObject, error) {
	object, err := m.client.GetObject(ctx, m.bucket, objectKey, miniosdk.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("get MinIO object: %w", err)
	}
	defer object.Close()
	info, err := object.Stat()
	if err != nil {
		return nil, fmt.Errorf("stat MinIO object: %w", err)
	}
	content, err := io.ReadAll(object)
	if err != nil {
		return nil, fmt.Errorf("read MinIO object: %w", err)
	}
	return &repo.StoredObject{ObjectKey: objectKey, MimeType: info.ContentType, Size: info.Size, Content: content}, nil
}

func (m *Minio) Put(ctx context.Context, req *repo.PutStoredObjectReq) (*repo.StoredObject, error) {
	if req == nil || req.ObjectKey == "" {
		return nil, fmt.Errorf("stored object key is required")
	}
	info, err := m.client.PutObject(ctx, m.bucket, req.ObjectKey, bytes.NewReader(req.Content), int64(len(req.Content)), miniosdk.PutObjectOptions{ContentType: req.MimeType})
	if err != nil {
		return nil, fmt.Errorf("put MinIO object: %w", err)
	}
	return &repo.StoredObject{ObjectKey: req.ObjectKey, MimeType: req.MimeType, Size: info.Size}, nil
}

func AssetObjectKey(hash string) string {
	return "assets/sha256/" + hash
}

func hexSHA256ToBase64(hash string) (string, error) {
	bytes, err := hex.DecodeString(hash)
	if err != nil || len(bytes) != 32 {
		return "", fmt.Errorf("invalid SHA-256")
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
