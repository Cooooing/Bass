package repo

import (
	"context"

	"bff_bbs/internal/biz/model"
)

type ContentMoonbreezeClient interface {
	CreateMoonbreeze(ctx context.Context, req *CreateMoonbreezeReq) (*model.Moonbreeze, error)
	PageMoonbreezes(ctx context.Context, req *PageMoonbreezesReq) (*model.PageMoonbreezesResp, error)
}

type CreateMoonbreezeReq struct {
	Content  string
	AuthorID int64
	City     *string
}

type PageMoonbreezesReq struct {
	AuthorIDs         []int64
	ExcludedAuthorIDs []int64
	Cursor            *string
	Size              uint32
}
