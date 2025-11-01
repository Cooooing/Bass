package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ArticleActionRecord 存储用户对文章的各种行为（点赞、收藏、感谢等）
type ArticleActionRecord struct {
	ent.Schema
}

// Fields 定义表字段
func (ArticleActionRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("关联的文章ID"),
		field.Int64("user_id").Comment("执行行为的用户ID"),
		field.Int32("type").Comment("行为类型 0-点赞 1-感谢 2-收藏 3-关注"),
	}
}

// Edges 边
func (ArticleActionRecord) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联文章 多对一
		edge.From("article", Article.Type).Ref("action_records").Required().Unique().Field("article_id"),
	}
}

// Indexes 定义索引
func (ArticleActionRecord) Indexes() []ent.Index {
	return []ent.Index{
		// 一个用户对一篇文章的某种行为只能有一条记录
		index.Fields("article_id", "user_id", "type").Unique(),
		// 常用查询索引
		index.Fields("article_id"),
		index.Fields("user_id"),
	}
}
