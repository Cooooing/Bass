package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/content/v1"
	"content/internal/biz/model"
	"context"
)

type DomainRepo interface {
	Save(ctx context.Context, domain *model.Domain) (*model.Domain, error)
	Saves(ctx context.Context, domains []*model.Domain) ([]*model.Domain, error)
	Update(ctx context.Context, domain *model.Domain) (*model.Domain, error)

	Get(ctx context.Context, req *DomainGetReq) (*model.Domain, error)
	GetList(ctx context.Context, req *DomainGetReq) ([]*model.Domain, error)
	Page(ctx context.Context, page *common.PageRequest, req *DomainGetReq) ([]*model.Domain, *common.PageReply, error)
}

type DomainGetReq struct {
	DomainId    *int64
	DomainIds   []int64
	Name        *string
	Description *string
	Status      *v1.DomainStatus
	Url         *string
	Icon        *string
	IsNav       *bool
}
