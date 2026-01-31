package model

import (
	v1 "common/api/content/v1"
	"content/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Domain struct {
	*gen.Domain
}

func (t *Domain) ConvertToRpc() *v1.Domain {
	domain := &v1.Domain{
		CreatedAt:   timestamppb.New(*t.CreatedAt),
		UpdatedAt:   timestamppb.New(*t.UpdatedAt),
		CreatedBy:   t.CreatedBy,
		UpdatedBy:   t.UpdatedBy,
		Id:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Status:      v1.DomainStatus(t.Status),
		Url:         t.URL,
		Icon:        t.Icon,
		TagCount:    t.TagCount,
		IsNav:       t.IsNav,
	}
	if len(t.Domain.Edges.Tags) > 0 {
		for _, tag := range t.Domain.Edges.Tags {
			domain.Tags = append(domain.Tags, (&Tag{Tag: tag}).ConvertToRpc())
		}
	}
	return domain
}
