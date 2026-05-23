package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TFA 存储账号双因素认证设置。
type TFA struct {
	ent.Schema
}

func (TFA) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "tfa"},
		entsql.WithComments(true),
	}
}

func (TFA) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID").Unique(),
		field.Bool("enable").Comment("是否启用双因素认证").Default(false),
		field.Time("enable_time").Comment("双因素认证启用时间").Optional().Nillable(),
		field.String("secret").Comment("双因素认证密钥").Default("").Sensitive(),
	}
}

func (TFA) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Account.Type).Ref("tfa").Field("user_id").Required().Unique(),
	}
}
