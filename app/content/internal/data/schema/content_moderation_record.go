package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	contentenum "content/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ContentModerationRecord 内容管理处置记录。
type ContentModerationRecord struct {
	ent.Schema
}

func (ContentModerationRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixContent.String() + "moderation_records",
		},
		entsql.WithComments(true),
	}
}

func (ContentModerationRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Enum("target").Values(contentenum.ContentModerationTargetMap.EnumValues()...).Comment("处置目标"),
		field.Int64("target_id").Comment("处置目标 ID"),
		field.Enum("action").Values(contentenum.ContentModerationActionMap.EnumValues()...).Comment("处置动作"),
		field.String("reason_code").Comment("原因编码").Nillable().Optional(),
		field.Text("reason").Comment("原因说明").Nillable().Optional(),
		field.Int64("operator_id").Comment("操作人 ID"),
	}
}

func (ContentModerationRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (ContentModerationRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("target", "target_id", "created_at", "id"),
		index.Fields("operator_id", "created_at", "id"),
	}
}
