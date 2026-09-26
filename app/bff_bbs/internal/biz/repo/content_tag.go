package repo

import (
	"context"
	"time"
)

type Tag struct {
	ID           int64
	Code         string
	Name         string
	Description  *string
	DomainID     *int64
	Status       *int32
	Icon         *string
	Sort         int32
	ArticleCount int32
	CreatedBy    *int64
	UpdatedBy    *int64
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type TagSave struct {
	Code        string
	Name        string
	Description *string
	DomainID    *int64
	Status      *int32
	Icon        *string
	Sort        int32
}

type TagQuery struct {
	IDs         []int64
	Code        *string
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

type UpdateTagReq struct {
	UserID int64
	TagID  int64
	Tag    *TagSave
}

type ListTagsReq struct {
	Page  *PageReq
	Query *TagQuery
}

type ListTagsResp struct {
	Page *PageResp
	Rows []*Tag
}

type BindArticleTagsReq struct {
	UserID    int64
	ArticleID int64
	TagIDs    []int64
}

type UnbindArticleTagsReq struct {
	UserID    int64
	ArticleID int64
	TagIDs    []int64
}

type ListArticleTagsReq struct {
	ArticleID int64
}

type ContentTagClient interface {
	CreateTag(ctx context.Context, req *CreateTagReq) (*Tag, error)
	UpdateTag(ctx context.Context, req *UpdateTagReq) (*Tag, error)
	BindArticleTags(ctx context.Context, req *BindArticleTagsReq) error
	UnbindArticleTags(ctx context.Context, req *UnbindArticleTagsReq) error
	ListArticleTags(ctx context.Context, req *ListArticleTagsReq) ([]*Tag, error)
	ListTags(ctx context.Context, req *ListTagsReq) (*ListTagsResp, error)
}
