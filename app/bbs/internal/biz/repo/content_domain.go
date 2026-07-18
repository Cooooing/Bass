package repo

import "context"

type Domain struct {
	ID          int64
	Name        string
	Description *string
	Status      int32
	URL         *string
	Icon        *string
	IsNav       bool
	CreatedBy   *int64
	UpdatedBy   *int64
	CreatedAt   string
	UpdatedAt   string
}

type DomainQuery struct {
	IDs         []int64
	Name        *string
	Description *string
	URL         *string
	Icon        *string
	IsNav       *bool
	Status      *int32
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
	ListDomains(ctx context.Context, req *ListDomainsReq) (*ListDomainsResp, error)
}
