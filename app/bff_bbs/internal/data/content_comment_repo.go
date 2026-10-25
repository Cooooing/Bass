package data

import (
	"common/pkg/client/rpc"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	contentv1 "common/proto/gen/content/v1"
	contentv1enum "common/proto/gen/content/v1/enum"
	"context"
	"fmt"

	"bff_bbs/internal/biz/model"

	"bff_bbs/internal/biz/repo"
	"bff_bbs/internal/enum"

	"github.com/samber/lo"
)

var _ repo.ContentCommentClient = (*ContentCommentClient)(nil)

type ContentCommentClient struct {
	accountProfileLoader
	contentClient        *rpc.ContentClient
	restrictionConverter *enum.ContentRestrictionConverter
}

func NewContentCommentClient(
	contentClient *rpc.ContentClient,
	userClient *rpc.UserClient,
) repo.ContentCommentClient {
	return &ContentCommentClient{
		accountProfileLoader: accountProfileLoader{userClient: userClient},
		contentClient:        contentClient,
		restrictionConverter: enum.NewContentRestrictionConverter(),
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
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return nil, err
	}
	comment := reply.GetComment()
	item := &contentv1.PageComments_Resp_Comment{}
	if comment != nil {
		item.CreatedAt = comment.GetCreatedAt()
		item.UpdatedAt = comment.GetUpdatedAt()
		item.CreatedBy = comment.CreatedBy
		item.UpdatedBy = comment.UpdatedBy
		item.Id = comment.GetId()
		item.ArticleId = comment.GetArticleId()
		item.Content = comment.GetContent()
		item.Level = comment.GetLevel()
		item.ParentId = comment.ParentId
		item.ReplyId = comment.ReplyId
		item.ThankCount = comment.GetThankCount()
		item.LikeCount = comment.GetLikeCount()
		item.ReplyCount = comment.GetReplyCount()
		item.ReplyUserId = comment.ReplyUserId
		item.Restriction = comment.GetRestriction()
		item.DeletedAt = comment.GetDeletedAt()
	}
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
	var pageReq *commonmodel.PageReq
	if req.Page != nil {
		pageReq = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	query := req.Query
	if query == nil {
		query = &repo.CommentQuery{}
	}
	contentQuery := &contentv1.PageComments_Req_CommentQueryParams{
		CommentId:       query.CommentID,
		ArticleId:       query.ArticleID,
		ParentId:        query.ParentID,
		ReplyId:         query.ReplyID,
		Level:           query.Level,
		UserId:          query.UserID,
		ExcludedUserIds: query.ExcludedUserIDs,
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
	accessScope := contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST
	contentAccess := &contentv1.ContentAccess{Scope: accessScope}
	if req.UserID > 0 {
		accessScope = contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER
		contentAccess = &contentv1.ContentAccess{Scope: accessScope, ActorUserId: new(req.UserID)}
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
		Page:   pageReq.ToProto(),
		Query:  contentQuery,
		Access: contentAccess,
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
	rows := make([]*model.CommentListItem, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()]))
	}
	var page *commonmodel.PageResp
	if reply.GetPage() != nil {
		page = commonmodel.PageRespFromProto(reply.GetPage())
	}
	return &repo.ListCommentsResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *ContentCommentClient) ListCommentThreads(ctx context.Context, req *repo.ListCommentThreadsReq) (*repo.ListCommentThreadsResp, error) {
	var pageReq *commonmodel.PageReq
	if req.Page != nil {
		pageReq = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	order := contentv1enum.CommentOrder_COMMENT_ORDER_HOTTEST
	if req.Order != nil && *req.Order != int32(contentv1enum.CommentOrder_COMMENT_ORDER_UNSPECIFIED) {
		order = contentv1enum.CommentOrder(*req.Order)
	}
	var acceptedRoot *contentv1.PageComments_Resp_Comment
	if req.AcceptedAnswerID != nil {
		acceptedReply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
			Page:   (&commonmodel.PageReq{Page: 1, Size: 1}).ToProto(),
			Query:  &contentv1.PageComments_Req_CommentQueryParams{CommentId: req.AcceptedAnswerID, ArticleId: new(req.ArticleID), Level: new(int32(1)), ExcludedUserIds: req.ExcludedUserIDs},
			Access: &contentv1.ContentAccess{Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST},
		})
		if err != nil {
			return nil, err
		}
		if len(acceptedReply.GetRows()) == 1 {
			acceptedRoot = acceptedReply.GetRows()[0]
		}
	}
	var pinnedRoot *contentv1.PageComments_Resp_Comment
	if req.PinnedCommentID != nil {
		pinnedReply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
			Page:   (&commonmodel.PageReq{Page: 1, Size: 1}).ToProto(),
			Query:  &contentv1.PageComments_Req_CommentQueryParams{CommentId: req.PinnedCommentID, ArticleId: new(req.ArticleID), Level: new(int32(1)), ExcludedUserIds: req.ExcludedUserIDs},
			Access: &contentv1.ContentAccess{Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST},
		})
		if err != nil {
			return nil, err
		}
		if len(pinnedReply.GetRows()) == 1 {
			pinnedRoot = pinnedReply.GetRows()[0]
		}
	}
	excludedCommentIDs := []int64(nil)
	if acceptedRoot != nil {
		excludedCommentIDs = append(excludedCommentIDs, acceptedRoot.GetId())
	}
	if pinnedRoot != nil {
		excludedCommentIDs = append(excludedCommentIDs, pinnedRoot.GetId())
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
		Page: pageReq.ToProto(),
		Query: &contentv1.PageComments_Req_CommentQueryParams{
			ArticleId:          new(req.ArticleID),
			Level:              new(int32(1)),
			Order:              new(order),
			ExcludedUserIds:    req.ExcludedUserIDs,
			ExcludedCommentIds: excludedCommentIDs,
		},
		Access: &contentv1.ContentAccess{Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST},
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
			ArticleId:       req.ArticleID,
			ParentIds:       rootIDs,
			LimitPerParent:  previewLimit,
			Order:           new(contentv1enum.CommentOrder_COMMENT_ORDER_OLDEST),
			Access:          &contentv1.ContentAccess{Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST},
			ExcludedUserIds: req.ExcludedUserIDs,
		})
		if err != nil {
			return nil, err
		}
		previewMap = lo.SliceToMap(previews.GetRows(), func(item *contentv1.ListCommentReplyPreviews_Resp_CommentReplyPreview) (int64, []*contentv1.PageComments_Resp_Comment) {
			rows := lo.Map(item.GetRows(), func(row *contentv1.ListCommentReplyPreviews_Resp_Comment, _ int) *contentv1.PageComments_Resp_Comment {
				return &contentv1.PageComments_Resp_Comment{
					CreatedAt:   row.GetCreatedAt(),
					UpdatedAt:   row.GetUpdatedAt(),
					CreatedBy:   row.CreatedBy,
					UpdatedBy:   row.UpdatedBy,
					Id:          row.GetId(),
					ArticleId:   row.GetArticleId(),
					Content:     row.GetContent(),
					Level:       row.GetLevel(),
					ParentId:    row.ParentId,
					ReplyId:     row.ReplyId,
					ThankCount:  row.GetThankCount(),
					LikeCount:   row.GetLikeCount(),
					ReplyCount:  row.GetReplyCount(),
					ReplyUserId: row.ReplyUserId,
					Restriction: row.GetRestriction(),
					DeletedAt:   row.GetDeletedAt(),
				}
			})
			return item.GetParentId(), rows
		})
	}
	allComments := append([]*contentv1.PageComments_Resp_Comment{}, reply.GetRows()...)
	if acceptedRoot != nil {
		allComments = append(allComments, acceptedRoot)
	}
	if pinnedRoot != nil {
		allComments = append(allComments, pinnedRoot)
	}
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
	if acceptedRoot != nil {
		userIDs = append(userIDs, r.commentProfileIDs(acceptedRoot, articles[acceptedRoot.GetArticleId()])...)
	}
	if pinnedRoot != nil {
		userIDs = append(userIDs, r.commentProfileIDs(pinnedRoot, articles[pinnedRoot.GetArticleId()])...)
	}
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
	if acceptedRoot != nil {
		commentIDs = append(commentIDs, acceptedRoot.GetId())
	}
	if pinnedRoot != nil {
		commentIDs = append(commentIDs, pinnedRoot.GetId())
	}
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
		previewRows := lo.Map(previewMap[item.GetId()], func(preview *contentv1.PageComments_Resp_Comment, _ int) *model.CommentListItem {
			return r.commentListItem(preview, articles[preview.GetArticleId()], profiles, states[preview.GetId()])
		})
		rows = append(rows, &repo.CommentThread{
			Root:           r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()]),
			PreviewReplies: previewRows,
			ReplyCount:     item.GetReplyCount(),
			HasMoreReplies: int64(item.GetReplyCount()) > int64(len(previewRows)),
		})
	}
	var page *commonmodel.PageResp
	if reply.GetPage() != nil {
		page = commonmodel.PageRespFromProto(reply.GetPage())
	}
	result := &repo.ListCommentThreadsResp{
		Page: page,
		Rows: rows,
	}
	if pinnedRoot != nil {
		result.PinnedComment = &repo.CommentThread{Root: r.commentListItem(pinnedRoot, articles[pinnedRoot.GetArticleId()], profiles, states[pinnedRoot.GetId()]), ReplyCount: pinnedRoot.GetReplyCount(), HasMoreReplies: pinnedRoot.GetReplyCount() > 0}
	}
	if acceptedRoot != nil {
		result.AcceptedAnswer = &repo.CommentThread{Root: r.commentListItem(acceptedRoot, articles[acceptedRoot.GetArticleId()], profiles, states[acceptedRoot.GetId()]), ReplyCount: acceptedRoot.GetReplyCount(), HasMoreReplies: acceptedRoot.GetReplyCount() > 0}
	}
	return result, nil
}

func (r *ContentCommentClient) ListCommentReplies(ctx context.Context, req *repo.ListCommentRepliesReq) (*repo.ListCommentRepliesResp, error) {
	var pageReq *commonmodel.PageReq
	if req.Page != nil {
		pageReq = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	order := contentv1enum.CommentOrder_COMMENT_ORDER_OLDEST
	if req.Order != nil && *req.Order != int32(contentv1enum.CommentOrder_COMMENT_ORDER_UNSPECIFIED) {
		order = contentv1enum.CommentOrder(*req.Order)
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
		Page: pageReq.ToProto(),
		Query: &contentv1.PageComments_Req_CommentQueryParams{
			ArticleId:       new(req.ArticleID),
			ParentId:        new(req.ParentID),
			Order:           new(order),
			ExcludedUserIds: req.ExcludedUserIDs,
		},
		Access: &contentv1.ContentAccess{Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST},
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
	rows := lo.Map(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) *model.CommentListItem {
		return r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()])
	})
	var page *commonmodel.PageResp
	if reply.GetPage() != nil {
		page = commonmodel.PageRespFromProto(reply.GetPage())
	}
	return &repo.ListCommentRepliesResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *ContentCommentClient) ListCommentTimeline(ctx context.Context, req *repo.ListCommentTimelineReq) (*repo.ListCommentTimelineResp, error) {
	var pageReq *commonmodel.PageReq
	if req.Page != nil {
		pageReq = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	order := contentv1enum.CommentOrder_COMMENT_ORDER_NEWEST
	if req.Order != nil && *req.Order != int32(contentv1enum.CommentOrder_COMMENT_ORDER_UNSPECIFIED) {
		order = contentv1enum.CommentOrder(*req.Order)
	}
	reply, err := r.contentClient.Comment.Page(ctx, &contentv1.PageComments_Req{
		Page: pageReq.ToProto(),
		Query: &contentv1.PageComments_Req_CommentQueryParams{
			ArticleId:       new(req.ArticleID),
			Order:           new(order),
			ExcludedUserIds: req.ExcludedUserIDs,
		},
		Access: &contentv1.ContentAccess{Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST},
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
	rows := lo.Map(reply.GetRows(), func(item *contentv1.PageComments_Resp_Comment, _ int) *model.CommentListItem {
		return r.commentListItem(item, articles[item.GetArticleId()], profiles, states[item.GetId()])
	})
	var page *commonmodel.PageResp
	if reply.GetPage() != nil {
		page = commonmodel.PageRespFromProto(reply.GetPage())
	}
	return &repo.ListCommentTimelineResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *ContentCommentClient) LikeComment(ctx context.Context, req *repo.LikeCommentReq) (bool, error) {
	reply, err := r.contentClient.Comment.Like(ctx, &contentv1.LikeComment_Req{
		Id:    req.ID,
		Liked: req.Active,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return false, err
	}
	return reply.GetLiked(), nil
}

func (r *ContentCommentClient) ThankComment(ctx context.Context, req *repo.ThankCommentReq) (bool, error) {
	reply, err := r.contentClient.Comment.Thank(ctx, &contentv1.ThankComment_Req{
		Id:              req.ID,
		Thanked:         true,
		ClientRequestId: req.ClientRequestID,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return false, err
	}
	return reply.GetThanked(), nil
}

func (r *ContentCommentClient) loadCommentViewerActionStates(ctx context.Context, userID int64, commentIDs []int64) (map[int64]*model.CommentViewerActionState, error) {
	commentIDs = lo.Uniq(commentIDs)
	if userID <= 0 || len(commentIDs) == 0 {
		return map[int64]*model.CommentViewerActionState{}, nil
	}
	reply, err := r.contentClient.Comment.MapViewerActionStates(ctx, &contentv1.MapCommentViewerActionStates_Req{
		CommentIds: commentIDs,
		UserId:     userID,
	})
	if err != nil {
		return nil, err
	}
	states := make(map[int64]*model.CommentViewerActionState, len(reply.GetStates()))
	for commentID, state := range reply.GetStates() {
		if state == nil {
			continue
		}
		states[commentID] = &model.CommentViewerActionState{
			Liked:   state.GetLiked(),
			Thanked: state.GetThanked(),
		}
	}
	return states, nil
}

func (r *ContentCommentClient) loadCommentArticles(ctx context.Context, comments []*contentv1.PageComments_Resp_Comment, viewerUserID int64) (map[int64]*contentv1.Article, error) {
	articleIDs := lo.Uniq(lo.FilterMap(comments, func(item *contentv1.PageComments_Resp_Comment, _ int) (int64, bool) {
		if item == nil || item.GetArticleId() == 0 {
			return 0, false
		}
		return item.GetArticleId(), true
	}))
	articles := make(map[int64]*contentv1.Article, len(articleIDs))
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

func (r *ContentCommentClient) commentProfileIDs(item *contentv1.PageComments_Resp_Comment, article *contentv1.Article) []int64 {
	if item == nil {
		return nil
	}
	userIDs := make([]int64, 0, 2)
	if item.CreatedBy != nil {
		userIDs = append(userIDs, *item.CreatedBy)
	}
	if item.ReplyUserId != nil {
		userIDs = append(userIDs, *item.ReplyUserId)
	}
	return userIDs
}

func (r *ContentCommentClient) commentListItem(item *contentv1.PageComments_Resp_Comment, article *contentv1.Article, profiles map[int64]*model.AccountProfile, state *model.CommentViewerActionState) *model.CommentListItem {
	if item == nil {
		return nil
	}
	if state == nil {
		state = &model.CommentViewerActionState{}
	}
	out := &model.CommentListItem{
		ID:                item.GetId(),
		ArticleID:         item.GetArticleId(),
		Content:           item.GetContent(),
		ContentRender:     util.LuteEngine.MarkdownStr(fmt.Sprintf("comment_%d", item.GetId()), item.GetContent()),
		Level:             item.GetLevel(),
		ParentID:          item.ParentId,
		ReplyID:           item.ReplyId,
		Restriction:       r.restrictionConverter.FromContentProto(item.GetRestriction()),
		ThankCount:        item.GetThankCount(),
		LikeCount:         item.GetLikeCount(),
		ReplyCount:        item.GetReplyCount(),
		ViewerActionState: state,
		CreatedAt:         new(item.GetCreatedAt().AsTime()),
		UpdatedAt:         new(item.GetUpdatedAt().AsTime()),
	}
	if item.GetDeletedAt() != nil {
		out.DeletedAt = new(item.GetDeletedAt().AsTime())
	}
	out.CreatedBy = item.CreatedBy
	out.UpdatedBy = item.UpdatedBy
	if item.CreatedBy != nil {
		out.User = profiles[*item.CreatedBy]
	}
	if item.ReplyUserId != nil {
		out.ReplyUser = profiles[*item.ReplyUserId]
	}
	if article != nil {
		out.Article = &model.ArticleBrief{
			ID:    article.GetId(),
			Title: article.GetTitle(),
		}
	}
	return out
}

func (r *ContentCommentClient) commentDetail(item *contentv1.PageComments_Resp_Comment, article *contentv1.Article, profiles map[int64]*model.AccountProfile) *repo.CommentDetail {
	if item == nil {
		return nil
	}
	out := &repo.CommentDetail{
		ID:                item.GetId(),
		ArticleID:         item.GetArticleId(),
		Content:           item.GetContent(),
		ContentRender:     util.LuteEngine.MarkdownStr(fmt.Sprintf("comment_%d", item.GetId()), item.GetContent()),
		Level:             item.GetLevel(),
		ParentID:          item.ParentId,
		ReplyID:           item.ReplyId,
		Restriction:       r.restrictionConverter.FromContentProto(item.GetRestriction()),
		ThankCount:        item.GetThankCount(),
		LikeCount:         item.GetLikeCount(),
		ReplyCount:        item.GetReplyCount(),
		ViewerActionState: &model.CommentViewerActionState{},
		CreatedAt:         new(item.GetCreatedAt().AsTime()),
		UpdatedAt:         new(item.GetUpdatedAt().AsTime()),
	}
	if item.GetDeletedAt() != nil {
		out.DeletedAt = new(item.GetDeletedAt().AsTime())
	}
	out.CreatedBy = item.CreatedBy
	out.UpdatedBy = item.UpdatedBy
	if item.CreatedBy != nil {
		out.User = profiles[*item.CreatedBy]
	}
	if item.ReplyUserId != nil {
		out.ReplyUser = profiles[*item.ReplyUserId]
	}
	return out
}
