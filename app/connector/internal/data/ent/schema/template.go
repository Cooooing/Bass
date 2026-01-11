package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// connector 文章实体定义
type connector struct {
	ent.Schema
}

func (connector) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixconnector.String() + "connectors"},
		entsql.WithComments(true),
	}
}

func (connector) Fields() []ent.Field {
	fields := []ent.Field{}
	return fields
}

func (connector) Edges() []ent.Edge {
	return []ent.Edge{}
}
