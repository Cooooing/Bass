package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbscontentv1enum "common/proto/gen/bbs/v1/content/enum"
	bbsuserv1enum "common/proto/gen/bbs/v1/user/enum"
	"common/proto/gen/common"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ContentCommentService struct {
	bbscontentv1.UnimplementedCommentServiceServer
	contentCommentUsecase *usecase.ContentCommentUsecase
}

func NewContentCommentService(contentCommentUsecase *usecase.ContentCommentUsecase) *ContentCommentService {
	return &ContentCommentService{contentCommentUsecase: contentCommentUsecase}
}

func (s *ContentCommentService) RegisterGrpc(gs *grpc.Server) {}

func (s *ContentCommentService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterCommentServiceHTTPServer(hs, s)
}

func (s *ContentCommentService) Create(ctx context.Context, req *bbscontentv1.CreateComment_Req) (*bbscontentv1.CreateComment_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentCommentUsecase.CreateComment(ctx, &usecase.CreateCommentReq{UserID: userID, ArticleID: req.GetArticleId(), Content: req.GetContent(), ReplyID: req.GetReplyId()})
	if err != nil {
		return nil, err
	}
	var comment *bbscontentv1.CreateComment_Resp_CommentDetail
	if resp != nil {
		comment = &bbscontentv1.CreateComment_Resp_CommentDetail{Id: resp.ID, ArticleId: resp.ArticleID, Content: resp.Content, ContentRender: resp.ContentRender, Level: resp.Level, ParentId: resp.ParentID, ReplyId: resp.ReplyID, ReplyCount: resp.ReplyCount, LikeCount: resp.LikeCount, ThankCount: resp.ThankCount, Restriction: bbscontentv1enum.ContentRestriction(resp.Restriction), DeletedAt: resp.DeletedAt, CreatedBy: resp.CreatedBy, UpdatedBy: resp.UpdatedBy, CreatedAt: resp.CreatedAt, UpdatedAt: resp.UpdatedAt}
		if resp.ViewerActionState != nil {
			comment.ViewerActionState = &bbscontentv1.CreateComment_Resp_CommentViewerActionState{Liked: resp.ViewerActionState.Liked, Thanked: resp.ViewerActionState.Thanked}
		}
		if resp.User != nil {
			comment.User = &bbscontentv1.CreateComment_Resp_AccountProfile{Id: resp.User.ID, Name: resp.User.Name, Nickname: resp.User.Nickname, Url: resp.User.URL, AvatarUrl: resp.User.AvatarURL, Introduction: resp.User.Introduction, Status: bbsuserv1enum.AccountStatus(resp.User.Status), Mbti: bbsuserv1enum.MBTI(resp.User.MBTI), FollowCount: resp.User.FollowCount, FollowerCount: resp.User.FollowerCount, CreatedAt: resp.User.CreatedAt, UpdatedAt: resp.User.UpdatedAt}
		}
		if resp.ReplyUser != nil {
			comment.ReplyUser = &bbscontentv1.CreateComment_Resp_AccountProfile{Id: resp.ReplyUser.ID, Name: resp.ReplyUser.Name, Nickname: resp.ReplyUser.Nickname, Url: resp.ReplyUser.URL, AvatarUrl: resp.ReplyUser.AvatarURL, Introduction: resp.ReplyUser.Introduction, Status: bbsuserv1enum.AccountStatus(resp.ReplyUser.Status), Mbti: bbsuserv1enum.MBTI(resp.ReplyUser.MBTI), FollowCount: resp.ReplyUser.FollowCount, FollowerCount: resp.ReplyUser.FollowerCount, CreatedAt: resp.ReplyUser.CreatedAt, UpdatedAt: resp.ReplyUser.UpdatedAt}
		}
	}
	return &bbscontentv1.CreateComment_Resp{Comment: comment}, nil
}

func (s *ContentCommentService) List(ctx context.Context, req *bbscontentv1.ListComments_Req) (*bbscontentv1.ListComments_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentCommentUsecase.ListComments(ctx, &usecase.ListCommentsReq{UserID: userID, Page: req.GetPage(), Query: req.GetQuery()})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{Page: resp.Page.Page, Size: resp.Page.Size, Total: resp.Page.Total}
	}
	rows := make([]*bbscontentv1.ListComments_Resp_CommentListItem, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		item := &bbscontentv1.ListComments_Resp_CommentListItem{Id: row.ID, ArticleId: row.ArticleID, Content: row.Content, ContentRender: row.ContentRender, Level: row.Level, ParentId: row.ParentID, ReplyId: row.ReplyID, ReplyCount: row.ReplyCount, LikeCount: row.LikeCount, ThankCount: row.ThankCount, Restriction: bbscontentv1enum.ContentRestriction(row.Restriction), DeletedAt: row.DeletedAt, CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
		if row.ViewerActionState != nil {
			item.ViewerActionState = &bbscontentv1.ListComments_Resp_CommentViewerActionState{Liked: row.ViewerActionState.Liked, Thanked: row.ViewerActionState.Thanked}
		}
		if row.User != nil {
			item.User = &bbscontentv1.ListComments_Resp_AccountProfile{Id: row.User.ID, Name: row.User.Name, Nickname: row.User.Nickname, Url: row.User.URL, AvatarUrl: row.User.AvatarURL, Introduction: row.User.Introduction, Status: bbsuserv1enum.AccountStatus(row.User.Status), Mbti: bbsuserv1enum.MBTI(row.User.MBTI), FollowCount: row.User.FollowCount, FollowerCount: row.User.FollowerCount, CreatedAt: row.User.CreatedAt, UpdatedAt: row.User.UpdatedAt}
		}
		if row.ReplyUser != nil {
			item.ReplyUser = &bbscontentv1.ListComments_Resp_AccountProfile{Id: row.ReplyUser.ID, Name: row.ReplyUser.Name, Nickname: row.ReplyUser.Nickname, Url: row.ReplyUser.URL, AvatarUrl: row.ReplyUser.AvatarURL, Introduction: row.ReplyUser.Introduction, Status: bbsuserv1enum.AccountStatus(row.ReplyUser.Status), Mbti: bbsuserv1enum.MBTI(row.ReplyUser.MBTI), FollowCount: row.ReplyUser.FollowCount, FollowerCount: row.ReplyUser.FollowerCount, CreatedAt: row.ReplyUser.CreatedAt, UpdatedAt: row.ReplyUser.UpdatedAt}
		}
		rows = append(rows, item)
	}
	return &bbscontentv1.ListComments_Resp{Page: page, Rows: rows}, nil
}

func (s *ContentCommentService) ListThreads(ctx context.Context, req *bbscontentv1.ListCommentThreads_Req) (*bbscontentv1.ListCommentThreads_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentCommentUsecase.ListCommentThreads(ctx, &usecase.ListCommentThreadsReq{UserID: userID, Page: req.GetPage(), ArticleID: req.GetArticleId(), Order: req.Order, ReplyPreviewLimit: req.ReplyPreviewLimit})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{Page: resp.Page.Page, Size: resp.Page.Size, Total: resp.Page.Total}
	}
	rows := make([]*bbscontentv1.ListCommentThreads_Resp_CommentThread, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		thread := &bbscontentv1.ListCommentThreads_Resp_CommentThread{ReplyCount: row.ReplyCount, HasMoreReplies: row.HasMoreReplies}
		if row.Root != nil {
			thread.Root = &bbscontentv1.ListCommentThreads_Resp_CommentListItem{Id: row.Root.ID, ArticleId: row.Root.ArticleID, Content: row.Root.Content, ContentRender: row.Root.ContentRender, Level: row.Root.Level, ParentId: row.Root.ParentID, ReplyId: row.Root.ReplyID, ReplyCount: row.Root.ReplyCount, LikeCount: row.Root.LikeCount, ThankCount: row.Root.ThankCount, Restriction: bbscontentv1enum.ContentRestriction(row.Root.Restriction), DeletedAt: row.Root.DeletedAt, CreatedBy: row.Root.CreatedBy, UpdatedBy: row.Root.UpdatedBy, CreatedAt: row.Root.CreatedAt, UpdatedAt: row.Root.UpdatedAt}
			if row.Root.ViewerActionState != nil {
				thread.Root.ViewerActionState = &bbscontentv1.ListCommentThreads_Resp_CommentViewerActionState{Liked: row.Root.ViewerActionState.Liked, Thanked: row.Root.ViewerActionState.Thanked}
			}
			if row.Root.User != nil {
				thread.Root.User = &bbscontentv1.ListCommentThreads_Resp_AccountProfile{Id: row.Root.User.ID, Name: row.Root.User.Name, Nickname: row.Root.User.Nickname, Url: row.Root.User.URL, AvatarUrl: row.Root.User.AvatarURL, Introduction: row.Root.User.Introduction, Status: bbsuserv1enum.AccountStatus(row.Root.User.Status), Mbti: bbsuserv1enum.MBTI(row.Root.User.MBTI), FollowCount: row.Root.User.FollowCount, FollowerCount: row.Root.User.FollowerCount, CreatedAt: row.Root.User.CreatedAt, UpdatedAt: row.Root.User.UpdatedAt}
			}
			if row.Root.ReplyUser != nil {
				thread.Root.ReplyUser = &bbscontentv1.ListCommentThreads_Resp_AccountProfile{Id: row.Root.ReplyUser.ID, Name: row.Root.ReplyUser.Name, Nickname: row.Root.ReplyUser.Nickname, Url: row.Root.ReplyUser.URL, AvatarUrl: row.Root.ReplyUser.AvatarURL, Introduction: row.Root.ReplyUser.Introduction, Status: bbsuserv1enum.AccountStatus(row.Root.ReplyUser.Status), Mbti: bbsuserv1enum.MBTI(row.Root.ReplyUser.MBTI), FollowCount: row.Root.ReplyUser.FollowCount, FollowerCount: row.Root.ReplyUser.FollowerCount, CreatedAt: row.Root.ReplyUser.CreatedAt, UpdatedAt: row.Root.ReplyUser.UpdatedAt}
			}
		}
		thread.PreviewReplies = make([]*bbscontentv1.ListCommentThreads_Resp_CommentListItem, 0, len(row.PreviewReplies))
		for _, preview := range row.PreviewReplies {
			if preview == nil {
				thread.PreviewReplies = append(thread.PreviewReplies, nil)
				continue
			}
			item := &bbscontentv1.ListCommentThreads_Resp_CommentListItem{Id: preview.ID, ArticleId: preview.ArticleID, Content: preview.Content, ContentRender: preview.ContentRender, Level: preview.Level, ParentId: preview.ParentID, ReplyId: preview.ReplyID, ReplyCount: preview.ReplyCount, LikeCount: preview.LikeCount, ThankCount: preview.ThankCount, Restriction: bbscontentv1enum.ContentRestriction(preview.Restriction), DeletedAt: preview.DeletedAt, CreatedBy: preview.CreatedBy, UpdatedBy: preview.UpdatedBy, CreatedAt: preview.CreatedAt, UpdatedAt: preview.UpdatedAt}
			if preview.ViewerActionState != nil {
				item.ViewerActionState = &bbscontentv1.ListCommentThreads_Resp_CommentViewerActionState{Liked: preview.ViewerActionState.Liked, Thanked: preview.ViewerActionState.Thanked}
			}
			if preview.User != nil {
				item.User = &bbscontentv1.ListCommentThreads_Resp_AccountProfile{Id: preview.User.ID, Name: preview.User.Name, Nickname: preview.User.Nickname, Url: preview.User.URL, AvatarUrl: preview.User.AvatarURL, Introduction: preview.User.Introduction, Status: bbsuserv1enum.AccountStatus(preview.User.Status), Mbti: bbsuserv1enum.MBTI(preview.User.MBTI), FollowCount: preview.User.FollowCount, FollowerCount: preview.User.FollowerCount, CreatedAt: preview.User.CreatedAt, UpdatedAt: preview.User.UpdatedAt}
			}
			if preview.ReplyUser != nil {
				item.ReplyUser = &bbscontentv1.ListCommentThreads_Resp_AccountProfile{Id: preview.ReplyUser.ID, Name: preview.ReplyUser.Name, Nickname: preview.ReplyUser.Nickname, Url: preview.ReplyUser.URL, AvatarUrl: preview.ReplyUser.AvatarURL, Introduction: preview.ReplyUser.Introduction, Status: bbsuserv1enum.AccountStatus(preview.ReplyUser.Status), Mbti: bbsuserv1enum.MBTI(preview.ReplyUser.MBTI), FollowCount: preview.ReplyUser.FollowCount, FollowerCount: preview.ReplyUser.FollowerCount, CreatedAt: preview.ReplyUser.CreatedAt, UpdatedAt: preview.ReplyUser.UpdatedAt}
			}
			thread.PreviewReplies = append(thread.PreviewReplies, item)
		}
		rows = append(rows, thread)
	}
	return &bbscontentv1.ListCommentThreads_Resp{Page: page, Rows: rows}, nil
}

func (s *ContentCommentService) ListReplies(ctx context.Context, req *bbscontentv1.ListCommentReplies_Req) (*bbscontentv1.ListCommentReplies_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentCommentUsecase.ListCommentReplies(ctx, &usecase.ListCommentRepliesReq{UserID: userID, Page: req.GetPage(), ArticleID: req.GetArticleId(), ParentID: req.GetParentId(), Order: req.Order})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{Page: resp.Page.Page, Size: resp.Page.Size, Total: resp.Page.Total}
	}
	rows := make([]*bbscontentv1.ListCommentReplies_Resp_CommentListItem, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		item := &bbscontentv1.ListCommentReplies_Resp_CommentListItem{Id: row.ID, ArticleId: row.ArticleID, Content: row.Content, ContentRender: row.ContentRender, Level: row.Level, ParentId: row.ParentID, ReplyId: row.ReplyID, ReplyCount: row.ReplyCount, LikeCount: row.LikeCount, ThankCount: row.ThankCount, Restriction: bbscontentv1enum.ContentRestriction(row.Restriction), DeletedAt: row.DeletedAt, CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
		if row.ViewerActionState != nil {
			item.ViewerActionState = &bbscontentv1.ListCommentReplies_Resp_CommentViewerActionState{Liked: row.ViewerActionState.Liked, Thanked: row.ViewerActionState.Thanked}
		}
		if row.User != nil {
			item.User = &bbscontentv1.ListCommentReplies_Resp_AccountProfile{Id: row.User.ID, Name: row.User.Name, Nickname: row.User.Nickname, Url: row.User.URL, AvatarUrl: row.User.AvatarURL, Introduction: row.User.Introduction, Status: bbsuserv1enum.AccountStatus(row.User.Status), Mbti: bbsuserv1enum.MBTI(row.User.MBTI), FollowCount: row.User.FollowCount, FollowerCount: row.User.FollowerCount, CreatedAt: row.User.CreatedAt, UpdatedAt: row.User.UpdatedAt}
		}
		if row.ReplyUser != nil {
			item.ReplyUser = &bbscontentv1.ListCommentReplies_Resp_AccountProfile{Id: row.ReplyUser.ID, Name: row.ReplyUser.Name, Nickname: row.ReplyUser.Nickname, Url: row.ReplyUser.URL, AvatarUrl: row.ReplyUser.AvatarURL, Introduction: row.ReplyUser.Introduction, Status: bbsuserv1enum.AccountStatus(row.ReplyUser.Status), Mbti: bbsuserv1enum.MBTI(row.ReplyUser.MBTI), FollowCount: row.ReplyUser.FollowCount, FollowerCount: row.ReplyUser.FollowerCount, CreatedAt: row.ReplyUser.CreatedAt, UpdatedAt: row.ReplyUser.UpdatedAt}
		}
		rows = append(rows, item)
	}
	return &bbscontentv1.ListCommentReplies_Resp{Page: page, Rows: rows}, nil
}

func (s *ContentCommentService) ListTimeline(ctx context.Context, req *bbscontentv1.ListCommentTimeline_Req) (*bbscontentv1.ListCommentTimeline_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentCommentUsecase.ListCommentTimeline(ctx, &usecase.ListCommentTimelineReq{UserID: userID, Page: req.GetPage(), ArticleID: req.GetArticleId(), Order: req.Order})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{Page: resp.Page.Page, Size: resp.Page.Size, Total: resp.Page.Total}
	}
	rows := make([]*bbscontentv1.ListCommentTimeline_Resp_CommentListItem, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		item := &bbscontentv1.ListCommentTimeline_Resp_CommentListItem{Id: row.ID, ArticleId: row.ArticleID, Content: row.Content, ContentRender: row.ContentRender, Level: row.Level, ParentId: row.ParentID, ReplyId: row.ReplyID, ReplyCount: row.ReplyCount, LikeCount: row.LikeCount, ThankCount: row.ThankCount, Restriction: bbscontentv1enum.ContentRestriction(row.Restriction), DeletedAt: row.DeletedAt, CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
		if row.ViewerActionState != nil {
			item.ViewerActionState = &bbscontentv1.ListCommentTimeline_Resp_CommentViewerActionState{Liked: row.ViewerActionState.Liked, Thanked: row.ViewerActionState.Thanked}
		}
		if row.User != nil {
			item.User = &bbscontentv1.ListCommentTimeline_Resp_AccountProfile{Id: row.User.ID, Name: row.User.Name, Nickname: row.User.Nickname, Url: row.User.URL, AvatarUrl: row.User.AvatarURL, Introduction: row.User.Introduction, Status: bbsuserv1enum.AccountStatus(row.User.Status), Mbti: bbsuserv1enum.MBTI(row.User.MBTI), FollowCount: row.User.FollowCount, FollowerCount: row.User.FollowerCount, CreatedAt: row.User.CreatedAt, UpdatedAt: row.User.UpdatedAt}
		}
		if row.ReplyUser != nil {
			item.ReplyUser = &bbscontentv1.ListCommentTimeline_Resp_AccountProfile{Id: row.ReplyUser.ID, Name: row.ReplyUser.Name, Nickname: row.ReplyUser.Nickname, Url: row.ReplyUser.URL, AvatarUrl: row.ReplyUser.AvatarURL, Introduction: row.ReplyUser.Introduction, Status: bbsuserv1enum.AccountStatus(row.ReplyUser.Status), Mbti: bbsuserv1enum.MBTI(row.ReplyUser.MBTI), FollowCount: row.ReplyUser.FollowCount, FollowerCount: row.ReplyUser.FollowerCount, CreatedAt: row.ReplyUser.CreatedAt, UpdatedAt: row.ReplyUser.UpdatedAt}
		}
		rows = append(rows, item)
	}
	return &bbscontentv1.ListCommentTimeline_Resp{Page: page, Rows: rows}, nil
}

func (s *ContentCommentService) Like(ctx context.Context, req *bbscontentv1.LikeComment_Req) (*bbscontentv1.LikeComment_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentCommentUsecase.LikeComment(ctx, &usecase.LikeCommentReq{UserID: userID, ID: req.GetId(), Active: req.GetActive()})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.LikeComment_Resp{Liked: resp}, nil
}

func (s *ContentCommentService) Thank(ctx context.Context, req *bbscontentv1.ThankComment_Req) (*bbscontentv1.ThankComment_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentCommentUsecase.ThankComment(ctx, &usecase.ThankCommentReq{UserID: userID, ID: req.GetId(), Active: req.GetActive()})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.ThankComment_Resp{Thanked: resp}, nil
}
