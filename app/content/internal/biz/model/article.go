package model

import (
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"content/internal/data/ent/gen"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Article gen.Article

// Summary 文章摘要
func (a *Article) Summary() {
	r := []rune(a.Content)
	if len(r) > 20 {
		a.Content = string(r[:20]) + "..."
	}
}

func (a *Article) ConvertToRpc(authorUser *userv1.User, lastReplyUser *userv1.User, repliedAt *time.Time) *v1.Article {
	article := &v1.Article{
		CreatedAt:               timestamppb.New(*a.CreatedAt),
		UpdatedAt:               timestamppb.New(*a.UpdatedAt),
		CreatedBy:               a.CreatedBy,
		UpdatedBy:               a.UpdatedBy,
		Id:                      a.ID,
		Title:                   a.Title,
		Content:                 a.Content,
		HasPostscript:           a.HasPostscript,
		RewardContent:           a.RewardContent,
		RewardPoints:            a.RewardPoints,
		Status:                  a.Status,
		Type:                    a.Type,
		Commentable:             a.Commentable,
		Anonymous:               a.Anonymous,
		ThankCount:              a.ThankCount,
		LikeCount:               a.LikeCount,
		CollectCount:            a.CollectCount,
		WatchCount:              a.WatchCount,
		ReplyCount:              a.ReplyCount,
		BountyPoints:            a.BountyPoints,
		AcceptedAnswerId:        a.AcceptedAnswerID,
		VoteTotal:               a.VoteTotal,
		LotteryParticipantCount: a.LotteryParticipantCount,
		LotteryWinnerCount:      a.LotteryWinnerCount,
		AuthorUser:              authorUser,
		LastReplyUser:           lastReplyUser,
	}
	if repliedAt != nil {
		article.LastReplyAt = timestamppb.New(*repliedAt)
	}
	entArticle := (*gen.Article)(a)
	if len(entArticle.Edges.Postscripts) > 0 {
		for _, postscript := range entArticle.Edges.Postscripts {
			article.Postscripts = append(article.Postscripts, (*ArticlePostscript)(postscript).ConvertToRpc())
		}
	}
	if len(entArticle.Edges.Tags) > 0 {
		for _, tag := range entArticle.Edges.Tags {
			article.Tags = append(article.Tags, (*Tag)(tag).ConvertToRpc())
		}
	}
	return article
}
