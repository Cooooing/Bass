package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	userenum "user/internal/enum"

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

func (CheckinRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (CheckinRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID"),
		field.Time("date").Comment("签到所属业务日"),
		field.Enum("status").Values(userenum.CheckinStatusMap.EnumValues()...).Default(userenum.CheckinStatusPending.String()).Comment("自动签到事务状态"),
		field.Int64("checkin_points").Comment("自动签到奖励积分快照").Default(0),
		field.Int64("streak_points").Comment("连续签到奖励积分快照").Default(0),
		field.Time("checked_at").Comment("自动签到确认时间").Optional().Nillable(),
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
