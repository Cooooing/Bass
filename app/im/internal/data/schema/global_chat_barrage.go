package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"im/internal/enum"
)

type GlobalChatBarrage struct{ ent.Schema }

func (GlobalChatBarrage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixIM.String() + "global_chat_barrages"},
		entsql.WithComments(true),
	}
}

func (GlobalChatBarrage) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("发送账号 ID"),
		field.String("client_request_id").NotEmpty().Immutable().Unique().Comment("客户端幂等请求 ID"),
		field.String("content").NotEmpty().Comment("纯文本弹幕内容"),
		field.Int64("points_cost").Positive().Comment("本次消耗积分"),
		field.Enum("status").Values(enum.GlobalChatBarrageStatusMap.EnumValues()...).Comment("弹幕状态"),
	}
}

func (GlobalChatBarrage) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (GlobalChatBarrage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status", "id"),
		index.Fields("user_id", "id"),
	}
}
