package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Account 定义用户积分账户
type Account struct {
	ent.Schema
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixEconomy.String() + "accounts"},
		entsql.WithComments(true),
	}
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("用户 ID"),
		field.Int64("balance").Default(0).NonNegative().Comment("当前积分"),
		field.Int64("total_income").Default(0).NonNegative().Comment("累计获得"),
		field.Int64("total_expense").Default(0).NonNegative().Comment("累计消耗"),
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").Unique().Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}
