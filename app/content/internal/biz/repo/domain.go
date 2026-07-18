package repo

import (
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
)

type DomainRepo interface {
	Save(ctx context.Context, domain *model.Domain) (*model.Domain, error)
	Saves(ctx context.Context, domains []*model.Domain) ([]*model.Domain, error)
	Update(ctx context.Context, domain *model.Domain) (*model.Domain, error)
	Get(ctx context.Context, req *DomainGetReq) (*model.Domain, error)
	List(ctx context.Context, req *DomainGetReq) ([]*model.Domain, error)
	Map(ctx context.Context, req *DomainGetReq) (map[int64]*model.Domain, error)
	Count(ctx context.Context, req *DomainGetReq) (int, error)
	Page(ctx context.Context, req *DomainGetReq) (*DomainPageResp, error)
}

type DomainPageResp struct {
	Rows []*model.Domain
	Page *base.PageResp
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
