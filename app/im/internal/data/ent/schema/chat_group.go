package schema

import (
	"common/pkg"
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ChatGroup struct {
	ent.Schema
}

func (ChatGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixIM.String() + "chat_groups"},
	}
}

func (ChatGroup) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("群名称").NotEmpty(),
		field.String("avatar").Comment("群头像").Default(""),
		field.Int64("owner_id").Comment("群主id"),
		field.Int32("status").Comment("群状态: 0-正常, 1-解散").Default(0),
		field.Int32("member_count").Comment("群成员数").Default(0),
		field.Int64("message_count").Comment("群消息数").Default(0),
		field.Int64("last_message_id").Comment("最后一条消息id").Default(0),
		field.String("last_message_content").Comment("消息内容").Default(""),
	}
	fields = append(fields, pkg.UserAuditFields()...)
	fields = append(fields, pkg.TimeAuditFields()...)
	fields = append(fields, pkg.UsernameAuditFields()...)
	return fields
}

func (ChatGroup) Edges() []ent.Edge {
	return []ent.Edge{
		// 成员列表 (O2M)
		edge.To("members", ChatGroupMember.Type),
		// 属于该群的所有会话 (O2M)
		edge.To("sessions", ChatSession.Type),
		// 群聊消息记录 (O2M)
		edge.To("messages", ChatMessage.Type),
		// 最后一条消息 (O2O)
		edge.To("last_message", ChatMessage.Type).Unique().Field("last_message_id"),
	}
}
