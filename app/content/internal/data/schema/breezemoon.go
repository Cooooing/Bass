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

// Breezemoon 定义明月清风短动态实体。
type Breezemoon struct {
	ent.Schema
}

func (Breezemoon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "breezemoons"},
		entsql.WithComments(true),
	}
}

func (Breezemoon) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Text("content").Comment("动态正文").NotEmpty(),
		field.Int64("author_id").Comment("作者账号 ID").Immutable(),
		field.String("city").Comment("发布时 IP 解析的城市快照").Optional().Nillable().MaxLen(128),
	}
}

func (Breezemoon) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (Breezemoon) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at", "id"),
		index.Fields("author_id", "created_at", "id"),
	}
}
