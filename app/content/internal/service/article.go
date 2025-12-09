package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"common/pkg/constant"
	"common/pkg/cutil/base"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"content/internal/biz"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"

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

func (s *ArticleService) Add(ctx context.Context, req *v1.AddArticleRequest) (rsp *v1.AddArticleReply, err error) {
	article := req.Article
	if article.Status != int32(v1.ArticleStatus_ArticleNormal) && article.Status != int32(v1.ArticleStatus_ArticleDrafts) {
		return nil, cv1.ErrorBadRequest("status only be 0(normal) or 3(drafts)")
	}
	if article.Type != int32(v1.ArticleType_ArticleTypeNormal) && article.Type != int32(v1.ArticleType_ArticleTypeQA) && article.Type != int32(v1.ArticleType_ArticleTypeVote) && article.Type != int32(v1.ArticleType_ArticleTypeLottery) {
		return nil, cv1.ErrorBadRequest("type only be 0(normal), 1(QA), 2(vote), 3(lottery)")
	}

	var tags []*model.Tag
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			tags = append(tags, &model.Tag{Tag: &gen.Tag{
				Name:         tag.Name,
				Description:  tag.Description,
				Status:       int32(v1.TagStatus_TagNormal),
				ArticleCount: 1,
			}})
		}
	}

	save, err := s.articleDomain.Add(ctx, &model.Article{Article: &gen.Article{
		Title:         article.Title,
		Content:       article.Content,
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Status:        article.Status,
		Type:          article.Type,
		BountyPoints:  base.If(article.Type != int32(v1.ArticleType_ArticleTypeQA), nil, article.BountyPoints),
		Statement:     article.Statement,
		Commentable:   base.DerefOrDefault(article.Commentable, true),
		Anonymous:     base.DerefOrDefault(article.Anonymous, false),
		Listable:      base.DerefOrDefault(article.Listable, true),
	}}, tags)
	if err != nil {
		return nil, err
	}
	return &v1.AddArticleReply{
		Article: save.ConvertToRpc(),
	}, nil
}

func (s *ArticleService) UpdateDraft(ctx context.Context, req *v1.UpdateArticleDraftRequest) (rsp *v1.UpdateArticleDraftReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	article := req.Article
	if article.Id == nil {
		return nil, cv1.ErrorBadRequest("article id is required")
	}
	if article.Status != int32(v1.ArticleStatus_ArticleNormal) && article.Status != int32(v1.ArticleStatus_ArticleDrafts) {
		return nil, cv1.ErrorBadRequest("status only be 0(normal) or 3(drafts)")
	}
	if article.Type != int32(v1.ArticleType_ArticleTypeNormal) && article.Type != int32(v1.ArticleType_ArticleTypeQA) && article.Type != int32(v1.ArticleType_ArticleTypeVote) && article.Type != int32(v1.ArticleType_ArticleTypeLottery) {
		return nil, cv1.ErrorBadRequest("type only be 0(normal), 1(QA), 2(vote), 3(lottery)")
	}

	var tags []*model.Tag
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			tags = append(tags, &model.Tag{Tag: &gen.Tag{
				Name:         tag.Name,
				Description:  tag.Description,
				Status:       int32(v1.TagStatus_TagNormal),
				ArticleCount: 1,
			}})
		}
	}

	update, err := s.articleDomain.UpdateDraft(ctx, &model.Article{Article: &gen.Article{
		ID:            *req.Article.Id,
		Title:         article.Title,
		Content:       article.Content,
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Status:        article.Status,
		Type:          article.Type,
		BountyPoints:  base.If(article.Type != int32(v1.ArticleType_ArticleTypeQA), nil, article.BountyPoints),
		Statement:     article.Statement,
		Commentable:   base.DerefOrDefault(article.Commentable, true),
		Anonymous:     base.DerefOrDefault(article.Anonymous, false),
		Listable:      base.DerefOrDefault(article.Listable, true),
		CreatedBy:     base.Ptr(user.ID),
	}}, tags)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateArticleDraftReply{
		Article: update.ConvertToRpc(),
	}, nil
}

func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticleRequest) (rsp *v1.PublishArticleReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	// 只有作者可以发布草稿
	exist, err := s.articleRepo.Exist(ctx, s.db, &repo.ArticleGetReq{
		ArticleId: base.Ptr(req.ArticleId),
		Status:    base.Ptr(v1.ArticleStatus_ArticleDrafts),
		CreatedBy: base.Ptr(user.ID),
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, cv1.ErrorBadRequest("article not exist")
	}
	err = s.articleDomain.Publish(ctx, req.ArticleId)
	return &v1.PublishArticleReply{}, err
}

func (s *ArticleService) Update(ctx context.Context, req *v1.UpdateArticleRequest) (rsp *v1.UpdateArticleReply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ArticleService) Delete(ctx context.Context, req *v1.DeleteArticleRequest) (rsp *v1.DeleteArticleReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	err = ent.WithTx(ctx, s.db, func(tx *gen.Client) error {
		// 只有作者可以删除草稿
		exist, err := s.articleRepo.Exist(ctx, s.db, &repo.ArticleGetReq{
			ArticleId: base.Ptr(req.ArticleId),
			Status:    base.Ptr(v1.ArticleStatus_ArticleDrafts),
			CreatedBy: base.Ptr(user.ID),
		})
		if err != nil {
			return err
		}
		if !exist {
			return cv1.ErrorBadRequest("article not exist")
		}
		err = s.articleRepo.Delete(ctx, s.db, req.ArticleId)
		return err
	})
	return &v1.DeleteArticleReply{}, err
}

func (s *ArticleService) Page(ctx context.Context, req *v1.PageArticleRequest) (rsp *v1.PageArticleReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	req.Query = base.OrDefault(req.Query, &v1.ArticleQueryParams{})

	/*
	 * 正常状态的只能查看公开
	 * 草稿状态的只能查看自己
	 */

	if req.Query.Status != nil && *req.Query.Status != int32(v1.ArticleStatus_ArticleNormal) && *req.Query.Status != int32(v1.ArticleStatus_ArticleDrafts) {
		return nil, cv1.ErrorBadRequest("status only be 0(normal) or 3(drafts)")
	}
	status := base.Ptr(v1.ArticleStatus_ArticleNormal)
	authorId := req.Query.AuthorId
	if req.Query.Status != nil && *req.Query.Status == int32(v1.ArticleStatus_ArticleDrafts) {
		if !ok {
			return nil, cv1.ErrorUnauthorized("login required to view drafts")
		}
		authorId = &user.ID
		status = base.Ptr(v1.ArticleStatus_ArticleDrafts)
	}

	reply, page, err := s.articleDomain.Page(ctx, req.Page, &repo.ArticleGetReq{
		TagId:    req.Query.TagId,
		DomainId: req.Query.DomainId,
		Status:   status,
		AuthorId: authorId,
		Order:    (*v1.ArticleOrder)(req.Query.Order),
		Type:     (*v1.ArticleType)(req.Query.Type),
		Keyword:  req.Query.Keyword,
		Listable: base.Ptr(true),
	})
	return &v1.PageArticleReply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(reply),
	}, err
}

func (s *ArticleService) GetOne(ctx context.Context, req *v1.GetArticleOneRequest) (rsp *v1.GetArticleOneReply, err error) {
	one, err := s.articleDomain.GetOne(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	return &v1.GetArticleOneReply{Article: one.ConvertToRpc()}, err
}

func (s *ArticleService) AddPostscript(ctx context.Context, req *v1.AddPostscriptArticleRequest) (rsp *v1.AddPostscriptArticleReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	// 只有作者可以添加附言
	exist, err := s.articleRepo.Exist(ctx, s.db, &repo.ArticleGetReq{
		ArticleId: base.Ptr(req.ArticleId),
		Status:    base.Ptr(v1.ArticleStatus_ArticleNormal),
		CreatedBy: base.Ptr(user.ID),
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, cv1.ErrorBadRequest("article not exist")
	}

	save, err := s.articleDomain.AddPostscript(ctx, req.ArticleId, req.Content)
	if err != nil {
		return nil, err
	}
	return &v1.AddPostscriptArticleReply{
		ArticlePostscript: save.ConvertToRpc(),
	}, err
}

func (s *ArticleService) Reward(ctx context.Context, req *v1.RewardArticleRequest) (rsp *v1.RewardArticleReply, err error) {
	return &v1.RewardArticleReply{}, nil
}

func (s *ArticleService) Thank(ctx context.Context, req *v1.ThankArticleRequest) (rsp *v1.ThankArticleReply, err error) {
	return &v1.ThankArticleReply{}, nil
}

func (s *ArticleService) Like(ctx context.Context, req *v1.LikeArticleRequest) (rsp *v1.LikeArticleReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ArticleActionLike, req.Active)
	return &v1.LikeArticleReply{}, err
}

func (s *ArticleService) Collect(ctx context.Context, req *v1.CollectArticleRequest) (rsp *v1.CollectArticleReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ArticleActionCollect, req.Active)
	return &v1.CollectArticleReply{}, err
}

func (s *ArticleService) Watch(ctx context.Context, req *v1.WatchArticleRequest) (rsp *v1.WatchArticleReply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ArticleService) AcceptAnswer(ctx context.Context, req *v1.AcceptAnswerArticleRequest) (rsp *v1.AcceptAnswerArticleReply, err error) {
	// TODO implement me
	panic("implement me")
}
