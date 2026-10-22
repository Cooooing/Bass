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

type GlobalChatMute struct{ ent.Schema }

func (GlobalChatMute) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixIM.String() + "global_chat_mutes"}, entsql.WithComments(true)}
}

func (GlobalChatMute) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("被禁言账号 ID"),
		field.Time("muted_until").Comment("禁言截止时间"),
		field.String("reason").Comment("禁言原因").Optional().Nillable(),
		field.Int64("operator_user_id").Comment("操作管理员账号 ID"),
	}
}

func (GlobalChatMute) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (GlobalChatMute) Indexes() []ent.Index {
	return []ent.Index{index.Fields("user_id").Unique(), index.Fields("muted_until")}
}
