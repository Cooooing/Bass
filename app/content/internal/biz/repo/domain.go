package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/content/v1"
	commonModel "common/pkg/model"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type DomainRepo interface {
	Save(ctx context.Context, tx *gen.Client, domain *model.Domain) (*model.Domain, error)
	Saves(ctx context.Context, tx *gen.Client, domains []*model.Domain) ([]*model.Domain, error)
	Update(ctx context.Context, tx *gen.Client, domain *model.Domain) (*model.Domain, error)

	GetOne(ctx context.Context, tx *gen.Client, req *DomainGetReq) (*model.Domain, error)
	GetList(ctx context.Context, tx *gen.Client, req *DomainGetReq) ([]*model.Domain, error)
	GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *DomainGetReq) ([]*model.Domain, *common.PageReply, error)
}

type DomainGetReq struct {
	DomainId    *int64
	DomainIds   []int64
	Name        *string
	Description *string
	Status      *v1.DomainStatus
	Url         *string
	Icon        *string
	TagCount    *commonModel.Range[int32]
	IsNav       *bool
}
