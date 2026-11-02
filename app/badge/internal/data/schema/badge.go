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

// Badge stores a reusable badge owned by the badge service.
type Badge struct {
	ent.Schema
}

func (Badge) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixBadge.String() + "badges"},
		entsql.WithComments(true),
	}
}
func (Badge) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("code").Optional().Nillable().Unique().Comment("发布后的唯一稳定徽章编码"),
		field.String("name").Comment("徽章名称"),
		field.String("description").Default("").Comment("带变量占位符的徽章说明"),
		field.JSON("render_params", json.RawMessage{}).Comment("符合 RenderSpec 的徽章渲染参数 JSON"),
		field.Int64("designer_user_id").Comment("设计该徽章的用户 ID"),
		field.Enum("status").Values("draft", "published").Default("draft").Comment("徽章定义发布状态"),
		field.Time("published_at").Optional().Nillable().Comment("定义发布为主站徽章的时间"),
	}
}
func (Badge) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }

func (Badge) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("designer_user_id", "status", "updated_at").
			StorageKey("badge_badges_designer_status_updated_idx"),
		index.Fields("status", "code").
			StorageKey("badge_badges_status_code_idx"),
	}
}
