package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/apperror"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	contentv1 "common/proto/gen/content/v1"
	userv1 "common/proto/gen/user/v1"
	"context"
	"fmt"

	"github.com/samber/lo"
)

var _ repo.ContentCommentClient = (*ContentCommentClient)(nil)

type ContentCommentClient struct {
	contentClient *rpc.ContentClient
	userClient    *rpc.UserClient
}

func NewContentCommentClient(contentClient *rpc.ContentClient, userClient *rpc.UserClient) repo.ContentCommentClient {
	return &ContentCommentClient{contentClient: contentClient, userClient: userClient}
}

func (r *ContentCommentClient) CreateComment(ctx context.Context, req *bbscontentv1.CreateComment_Request) (*bbscontentv1.CreateComment_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	var replyID *int64
	if req.GetReplyId() != 0 {
		replyID = new(req.GetReplyId())
	}
	reply, err := r.contentClient.Comment.Create(ctx, &contentv1.CreateComment_Request{ArticleId: req.GetArticleId(), Content: req.GetContent(), ReplyId: replyID, UserId: userID})
	if err != nil {
		return nil, err
	}
	item := reply.GetComment()
	articles, err := r.loadCommentArticles(ctx, []*contentv1.Comment{item}, userID)
	if err != nil {
		return nil, err
	}
	profiles, err := r.loadAccountProfiles(ctx, r.commentProfileIDs(item, articles[item.GetArticleId()])...)
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.CreateComment_Reply{Comment: r.commentDetail(item, articles[item.GetArticleId()], profiles)}, nil
}

func (r *ContentCommentClient) ListComments(ctx context.Context, req *bbscontentv1.ListComments_Request) (*bbscontentv1.ListComments_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	query := req.GetQuery()
	if query == nil {
		query = &bbscontentv1.CommentQuery{}
	}
	contentQuery := &contentv1.CommentQueryParams{
		CommentId: query.CommentId,
		ArticleId: query.ArticleId,
		ParentId:  query.ParentId,
		ReplyId:   query.ReplyId,
		Level:     query.Level,
		UserId:    query.UserId,
	}
	if query.Restriction != nil {
		contentQuery.Restriction = new(contentv1.ContentRestriction(*query.Restriction))
	}
	if len(query.Restrictions) > 0 {
		contentQuery.Restrictions = lo.Map(query.Restrictions, func(item bbscontentv1.ContentRestriction, _ int) contentv1.ContentRestriction {
			return contentv1.ContentRestriction(item)
		})
	}
	if query.Order != nil {
		contentQuery.Order = new(contentv1.CommentOrder(*query.Order))
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Request{
		Page:  req.Page,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	articles, err := r.loadCommentArticles(ctx, reply.GetRows(), userID)
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, 0, len(reply.GetRows())*2)
	for _, item := range reply.GetRows() {
		userIDs = append(userIDs, r.commentProfileIDs(item, articles[item.GetArticleId()])...)
	}
	profiles, err := r.loadAccountProfiles(ctx, userIDs...)
	if err != nil {
		return nil, err
	}
	states, err := r.loadCommentViewerActionStates(ctx, lo.Map(reply.GetRows(), func(item *contentv1.Comment, _ int) int64 {
		return item.GetId()
	}))
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.CommentListItem, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()]))
	}
	return &bbscontentv1.ListComments_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *ContentCommentClient) ListCommentThreads(ctx context.Context, req *bbscontentv1.ListCommentThreads_Request) (*bbscontentv1.ListCommentThreads_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	normal := contentv1.ContentRestriction_CONTENT_RESTRICTION_NONE
	locked := contentv1.ContentRestriction_CONTENT_RESTRICTION_LOCKED
	restrictions := []contentv1.ContentRestriction{normal, locked}
	order := contentv1.CommentOrder_COMMENT_ORDER_HOTTEST
	if req.Order != nil && *req.Order != bbscontentv1.CommentOrder_COMMENT_ORDER_UNSPECIFIED {
		order = contentv1.CommentOrder(*req.Order)
	}
	level := int32(1)
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Request{
		Page: req.Page,
		Query: &contentv1.CommentQueryParams{
			ArticleId:    new(req.ArticleId),
			Level:        &level,
			Restrictions: restrictions,
			Order:        &order,
		},
	})
	if err != nil {
		return nil, err
	}
	rootIDs := lo.Map(reply.GetRows(), func(item *contentv1.Comment, _ int) int64 {
		return item.GetId()
	})
	previewMap := map[int64][]*contentv1.Comment{}
	if len(rootIDs) > 0 {
		previewLimit := req.GetReplyPreviewLimit()
		if previewLimit <= 0 {
			previewLimit = 3
		}
		if previewLimit > 5 {
			previewLimit = 5
		}
		previewOrder := contentv1.CommentOrder_COMMENT_ORDER_OLDEST
		previews, err := r.contentClient.Comment.ListReplyPreviews(ctx, &contentv1.ListCommentReplyPreviews_Request{
			ArticleId:      req.GetArticleId(),
			ParentIds:      rootIDs,
			LimitPerParent: previewLimit,
			Order:          &previewOrder,
			Restrictions:   restrictions,
		})
		if err != nil {
			return nil, err
		}
		previewMap = lo.SliceToMap(previews.GetRows(), func(item *contentv1.CommentReplyPreview) (int64, []*contentv1.Comment) {
			return item.GetParentId(), item.GetRows()
		})
	}
	allComments := append([]*contentv1.Comment{}, reply.GetRows()...)
	for _, rows := range previewMap {
		allComments = append(allComments, rows...)
	}
	articles, err := r.loadCommentArticles(ctx, allComments, userID)
	if err != nil {
		return nil, err
	}
	userIDs := lo.FlatMap(reply.GetRows(), func(item *contentv1.Comment, _ int) []int64 {
		return r.commentProfileIDs(item, articles[item.GetArticleId()])
	})
	for _, rows := range previewMap {
		for _, item := range rows {
			userIDs = append(userIDs, r.commentProfileIDs(item, articles[item.GetArticleId()])...)
		}
	}
	profiles, err := r.loadAccountProfiles(ctx, userIDs...)
	if err != nil {
		return nil, err
	}
	commentIDs := append([]int64{}, rootIDs...)
	for _, rows := range previewMap {
		for _, item := range rows {
			commentIDs = append(commentIDs, item.GetId())
		}
	}
	states, err := r.loadCommentViewerActionStates(ctx, commentIDs)
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.CommentThread, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		previewRows := lo.Map(previewMap[item.GetId()], func(preview *contentv1.Comment, _ int) *bbscontentv1.CommentListItem {
			return r.commentListItem(preview, articles[preview.GetArticleId()], profiles, states[preview.GetId()])
		})
		rows = append(rows, &bbscontentv1.CommentThread{
			Root:           r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()]),
			PreviewReplies: previewRows,
			ReplyCount:     item.GetReplyCount(),
			HasMoreReplies: item.GetReplyCount() > int32(len(previewRows)),
		})
	}
	return &bbscontentv1.ListCommentThreads_Reply{
		Page: reply.GetPage(),
		Rows: rows,
	}, nil
}

func (r *ContentCommentClient) ListCommentReplies(ctx context.Context, req *bbscontentv1.ListCommentReplies_Request) (*bbscontentv1.ListCommentReplies_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	normal := contentv1.ContentRestriction_CONTENT_RESTRICTION_NONE
	locked := contentv1.ContentRestriction_CONTENT_RESTRICTION_LOCKED
	restrictions := []contentv1.ContentRestriction{normal, locked}
	order := contentv1.CommentOrder_COMMENT_ORDER_OLDEST
	if req.Order != nil && *req.Order != bbscontentv1.CommentOrder_COMMENT_ORDER_UNSPECIFIED {
		order = contentv1.CommentOrder(*req.Order)
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Request{
		Page: req.Page,
		Query: &contentv1.CommentQueryParams{
			ArticleId:    new(req.ArticleId),
			ParentId:     new(req.ParentId),
			Restrictions: restrictions,
			Order:        &order,
		},
	})
	if err != nil {
		return nil, err
	}
	articles, err := r.loadCommentArticles(ctx, reply.GetRows(), userID)
	if err != nil {
		return nil, err
	}
	userIDs := lo.FlatMap(reply.GetRows(), func(item *contentv1.Comment, _ int) []int64 {
		return r.commentProfileIDs(item, articles[item.GetArticleId()])
	})
	profiles, err := r.loadAccountProfiles(ctx, userIDs...)
	if err != nil {
		return nil, err
	}
	states, err := r.loadCommentViewerActionStates(ctx, lo.Map(reply.GetRows(), func(item *contentv1.Comment, _ int) int64 {
		return item.GetId()
	}))
	if err != nil {
		return nil, err
	}
	rows := lo.Map(reply.GetRows(), func(item *contentv1.Comment, _ int) *bbscontentv1.CommentListItem {
		return r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()])
	})
	return &bbscontentv1.ListCommentReplies_Reply{
		Page: reply.GetPage(),
		Rows: rows,
	}, nil
}

func (r *ContentCommentClient) ListCommentTimeline(ctx context.Context, req *bbscontentv1.ListCommentTimeline_Request) (*bbscontentv1.ListCommentTimeline_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	normal := contentv1.ContentRestriction_CONTENT_RESTRICTION_NONE
	locked := contentv1.ContentRestriction_CONTENT_RESTRICTION_LOCKED
	restrictions := []contentv1.ContentRestriction{normal, locked}
	order := contentv1.CommentOrder_COMMENT_ORDER_NEWEST
	if req.Order != nil && *req.Order != bbscontentv1.CommentOrder_COMMENT_ORDER_UNSPECIFIED {
		order = contentv1.CommentOrder(*req.Order)
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Request{
		Page: req.Page,
		Query: &contentv1.CommentQueryParams{
			ArticleId:    new(req.ArticleId),
			Restrictions: restrictions,
			Order:        &order,
		},
	})
	if err != nil {
		return nil, err
	}
	articles, err := r.loadCommentArticles(ctx, reply.GetRows(), userID)
	if err != nil {
		return nil, err
	}
	userIDs := lo.FlatMap(reply.GetRows(), func(item *contentv1.Comment, _ int) []int64 {
		return r.commentProfileIDs(item, articles[item.GetArticleId()])
	})
	profiles, err := r.loadAccountProfiles(ctx, userIDs...)
	if err != nil {
		return nil, err
	}
	states, err := r.loadCommentViewerActionStates(ctx, lo.Map(reply.GetRows(), func(item *contentv1.Comment, _ int) int64 {
		return item.GetId()
	}))
	if err != nil {
		return nil, err
	}
	rows := lo.Map(reply.GetRows(), func(item *contentv1.Comment, _ int) *bbscontentv1.CommentListItem {
		return r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()])
	})
	return &bbscontentv1.ListCommentTimeline_Reply{
		Page: reply.GetPage(),
		Rows: rows,
	}, nil
}

func (r *ContentCommentClient) LikeComment(ctx context.Context, req *bbscontentv1.LikeComment_Request) (*bbscontentv1.LikeComment_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Comment.Like(ctx, &contentv1.LikeComment_Request{Id: req.GetId(), Liked: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.LikeComment_Reply{Liked: reply.GetLiked()}, nil
}

func (r *ContentCommentClient) ThankComment(ctx context.Context, req *bbscontentv1.ThankComment_Request) (*bbscontentv1.ThankComment_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Comment.Thank(ctx, &contentv1.ThankComment_Request{Id: req.GetId(), Thanked: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.ThankComment_Reply{Thanked: reply.GetThanked()}, nil
}

func (r *ContentCommentClient) loadCommentViewerActionStates(ctx context.Context, commentIDs []int64) (map[int64]*bbscontentv1.CommentViewerActionState, error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil || user.ID <= 0 {
		return map[int64]*bbscontentv1.CommentViewerActionState{}, nil
	}
	commentIDs = lo.Uniq(commentIDs)
	if len(commentIDs) == 0 {
		return map[int64]*bbscontentv1.CommentViewerActionState{}, nil
	}
	reply, err := r.contentClient.Comment.MapViewerActionStates(ctx, &contentv1.MapCommentViewerActionStates_Request{
		CommentIds: commentIDs,
		UserId:     user.ID,
	})
	if err != nil {
		return nil, err
	}
	states := make(map[int64]*bbscontentv1.CommentViewerActionState, len(reply.GetStates()))
	for commentID, state := range reply.GetStates() {
		if state == nil {
			continue
		}
		states[commentID] = &bbscontentv1.CommentViewerActionState{
			Liked:   state.GetLiked(),
			Thanked: state.GetThanked(),
		}
	}
	return states, nil
}

func (r *ContentCommentClient) loadCommentArticles(ctx context.Context, comments []*contentv1.Comment, viewerUserID int64) (map[int64]*contentv1.Article, error) {
	articleIDs := lo.Uniq(lo.FilterMap(comments, func(item *contentv1.Comment, _ int) (int64, bool) {
		if item == nil || item.GetArticleId() == 0 {
			return 0, false
		}
		return item.GetArticleId(), true
	}))
	articles := make(map[int64]*contentv1.Article, len(articleIDs))
	for _, articleID := range articleIDs {
		reply, err := r.contentClient.Article.Get(ctx, &contentv1.GetArticle_Request{
			ArticleId: articleID,
		})
		if err != nil {
			return nil, err
		}
		article := reply.GetArticle()
		isAuthor := article.CreatedBy != nil && *article.CreatedBy == viewerUserID
		if article.GetRestriction() == contentv1.ContentRestriction_CONTENT_RESTRICTION_HIDDEN {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		switch article.GetPublishStatus() {
		case contentv1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_DRAFT:
			if !isAuthor {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
			}
		case contentv1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_PUBLISHED, contentv1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_ARCHIVED:
			if article.GetVisibility() == contentv1.ArticleVisibility_ARTICLE_VISIBILITY_PRIVATE && !isAuthor {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
			}
		default:
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		articles[articleID] = article
	}
	return articles, nil
}

func (r *ContentCommentClient) commentProfileIDs(item *contentv1.Comment, article *contentv1.Article) []int64 {
	if item == nil {
		return nil
	}
	userIDs := make([]int64, 0, 2)
	if item.CreatedBy != nil && !r.anonymousArticleUser(article, item.CreatedBy) {
		userIDs = append(userIDs, *item.CreatedBy)
	}
	if item.ReplyUserId != nil && !r.anonymousArticleUser(article, item.ReplyUserId) {
		userIDs = append(userIDs, *item.ReplyUserId)
	}
	return userIDs
}

func (r *ContentCommentClient) commentListItem(item *contentv1.Comment, article *contentv1.Article, profiles map[int64]*bbsuserv1.AccountProfile, state *bbscontentv1.CommentViewerActionState) *bbscontentv1.CommentListItem {
	if item == nil {
		return nil
	}
	if state == nil {
		state = &bbscontentv1.CommentViewerActionState{}
	}
	out := &bbscontentv1.CommentListItem{
		Id:                item.GetId(),
		ArticleId:         item.GetArticleId(),
		Content:           item.GetContent(),
		ContentRender:     r.commentContentRender(item.GetId(), item.GetContent()),
		Level:             item.GetLevel(),
		ParentId:          item.ParentId,
		ReplyId:           item.ReplyId,
		Restriction:       bbscontentv1.ContentRestriction(item.GetRestriction()),
		DeletedAt:         formatProtoTime(item.GetDeletedAt()),
		ThankCount:        item.GetThankCount(),
		LikeCount:         item.GetLikeCount(),
		ReplyCount:        item.GetReplyCount(),
		ViewerActionState: state,
		CreatedAt:         formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:         formatProtoTime(item.GetUpdatedAt()),
	}
	if !r.anonymousArticleUser(article, item.CreatedBy) {
		out.CreatedBy = item.CreatedBy
		out.UpdatedBy = item.UpdatedBy
	}
	if item.CreatedBy != nil && !r.anonymousArticleUser(article, item.CreatedBy) {
		out.User = profiles[*item.CreatedBy]
	}
	if item.ReplyUserId != nil && !r.anonymousArticleUser(article, item.ReplyUserId) {
		out.ReplyUser = profiles[*item.ReplyUserId]
	}
	return out
}

func (r *ContentCommentClient) commentDetail(item *contentv1.Comment, article *contentv1.Article, profiles map[int64]*bbsuserv1.AccountProfile) *bbscontentv1.CommentDetail {
	if item == nil {
		return nil
	}
	out := &bbscontentv1.CommentDetail{
		Id:                item.GetId(),
		ArticleId:         item.GetArticleId(),
		Content:           item.GetContent(),
		ContentRender:     r.commentContentRender(item.GetId(), item.GetContent()),
		Level:             item.GetLevel(),
		ParentId:          item.ParentId,
		ReplyId:           item.ReplyId,
		Restriction:       bbscontentv1.ContentRestriction(item.GetRestriction()),
		DeletedAt:         formatProtoTime(item.GetDeletedAt()),
		ThankCount:        item.GetThankCount(),
		LikeCount:         item.GetLikeCount(),
		ReplyCount:        item.GetReplyCount(),
		ViewerActionState: &bbscontentv1.CommentViewerActionState{},
		CreatedAt:         formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:         formatProtoTime(item.GetUpdatedAt()),
	}
	if !r.anonymousArticleUser(article, item.CreatedBy) {
		out.CreatedBy = item.CreatedBy
		out.UpdatedBy = item.UpdatedBy
	}
	if item.CreatedBy != nil && !r.anonymousArticleUser(article, item.CreatedBy) {
		out.User = profiles[*item.CreatedBy]
	}
	if item.ReplyUserId != nil && !r.anonymousArticleUser(article, item.ReplyUserId) {
		out.ReplyUser = profiles[*item.ReplyUserId]
	}
	return out
}

func (r *ContentCommentClient) anonymousArticleUser(article *contentv1.Article, userID *int64) bool {
	return article != nil && article.GetAnonymous() && userID != nil && article.CreatedBy != nil && *article.CreatedBy == *userID
}

func (r *ContentCommentClient) commentContentRender(commentID int64, content string) string {
	return util.LuteEngine.MarkdownStr(fmt.Sprintf("comment_%d", commentID), content)
}

func (r *ContentCommentClient) loadAccountProfiles(ctx context.Context, userIDs ...int64) (map[int64]*bbsuserv1.AccountProfile, error) {
	if r.userClient == nil {
		return map[int64]*bbsuserv1.AccountProfile{}, nil
	}
	ids := lo.Filter(lo.Uniq(userIDs), func(userID int64, _ int) bool {
		return userID != 0
	})
	if len(ids) == 0 {
		return map[int64]*bbsuserv1.AccountProfile{}, nil
	}
	reply, err := r.userClient.Account.Map(ctx, &userv1.MapAccounts_Request{
		Query: &userv1.AccountQuery{UserIds: ids},
	})
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*bbsuserv1.AccountProfile, len(reply.GetAccounts()))
	for userID, account := range reply.GetAccounts() {
		basic := account.GetBasic()
		if basic == nil {
			continue
		}
		out[userID] = &bbsuserv1.AccountProfile{
			Id:            basic.GetId(),
			Name:          basic.GetName(),
			Nickname:      basic.Nickname,
			Url:           basic.Url,
			AvatarUrl:     basic.AvatarUrl,
			Introduction:  basic.Introduction,
			Mbti:          bbsuserv1.MBTI(basic.GetMbti()),
			Status:        bbsuserv1.AccountStatus(basic.GetStatus()),
			FollowCount:   basic.FollowCount,
			FollowerCount: basic.FollowerCount,
			CreatedAt:     formatProtoTime(basic.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(basic.GetUpdatedAt()),
		}
	}
	return out, nil
}
