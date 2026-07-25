package service

import (
	"common/pkg/apperror"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/usecase"
	"content/internal/enum"
	"context"

	v1 "common/proto/gen/content/v1"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleService struct {
	v1.UnimplementedContentArticleServiceServer

	articleUsecase *usecase.ArticleUsecase
}

func (s *ArticleService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentArticleServiceServer(gs, s)
}

func (s *ArticleService) RegisterHttp(hs *http.Server) {
}

func NewArticleService(
	articleUsecase *usecase.ArticleUsecase,
) *ArticleService {
	return &ArticleService{
		articleUsecase: articleUsecase,
	}
}

func (s *ArticleService) Create(ctx context.Context, req *v1.CreateArticle_Req) (rsp *v1.CreateArticle_Resp, err error) {
	article := req.Article
	if article == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	dbType, ok := enum.ArticleTypeMap.ToEnum(article.Type)
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	var bountyPoints *int32
	switch dbType {
	case enum.ArticleTypeQA:
		if qa := article.GetQa(); qa != nil {
			bountyPoints = new(qa.GetBountyPoints())
		}
	}
	addResp, err := s.articleUsecase.Add(ctx, &usecase.ArticleAddReq{
		Article: &model.Article{
			Title:         article.Title,
			Content:       article.Content,
			RewardContent: article.RewardContent,
			RewardPoints:  article.RewardPoints,
			Type:          dbType,
			BountyPoints:  bountyPoints,
			Statement:     article.Statement,
			Commentable:   util.DerefOrDefault(article.Commentable, true),
			Anonymous:     util.DerefOrDefault(article.Anonymous, false),
			CreatedBy:     new(req.UserId),
			UpdatedBy:     new(req.UserId),
		},
	})
	if err != nil {
		return nil, err
	}
	save := addResp
	articleReply := &v1.CreateArticle_Resp_Article{
		CreatedAt:     timestamppb.New(*save.CreatedAt),
		UpdatedAt:     timestamppb.New(*save.UpdatedAt),
		CreatedBy:     save.CreatedBy,
		UpdatedBy:     save.UpdatedBy,
		Id:            save.ID,
		Title:         save.Title,
		Content:       save.Content,
		RewardContent: save.RewardContent,
		HasPostscript: save.HasPostscript,
		HasReward:     util.IsNotNil(save.RewardPoints),
		RewardPoints:  save.RewardPoints,
		PublishStatus: enum.ArticlePublishStatusMap.MustToProto(save.PublishStatus),
		Visibility:    enum.ArticleVisibilityMap.MustToProto(save.Visibility),
		Restriction:   enum.ContentRestrictionMap.MustToProto(save.Restriction),
		Type:          enum.ArticleTypeMap.MustToProto(save.Type),
		Statement:     save.Statement,
		Commentable:   save.Commentable,
		Anonymous:     save.Anonymous,
		ViewCount:     save.ViewCount,
		ThankCount:    save.ThankCount,
		LikeCount:     save.LikeCount,
		CollectCount:  save.CollectCount,
		WatchCount:    save.WatchCount,
		ReplyCount:    save.ReplyCount,
	}
	switch save.Type {
	case enum.ArticleTypeQA:
		articleReply.TypeParams = &v1.CreateArticle_Resp_Article_Qa{
			Qa: &v1.CreateArticle_Resp_Article_QA{
				BountyPoints:     save.BountyPoints,
				AcceptedAnswerId: save.AcceptedAnswerID,
			},
		}
	}
	if save.PublishedAt != nil {
		articleReply.PublishedAt = timestamppb.New(*save.PublishedAt)
	}
	if save.EditedAt != nil {
		articleReply.EditedAt = timestamppb.New(*save.EditedAt)
	}
	return &v1.CreateArticle_Resp{
		Article: articleReply,
	}, nil
}

func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticle_Req) (rsp *v1.PublishArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	visibility, ok := enum.ArticleVisibilityMap.ToEnum(req.Visibility)
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
	}
	err = s.articleUsecase.Publish(ctx, &usecase.ArticlePublishReq{
		ArticleID:  req.ArticleId,
		UserID:     req.UserId,
		Visibility: visibility,
	})
	return &v1.PublishArticle_Resp{}, err
}

func (s *ArticleService) AddPostscript(ctx context.Context, req *v1.AddPostscriptArticle_Req) (rsp *v1.AddPostscriptArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	addPostscriptResp, err := s.articleUsecase.AddPostscript(ctx, &usecase.ArticleAddPostscriptReq{
		ArticleID: req.ArticleId,
		Content:   req.Content,
		UserID:    req.UserId,
	})
	if err != nil {
		return nil, err
	}
	save := addPostscriptResp
	return &v1.AddPostscriptArticle_Resp{
		ArticlePostscript: &v1.AddPostscriptArticle_Resp_ArticlePostscript{
			CreatedAt:   timestamppb.New(*save.CreatedAt),
			UpdatedAt:   timestamppb.New(*save.UpdatedAt),
			CreatedBy:   save.CreatedBy,
			UpdatedBy:   save.UpdatedBy,
			Id:          save.ID,
			ArticleId:   save.ArticleID,
			Content:     save.Content,
			Restriction: enum.ContentRestrictionMap.MustToProto(save.Restriction),
		},
	}, err
}

func (s *ArticleService) ListPostscripts(ctx context.Context, req *v1.ListArticlePostscripts_Req) (rsp *v1.ListArticlePostscripts_Resp, err error) {
	listPostscriptsResp, err := s.articleUsecase.ListPostscripts(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	rows := listPostscriptsResp
	reply := make([]*v1.ListArticlePostscripts_Resp_ArticlePostscript, 0, len(rows))
	for _, item := range rows {
		reply = append(reply, &v1.ListArticlePostscripts_Resp_ArticlePostscript{
			CreatedAt:   timestamppb.New(*item.CreatedAt),
			UpdatedAt:   timestamppb.New(*item.UpdatedAt),
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			Id:          item.ID,
			ArticleId:   item.ArticleID,
			Content:     item.Content,
			Restriction: enum.ContentRestrictionMap.MustToProto(item.Restriction),
		})
	}
	return &v1.ListArticlePostscripts_Resp{
		Rows: reply,
	}, nil
}

func (s *ArticleService) Update(ctx context.Context, req *v1.UpdateArticle_Req) (rsp *v1.UpdateArticle_Resp, err error) {
	article := req.Article
	if article == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	dbType, ok := enum.ArticleTypeMap.ToEnum(article.Type)
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	var bountyPoints *int32
	switch dbType {
	case enum.ArticleTypeQA:
		if qa := article.GetQa(); qa != nil {
			bountyPoints = new(qa.GetBountyPoints())
		}
	}
	updateResp, err := s.articleUsecase.Update(ctx, &model.Article{
		ID:            req.ArticleId,
		Title:         article.Title,
		Content:       article.Content,
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Type:          dbType,
		BountyPoints:  bountyPoints,
		Statement:     article.Statement,
		Commentable:   util.DerefOrDefault(article.Commentable, true),
		Anonymous:     util.DerefOrDefault(article.Anonymous, false),
		UpdatedBy:     new(req.UserId),
	})
	if err != nil {
		return nil, err
	}
	update := updateResp
	articleReply := &v1.UpdateArticle_Resp_Article{
		CreatedAt:     timestamppb.New(*update.CreatedAt),
		UpdatedAt:     timestamppb.New(*update.UpdatedAt),
		CreatedBy:     update.CreatedBy,
		UpdatedBy:     update.UpdatedBy,
		Id:            update.ID,
		Title:         update.Title,
		Content:       update.Content,
		RewardContent: update.RewardContent,
		HasPostscript: update.HasPostscript,
		HasReward:     util.IsNotNil(update.RewardPoints),
		RewardPoints:  update.RewardPoints,
		PublishStatus: enum.ArticlePublishStatusMap.MustToProto(update.PublishStatus),
		Visibility:    enum.ArticleVisibilityMap.MustToProto(update.Visibility),
		Restriction:   enum.ContentRestrictionMap.MustToProto(update.Restriction),
		Type:          enum.ArticleTypeMap.MustToProto(update.Type),
		Statement:     update.Statement,
		Commentable:   update.Commentable,
		Anonymous:     update.Anonymous,
		ViewCount:     update.ViewCount,
		ThankCount:    update.ThankCount,
		LikeCount:     update.LikeCount,
		CollectCount:  update.CollectCount,
		WatchCount:    update.WatchCount,
		ReplyCount:    update.ReplyCount,
	}
	switch update.Type {
	case enum.ArticleTypeQA:
		articleReply.TypeParams = &v1.UpdateArticle_Resp_Article_Qa{
			Qa: &v1.UpdateArticle_Resp_Article_QA{
				BountyPoints:     update.BountyPoints,
				AcceptedAnswerId: update.AcceptedAnswerID,
			},
		}
	}
	if update.PublishedAt != nil {
		articleReply.PublishedAt = timestamppb.New(*update.PublishedAt)
	}
	if update.EditedAt != nil {
		articleReply.EditedAt = timestamppb.New(*update.EditedAt)
	}
	return &v1.UpdateArticle_Resp{
		Article: articleReply,
	}, nil
}

func (s *ArticleService) DiscardDraft(ctx context.Context, req *v1.DiscardDraftArticle_Req) (rsp *v1.DiscardDraftArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.DiscardDraft(ctx, &usecase.ArticleDiscardDraftReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
	})
	return &v1.DiscardDraftArticle_Resp{}, err
}

func (s *ArticleService) MakePrivate(ctx context.Context, req *v1.MakePrivateArticle_Req) (*v1.MakePrivateArticle_Resp, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.MakePrivate(ctx, &usecase.ArticleMakePrivateReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
	})
	return &v1.MakePrivateArticle_Resp{}, err
}

func (s *ArticleService) MakePublic(ctx context.Context, req *v1.MakePublicArticle_Req) (*v1.MakePublicArticle_Resp, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.MakePublic(ctx, &usecase.ArticleMakePublicReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
	})
	return &v1.MakePublicArticle_Resp{}, err
}

func (s *ArticleService) Archive(ctx context.Context, req *v1.ArchiveArticle_Req) (*v1.ArchiveArticle_Resp, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.Archive(ctx, &usecase.ArticleArchiveReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Reason:    req.Reason,
	})
	return &v1.ArchiveArticle_Resp{}, err
}

func (s *ArticleService) Unarchive(ctx context.Context, req *v1.UnarchiveArticle_Req) (*v1.UnarchiveArticle_Resp, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.Unarchive(ctx, &usecase.ArticleUnarchiveReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Reason:    req.Reason,
	})
	return &v1.UnarchiveArticle_Resp{}, err
}

func (s *ArticleService) Hide(ctx context.Context, req *v1.HideArticle_Req) (rsp *v1.HideArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.Hide(ctx, &usecase.ArticleHideReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Reason:    req.Reason,
	})
	return &v1.HideArticle_Resp{}, err
}

func (s *ArticleService) Unhide(ctx context.Context, req *v1.UnhideArticle_Req) (*v1.UnhideArticle_Resp, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.Unhide(ctx, &usecase.ArticleUnhideReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Reason:    req.Reason,
	})
	return &v1.UnhideArticle_Resp{}, err
}

func (s *ArticleService) Lock(ctx context.Context, req *v1.LockArticle_Req) (rsp *v1.LockArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.Lock(ctx, &usecase.ArticleLockReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Reason:    req.Reason,
	})
	return &v1.LockArticle_Resp{}, err
}

func (s *ArticleService) Unlock(ctx context.Context, req *v1.UnlockArticle_Req) (rsp *v1.UnlockArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.Unlock(ctx, &usecase.ArticleUnlockReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Reason:    req.Reason,
	})
	return &v1.UnlockArticle_Resp{}, err
}

func (s *ArticleService) List(ctx context.Context, req *v1.ListArticles_Req) (rsp *v1.ListArticles_Resp, err error) {
	req.Query = util.OrDefault(req.Query, &v1.ListArticles_Req_ArticleQueryParams{})
	var publishStatus *enum.ArticlePublishStatus
	if req.Query.PublishStatus != nil {
		status, ok := enum.ArticlePublishStatusMap.ToEnum(*req.Query.PublishStatus)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		publishStatus = new(status)
	}
	publishStatuses := make([]enum.ArticlePublishStatus, 0, len(req.Query.PublishStatuses))
	for _, item := range req.Query.PublishStatuses {
		status, ok := enum.ArticlePublishStatusMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		publishStatuses = append(publishStatuses, status)
	}
	var visibility *enum.ArticleVisibility
	if req.Query.Visibility != nil {
		item, ok := enum.ArticleVisibilityMap.ToEnum(*req.Query.Visibility)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		visibility = new(item)
	}
	visibilities := make([]enum.ArticleVisibility, 0, len(req.Query.Visibilities))
	for _, item := range req.Query.Visibilities {
		visibility, ok := enum.ArticleVisibilityMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		visibilities = append(visibilities, visibility)
	}
	var restriction *enum.ContentRestriction
	if req.Query.Restriction != nil {
		status, ok := enum.ContentRestrictionMap.ToEnum(*req.Query.Restriction)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		restriction = new(status)
	}
	restrictions := make([]enum.ContentRestriction, 0, len(req.Query.Restrictions))
	for _, item := range req.Query.Restrictions {
		status, ok := enum.ContentRestrictionMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		restrictions = append(restrictions, status)
	}
	var dbType *enum.ArticleType
	if req.Query.Type != nil {
		articleType, ok := enum.ArticleTypeMap.ToEnum(*req.Query.Type)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
		}
		dbType = new(articleType)
	}
	var dbOrder *enum.ArticleOrder
	if req.Query.Order != nil {
		order, ok := enum.ArticleOrderMap.ToEnum(*req.Query.Order)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		dbOrder = new(order)
	}

	query := &usecase.ArticlePageReq{
		TagID:           req.Query.TagId,
		DomainID:        req.Query.DomainId,
		PublishStatus:   publishStatus,
		PublishStatuses: publishStatuses,
		Visibility:      visibility,
		Visibilities:    visibilities,
		Restriction:     restriction,
		Restrictions:    restrictions,
		AuthorID:        req.Query.AuthorId,
		Order:           dbOrder,
		Type:            dbType,
		Keyword:         req.Query.Keyword,
	}
	query.Page = &base.PageRequest{
		Page: 1,
		Size: 1000,
	}
	pageResp, err := s.articleUsecase.Page(ctx, query)
	if err != nil {
		return nil, err
	}
	reply := pageResp.Rows
	rows := make([]*v1.ListArticles_Resp_Article, 0, len(reply))
	for _, item := range reply {
		row := &v1.ListArticles_Resp_Article{
			CreatedAt:     timestamppb.New(*item.CreatedAt),
			UpdatedAt:     timestamppb.New(*item.UpdatedAt),
			CreatedBy:     item.CreatedBy,
			UpdatedBy:     item.UpdatedBy,
			Id:            item.ID,
			Title:         item.Title,
			Content:       item.Content,
			RewardContent: item.RewardContent,
			HasPostscript: item.HasPostscript,
			HasReward:     util.IsNotNil(item.RewardPoints),
			RewardPoints:  item.RewardPoints,
			PublishStatus: enum.ArticlePublishStatusMap.MustToProto(item.PublishStatus),
			Visibility:    enum.ArticleVisibilityMap.MustToProto(item.Visibility),
			Restriction:   enum.ContentRestrictionMap.MustToProto(item.Restriction),
			Type:          enum.ArticleTypeMap.MustToProto(item.Type),
			Statement:     item.Statement,
			Commentable:   item.Commentable,
			Anonymous:     item.Anonymous,
			ViewCount:     item.ViewCount,
			ThankCount:    item.ThankCount,
			LikeCount:     item.LikeCount,
			CollectCount:  item.CollectCount,
			WatchCount:    item.WatchCount,
			ReplyCount:    item.ReplyCount,
		}
		switch item.Type {
		case enum.ArticleTypeQA:
			row.TypeParams = &v1.ListArticles_Resp_Article_Qa{
				Qa: &v1.ListArticles_Resp_Article_QA{
					BountyPoints:     item.BountyPoints,
					AcceptedAnswerId: item.AcceptedAnswerID,
				},
			}
		}
		if item.PublishedAt != nil {
			row.PublishedAt = timestamppb.New(*item.PublishedAt)
		}
		if item.EditedAt != nil {
			row.EditedAt = timestamppb.New(*item.EditedAt)
		}
		rows = append(rows, row)
	}
	return &v1.ListArticles_Resp{
		Rows: rows,
	}, err
}

func (s *ArticleService) Page(ctx context.Context, req *v1.PageArticles_Req) (rsp *v1.PageArticles_Resp, err error) {
	req.Query = util.OrDefault(req.Query, &v1.PageArticles_Req_ArticleQueryParams{})
	var publishStatus *enum.ArticlePublishStatus
	if req.Query.PublishStatus != nil {
		status, ok := enum.ArticlePublishStatusMap.ToEnum(*req.Query.PublishStatus)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		publishStatus = new(status)
	}
	publishStatuses := make([]enum.ArticlePublishStatus, 0, len(req.Query.PublishStatuses))
	for _, item := range req.Query.PublishStatuses {
		status, ok := enum.ArticlePublishStatusMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		publishStatuses = append(publishStatuses, status)
	}
	var visibility *enum.ArticleVisibility
	if req.Query.Visibility != nil {
		item, ok := enum.ArticleVisibilityMap.ToEnum(*req.Query.Visibility)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		visibility = new(item)
	}
	visibilities := make([]enum.ArticleVisibility, 0, len(req.Query.Visibilities))
	for _, item := range req.Query.Visibilities {
		visibility, ok := enum.ArticleVisibilityMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		visibilities = append(visibilities, visibility)
	}
	var restriction *enum.ContentRestriction
	if req.Query.Restriction != nil {
		status, ok := enum.ContentRestrictionMap.ToEnum(*req.Query.Restriction)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		restriction = new(status)
	}
	restrictions := make([]enum.ContentRestriction, 0, len(req.Query.Restrictions))
	for _, item := range req.Query.Restrictions {
		status, ok := enum.ContentRestrictionMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		restrictions = append(restrictions, status)
	}
	var dbType *enum.ArticleType
	if req.Query.Type != nil {
		articleType, ok := enum.ArticleTypeMap.ToEnum(*req.Query.Type)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
		}
		dbType = new(articleType)
	}
	var dbOrder *enum.ArticleOrder
	if req.Query.Order != nil {
		order, ok := enum.ArticleOrderMap.ToEnum(*req.Query.Order)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		dbOrder = new(order)
	}

	query := &usecase.ArticlePageReq{
		TagID:           req.Query.TagId,
		DomainID:        req.Query.DomainId,
		PublishStatus:   publishStatus,
		PublishStatuses: publishStatuses,
		Visibility:      visibility,
		Visibilities:    visibilities,
		Restriction:     restriction,
		Restrictions:    restrictions,
		AuthorID:        req.Query.AuthorId,
		Order:           dbOrder,
		Type:            dbType,
		Keyword:         req.Query.Keyword,
	}
	query.Page = &base.PageRequest{
		Page: int64(req.GetPage().GetPage()),
		Size: int64(req.GetPage().GetSize()),
	}
	pageResp, err := s.articleUsecase.Page(ctx, query)
	if err != nil {
		return nil, err
	}
	reply := pageResp.Rows
	page := pageResp.Page
	rows := make([]*v1.PageArticles_Resp_Article, 0, len(reply))
	for _, item := range reply {
		row := &v1.PageArticles_Resp_Article{
			CreatedAt:     timestamppb.New(*item.CreatedAt),
			UpdatedAt:     timestamppb.New(*item.UpdatedAt),
			CreatedBy:     item.CreatedBy,
			UpdatedBy:     item.UpdatedBy,
			Id:            item.ID,
			Title:         item.Title,
			Content:       item.Content,
			RewardContent: item.RewardContent,
			HasPostscript: item.HasPostscript,
			HasReward:     util.IsNotNil(item.RewardPoints),
			RewardPoints:  item.RewardPoints,
			PublishStatus: enum.ArticlePublishStatusMap.MustToProto(item.PublishStatus),
			Visibility:    enum.ArticleVisibilityMap.MustToProto(item.Visibility),
			Restriction:   enum.ContentRestrictionMap.MustToProto(item.Restriction),
			Type:          enum.ArticleTypeMap.MustToProto(item.Type),
			Statement:     item.Statement,
			Commentable:   item.Commentable,
			Anonymous:     item.Anonymous,
			ViewCount:     item.ViewCount,
			ThankCount:    item.ThankCount,
			LikeCount:     item.LikeCount,
			CollectCount:  item.CollectCount,
			WatchCount:    item.WatchCount,
			ReplyCount:    item.ReplyCount,
		}
		switch item.Type {
		case enum.ArticleTypeQA:
			row.TypeParams = &v1.PageArticles_Resp_Article_Qa{
				Qa: &v1.PageArticles_Resp_Article_QA{
					BountyPoints:     item.BountyPoints,
					AcceptedAnswerId: item.AcceptedAnswerID,
				},
			}
		}
		if item.PublishedAt != nil {
			row.PublishedAt = timestamppb.New(*item.PublishedAt)
		}
		if item.EditedAt != nil {
			row.EditedAt = timestamppb.New(*item.EditedAt)
		}
		rows = append(rows, row)
	}
	return &v1.PageArticles_Resp{
		Page: &common.PageResp{
			Page:  uint32(page.Page),
			Size:  uint32(page.Size),
			Total: uint32(page.Total),
		},
		Rows: rows,
	}, err
}

func (s *ArticleService) Get(ctx context.Context, req *v1.GetArticle_Req) (rsp *v1.GetArticle_Resp, err error) {
	getResp, err := s.articleUsecase.Get(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	one := getResp
	article := &v1.GetArticle_Resp_Article{
		CreatedAt:     timestamppb.New(*one.CreatedAt),
		UpdatedAt:     timestamppb.New(*one.UpdatedAt),
		CreatedBy:     one.CreatedBy,
		UpdatedBy:     one.UpdatedBy,
		Id:            one.ID,
		Title:         one.Title,
		Content:       one.Content,
		RewardContent: one.RewardContent,
		HasPostscript: one.HasPostscript,
		HasReward:     util.IsNotNil(one.RewardPoints),
		RewardPoints:  one.RewardPoints,
		PublishStatus: enum.ArticlePublishStatusMap.MustToProto(one.PublishStatus),
		Visibility:    enum.ArticleVisibilityMap.MustToProto(one.Visibility),
		Restriction:   enum.ContentRestrictionMap.MustToProto(one.Restriction),
		Type:          enum.ArticleTypeMap.MustToProto(one.Type),
		Statement:     one.Statement,
		Commentable:   one.Commentable,
		Anonymous:     one.Anonymous,
		ViewCount:     one.ViewCount,
		ThankCount:    one.ThankCount,
		LikeCount:     one.LikeCount,
		CollectCount:  one.CollectCount,
		WatchCount:    one.WatchCount,
		ReplyCount:    one.ReplyCount,
	}
	switch one.Type {
	case enum.ArticleTypeQA:
		article.TypeParams = &v1.GetArticle_Resp_Article_Qa{
			Qa: &v1.GetArticle_Resp_Article_QA{
				BountyPoints:     one.BountyPoints,
				AcceptedAnswerId: one.AcceptedAnswerID,
			},
		}
	}
	if one.PublishedAt != nil {
		article.PublishedAt = timestamppb.New(*one.PublishedAt)
	}
	if one.EditedAt != nil {
		article.EditedAt = timestamppb.New(*one.EditedAt)
	}
	return &v1.GetArticle_Resp{
		Article: article,
	}, err
}

func (s *ArticleService) MapViewerActionStates(ctx context.Context, req *v1.MapArticleViewerActionStates_Req) (rsp *v1.MapArticleViewerActionStates_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	statesResp, err := s.articleUsecase.MapViewerActionStates(ctx, &usecase.ArticleMapViewerActionStatesReq{
		ArticleIDs: req.GetArticleIds(),
		UserID:     req.UserId,
	})
	if err != nil {
		return nil, err
	}
	states := statesResp
	reply := make(map[int64]*v1.MapArticleViewerActionStates_Resp_ArticleViewerActionState, len(states))
	for articleID, state := range states {
		reply[articleID] = &v1.MapArticleViewerActionStates_Resp_ArticleViewerActionState{
			Liked:     state.Liked,
			Thanked:   state.Thanked,
			Collected: state.Collected,
			Watched:   state.Watched,
		}
	}
	return &v1.MapArticleViewerActionStates_Resp{
		States: reply,
	}, nil
}

func (s *ArticleService) Reward(ctx context.Context, req *v1.RewardArticle_Req) (rsp *v1.RewardArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.Reward(ctx, &usecase.ArticleRewardReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Points:    req.Points,
	})
	return &v1.RewardArticle_Resp{}, err
}

func (s *ArticleService) View(ctx context.Context, req *v1.ViewArticle_Req) (rsp *v1.ViewArticle_Resp, err error) {
	err = s.articleUsecase.View(ctx, &usecase.ArticleViewReq{
		ArticleID:    req.ArticleId,
		ViewerUserID: req.ViewerUserId,
	})
	return &v1.ViewArticle_Resp{}, err
}

func (s *ArticleService) Like(ctx context.Context, req *v1.LikeArticle_Req) (rsp *v1.LikeArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	likeResp, err := s.articleUsecase.Like(ctx, &usecase.ArticleLikeReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Active:    req.Liked,
	})
	return &v1.LikeArticle_Resp{
		Liked: likeResp,
	}, err
}

func (s *ArticleService) Thank(ctx context.Context, req *v1.ThankArticle_Req) (rsp *v1.ThankArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	thankResp, err := s.articleUsecase.Thank(ctx, &usecase.ArticleThankReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Active:    req.Thanked,
	})
	return &v1.ThankArticle_Resp{
		Thanked: thankResp,
	}, err
}

func (s *ArticleService) Collect(ctx context.Context, req *v1.CollectArticle_Req) (rsp *v1.CollectArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	collectResp, err := s.articleUsecase.Collect(ctx, &usecase.ArticleCollectReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Active:    req.Collected,
	})
	return &v1.CollectArticle_Resp{
		Collected: collectResp,
	}, err
}

func (s *ArticleService) Watch(ctx context.Context, req *v1.WatchArticle_Req) (rsp *v1.WatchArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	watchResp, err := s.articleUsecase.Watch(ctx, &usecase.ArticleWatchReq{
		ArticleID: req.ArticleId,
		UserID:    req.UserId,
		Active:    req.Watched,
	})
	return &v1.WatchArticle_Resp{
		Watched: watchResp,
	}, err
}

func (s *ArticleService) AcceptAnswer(ctx context.Context, req *v1.AcceptAnswerArticle_Req) (rsp *v1.AcceptAnswerArticle_Resp, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.AcceptAnswer(ctx, &usecase.ArticleAcceptAnswerReq{
		ArticleID: req.ArticleId,
		CommentID: req.CommentId,
		UserID:    req.UserId,
	})
	return &v1.AcceptAnswerArticle_Resp{}, err
}
