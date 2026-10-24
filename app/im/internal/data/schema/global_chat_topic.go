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

// GlobalChatTopic 保存一次全局聊天室话题设置请求。
type GlobalChatTopic struct{ ent.Schema }

func (GlobalChatTopic) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixIM.String() + "global_chat_topics"}, entsql.WithComments(true)}
}

func (GlobalChatTopic) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("设置话题的账号 ID"),
		field.String("client_request_id").NotEmpty().Immutable().Unique().Comment("客户端幂等请求 ID"),
		field.String("content").NotEmpty().Comment("话题内容"),
		field.Int64("points_cost").Positive().Comment("本次设置消耗的积分"),
		field.Enum("status").Values(enum.GlobalChatTopicStatusMap.EnumValues()...).Comment("话题设置状态"),
	}
}

func (GlobalChatTopic) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (GlobalChatTopic) Indexes() []ent.Index {
	return []ent.Index{index.Fields("status", "id")}
}
