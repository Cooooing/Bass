package service

import (
	"common/api/gen/common"
	v1 "common/api/gen/content/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"content/internal/biz/domain"
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

	articleDomain *domain.ArticleDomain
	articleRepo   repo.ArticleRepo
}

func (s *ArticleService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentArticleServiceServer(gs, s)
}

func (s *ArticleService) RegisterHttp(hs *http.Server) {
	v1.RegisterContentArticleServiceHTTPServer(hs, s)
}

func NewArticleService(baseService *BaseService, articleDomain *domain.ArticleDomain, articleRepo repo.ArticleRepo) *ArticleService {
	return &ArticleService{
		BaseService:   baseService,
		articleDomain: articleDomain,
		articleRepo:   articleRepo,
	}
}

func (s *ArticleService) AddArticle(ctx context.Context, req *v1.AddArticle_Request) (rsp *v1.AddArticle_Reply, err error) {
	article := req.Article
	if article.Status != v1.ArticleStatus_ARTICLE_STATUS_NORMAL && article.Status != v1.ArticleStatus_ARTICLE_STATUS_DRAFTS {
		return nil, common.ErrorBadRequest("status only be 1(normal) or 4(drafts)")
	}
	if article.Type != v1.ArticleType_ARTICLE_TYPE_NORMAL && article.Type != v1.ArticleType_ARTICLE_TYPE_QA && article.Type != v1.ArticleType_ARTICLE_TYPE_VOTE && article.Type != v1.ArticleType_ARTICLE_TYPE_LOTTERY {
		return nil, common.ErrorBadRequest("type only be 1(normal), 2(QA), 3(vote), 4(lottery)")
	}
	if article.Type != v1.ArticleType_ARTICLE_TYPE_QA && article.BountyPoints != nil {
		return nil, common.ErrorBadRequest("bounty points only be set when type is 2(QA)")
	}

	var tags []*model.Tag
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			tags = append(tags, &model.Tag{Tag: &gen.Tag{
				Name:         tag.Name,
				Description:  tag.Description,
				Status:       int32(v1.TagStatus_TAG_STATUS_NORMAL),
				ArticleCount: 1,
			}})
		}
	}

	save, err := s.articleDomain.Add(ctx, &model.Article{Article: &gen.Article{
		Title:         article.Title,
		Content:       article.Content,
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Status:        int32(article.Status),
		Type:          int32(article.Type),
		BountyPoints:  util.If(article.Type != v1.ArticleType_ARTICLE_TYPE_QA, nil, article.BountyPoints),
		Statement:     article.Statement,
		Commentable:   util.DerefOrDefault(article.Commentable, true),
		Anonymous:     util.DerefOrDefault(article.Anonymous, false),
		Listable:      util.DerefOrDefault(article.Listable, true),
	}}, tags)
	if err != nil {
		return nil, err
	}
	return &v1.AddArticle_Reply{
		Article: save.ConvertToRpc(),
	}, nil
}

func (s *ArticleService) UpdateDraftArticle(ctx context.Context, req *v1.UpdateDraftArticle_Request) (rsp *v1.UpdateDraftArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	article := req.Article
	if article.Id == nil {
		return nil, common.ErrorBadRequest("article id is required")
	}
	if article.Status != v1.ArticleStatus_ARTICLE_STATUS_NORMAL && article.Status != v1.ArticleStatus_ARTICLE_STATUS_DRAFTS {
		return nil, common.ErrorBadRequest("status only be 1(normal) or 4(drafts)")
	}
	if article.Type != v1.ArticleType_ARTICLE_TYPE_NORMAL && article.Type != v1.ArticleType_ARTICLE_TYPE_QA && article.Type != v1.ArticleType_ARTICLE_TYPE_VOTE && article.Type != v1.ArticleType_ARTICLE_TYPE_LOTTERY {
		return nil, common.ErrorBadRequest("type only be 1(normal), 2(QA), 3(vote), 4(lottery)")
	}
	if article.Type != v1.ArticleType_ARTICLE_TYPE_QA && article.BountyPoints != nil {
		return nil, common.ErrorBadRequest("bounty points only be set when type is 2(QA)")
	}

	var tags []*model.Tag
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			tags = append(tags, &model.Tag{Tag: &gen.Tag{
				Name:         tag.Name,
				Description:  tag.Description,
				Status:       int32(v1.TagStatus_TAG_STATUS_NORMAL),
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
		Status:        int32(article.Status),
		Type:          int32(article.Type),
		BountyPoints:  util.If(article.Type != v1.ArticleType_ARTICLE_TYPE_QA, nil, article.BountyPoints),
		Statement:     article.Statement,
		Commentable:   util.DerefOrDefault(article.Commentable, true),
		Anonymous:     util.DerefOrDefault(article.Anonymous, false),
		Listable:      util.DerefOrDefault(article.Listable, true),
		CreatedBy:     new(user.ID),
	}}, tags)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDraftArticle_Reply{
		Article: update.ConvertToRpc(),
	}, nil
}

func (s *ArticleService) PublishArticle(ctx context.Context, req *v1.PublishArticle_Request) (rsp *v1.PublishArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	// 只有作者可以发布草稿
	exist, err := s.articleRepo.Exist(ctx, s.Db, &repo.ArticleGetReq{
		ArticleId: new(req.ArticleId),
		Status:    new(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
		CreatedBy: new(user.ID),
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, common.ErrorBadRequest("article not exist")
	}
	err = s.articleDomain.Publish(ctx, s.Db, req.ArticleId)
	return &v1.PublishArticle_Reply{}, err
}

func (s *ArticleService) AddPostscriptArticle(ctx context.Context, req *v1.AddPostscriptArticle_Request) (rsp *v1.AddPostscriptArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	// 只有作者可以添加附言
	exist, err := s.articleRepo.Exist(ctx, s.Db, &repo.ArticleGetReq{
		ArticleId: new(req.ArticleId),
		Status:    new(v1.ArticleStatus_ARTICLE_STATUS_NORMAL),
		CreatedBy: new(user.ID),
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, common.ErrorBadRequest("article not exist")
	}

	save, err := s.articleDomain.AddPostscript(ctx, req.ArticleId, req.Content)
	if err != nil {
		return nil, err
	}
	return &v1.AddPostscriptArticle_Reply{
		ArticlePostscript: save.ConvertToRpc(),
	}, err
}

func (s *ArticleService) UpdateArticleArticle(ctx context.Context, req *v1.UpdateArticleArticle_Request) (rsp *v1.UpdateArticleArticle_Reply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ArticleService) DeleteArticle(ctx context.Context, req *v1.DeleteArticle_Request) (rsp *v1.DeleteArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	err = ent.WithTx(ctx, s.Db, func(tx *gen.Client) error {
		// 只有作者可以删除草稿
		exist, err := s.articleRepo.Exist(ctx, s.Db, &repo.ArticleGetReq{
			ArticleId: new(req.ArticleId),
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
			CreatedBy: new(user.ID),
		})
		if err != nil {
			return err
		}
		if !exist {
			return common.ErrorBadRequest("article not exist")
		}
		err = s.articleRepo.Delete(ctx, s.Db, req.ArticleId)
		return err
	})
	return &v1.DeleteArticle_Reply{}, err
}

func (s *ArticleService) PageArticle(ctx context.Context, req *v1.PageArticle_Request) (rsp *v1.PageArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	req.Query = util.OrDefault(req.Query, &v1.ArticleQueryParams{})

	/*
	 * 正常状态的只能查看公开
	 * 草稿状态的只能查看自己
	 */

	if req.Query.Status != nil && *req.Query.Status != v1.ArticleStatus_ARTICLE_STATUS_NORMAL && *req.Query.Status != v1.ArticleStatus_ARTICLE_STATUS_DRAFTS {
		return nil, common.ErrorBadRequest("status only be 1(normal) or 4(drafts)")
	}
	status := v1.ArticleStatus_ARTICLE_STATUS_NORMAL
	authorId := req.Query.AuthorId
	if req.Query.Status != nil && *req.Query.Status == v1.ArticleStatus_ARTICLE_STATUS_DRAFTS {
		if !ok {
			return nil, common.ErrorUnauthorized("login required to view drafts")
		}
		authorId = &user.ID
		status = v1.ArticleStatus_ARTICLE_STATUS_DRAFTS
	}

	reply, page, err := s.articleDomain.Page(ctx, req.Page, &repo.ArticleGetReq{
		TagId:    req.Query.TagId,
		DomainId: req.Query.DomainId,
		Status:   &status,
		AuthorId: authorId,
		Order:    req.Query.Order,
		Type:     req.Query.Type,
		Keyword:  req.Query.Keyword,
		Listable: new(true),
	})
	return &v1.PageArticle_Reply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(reply),
	}, err
}

func (s *ArticleService) GetOneArticle(ctx context.Context, req *v1.GetOneArticle_Request) (rsp *v1.GetOneArticle_Reply, err error) {
	one, err := s.articleDomain.GetOne(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	return &v1.GetOneArticle_Reply{Article: one.ConvertToRpc()}, err
}

func (s *ArticleService) RewardArticle(ctx context.Context, req *v1.RewardArticle_Request) (rsp *v1.RewardArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ARTICLE_ACTION_REWARD, true)
	return &v1.RewardArticle_Reply{}, nil
}

func (s *ArticleService) LikeArticle(ctx context.Context, req *v1.LikeArticle_Request) (rsp *v1.LikeArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ARTICLE_ACTION_LIKE, req.Active)
	return &v1.LikeArticle_Reply{}, err
}

func (s *ArticleService) ThankArticle(ctx context.Context, req *v1.ThankArticle_Request) (rsp *v1.ThankArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ARTICLE_ACTION_THANK, req.Active)
	return &v1.ThankArticle_Reply{}, nil
}

func (s *ArticleService) CollectArticle(ctx context.Context, req *v1.CollectArticle_Request) (rsp *v1.CollectArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ARTICLE_ACTION_COLLECT, req.Active)
	return &v1.CollectArticle_Reply{}, err
}

func (s *ArticleService) WatchArticle(ctx context.Context, req *v1.WatchArticle_Request) (rsp *v1.WatchArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ARTICLE_ACTION_WATCH, req.Active)
	return &v1.WatchArticle_Reply{}, err
}

func (s *ArticleService) AcceptAnswerArticle(ctx context.Context, req *v1.AcceptAnswerArticle_Request) (rsp *v1.AcceptAnswerArticle_Reply, err error) {
	err = s.articleDomain.AcceptAnswer(ctx, req.ArticleId, req.CommentId)
	return &v1.AcceptAnswerArticle_Reply{}, err
}
