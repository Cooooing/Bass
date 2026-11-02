package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"encoding/json"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserBadge stores a user's owned badge and display preferences.
type UserBadge struct {
	ent.Schema
}

func (UserBadge) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixBadge.String() + "user_badges"},
		entsql.WithComments(true),
	}
}
func (UserBadge) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("用户服务账户 ID"),
		field.Int64("badge_id").Comment("徽章定义 ID"),
		field.JSON("variables", json.RawMessage{}).Comment("说明模板变量 JSON 对象"),
		field.Time("obtained_at").Comment("用户获得或刷新徽章的时间"),
		field.Time("expires_at").Optional().Nillable().Comment("到期时间"),
		field.Bool("displayed").Default(true).Comment("是否公开展示"),
		field.Int32("display_order").Default(0).Comment("展示排序"),
	}
}
func (UserBadge) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }
func (UserBadge) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "badge_id").
			Unique().
			StorageKey("badge_user_badges_user_badge_uq"),
		index.Fields("user_id", "displayed", "display_order").
			StorageKey("badge_user_badges_display_order_idx"),
	}
}
