package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"im/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ChatGroup struct {
	ent.Schema
}

func (ChatGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixIM.String() + "chat_groups",
		},
	}
}

func (ChatGroup) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("群名称").NotEmpty(),
		field.String("avatar").Comment("群头像").Optional().Nillable(),
		field.String("introduction").Comment("群简介").Optional().Nillable(),
		field.Int64("owner_id").Comment("群主id"),
		field.Enum("status").Values(enum.ChatGroupStatusMap.EnumValues()...).Default(string(enum.ChatGroupStatusNormal)).Comment("群状态"),
		field.Uint32("member_count").Comment("群成员数").Default(0),

		field.Uint32("message_count").Comment("群消息数").Default(0),
		field.Int64("last_message_id").Comment("最后一条消息id").Unique().Optional().Nillable(),
		field.Int64("created_by").Comment("创建人ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人ID").Nillable().Optional(),
	}
	return fields
}

func (ChatGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (ChatGroup) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联群成员 一对多
		edge.To("members", ChatGroupMember.Type),
		// 关联会话 一对多
		edge.To("group_sessions", ChatSession.Type),
		// 关联消息 一对多
		edge.To("group_messages", ChatMessage.Type),
		// 关联最后一条群消息 一对一
		edge.From("last_message_of_group", ChatMessage.Type).Ref("last_message_group").Unique().Field("last_message_id"),
	}
}

func (ChatGroup) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
	}
}
