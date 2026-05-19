package usecase

import "context"

const defaultLanguage = "zh_CN"

func (s *NotifyUsecase) getLanguage(_ context.Context, _ int64) string {
	// TODO: Redis 缓存 → gRPC 回源 → 硬编码兜底
	return defaultLanguage
}
