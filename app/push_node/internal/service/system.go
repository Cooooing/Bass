package service

import (
	"context"
	"fmt"

	"push_node/internal/conf"

	"github.com/go-kratos/kratos/v3/transport/http"
)

// SystemService 提供健康检查 HTTP 端点。
type SystemService struct {
	conf *conf.Bootstrap
}

// NewSystemService 创建 SystemService。
func NewSystemService(conf *conf.Bootstrap) *SystemService {
	return &SystemService{conf: conf}
}

// RegisterHttp 注册 HTTP 路由。
func (s *SystemService) RegisterHttp(hs *http.Server) {
	hs.HandleFunc("/push/v1/system/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		msg := fmt.Sprintf("%s %s is ok", s.conf.Server.Name, s.conf.Server.Version)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"code":0,"message":"%s","data":null}`, msg)))
	})
}

// Health 健康检查（保留以兼容接口）。
func (s *SystemService) Health(ctx context.Context) string {
	return fmt.Sprintf("%s %s is ok", s.conf.Server.Name, s.conf.Server.Version)
}
