package repo

import (
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
)

type DomainRepo interface {
	Save(ctx context.Context, req *DomainSaveReq) (*DomainSaveResponse, error)
	Saves(ctx context.Context, req *DomainSavesReq) (*DomainSavesResponse, error)
	Update(ctx context.Context, req *DomainUpdateReq) (*DomainUpdateResponse, error)
	Get(ctx context.Context, req *DomainGetReq) (*DomainGetResponse, error)
	List(ctx context.Context, req *DomainGetReq) (*DomainListResponse, error)
	Map(ctx context.Context, req *DomainGetReq) (*DomainMapResponse, error)
	Count(ctx context.Context, req *DomainGetReq) (*DomainCountResponse, error)
	Page(ctx context.Context, req *DomainGetReq) (*DomainPageResponse, error)
}

type DomainSaveReq struct {
	Domain *model.Domain
}

type DomainSaveResponse struct {
	Domain *model.Domain
}

type DomainSavesReq struct {
	Domains []*model.Domain
}

type DomainSavesResponse struct {
	Rows []*model.Domain
}

type DomainUpdateReq struct {
	Domain *model.Domain
}

type DomainUpdateResponse struct {
	Domain *model.Domain
}

type DomainGetResponse struct {
	Domain *model.Domain
}

type DomainListResponse struct {
	Rows []*model.Domain
}

type DomainMapResponse struct {
	Rows map[int64]*model.Domain
}

type DomainCountResponse struct {
	Count int
}

type DomainPageResponse struct {
	Rows []*model.Domain
	Page *base.PageResponse
}

type DomainGetReq struct {
	Page        *base.PageRequest
	DomainId    *int64
	DomainIds   []int64
	Name        *string
	Description *string
	Status      *enum.DomainStatus
	Url         *string
	Icon        *string
	IsNav       *bool
}
