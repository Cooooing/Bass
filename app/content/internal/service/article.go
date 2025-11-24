package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"common/pkg/util/base"
	"content/internal/biz"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ArticleService struct {
	v1.UnimplementedContentArticleServiceServer
	*BaseService

	articleDomain *biz.ArticleDomain
	articleRepo   repo.ArticleRepo
}

func (s *ArticleService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentArticleServiceServer(gs, s)
}

func (s *ArticleService) RegisterHttp(hs *http.Server) {
	v1.RegisterContentArticleServiceHTTPServer(hs, s)
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
	article := req.Article
	if article.Status != int32(cv1.ArticleStatus_ArticleNormal) && article.Status != int32(cv1.ArticleStatus_ArticleDrafts) {
		return nil, errors.New("status only be 0(normal) or 3(drafts)")
	}
	if article.Type != int32(cv1.ArticleType_ArticleTypeNormal) && article.Type != int32(cv1.ArticleType_ArticleTypeQA) && article.Type != int32(cv1.ArticleType_ArticleTypeVote) && article.Type != int32(cv1.ArticleType_ArticleTypeLottery) {
		return nil, errors.New("type only be 0(normal), 1(QA), 2(vote), 3(lottery)")
	}

	var tags []*model.Tag
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			tags = append(tags, &model.Tag{
				Name:         tag.Name,
				Description:  tag.Description,
				Status:       int32(cv1.TagStatus_TagNormal),
				ArticleCount: 1,
			})
		}
	}

	save, err := s.articleDomain.Add(ctx, &model.Article{
		Title:         article.Title,
		Content:       article.Content,
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Status:        article.Status,
		Type:          article.Type,
		BountyPoints:  base.If(article.Type != int32(cv1.ArticleType_ArticleTypeQA), nil, article.BountyPoints),
		Statement:     article.Statement,
		Commentable:   base.DerefOrDefault(article.Commentable, true),
		Anonymous:     base.DerefOrDefault(article.Anonymous, false),
		Listable:      base.DerefOrDefault(article.Listable, true),
	}, tags)
	if err != nil {
		return nil, err
	}
	return &v1.AddArticleReply{
		ArticleId: save.ID,
	}, nil
}

func (s *ArticleService) UpdateDraft(ctx context.Context, req *v1.UpdateArticleDraftRequest) (rsp *v1.UpdateArticleDraftReply, err error) {
	article := req.Article
	if article.Status != int32(cv1.ArticleStatus_ArticleNormal) && article.Status != int32(cv1.ArticleStatus_ArticleDrafts) {
		return nil, errors.New("status only be 0(normal) or 3(drafts)")
	}
	if article.Type != int32(cv1.ArticleType_ArticleTypeNormal) && article.Type != int32(cv1.ArticleType_ArticleTypeQA) && article.Type != int32(cv1.ArticleType_ArticleTypeVote) && article.Type != int32(cv1.ArticleType_ArticleTypeLottery) {
		return nil, errors.New("type only be 0(normal), 1(QA), 2(vote), 3(lottery)")
	}

	var tags []*model.Tag
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			tags = append(tags, &model.Tag{
				Name:         tag.Name,
				Description:  tag.Description,
				Status:       int32(cv1.TagStatus_TagNormal),
				ArticleCount: 1,
			})
		}
	}

	update, err := s.articleDomain.UpdateDraft(ctx, &model.Article{
		Title:         article.Title,
		Content:       article.Content,
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Status:        article.Status,
		Type:          article.Type,
		BountyPoints:  base.If(article.Type != int32(cv1.ArticleType_ArticleTypeQA), nil, article.BountyPoints),
		Statement:     article.Statement,
		Commentable:   base.DerefOrDefault(article.Commentable, true),
		Anonymous:     base.DerefOrDefault(article.Anonymous, false),
		Listable:      base.DerefOrDefault(article.Listable, true),
	}, tags)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateArticleDraftReply{
		ArticleId: update.ID,
	}, nil
}

func (s *ArticleService) AddPostscript(ctx context.Context, req *v1.AddPostscriptArticleRequest) (rsp *v1.AddPostscriptArticleReply, err error) {
	user := util.MustGetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	// 只有作者可以添加附言
	if article, err := s.articleRepo.GetById(ctx, s.db, req.ArticleId); err != nil || article == nil || *article.CreatedBy != user.ID {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("you are not the author")
	}

	err = s.articleDomain.AddPostscript(ctx, req.ArticleId, req.Content)
	return &v1.AddPostscriptArticleReply{}, err
}

func (s *ArticleService) Collect(ctx context.Context, req *v1.CollectArticleRequest) (rsp *v1.CollectArticleReply, err error) {
	user := util.MustGetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, cv1.ArticleAction_ArticleActionCollect, req.Active)
	return &v1.CollectArticleReply{}, err
}

func (s *ArticleService) Delete(ctx context.Context, req *v1.DeleteArticleRequest) (rsp *v1.DeleteArticleReply, err error) {
	user := util.MustGetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	err = ent.WithTx(ctx, s.db, func(tx *gen.Client) error {
		article, err := s.articleRepo.GetById(ctx, s.db, req.ArticleId)
		if err != nil {
			return err
		}
		// 只能删除草稿
		if article.Status != int32(cv1.ArticleStatus_ArticleDrafts) {
			return errors.New("only drafts can be deleted")
		}
		// 只有作者可以删除草稿
		if *article.CreatedBy != user.ID {
			return errors.New("you are not the author")
		}
		err = s.articleRepo.Delete(ctx, s.db, req.ArticleId)
		return err
	})
	return &v1.DeleteArticleReply{}, err
}

func (s *ArticleService) Page(ctx context.Context, req *v1.PageArticleRequest) (rsp *v1.PageArticleReply, err error) {
	if req.Status != nil && *req.Status != int32(cv1.ArticleStatus_ArticleNormal) && *req.Status != int32(cv1.ArticleStatus_ArticleDrafts) {
		return nil, errors.New("status only be 0(normal) or 3(drafts)")
	}
	rsp, err = s.articleDomain.Page(ctx, req.Page, &repo.ArticleGetReq{
		TagId:    req.TagId,
		DomainId: req.DomainId,
		Status:   (*cv1.ArticleStatus)(req.Status),
		AuthorId: req.AuthorId,
		Order:    (*cv1.ArticleOrder)(req.Order),
		Type:     (*cv1.ArticleType)(req.Type),
		Keyword:  req.Keyword,
	})
	return rsp, err
}

func (s *ArticleService) GetOne(ctx context.Context, req *v1.GetArticleOneRequest) (rsp *v1.GetArticleOneReply, err error) {
	return s.articleDomain.GetOne(ctx, req.ArticleId)
}

func (s *ArticleService) Like(ctx context.Context, req *v1.LikeArticleRequest) (rsp *v1.LikeArticleReply, err error) {
	user := util.MustGetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, cv1.ArticleAction_ArticleActionLike, req.Active)
	return &v1.LikeArticleReply{}, err
}

func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticleRequest) (rsp *v1.PublishArticleReply, err error) {
	user := util.MustGetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	article, err := s.articleRepo.GetById(ctx, s.db, req.ArticleId)
	if err != nil {
		return nil, err
	}
	// 只能发布草稿
	if article.Status != int32(cv1.ArticleStatus_ArticleDrafts) {
		return nil, errors.New("only drafts can be publish")
	}
	// 只有作者可以发布草稿
	if *article.CreatedBy != user.ID {
		return nil, errors.New("you are not the author")
	}
	err = s.articleDomain.Publish(ctx, req.ArticleId)
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
