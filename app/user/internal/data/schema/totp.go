package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Totp 存储账号 TOTP 认证设置。
type Totp struct {
	ent.Schema
}

func (Totp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixUser.String() + "totp",
		},
		entsql.WithComments(true),
	}
}

func (Totp) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID").Unique(),
		field.Bool("enable").Comment("是否启用 TOTP 认证").Default(false),
		field.Time("enable_time").Comment("TOTP 认证启用时间").Optional().Nillable(),
		field.String("secret").Comment("TOTP 认证密钥").Default("").Sensitive(),
	}
}

func (Totp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Account.Type).Ref("totp").Field("user_id").Required().Unique(),
	}
}
