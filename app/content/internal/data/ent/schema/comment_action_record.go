package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CommentActionRecord 存储用户对文章的各种行为（点赞、收藏、感谢等）
type CommentActionRecord struct {
	ent.Schema
}

// Fields 定义表字段
func (CommentActionRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("comment_id").Comment("关联的文章ID"),
		field.Int("user_id").Comment("执行行为的用户ID"),
		field.Int("type").Comment("行为类型 0-点赞 1收藏"),
	}
}

// Edges 边
func (CommentActionRecord) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联评论 多对一
		edge.From("comment", Comment.Type).Ref("action_records").Required().Unique().Field("comment_id"),
	}
}

// Indexes 定义索引
func (CommentActionRecord) Indexes() []ent.Index {
	return []ent.Index{
		// 一个用户对一条评论的某种行为只能有一条记录
		index.Fields("comment_id", "user_id", "type").Unique(),
		// 常用查询索引
		index.Fields("comment_id"),
		index.Fields("user_id"),
	}
}
