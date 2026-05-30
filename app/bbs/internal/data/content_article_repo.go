package data

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	contentv1 "common/api/gen/content/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.ContentArticleRepo = (*ContentArticleRepo)(nil)

type ContentArticleRepo struct {
	contentClient *rpc.ContentClient
}

func NewContentArticleRepo(contentClient *rpc.ContentClient) repo.ContentArticleRepo {
	return &ContentArticleRepo{contentClient: contentClient}
}

func (r *ContentArticleRepo) CreateArticle(ctx context.Context, req *bbscontentv1.CreateArticle_Request) (*bbscontentv1.CreateArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	article := req.GetArticle()
	save := &contentv1.ArticleSave{}
	if article != nil {
		save.Title = article.GetTitle()
		save.Content = article.GetContent()
		save.RewardContent = article.RewardContent
		save.RewardPoints = article.RewardPoints
		save.Status = contentv1.ArticleStatus(article.GetStatus())
		save.Type = contentv1.ArticleType(article.GetType())
		save.BountyPoints = article.BountyPoints
		save.Statement = article.Statement
		save.Commentable = article.Commentable
		save.Anonymous = article.Anonymous
		save.Listable = article.Listable
		save.Tags = make([]*contentv1.TagSave, 0, len(article.GetTags()))
		for _, tag := range article.GetTags() {
			item := &contentv1.TagSave{
				Id:          tag.Id,
				DomainId:    tag.DomainId,
				Name:        tag.GetName(),
				Description: tag.Description,
			}
			if tag.Status != nil {
				item.Status = new(contentv1.TagStatus(*tag.Status))
			}
			save.Tags = append(save.Tags, item)
		}
	}
	reply, err := r.contentClient.Article.Create(ctx, &contentv1.CreateArticle_Request{Article: save, UserId: userID})
	if err != nil {
		return nil, err
	}
	item := reply.GetArticle()
	out := &bbscontentv1.Article{
		Id:                  item.GetId(),
		Title:               item.GetTitle(),
		Content:             item.GetContent(),
		ContentRender:       item.GetContentRender(),
		HasPostscript:       item.GetHasPostscript(),
		HasReward:           item.GetHasReward(),
		RewardContent:       item.RewardContent,
		RewardContentRender: item.RewardContentRender,
		RewardPoints:        item.RewardPoints,
		Status:              bbscontentv1.ArticleStatus(item.GetStatus()),
		Type:                bbscontentv1.ArticleType(item.GetType()),
		Statement:           item.Statement,
		Commentable:         item.GetCommentable(),
		Anonymous:           item.GetAnonymous(),
		Listable:            item.GetListable(),
		ViewCount:           item.GetViewCount(),
		ThankCount:          item.GetThankCount(),
		LikeCount:           item.GetLikeCount(),
		CollectCount:        item.GetCollectCount(),
		WatchCount:          item.GetWatchCount(),
		ReplyCount:          item.GetReplyCount(),
		BountyPoints:        item.BountyPoints,
		AcceptedAnswerId:    item.AcceptedAnswerId,
		CoverImageUrl:       item.CoverImageUrl,
		CreatedBy:           item.CreatedBy,
		UpdatedBy:           item.UpdatedBy,
		CreatedAt:           formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:           formatProtoTime(item.GetUpdatedAt()),
		LastReplyAt:         formatProtoTime(item.GetLastReplyAt()),
	}
	if author := item.GetAuthorUser(); author != nil {
		out.AuthorUser = &bbsuserv1.AccountProfile{
			Id:            author.GetId(),
			Name:          author.GetName(),
			Nickname:      author.Nickname,
			Url:           author.Url,
			AvatarUrl:     author.AvatarUrl,
			Introduction:  author.Introduction,
			Mbti:          bbsuserv1.MBTI(author.GetMbti()),
			Status:        bbsuserv1.AccountStatus(author.GetStatus()),
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			CreatedAt:     formatProtoTime(author.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(author.GetUpdatedAt()),
		}
	}
	return &bbscontentv1.CreateArticle_Reply{Article: out}, nil
}

func (r *ContentArticleRepo) UpdateDraftArticle(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Request) (*bbscontentv1.UpdateDraftArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	article := req.GetArticle()
	save := &contentv1.ArticleSave{}
	if article != nil {
		save.Id = article.Id
		save.Title = article.GetTitle()
		save.Content = article.GetContent()
		save.RewardContent = article.RewardContent
		save.RewardPoints = article.RewardPoints
		save.Status = contentv1.ArticleStatus(article.GetStatus())
		save.Type = contentv1.ArticleType(article.GetType())
		save.BountyPoints = article.BountyPoints
		save.Statement = article.Statement
		save.Commentable = article.Commentable
		save.Anonymous = article.Anonymous
		save.Listable = article.Listable
		for _, tag := range article.GetTags() {
			item := &contentv1.TagSave{
				Id:          tag.Id,
				DomainId:    tag.DomainId,
				Name:        tag.GetName(),
				Description: tag.Description,
			}
			if tag.Status != nil {
				item.Status = new(contentv1.TagStatus(*tag.Status))
			}
			save.Tags = append(save.Tags, item)
		}
	}
	reply, err := r.contentClient.Article.UpdateDraft(ctx, &contentv1.UpdateDraftArticle_Request{Article: save, UserId: userID})
	if err != nil {
		return nil, err
	}
	item := reply.GetArticle()
	out := &bbscontentv1.Article{
		Id:                  item.GetId(),
		Title:               item.GetTitle(),
		Content:             item.GetContent(),
		ContentRender:       item.GetContentRender(),
		HasPostscript:       item.GetHasPostscript(),
		HasReward:           item.GetHasReward(),
		RewardContent:       item.RewardContent,
		RewardContentRender: item.RewardContentRender,
		RewardPoints:        item.RewardPoints,
		Status:              bbscontentv1.ArticleStatus(item.GetStatus()),
		Type:                bbscontentv1.ArticleType(item.GetType()),
		Statement:           item.Statement,
		Commentable:         item.GetCommentable(),
		Anonymous:           item.GetAnonymous(),
		Listable:            item.GetListable(),
		ViewCount:           item.GetViewCount(),
		ThankCount:          item.GetThankCount(),
		LikeCount:           item.GetLikeCount(),
		CollectCount:        item.GetCollectCount(),
		WatchCount:          item.GetWatchCount(),
		ReplyCount:          item.GetReplyCount(),
		BountyPoints:        item.BountyPoints,
		AcceptedAnswerId:    item.AcceptedAnswerId,
		CoverImageUrl:       item.CoverImageUrl,
		CreatedBy:           item.CreatedBy,
		UpdatedBy:           item.UpdatedBy,
		CreatedAt:           formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:           formatProtoTime(item.GetUpdatedAt()),
		LastReplyAt:         formatProtoTime(item.GetLastReplyAt()),
	}
	return &bbscontentv1.UpdateDraftArticle_Reply{Article: out}, nil
}

func (r *ContentArticleRepo) PublishArticle(ctx context.Context, req *bbscontentv1.PublishArticle_Request) (*bbscontentv1.PublishArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.Publish(ctx, &contentv1.PublishArticle_Request{ArticleId: req.GetArticleId(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.PublishArticle_Reply{}, nil
}

func (r *ContentArticleRepo) DeleteArticle(ctx context.Context, req *bbscontentv1.DeleteArticle_Request) (*bbscontentv1.DeleteArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.Delete(ctx, &contentv1.DeleteArticle_Request{ArticleId: req.GetArticleId(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.DeleteArticle_Reply{}, nil
}

func (r *ContentArticleRepo) ListArticles(ctx context.Context, req *bbscontentv1.ListArticles_Request) (*bbscontentv1.ListArticles_Reply, error) {
	query := req.GetQuery()
	if query == nil {
		query = &bbscontentv1.ArticleQuery{}
	}
	contentQuery := &contentv1.ArticleQueryParams{
		TagId:    query.TagId,
		DomainId: query.DomainId,
		Keyword:  query.Keyword,
		AuthorId: query.AuthorId,
		Listable: query.Listable,
	}
	if query.Status != nil {
		contentQuery.Status = new(contentv1.ArticleStatus(*query.Status))
	}
	if query.Type != nil {
		contentQuery.Type = new(contentv1.ArticleType(*query.Type))
	}
	if query.Order != nil {
		contentQuery.Order = new(contentv1.ArticleOrder(*query.Order))
	}
	reply, err := r.contentClient.Article.List(ctx, &contentv1.ListArticles_Request{
		Page:  req.GetPage(),
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.Article, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		row := &bbscontentv1.Article{
			Id:                  item.GetId(),
			Title:               item.GetTitle(),
			Content:             item.GetContent(),
			ContentRender:       item.GetContentRender(),
			HasPostscript:       item.GetHasPostscript(),
			HasReward:           item.GetHasReward(),
			RewardContent:       item.RewardContent,
			RewardContentRender: item.RewardContentRender,
			RewardPoints:        item.RewardPoints,
			Status:              bbscontentv1.ArticleStatus(item.GetStatus()),
			Type:                bbscontentv1.ArticleType(item.GetType()),
			Statement:           item.Statement,
			Commentable:         item.GetCommentable(),
			Anonymous:           item.GetAnonymous(),
			Listable:            item.GetListable(),
			ViewCount:           item.GetViewCount(),
			ThankCount:          item.GetThankCount(),
			LikeCount:           item.GetLikeCount(),
			CollectCount:        item.GetCollectCount(),
			WatchCount:          item.GetWatchCount(),
			ReplyCount:          item.GetReplyCount(),
			BountyPoints:        item.BountyPoints,
			AcceptedAnswerId:    item.AcceptedAnswerId,
			CoverImageUrl:       item.CoverImageUrl,
			CreatedBy:           item.CreatedBy,
			UpdatedBy:           item.UpdatedBy,
			CreatedAt:           formatProtoTime(item.GetCreatedAt()),
			UpdatedAt:           formatProtoTime(item.GetUpdatedAt()),
			LastReplyAt:         formatProtoTime(item.GetLastReplyAt()),
		}
		if author := item.GetAuthorUser(); author != nil {
			row.AuthorUser = &bbsuserv1.AccountProfile{Id: author.GetId(), Name: author.GetName(), Nickname: author.Nickname, AvatarUrl: author.AvatarUrl}
		}
		if lastReplyUser := item.GetLastReplyUser(); lastReplyUser != nil {
			row.LastReplyUser = &bbsuserv1.AccountProfile{Id: lastReplyUser.GetId(), Name: lastReplyUser.GetName(), Nickname: lastReplyUser.Nickname, AvatarUrl: lastReplyUser.AvatarUrl}
		}
		rows = append(rows, row)
	}
	return &bbscontentv1.ListArticles_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *ContentArticleRepo) GetArticle(ctx context.Context, req *bbscontentv1.GetArticle_Request) (*bbscontentv1.GetArticle_Reply, error) {
	reply, err := r.contentClient.Article.Get(ctx, &contentv1.GetArticle_Request{ArticleId: req.GetArticleId()})
	if err != nil {
		return nil, err
	}
	item := reply.GetArticle()
	out := &bbscontentv1.Article{
		Id:                  item.GetId(),
		Title:               item.GetTitle(),
		Content:             item.GetContent(),
		ContentRender:       item.GetContentRender(),
		HasPostscript:       item.GetHasPostscript(),
		HasReward:           item.GetHasReward(),
		RewardContent:       item.RewardContent,
		RewardContentRender: item.RewardContentRender,
		RewardPoints:        item.RewardPoints,
		Status:              bbscontentv1.ArticleStatus(item.GetStatus()),
		Type:                bbscontentv1.ArticleType(item.GetType()),
		Statement:           item.Statement,
		Commentable:         item.GetCommentable(),
		Anonymous:           item.GetAnonymous(),
		Listable:            item.GetListable(),
		ViewCount:           item.GetViewCount(),
		ThankCount:          item.GetThankCount(),
		LikeCount:           item.GetLikeCount(),
		CollectCount:        item.GetCollectCount(),
		WatchCount:          item.GetWatchCount(),
		ReplyCount:          item.GetReplyCount(),
		BountyPoints:        item.BountyPoints,
		AcceptedAnswerId:    item.AcceptedAnswerId,
		CoverImageUrl:       item.CoverImageUrl,
		CreatedBy:           item.CreatedBy,
		UpdatedBy:           item.UpdatedBy,
		CreatedAt:           formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:           formatProtoTime(item.GetUpdatedAt()),
		LastReplyAt:         formatProtoTime(item.GetLastReplyAt()),
	}
	if author := item.GetAuthorUser(); author != nil {
		out.AuthorUser = &bbsuserv1.AccountProfile{Id: author.GetId(), Name: author.GetName(), Nickname: author.Nickname, AvatarUrl: author.AvatarUrl}
	}
	return &bbscontentv1.GetArticle_Reply{Article: out}, nil
}

func (r *ContentArticleRepo) LikeArticle(ctx context.Context, req *bbscontentv1.LikeArticle_Request) (*bbscontentv1.LikeArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.Like(ctx, &contentv1.LikeArticle_Request{ArticleId: req.GetArticleId(), Active: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.LikeArticle_Reply{}, nil
}

func (r *ContentArticleRepo) ThankArticle(ctx context.Context, req *bbscontentv1.ThankArticle_Request) (*bbscontentv1.ThankArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.Thank(ctx, &contentv1.ThankArticle_Request{ArticleId: req.GetArticleId(), Active: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.ThankArticle_Reply{}, nil
}

func (r *ContentArticleRepo) CollectArticle(ctx context.Context, req *bbscontentv1.CollectArticle_Request) (*bbscontentv1.CollectArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.Collect(ctx, &contentv1.CollectArticle_Request{ArticleId: req.GetArticleId(), Active: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.CollectArticle_Reply{}, nil
}

func (r *ContentArticleRepo) WatchArticle(ctx context.Context, req *bbscontentv1.WatchArticle_Request) (*bbscontentv1.WatchArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.Watch(ctx, &contentv1.WatchArticle_Request{ArticleId: req.GetArticleId(), Active: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.WatchArticle_Reply{}, nil
}

func (r *ContentArticleRepo) RewardArticle(ctx context.Context, req *bbscontentv1.RewardArticle_Request) (*bbscontentv1.RewardArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.Reward(ctx, &contentv1.RewardArticle_Request{ArticleId: req.GetArticleId(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.RewardArticle_Reply{}, nil
}

func (r *ContentArticleRepo) AcceptAnswerArticle(ctx context.Context, req *bbscontentv1.AcceptAnswerArticle_Request) (*bbscontentv1.AcceptAnswerArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.AcceptAnswer(ctx, &contentv1.AcceptAnswerArticle_Request{ArticleId: req.GetArticleId(), CommentId: req.GetCommentId(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.AcceptAnswerArticle_Reply{}, nil
}
