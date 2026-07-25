package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/pkg/util"
	"common/proto/gen/common"
	contentv1 "common/proto/gen/content/v1"
	contentv1enum "common/proto/gen/content/v1/enum"
	userv1 "common/proto/gen/user/v1"
	"context"
	"fmt"

	"github.com/samber/lo"
)

var _ repo.ContentCommentClient = (*ContentCommentClient)(nil)

type ContentCommentClient struct {
	protoTimeFormatter
	contentClient *rpc.ContentClient
	userClient    *rpc.UserClient
}

func NewContentCommentClient(
	contentClient *rpc.ContentClient,
	userClient *rpc.UserClient,
) repo.ContentCommentClient {
	return &ContentCommentClient{
		contentClient: contentClient,
		userClient:    userClient,
	}
}

func (r *ContentCommentClient) CreateComment(ctx context.Context, req *repo.CreateCommentReq) (*repo.CommentDetail, error) {
	var replyID *int64
	if req.ReplyID != 0 {
		replyID = new(req.ReplyID)
	}
	reply, err := r.contentClient.Comment.Create(ctx, &contentv1.CreateComment_Req{
		ArticleId: req.ArticleID,
		Content:   req.Content,
		ReplyId:   replyID,
		UserId:    req.UserID,
	})
	if err != nil {
		return nil, err
	}
	item := cloneDataMessage(reply.GetComment(), &contentv1.PageComments_Resp_Comment{})
	articles, err := r.loadCommentArticles(ctx, []*contentv1.PageComments_Resp_Comment{item}, req.UserID)
	if err != nil {
		return nil, err
	}
	profiles, err := r.loadAccountProfiles(ctx, r.commentProfileIDs(item, articles[item.GetArticleId()])...)
	if err != nil {
		return nil, err
	}
	return r.commentDetail(item, articles[item.GetArticleId()], profiles), nil
}

func (r *ContentCommentClient) ListComments(ctx context.Context, req *repo.ListCommentsReq) (*repo.ListCommentsResp, error) {
	var pageReq *common.PageReq
	if req.Page != nil {
		pageReq = &common.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	query := req.Query
	if query == nil {
		query = &repo.CommentQuery{}
	}
	contentQuery := &contentv1.PageComments_Req_CommentQueryParams{
		CommentId: query.CommentID,
		ArticleId: query.ArticleID,
		ParentId:  query.ParentID,
		ReplyId:   query.ReplyID,
		Level:     query.Level,
		UserId:    query.UserID,
	}
	if query.Restriction != nil {
		contentQuery.Restriction = new(contentv1enum.ContentRestriction(*query.Restriction))
	}
	if len(query.Restrictions) > 0 {
		contentQuery.Restrictions = lo.Map(query.Restrictions, func(item int32, _ int) contentv1enum.ContentRestriction {
			return contentv1enum.ContentRestriction(item)
		})
	}
	if query.Order != nil {
		contentQuery.Order = new(contentv1enum.CommentOrder(*query.Order))
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
		Page:  pageReq,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	articles, err := r.loadCommentArticles(ctx, reply.GetRows(), req.UserID)
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
	states, err := r.loadCommentViewerActionStates(ctx, req.UserID, lo.Map(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) int64 {
		return item.GetId()
	}))
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.CommentListItem, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()]))
	}
	var page *repo.PageResp
	if reply.GetPage() != nil {
		page = &repo.PageResp{
			Page:  reply.GetPage().GetPage(),
			Size:  reply.GetPage().GetSize(),
			Total: reply.GetPage().GetTotal(),
		}
	}
	return &repo.ListCommentsResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *ContentCommentClient) ListCommentThreads(ctx context.Context, req *repo.ListCommentThreadsReq) (*repo.ListCommentThreadsResp, error) {
	var pageReq *common.PageReq
	if req.Page != nil {
		pageReq = &common.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	normal := contentv1enum.ContentRestriction_CONTENT_RESTRICTION_NONE
	locked := contentv1enum.ContentRestriction_CONTENT_RESTRICTION_LOCKED
	restrictions := []contentv1enum.ContentRestriction{normal, locked}
	order := contentv1enum.CommentOrder_COMMENT_ORDER_HOTTEST
	if req.Order != nil && *req.Order != int32(contentv1enum.CommentOrder_COMMENT_ORDER_UNSPECIFIED) {
		order = contentv1enum.CommentOrder(*req.Order)
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
		Page: pageReq,
		Query: &contentv1.PageComments_Req_CommentQueryParams{
			ArticleId:    new(req.ArticleID),
			Level:        new(int32(1)),
			Restrictions: restrictions,
			Order:        new(order),
		},
	})
	if err != nil {
		return nil, err
	}
	rootIDs := lo.Map(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) int64 {
		return item.GetId()
	})
	previewMap := map[int64][]*contentv1.PageComments_Resp_Comment{}
	if len(rootIDs) > 0 {
		previewLimit := int32(0)
		if req.ReplyPreviewLimit != nil {
			previewLimit = *req.ReplyPreviewLimit
		}
		if previewLimit <= 0 {
			previewLimit = 3
		}
		if previewLimit > 5 {
			previewLimit = 5
		}
		previews, err := r.contentClient.Comment.ListReplyPreviews(ctx, &contentv1.ListCommentReplyPreviews_Req{
			ArticleId:      req.ArticleID,
			ParentIds:      rootIDs,
			LimitPerParent: previewLimit,
			Order:          new(contentv1enum.CommentOrder_COMMENT_ORDER_OLDEST),
			Restrictions:   restrictions,
		})
		if err != nil {
			return nil, err
		}
		previewMap = lo.SliceToMap(previews.GetRows(), func(item *contentv1.ListCommentReplyPreviews_Resp_CommentReplyPreview) (int64, []*contentv1.PageComments_Resp_Comment) {
			rows := lo.Map(item.GetRows(), func(row *contentv1.ListCommentReplyPreviews_Resp_Comment, _ int) *contentv1.PageComments_Resp_Comment {
				return cloneDataMessage(row, &contentv1.PageComments_Resp_Comment{})
			})
			return item.GetParentId(), rows
		})
	}
	allComments := append([]*contentv1.PageComments_Resp_Comment{}, reply.GetRows()...)
	for _, rows := range previewMap {
		allComments = append(allComments, rows...)
	}
	articles, err := r.loadCommentArticles(ctx, allComments, req.UserID)
	if err != nil {
		return nil, err
	}
	userIDs := lo.FlatMap(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) []int64 {
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
	states, err := r.loadCommentViewerActionStates(ctx, req.UserID, commentIDs)
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.CommentThread, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		previewRows := lo.Map(previewMap[item.GetId()], func(preview *contentv1.PageComments_Resp_Comment, _ int) *repo.CommentListItem {
			return r.commentListItem(preview, articles[preview.GetArticleId()], profiles, states[preview.GetId()])
		})
		rows = append(rows, &repo.CommentThread{
			Root:           r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()]),
			PreviewReplies: previewRows,
			ReplyCount:     item.GetReplyCount(),
			HasMoreReplies: item.GetReplyCount() > int32(len(previewRows)),
		})
	}
	var page *repo.PageResp
	if reply.GetPage() != nil {
		page = &repo.PageResp{
			Page:  reply.GetPage().GetPage(),
			Size:  reply.GetPage().GetSize(),
			Total: reply.GetPage().GetTotal(),
		}
	}
	return &repo.ListCommentThreadsResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *ContentCommentClient) ListCommentReplies(ctx context.Context, req *repo.ListCommentRepliesReq) (*repo.ListCommentRepliesResp, error) {
	var pageReq *common.PageReq
	if req.Page != nil {
		pageReq = &common.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	normal := contentv1enum.ContentRestriction_CONTENT_RESTRICTION_NONE
	locked := contentv1enum.ContentRestriction_CONTENT_RESTRICTION_LOCKED
	restrictions := []contentv1enum.ContentRestriction{normal, locked}
	order := contentv1enum.CommentOrder_COMMENT_ORDER_OLDEST
	if req.Order != nil && *req.Order != int32(contentv1enum.CommentOrder_COMMENT_ORDER_UNSPECIFIED) {
		order = contentv1enum.CommentOrder(*req.Order)
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
		Page: pageReq,
		Query: &contentv1.PageComments_Req_CommentQueryParams{
			ArticleId:    new(req.ArticleID),
			ParentId:     new(req.ParentID),
			Restrictions: restrictions,
			Order:        new(order),
		},
	})
	if err != nil {
		return nil, err
	}
	articles, err := r.loadCommentArticles(ctx, reply.GetRows(), req.UserID)
	if err != nil {
		return nil, err
	}
	userIDs := lo.FlatMap(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) []int64 {
		return r.commentProfileIDs(item, articles[item.GetArticleId()])
	})
	profiles, err := r.loadAccountProfiles(ctx, userIDs...)
	if err != nil {
		return nil, err
	}
	states, err := r.loadCommentViewerActionStates(ctx, req.UserID, lo.Map(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) int64 {
		return item.GetId()
	}))
	if err != nil {
		return nil, err
	}
	rows := lo.Map(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) *repo.CommentListItem {
		return r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()])
	})
	var page *repo.PageResp
	if reply.GetPage() != nil {
		page = &repo.PageResp{
			Page:  reply.GetPage().GetPage(),
			Size:  reply.GetPage().GetSize(),
			Total: reply.GetPage().GetTotal(),
		}
	}
	return &repo.ListCommentRepliesResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *ContentCommentClient) ListCommentTimeline(ctx context.Context, req *repo.ListCommentTimelineReq) (*repo.ListCommentTimelineResp, error) {
	var pageReq *common.PageReq
	if req.Page != nil {
		pageReq = &common.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	normal := contentv1enum.ContentRestriction_CONTENT_RESTRICTION_NONE
	locked := contentv1enum.ContentRestriction_CONTENT_RESTRICTION_LOCKED
	restrictions := []contentv1enum.ContentRestriction{normal, locked}
	order := contentv1enum.CommentOrder_COMMENT_ORDER_NEWEST
	if req.Order != nil && *req.Order != int32(contentv1enum.CommentOrder_COMMENT_ORDER_UNSPECIFIED) {
		order = contentv1enum.CommentOrder(*req.Order)
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
		Page: pageReq,
		Query: &contentv1.PageComments_Req_CommentQueryParams{
			ArticleId:    new(req.ArticleID),
			Restrictions: restrictions,
			Order:        new(order),
		},
	})
	if err != nil {
		return nil, err
	}
	articles, err := r.loadCommentArticles(ctx, reply.GetRows(), req.UserID)
	if err != nil {
		return nil, err
	}
	userIDs := lo.FlatMap(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) []int64 {
		return r.commentProfileIDs(item, articles[item.GetArticleId()])
	})
	profiles, err := r.loadAccountProfiles(ctx, userIDs...)
	if err != nil {
		return nil, err
	}
	states, err := r.loadCommentViewerActionStates(ctx, req.UserID, lo.Map(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) int64 {
		return item.GetId()
	}))
	if err != nil {
		return nil, err
	}
	rows := lo.Map(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) *repo.CommentListItem {
		return r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()])
	})
	var page *repo.PageResp
	if reply.GetPage() != nil {
		page = &repo.PageResp{
			Page:  reply.GetPage().GetPage(),
			Size:  reply.GetPage().GetSize(),
			Total: reply.GetPage().GetTotal(),
		}
	}
	return &repo.ListCommentTimelineResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *ContentCommentClient) LikeComment(ctx context.Context, req *repo.LikeCommentReq) (bool, error) {
	reply, err := r.contentClient.Comment.Like(ctx, &contentv1.LikeComment_Req{
		Id:     req.ID,
		Liked:  req.Active,
		UserId: req.UserID,
	})
	if err != nil {
		return false, err
	}
	return reply.GetLiked(), nil
}

func (r *ContentCommentClient) ThankComment(ctx context.Context, req *repo.ThankCommentReq) (bool, error) {
	reply, err := r.contentClient.Comment.Thank(ctx, &contentv1.ThankComment_Req{
		Id:      req.ID,
		Thanked: req.Active,
		UserId:  req.UserID,
	})
	if err != nil {
		return false, err
	}
	return reply.GetThanked(), nil
}

func (r *ContentCommentClient) loadCommentViewerActionStates(ctx context.Context, userID int64, commentIDs []int64) (map[int64]*repo.CommentViewerActionState, error) {
	commentIDs = lo.Uniq(commentIDs)
	if len(commentIDs) == 0 {
		return map[int64]*repo.CommentViewerActionState{}, nil
	}
	reply, err := r.contentClient.Comment.MapViewerActionStates(ctx, &contentv1.MapCommentViewerActionStates_Req{
		CommentIds: commentIDs,
		UserId:     userID,
	})
	if err != nil {
		return nil, err
	}
	states := make(map[int64]*repo.CommentViewerActionState, len(reply.GetStates()))
	for commentID, state := range reply.GetStates() {
		if state == nil {
			continue
		}
		states[commentID] = &repo.CommentViewerActionState{
			Liked:   state.GetLiked(),
			Thanked: state.GetThanked(),
		}
	}
	return states, nil
}

func (r *ContentCommentClient) loadCommentArticles(ctx context.Context, comments []*contentv1.PageComments_Resp_Comment, viewerUserID int64) (map[int64]*contentv1.GetArticle_Resp_Article, error) {
	articleIDs := lo.Uniq(lo.FilterMap(comments, func(item *contentv1.PageComments_Resp_Comment, _ int) (int64, bool) {
		if item == nil || item.GetArticleId() == 0 {
			return 0, false
		}
		return item.GetArticleId(), true
	}))
	articles := make(map[int64]*contentv1.GetArticle_Resp_Article, len(articleIDs))
	for _, articleID := range articleIDs {
		reply, err := r.contentClient.Article.Get(ctx, &contentv1.GetArticle_Req{
			ArticleId: articleID,
		})
		if err != nil {
			return nil, err
		}
		article := reply.GetArticle()

		articles[articleID] = article
	}
	return articles, nil
}

func (r *ContentCommentClient) commentProfileIDs(item *contentv1.PageComments_Resp_Comment, article *contentv1.GetArticle_Resp_Article) []int64 {
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

func (r *ContentCommentClient) commentListItem(item *contentv1.PageComments_Resp_Comment, article *contentv1.GetArticle_Resp_Article, profiles map[int64]*repo.AccountProfile, state *repo.CommentViewerActionState) *repo.CommentListItem {
	if item == nil {
		return nil
	}
	if state == nil {
		state = &repo.CommentViewerActionState{}
	}
	out := &repo.CommentListItem{
		ID:                item.GetId(),
		ArticleID:         item.GetArticleId(),
		Content:           item.GetContent(),
		ContentRender:     r.commentContentRender(item.GetId(), item.GetContent()),
		Level:             item.GetLevel(),
		ParentID:          item.ParentId,
		ReplyID:           item.ReplyId,
		Restriction:       int32(item.GetRestriction()),
		DeletedAt:         r.formatProtoTime(item.GetDeletedAt()),
		ThankCount:        item.GetThankCount(),
		LikeCount:         item.GetLikeCount(),
		ReplyCount:        item.GetReplyCount(),
		ViewerActionState: state,
		CreatedAt:         r.formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:         r.formatProtoTime(item.GetUpdatedAt()),
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

func (r *ContentCommentClient) commentDetail(item *contentv1.PageComments_Resp_Comment, article *contentv1.GetArticle_Resp_Article, profiles map[int64]*repo.AccountProfile) *repo.CommentDetail {
	if item == nil {
		return nil
	}
	out := &repo.CommentDetail{
		ID:                item.GetId(),
		ArticleID:         item.GetArticleId(),
		Content:           item.GetContent(),
		ContentRender:     r.commentContentRender(item.GetId(), item.GetContent()),
		Level:             item.GetLevel(),
		ParentID:          item.ParentId,
		ReplyID:           item.ReplyId,
		Restriction:       int32(item.GetRestriction()),
		DeletedAt:         r.formatProtoTime(item.GetDeletedAt()),
		ThankCount:        item.GetThankCount(),
		LikeCount:         item.GetLikeCount(),
		ReplyCount:        item.GetReplyCount(),
		ViewerActionState: &repo.CommentViewerActionState{},
		CreatedAt:         r.formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:         r.formatProtoTime(item.GetUpdatedAt()),
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

func (r *ContentCommentClient) anonymousArticleUser(article *contentv1.GetArticle_Resp_Article, userID *int64) bool {
	return article != nil && article.GetAnonymous() && userID != nil && article.CreatedBy != nil && *article.CreatedBy == *userID
}

func (r *ContentCommentClient) commentContentRender(commentID int64, content string) string {
	return util.LuteEngine.MarkdownStr(fmt.Sprintf("comment_%d", commentID), content)
}

func (r *ContentCommentClient) loadAccountProfiles(ctx context.Context, userIDs ...int64) (map[int64]*repo.AccountProfile, error) {
	if r.userClient == nil {
		return map[int64]*repo.AccountProfile{}, nil
	}
	ids := lo.Filter(lo.Uniq(userIDs), func(userID int64, _ int) bool {
		return userID != 0
	})
	if len(ids) == 0 {
		return map[int64]*repo.AccountProfile{}, nil
	}
	reply, err := r.userClient.Account.Map(ctx, &userv1.MapAccounts_Req{
		Query: &userv1.MapAccounts_Req_AccountQuery{
			UserIds: ids,
		},
	})
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*repo.AccountProfile, len(reply.GetAccounts()))
	for userID, account := range reply.GetAccounts() {
		basic := account.GetBasic()
		if basic == nil {
			continue
		}
		out[userID] = &repo.AccountProfile{
			ID:            basic.GetId(),
			Name:          basic.GetName(),
			Nickname:      basic.Nickname,
			URL:           basic.Url,
			AvatarURL:     basic.AvatarUrl,
			Introduction:  basic.Introduction,
			MBTI:          int32(basic.GetMbti()),
			Status:        int32(basic.GetStatus()),
			FollowCount:   basic.FollowCount,
			FollowerCount: basic.FollowerCount,
			CreatedAt:     r.formatProtoTime(basic.GetCreatedAt()),
			UpdatedAt:     r.formatProtoTime(basic.GetUpdatedAt()),
		}
	}
	return out, nil
}
