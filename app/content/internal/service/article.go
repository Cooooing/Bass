package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"content/internal/biz"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ArticleService struct {
	v1.UnimplementedArticleServiceServer
	*BaseService

	articleDomain *biz.ArticleDomain
	articleRepo   repo.ArticleRepo
}

func (s *ArticleService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterArticleServiceServer(gs, s)
}

func (s *ArticleService) RegisterHttp(hs *http.Server) {
	v1.RegisterArticleServiceHTTPServer(hs, s)
}

func NewArticleService(baseService *BaseService, articleDomain *biz.ArticleDomain, articleRepo repo.ArticleRepo) *ArticleService {
	return &ArticleService{
		BaseService:   baseService,
		articleDomain: articleDomain,
		articleRepo:   articleRepo,
	}
}

func (s *ArticleService) AcceptAnswer(ctx context.Context, req *v1.AcceptAnswerArticleRequest) (rsp *v1.AcceptAnswerArticleReply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ArticleService) Add(ctx context.Context, req *v1.AddArticleRequest) (rsp *v1.AddArticleReply, err error) {
	if req.Status != int32(cv1.ArticleStatus_ArticleNormal) || req.Status != int32(cv1.ArticleStatus_ArticleDrafts) {
		return nil, errors.New("status only be 0(normal) or 3(drafts)")
	}
	user, err := s.tokenRepo.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	_, err = s.articleDomain.Add(ctx, &model.Article{
		UserID:        user.ID,
		Title:         req.Title,
		Content:       req.Content,
		RewardContent: &req.RewardContent,
		RewardPoints:  int(req.RewardPoints),
		Status:        int(req.Status),
		Type:          int(req.Type),
		BountyPoints:  int(req.BountyPoints),
	})
	if err != nil {
		return nil, err
	}
	return &v1.AddArticleReply{}, nil
}

func (s *ArticleService) AddPostscript(ctx context.Context, req *v1.AddPostscriptArticleRequest) (rsp *v1.AddPostscriptArticleReply, err error) {
	user, err := s.tokenRepo.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	// 只有作者可以添加附言
	if article, err := s.articleRepo.GetArticleById(ctx, s.db, int(req.ArticleId)); err != nil || article.UserID != user.ID {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("you are not the author")
	}

	err = s.articleDomain.AddPostscript(ctx, int(req.ArticleId), req.Content)
	return &v1.AddPostscriptArticleReply{}, err
}

func (s *ArticleService) Collect(ctx context.Context, req *v1.CollectArticleRequest) (rsp *v1.CollectArticleReply, err error) {
	user, err := s.tokenRepo.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	err = s.articleDomain.Action(ctx, int(req.ArticleId), user.ID, cv1.ArticleAction_ArticleActionCollect, req.Active)
	return &v1.CollectArticleReply{}, err
}

func (s *ArticleService) Delete(ctx context.Context, req *v1.DeleteArticleRequest) (rsp *v1.DeleteArticleReply, err error) {
	user, err := s.tokenRepo.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	article, err := s.articleRepo.GetArticleById(ctx, s.db, int(req.ArticleId))
	if err != nil {
		return nil, err
	}
	// 只能删除草稿
	if article.Status != int(cv1.ArticleStatus_ArticleDrafts) {
		return nil, errors.New("only drafts can be deleted")
	}
	// 只有作者可以删除草稿
	if article.UserID != user.ID {
		return nil, errors.New("you are not the author")
	}
	err = s.articleRepo.Delete(ctx, int(req.ArticleId))
	return &v1.DeleteArticleReply{}, err
}

func (s *ArticleService) Get(ctx context.Context, req *v1.GetArticleRequest) (rsp *v1.GetArticleReply, err error) {
	return &v1.GetArticleReply{}, nil
}

func (s *ArticleService) Like(ctx context.Context, req *v1.LikeArticleRequest) (rsp *v1.LikeArticleReply, err error) {
	user, err := s.tokenRepo.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	err = s.articleDomain.Action(ctx, int(req.ArticleId), user.ID, cv1.ArticleAction_ArticleActionLike, req.Active)
	return &v1.LikeArticleReply{}, err
}

func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticleRequest) (rsp *v1.PublishArticleReply, err error) {
	user, err := s.tokenRepo.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	article, err := s.articleRepo.GetArticleById(ctx, s.db, int(req.ArticleId))
	if err != nil {
		return nil, err
	}
	// 只能发布草稿
	if article.Status != int(cv1.ArticleStatus_ArticleDrafts) {
		return nil, errors.New("only drafts can be publish")
	}
	// 只有作者可以发布草稿
	if article.UserID != user.ID {
		return nil, errors.New("you are not the author")
	}
	err = s.articleDomain.Publish(ctx, int(req.ArticleId))
	return &v1.PublishArticleReply{}, err
}

func (s *ArticleService) Reward(ctx context.Context, req *v1.RewardArticleRequest) (rsp *v1.RewardArticleReply, err error) {
	return &v1.RewardArticleReply{}, nil
}

func (s *ArticleService) Thank(ctx context.Context, req *v1.ThankArticleRequest) (rsp *v1.ThankArticleReply, err error) {
	return &v1.ThankArticleReply{}, nil
}

func (s *ArticleService) Update(ctx context.Context, req *v1.UpdateArticleRequest) (rsp *v1.UpdateArticleReply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ArticleService) Watch(ctx context.Context, req *v1.WatchArticleRequest) (rsp *v1.WatchArticleReply, err error) {
	// TODO implement me
	panic("implement me")
}
