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

	"common/pkg/util"

	"platform/internal/biz/model"

	"platform/internal/biz/repo"
	"platform/internal/config"

	"github.com/lionsoul2014/ip2region/binding/golang/service"
)

type IPResolutionUsecase struct {
	conf          *config.Bootstrap
	log           *slog.Logger
	httpClient    *http.Client
	storageClient repo.StorageClient
	ipDataDir     string
	mu            sync.RWMutex
	ip2region     *service.Ip2Region
}

func NewIPResolutionUsecase(
	conf *config.Bootstrap,
	logger *slog.Logger,
	httpClient *http.Client,
	storageClient repo.StorageClient,
) (*IPResolutionUsecase, func(), error) {
	u := &IPResolutionUsecase{
		conf:          conf,
		log:           logger,
		httpClient:    httpClient,
		storageClient: storageClient,
		ipDataDir:     "configs/ip2region",
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
	return u, cleanup, nil
}

// Start 在应用生命周期启动时加载 IP 数据集，避免在依赖装配阶段创建根 context。
func (u *IPResolutionUsecase) Start(ctx context.Context) error {
	if !u.conf.GetBusiness().GetIpData().GetEnable() {
		return nil
	}
	ipData := u.conf.GetBusiness().GetIpData()
	ipv4Content, ipv4Err := u.downloadIPDataFromStorage(ctx, ipData.GetIpv4XdbPath())
	ipv6Content, ipv6Err := u.downloadIPDataFromStorage(ctx, ipData.GetIpv6XdbPath())
	if ipv4Err == nil && ipv6Err == nil {
		err := u.uploadIPDataToLocal(ctx, ipv4Content, ipv6Content)
		if err != nil {
			return fmt.Errorf("upload ip data to local: %w", err)
		}
		return nil
	}
	u.log.Warn("download ip data from object storage failed, use local cache", "ipv4_error", ipv4Err, "ipv6_error", ipv6Err)
	if err := u.loadIPDataFromLocal(ctx); err != nil {
		u.log.Warn("load local ip data cache failed, IP resolution is unavailable", "error", err)
	}
	return nil
}

func (u *IPResolutionUsecase) downloadIPDataFromSource(ctx context.Context, url string) (content []byte, err error) {
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
	defer func() {
		if closeErr := httpResp.Body.Close(); err == nil && closeErr != nil {
			err = fmt.Errorf("close ip data source response body: %w", closeErr)
		}
	}()
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return nil, fmt.Errorf("download ip data source %s: status %d", url, httpResp.StatusCode)
	}
	content, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read ip data source %s: %w", url, err)
	}
	if len(content) == 0 {
		return nil, fmt.Errorf("ip data source %s is empty", url)
	}
	return content, nil
}

func (u *IPResolutionUsecase) downloadIPDataFromStorage(ctx context.Context, key string) ([]byte, error) {
	key = strings.TrimSpace(key)
	if key == "" {
		return nil, fmt.Errorf("ip data oss object key is empty")
	}
	downloadResp, err := u.storageClient.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("download IP data from object storage %s: %w", key, err)
	}
	if len(downloadResp.Content) == 0 {
		return nil, fmt.Errorf("IP data from object storage %s is empty", key)
	}
	return downloadResp.Content, nil
}

func (u *IPResolutionUsecase) uploadIPDataToStorage(ctx context.Context, key string, content []byte) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return fmt.Errorf("ip data oss object key is empty")
	}
	if len(content) == 0 {
		return fmt.Errorf("ip data content is empty")
	}
	current, err := u.downloadIPDataFromStorage(ctx, key)
	if err == nil && util.Sha256Bytes(current) == util.Sha256Bytes(content) {
		return nil
	}
	_, err = u.storageClient.Put(ctx, &repo.PutStoredObjectReq{
		ObjectKey: key,
		MimeType:  "application/octet-stream",
		Content:   content,
	})
	if err != nil {
		return fmt.Errorf("upload IP data to object storage %s: %w", key, err)
	}
	return nil
}

func (u *IPResolutionUsecase) uploadIPDataToLocal(ctx context.Context, ipv4Content []byte, ipv6Content []byte) error {
	ipData := u.conf.GetBusiness().GetIpData()
	ipv4Path, err := u.localIPDataPath(ipData.GetIpv4XdbPath())
	if err != nil {
		return fmt.Errorf("validate IPv4 ip data path: %w", err)
	}
	ipv6Path, err := u.localIPDataPath(ipData.GetIpv6XdbPath())
	if err != nil {
		return fmt.Errorf("validate IPv6 ip data path: %w", err)
	}
	if len(ipv4Content) == 0 && len(ipv6Content) == 0 {
		return fmt.Errorf("ip data content is empty")
	}
	ipv4Changed := false
	if len(ipv4Content) > 0 {
		changed, err := u.writeIPDataFile(ctx, ipv4Path, ipv4Content)
		if err != nil {
			return fmt.Errorf("write IPv4 ip data file: %w", err)
		}
		ipv4Changed = changed
	}
	ipv6Changed := false
	if len(ipv6Content) > 0 {
		changed, err := u.writeIPDataFile(ctx, ipv6Path, ipv6Content)
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

func (u *IPResolutionUsecase) loadIPDataFromLocal(ctx context.Context) error {
	ipData := u.conf.GetBusiness().GetIpData()
	ipv4Path, err := u.localIPDataPath(ipData.GetIpv4XdbPath())
	if err != nil {
		return fmt.Errorf("validate IPv4 ip data path: %w", err)
	}
	ipv6Path, err := u.localIPDataPath(ipData.GetIpv6XdbPath())
	if err != nil {
		return fmt.Errorf("validate IPv6 ip data path: %w", err)
	}
	// #nosec G304 -- 路径已限制为项目内 configs/ip2region 的 .xdb 数据文件。
	ipv4Content, err := os.ReadFile(ipv4Path)
	if err != nil {
		return fmt.Errorf("read local IPv4 ip data %s: %w", ipv4Path, err)
	}
	// #nosec G304 -- 路径已限制为项目内 configs/ip2region 的 .xdb 数据文件。
	ipv6Content, err := os.ReadFile(ipv6Path)
	if err != nil {
		return fmt.Errorf("read local IPv6 ip data %s: %w", ipv6Path, err)
	}
	return u.uploadIPDataToLocal(ctx, ipv4Content, ipv6Content)
}

func (u *IPResolutionUsecase) writeIPDataFile(ctx context.Context, path string, content []byte) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if localHash, err := util.Sha256File(path); err == nil && localHash == util.Sha256Bytes(content) {
		return false, nil
	} else if err != nil && !os.IsNotExist(err) {
		return false, fmt.Errorf("hash local ip data file %s: %w", path, err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return false, fmt.Errorf("create ip data dir %s: %w", filepath.Dir(path), err)
	}
	tmpPath := path + ".tmp"
	// #nosec G703 -- path 已由 localIPDataPath 验证，tmpPath 仅追加固定后缀。
	if err := os.WriteFile(tmpPath, content, 0o600); err != nil {
		return false, fmt.Errorf("write ip data temp file %s: %w", tmpPath, err)
	}
	if err := os.Rename(tmpPath, path); err != nil {
		if cleanupErr := util.RemoveFile(tmpPath); cleanupErr != nil && !os.IsNotExist(cleanupErr) {
			u.log.Warn("clean IP data temp file failed", "error", cleanupErr, "path", tmpPath)
		}
		return false, fmt.Errorf("replace ip data file %s: %w", path, err)
	}
	return true, nil
}

// localIPDataPath 仅接受项目内 IP 数据目录的相对 xdb 文件，阻止配置穿越到任意路径。
func (u *IPResolutionUsecase) localIPDataPath(path string) (string, error) {
	path = filepath.Clean(strings.TrimSpace(path))
	if path == "." || filepath.IsAbs(path) || filepath.Ext(path) != ".xdb" {
		return "", fmt.Errorf("invalid local ip data path")
	}
	relativePath, err := filepath.Rel(u.ipDataDir, path)
	if err != nil || relativePath == ".." || strings.HasPrefix(relativePath, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("ip data path must be under %s", u.ipDataDir)
	}
	return path, nil
}

func (u *IPResolutionUsecase) Get(ctx context.Context, ip string) (*model.IpInfo, error) {
	def := "unknown"
	u.mu.RLock()
	defer u.mu.RUnlock()
	if u.ip2region == nil || ip == "" {
		return &model.IpInfo{
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
	if strings.TrimSpace(region) == "" {
		return &model.IpInfo{
			Ip:          ip,
			Country:     def,
			Province:    def,
			City:        def,
			ISP:         def,
			CountryCode: def,
		}, nil
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
	return &model.IpInfo{
		Ip:          ip,
		Country:     clean(parts[0]),
		Province:    clean(parts[1]),
		City:        clean(parts[2]),
		ISP:         clean(parts[3]),
		CountryCode: clean(parts[4]),
	}, nil
}

func (u *IPResolutionUsecase) UpdateIPDataFromSource(ctx context.Context) error {
	ipData := u.conf.GetBusiness().GetIpData()
	ipv4Content, err := u.downloadIPDataFromSource(ctx, ipData.GetIpv4SourceUrl())
	if err != nil {
		return fmt.Errorf("download IPv4 ip data from source: %w", err)
	}
	ipv6Content, err := u.downloadIPDataFromSource(ctx, ipData.GetIpv6SourceUrl())
	if err != nil {
		return fmt.Errorf("download IPv6 ip data from source: %w", err)
	}
	if err := u.uploadIPDataToStorage(ctx, ipData.GetIpv4XdbPath(), ipv4Content); err != nil {
		return fmt.Errorf("upload IPv4 ip data to oss: %w", err)
	}
	if err := u.uploadIPDataToStorage(ctx, ipData.GetIpv6XdbPath(), ipv6Content); err != nil {
		return fmt.Errorf("upload IPv6 ip data to oss: %w", err)
	}
	return u.uploadIPDataToLocal(ctx, ipv4Content, ipv6Content)
}

func (u *IPResolutionUsecase) UpdateIPDataFromOss(ctx context.Context) error {
	ipData := u.conf.GetBusiness().GetIpData()
	ipv4Content, err := u.downloadIPDataFromStorage(ctx, ipData.GetIpv4XdbPath())
	if err != nil {
		return fmt.Errorf("download IPv4 ip data from oss: %w", err)
	}
	ipv6Content, err := u.downloadIPDataFromStorage(ctx, ipData.GetIpv6XdbPath())
	if err != nil {
		return fmt.Errorf("download IPv6 ip data from oss: %w", err)
	}
	return u.uploadIPDataToLocal(ctx, ipv4Content, ipv6Content)
}
