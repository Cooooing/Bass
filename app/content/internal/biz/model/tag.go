package model

import (
	v1 "common/api/content/v1"
	"common/pkg/cutil/base"
	"content/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Tag struct {
	*gen.Tag
}

func (t *Tag) ConvertToRpc() *v1.Tag {
	return &v1.Tag{
		CreatedAt:    timestamppb.New(*t.CreatedAt),
		UpdatedAt:    timestamppb.New(*t.UpdatedAt),
		CreatedBy:    t.CreatedBy,
		UpdatedBy:    t.UpdatedBy,
		Id:           t.ID,
		Name:         t.Name,
		Description:  t.Description,
		DomainId:     t.DomainID,
		Status:       base.Ptr(v1.TagStatus(t.Status)),
		ArticleCount: t.ArticleCount,
	}
}
