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

type Tag struct {
	ent.Schema
}

func (Tag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixContent.String() + "tags",
		},
		entsql.WithComments(true),
	}
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("code").Comment("稳定编码").NotEmpty(),
		field.String("name").Comment("标签名称").NotEmpty(),
		field.String("description").Comment("标签描述").Optional().Nillable(),
		field.Int64("domain_id").Comment("所属领域 ID").Optional().Nillable(),
		field.String("icon").Comment("图标").Optional().Nillable(),
		field.Int32("sort").Comment("排序值").Default(0),
		field.Int32("article_count").Comment("文章数量").Default(0),
		field.Enum("status").Values(contentenum.TagStatusMap.EnumValues()...).Default(contentenum.TagStatusEnabled.String()).Comment("标签启停状态"),
		field.Int64("created_by").Comment("创建人 ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人 ID").Nillable().Optional(),
	}
}

func (Tag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("domain_id", "code").Unique().Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("domain_id", "name").Unique().Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("domain_id", "status", "sort").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}

func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).Ref("tags"),
		edge.From("domain", Domain.Type).Ref("tags").Field("domain_id").Unique(),
	}
}
