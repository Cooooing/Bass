package model

import commonpb "common/proto/gen/common"

const (
	DefaultPage     int64 = 1
	DefaultPageSize int64 = 10
	MaxPageSize     int64 = 1000
)

// PageReq describes offset pagination parameters.
type PageReq struct {
	Page int64
	Size int64
}

// PageResp describes an offset pagination result.
type PageResp struct {
	Page  int64
	Size  int64
	Total int64
}

// NormalizePage returns an independent, valid page request.
func NormalizePage(req *PageReq) *PageReq {
	result := &PageReq{
		Page: DefaultPage,
		Size: DefaultPageSize,
	}
	if req == nil {
		return result
	}
	if req.Page >= DefaultPage {
		result.Page = req.Page
	}
	if req.Size > 0 {
		result.Size = req.Size
	}
	if result.Size > MaxPageSize {
		result.Size = MaxPageSize
	}
	return result
}

// Limit returns a database-safe limit for a normalized page request.
func (p *PageReq) Limit() int {
	return int(NormalizePage(p).Size)
}

// Offset returns a database-safe offset for a normalized page request.
func (p *PageReq) Offset() int {
	page := NormalizePage(p)
	maxInt := int64(^uint(0) >> 1)
	if page.Page-DefaultPage > maxInt/page.Size {
		return int(maxInt)
	}
	return int((page.Page - DefaultPage) * page.Size)
}

// PageReqFromProto converts the shared protocol representation into its domain model.
func PageReqFromProto(page *commonpb.PageReq) *PageReq {
	if page == nil {
		return nil
	}
	return &PageReq{
		Page: page.GetPage(),
		Size: page.GetSize(),
	}
}

// ToProto converts the page request into its shared protocol representation.
func (p *PageReq) ToProto() *commonpb.PageReq {
	if p == nil {
		return nil
	}
	return &commonpb.PageReq{
		Page: p.Page,
		Size: p.Size,
	}
}

// PageRespFromProto converts the shared protocol representation into its domain model.
func PageRespFromProto(page *commonpb.PageResp) *PageResp {
	if page == nil {
		return nil
	}
	return &PageResp{
		Page:  page.GetPage(),
		Size:  page.GetSize(),
		Total: page.GetTotal(),
	}
}

// ToProto converts the page response into its shared protocol representation.
func (p *PageResp) ToProto() *commonpb.PageResp {
	if p == nil {
		return nil
	}
	return &commonpb.PageResp{
		Page:  p.Page,
		Size:  p.Size,
		Total: p.Total,
	}
}
