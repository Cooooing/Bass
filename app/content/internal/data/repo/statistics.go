package repo

import (
	"context"
	"time"

	"content/internal/biz/model"
	bizrepo "content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/article"
	"content/internal/data/gen/articleactionrecord"
	"content/internal/data/gen/articleviewrecord"
	"content/internal/data/gen/comment"
	"content/internal/data/gen/commentactionrecord"
	"content/internal/enum"
)

var _ bizrepo.StatisticsRepo = (*StatisticsRepo)(nil)

// StatisticsRepo 以分组查询批量统计内容领域事实表。
type StatisticsRepo struct{ db *gen.Client }

// NewStatisticsRepo 创建内容统计仓储。
func NewStatisticsRepo(db *gen.Client) bizrepo.StatisticsRepo { return &StatisticsRepo{db: db} }

// BatchCountUserBehaviors 以用户维度聚合所有内容行为，避免为每个账号重复查询。
func (r *StatisticsRepo) BatchCountUserBehaviors(ctx context.Context, userIDs []int64, start, end time.Time, actionStatuses []enum.ContentActionStatus) ([]*model.UserBehaviorCount, error) {
	counts := make(map[int64]*model.UserBehaviorCount, len(userIDs))
	for _, userID := range userIDs {
		counts[userID] = &model.UserBehaviorCount{UserID: userID}
	}
	type row struct {
		UserID int64  `json:"user_id"`
		Count  int    `json:"count"`
		Type   string `json:"type"`
	}
	type creatorRow struct {
		UserID int64 `json:"created_by"`
		Count  int   `json:"count"`
	}
	var rows []row
	var creators []creatorRow
	if err := r.db.Article.Query().Where(article.CreatedByIn(userIDs...), article.CreatedAtGTE(start), article.CreatedAtLT(end), article.DeletedAtIsNil(), article.PublishStatusEQ(article.PublishStatus(enum.ArticlePublishStatusPublished))).GroupBy(article.FieldCreatedBy).Aggregate(gen.Count()).Scan(ctx, &creators); err != nil {
		return nil, err
	}
	for _, item := range creators {
		counts[item.UserID].ArticleCount = int32(item.Count)
	}
	creators = nil
	if err := r.db.Comment.Query().Where(comment.CreatedByIn(userIDs...), comment.CreatedAtGTE(start), comment.CreatedAtLT(end), comment.DeletedAtIsNil()).GroupBy(comment.FieldCreatedBy).Aggregate(gen.Count()).Scan(ctx, &creators); err != nil {
		return nil, err
	}
	for _, item := range creators {
		counts[item.UserID].CommentCount = int32(item.Count)
	}
	rows = nil
	articleStatuses := make([]articleactionrecord.Status, 0, len(actionStatuses))
	commentStatuses := make([]commentactionrecord.Status, 0, len(actionStatuses))
	for _, status := range actionStatuses {
		articleStatuses = append(articleStatuses, articleactionrecord.Status(status))
		commentStatuses = append(commentStatuses, commentactionrecord.Status(status))
	}
	if err := r.db.ArticleActionRecord.Query().Where(articleactionrecord.UserIDIn(userIDs...), articleactionrecord.StatusIn(articleStatuses...), articleactionrecord.CreatedAtGTE(start), articleactionrecord.CreatedAtLT(end)).GroupBy(articleactionrecord.FieldUserID, articleactionrecord.FieldType).Aggregate(gen.Count()).Scan(ctx, &rows); err != nil {
		return nil, err
	}
	for _, item := range rows {
		switch enum.ArticleAction(item.Type) {
		case enum.ArticleActionLike:
			counts[item.UserID].ArticleLikeCount = int32(item.Count)
		case enum.ArticleActionThank:
			counts[item.UserID].ArticleThankCount = int32(item.Count)
		case enum.ArticleActionReward:
			counts[item.UserID].ArticleRewardCount = int32(item.Count)
		}
	}
	rows = nil
	if err := r.db.CommentActionRecord.Query().Where(commentactionrecord.UserIDIn(userIDs...), commentactionrecord.StatusIn(commentStatuses...), commentactionrecord.CreatedAtGTE(start), commentactionrecord.CreatedAtLT(end)).GroupBy(commentactionrecord.FieldUserID, commentactionrecord.FieldType).Aggregate(gen.Count()).Scan(ctx, &rows); err != nil {
		return nil, err
	}
	for _, item := range rows {
		switch enum.CommentAction(item.Type) {
		case enum.CommentActionLike:
			counts[item.UserID].CommentLikeCount = int32(item.Count)
		case enum.CommentActionThank:
			counts[item.UserID].CommentThankCount = int32(item.Count)
		}
	}
	rows = nil
	if err := r.db.ArticleViewRecord.Query().Where(articleviewrecord.UserIDIn(userIDs...), articleviewrecord.ViewedAtGTE(start), articleviewrecord.ViewedAtLT(end)).GroupBy(articleviewrecord.FieldUserID).Aggregate(gen.Count()).Scan(ctx, &rows); err != nil {
		return nil, err
	}
	for _, item := range rows {
		counts[item.UserID].ViewCount = int32(item.Count)
	}
	result := make([]*model.UserBehaviorCount, 0, len(userIDs))
	for _, userID := range userIDs {
		result = append(result, counts[userID])
	}
	return result, nil
}
