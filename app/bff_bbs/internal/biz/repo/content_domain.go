package repo

import (
	"context"
	"time"
)

type Domain struct {
	ID          int64
	Code        string
	Name        string
	Description *string
	Status      int32
	URL         *string
	Icon        *string
	IsNav       bool
	Sort        int32
	CreatedBy   *int64
	UpdatedBy   *int64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type DomainSave struct {
	Code        string
	Name        string
	Description *string
	Status      *int32
	URL         *string
	Icon        *string
	IsNav       bool
	Sort        int32
}

type DomainQuery struct {
	IDs         []int64
	Code        *string
	Name        *string
	Description *string
	URL         *string
	Icon        *string
	IsNav       *bool
	Status      *int32
}

type CreateDomainReq struct {
	UserID int64
	Domain *DomainSave
}

type UpdateDomainReq struct {
	UserID   int64
	DomainID int64
	Domain   *DomainSave
}

type ListDomainsReq struct {
	Page  *PageReq
	Query *DomainQuery
}

type ListDomainsResp struct {
	Page *PageResp
	Rows []*Domain
}

type ContentDomainClient interface {
	CreateDomain(ctx context.Context, req *CreateDomainReq) (*Domain, error)
	UpdateDomain(ctx context.Context, req *UpdateDomainReq) (*Domain, error)
	ListDomains(ctx context.Context, req *ListDomainsReq) (*ListDomainsResp, error)
}
