package server

import "common/proto/gen/common"

const (
	defaultPage uint32 = 1
	defaultSize uint32 = 10
	maxPageSize uint32 = 1000
)

func PageValid(p *common.PageReq) *common.PageReq {
	if p == nil {
		return GetPageDefault()
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

func GetPageDefault() *common.PageReq {
	return &common.PageReq{
		Page: defaultPage,
		Size: defaultSize,
	}
}

func GetPageMax() *common.PageReq {
	return &common.PageReq{
		Page: defaultPage,
		Size: maxPageSize,
	}
}
