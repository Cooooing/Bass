package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CheckinStat 存储账号签到聚合统计。
type CheckinStat struct {
	ent.Schema
}

func (CheckinStat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "checkin_stats"},
		entsql.WithComments(true),
	}
}

func (CheckinStat) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID").Unique(),
		field.Int32("total_online_minutes").Comment("累计在线时长，单位为分钟").Default(0).Nillable(),
		field.Int32("current_streak").Comment("当前连续签到天数").Default(0).Nillable(),
		field.Int32("longest_streak").Comment("最长连续签到天数").Default(0).Nillable(),
	}
}

func (CheckinStat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Account.Type).Ref("checkin_stat").Field("user_id").Required().Unique(),
	}
}
