package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Moonbreeze 定义清风明月短动态实体。
type Moonbreeze struct {
	ent.Schema
}

func (Moonbreeze) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "moonbreezes"},
		entsql.WithComments(true),
	}
}

func (Moonbreeze) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Text("content").Comment("动态正文").NotEmpty(),
		field.Int64("author_id").Comment("作者账号 ID").Immutable(),
		field.String("city").Comment("发布时 IP 解析的城市快照").Optional().Nillable().MaxLen(128),
	}
}

func (Moonbreeze) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (Moonbreeze) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at", "id"),
		index.Fields("author_id", "created_at", "id"),
	}
}
