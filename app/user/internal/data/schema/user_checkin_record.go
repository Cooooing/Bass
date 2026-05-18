package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserCheckinRecord 用户每日签到记录
type UserCheckinRecord struct {
	ent.Schema
}

func (UserCheckinRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "user_checkin_records"},
		entsql.WithComments(true),
	}
}

func (UserCheckinRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("用户ID"),
		field.Time("date").Comment("签到日期").Nillable(),
		field.Int32("online_minutes").Comment("当日在线时长（分钟）").Default(0).Nillable(),
		field.Int32("activity").Comment("当日活跃度").Default(0).Nillable(),
		field.Bool("checked").Comment("是否达标签到").Default(false),
	}
}

func (UserCheckinRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "date").Unique(),
		index.Fields("user_id"),
	}
}

func (UserCheckinRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("checkin_records").Field("user_id").Required().Unique(),
	}
}
