package repo

import "notify/internal/biz/base"

const (
	defaultPage int64 = 1
	defaultSize int64 = 10
	maxPageSize int64 = 1000
)

func normalizePage(p *base.PageRequest) *base.PageRequest {
	if p == nil {
		return &base.PageRequest{Page: defaultPage, Size: defaultSize}
	}
	if p.Page <= 0 {
		p.Page = defaultPage
	}
	if p.Size <= 0 {
		p.Size = defaultSize
	}
	if p.Size > maxPageSize {
		p.Size = maxPageSize
	}
	return p
}
