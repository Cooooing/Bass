package usecase

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	commonModel "common/pkg/model"
	"common/pkg/util"
	"platform/internal/config"

	"github.com/lionsoul2014/ip2region/binding/golang/service"
)

type IpResolutionUsecase struct {
	conf                 *config.Bootstrap
	log                  *slog.Logger
	httpClient           *http.Client
	objectStorageUsecase *ObjectStorageUsecase
	mu                   sync.RWMutex
	ip2region            *service.Ip2Region
}

func NewIpResolutionUsecase(
	conf *config.Bootstrap,
	logger *slog.Logger,
	httpClient *http.Client,
	objectStorageUsecase *ObjectStorageUsecase,
) (*IpResolutionUsecase, func(), error) {
	ctx := context.Background()
	u := &IpResolutionUsecase{
		conf:                 conf,
		log:                  logger,
		httpClient:           httpClient,
		objectStorageUsecase: objectStorageUsecase,
	}
	cleanup := func() {
		u.mu.Lock()
		old := u.ip2region
		u.ip2region = nil
		u.mu.Unlock()
		if old != nil {
			old.Close()
		}
	}
	if !conf.GetPlatform().GetIpData().GetEnable() {
		return u, cleanup, nil
	}
	ipData := u.conf.GetPlatform().GetIpData()
	ipv4Content, ipv4Err := u.downloadIpDataFromOss(ctx, ipData.GetIpv4XdbPath())
	ipv6Content, ipv6Err := u.downloadIpDataFromOss(ctx, ipData.GetIpv6XdbPath())
	if ipv4Err == nil && ipv6Err == nil {
		err := u.uploadIpDataToLocal(ctx, ipv4Content, ipv6Content)
		if err != nil {
			return u, cleanup, fmt.Errorf("upload ip data to local: %w", err)
		}
		return u, cleanup, nil
	}
	u.log.Warn("download ip data from oss failed, use source", "ipv4_error", ipv4Err, "ipv6_error", ipv6Err)
	err := u.UpdateIpDataFromSource(ctx)
	if err != nil {
		return u, cleanup, fmt.Errorf("update ip data from source: %w", err)
	}
	return u, cleanup, nil
}

func (u *IpResolutionUsecase) downloadIpDataFromSource(
	ctx context.Context,
	url string,
) ([]byte, error) {
	url = strings.TrimSpace(url)
	if url == "" {
		return nil, fmt.Errorf("ip data source url is empty")
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create ip data source request: %w", err)
	}
	client := u.httpClient
	if client == nil {
		client = http.DefaultClient
	}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("download ip data source %s: %w", url, err)
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return nil, fmt.Errorf("download ip data source %s: status %d", url, httpResp.StatusCode)
	}
	content, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read ip data source %s: %w", url, err)
	}
	if len(content) == 0 {
		return nil, fmt.Errorf("ip data source %s is empty", url)
	}
	return content, nil
}

func (u *IpResolutionUsecase) downloadIpDataFromOss(
	ctx context.Context,
	key string,
) ([]byte, error) {
	key = strings.TrimSpace(key)
	if key == "" {
		return nil, fmt.Errorf("ip data oss object key is empty")
	}
	downloadResp, err := u.objectStorageUsecase.Download(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("download ip data from oss %s: %w", key, err)
	}
	if len(downloadResp.Content) == 0 {
		return nil, fmt.Errorf("ip data from oss %s is empty", key)
	}
	return downloadResp.Content, nil
}

func (u *IpResolutionUsecase) uploadIpDataToOss(
	ctx context.Context,
	key string,
	content []byte,
) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return fmt.Errorf("ip data oss object key is empty")
	}
	if len(content) == 0 {
		return fmt.Errorf("ip data content is empty")
	}
	current, err := u.downloadIpDataFromOss(ctx, key)
	if err == nil && util.Sha256Bytes(current) == util.Sha256Bytes(content) {
		return nil
	}
	_, err = u.objectStorageUsecase.Upload(ctx, &UploadReq{
		Key:      key,
		FileName: filepath.Base(key),
		MimeType: "application/octet-stream",
		Content:  content,
	})
	if err != nil {
		return fmt.Errorf("upload ip data to oss %s: %w", key, err)
	}
	return nil
}

func (u *IpResolutionUsecase) uploadIpDataToLocal(
	ctx context.Context,
	ipv4Content []byte,
	ipv6Content []byte,
) error {
	ipData := u.conf.GetPlatform().GetIpData()
	ipv4Path := strings.TrimSpace(ipData.GetIpv4XdbPath())
	if ipv4Path == "" {
		return fmt.Errorf("IPv4 ip data path is empty")
	}
	ipv6Path := strings.TrimSpace(ipData.GetIpv6XdbPath())
	if ipv6Path == "" {
		return fmt.Errorf("IPv6 ip data path is empty")
	}
	if len(ipv4Content) == 0 && len(ipv6Content) == 0 {
		return fmt.Errorf("ip data content is empty")
	}
	ipv4Changed := false
	if len(ipv4Content) > 0 {
		changed, err := u.writeIpDataFile(ctx, ipv4Path, ipv4Content)
		if err != nil {
			return fmt.Errorf("write IPv4 ip data file: %w", err)
		}
		ipv4Changed = changed
	}
	ipv6Changed := false
	if len(ipv6Content) > 0 {
		changed, err := u.writeIpDataFile(ctx, ipv6Path, ipv6Content)
		if err != nil {
			return fmt.Errorf("write IPv6 ip data file: %w", err)
		}
		ipv6Changed = changed
	}
	u.mu.RLock()
	needsReload := u.ip2region == nil || ipv4Changed || ipv6Changed
	u.mu.RUnlock()
	if !needsReload {
		return nil
	}

	cachePolicy := service.VIndexCache
	cachePolicy, err := service.CachePolicyFromName(strings.ToLower(strings.TrimSpace(ipData.GetCachePolicy())))
	if err != nil {
		return fmt.Errorf("invalid ip data cache policy: %s", ipData.GetCachePolicy())
	}
	v4Config, err := service.NewV4Config(cachePolicy, ipv4Path, 20)
	if err != nil {
		return fmt.Errorf("create IPv4 config: %w", err)
	}
	v6Config, err := service.NewV6Config(cachePolicy, ipv6Path, 20)
	if err != nil {
		return fmt.Errorf("create IPv6 config: %w", err)
	}
	next, err := service.NewIp2Region(v4Config, v6Config)
	if err != nil {
		return fmt.Errorf("create ip2region service: %w", err)
	}
	u.mu.Lock()
	old := u.ip2region
	u.ip2region = next
	u.mu.Unlock()
	if old != nil {
		old.Close()
	}
	return nil
}

func (u *IpResolutionUsecase) writeIpDataFile(
	ctx context.Context,
	path string,
	content []byte,
) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if localHash, err := util.Sha256File(path); err == nil && localHash == util.Sha256Bytes(content) {
		return false, nil
	} else if err != nil && !os.IsNotExist(err) {
		return false, fmt.Errorf("hash local ip data file %s: %w", path, err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return false, fmt.Errorf("create ip data dir %s: %w", filepath.Dir(path), err)
	}
	tmpPath := path + ".tmp"
	if err := os.WriteFile(tmpPath, content, 0o644); err != nil {
		return false, fmt.Errorf("write ip data temp file %s: %w", tmpPath, err)
	}
	if err := os.Rename(tmpPath, path); err != nil {
		_ = os.Remove(tmpPath)
		return false, fmt.Errorf("replace ip data file %s: %w", path, err)
	}
	return true, nil
}

func (u *IpResolutionUsecase) Get(
	ctx context.Context,
	ip string,
) (*commonModel.IpInfo, error) {
	def := "unknown"
	u.mu.RLock()
	defer u.mu.RUnlock()
	if u.ip2region == nil || ip == "" {
		return &commonModel.IpInfo{
			Ip:          ip,
			Country:     def,
			Province:    def,
			City:        def,
			ISP:         def,
			CountryCode: def,
		}, nil
	}
	region, err := u.ip2region.SearchByStr(ip)
	if err != nil {
		return nil, fmt.Errorf("resolve IP: %w", err)
	}
	parts := strings.Split(region, "|")
	if len(parts) < 5 {
		return nil, fmt.Errorf("invalid IP region: %s", region)
	}
	clean := func(s string) string {
		if s == "0" {
			return def
		}
		return s
	}
	return &commonModel.IpInfo{
		Ip:          ip,
		Country:     clean(parts[0]),
		Province:    clean(parts[1]),
		City:        clean(parts[2]),
		ISP:         clean(parts[3]),
		CountryCode: clean(parts[4]),
	}, nil
}

func (u *IpResolutionUsecase) UpdateIpDataFromSource(
	ctx context.Context,
) error {
	ipData := u.conf.GetPlatform().GetIpData()
	ipv4Content, err := u.downloadIpDataFromSource(ctx, ipData.GetIpv4SourceUrl())
	if err != nil {
		return fmt.Errorf("download IPv4 ip data from source: %w", err)
	}
	ipv6Content, err := u.downloadIpDataFromSource(ctx, ipData.GetIpv6SourceUrl())
	if err != nil {
		return fmt.Errorf("download IPv6 ip data from source: %w", err)
	}
	if err := u.uploadIpDataToOss(ctx, ipData.GetIpv4XdbPath(), ipv4Content); err != nil {
		return fmt.Errorf("upload IPv4 ip data to oss: %w", err)
	}
	if err := u.uploadIpDataToOss(ctx, ipData.GetIpv6XdbPath(), ipv6Content); err != nil {
		return fmt.Errorf("upload IPv6 ip data to oss: %w", err)
	}
	return u.uploadIpDataToLocal(ctx, ipv4Content, ipv6Content)
}

func (u *IpResolutionUsecase) UpdateIpDataFromOss(
	ctx context.Context,
) error {
	ipData := u.conf.GetPlatform().GetIpData()
	ipv4Content, err := u.downloadIpDataFromOss(ctx, ipData.GetIpv4XdbPath())
	if err != nil {
		return fmt.Errorf("download IPv4 ip data from oss: %w", err)
	}
	ipv6Content, err := u.downloadIpDataFromOss(ctx, ipData.GetIpv6XdbPath())
	if err != nil {
		return fmt.Errorf("download IPv6 ip data from oss: %w", err)
	}
	return u.uploadIpDataToLocal(ctx, ipv4Content, ipv6Content)
}
