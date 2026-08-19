package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	economyenum "economy/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Record 定义积分流水
type Record struct {
	ent.Schema
}

func (Record) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixEconomy.String() + "records"},
		entsql.WithComments(true),
	}
}

func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("transaction_no").Comment("交易号"),
		field.Int64("user_id").Comment("用户 ID"),
		field.Enum("record_type").Values(economyenum.EconomyRecordTypeMap.EnumValues()...).Comment("流水类型"),
		field.Enum("direction").Values(economyenum.EconomyRecordDirectionMap.EnumValues()...).Comment("流水方向"),
		field.Int64("amount").Positive().Comment("积分数量"),
		field.Int64("balance_before").NonNegative().Comment("变动前积分"),
		field.Int64("balance_after").NonNegative().Comment("变动后积分"),
		field.String("remark").Nillable().Optional().Comment("备注"),
	}
}

func (Record) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (Record) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("transaction_no"),
		index.Fields("user_id", "created_at", "id"),
		index.Fields("record_type", "user_id", "created_at", "id"),
	}
}
