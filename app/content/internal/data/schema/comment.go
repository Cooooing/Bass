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

// Comment 定义评论实体。
type Comment struct {
	ent.Schema
}

func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixContent.String() + "comments",
		},
		entsql.WithComments(true),
	}
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("所属文章 ID"),
		field.Text("content").Comment("评论内容").NotEmpty(),
		field.Int32("level").Comment("评论层级"),
		field.Int64("parent_id").Comment("父级评论 ID").Optional().Nillable(),
		field.Int64("reply_id").Comment("回复评论 ID").Optional().Nillable(),
		field.Enum("restriction").Values(contentenum.ContentRestrictionMap.EnumValues()...).Default(string(contentenum.ContentRestrictionNone)).Comment("管理限制"),

		field.Int32("thank_count").Comment("感谢数").Default(0),
		field.Int32("like_count").Comment("点赞数").Default(0),
		field.Int32("reply_count").Comment("回复数").Default(0),
		field.Int64("created_by").Comment("创建人 ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人 ID").Nillable().Optional(),
	}
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).Ref("comments").Field("article_id").Required().Unique(),
		edge.From("parent", Comment.Type).Ref("parent_replies").Field("parent_id").Unique(),
		edge.To("parent_replies", Comment.Type),
		edge.From("reply", Comment.Type).Ref("reply_replies").Field("reply_id").Unique(),
		edge.To("reply_replies", Comment.Type),
		edge.To("action_records", CommentActionRecord.Type),
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("article_id", "level", "restriction", "created_at", "id").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("article_id", "parent_id", "restriction", "created_at", "id").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("article_id", "restriction", "created_at", "id").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}
