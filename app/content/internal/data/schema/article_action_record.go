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

// ArticleActionRecord 存储用户对文章的各种行为（点赞、收藏、感谢等）
type ArticleActionRecord struct {
	ent.Schema
}

func (ArticleActionRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixContent.String() + "article_action_records",
		},
		entsql.WithComments(true),
	}
}

func (ArticleActionRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

// Fields 定义表字段
func (ArticleActionRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("article_id").Comment("关联的文章ID"),
		field.Int64("user_id").Comment("执行行为的用户ID"),
		field.Enum("type").Values(contentenum.ArticleActionMap.EnumValues()...).Comment("行为类型"),
		field.String("client_request_id").Nillable().Optional().Comment("客户端幂等请求 ID"),
		field.Int64("points").Nillable().Optional().Comment("本次互动涉及的积分数量"),
		field.Enum("status").Values(contentenum.ContentActionStatusMap.EnumValues()...).Default(contentenum.ContentActionStatusPublished.String()).Comment("互动命令状态"),
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
		index.Fields("article_id", "user_id", "type").Unique().Annotations(entsql.IndexWhere("status IN ('pending', 'published')")),
		index.Fields("client_request_id").Unique(),
		index.Fields("user_id"),
	}
}
