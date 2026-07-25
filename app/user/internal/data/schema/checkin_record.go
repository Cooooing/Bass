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

// CheckinRecord 存储账号每日签到记录。
type CheckinRecord struct {
	ent.Schema
}

func (CheckinRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixUser.String() + "checkin_records",
		},
		entsql.WithComments(true),
	}
}

func (CheckinRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID"),
		field.Time("date").Comment("签到日期"),
		field.Int32("online_minutes").Comment("当天在线时长，单位为分钟").Default(0),
		field.Int32("activity").Comment("当天活跃度").Default(0),
		field.Bool("checked").Comment("是否完成签到").Default(false),
	}
}

func (CheckinRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "date").Unique(),
	}
}

func (CheckinRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Account.Type).Ref("checkin_records").Field("user_id").Required().Unique(),
	}
}
