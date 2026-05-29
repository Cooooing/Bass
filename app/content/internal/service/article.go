package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	"common/pkg/util"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/biz/usecase"
	"content/internal/enum"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
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
		return nil, cerrors.ErrorBadRequest("article is required")
	}

	var tags []*model.Tag
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			if tag.Id == nil && tag.Name == "" {
				return nil, cerrors.ErrorBadRequest("tag id or name is required")
			}
			tagStatus, ok := enum.TagStatusMap.ToEnum(util.DerefOrDefault(tag.Status, v1.TagStatus_TAG_STATUS_NORMAL))
			if !ok {
				return nil, cerrors.ErrorBadRequest("invalid tag status")
			}
			saveTag := &model.Tag{
				Name:        tag.Name,
				Description: tag.Description,
				DomainID:    tag.DomainId,
				Status:      tagStatus,
				CreatedBy:   new(req.UserId),
				UpdatedBy:   new(req.UserId),
			}
			if tag.Id != nil {
				saveTag.ID = *tag.Id
			}
			tags = append(tags, saveTag)
		}
	}

	dbStatus, ok := enum.ArticleStatusMap.ToEnum(article.Status)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid article status")
	}
	dbType, ok := enum.ArticleTypeMap.ToEnum(article.Type)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid article type")
	}
	save, err := s.articleUsecase.Add(ctx, &model.Article{
		Title:         article.Title,
		Content:       article.Content,
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Status:        dbStatus,
		Type:          dbType,
		BountyPoints:  util.If(article.Type != v1.ArticleType_ARTICLE_TYPE_QA, nil, article.BountyPoints),
		Statement:     article.Statement,
		Commentable:   util.DerefOrDefault(article.Commentable, true),
		Anonymous:     util.DerefOrDefault(article.Anonymous, false),
		Listable:      util.DerefOrDefault(article.Listable, true),
		CreatedBy:     new(req.UserId),
		UpdatedBy:     new(req.UserId),
	}, tags)
	if err != nil {
		return nil, err
	}
	save.ParseContent()
	save.ParseRewardContent()
	articleReply := &v1.Article{
		CreatedAt:           timestamppb.New(*save.CreatedAt),
		UpdatedAt:           timestamppb.New(*save.UpdatedAt),
		CreatedBy:           save.CreatedBy,
		UpdatedBy:           save.UpdatedBy,
		Id:                  save.ID,
		Title:               save.Title,
		Content:             save.Content,
		ContentRender:       save.ContentRender,
		RewardContent:       save.RewardContent,
		RewardContentRender: save.RewardContentRender,
		HasPostscript:       save.HasPostscript,
		HasReward:           util.IsNotNil(save.RewardPoints),
		RewardPoints:        save.RewardPoints,
		Status:              enum.ArticleStatusMap.MustToProto(save.Status),
		Type:                enum.ArticleTypeMap.MustToProto(save.Type),
		Statement:           save.Statement,
		Commentable:         save.Commentable,
		Anonymous:           save.Anonymous,
		ViewCount:           save.ViewCount,
		ThankCount:          save.ThankCount,
		LikeCount:           save.LikeCount,
		CollectCount:        save.CollectCount,
		WatchCount:          save.WatchCount,
		ReplyCount:          save.ReplyCount,
		BountyPoints:        save.BountyPoints,
		AcceptedAnswerId:    save.AcceptedAnswerID,
		AuthorUser:          save.AuthorUser,
		LastReplyUser:       save.LastReplyCommentUser,
		CoverImageUrl:       save.CoverImageUrl,
	}
	if save.LastReplyCommentAt != nil {
		articleReply.LastReplyAt = timestamppb.New(*save.LastReplyCommentAt)
	}
	return &v1.CreateArticle_Reply{
		Article: articleReply,
	}, nil
}

func (s *ArticleService) UpdateDraft(ctx context.Context, req *v1.UpdateDraftArticle_Request) (rsp *v1.UpdateDraftArticle_Reply, err error) {
	article := req.Article
	if article == nil || article.Id == nil {
		return nil, cerrors.ErrorBadRequest("article is required")
	}

	var tags []*model.Tag
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			if tag.Id == nil && tag.Name == "" {
				return nil, cerrors.ErrorBadRequest("tag id or name is required")
			}
			tagStatus, ok := enum.TagStatusMap.ToEnum(util.DerefOrDefault(tag.Status, v1.TagStatus_TAG_STATUS_NORMAL))
			if !ok {
				return nil, cerrors.ErrorBadRequest("invalid tag status")
			}
			saveTag := &model.Tag{
				Name:        tag.Name,
				Description: tag.Description,
				DomainID:    tag.DomainId,
				Status:      tagStatus,
				CreatedBy:   new(req.UserId),
				UpdatedBy:   new(req.UserId),
			}
			if tag.Id != nil {
				saveTag.ID = *tag.Id
			}
			tags = append(tags, saveTag)
		}
	}

	dbStatus2, ok := enum.ArticleStatusMap.ToEnum(article.Status)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid article status")
	}
	dbType2, ok := enum.ArticleTypeMap.ToEnum(article.Type)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid article type")
	}
	update, err := s.articleUsecase.UpdateDraft(ctx, &model.Article{
		ID:            *req.Article.Id,
		Title:         article.Title,
		Content:       article.Content,
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Status:        dbStatus2,
		Type:          dbType2,
		BountyPoints:  util.If(article.Type != v1.ArticleType_ARTICLE_TYPE_QA, nil, article.BountyPoints),
		Statement:     article.Statement,
		Commentable:   util.DerefOrDefault(article.Commentable, true),
		Anonymous:     util.DerefOrDefault(article.Anonymous, false),
		Listable:      util.DerefOrDefault(article.Listable, true),
		UpdatedBy:     new(req.UserId),
	}, tags)
	if err != nil {
		return nil, err
	}
	update.ParseContent()
	update.ParseRewardContent()
	articleReply := &v1.Article{
		CreatedAt:           timestamppb.New(*update.CreatedAt),
		UpdatedAt:           timestamppb.New(*update.UpdatedAt),
		CreatedBy:           update.CreatedBy,
		UpdatedBy:           update.UpdatedBy,
		Id:                  update.ID,
		Title:               update.Title,
		Content:             update.Content,
		ContentRender:       update.ContentRender,
		RewardContent:       update.RewardContent,
		RewardContentRender: update.RewardContentRender,
		HasPostscript:       update.HasPostscript,
		HasReward:           util.IsNotNil(update.RewardPoints),
		RewardPoints:        update.RewardPoints,
		Status:              enum.ArticleStatusMap.MustToProto(update.Status),
		Type:                enum.ArticleTypeMap.MustToProto(update.Type),
		Statement:           update.Statement,
		Commentable:         update.Commentable,
		Anonymous:           update.Anonymous,
		ViewCount:           update.ViewCount,
		ThankCount:          update.ThankCount,
		LikeCount:           update.LikeCount,
		CollectCount:        update.CollectCount,
		WatchCount:          update.WatchCount,
		ReplyCount:          update.ReplyCount,
		BountyPoints:        update.BountyPoints,
		AcceptedAnswerId:    update.AcceptedAnswerID,
		AuthorUser:          update.AuthorUser,
		LastReplyUser:       update.LastReplyCommentUser,
		CoverImageUrl:       update.CoverImageUrl,
	}
	if update.LastReplyCommentAt != nil {
		articleReply.LastReplyAt = timestamppb.New(*update.LastReplyCommentAt)
	}
	return &v1.UpdateDraftArticle_Reply{
		Article: articleReply,
	}, nil
}

func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticle_Request) (rsp *v1.PublishArticle_Reply, err error) {
	err = s.articleUsecase.Publish(ctx, req.ArticleId, req.UserId)
	return &v1.PublishArticle_Reply{}, err
}

func (s *ArticleService) AddPostscript(ctx context.Context, req *v1.AddPostscriptArticle_Request) (rsp *v1.AddPostscriptArticle_Reply, err error) {
	save, err := s.articleUsecase.AddPostscript(ctx, req.ArticleId, req.UserId, req.Content)
	if err != nil {
		return nil, err
	}
	save.ParseContent()
	return &v1.AddPostscriptArticle_Reply{
		ArticlePostscript: &v1.ArticlePostscript{
			CreatedAt:     timestamppb.New(*save.CreatedAt),
			UpdatedAt:     timestamppb.New(*save.UpdatedAt),
			CreatedBy:     save.CreatedBy,
			UpdatedBy:     save.UpdatedBy,
			Id:            save.ID,
			ArticleId:     save.ArticleID,
			Content:       save.Content,
			ContentRender: save.ContentRender,
		},
	}, err
}

func (s *ArticleService) Update(ctx context.Context, req *v1.UpdateArticle_Request) (rsp *v1.UpdateArticle_Reply, err error) {
	err = s.articleUsecase.UpdateArticle(ctx, req.ArticleId, req.UserId, req.Status, req.Commentable, req.Anonymous, req.Listable)
	return &v1.UpdateArticle_Reply{}, err
}

func (s *ArticleService) Delete(ctx context.Context, req *v1.DeleteArticle_Request) (rsp *v1.DeleteArticle_Reply, err error) {
	err = s.articleUsecase.Delete(ctx, req.ArticleId, req.UserId)
	return &v1.DeleteArticle_Reply{}, err
}

func (s *ArticleService) List(ctx context.Context, req *v1.ListArticles_Request) (rsp *v1.ListArticles_Reply, err error) {
	req.Query = util.OrDefault(req.Query, &v1.ArticleQueryParams{})

	if req.Query.Status != nil {
		if _, ok := enum.ArticleStatusMap.ToEnum(*req.Query.Status); !ok {
			return nil, cerrors.ErrorBadRequest("invalid article status")
		}
	}

	reply, page, err := s.articleUsecase.Page(ctx, req.Page, &repo.ArticleGetReq{
		TagId:    req.Query.TagId,
		DomainId: req.Query.DomainId,
		Status:   req.Query.Status,
		AuthorId: req.Query.AuthorId,
		Order:    req.Query.Order,
		Type:     req.Query.Type,
		Keyword:  req.Query.Keyword,
		Listable: req.Query.Listable,
	})
	rows := make([]*v1.Article, 0, len(reply))
	for _, item := range reply {
		item.ParseContent()
		item.ParseRewardContent()
		if item.IsSummary {
			item.Summary()
		}
		row := &v1.Article{
			CreatedAt:           timestamppb.New(*item.CreatedAt),
			UpdatedAt:           timestamppb.New(*item.UpdatedAt),
			CreatedBy:           item.CreatedBy,
			UpdatedBy:           item.UpdatedBy,
			Id:                  item.ID,
			Title:               item.Title,
			Content:             item.Content,
			ContentRender:       item.ContentRender,
			RewardContent:       item.RewardContent,
			RewardContentRender: item.RewardContentRender,
			HasPostscript:       item.HasPostscript,
			HasReward:           util.IsNotNil(item.RewardPoints),
			RewardPoints:        item.RewardPoints,
			Status:              enum.ArticleStatusMap.MustToProto(item.Status),
			Type:                enum.ArticleTypeMap.MustToProto(item.Type),
			Statement:           item.Statement,
			Commentable:         item.Commentable,
			Anonymous:           item.Anonymous,
			ViewCount:           item.ViewCount,
			ThankCount:          item.ThankCount,
			LikeCount:           item.LikeCount,
			CollectCount:        item.CollectCount,
			WatchCount:          item.WatchCount,
			ReplyCount:          item.ReplyCount,
			BountyPoints:        item.BountyPoints,
			AcceptedAnswerId:    item.AcceptedAnswerID,
			AuthorUser:          item.AuthorUser,
			LastReplyUser:       item.LastReplyCommentUser,
			CoverImageUrl:       item.CoverImageUrl,
		}
		if item.LastReplyCommentAt != nil {
			row.LastReplyAt = timestamppb.New(*item.LastReplyCommentAt)
		}
		rows = append(rows, row)
	}
	return &v1.ListArticles_Reply{
		Page: page,
		Rows: rows,
	}, err
}

func (s *ArticleService) Get(ctx context.Context, req *v1.GetArticle_Request) (rsp *v1.GetArticle_Reply, err error) {
	one, err := s.articleUsecase.Get(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	one.ParseContent()
	one.ParseRewardContent()
	if one.IsSummary {
		one.Summary()
	}
	article := &v1.Article{
		CreatedAt:           timestamppb.New(*one.CreatedAt),
		UpdatedAt:           timestamppb.New(*one.UpdatedAt),
		CreatedBy:           one.CreatedBy,
		UpdatedBy:           one.UpdatedBy,
		Id:                  one.ID,
		Title:               one.Title,
		Content:             one.Content,
		ContentRender:       one.ContentRender,
		RewardContent:       one.RewardContent,
		RewardContentRender: one.RewardContentRender,
		HasPostscript:       one.HasPostscript,
		HasReward:           util.IsNotNil(one.RewardPoints),
		RewardPoints:        one.RewardPoints,
		Status:              enum.ArticleStatusMap.MustToProto(one.Status),
		Type:                enum.ArticleTypeMap.MustToProto(one.Type),
		Statement:           one.Statement,
		Commentable:         one.Commentable,
		Anonymous:           one.Anonymous,
		ViewCount:           one.ViewCount,
		ThankCount:          one.ThankCount,
		LikeCount:           one.LikeCount,
		CollectCount:        one.CollectCount,
		WatchCount:          one.WatchCount,
		ReplyCount:          one.ReplyCount,
		BountyPoints:        one.BountyPoints,
		AcceptedAnswerId:    one.AcceptedAnswerID,
		AuthorUser:          one.AuthorUser,
		LastReplyUser:       one.LastReplyCommentUser,
		CoverImageUrl:       one.CoverImageUrl,
	}
	if one.LastReplyCommentAt != nil {
		article.LastReplyAt = timestamppb.New(*one.LastReplyCommentAt)
	}
	return &v1.GetArticle_Reply{Article: article}, err
}

func (s *ArticleService) Reward(ctx context.Context, req *v1.RewardArticle_Request) (rsp *v1.RewardArticle_Reply, err error) {
	err = s.articleUsecase.Action(ctx, req.ArticleId, req.UserId, v1.ArticleAction_ARTICLE_ACTION_REWARD, true)
	return &v1.RewardArticle_Reply{}, err
}

func (s *ArticleService) Like(ctx context.Context, req *v1.LikeArticle_Request) (rsp *v1.LikeArticle_Reply, err error) {
	err = s.articleUsecase.Action(ctx, req.ArticleId, req.UserId, v1.ArticleAction_ARTICLE_ACTION_LIKE, req.Active)
	return &v1.LikeArticle_Reply{}, err
}

func (s *ArticleService) Thank(ctx context.Context, req *v1.ThankArticle_Request) (rsp *v1.ThankArticle_Reply, err error) {
	err = s.articleUsecase.Action(ctx, req.ArticleId, req.UserId, v1.ArticleAction_ARTICLE_ACTION_THANK, req.Active)
	return &v1.ThankArticle_Reply{}, err
}

func (s *ArticleService) Collect(ctx context.Context, req *v1.CollectArticle_Request) (rsp *v1.CollectArticle_Reply, err error) {
	err = s.articleUsecase.Action(ctx, req.ArticleId, req.UserId, v1.ArticleAction_ARTICLE_ACTION_COLLECT, req.Active)
	return &v1.CollectArticle_Reply{}, err
}

func (s *ArticleService) Watch(ctx context.Context, req *v1.WatchArticle_Request) (rsp *v1.WatchArticle_Reply, err error) {
	err = s.articleUsecase.Action(ctx, req.ArticleId, req.UserId, v1.ArticleAction_ARTICLE_ACTION_WATCH, req.Active)
	return &v1.WatchArticle_Reply{}, err
}

func (s *ArticleService) AcceptAnswer(ctx context.Context, req *v1.AcceptAnswerArticle_Request) (rsp *v1.AcceptAnswerArticle_Reply, err error) {
	err = s.articleUsecase.AcceptAnswer(ctx, req.ArticleId, req.UserId, req.CommentId)
	return &v1.AcceptAnswerArticle_Reply{}, err
}
