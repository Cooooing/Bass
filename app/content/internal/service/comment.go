package service

import (
	"common/pkg/apperror"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/content/v1"
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/usecase"
	"content/internal/enum"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CommentService struct {
	v1.UnimplementedContentCommentServiceServer

	commentUsecase *usecase.CommentUsecase
}

func (s *CommentService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentCommentServiceServer(gs, s)
}

func (s *CommentService) RegisterHttp(hs *http.Server) {}

func NewCommentService(
	commentUsecase *usecase.CommentUsecase,
) *CommentService {
	return &CommentService{
		commentUsecase: commentUsecase,
	}
}

func (s *CommentService) Create(ctx context.Context, req *v1.CreateComment_Request) (rsp *v1.CreateComment_Response, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	addResponse, err := s.commentUsecase.Add(ctx, &usecase.CommentAddReq{Comment: &model.Comment{
		ArticleID: req.ArticleId,
		Content:   req.Content,
		ReplyID:   req.ReplyId,
		CreatedBy: new(req.UserId),
		UpdatedBy: new(req.UserId),
	}})
	if err != nil {
		return nil, err
	}
	comment := addResponse.Comment
	reply := &v1.CreateComment_Response_Comment{
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
	return &v1.CreateComment_Response{
		Comment: reply,
	}, err
}

func (s *CommentService) Hide(ctx context.Context, req *v1.HideComment_Request) (*v1.HideComment_Response, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.commentUsecase.Hide(ctx, &usecase.CommentHideReq{CommentID: req.Id, UserID: req.UserId, Reason: req.Reason})
	return &v1.HideComment_Response{}, err
}

func (s *CommentService) Unhide(ctx context.Context, req *v1.UnhideComment_Request) (*v1.UnhideComment_Response, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.commentUsecase.Unhide(ctx, &usecase.CommentUnhideReq{CommentID: req.Id, UserID: req.UserId, Reason: req.Reason})
	return &v1.UnhideComment_Response{}, err
}

func (s *CommentService) Lock(ctx context.Context, req *v1.LockComment_Request) (*v1.LockComment_Response, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.commentUsecase.Lock(ctx, &usecase.CommentLockReq{CommentID: req.Id, UserID: req.UserId, Reason: req.Reason})
	return &v1.LockComment_Response{}, err
}

func (s *CommentService) Unlock(ctx context.Context, req *v1.UnlockComment_Request) (*v1.UnlockComment_Response, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.commentUsecase.Unlock(ctx, &usecase.CommentUnlockReq{CommentID: req.Id, UserID: req.UserId, Reason: req.Reason})
	return &v1.UnlockComment_Response{}, err
}

func (s *CommentService) List(ctx context.Context, req *v1.ListComments_Request) (*v1.ListComments_Response, error) {
	req.Query = util.OrDefault(req.Query, &v1.ListComments_Request_CommentQueryParams{})
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
	pageResponse, err := s.commentUsecase.Page(ctx, &usecase.CommentPageReq{
		Page:         &base.PageRequest{Page: 1, Size: 1000},
		CommentID:    req.Query.CommentId,
		ParentID:     req.Query.ParentId,
		ReplyID:      req.Query.ReplyId,
		ArticleID:    req.Query.ArticleId,
		CreatedBy:    req.Query.UserId,
		Restriction:  restriction,
		Restrictions: restrictions,
		Level:        req.Query.Level,
		Order:        dbOrder,
	})
	if err != nil {
		return nil, err
	}
	comments := pageResponse.Rows
	rows := make([]*v1.ListComments_Response_Comment, 0, len(comments))
	for _, comment := range comments {
		row := &v1.ListComments_Response_Comment{
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
	return &v1.ListComments_Response{
		Rows: rows,
	}, err
}

func (s *CommentService) Page(ctx context.Context, req *v1.PageComments_Request) (*v1.PageComments_Response, error) {
	req.Query = util.OrDefault(req.Query, &v1.PageComments_Request_CommentQueryParams{})
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
	pageResponse, err := s.commentUsecase.Page(ctx, &usecase.CommentPageReq{
		Page:         &base.PageRequest{Page: int64(req.GetPage().GetPage()), Size: int64(req.GetPage().GetSize())},
		CommentID:    req.Query.CommentId,
		ParentID:     req.Query.ParentId,
		ReplyID:      req.Query.ReplyId,
		ArticleID:    req.Query.ArticleId,
		CreatedBy:    req.Query.UserId,
		Restriction:  restriction,
		Restrictions: restrictions,
		Level:        req.Query.Level,
		Order:        dbOrder,
	})
	if err != nil {
		return nil, err
	}
	comments := pageResponse.Rows
	rows := make([]*v1.PageComments_Response_Comment, 0, len(comments))
	for _, comment := range comments {
		row := &v1.PageComments_Response_Comment{
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
	return &v1.PageComments_Response{
		Page: &common.PageResponse{Page: uint32(pageResponse.Page.Page), Size: uint32(pageResponse.Page.Size), Total: uint32(pageResponse.Page.Total)},
		Rows: rows,
	}, err
}

func (s *CommentService) ListReplyPreviews(ctx context.Context, req *v1.ListCommentReplyPreviews_Request) (*v1.ListCommentReplyPreviews_Response, error) {
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
	previewsResponse, err := s.commentUsecase.ListReplyPreviews(ctx, &usecase.CommentListReplyPreviewsReq{ArticleID: req.ArticleId, ParentIDs: req.ParentIds, LimitPerParent: req.LimitPerParent, Restriction: restriction, Restrictions: restrictions, Order: dbOrder})
	if err != nil {
		return nil, err
	}
	previews := previewsResponse.Rows
	rows := make([]*v1.ListCommentReplyPreviews_Response_CommentReplyPreview, 0, len(previews))
	for _, preview := range previews {
		row := &v1.ListCommentReplyPreviews_Response_CommentReplyPreview{
			ParentId: preview.ParentID,
			Rows:     make([]*v1.ListCommentReplyPreviews_Response_Comment, 0, len(preview.Rows)),
		}
		for _, comment := range preview.Rows {
			reply := &v1.ListCommentReplyPreviews_Response_Comment{
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
	return &v1.ListCommentReplyPreviews_Response{
		Rows: rows,
	}, nil
}

func (s *CommentService) MapViewerActionStates(ctx context.Context, req *v1.MapCommentViewerActionStates_Request) (*v1.MapCommentViewerActionStates_Response, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	statesResponse, err := s.commentUsecase.MapViewerActionStates(ctx, &usecase.CommentMapViewerActionStatesReq{CommentIDs: req.CommentIds, UserID: req.UserId})
	if err != nil {
		return nil, err
	}
	reply := &v1.MapCommentViewerActionStates_Response{
		States: make(map[int64]*v1.MapCommentViewerActionStates_Response_State, len(statesResponse.States)),
	}
	for commentID, state := range statesResponse.States {
		reply.States[commentID] = &v1.MapCommentViewerActionStates_Response_State{
			Liked:   state.Liked,
			Thanked: state.Thanked,
		}
	}
	return reply, nil
}

func (s *CommentService) MapArticleLastComments(ctx context.Context, req *v1.MapArticleLastComments_Request) (*v1.MapArticleLastComments_Response, error) {
	commentsResponse, err := s.commentUsecase.MapArticleLastComments(ctx, &usecase.CommentMapArticleLastCommentsReq{ArticleIDs: req.ArticleIds})
	if err != nil {
		return nil, err
	}
	comments := commentsResponse.Comments
	reply := &v1.MapArticleLastComments_Response{
		Comments: make(map[int64]*v1.MapArticleLastComments_Response_Comment, len(comments)),
	}
	for articleID, comment := range comments {
		reply.Comments[articleID] = &v1.MapArticleLastComments_Response_Comment{
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

func (s *CommentService) Like(ctx context.Context, req *v1.LikeComment_Request) (rsp *v1.LikeComment_Response, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	likeResponse, err := s.commentUsecase.Like(ctx, &usecase.CommentLikeReq{CommentID: req.Id, UserID: req.UserId, Active: req.Liked})
	return &v1.LikeComment_Response{Liked: likeResponse.Liked}, err
}

func (s *CommentService) Thank(ctx context.Context, req *v1.ThankComment_Request) (rsp *v1.ThankComment_Response, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	thankResponse, err := s.commentUsecase.Thank(ctx, &usecase.CommentThankReq{CommentID: req.Id, UserID: req.UserId, Active: req.Thanked})
	return &v1.ThankComment_Response{Thanked: thankResponse.Thanked}, err
}
