package service

import (
	"common/pkg/apperror"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors/v1"
	"common/proto/gen/common/v1"
	v1 "common/proto/gen/content/v1"
	"context"

	"content/internal/biz/model"
	"content/internal/biz/usecase"
	"content/internal/enum"

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

func (s *CommentService) RegisterHTTP(hs *http.Server) {
}

func NewCommentService(
	commentUsecase *usecase.CommentUsecase,
) *CommentService {
	return &CommentService{
		commentUsecase: commentUsecase,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, req *v1.CreateComment_Req) (rsp *v1.CreateComment_Resp, err error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err = access.Normalize("")
	if err != nil {
		return nil, err
	}
	addResp, err := s.commentUsecase.Add(ctx, &usecase.CommentAddReq{
		Access: access,
		Comment: &model.Comment{
			ArticleID: req.GetArticleId(),
			Content:   req.GetContent(),
			ReplyID:   req.ReplyId,
		},
	})
	if err != nil {
		return nil, err
	}
	comment := addResp
	reply := &v1.CreateComment_Resp_Comment{
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
	return &v1.CreateComment_Resp{
		Comment: reply,
	}, err
}

func (s *CommentService) HideComment(ctx context.Context, req *v1.HideComment_Req) (*v1.HideComment_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	err = s.commentUsecase.Hide(ctx, &usecase.CommentHideReq{
		CommentID: req.GetId(),
		Access:    access,
		Reason:    req.Reason,
	})
	return &v1.HideComment_Resp{}, err
}

func (s *CommentService) UnhideComment(ctx context.Context, req *v1.UnhideComment_Req) (*v1.UnhideComment_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	err = s.commentUsecase.Unhide(ctx, &usecase.CommentUnhideReq{
		CommentID: req.GetId(),
		Access:    access,
		Reason:    req.Reason,
	})
	return &v1.UnhideComment_Resp{}, err
}

func (s *CommentService) LockComment(ctx context.Context, req *v1.LockComment_Req) (*v1.LockComment_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	err = s.commentUsecase.Lock(ctx, &usecase.CommentLockReq{
		CommentID: req.GetId(),
		Access:    access,
		Reason:    req.Reason,
	})
	return &v1.LockComment_Resp{}, err
}

func (s *CommentService) UnlockComment(ctx context.Context, req *v1.UnlockComment_Req) (*v1.UnlockComment_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	err = s.commentUsecase.Unlock(ctx, &usecase.CommentUnlockReq{
		CommentID: req.GetId(),
		Access:    access,
		Reason:    req.Reason,
	})
	return &v1.UnlockComment_Resp{}, err
}

func (s *CommentService) ListComments(ctx context.Context, req *v1.ListComments_Req) (*v1.ListComments_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	req.Query = util.OrDefault(req.GetQuery(), &v1.ListComments_Req_CommentQueryParams{})
	var restriction *enum.ContentRestriction
	if req.Query.Restriction != nil {
		status, ok := enum.ContentRestrictionMap.ToEnum(req.GetQuery().GetRestriction())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restriction = new(status)
	}
	restrictions := make([]enum.ContentRestriction, 0, len(req.GetQuery().GetRestrictions()))
	for _, item := range req.GetQuery().GetRestrictions() {
		status, ok := enum.ContentRestrictionMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restrictions = append(restrictions, status)
	}
	var dbOrder *enum.CommentOrder
	if req.Query.Order != nil {
		order, ok := enum.CommentOrderMap.ToEnum(req.GetQuery().GetOrder())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		dbOrder = new(order)
	}
	pageResp, err := s.commentUsecase.Page(ctx, &usecase.CommentPageReq{
		Access: access,
		Page: &commonmodel.PageReq{
			Page: 1,
			Size: 1000,
		},
		Filter: &model.CommentFilter{
			CommentID:       req.GetQuery().CommentId,
			ParentID:        req.GetQuery().ParentId,
			ReplyID:         req.GetQuery().ReplyId,
			ArticleID:       req.GetQuery().ArticleId,
			CreatedBy:       req.GetQuery().UserId,
			Restriction:     restriction,
			Restrictions:    restrictions,
			Level:           req.GetQuery().Level,
			Order:           dbOrder,
			ExcludedUserIDs: req.GetQuery().GetExcludedUserIds(),
			ExcludedIDs:     req.GetQuery().GetExcludedCommentIds(),
		},
	})
	if err != nil {
		return nil, err
	}
	comments := pageResp.Rows
	rows := make([]*v1.ListComments_Resp_Comment, 0, len(comments))
	for _, comment := range comments {
		row := &v1.ListComments_Resp_Comment{
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
	return &v1.ListComments_Resp{
		Rows: rows,
	}, err
}

func (s *CommentService) PageComments(ctx context.Context, req *v1.PageComments_Req) (*v1.PageComments_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	req.Query = util.OrDefault(req.GetQuery(), &v1.PageComments_Req_CommentQueryParams{})
	var restriction *enum.ContentRestriction
	if req.Query.Restriction != nil {
		status, ok := enum.ContentRestrictionMap.ToEnum(req.GetQuery().GetRestriction())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restriction = new(status)
	}
	restrictions := make([]enum.ContentRestriction, 0, len(req.GetQuery().GetRestrictions()))
	for _, item := range req.GetQuery().GetRestrictions() {
		status, ok := enum.ContentRestrictionMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restrictions = append(restrictions, status)
	}
	var dbOrder *enum.CommentOrder
	if req.Query.Order != nil {
		order, ok := enum.CommentOrderMap.ToEnum(req.GetQuery().GetOrder())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		dbOrder = new(order)
	}
	pageResp, err := s.commentUsecase.Page(ctx, &usecase.CommentPageReq{
		Access: access,
		Page:   commonmodel.PageReqFromProto(req.GetPage()),
		Filter: &model.CommentFilter{
			CommentID:       req.GetQuery().CommentId,
			ParentID:        req.GetQuery().ParentId,
			ReplyID:         req.GetQuery().ReplyId,
			ArticleID:       req.GetQuery().ArticleId,
			CreatedBy:       req.GetQuery().UserId,
			Restriction:     restriction,
			Restrictions:    restrictions,
			Level:           req.GetQuery().Level,
			Order:           dbOrder,
			ExcludedUserIDs: req.GetQuery().GetExcludedUserIds(),
			ExcludedIDs:     req.GetQuery().GetExcludedCommentIds(),
		},
	})
	if err != nil {
		return nil, err
	}
	comments := pageResp.Rows
	rows := make([]*v1.PageComments_Resp_Comment, 0, len(comments))
	for _, comment := range comments {
		row := &v1.PageComments_Resp_Comment{
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
	return &v1.PageComments_Resp{
		Page: &common.PageResp{
			Page:  pageResp.Page.Page,
			Size:  pageResp.Page.Size,
			Total: pageResp.Page.Total,
		},
		Rows: rows,
	}, err
}

func (s *CommentService) ListCommentReplyPreviews(ctx context.Context, req *v1.ListCommentReplyPreviews_Req) (*v1.ListCommentReplyPreviews_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	var restriction *enum.ContentRestriction
	if req.Restriction != nil {
		status, ok := enum.ContentRestrictionMap.ToEnum(req.GetRestriction())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restriction = new(status)
	}
	restrictions := make([]enum.ContentRestriction, 0, len(req.GetRestrictions()))
	for _, item := range req.GetRestrictions() {
		status, ok := enum.ContentRestrictionMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
		}
		restrictions = append(restrictions, status)
	}
	var dbOrder *enum.CommentOrder
	if req.Order != nil {
		order, ok := enum.CommentOrderMap.ToEnum(req.GetOrder())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		dbOrder = new(order)
	}
	previewsResp, err := s.commentUsecase.ListReplyPreviews(ctx, &usecase.CommentListReplyPreviewsReq{
		Access:          access,
		ArticleID:       req.GetArticleId(),
		ParentIDs:       req.GetParentIds(),
		LimitPerParent:  req.GetLimitPerParent(),
		Restriction:     restriction,
		Restrictions:    restrictions,
		Order:           dbOrder,
		ExcludedUserIDs: req.GetExcludedUserIds(),
	})
	if err != nil {
		return nil, err
	}
	previews := previewsResp
	rows := make([]*v1.ListCommentReplyPreviews_Resp_CommentReplyPreview, 0, len(previews))
	for _, preview := range previews {
		row := &v1.ListCommentReplyPreviews_Resp_CommentReplyPreview{
			ParentId: preview.ParentID,
			Rows:     make([]*v1.ListCommentReplyPreviews_Resp_Comment, 0, len(preview.Rows)),
		}
		for _, comment := range preview.Rows {
			reply := &v1.ListCommentReplyPreviews_Resp_Comment{
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
	return &v1.ListCommentReplyPreviews_Resp{
		Rows: rows,
	}, nil
}

func (s *CommentService) MapCommentViewerActionStates(ctx context.Context, req *v1.MapCommentViewerActionStates_Req) (*v1.MapCommentViewerActionStates_Resp, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	statesResp, err := s.commentUsecase.MapViewerActionStates(ctx, &usecase.CommentMapViewerActionStatesReq{
		CommentIDs: req.GetCommentIds(),
		UserID:     req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.MapCommentViewerActionStates_Resp{
		States: make(map[int64]*v1.MapCommentViewerActionStates_Resp_State, len(statesResp)),
	}
	for commentID, state := range statesResp {
		reply.States[commentID] = &v1.MapCommentViewerActionStates_Resp_State{
			Liked:   state.Liked,
			Thanked: state.Thanked,
		}
	}
	return reply, nil
}

func (s *CommentService) MapArticleLastComments(ctx context.Context, req *v1.MapArticleLastComments_Req) (*v1.MapArticleLastComments_Resp, error) {
	commentsResp, err := s.commentUsecase.MapArticleLastComments(ctx, req.GetArticleIds(), req.GetExcludedUserIds())
	if err != nil {
		return nil, err
	}
	comments := commentsResp
	reply := &v1.MapArticleLastComments_Resp{
		Comments: make(map[int64]*v1.MapArticleLastComments_Resp_Comment, len(comments)),
	}
	for articleID, comment := range comments {
		reply.Comments[articleID] = &v1.MapArticleLastComments_Resp_Comment{
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

func (s *CommentService) LikeComment(ctx context.Context, req *v1.LikeComment_Req) (rsp *v1.LikeComment_Resp, err error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err = access.Normalize("")
	if err != nil {
		return nil, err
	}
	likeResp, err := s.commentUsecase.Like(ctx, &usecase.CommentLikeReq{
		CommentID: req.GetId(),
		Access:    access,
		Active:    req.GetLiked(),
	})
	return &v1.LikeComment_Resp{
		Liked: likeResp,
	}, err
}

func (s *CommentService) ThankComment(ctx context.Context, req *v1.ThankComment_Req) (rsp *v1.ThankComment_Resp, err error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err = access.Normalize("")
	if err != nil {
		return nil, err
	}
	thankResp, err := s.commentUsecase.Thank(ctx, &usecase.CommentThankReq{
		CommentID: req.GetId(),
		Access:    access,
		Active:    req.GetThanked(),
	})
	return &v1.ThankComment_Resp{
		Thanked: thankResp,
	}, err
}
