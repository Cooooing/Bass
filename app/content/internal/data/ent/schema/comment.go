package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Comment 评论实体定义
type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("所属文章ID"),
		field.Text("content").Comment("评论内容").NotEmpty(),
		field.Int32("level").Comment("评论层级"),
		field.Int64("parent_id").Comment("父级评论ID").Optional().Nillable(),
		field.Int64("reply_id").Comment("回复评论ID").Optional().Nillable(),
		field.Int32("status").Comment("状态 0-正常 1-隐藏").Default(0),

		field.Int32("reply_count").Comment("回复数").Default(0),
		field.Int32("like_count").Comment("点赞数").Default(0),
		field.Int32("collect_count").Comment("收藏数").Default(0),
	}
	fields = append(fields, pkg.UserAuditFields()...)
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联文章 多对一
		edge.From("article", Article.Type).Ref("comments").Field("article_id").Required().Unique(),
		// 关联父评论 多对一
		edge.From("parent", Comment.Type).Ref("parent_replies").Field("parent_id").Unique(),
		// 关联子评论 一对多
		edge.To("parent_replies", Comment.Type),
		// 关联回复评论 多对一
		edge.From("reply", Comment.Type).Ref("reply_replies").Field("reply_id").Unique(),
		// 关联子评论 一对多
		edge.To("reply_replies", Comment.Type),
		// 关联操作 一对多
		edge.To("action_records", CommentActionRecord.Type),
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("article_id", "parent_id", "status"),
	}
}
