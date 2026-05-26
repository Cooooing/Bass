package data

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	contentv1 "common/api/gen/content/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.ContentRepo = (*ContentRepo)(nil)

type ContentRepo struct {
	contentClient *rpc.ContentClient
}

func NewContentRepo(contentClient *rpc.ContentClient) repo.ContentRepo {
	return &ContentRepo{contentClient: contentClient}
}

func (r *ContentRepo) CreateArticle(ctx context.Context, req *bbscontentv1.CreateArticle_Request) (*bbscontentv1.CreateArticle_Reply, error) {
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
	for _, tag := range item.GetTags() {
		out.Tags = append(out.Tags, &bbscontentv1.Tag{
			Id:          tag.GetId(),
			Name:        tag.GetName(),
			Description: tag.Description,
			DomainId:    tag.DomainId,
			Status:      new(bbscontentv1.TagStatus(tag.GetStatus())),
			CreatedBy:   tag.CreatedBy,
			UpdatedBy:   tag.UpdatedBy,
			CreatedAt:   formatProtoTime(tag.GetCreatedAt()),
			UpdatedAt:   formatProtoTime(tag.GetUpdatedAt()),
		})
	}
	return &bbscontentv1.CreateArticle_Reply{Article: out}, nil
}

func (r *ContentRepo) UpdateDraftArticle(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Request) (*bbscontentv1.UpdateDraftArticle_Reply, error) {
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

func (r *ContentRepo) PublishArticle(ctx context.Context, req *bbscontentv1.PublishArticle_Request) (*bbscontentv1.PublishArticle_Reply, error) {
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

func (r *ContentRepo) DeleteArticle(ctx context.Context, req *bbscontentv1.DeleteArticle_Request) (*bbscontentv1.DeleteArticle_Reply, error) {
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

func (r *ContentRepo) ListArticles(ctx context.Context, req *bbscontentv1.ListArticles_Request) (*bbscontentv1.ListArticles_Reply, error) {
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

func (r *ContentRepo) GetArticle(ctx context.Context, req *bbscontentv1.GetArticle_Request) (*bbscontentv1.GetArticle_Reply, error) {
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
	for _, tag := range item.GetTags() {
		out.Tags = append(out.Tags, &bbscontentv1.Tag{
			Id:          tag.GetId(),
			Name:        tag.GetName(),
			Description: tag.Description,
			DomainId:    tag.DomainId,
			Status:      new(bbscontentv1.TagStatus(tag.GetStatus())),
			CreatedAt:   formatProtoTime(tag.GetCreatedAt()),
			UpdatedAt:   formatProtoTime(tag.GetUpdatedAt()),
		})
	}
	for _, postscript := range item.GetPostscripts() {
		out.Postscripts = append(out.Postscripts, &bbscontentv1.ArticlePostscript{
			Id:            postscript.GetId(),
			ArticleId:     postscript.GetArticleId(),
			Content:       postscript.GetContent(),
			ContentRender: postscript.GetContentRender(),
			CreatedBy:     postscript.CreatedBy,
			UpdatedBy:     postscript.UpdatedBy,
			CreatedAt:     formatProtoTime(postscript.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(postscript.GetUpdatedAt()),
		})
	}
	return &bbscontentv1.GetArticle_Reply{Article: out}, nil
}

func (r *ContentRepo) AddPostscript(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Article.AddPostscript(ctx, &contentv1.AddPostscriptArticle_Request{ArticleId: req.GetArticleId(), Content: req.GetContent(), UserId: userID})
	if err != nil {
		return nil, err
	}
	item := reply.GetArticlePostscript()
	return &bbscontentv1.AddPostscript_Reply{Postscript: &bbscontentv1.ArticlePostscript{
		Id:            item.GetId(),
		ArticleId:     item.GetArticleId(),
		Content:       item.GetContent(),
		ContentRender: item.GetContentRender(),
		CreatedBy:     item.CreatedBy,
		UpdatedBy:     item.UpdatedBy,
		CreatedAt:     formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:     formatProtoTime(item.GetUpdatedAt()),
	}}, nil
}

func (r *ContentRepo) LikeArticle(ctx context.Context, req *bbscontentv1.LikeArticle_Request) (*bbscontentv1.LikeArticle_Reply, error) {
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

func (r *ContentRepo) ThankArticle(ctx context.Context, req *bbscontentv1.ThankArticle_Request) (*bbscontentv1.ThankArticle_Reply, error) {
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

func (r *ContentRepo) CollectArticle(ctx context.Context, req *bbscontentv1.CollectArticle_Request) (*bbscontentv1.CollectArticle_Reply, error) {
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

func (r *ContentRepo) WatchArticle(ctx context.Context, req *bbscontentv1.WatchArticle_Request) (*bbscontentv1.WatchArticle_Reply, error) {
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

func (r *ContentRepo) RewardArticle(ctx context.Context, req *bbscontentv1.RewardArticle_Request) (*bbscontentv1.RewardArticle_Reply, error) {
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

func (r *ContentRepo) AcceptAnswerArticle(ctx context.Context, req *bbscontentv1.AcceptAnswerArticle_Request) (*bbscontentv1.AcceptAnswerArticle_Reply, error) {
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

func (r *ContentRepo) CreateComment(ctx context.Context, req *bbscontentv1.CreateComment_Request) (*bbscontentv1.CreateComment_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Comment.Create(ctx, &contentv1.CreateComment_Request{ArticleId: req.GetArticleId(), Content: req.GetContent(), ReplyId: req.GetReplyId(), UserId: userID})
	if err != nil {
		return nil, err
	}
	item := reply.GetComment()
	return &bbscontentv1.CreateComment_Reply{Comment: &bbscontentv1.Comment{
		Id:            item.GetId(),
		ArticleId:     item.GetArticleId(),
		Content:       item.GetContent(),
		ContentRender: item.GetContentRender(),
		Level:         item.GetLevel(),
		ParentId:      item.ParentId,
		ReplyId:       item.ReplyId,
		Status:        bbscontentv1.CommentStatus(item.GetStatus()),
		ThankCount:    item.GetThankCount(),
		LikeCount:     item.GetLikeCount(),
		ReplyCount:    item.GetReplyCount(),
		CreatedBy:     item.CreatedBy,
		UpdatedBy:     item.UpdatedBy,
		CreatedAt:     formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:     formatProtoTime(item.GetUpdatedAt()),
	}}, nil
}

func (r *ContentRepo) ListComments(ctx context.Context, req *bbscontentv1.ListComments_Request) (*bbscontentv1.ListComments_Reply, error) {
	query := req.GetQuery()
	if query == nil {
		query = &bbscontentv1.CommentQuery{}
	}
	contentQuery := &contentv1.CommentQueryParams{
		CommentId:   query.CommentId,
		ArticleId:   query.ArticleId,
		ParentId:    query.ParentId,
		ReplyId:     query.ReplyId,
		UserId:      query.UserId,
		Level:       query.Level,
		WithArticle: query.GetWithArticle(),
	}
	if query.Order != nil {
		contentQuery.Order = new(contentv1.CommentOrder(*query.Order))
	}
	if query.Status != nil {
		contentQuery.Status = new(contentv1.CommentStatus(*query.Status))
	}
	reply, err := r.contentClient.Comment.List(ctx, &contentv1.ListComments_Request{
		Page:  req.Page,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.Comment, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		row := &bbscontentv1.Comment{
			Id:            item.GetId(),
			ArticleId:     item.GetArticleId(),
			Content:       item.GetContent(),
			ContentRender: item.GetContentRender(),
			Level:         item.GetLevel(),
			ParentId:      item.ParentId,
			ReplyId:       item.ReplyId,
			Status:        bbscontentv1.CommentStatus(item.GetStatus()),
			ThankCount:    item.GetThankCount(),
			LikeCount:     item.GetLikeCount(),
			ReplyCount:    item.GetReplyCount(),
			CreatedBy:     item.CreatedBy,
			UpdatedBy:     item.UpdatedBy,
			CreatedAt:     formatProtoTime(item.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(item.GetUpdatedAt()),
		}
		if user := item.GetUser(); user != nil {
			row.User = &bbsuserv1.AccountProfile{Id: user.GetId(), Name: user.GetName(), Nickname: user.Nickname, AvatarUrl: user.AvatarUrl}
		}
		if replyUser := item.GetReplyUser(); replyUser != nil {
			row.ReplyUser = &bbsuserv1.AccountProfile{Id: replyUser.GetId(), Name: replyUser.GetName(), Nickname: replyUser.Nickname, AvatarUrl: replyUser.AvatarUrl}
		}
		rows = append(rows, row)
	}
	return &bbscontentv1.ListComments_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *ContentRepo) LikeComment(ctx context.Context, req *bbscontentv1.LikeComment_Request) (*bbscontentv1.LikeComment_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Comment.Like(ctx, &contentv1.LikeComment_Request{Id: req.GetId(), Active: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.LikeComment_Reply{}, nil
}

func (r *ContentRepo) ThankComment(ctx context.Context, req *bbscontentv1.ThankComment_Request) (*bbscontentv1.ThankComment_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Comment.Thank(ctx, &contentv1.ThankComment_Request{Id: req.GetId(), Active: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.ThankComment_Reply{}, nil
}

func (r *ContentRepo) ListDomains(ctx context.Context, req *bbscontentv1.ListDomains_Request) (*bbscontentv1.ListDomains_Reply, error) {
	query := req.GetQuery()
	if query == nil {
		query = &bbscontentv1.DomainQuery{}
	}
	contentQuery := &contentv1.DomainQueryParams{
		Ids:         query.GetIds(),
		Name:        query.Name,
		Description: query.Description,
		Url:         query.Url,
		Icon:        query.Icon,
		IsNav:       query.IsNav,
	}
	if query.Status != nil {
		contentQuery.Status = new(contentv1.DomainStatus(*query.Status))
	}
	reply, err := r.contentClient.Domain.List(ctx, &contentv1.ListDomains_Request{
		Page:  req.Page,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.Domain, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		row := &bbscontentv1.Domain{
			Id:          item.GetId(),
			Name:        item.GetName(),
			Description: item.Description,
			Status:      bbscontentv1.DomainStatus(item.GetStatus()),
			Url:         item.Url,
			Icon:        item.Icon,
			IsNav:       item.GetIsNav(),
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			CreatedAt:   formatProtoTime(item.GetCreatedAt()),
			UpdatedAt:   formatProtoTime(item.GetUpdatedAt()),
		}
		for _, tag := range item.GetTags() {
			row.Tags = append(row.Tags, &bbscontentv1.Tag{
				Id:          tag.GetId(),
				Name:        tag.GetName(),
				Description: tag.Description,
				DomainId:    tag.DomainId,
				Status:      new(bbscontentv1.TagStatus(tag.GetStatus())),
				CreatedAt:   formatProtoTime(tag.GetCreatedAt()),
				UpdatedAt:   formatProtoTime(tag.GetUpdatedAt()),
			})
		}
		rows = append(rows, row)
	}
	return &bbscontentv1.ListDomains_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *ContentRepo) ListTags(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error) {
	query := req.GetQuery()
	if query == nil {
		query = &bbscontentv1.TagQuery{}
	}
	contentQuery := &contentv1.TagQueryParams{
		Ids:         query.GetIds(),
		Name:        query.Name,
		Names:       query.GetNames(),
		Description: query.Description,
		DomainId:    query.DomainId,
	}
	if query.Status != nil {
		contentQuery.Status = new(contentv1.TagStatus(*query.Status))
	}
	reply, err := r.contentClient.Tag.List(ctx, &contentv1.ListTags_Request{
		Page:  req.Page,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.Tag, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, &bbscontentv1.Tag{
			Id:          item.GetId(),
			Name:        item.GetName(),
			Description: item.Description,
			DomainId:    item.DomainId,
			Status:      new(bbscontentv1.TagStatus(item.GetStatus())),
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			CreatedAt:   formatProtoTime(item.GetCreatedAt()),
			UpdatedAt:   formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &bbscontentv1.ListTags_Reply{Page: reply.GetPage(), Rows: rows}, nil
}
