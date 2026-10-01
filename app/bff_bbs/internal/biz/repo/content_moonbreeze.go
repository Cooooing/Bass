package repo

import (
	"context"
	"time"
)

type ContentMoonbreezeClient interface {
	CreateMoonbreeze(ctx context.Context, req *CreateMoonbreezeReq) (*Moonbreeze, error)
	PageMoonbreezes(ctx context.Context, req *PageMoonbreezesReq) (*PageMoonbreezesResp, error)
}

type Moonbreeze struct {
	ID        int64
	Content   string
	AuthorID  int64
	City      *string
	CreatedAt *time.Time
	Author    *AccountProfile
	ShowCity  bool
}

type CreateMoonbreezeReq struct {
	Content  string
	AuthorID int64
	City     *string
}

type PageMoonbreezesReq struct {
	AuthorIDs []int64
	Cursor    *string
	Size      uint32
}

type PageMoonbreezesResp struct {
	Rows       []*Moonbreeze
	NextCursor *string
}
