package repo

import (
	"common/proto/gen/common"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
)

type TagRepo interface {
	Save(ctx context.Context, tag *model.Tag) (*model.Tag, error)
	Saves(ctx context.Context, tags []*model.Tag) ([]*model.Tag, error)
	Update(ctx context.Context, tag *model.Tag) (*model.Tag, error)

	Get(ctx context.Context, req *TagGetReq) (*model.Tag, error)
	List(ctx context.Context, req *TagGetReq) ([]*model.Tag, error)
	Map(ctx context.Context, req *TagGetReq) (map[int64]*model.Tag, error)
	Count(ctx context.Context, req *TagGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *TagGetReq) ([]*model.Tag, *common.PageReply, error)
}

type TagGetReq struct {
	TagId       *int64
	TagIds      []int64
	UserId      *int64
	Name        *string
	Names       []string
	Description *string
	Status      *enum.TagStatus
	DomainId    *int64
}
