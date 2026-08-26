package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	gameenum "game_idle/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ChatMessage 定义挂机游戏聊天消息。
type ChatMessage struct {
	ent.Schema
}

func (ChatMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "chat_messages",
		},
		entsql.WithComments(true),
	}
}

func (ChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Enum("channel_type").
			Values(gameenum.ChatChannelTypeValues()...).
			Default(gameenum.ChatChannelTypeWorld.String()).
			Comment("频道类型"),
		field.String("channel_id").Comment("频道标识").MaxLen(128).NotEmpty(),
		field.Int64("sender_character_id").Comment("发送角色 ID").Positive(),
		field.Int64("receiver_character_id").Comment("接收角色 ID").Positive().Optional().Nillable(),
		field.String("content").Comment("消息内容").MaxRuneLen(500).NotEmpty(),
		field.Enum("status").
			Values(gameenum.ChatMessageStatusValues()...).
			Default(gameenum.ChatMessageStatusNormal.String()).
			Comment("消息状态"),
	}
}

func (ChatMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (ChatMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("channel_type", "channel_id", "id").
			StorageKey("game_idle_chat_messages_channel_id_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("sender_character_id", "id").
			StorageKey("game_idle_chat_messages_sender_id_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("status", "id").
			StorageKey("game_idle_chat_messages_status_id_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}
