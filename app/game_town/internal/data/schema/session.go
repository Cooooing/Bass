package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Session 保存文字客户端会话。
type Session struct{ ent.Schema }

func (Session) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "sessions"}, entsql.WithComments(true)}
}
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("player_id").Comment("玩家 ID").Nillable().Optional(),
		field.Int64("current_world_id").Comment("当前世界 ID").Nillable().Optional(),
		field.String("client_type").Comment("客户端类型").MaxLen(32).NotEmpty(),
		field.Time("started_at").Comment("开始时间"),
		field.Time("last_seen_at").Comment("最近活跃时间"),
		field.Time("ended_at").Comment("结束时间").Nillable().Optional(),
	}
}
func (Session) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }
func (Session) Edges() []ent.Edge  { return nil }
