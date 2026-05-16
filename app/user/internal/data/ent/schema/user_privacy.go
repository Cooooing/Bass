package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserPrivacy 用户隐私设置
type UserPrivacy struct {
	ent.Schema
}

func (UserPrivacy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "user_privacies"},
		entsql.WithComments(true),
	}
}

func (UserPrivacy) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("用户ID").Unique(),
		field.Bool("public_points").Comment("是否公开积分榜").Default(true).Nillable(),
		field.Bool("public_followers").Comment("是否公开粉丝列表").Default(true).Nillable(),
		field.Bool("public_articles").Comment("是否公开帖子列表").Default(true).Nillable(),
		field.Bool("public_comments").Comment("是否公开评论列表").Default(true).Nillable(),
		field.Bool("public_online_status").Comment("是否公开在线状态").Default(true).Nillable(),
		field.Bool("public_location").Comment("是否公开地理位置").Default(true).Nillable(),
	}
}

func (UserPrivacy) Indexes() []ent.Index {
	return nil
}

func (UserPrivacy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("privacy").Field("user_id").Required().Unique(),
	}
}
