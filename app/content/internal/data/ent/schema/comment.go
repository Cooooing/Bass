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
	return append([]ent.Field{
		field.Int("article_id").Comment("所属文章ID"),
		field.Int("user_id").Comment("用户ID"),
		field.Text("content").Comment("评论内容").NotEmpty(),
		field.Int("level").Comment("评论层级"),
		field.Int("parent_id").Comment("父级评论ID").Optional().Default(0),
		field.Int("status").Comment("状态 0-正常 1-隐藏").Default(0),

		field.Int("reply_count").Comment("回复数").Default(0),
		field.Int("like_count").Comment("点赞数").Default(0),
		field.Int("dislike_count").Comment("点踩数").Default(0),
		field.Int("collect_count").Comment("收藏数").Default(0),
	}, pkg.TimeAuditFields()...)
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联文章 多对一
		edge.From("article", Article.Type).Ref("comments").Field("article_id").Required().Unique(),
		// 关联父评论 多对一
		edge.From("parent", Comment.Type).Ref("replies").Field("parent_id").Unique(),
		// 关联子评论 一对多
		edge.To("replies", Comment.Type),
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("article_id", "parent_id", "status"),
	}
}
