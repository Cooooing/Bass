package repo

import (
	"context"
	"time"
)

type ContentBreezemoonClient interface {
	CreateBreezemoon(ctx context.Context, req *CreateBreezemoonReq) (*Breezemoon, error)
	PageBreezemoons(ctx context.Context, req *PageBreezemoonsReq) (*PageBreezemoonsResp, error)
}

type Breezemoon struct {
	ID        int64
	Content   string
	AuthorID  int64
	City      *string
	CreatedAt *time.Time
	Author    *AccountProfile
	ShowCity  bool
}

type CreateBreezemoonReq struct {
	Content  string
	AuthorID int64
	City     *string
}

type PageBreezemoonsReq struct {
	AuthorIDs []int64
	Cursor    *string
	Size      uint32
}

type PageBreezemoonsResp struct {
	Rows       []*Breezemoon
	NextCursor *string
}
