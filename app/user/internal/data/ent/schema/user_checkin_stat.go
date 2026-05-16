package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserCheckinStat 用户签到聚合统计
type UserCheckinStat struct {
	ent.Schema
}

func (UserCheckinStat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "user_checkin_stats"},
		entsql.WithComments(true),
	}
}

func (UserCheckinStat) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("用户ID").Unique(),
		field.Int32("total_online_minutes").Comment("累计在线总时长（分钟）").Default(0).Nillable(),
		field.Int32("current_streak").Comment("当前连续签到天数").Default(0).Nillable(),
		field.Int32("longest_streak").Comment("最长连续签到天数").Default(0).Nillable(),
	}
}

func (UserCheckinStat) Indexes() []ent.Index {
	return nil
}

func (UserCheckinStat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("checkin_stat").Field("user_id").Required().Unique(),
	}
}
