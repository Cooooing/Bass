package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"content/internal/data/client"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/biz/usecase"
	"content/internal/data/gen"
	articleent "content/internal/data/gen/article"
	tagent "content/internal/data/gen/tag"
	"content/internal/enum"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type ArticleService struct {
	v1.UnimplementedContentArticleServiceServer

	articleDomain *usecase.ArticleUsecase
	articleRepo   repo.ArticleRepo
	db            *gen.Client
}

func (s *ArticleService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentArticleServiceServer(gs, s)
}

func NewArticleService(
	articleDomain *usecase.ArticleUsecase,
	articleRepo repo.ArticleRepo,
	db *gen.Client,
) *ArticleService {
	return &ArticleService{
		articleDomain: articleDomain,
		articleRepo:   articleRepo,
		db:            db,
	}
}

func (s *ArticleService) AddArticle(ctx context.Context, req *v1.AddArticle_Request) (rsp *v1.AddArticle_Reply, err error) {
	article := req.Article
	if article.Status != v1.ArticleStatus_ARTICLE_STATUS_NORMAL && article.Status != v1.ArticleStatus_ARTICLE_STATUS_DRAFTS {
		return nil, cerrors.ErrorBadRequest("status only be 1(normal) or 4(drafts)")
	}
	if article.Type != v1.ArticleType_ARTICLE_TYPE_NORMAL && article.Type != v1.ArticleType_ARTICLE_TYPE_QA && article.Type != v1.ArticleType_ARTICLE_TYPE_VOTE && article.Type != v1.ArticleType_ARTICLE_TYPE_LOTTERY {
		return nil, cerrors.ErrorBadRequest("type only be 1(normal), 2(QA), 3(vote), 4(lottery)")
	}
	if article.Type != v1.ArticleType_ARTICLE_TYPE_QA && article.BountyPoints != nil {
		return nil, cerrors.ErrorBadRequest("bounty points only be set when type is 2(QA)")
	}

	var tags []*model.Tag
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			tags = append(tags, &model.Tag{Tag: &gen.Tag{
				Name:         tag.Name,
				Description:  tag.Description,
				Status:       tagent.StatusNormal,
				ArticleCount: 1,
			}})
		}
	}

	dbStatus, _ := enum.ArticleStatusMap.ToEnum(article.Status)
	dbType, _ := enum.ArticleTypeMap.ToEnum(article.Type)
	save, err := s.articleDomain.Add(ctx, &model.Article{Article: &gen.Article{
		Title:         article.Title,
		Content:       article.Content,
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Status:        articleent.Status(dbStatus),
		Type:          articleent.Type(dbType),
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

func (s *ArticleService) UpdateDraft(ctx context.Context, req *v1.UpdateDraftArticle_Request) (rsp *v1.UpdateDraftArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	article := req.Article
	if article.Id == nil {
		return nil, cerrors.ErrorBadRequest("article id is required")
	}
	if article.Status != v1.ArticleStatus_ARTICLE_STATUS_NORMAL && article.Status != v1.ArticleStatus_ARTICLE_STATUS_DRAFTS {
		return nil, cerrors.ErrorBadRequest("status only be 1(normal) or 4(drafts)")
	}
	if article.Type != v1.ArticleType_ARTICLE_TYPE_NORMAL && article.Type != v1.ArticleType_ARTICLE_TYPE_QA && article.Type != v1.ArticleType_ARTICLE_TYPE_VOTE && article.Type != v1.ArticleType_ARTICLE_TYPE_LOTTERY {
		return nil, cerrors.ErrorBadRequest("type only be 1(normal), 2(QA), 3(vote), 4(lottery)")
	}
	if article.Type != v1.ArticleType_ARTICLE_TYPE_QA && article.BountyPoints != nil {
		return nil, cerrors.ErrorBadRequest("bounty points only be set when type is 2(QA)")
	}

	var tags []*model.Tag
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			tags = append(tags, &model.Tag{Tag: &gen.Tag{
				Name:         tag.Name,
				Description:  tag.Description,
				Status:       tagent.StatusNormal,
				ArticleCount: 1,
			}})
		}
	}

	dbStatus2, _ := enum.ArticleStatusMap.ToEnum(article.Status)
	dbType2, _ := enum.ArticleTypeMap.ToEnum(article.Type)
	update, err := s.articleDomain.UpdateDraft(ctx, &model.Article{Article: &gen.Article{
		ID:            *req.Article.Id,
		Title:         article.Title,
		Content:       article.Content,
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Status:        articleent.Status(dbStatus2),
		Type:          articleent.Type(dbType2),
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

func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticle_Request) (rsp *v1.PublishArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	// 只有作者可以发布草稿
	exist, err := s.articleRepo.Exist(ctx, s.db, &repo.ArticleGetReq{
		ArticleId: new(req.ArticleId),
		Status:    new(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
		CreatedBy: new(user.ID),
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, cerrors.ErrorBadRequest("article not exist")
	}
	err = s.articleDomain.Publish(ctx, s.db, req.ArticleId)
	return &v1.PublishArticle_Reply{}, err
}

func (s *ArticleService) AddPostscript(ctx context.Context, req *v1.AddPostscriptArticle_Request) (rsp *v1.AddPostscriptArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	// 只有作者可以添加附言
	exist, err := s.articleRepo.Exist(ctx, s.db, &repo.ArticleGetReq{
		ArticleId: new(req.ArticleId),
		Status:    new(v1.ArticleStatus_ARTICLE_STATUS_NORMAL),
		CreatedBy: new(user.ID),
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, cerrors.ErrorBadRequest("article not exist")
	}

	save, err := s.articleDomain.AddPostscript(ctx, req.ArticleId, req.Content)
	if err != nil {
		return nil, err
	}
	return &v1.AddPostscriptArticle_Reply{
		ArticlePostscript: save.ConvertToRpc(),
	}, err
}

func (s *ArticleService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleArticle_Request) (rsp *v1.UpdateArticleArticle_Reply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ArticleService) Delete(ctx context.Context, req *v1.DeleteArticle_Request) (rsp *v1.DeleteArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err = client.WithTx(ctx, s.db, func(tx *gen.Client) error {
		// 只有作者可以删除草稿
		exist, err := s.articleRepo.Exist(ctx, s.db, &repo.ArticleGetReq{
			ArticleId: new(req.ArticleId),
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
			CreatedBy: new(user.ID),
		})
		if err != nil {
			return err
		}
		if !exist {
			return cerrors.ErrorBadRequest("article not exist")
		}
		err = s.articleRepo.Delete(ctx, s.db, req.ArticleId)
		return err
	})
	return &v1.DeleteArticle_Reply{}, err
}

func (s *ArticleService) Page(ctx context.Context, req *v1.PageArticle_Request) (rsp *v1.PageArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	req.Query = util.OrDefault(req.Query, &v1.ArticleQueryParams{})

	/*
	 * 正常状态的只能查看公开
	 * 草稿状态的只能查看自己
	 */

	if req.Query.Status != nil && *req.Query.Status != v1.ArticleStatus_ARTICLE_STATUS_NORMAL && *req.Query.Status != v1.ArticleStatus_ARTICLE_STATUS_DRAFTS {
		return nil, cerrors.ErrorBadRequest("status only be 1(normal) or 4(drafts)")
	}
	status := v1.ArticleStatus_ARTICLE_STATUS_NORMAL
	authorId := req.Query.AuthorId
	if req.Query.Status != nil && *req.Query.Status == v1.ArticleStatus_ARTICLE_STATUS_DRAFTS {
		if !ok {
			return nil, cerrors.ErrorUnauthorized("login required to view drafts")
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

func (s *ArticleService) GetOne(ctx context.Context, req *v1.GetOneArticle_Request) (rsp *v1.GetOneArticle_Reply, err error) {
	one, err := s.articleDomain.GetOne(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	return &v1.GetOneArticle_Reply{Article: one.ConvertToRpc()}, err
}

func (s *ArticleService) Reward(ctx context.Context, req *v1.RewardArticle_Request) (rsp *v1.RewardArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ARTICLE_ACTION_REWARD, true)
	return &v1.RewardArticle_Reply{}, nil
}

func (s *ArticleService) Like(ctx context.Context, req *v1.LikeArticle_Request) (rsp *v1.LikeArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ARTICLE_ACTION_LIKE, req.Active)
	return &v1.LikeArticle_Reply{}, err
}

func (s *ArticleService) Thank(ctx context.Context, req *v1.ThankArticle_Request) (rsp *v1.ThankArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ARTICLE_ACTION_THANK, req.Active)
	return &v1.ThankArticle_Reply{}, nil
}

func (s *ArticleService) Collect(ctx context.Context, req *v1.CollectArticle_Request) (rsp *v1.CollectArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ARTICLE_ACTION_COLLECT, req.Active)
	return &v1.CollectArticle_Reply{}, err
}

func (s *ArticleService) Watch(ctx context.Context, req *v1.WatchArticle_Request) (rsp *v1.WatchArticle_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err = s.articleDomain.Action(ctx, req.ArticleId, user.ID, v1.ArticleAction_ARTICLE_ACTION_WATCH, req.Active)
	return &v1.WatchArticle_Reply{}, err
}

func (s *ArticleService) AcceptAnswer(ctx context.Context, req *v1.AcceptAnswerArticle_Request) (rsp *v1.AcceptAnswerArticle_Reply, err error) {
	err = s.articleDomain.AcceptAnswer(ctx, req.ArticleId, req.CommentId)
	return &v1.AcceptAnswerArticle_Reply{}, err
}
