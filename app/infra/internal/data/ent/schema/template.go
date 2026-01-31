package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// infra 文章实体定义
type infra struct {
	ent.Schema
}

func (infra) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixinfra.String() + "infras"},
		entsql.WithComments(true),
	}
}

func (infra) Fields() []ent.Field {
	fields := []ent.Field{}
	return fields
}

func (infra) Edges() []ent.Edge {
	return []ent.Edge{}
}
