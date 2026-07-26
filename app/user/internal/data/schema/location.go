package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Location 存储账号可选位置信息。
type Location struct {
	ent.Schema
}

func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixUser.String() + "locations",
		},
		entsql.WithComments(true),
	}
}

func (Location) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID").Unique(),
		field.String("country").Comment("国家").Optional().Nillable(),
		field.String("province").Comment("省份").Optional().Nillable(),
		field.String("city").Comment("城市").Optional().Nillable(),
	}
}

func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Account.Type).Ref("location").Field("user_id").Required().Unique(),
	}
}
