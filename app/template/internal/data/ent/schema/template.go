package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// Template 文章实体定义
type Template struct {
	ent.Schema
}

func (Template) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixTemplate.String() + "templates"},
		entsql.WithComments(true),
	}
}

func (Template) Fields() []ent.Field {
	fields := []ent.Field{}
	return fields
}

func (Template) Edges() []ent.Edge {
	return []ent.Edge{}
}
