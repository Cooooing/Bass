package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"im/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type GlobalChatMessage struct{ ent.Schema }

func (GlobalChatMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixIM.String() + "global_messages"}, entsql.WithComments(true)}
}

func (GlobalChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("sender_id").Comment("发送账号 ID"),
		field.String("client_message_id").Comment("客户端幂等消息 ID").NotEmpty(),
		field.Text("content").Comment("原始 Markdown 内容"),
		field.Enum("message_type").Values(enum.MessageTypeMap.EnumValues()...).Comment("消息类型").Default(enum.MessageTypeText.String()),
		field.Bool("is_revoked").Comment("是否已撤回").Default(false),
		field.Time("revoked_at").Comment("撤回时间").Optional().Nillable(),
		field.Int64("revoked_by").Comment("撤回操作账号 ID").Optional().Nillable(),
	}
}

func (GlobalChatMessage) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (GlobalChatMessage) Indexes() []ent.Index {
	return []ent.Index{index.Fields("sender_id", "client_message_id").Unique()}
}
