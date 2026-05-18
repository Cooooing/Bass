package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	contentenum "content/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Domain struct {
	ent.Schema
}

func (Domain) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "domains"},
		entsql.WithComments(true),
	}
}

func (Domain) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("域名名称").NotEmpty(),
		field.String("description").Comment("域名描述").Nillable().Optional(),
		field.Enum("status").Values(contentenum.DomainStatusMap.EnumValues()...).Default(string(contentenum.DomainStatusNormal)).Comment("状态"),
		field.String("url").Comment("领域地址").Nillable().Optional(),
		field.String("icon").Comment("图标").Nillable().Optional(),
		field.Int32("tag_count").Comment("标签数").Default(0),
		field.Bool("is_nav").Comment("是否导航").Default(false).Optional(),
	}
	return fields
}

func (Domain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.UserAuditMixin{},
	}
}

func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联标签 一对多
		edge.To("tags", Tag.Type),
	}
}

func (Domain) Indexes() []ent.Index {
	return []ent.Index{}
}
