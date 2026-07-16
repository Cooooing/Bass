package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"common/proto/gen/common"
	"context"
)

type ContentCommentUsecase struct {
	contentCommentClient repo.ContentCommentClient
}

func NewContentCommentUsecase(contentCommentClient repo.ContentCommentClient) *ContentCommentUsecase {
	return &ContentCommentUsecase{contentCommentClient: contentCommentClient}
}

type CreateCommentReq struct {
	UserID    int64
	ArticleID int64
	Content   string
	ReplyID   int64
}

type CreateCommentResponse struct {
	Comment *repo.CommentDetail
}

func (u *ContentCommentUsecase) CreateComment(ctx context.Context, req *CreateCommentReq) (*CreateCommentResponse, error) {
	response, err := u.contentCommentClient.CreateComment(ctx, &repo.CreateCommentReq{UserID: req.UserID, ArticleID: req.ArticleID, Content: req.Content, ReplyID: req.ReplyID})
	if err != nil {
		return nil, err
	}
	return &CreateCommentResponse{Comment: response.Comment}, nil
}

type ListCommentsReq struct {
	UserID int64
	Page   *common.PageRequest
	Query  *bbscontentv1.ListComments_Request_CommentQuery
}

type ListCommentsResponse struct {
	Page *repo.PageResponse
	Rows []*repo.CommentListItem
}

func (u *ContentCommentUsecase) ListComments(ctx context.Context, req *ListCommentsReq) (*ListCommentsResponse, error) {
	if req == nil {
		req = &ListCommentsReq{}
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	query := &repo.CommentQuery{
		Restrictions: []int32{
			int32(bbscontentv1.ContentRestriction_CONTENT_RESTRICTION_NONE),
			int32(bbscontentv1.ContentRestriction_CONTENT_RESTRICTION_LOCKED),
		},
	}
	if req.Query != nil {
		query.CommentID = req.Query.CommentId
		query.ArticleID = req.Query.ArticleId
		query.ParentID = req.Query.ParentId
		query.ReplyID = req.Query.ReplyId
		query.Level = req.Query.Level
		query.UserID = req.Query.UserId
		if req.Query.Order != nil {
			value := int32(*req.Query.Order)
			query.Order = &value
		}
	}
	response, err := u.contentCommentClient.ListComments(ctx, &repo.ListCommentsReq{UserID: req.UserID, Page: page, Query: query})
	if err != nil {
		return nil, err
	}
	return &ListCommentsResponse{Page: response.Page, Rows: response.Rows}, nil
}

type ListCommentThreadsReq struct {
	UserID            int64
	Page              *common.PageRequest
	ArticleID         int64
	Order             *bbscontentv1.CommentOrder
	ReplyPreviewLimit *int32
}

type ListCommentThreadsResponse struct {
	Page *repo.PageResponse
	Rows []*repo.CommentThread
}

func (u *ContentCommentUsecase) ListCommentThreads(ctx context.Context, req *ListCommentThreadsReq) (*ListCommentThreadsResponse, error) {
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	var order *int32
	if req.Order != nil {
		value := int32(*req.Order)
		order = &value
	}
	response, err := u.contentCommentClient.ListCommentThreads(ctx, &repo.ListCommentThreadsReq{UserID: req.UserID, Page: page, ArticleID: req.ArticleID, Order: order, ReplyPreviewLimit: req.ReplyPreviewLimit})
	if err != nil {
		return nil, err
	}
	return &ListCommentThreadsResponse{Page: response.Page, Rows: response.Rows}, nil
}

type ListCommentRepliesReq struct {
	UserID    int64
	Page      *common.PageRequest
	ArticleID int64
	ParentID  int64
	Order     *bbscontentv1.CommentOrder
}

type ListCommentRepliesResponse struct {
	Page *repo.PageResponse
	Rows []*repo.CommentListItem
}

func (u *ContentCommentUsecase) ListCommentReplies(ctx context.Context, req *ListCommentRepliesReq) (*ListCommentRepliesResponse, error) {
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	var order *int32
	if req.Order != nil {
		value := int32(*req.Order)
		order = &value
	}
	response, err := u.contentCommentClient.ListCommentReplies(ctx, &repo.ListCommentRepliesReq{UserID: req.UserID, Page: page, ArticleID: req.ArticleID, ParentID: req.ParentID, Order: order})
	if err != nil {
		return nil, err
	}
	return &ListCommentRepliesResponse{Page: response.Page, Rows: response.Rows}, nil
}

type ListCommentTimelineReq struct {
	UserID    int64
	Page      *common.PageRequest
	ArticleID int64
	Order     *bbscontentv1.CommentOrder
}

type ListCommentTimelineResponse struct {
	Page *repo.PageResponse
	Rows []*repo.CommentListItem
}

func (u *ContentCommentUsecase) ListCommentTimeline(ctx context.Context, req *ListCommentTimelineReq) (*ListCommentTimelineResponse, error) {
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	var order *int32
	if req.Order != nil {
		value := int32(*req.Order)
		order = &value
	}
	response, err := u.contentCommentClient.ListCommentTimeline(ctx, &repo.ListCommentTimelineReq{UserID: req.UserID, Page: page, ArticleID: req.ArticleID, Order: order})
	if err != nil {
		return nil, err
	}
	return &ListCommentTimelineResponse{Page: response.Page, Rows: response.Rows}, nil
}

type LikeCommentReq struct {
	UserID int64
	ID     int64
	Active bool
}

type LikeCommentResponse struct {
	Liked bool
}

func (u *ContentCommentUsecase) LikeComment(ctx context.Context, req *LikeCommentReq) (*LikeCommentResponse, error) {
	response, err := u.contentCommentClient.LikeComment(ctx, &repo.LikeCommentReq{UserID: req.UserID, ID: req.ID, Active: req.Active})
	if err != nil {
		return nil, err
	}
	return &LikeCommentResponse{Liked: response.Liked}, nil
}

type ThankCommentReq struct {
	UserID int64
	ID     int64
	Active bool
}

type ThankCommentResponse struct {
	Thanked bool
}

func (u *ContentCommentUsecase) ThankComment(ctx context.Context, req *ThankCommentReq) (*ThankCommentResponse, error) {
	response, err := u.contentCommentClient.ThankComment(ctx, &repo.ThankCommentReq{UserID: req.UserID, ID: req.ID, Active: req.Active})
	if err != nil {
		return nil, err
	}
	return &ThankCommentResponse{Thanked: response.Thanked}, nil
}
