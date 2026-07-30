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
	"entgo.io/ent/schema/index"
)

type Domain struct {
	ent.Schema
}

func (Domain) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixContent.String() + "domains",
		},
		entsql.WithComments(true),
	}
}

func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("code").Comment("稳定编码").NotEmpty(),
		field.String("name").Comment("领域名称").NotEmpty(),
		field.String("description").Comment("领域描述").Nillable().Optional(),
		field.Enum("status").Values(contentenum.DomainStatusMap.EnumValues()...).Default(contentenum.DomainStatusEnabled.String()).Comment("启停状态"),
		field.String("url").Comment("领域地址").Nillable().Optional(),
		field.String("icon").Comment("图标").Nillable().Optional(),
		field.Bool("is_nav").Comment("是否导航展示").Default(false).Optional(),
		field.Int32("sort").Comment("排序值").Default(0),
		field.Int64("created_by").Comment("创建人 ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人 ID").Nillable().Optional(),
	}
}

func (Domain) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique().Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("name").Unique().Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("status", "is_nav", "sort").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}

func (Domain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
	}
}
