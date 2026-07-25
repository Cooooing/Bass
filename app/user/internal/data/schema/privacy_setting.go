package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PrivacySetting 存储账号可选隐私设置。
type PrivacySetting struct {
	ent.Schema
}

func (PrivacySetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixUser.String() + "privacies",
		},
		entsql.WithComments(true),
	}
}

func (PrivacySetting) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID").Unique(),
		field.Bool("public_points").Comment("积分是否公开").Default(true),
		field.Bool("public_followers").Comment("粉丝列表是否公开").Default(true),
		field.Bool("public_articles").Comment("文章列表是否公开").Default(true),
		field.Bool("public_comments").Comment("评论列表是否公开").Default(true),
		field.Bool("public_online_status").Comment("在线状态是否公开").Default(true),
		field.Bool("public_location").Comment("位置是否公开").Default(true),
	}
}

func (PrivacySetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Account.Type).Ref("privacy_settings").Field("user_id").Required().Unique(),
	}
}
