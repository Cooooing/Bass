package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/content/v1"
	"content/internal/biz/model"
	"context"
)

type TagRepo interface {
	Save(ctx context.Context, tag *model.Tag) (*model.Tag, error)
	Saves(ctx context.Context, tags []*model.Tag) ([]*model.Tag, error)
	Update(ctx context.Context, tag *model.Tag) (*model.Tag, error)

	Get(ctx context.Context, req *TagGetReq) (*model.Tag, error)
	GetList(ctx context.Context, req *TagGetReq) ([]*model.Tag, error)
	GetPage(ctx context.Context, page *common.PageRequest, req *TagGetReq) ([]*model.Tag, *common.PageReply, error)
}

type TagGetReq struct {
	TagId       *int64
	TagIds      []int64
	UserId      *int64
	Name        *string
	Names       []string
	Description *string
	Status      *v1.TagStatus
	DomainId    *int64
}
