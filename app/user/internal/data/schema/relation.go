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

// Relation 存储账号之间的有向关系，如关注和拉黑。
type Relation struct {
	ent.Schema
}

func (Relation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "relations"},
		entsql.WithComments(true),
	}
}

func (Relation) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("actor_id").Comment("关系发起账号 ID"),
		field.Int64("target_id").Comment("关系目标账号 ID"),
		field.Enum("type").Values(userenum.RelationTypeMap.EnumValues()...).Comment("关系类型"),
	}
}

func (Relation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (Relation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("actor_id", "target_id", "type").Unique(),
		index.Fields("actor_id", "type"),
		index.Fields("target_id", "type"),
	}
}

func (Relation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("actor", Account.Type).Ref("relations_as_actor").Field("actor_id").Required().Unique(),
		edge.From("target", Account.Type).Ref("relations_as_target").Field("target_id").Required().Unique(),
	}
}
