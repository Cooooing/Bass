package model

import (
	v1 "common/api/content/v1"
	"content/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticlePostscript gen.ArticlePostscript

func (p ArticlePostscript) ConvertToRpc() *v1.ArticlePostscript {
	return &v1.ArticlePostscript{
		CreatedAt: timestamppb.New(*p.CreatedAt),
		UpdatedAt: timestamppb.New(*p.UpdatedAt),
		CreatedBy: p.CreatedBy,
		UpdatedBy: p.UpdatedBy,
		Id:        p.ID,
		ArticleId: p.ArticleID,
		Content:   p.Content,
	}
}
