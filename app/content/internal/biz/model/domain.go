package model

import (
	v1 "common/api/content/v1"
	"content/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Domain gen.Domain

func (t *Domain) ConvertToRpc() *v1.Domain {
	domain := &v1.Domain{
		CreatedAt:   timestamppb.New(*t.CreatedAt),
		UpdatedAt:   timestamppb.New(*t.UpdatedAt),
		CreatedBy:   t.CreatedBy,
		UpdatedBy:   t.UpdatedBy,
		Id:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Status:      t.Status,
		Url:         t.URL,
		Icon:        t.Icon,
		TagCount:    t.TagCount,
		IsNav:       t.IsNav,
	}
	entDomain := (*gen.Domain)(t)
	if len(entDomain.Edges.Tags) > 0 {
		for _, tag := range entDomain.Edges.Tags {
			domain.Tags = append(domain.Tags, (*Tag)(tag).ConvertToRpc())
		}
	}
	return domain
}
