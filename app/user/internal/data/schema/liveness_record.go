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

// LivenessRecord 存储账号某业务日冻结后的活跃度及手动奖励事实。
type LivenessRecord struct {
	ent.Schema
}

func (LivenessRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "liveness_records"},
		entsql.WithComments(true),
	}
}

func (LivenessRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (LivenessRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID"),
		field.Time("date").Comment("活跃度所属业务日"),
		field.Int32("score").Comment("冻结后的活跃度总分"),
		field.Enum("reward_status").Values(userenum.LivenessRewardStatusMap.EnumValues()...).Default(userenum.LivenessRewardStatusUnclaimed.String()).Comment("活跃奖励领取状态"),
		field.Time("claimed_at").Comment("活跃奖励领取确认时间").Optional().Nillable(),
	}
}

func (LivenessRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "date").Unique(),
		index.Fields("date", "reward_status"),
	}
}

func (LivenessRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Account.Type).Ref("liveness_records").Field("user_id").Required().Unique(),
	}
}
