package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserLocation 用户地理信息
type UserLocation struct {
	ent.Schema
}

func (UserLocation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "user_locations"},
		entsql.WithComments(true),
	}
}

func (UserLocation) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("用户ID").Unique(),
		field.String("country").Comment("所在国家").Optional().Nillable(),
		field.String("province").Comment("所在省份").Optional().Nillable(),
		field.String("city").Comment("所在城市").Optional().Nillable(),
	}
}

func (UserLocation) Indexes() []ent.Index {
	return nil
}

func (UserLocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("location").Field("user_id").Required().Unique(),
	}
}
