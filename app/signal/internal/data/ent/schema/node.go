package schema

import (
	"common/pkg"
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Node 文章实体定义
type Node struct {
	ent.Schema
}

func (Node) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixSignal.String() + "nodes"},
		entsql.WithComments(true),
	}
}

func (Node) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("owner_id").Comment("节点拥有者 ID").Optional().Nillable(),
		field.String("key").Comment("节点标识，请求头唯一标识").NotEmpty(),
		field.String("name").Comment("节点名称").NotEmpty(),
		field.String("description").Comment("节点描述").Optional().Nillable(),
		field.String("secret").Comment("节点密钥"),
		field.String("callback_url").Comment("节点回调公网地址，如 example.com/api"),
		field.Int32("status").Comment("节点状态: 1-正常, 2-禁用").Default(1),
		field.Float("weight").Comment("节点权重").Default(1.0),
	}
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
}

func (Node) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").Unique(),
	}
}
