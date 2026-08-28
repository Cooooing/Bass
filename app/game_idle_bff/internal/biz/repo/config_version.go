package repo

import "context"

// ConfigVersionRepo 管理前端聚合配置版本缓存。
type ConfigVersionRepo interface {
	Get(ctx context.Context) (string, error)
	Save(ctx context.Context, version string) error
}
