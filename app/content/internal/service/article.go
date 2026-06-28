package service

import (
	"common/pkg/apperror"
	"common/pkg/server"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/biz/usecase"
	"content/internal/enum"
	"context"

	v1 "common/proto/gen/content/v1"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleService struct {
	v1.UnimplementedContentArticleServiceServer

	articleUsecase *usecase.ArticleUsecase
}

func (s *ArticleService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentArticleServiceServer(gs, s)
}

func NewArticleService(
	articleUsecase *usecase.ArticleUsecase,
) *ArticleService {
	return &ArticleService{
		articleUsecase: articleUsecase,
	}
}

func (s *ArticleService) Create(ctx context.Context, req *v1.CreateArticle_Request) (rsp *v1.CreateArticle_Reply, err error) {
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
	save, err := s.articleUsecase.Add(ctx, &model.Article{
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
	}, article.GetTagIds())
	if err != nil {
		return nil, err
	}
	articleReply := &v1.Article{
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
		articleReply.TypeParams = &v1.Article_Qa{Qa: &v1.Article_QA{
			BountyPoints:     save.BountyPoints,
			AcceptedAnswerId: save.AcceptedAnswerID,
		}}
	}
	if save.PublishedAt != nil {
		articleReply.PublishedAt = timestamppb.New(*save.PublishedAt)
	}
	if save.EditedAt != nil {
		articleReply.EditedAt = timestamppb.New(*save.EditedAt)
	}
	return &v1.CreateArticle_Reply{
		Article: articleReply,
	}, nil
}

func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticle_Request) (rsp *v1.PublishArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	visibility, ok := enum.ArticleVisibilityMap.ToEnum(req.Visibility)
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
	}
	err = s.articleUsecase.Publish(ctx, req.ArticleId, req.UserId, visibility)
	return &v1.PublishArticle_Reply{}, err
}

func (s *ArticleService) AddPostscript(ctx context.Context, req *v1.AddPostscriptArticle_Request) (rsp *v1.AddPostscriptArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	save, err := s.articleUsecase.AddPostscript(ctx, req.ArticleId, req.Content, req.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.AddPostscriptArticle_Reply{
		ArticlePostscript: &v1.ArticlePostscript{
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

func (s *ArticleService) ListPostscripts(ctx context.Context, req *v1.ListArticlePostscripts_Request) (rsp *v1.ListArticlePostscripts_Reply, err error) {
	rows, err := s.articleUsecase.ListPostscripts(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	reply := make([]*v1.ArticlePostscript, 0, len(rows))
	for _, item := range rows {
		reply = append(reply, &v1.ArticlePostscript{
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
	return &v1.ListArticlePostscripts_Reply{Rows: reply}, nil
}

func (s *ArticleService) Update(ctx context.Context, req *v1.UpdateArticle_Request) (rsp *v1.UpdateArticle_Reply, err error) {
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
	update, err := s.articleUsecase.Update(ctx, &model.Article{
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
	}, article.GetTagIds())
	if err != nil {
		return nil, err
	}
	articleReply := &v1.Article{
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
		articleReply.TypeParams = &v1.Article_Qa{Qa: &v1.Article_QA{
			BountyPoints:     update.BountyPoints,
			AcceptedAnswerId: update.AcceptedAnswerID,
		}}
	}
	if update.PublishedAt != nil {
		articleReply.PublishedAt = timestamppb.New(*update.PublishedAt)
	}
	if update.EditedAt != nil {
		articleReply.EditedAt = timestamppb.New(*update.EditedAt)
	}
	return &v1.UpdateArticle_Reply{
		Article: articleReply,
	}, nil
}

func (s *ArticleService) DiscardDraft(ctx context.Context, req *v1.DiscardDraftArticle_Request) (rsp *v1.DiscardDraftArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.DiscardDraft(ctx, req.ArticleId, req.UserId)
	return &v1.DiscardDraftArticle_Reply{}, err
}

func (s *ArticleService) MakePrivate(ctx context.Context, req *v1.MakePrivateArticle_Request) (*v1.MakePrivateArticle_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.MakePrivate(ctx, req.ArticleId, req.UserId)
	return &v1.MakePrivateArticle_Reply{}, err
}

func (s *ArticleService) MakePublic(ctx context.Context, req *v1.MakePublicArticle_Request) (*v1.MakePublicArticle_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.MakePublic(ctx, req.ArticleId, req.UserId)
	return &v1.MakePublicArticle_Reply{}, err
}

func (s *ArticleService) Archive(ctx context.Context, req *v1.ArchiveArticle_Request) (*v1.ArchiveArticle_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.Archive(ctx, req.ArticleId, req.UserId, req.Reason)
	return &v1.ArchiveArticle_Reply{}, err
}

func (s *ArticleService) Unarchive(ctx context.Context, req *v1.UnarchiveArticle_Request) (*v1.UnarchiveArticle_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.Unarchive(ctx, req.ArticleId, req.UserId, req.Reason)
	return &v1.UnarchiveArticle_Reply{}, err
}

func (s *ArticleService) Hide(ctx context.Context, req *v1.HideArticle_Request) (rsp *v1.HideArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.Hide(ctx, req.ArticleId, req.UserId, req.Reason)
	return &v1.HideArticle_Reply{}, err
}

func (s *ArticleService) Unhide(ctx context.Context, req *v1.UnhideArticle_Request) (*v1.UnhideArticle_Reply, error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.Unhide(ctx, req.ArticleId, req.UserId, req.Reason)
	return &v1.UnhideArticle_Reply{}, err
}

func (s *ArticleService) Lock(ctx context.Context, req *v1.LockArticle_Request) (rsp *v1.LockArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.Lock(ctx, req.ArticleId, req.UserId, req.Reason)
	return &v1.LockArticle_Reply{}, err
}

func (s *ArticleService) Unlock(ctx context.Context, req *v1.UnlockArticle_Request) (rsp *v1.UnlockArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.Unlock(ctx, req.ArticleId, req.UserId, req.Reason)
	return &v1.UnlockArticle_Reply{}, err
}

func (s *ArticleService) List(ctx context.Context, req *v1.ListArticles_Request) (rsp *v1.ListArticles_Reply, err error) {
	req.Query = util.OrDefault(req.Query, &v1.ArticleQueryParams{})
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

	query := &repo.ArticleGetReq{
		TagId:           req.Query.TagId,
		DomainId:        req.Query.DomainId,
		PublishStatus:   publishStatus,
		PublishStatuses: publishStatuses,
		Visibility:      visibility,
		Visibilities:    visibilities,
		Restriction:     restriction,
		Restrictions:    restrictions,
		AuthorId:        req.Query.AuthorId,
		Order:           dbOrder,
		Type:            dbType,
		Keyword:         req.Query.Keyword,
	}
	reply, _, err := s.articleUsecase.Page(ctx, server.GetPageMax(), query)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Article, 0, len(reply))
	for _, item := range reply {
		row := &v1.Article{
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
			row.TypeParams = &v1.Article_Qa{Qa: &v1.Article_QA{
				BountyPoints:     item.BountyPoints,
				AcceptedAnswerId: item.AcceptedAnswerID,
			}}
		}
		if item.PublishedAt != nil {
			row.PublishedAt = timestamppb.New(*item.PublishedAt)
		}
		if item.EditedAt != nil {
			row.EditedAt = timestamppb.New(*item.EditedAt)
		}
		rows = append(rows, row)
	}
	return &v1.ListArticles_Reply{
		Rows: rows,
	}, err
}

func (s *ArticleService) Page(ctx context.Context, req *v1.PageArticles_Request) (rsp *v1.PageArticles_Reply, err error) {
	req.Query = util.OrDefault(req.Query, &v1.ArticleQueryParams{})
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

	query := &repo.ArticleGetReq{
		TagId:           req.Query.TagId,
		DomainId:        req.Query.DomainId,
		PublishStatus:   publishStatus,
		PublishStatuses: publishStatuses,
		Visibility:      visibility,
		Visibilities:    visibilities,
		Restriction:     restriction,
		Restrictions:    restrictions,
		AuthorId:        req.Query.AuthorId,
		Order:           dbOrder,
		Type:            dbType,
		Keyword:         req.Query.Keyword,
	}
	reply, page, err := s.articleUsecase.Page(ctx, req.Page, query)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Article, 0, len(reply))
	for _, item := range reply {
		row := &v1.Article{
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
			row.TypeParams = &v1.Article_Qa{Qa: &v1.Article_QA{
				BountyPoints:     item.BountyPoints,
				AcceptedAnswerId: item.AcceptedAnswerID,
			}}
		}
		if item.PublishedAt != nil {
			row.PublishedAt = timestamppb.New(*item.PublishedAt)
		}
		if item.EditedAt != nil {
			row.EditedAt = timestamppb.New(*item.EditedAt)
		}
		rows = append(rows, row)
	}
	return &v1.PageArticles_Reply{
		Page: page,
		Rows: rows,
	}, err
}

func (s *ArticleService) Get(ctx context.Context, req *v1.GetArticle_Request) (rsp *v1.GetArticle_Reply, err error) {
	one, err := s.articleUsecase.Get(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	article := &v1.Article{
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
		article.TypeParams = &v1.Article_Qa{Qa: &v1.Article_QA{
			BountyPoints:     one.BountyPoints,
			AcceptedAnswerId: one.AcceptedAnswerID,
		}}
	}
	if one.PublishedAt != nil {
		article.PublishedAt = timestamppb.New(*one.PublishedAt)
	}
	if one.EditedAt != nil {
		article.EditedAt = timestamppb.New(*one.EditedAt)
	}
	return &v1.GetArticle_Reply{Article: article}, err
}

func (s *ArticleService) MapViewerActionStates(ctx context.Context, req *v1.MapArticleViewerActionStates_Request) (rsp *v1.MapArticleViewerActionStates_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	states, err := s.articleUsecase.MapViewerActionStates(ctx, req.GetArticleIds(), req.UserId)
	if err != nil {
		return nil, err
	}
	reply := make(map[int64]*v1.ArticleViewerActionState, len(states))
	for articleID, state := range states {
		reply[articleID] = &v1.ArticleViewerActionState{
			Liked:     state.Liked,
			Thanked:   state.Thanked,
			Collected: state.Collected,
			Watched:   state.Watched,
		}
	}
	return &v1.MapArticleViewerActionStates_Reply{States: reply}, nil
}

func (s *ArticleService) Reward(ctx context.Context, req *v1.RewardArticle_Request) (rsp *v1.RewardArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.Reward(ctx, req.ArticleId, req.UserId, req.Points)
	return &v1.RewardArticle_Reply{}, err
}

func (s *ArticleService) View(ctx context.Context, req *v1.ViewArticle_Request) (rsp *v1.ViewArticle_Reply, err error) {
	err = s.articleUsecase.View(ctx, req.ArticleId, req.ViewerUserId)
	return &v1.ViewArticle_Reply{}, err
}

func (s *ArticleService) Like(ctx context.Context, req *v1.LikeArticle_Request) (rsp *v1.LikeArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	liked, err := s.articleUsecase.Like(ctx, req.ArticleId, req.UserId, req.Liked)
	return &v1.LikeArticle_Reply{Liked: liked}, err
}

func (s *ArticleService) Thank(ctx context.Context, req *v1.ThankArticle_Request) (rsp *v1.ThankArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	thanked, err := s.articleUsecase.Thank(ctx, req.ArticleId, req.UserId, req.Thanked)
	return &v1.ThankArticle_Reply{Thanked: thanked}, err
}

func (s *ArticleService) Collect(ctx context.Context, req *v1.CollectArticle_Request) (rsp *v1.CollectArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	collected, err := s.articleUsecase.Collect(ctx, req.ArticleId, req.UserId, req.Collected)
	return &v1.CollectArticle_Reply{Collected: collected}, err
}

func (s *ArticleService) Watch(ctx context.Context, req *v1.WatchArticle_Request) (rsp *v1.WatchArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	watched, err := s.articleUsecase.Watch(ctx, req.ArticleId, req.UserId, req.Watched)
	return &v1.WatchArticle_Reply{Watched: watched}, err
}

func (s *ArticleService) AcceptAnswer(ctx context.Context, req *v1.AcceptAnswerArticle_Request) (rsp *v1.AcceptAnswerArticle_Reply, err error) {
	if req.UserId <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.AcceptAnswer(ctx, req.ArticleId, req.CommentId, req.UserId)
	return &v1.AcceptAnswerArticle_Reply{}, err
}
