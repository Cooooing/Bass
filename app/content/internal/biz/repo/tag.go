package repo

import (
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
)

type TagRepo interface {
	Save(ctx context.Context, req *TagSaveReq) (*TagSaveResponse, error)
	Saves(ctx context.Context, req *TagSavesReq) (*TagSavesResponse, error)
	Update(ctx context.Context, req *TagUpdateReq) (*TagUpdateResponse, error)
	Get(ctx context.Context, req *TagGetReq) (*TagGetResponse, error)
	List(ctx context.Context, req *TagGetReq) (*TagListResponse, error)
	Map(ctx context.Context, req *TagGetReq) (*TagMapResponse, error)
	Count(ctx context.Context, req *TagGetReq) (*TagCountResponse, error)
	Page(ctx context.Context, req *TagGetReq) (*TagPageResponse, error)
}

type TagSaveReq struct {
	Tag *model.Tag
}

type TagSaveResponse struct {
	Tag *model.Tag
}

type TagSavesReq struct {
	Tags []*model.Tag
}

type TagSavesResponse struct {
	Rows []*model.Tag
}

type TagUpdateReq struct {
	Tag *model.Tag
}

type TagUpdateResponse struct {
	Tag *model.Tag
}

type TagGetResponse struct {
	Tag *model.Tag
}

type TagListResponse struct {
	Rows []*model.Tag
}

type TagMapResponse struct {
	Rows map[int64]*model.Tag
}

type TagCountResponse struct {
	Count int
}

type TagPageResponse struct {
	Rows []*model.Tag
	Page *base.PageResponse
}

type TagGetReq struct {
	Page        *base.PageRequest
	TagId       *int64
	TagIds      []int64
	UserId      *int64
	Name        *string
	Names       []string
	Description *string
	Status      *enum.TagStatus
	DomainId    *int64
}
