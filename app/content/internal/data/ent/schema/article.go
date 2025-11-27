package schema

import (
	"common/pkg"
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article 文章实体定义
type Article struct {
	ent.Schema
}

func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "articles"},
	}
}

func (Article) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("title").Comment("标题").NotEmpty(),
		field.Text("content").Comment("正文内容").NotEmpty(),

		field.Bool("has_postscript").Comment("是否有附言").Default(false),
		field.Text("reward_content").Comment("打赏区内容").Nillable().Optional(),
		field.Int32("reward_points").Comment("打赏积分").Default(0).Nillable().Optional(),

		field.Int32("status").Comment("状态 0-正常 1-隐藏 2-锁定 3-草稿 4-删除").Default(0),
		field.Int32("type").Comment("类型 0-普通 1-问答 2-投票 3-抽奖").Default(0),
		field.String("statement").Comment("创作声明").Nillable().Optional(),
		field.Bool("commentable").Comment("是否允许评论").Default(true),
		field.Bool("anonymous").Comment("是否匿名").Default(false),
		field.Bool("listable").Comment("是否在列表展示").Default(true),

		// 统计信息
		field.Int32("thank_count").Comment("感谢数").Default(0),
		field.Int32("like_count").Comment("点赞数").Default(0),
		field.Int32("collect_count").Comment("收藏数").Default(0),
		field.Int32("watch_count").Comment("关注数").Default(0),
		field.Int32("reply_count").Comment("回复数").Default(0),

		// 问答
		field.Int32("bounty_points").Comment("悬赏积分").Default(0).Nillable().Optional(),
		field.Int64("accepted_answer_id").Comment("采纳评论ID").Nillable().Optional(),

		// 投票 / 抽奖统计字段
		field.Int32("vote_total").Comment("总投票数").Default(0),
		field.Int32("lottery_participant_count").Comment("抽奖参与人数").Default(0),
		field.Int32("lottery_winner_count").Comment("抽奖获奖人数").Default(0),
	}
	fields = append(fields, pkg.UserAuditFields()...)
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
}

func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联附言 一对多
		edge.To("postscripts", ArticlePostscript.Type),
		// 关联投票 一对多
		edge.To("votes", ArticleVote.Type),
		// 关联抽奖 一对多
		edge.To("lotteries", ArticleLottery.Type),
		// 关联评论 一对多
		edge.To("comments", Comment.Type),
		// 关联标签 多对多
		edge.To("tags", Tag.Type).
			StorageKey(
				edge.Table(constant.TablePrefixContent.String() + "article_tags"), // 自定义中间表名
			),
		// 关联操作 一对多
		edge.To("action_records", ArticleActionRecord.Type),
	}
}
