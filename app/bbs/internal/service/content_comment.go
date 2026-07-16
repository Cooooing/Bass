package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"common/proto/gen/common"
	"context"

	"github.com/go-kratos/kratos/v3/transport/http"
)

type ContentCommentService struct {
	bbscontentv1.UnimplementedCommentServiceServer
	contentCommentUsecase *usecase.ContentCommentUsecase
}

func NewContentCommentService(contentCommentUsecase *usecase.ContentCommentUsecase) *ContentCommentService {
	return &ContentCommentService{contentCommentUsecase: contentCommentUsecase}
}

func (s *ContentCommentService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterCommentServiceHTTPServer(hs, s)
}

func (s *ContentCommentService) Create(ctx context.Context, req *bbscontentv1.CreateComment_Request) (*bbscontentv1.CreateComment_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.contentCommentUsecase.CreateComment(ctx, &usecase.CreateCommentReq{UserID: userID, ArticleID: req.GetArticleId(), Content: req.GetContent(), ReplyID: req.GetReplyId()})
	if err != nil {
		return nil, err
	}
	var comment *bbscontentv1.CreateComment_Response_CommentDetail
	if response.Comment != nil {
		comment = &bbscontentv1.CreateComment_Response_CommentDetail{Id: response.Comment.ID, ArticleId: response.Comment.ArticleID, Content: response.Comment.Content, ContentRender: response.Comment.ContentRender, Level: response.Comment.Level, ParentId: response.Comment.ParentID, ReplyId: response.Comment.ReplyID, ReplyCount: response.Comment.ReplyCount, LikeCount: response.Comment.LikeCount, ThankCount: response.Comment.ThankCount, Restriction: bbscontentv1.ContentRestriction(response.Comment.Restriction), DeletedAt: response.Comment.DeletedAt, CreatedBy: response.Comment.CreatedBy, UpdatedBy: response.Comment.UpdatedBy, CreatedAt: response.Comment.CreatedAt, UpdatedAt: response.Comment.UpdatedAt}
		if response.Comment.ViewerActionState != nil {
			comment.ViewerActionState = &bbscontentv1.CreateComment_Response_CommentViewerActionState{Liked: response.Comment.ViewerActionState.Liked, Thanked: response.Comment.ViewerActionState.Thanked}
		}
		if response.Comment.User != nil {
			comment.User = &bbscontentv1.CreateComment_Response_AccountProfile{Id: response.Comment.User.ID, Name: response.Comment.User.Name, Nickname: response.Comment.User.Nickname, Url: response.Comment.User.URL, AvatarUrl: response.Comment.User.AvatarURL, Introduction: response.Comment.User.Introduction, Status: bbsuserv1.AccountStatus(response.Comment.User.Status), Mbti: bbsuserv1.MBTI(response.Comment.User.MBTI), FollowCount: response.Comment.User.FollowCount, FollowerCount: response.Comment.User.FollowerCount, CreatedAt: response.Comment.User.CreatedAt, UpdatedAt: response.Comment.User.UpdatedAt}
		}
		if response.Comment.ReplyUser != nil {
			comment.ReplyUser = &bbscontentv1.CreateComment_Response_AccountProfile{Id: response.Comment.ReplyUser.ID, Name: response.Comment.ReplyUser.Name, Nickname: response.Comment.ReplyUser.Nickname, Url: response.Comment.ReplyUser.URL, AvatarUrl: response.Comment.ReplyUser.AvatarURL, Introduction: response.Comment.ReplyUser.Introduction, Status: bbsuserv1.AccountStatus(response.Comment.ReplyUser.Status), Mbti: bbsuserv1.MBTI(response.Comment.ReplyUser.MBTI), FollowCount: response.Comment.ReplyUser.FollowCount, FollowerCount: response.Comment.ReplyUser.FollowerCount, CreatedAt: response.Comment.ReplyUser.CreatedAt, UpdatedAt: response.Comment.ReplyUser.UpdatedAt}
		}
	}
	return &bbscontentv1.CreateComment_Response{Comment: comment}, nil
}

func (s *ContentCommentService) List(ctx context.Context, req *bbscontentv1.ListComments_Request) (*bbscontentv1.ListComments_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.contentCommentUsecase.ListComments(ctx, &usecase.ListCommentsReq{UserID: userID, Page: req.GetPage(), Query: req.GetQuery()})
	if err != nil {
		return nil, err
	}
	var page *common.PageResponse
	if response.Page != nil {
		page = &common.PageResponse{Page: response.Page.Page, Size: response.Page.Size, Total: response.Page.Total}
	}
	rows := make([]*bbscontentv1.ListComments_Response_CommentListItem, 0, len(response.Rows))
	for _, row := range response.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		item := &bbscontentv1.ListComments_Response_CommentListItem{Id: row.ID, ArticleId: row.ArticleID, Content: row.Content, ContentRender: row.ContentRender, Level: row.Level, ParentId: row.ParentID, ReplyId: row.ReplyID, ReplyCount: row.ReplyCount, LikeCount: row.LikeCount, ThankCount: row.ThankCount, Restriction: bbscontentv1.ContentRestriction(row.Restriction), DeletedAt: row.DeletedAt, CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
		if row.ViewerActionState != nil {
			item.ViewerActionState = &bbscontentv1.ListComments_Response_CommentViewerActionState{Liked: row.ViewerActionState.Liked, Thanked: row.ViewerActionState.Thanked}
		}
		if row.User != nil {
			item.User = &bbscontentv1.ListComments_Response_AccountProfile{Id: row.User.ID, Name: row.User.Name, Nickname: row.User.Nickname, Url: row.User.URL, AvatarUrl: row.User.AvatarURL, Introduction: row.User.Introduction, Status: bbsuserv1.AccountStatus(row.User.Status), Mbti: bbsuserv1.MBTI(row.User.MBTI), FollowCount: row.User.FollowCount, FollowerCount: row.User.FollowerCount, CreatedAt: row.User.CreatedAt, UpdatedAt: row.User.UpdatedAt}
		}
		if row.ReplyUser != nil {
			item.ReplyUser = &bbscontentv1.ListComments_Response_AccountProfile{Id: row.ReplyUser.ID, Name: row.ReplyUser.Name, Nickname: row.ReplyUser.Nickname, Url: row.ReplyUser.URL, AvatarUrl: row.ReplyUser.AvatarURL, Introduction: row.ReplyUser.Introduction, Status: bbsuserv1.AccountStatus(row.ReplyUser.Status), Mbti: bbsuserv1.MBTI(row.ReplyUser.MBTI), FollowCount: row.ReplyUser.FollowCount, FollowerCount: row.ReplyUser.FollowerCount, CreatedAt: row.ReplyUser.CreatedAt, UpdatedAt: row.ReplyUser.UpdatedAt}
		}
		rows = append(rows, item)
	}
	return &bbscontentv1.ListComments_Response{Page: page, Rows: rows}, nil
}

func (s *ContentCommentService) ListThreads(ctx context.Context, req *bbscontentv1.ListCommentThreads_Request) (*bbscontentv1.ListCommentThreads_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.contentCommentUsecase.ListCommentThreads(ctx, &usecase.ListCommentThreadsReq{UserID: userID, Page: req.GetPage(), ArticleID: req.GetArticleId(), Order: req.Order, ReplyPreviewLimit: req.ReplyPreviewLimit})
	if err != nil {
		return nil, err
	}
	var page *common.PageResponse
	if response.Page != nil {
		page = &common.PageResponse{Page: response.Page.Page, Size: response.Page.Size, Total: response.Page.Total}
	}
	rows := make([]*bbscontentv1.ListCommentThreads_Response_CommentThread, 0, len(response.Rows))
	for _, row := range response.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		thread := &bbscontentv1.ListCommentThreads_Response_CommentThread{ReplyCount: row.ReplyCount, HasMoreReplies: row.HasMoreReplies}
		if row.Root != nil {
			thread.Root = &bbscontentv1.ListCommentThreads_Response_CommentListItem{Id: row.Root.ID, ArticleId: row.Root.ArticleID, Content: row.Root.Content, ContentRender: row.Root.ContentRender, Level: row.Root.Level, ParentId: row.Root.ParentID, ReplyId: row.Root.ReplyID, ReplyCount: row.Root.ReplyCount, LikeCount: row.Root.LikeCount, ThankCount: row.Root.ThankCount, Restriction: bbscontentv1.ContentRestriction(row.Root.Restriction), DeletedAt: row.Root.DeletedAt, CreatedBy: row.Root.CreatedBy, UpdatedBy: row.Root.UpdatedBy, CreatedAt: row.Root.CreatedAt, UpdatedAt: row.Root.UpdatedAt}
			if row.Root.ViewerActionState != nil {
				thread.Root.ViewerActionState = &bbscontentv1.ListCommentThreads_Response_CommentViewerActionState{Liked: row.Root.ViewerActionState.Liked, Thanked: row.Root.ViewerActionState.Thanked}
			}
			if row.Root.User != nil {
				thread.Root.User = &bbscontentv1.ListCommentThreads_Response_AccountProfile{Id: row.Root.User.ID, Name: row.Root.User.Name, Nickname: row.Root.User.Nickname, Url: row.Root.User.URL, AvatarUrl: row.Root.User.AvatarURL, Introduction: row.Root.User.Introduction, Status: bbsuserv1.AccountStatus(row.Root.User.Status), Mbti: bbsuserv1.MBTI(row.Root.User.MBTI), FollowCount: row.Root.User.FollowCount, FollowerCount: row.Root.User.FollowerCount, CreatedAt: row.Root.User.CreatedAt, UpdatedAt: row.Root.User.UpdatedAt}
			}
			if row.Root.ReplyUser != nil {
				thread.Root.ReplyUser = &bbscontentv1.ListCommentThreads_Response_AccountProfile{Id: row.Root.ReplyUser.ID, Name: row.Root.ReplyUser.Name, Nickname: row.Root.ReplyUser.Nickname, Url: row.Root.ReplyUser.URL, AvatarUrl: row.Root.ReplyUser.AvatarURL, Introduction: row.Root.ReplyUser.Introduction, Status: bbsuserv1.AccountStatus(row.Root.ReplyUser.Status), Mbti: bbsuserv1.MBTI(row.Root.ReplyUser.MBTI), FollowCount: row.Root.ReplyUser.FollowCount, FollowerCount: row.Root.ReplyUser.FollowerCount, CreatedAt: row.Root.ReplyUser.CreatedAt, UpdatedAt: row.Root.ReplyUser.UpdatedAt}
			}
		}
		thread.PreviewReplies = make([]*bbscontentv1.ListCommentThreads_Response_CommentListItem, 0, len(row.PreviewReplies))
		for _, preview := range row.PreviewReplies {
			if preview == nil {
				thread.PreviewReplies = append(thread.PreviewReplies, nil)
				continue
			}
			item := &bbscontentv1.ListCommentThreads_Response_CommentListItem{Id: preview.ID, ArticleId: preview.ArticleID, Content: preview.Content, ContentRender: preview.ContentRender, Level: preview.Level, ParentId: preview.ParentID, ReplyId: preview.ReplyID, ReplyCount: preview.ReplyCount, LikeCount: preview.LikeCount, ThankCount: preview.ThankCount, Restriction: bbscontentv1.ContentRestriction(preview.Restriction), DeletedAt: preview.DeletedAt, CreatedBy: preview.CreatedBy, UpdatedBy: preview.UpdatedBy, CreatedAt: preview.CreatedAt, UpdatedAt: preview.UpdatedAt}
			if preview.ViewerActionState != nil {
				item.ViewerActionState = &bbscontentv1.ListCommentThreads_Response_CommentViewerActionState{Liked: preview.ViewerActionState.Liked, Thanked: preview.ViewerActionState.Thanked}
			}
			if preview.User != nil {
				item.User = &bbscontentv1.ListCommentThreads_Response_AccountProfile{Id: preview.User.ID, Name: preview.User.Name, Nickname: preview.User.Nickname, Url: preview.User.URL, AvatarUrl: preview.User.AvatarURL, Introduction: preview.User.Introduction, Status: bbsuserv1.AccountStatus(preview.User.Status), Mbti: bbsuserv1.MBTI(preview.User.MBTI), FollowCount: preview.User.FollowCount, FollowerCount: preview.User.FollowerCount, CreatedAt: preview.User.CreatedAt, UpdatedAt: preview.User.UpdatedAt}
			}
			if preview.ReplyUser != nil {
				item.ReplyUser = &bbscontentv1.ListCommentThreads_Response_AccountProfile{Id: preview.ReplyUser.ID, Name: preview.ReplyUser.Name, Nickname: preview.ReplyUser.Nickname, Url: preview.ReplyUser.URL, AvatarUrl: preview.ReplyUser.AvatarURL, Introduction: preview.ReplyUser.Introduction, Status: bbsuserv1.AccountStatus(preview.ReplyUser.Status), Mbti: bbsuserv1.MBTI(preview.ReplyUser.MBTI), FollowCount: preview.ReplyUser.FollowCount, FollowerCount: preview.ReplyUser.FollowerCount, CreatedAt: preview.ReplyUser.CreatedAt, UpdatedAt: preview.ReplyUser.UpdatedAt}
			}
			thread.PreviewReplies = append(thread.PreviewReplies, item)
		}
		rows = append(rows, thread)
	}
	return &bbscontentv1.ListCommentThreads_Response{Page: page, Rows: rows}, nil
}

func (s *ContentCommentService) ListReplies(ctx context.Context, req *bbscontentv1.ListCommentReplies_Request) (*bbscontentv1.ListCommentReplies_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.contentCommentUsecase.ListCommentReplies(ctx, &usecase.ListCommentRepliesReq{UserID: userID, Page: req.GetPage(), ArticleID: req.GetArticleId(), ParentID: req.GetParentId(), Order: req.Order})
	if err != nil {
		return nil, err
	}
	var page *common.PageResponse
	if response.Page != nil {
		page = &common.PageResponse{Page: response.Page.Page, Size: response.Page.Size, Total: response.Page.Total}
	}
	rows := make([]*bbscontentv1.ListCommentReplies_Response_CommentListItem, 0, len(response.Rows))
	for _, row := range response.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		item := &bbscontentv1.ListCommentReplies_Response_CommentListItem{Id: row.ID, ArticleId: row.ArticleID, Content: row.Content, ContentRender: row.ContentRender, Level: row.Level, ParentId: row.ParentID, ReplyId: row.ReplyID, ReplyCount: row.ReplyCount, LikeCount: row.LikeCount, ThankCount: row.ThankCount, Restriction: bbscontentv1.ContentRestriction(row.Restriction), DeletedAt: row.DeletedAt, CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
		if row.ViewerActionState != nil {
			item.ViewerActionState = &bbscontentv1.ListCommentReplies_Response_CommentViewerActionState{Liked: row.ViewerActionState.Liked, Thanked: row.ViewerActionState.Thanked}
		}
		if row.User != nil {
			item.User = &bbscontentv1.ListCommentReplies_Response_AccountProfile{Id: row.User.ID, Name: row.User.Name, Nickname: row.User.Nickname, Url: row.User.URL, AvatarUrl: row.User.AvatarURL, Introduction: row.User.Introduction, Status: bbsuserv1.AccountStatus(row.User.Status), Mbti: bbsuserv1.MBTI(row.User.MBTI), FollowCount: row.User.FollowCount, FollowerCount: row.User.FollowerCount, CreatedAt: row.User.CreatedAt, UpdatedAt: row.User.UpdatedAt}
		}
		if row.ReplyUser != nil {
			item.ReplyUser = &bbscontentv1.ListCommentReplies_Response_AccountProfile{Id: row.ReplyUser.ID, Name: row.ReplyUser.Name, Nickname: row.ReplyUser.Nickname, Url: row.ReplyUser.URL, AvatarUrl: row.ReplyUser.AvatarURL, Introduction: row.ReplyUser.Introduction, Status: bbsuserv1.AccountStatus(row.ReplyUser.Status), Mbti: bbsuserv1.MBTI(row.ReplyUser.MBTI), FollowCount: row.ReplyUser.FollowCount, FollowerCount: row.ReplyUser.FollowerCount, CreatedAt: row.ReplyUser.CreatedAt, UpdatedAt: row.ReplyUser.UpdatedAt}
		}
		rows = append(rows, item)
	}
	return &bbscontentv1.ListCommentReplies_Response{Page: page, Rows: rows}, nil
}

func (s *ContentCommentService) ListTimeline(ctx context.Context, req *bbscontentv1.ListCommentTimeline_Request) (*bbscontentv1.ListCommentTimeline_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.contentCommentUsecase.ListCommentTimeline(ctx, &usecase.ListCommentTimelineReq{UserID: userID, Page: req.GetPage(), ArticleID: req.GetArticleId(), Order: req.Order})
	if err != nil {
		return nil, err
	}
	var page *common.PageResponse
	if response.Page != nil {
		page = &common.PageResponse{Page: response.Page.Page, Size: response.Page.Size, Total: response.Page.Total}
	}
	rows := make([]*bbscontentv1.ListCommentTimeline_Response_CommentListItem, 0, len(response.Rows))
	for _, row := range response.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		item := &bbscontentv1.ListCommentTimeline_Response_CommentListItem{Id: row.ID, ArticleId: row.ArticleID, Content: row.Content, ContentRender: row.ContentRender, Level: row.Level, ParentId: row.ParentID, ReplyId: row.ReplyID, ReplyCount: row.ReplyCount, LikeCount: row.LikeCount, ThankCount: row.ThankCount, Restriction: bbscontentv1.ContentRestriction(row.Restriction), DeletedAt: row.DeletedAt, CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
		if row.ViewerActionState != nil {
			item.ViewerActionState = &bbscontentv1.ListCommentTimeline_Response_CommentViewerActionState{Liked: row.ViewerActionState.Liked, Thanked: row.ViewerActionState.Thanked}
		}
		if row.User != nil {
			item.User = &bbscontentv1.ListCommentTimeline_Response_AccountProfile{Id: row.User.ID, Name: row.User.Name, Nickname: row.User.Nickname, Url: row.User.URL, AvatarUrl: row.User.AvatarURL, Introduction: row.User.Introduction, Status: bbsuserv1.AccountStatus(row.User.Status), Mbti: bbsuserv1.MBTI(row.User.MBTI), FollowCount: row.User.FollowCount, FollowerCount: row.User.FollowerCount, CreatedAt: row.User.CreatedAt, UpdatedAt: row.User.UpdatedAt}
		}
		if row.ReplyUser != nil {
			item.ReplyUser = &bbscontentv1.ListCommentTimeline_Response_AccountProfile{Id: row.ReplyUser.ID, Name: row.ReplyUser.Name, Nickname: row.ReplyUser.Nickname, Url: row.ReplyUser.URL, AvatarUrl: row.ReplyUser.AvatarURL, Introduction: row.ReplyUser.Introduction, Status: bbsuserv1.AccountStatus(row.ReplyUser.Status), Mbti: bbsuserv1.MBTI(row.ReplyUser.MBTI), FollowCount: row.ReplyUser.FollowCount, FollowerCount: row.ReplyUser.FollowerCount, CreatedAt: row.ReplyUser.CreatedAt, UpdatedAt: row.ReplyUser.UpdatedAt}
		}
		rows = append(rows, item)
	}
	return &bbscontentv1.ListCommentTimeline_Response{Page: page, Rows: rows}, nil
}

func (s *ContentCommentService) Like(ctx context.Context, req *bbscontentv1.LikeComment_Request) (*bbscontentv1.LikeComment_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.contentCommentUsecase.LikeComment(ctx, &usecase.LikeCommentReq{UserID: userID, ID: req.GetId(), Active: req.GetActive()})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.LikeComment_Response{Liked: response.Liked}, nil
}

func (s *ContentCommentService) Thank(ctx context.Context, req *bbscontentv1.ThankComment_Request) (*bbscontentv1.ThankComment_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.contentCommentUsecase.ThankComment(ctx, &usecase.ThankCommentReq{UserID: userID, ID: req.GetId(), Active: req.GetActive()})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.ThankComment_Response{Thanked: response.Thanked}, nil
}
