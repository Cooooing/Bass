package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserTFA 用户二步验证
type UserTFA struct {
	ent.Schema
}

func (UserTFA) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "user_tfa"},
		entsql.WithComments(true),
	}
}

func (UserTFA) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("用户ID").Unique(),
		field.Bool("enable").Comment("是否开启二步验证").Default(false),
		field.Time("enable_time").Comment("二步验证启用时间").Optional().Nillable(),
		field.String("secret").Comment("二步验证Secret").Default("").Sensitive(),
	}
}

func (UserTFA) Indexes() []ent.Index {
	return nil
}

func (UserTFA) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("tfa").Field("user_id").Required().Unique(),
	}
}
