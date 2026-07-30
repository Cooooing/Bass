package repo

import (
	"content/internal/biz/base"
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
	AddArticleCount(ctx context.Context, req *TagAddArticleCountReq) error
	Page(ctx context.Context, req *TagGetReq) (*TagPageResp, error)
}

type TagPageResp struct {
	Rows []*model.Tag
	Page *base.PageResp
}

type TagGetReq struct {
	Page        *base.PageRequest
	TagId       *int64
	TagIds      []int64
	UserId      *int64
	Code        *string
	Name        *string
	Names       []string
	Description *string
	Status      *enum.TagStatus
	DomainId    *int64
}

type TagAddArticleCountReq struct {
	TagIDs []int64
	Delta  int32
}
