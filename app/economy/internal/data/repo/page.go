package repo

import "economy/internal/biz/base"

type pageNormalizer struct{}

func (pageNormalizer) normalizePage(p *base.PageRequest) *base.PageRequest {
	if p == nil {
		return &base.PageRequest{Page: 1, Size: 10}
	}
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Size <= 0 {
		p.Size = 10
	}
	if p.Size > 100 {
		p.Size = 100
	}
	return p
}
