package repo

import "context"

type Tag struct {
	ID          int64
	Name        string
	Description *string
	DomainID    *int64
	Status      *int32
	CreatedBy   *int64
	UpdatedBy   *int64
	CreatedAt   string
	UpdatedAt   string
}

type TagSave struct {
	Name        string
	Description *string
	DomainID    *int64
	Status      *int32
}

type TagQuery struct {
	IDs         []int64
	Name        *string
	Names       []string
	Description *string
	DomainID    *int64
	Status      *int32
}

type CreateTagReq struct {
	UserID int64
	Tag    *TagSave
}

type CreateTagResponse struct {
	Tag *Tag
}

type UpdateTagReq struct {
	UserID int64
	TagID  int64
	Tag    *TagSave
}

type UpdateTagResponse struct {
	Tag *Tag
}

type ListTagsReq struct {
	Page  *PageReq
	Query *TagQuery
}

type ListTagsResponse struct {
	Page *PageResponse
	Rows []*Tag
}

type ContentTagClient interface {
	CreateTag(ctx context.Context, req *CreateTagReq) (*CreateTagResponse, error)
	UpdateTag(ctx context.Context, req *UpdateTagReq) (*UpdateTagResponse, error)
	ListTags(ctx context.Context, req *ListTagsReq) (*ListTagsResponse, error)
}
