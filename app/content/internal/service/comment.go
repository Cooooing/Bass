package service

import (
	"common/pkg/apperror"
	"common/pkg/server"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/content/v1"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/biz/usecase"
	"content/internal/enum"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CommentService struct {
	v1.UnimplementedContentCommentServiceServer

	commentUsecase *usecase.CommentUsecase
}

func (s *CommentService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentCommentServiceServer(gs, s)
}

func NewCommentService(
	commentUsecase *usecase.CommentUsecase,
) *CommentService {
	return &CommentService{
		commentUsecase: commentUsecase,
	}
}

func (s *CommentService) Create(ctx context.Context, req *v1.CreateComment_Request) (rsp *v1.CreateComment_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	comment, err := s.commentUsecase.Add(ctx, &model.Comment{
		ArticleID: req.ArticleId,
		Content:   req.Content,
		ReplyID:   req.ReplyId,
		CreatedBy: new(req.UserId),
		UpdatedBy: new(req.UserId),
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.Comment{
		CreatedAt:   timestamppb.New(*comment.CreatedAt),
		UpdatedAt:   timestamppb.New(*comment.UpdatedAt),
		CreatedBy:   comment.CreatedBy,
		UpdatedBy:   comment.UpdatedBy,
		Id:          comment.ID,
		ArticleId:   comment.ArticleID,
		Content:     comment.Content,
		Level:       comment.Level,
		ParentId:    comment.ParentID,
		ReplyId:     comment.ReplyID,
		Restriction: enum.ContentRestrictionMap.MustToProto(comment.Restriction),
		ThankCount:  comment.ThankCount,
		LikeCount:   comment.LikeCount,
		ReplyCount:  comment.ReplyCount,
		ReplyUserId: comment.ReplyUserID,
	}
	if comment.DeletedAt != nil {
		reply.DeletedAt = timestamppb.New(*comment.DeletedAt)
	}
	return &v1.CreateComment_Reply{
		Comment: reply,
	}, err
}

func (s *CommentService) Hide(ctx context.Context, req *v1.HideComment_Request) (*v1.HideComment_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.commentUsecase.Hide(ctx, req.Id, req.UserId, req.Reason)
	return &v1.HideComment_Reply{}, err
}

func (s *CommentService) Unhide(ctx context.Context, req *v1.UnhideComment_Request) (*v1.UnhideComment_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.commentUsecase.Unhide(ctx, req.Id, req.UserId, req.Reason)
	return &v1.UnhideComment_Reply{}, err
}

func (s *CommentService) Lock(ctx context.Context, req *v1.LockComment_Request) (*v1.LockComment_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.commentUsecase.Lock(ctx, req.Id, req.UserId, req.Reason)
	return &v1.LockComment_Reply{}, err
}

func (s *CommentService) Unlock(ctx context.Context, req *v1.UnlockComment_Request) (*v1.UnlockComment_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.commentUsecase.Unlock(ctx, req.Id, req.UserId, req.Reason)
	return &v1.UnlockComment_Reply{}, err
}

func (s *CommentService) List(ctx context.Context, req *v1.ListComments_Request) (*v1.ListComments_Reply, error) {
	req.Query = util.OrDefault(req.Query, &v1.CommentQueryParams{})
	var restriction *enum.ContentRestriction
	if req.Query.Restriction != nil {
		status, ok := enum.ContentRestrictionMap.ToEnum(*req.Query.Restriction)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restriction = new(status)
	}
	restrictions := make([]enum.ContentRestriction, 0, len(req.Query.Restrictions))
	for _, item := range req.Query.Restrictions {
		status, ok := enum.ContentRestrictionMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restrictions = append(restrictions, status)
	}
	var dbOrder *enum.CommentOrder
	if req.Query.Order != nil {
		order, ok := enum.CommentOrderMap.ToEnum(*req.Query.Order)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		dbOrder = new(order)
	}
	_, comments, err := s.commentUsecase.Page(ctx, server.GetPageMax(), &repo.CommentGetReq{
		CommentId:    req.Query.CommentId,
		CommentIds:   nil,
		ParentId:     req.Query.ParentId,
		ReplyId:      req.Query.ReplyId,
		ArticleId:    req.Query.ArticleId,
		ArticleIds:   nil,
		CreatedBy:    req.Query.UserId,
		Restriction:  restriction,
		Restrictions: restrictions,
		Level:        req.Query.Level,
		Order:        dbOrder,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Comment, 0, len(comments))
	for _, comment := range comments {
		row := &v1.Comment{
			CreatedAt:   timestamppb.New(*comment.CreatedAt),
			UpdatedAt:   timestamppb.New(*comment.UpdatedAt),
			CreatedBy:   comment.CreatedBy,
			UpdatedBy:   comment.UpdatedBy,
			Id:          comment.ID,
			ArticleId:   comment.ArticleID,
			Content:     comment.Content,
			Level:       comment.Level,
			ParentId:    comment.ParentID,
			ReplyId:     comment.ReplyID,
			Restriction: enum.ContentRestrictionMap.MustToProto(comment.Restriction),
			ThankCount:  comment.ThankCount,
			LikeCount:   comment.LikeCount,
			ReplyCount:  comment.ReplyCount,
			ReplyUserId: comment.ReplyUserID,
		}
		if comment.DeletedAt != nil {
			row.DeletedAt = timestamppb.New(*comment.DeletedAt)
		}
		rows = append(rows, row)
	}
	return &v1.ListComments_Reply{
		Rows: rows,
	}, err
}

func (s *CommentService) Page(ctx context.Context, req *v1.PageComments_Request) (*v1.PageComments_Reply, error) {
	req.Query = util.OrDefault(req.Query, &v1.CommentQueryParams{})
	var restriction *enum.ContentRestriction
	if req.Query.Restriction != nil {
		status, ok := enum.ContentRestrictionMap.ToEnum(*req.Query.Restriction)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restriction = new(status)
	}
	restrictions := make([]enum.ContentRestriction, 0, len(req.Query.Restrictions))
	for _, item := range req.Query.Restrictions {
		status, ok := enum.ContentRestrictionMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restrictions = append(restrictions, status)
	}
	var dbOrder *enum.CommentOrder
	if req.Query.Order != nil {
		order, ok := enum.CommentOrderMap.ToEnum(*req.Query.Order)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		dbOrder = new(order)
	}
	page, comments, err := s.commentUsecase.Page(ctx, req.Page, &repo.CommentGetReq{
		CommentId:    req.Query.CommentId,
		CommentIds:   nil,
		ParentId:     req.Query.ParentId,
		ReplyId:      req.Query.ReplyId,
		ArticleId:    req.Query.ArticleId,
		ArticleIds:   nil,
		CreatedBy:    req.Query.UserId,
		Restriction:  restriction,
		Restrictions: restrictions,
		Level:        req.Query.Level,
		Order:        dbOrder,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Comment, 0, len(comments))
	for _, comment := range comments {
		row := &v1.Comment{
			CreatedAt:   timestamppb.New(*comment.CreatedAt),
			UpdatedAt:   timestamppb.New(*comment.UpdatedAt),
			CreatedBy:   comment.CreatedBy,
			UpdatedBy:   comment.UpdatedBy,
			Id:          comment.ID,
			ArticleId:   comment.ArticleID,
			Content:     comment.Content,
			Level:       comment.Level,
			ParentId:    comment.ParentID,
			ReplyId:     comment.ReplyID,
			Restriction: enum.ContentRestrictionMap.MustToProto(comment.Restriction),
			ThankCount:  comment.ThankCount,
			LikeCount:   comment.LikeCount,
			ReplyCount:  comment.ReplyCount,
			ReplyUserId: comment.ReplyUserID,
		}
		if comment.DeletedAt != nil {
			row.DeletedAt = timestamppb.New(*comment.DeletedAt)
		}
		rows = append(rows, row)
	}
	return &v1.PageComments_Reply{
		Page: page,
		Rows: rows,
	}, err
}

func (s *CommentService) ListReplyPreviews(ctx context.Context, req *v1.ListCommentReplyPreviews_Request) (*v1.ListCommentReplyPreviews_Reply, error) {
	var restriction *enum.ContentRestriction
	if req.Restriction != nil {
		status, ok := enum.ContentRestrictionMap.ToEnum(*req.Restriction)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restriction = new(status)
	}
	restrictions := make([]enum.ContentRestriction, 0, len(req.Restrictions))
	for _, item := range req.Restrictions {
		status, ok := enum.ContentRestrictionMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restrictions = append(restrictions, status)
	}
	var dbOrder *enum.CommentOrder
	if req.Order != nil {
		order, ok := enum.CommentOrderMap.ToEnum(*req.Order)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		dbOrder = new(order)
	}
	previews, err := s.commentUsecase.ListReplyPreviews(ctx, req.ArticleId, req.ParentIds, req.LimitPerParent, restriction, restrictions, dbOrder)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.CommentReplyPreview, 0, len(previews))
	for _, preview := range previews {
		row := &v1.CommentReplyPreview{
			ParentId: preview.ParentId,
			Rows:     make([]*v1.Comment, 0, len(preview.Rows)),
		}
		for _, comment := range preview.Rows {
			reply := &v1.Comment{
				CreatedAt:   timestamppb.New(*comment.CreatedAt),
				UpdatedAt:   timestamppb.New(*comment.UpdatedAt),
				CreatedBy:   comment.CreatedBy,
				UpdatedBy:   comment.UpdatedBy,
				Id:          comment.ID,
				ArticleId:   comment.ArticleID,
				Content:     comment.Content,
				Level:       comment.Level,
				ParentId:    comment.ParentID,
				ReplyId:     comment.ReplyID,
				Restriction: enum.ContentRestrictionMap.MustToProto(comment.Restriction),
				ThankCount:  comment.ThankCount,
				LikeCount:   comment.LikeCount,
				ReplyCount:  comment.ReplyCount,
				ReplyUserId: comment.ReplyUserID,
			}
			if comment.DeletedAt != nil {
				reply.DeletedAt = timestamppb.New(*comment.DeletedAt)
			}
			row.Rows = append(row.Rows, reply)
		}
		rows = append(rows, row)
	}
	return &v1.ListCommentReplyPreviews_Reply{
		Rows: rows,
	}, nil
}

func (s *CommentService) MapViewerActionStates(ctx context.Context, req *v1.MapCommentViewerActionStates_Request) (*v1.MapCommentViewerActionStates_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	states, err := s.commentUsecase.MapViewerActionStates(ctx, req.CommentIds, req.UserId)
	if err != nil {
		return nil, err
	}
	reply := &v1.MapCommentViewerActionStates_Reply{
		States: make(map[int64]*v1.MapCommentViewerActionStates_State, len(states)),
	}
	for commentID, state := range states {
		reply.States[commentID] = &v1.MapCommentViewerActionStates_State{
			Liked:   state.Liked,
			Thanked: state.Thanked,
		}
	}
	return reply, nil
}

func (s *CommentService) MapArticleLastComments(ctx context.Context, req *v1.MapArticleLastComments_Request) (*v1.MapArticleLastComments_Reply, error) {
	comments, err := s.commentUsecase.MapArticleLastComments(ctx, req.ArticleIds)
	if err != nil {
		return nil, err
	}
	reply := &v1.MapArticleLastComments_Reply{
		Comments: make(map[int64]*v1.Comment, len(comments)),
	}
	for articleID, comment := range comments {
		reply.Comments[articleID] = &v1.Comment{
			CreatedAt:   timestamppb.New(*comment.CreatedAt),
			UpdatedAt:   timestamppb.New(*comment.UpdatedAt),
			CreatedBy:   comment.CreatedBy,
			UpdatedBy:   comment.UpdatedBy,
			Id:          comment.ID,
			ArticleId:   comment.ArticleID,
			Content:     comment.Content,
			Level:       comment.Level,
			ParentId:    comment.ParentID,
			ReplyId:     comment.ReplyID,
			Restriction: enum.ContentRestrictionMap.MustToProto(comment.Restriction),
			ThankCount:  comment.ThankCount,
			LikeCount:   comment.LikeCount,
			ReplyCount:  comment.ReplyCount,
			ReplyUserId: comment.ReplyUserID,
		}
		if comment.DeletedAt != nil {
			reply.Comments[articleID].DeletedAt = timestamppb.New(*comment.DeletedAt)
		}
	}
	return reply, nil
}

func (s *CommentService) Like(ctx context.Context, req *v1.LikeComment_Request) (rsp *v1.LikeComment_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	liked, err := s.commentUsecase.Like(ctx, req.Id, req.UserId, req.Liked)
	return &v1.LikeComment_Reply{Liked: liked}, err
}

func (s *CommentService) Thank(ctx context.Context, req *v1.ThankComment_Request) (rsp *v1.ThankComment_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	thanked, err := s.commentUsecase.Thank(ctx, req.Id, req.UserId, req.Thanked)
	return &v1.ThankComment_Reply{Thanked: thanked}, err
}
