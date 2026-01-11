package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// IM 文章实体定义
type IM struct {
	ent.Schema
}

func (IM) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixIM.String() + "ims"},
		entsql.WithComments(true),
	}
}

func (IM) Fields() []ent.Field {
	fields := []ent.Field{}
	return fields
}

func (IM) Edges() []ent.Edge {
	return []ent.Edge{}
}
