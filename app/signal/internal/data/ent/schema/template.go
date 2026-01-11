package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// Signal 文章实体定义
type Signal struct {
	ent.Schema
}

func (Signal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixSignal.String() + "signals"},
		entsql.WithComments(true),
	}
}

func (Signal) Fields() []ent.Field {
	fields := []ent.Field{}
	return fields
}

func (Signal) Edges() []ent.Edge {
	return []ent.Edge{}
}
