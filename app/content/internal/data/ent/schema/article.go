package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article 文章实体定义
type Article struct {
	ent.Schema
}

func (Article) Fields() []ent.Field {
	return append([]ent.Field{
		field.String("user_id").Comment("作者ID").NotEmpty(),
		field.String("title").Comment("标题").NotEmpty(),
		field.Text("content").Comment("正文内容").NotEmpty(),

		field.Bool("has_postscript").Comment("是否有附言").Default(false),
		field.Text("reward_content").Comment("打赏区内容").Optional(),
		field.Int("reward_points").Comment("打赏积分").Default(0),

		field.Int("status").Comment("状态 0-正常 1-隐藏 2-锁定 3-草稿 4-删除").Default(0),
		field.Int("type").Comment("类型 0-普通 1-问答 2-投票 3-抽奖").Default(0),
		field.Bool("commentable").Comment("是否允许评论").Default(true),
		field.Bool("anonymous").Comment("是否允许评论").Default(true),

		// 统计信息
		field.Int("thank_count").Comment("帖子感谢数").Default(0),
		field.Int("like_count").Comment("点赞数").Default(0),
		field.Int("dislike_count").Comment("点踩数").Default(0),
		field.Int("collect_count").Comment("收藏数").Default(0),
		field.Int("watch_count").Comment("关注数").Default(0),

		// 问答
		field.Int("bounty_points").Comment("悬赏积分").Default(0),
		field.Int("accepted_answer_id").Comment("采纳评论ID").Optional(),

		// 投票 / 抽奖统计字段
		field.Int("vote_total").Comment("总投票数").Default(0),
		field.Int("lottery_participant_count").Comment("抽奖参与人数").Default(0),
		field.Int("lottery_winner_count").Comment("抽奖获奖人数").Default(0),
	}, pkg.TimeAuditFields()...)
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
		edge.To("tags", Tag.Type),
	}
}
