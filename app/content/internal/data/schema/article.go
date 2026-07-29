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

// Article 定义文章实体。
type Article struct {
	ent.Schema
}

func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixContent.String() + "articles",
		},
		entsql.WithComments(true),
	}
}

func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("title").Comment("标题").NotEmpty(),
		field.Text("content").Comment("正文内容").NotEmpty(),
		field.String("cover_image_url").Comment("封面图片 URL").Nillable().Optional(),

		field.Bool("has_postscript").Comment("是否有附言").Default(false),
		field.Text("reward_content").Comment("打赏区内容").Nillable().Optional(),
		field.Int32("reward_points").Comment("打赏积分").Nillable().Optional(),

		field.Enum("publish_status").Values(contentenum.ArticlePublishStatusMap.EnumValues()...).Default(contentenum.ArticlePublishStatusDraft.String()).Comment("发布状态"),
		field.Enum("visibility").Values(contentenum.ArticleVisibilityMap.EnumValues()...).Default(contentenum.ArticleVisibilityPublic.String()).Comment("可见范围"),
		field.Enum("restriction").Values(contentenum.ContentRestrictionMap.EnumValues()...).Default(contentenum.ContentRestrictionNone.String()).Comment("管理限制"),
		field.Enum("type").Values(contentenum.ArticleTypeMap.EnumValues()...).Default(contentenum.ArticleTypeNormal.String()).Comment("类型"),
		field.String("statement").Comment("创作声明").Nillable().Optional(),
		field.Bool("commentable").Comment("是否允许评论").Default(true),
		field.Bool("anonymous").Comment("是否匿名").Default(false),
		field.Time("published_at").Comment("发布时间").Nillable().Optional(),
		field.Time("edited_at").Comment("内容编辑时间").Nillable().Optional(),

		field.Int32("view_count").Comment("浏览数").Default(0),
		field.Int32("thank_count").Comment("感谢数").Default(0),
		field.Int32("like_count").Comment("点赞数").Default(0),
		field.Int32("collect_count").Comment("收藏数").Default(0),
		field.Int32("watch_count").Comment("关注数").Default(0),
		field.Int32("reply_count").Comment("回复数").Default(0),

		field.Int32("bounty_points").Comment("悬赏积分").Nillable().Optional(),
		field.Int64("accepted_answer_id").Comment("采纳评论 ID").Nillable().Optional(),
		field.Int64("created_by").Comment("创建人 ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人 ID").Nillable().Optional(),
	}
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("postscripts", ArticlePostscript.Type),
		edge.To("comments", Comment.Type),
		edge.To("tags", Tag.Type).
			StorageKey(edge.Table(constant.TablePrefixContent.String() + "article_tags")),
		edge.To("action_records", ArticleActionRecord.Type),
	}
}
