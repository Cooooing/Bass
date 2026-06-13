package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/apperror"
	"common/pkg/client/rpc"
	"common/pkg/util"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	contentv1 "common/proto/gen/content/v1"
	userv1 "common/proto/gen/user/v1"
	"context"
	"fmt"

	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
	"github.com/samber/lo"
)

var _ repo.ContentArticleClient = (*ContentArticleClient)(nil)

type ContentArticleClient struct {
	contentClient *rpc.ContentClient
	userClient    *rpc.UserClient
}

func NewContentArticleClient(contentClient *rpc.ContentClient, userClient *rpc.UserClient) repo.ContentArticleClient {
	return &ContentArticleClient{contentClient: contentClient, userClient: userClient}
}

func (r *ContentArticleClient) CreateArticle(ctx context.Context, req *bbscontentv1.CreateArticle_Request) (*bbscontentv1.CreateArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	article := req.GetArticle()
	save := &contentv1.CreateArticle_Request_Article{}
	if article != nil {
		save.Title = article.GetTitle()
		save.Content = article.GetContent()
		save.RewardContent = article.RewardContent
		save.RewardPoints = article.RewardPoints
		save.Type = contentv1.ArticleType(article.GetType())
		save.Statement = article.Statement
		save.Commentable = article.Commentable
		save.Anonymous = article.Anonymous
		save.TagIds = article.GetTagIds()
		switch article.GetType() {
		case bbscontentv1.ArticleType_ARTICLE_TYPE_QA:
			if article.BountyPoints != nil {
				save.TypeParams = &contentv1.CreateArticle_Request_Article_Qa{
					Qa: &contentv1.CreateArticle_Request_Article_QA{BountyPoints: article.GetBountyPoints()},
				}
			}
		}
	}
	reply, err := r.contentClient.Article.Create(ctx, &contentv1.CreateArticle_Request{Article: save, UserId: userID})
	if err != nil {
		return nil, err
	}
	item := reply.GetArticle()
	lastComments, states, err := r.loadArticleFacts(ctx, []int64{item.GetId()}, userID)
	if err != nil {
		return nil, err
	}
	profiles, err := r.loadAccountProfiles(ctx, r.articleProfileIDs(item, lastComments[item.GetId()])...)
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.CreateArticle_Reply{Article: r.articleDetail(item, profiles, lastComments[item.GetId()], states[item.GetId()])}, nil
}

func (r *ContentArticleClient) UpdateArticle(ctx context.Context, req *bbscontentv1.UpdateArticle_Request) (*bbscontentv1.UpdateArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	article := req.GetArticle()
	save := &contentv1.UpdateArticle_Request_Article{}
	articleID := req.GetArticleId()
	if article != nil {
		save.Title = article.GetTitle()
		save.Content = article.GetContent()
		save.RewardContent = article.RewardContent
		save.RewardPoints = article.RewardPoints
		save.Type = contentv1.ArticleType(article.GetType())
		save.Statement = article.Statement
		save.Commentable = article.Commentable
		save.Anonymous = article.Anonymous
		save.TagIds = article.GetTagIds()
		switch article.GetType() {
		case bbscontentv1.ArticleType_ARTICLE_TYPE_QA:
			if article.BountyPoints != nil {
				save.TypeParams = &contentv1.UpdateArticle_Request_Article_Qa{
					Qa: &contentv1.UpdateArticle_Request_Article_QA{BountyPoints: article.GetBountyPoints()},
				}
			}
		}
	}
	reply, err := r.contentClient.Article.Update(ctx, &contentv1.UpdateArticle_Request{ArticleId: articleID, Article: save, UserId: userID})
	if err != nil {
		return nil, err
	}
	item := reply.GetArticle()
	lastComments, states, err := r.loadArticleFacts(ctx, []int64{item.GetId()}, userID)
	if err != nil {
		return nil, err
	}
	profiles, err := r.loadAccountProfiles(ctx, r.articleProfileIDs(item, lastComments[item.GetId()])...)
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.UpdateArticle_Reply{Article: r.articleDetail(item, profiles, lastComments[item.GetId()], states[item.GetId()])}, nil
}

func (r *ContentArticleClient) UpdateDraftArticle(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Request) (*bbscontentv1.UpdateDraftArticle_Reply, error) {
	article := req.GetArticle()
	update := &bbscontentv1.UpdateArticle_Request{ArticleId: req.GetArticleId()}
	if article != nil {
		update.Article = &bbscontentv1.UpdateArticle_Request_Article{
			Title:         article.GetTitle(),
			Content:       article.GetContent(),
			RewardContent: article.RewardContent,
			RewardPoints:  article.RewardPoints,
			Type:          bbscontentv1.ArticleType(article.GetType()),
			BountyPoints:  article.BountyPoints,
			Statement:     article.Statement,
			Commentable:   article.Commentable,
			Anonymous:     article.Anonymous,
			TagIds:        article.GetTagIds(),
		}
	}
	reply, err := r.UpdateArticle(ctx, update)
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.UpdateDraftArticle_Reply{Article: reply.GetArticle()}, nil
}

func (r *ContentArticleClient) PublishArticle(ctx context.Context, req *bbscontentv1.PublishArticle_Request) (*bbscontentv1.PublishArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.Publish(ctx, &contentv1.PublishArticle_Request{
		ArticleId:  req.GetArticleId(),
		UserId:     userID,
		Visibility: contentv1.ArticleVisibility(req.GetVisibility()),
	})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.PublishArticle_Reply{}, nil
}

func (r *ContentArticleClient) DiscardDraftArticle(ctx context.Context, req *bbscontentv1.DiscardDraftArticle_Request) (*bbscontentv1.DiscardDraftArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.DiscardDraft(ctx, &contentv1.DiscardDraftArticle_Request{ArticleId: req.GetArticleId(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.DiscardDraftArticle_Reply{}, nil
}

func (r *ContentArticleClient) ListArticles(ctx context.Context, req *bbscontentv1.ListArticles_Request) (*bbscontentv1.ListArticles_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	query := req.GetQuery()
	if query == nil {
		query = &bbscontentv1.ArticleQuery{}
	}
	contentQuery := &contentv1.ArticleQueryParams{
		TagId:    query.TagId,
		DomainId: query.DomainId,
		Keyword:  query.Keyword,
		AuthorId: query.AuthorId,
	}
	if query.Type != nil {
		contentQuery.Type = new(contentv1.ArticleType(*query.Type))
	}
	if query.Order != nil {
		contentQuery.Order = new(contentv1.ArticleOrder(*query.Order))
	}
	normal := contentv1.ContentRestriction_CONTENT_RESTRICTION_NONE
	locked := contentv1.ContentRestriction_CONTENT_RESTRICTION_LOCKED
	contentQuery.Restrictions = []contentv1.ContentRestriction{normal, locked}
	if query.AuthorId != nil && *query.AuthorId == userID {
		if query.PublishStatus != nil {
			contentQuery.PublishStatus = new(contentv1.ArticlePublishStatus(*query.PublishStatus))
		}
		if len(query.PublishStatuses) > 0 {
			contentQuery.PublishStatuses = lo.Map(query.PublishStatuses, func(item bbscontentv1.ArticlePublishStatus, _ int) contentv1.ArticlePublishStatus {
				return contentv1.ArticlePublishStatus(item)
			})
		}
		if query.Visibility != nil {
			contentQuery.Visibility = new(contentv1.ArticleVisibility(*query.Visibility))
		}
		if len(query.Visibilities) > 0 {
			contentQuery.Visibilities = lo.Map(query.Visibilities, func(item bbscontentv1.ArticleVisibility, _ int) contentv1.ArticleVisibility {
				return contentv1.ArticleVisibility(item)
			})
		}
	} else {
		published := contentv1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_PUBLISHED
		public := contentv1.ArticleVisibility_ARTICLE_VISIBILITY_PUBLIC
		contentQuery.PublishStatus = &published
		contentQuery.Visibility = &public
	}
	reply, err := r.contentClient.Article.Page(ctx, &contentv1.PageArticles_Request{
		Page:  req.GetPage(),
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	articleIDs := make([]int64, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		articleIDs = append(articleIDs, item.GetId())
	}
	lastComments, states, err := r.loadArticleFacts(ctx, articleIDs, userID)
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, 0, len(reply.GetRows())*2)
	for _, item := range reply.GetRows() {
		userIDs = append(userIDs, r.articleProfileIDs(item, lastComments[item.GetId()])...)
	}
	profiles, err := r.loadAccountProfiles(ctx, userIDs...)
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.ArticleListItem, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, r.articleListItem(item, profiles, lastComments[item.GetId()], states[item.GetId()]))
	}
	return &bbscontentv1.ListArticles_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *ContentArticleClient) GetArticle(ctx context.Context, req *bbscontentv1.GetArticle_Request) (*bbscontentv1.GetArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Article.Get(ctx, &contentv1.GetArticle_Request{ArticleId: req.GetArticleId()})
	if err != nil {
		return nil, err
	}
	item := reply.GetArticle()
	isAuthor := item.CreatedBy != nil && *item.CreatedBy == userID
	if item.GetRestriction() == contentv1.ContentRestriction_CONTENT_RESTRICTION_HIDDEN {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	switch item.GetPublishStatus() {
	case contentv1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_DRAFT:
		if !isAuthor {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
	case contentv1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_PUBLISHED, contentv1.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_ARCHIVED:
		if item.GetVisibility() == contentv1.ArticleVisibility_ARTICLE_VISIBILITY_PRIVATE && !isAuthor {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
	}
	lastComments, states, err := r.loadArticleFacts(ctx, []int64{item.GetId()}, userID)
	if err != nil {
		return nil, err
	}
	profiles, err := r.loadAccountProfiles(ctx, r.articleProfileIDs(item, lastComments[item.GetId()])...)
	if err != nil {
		return nil, err
	}
	detail := r.articleDetail(item, profiles, lastComments[item.GetId()], states[item.GetId()])
	if item.GetHasPostscript() {
		postscriptReply, err := r.contentClient.Article.ListPostscripts(ctx, &contentv1.ListArticlePostscripts_Request{
			ArticleId: item.GetId(),
		})
		if err != nil {
			return nil, err
		}
		detail.Postscripts = make([]*bbscontentv1.ArticlePostscript, 0, len(postscriptReply.GetRows()))
		for _, postscript := range postscriptReply.GetRows() {
			detail.Postscripts = append(detail.Postscripts, &bbscontentv1.ArticlePostscript{
				Id:            postscript.GetId(),
				ArticleId:     postscript.GetArticleId(),
				Content:       postscript.GetContent(),
				ContentRender: r.articlePostscriptContentRender(postscript.GetId(), postscript.GetContent()),
				Restriction:   bbscontentv1.ContentRestriction(postscript.GetRestriction()),
				CreatedBy:     postscript.CreatedBy,
				UpdatedBy:     postscript.UpdatedBy,
				CreatedAt:     formatProtoTime(postscript.GetCreatedAt()),
				UpdatedAt:     formatProtoTime(postscript.GetUpdatedAt()),
			})
		}
	}
	return &bbscontentv1.GetArticle_Reply{Article: detail}, nil
}

func (r *ContentArticleClient) ViewArticle(ctx context.Context, articleID int64) error {
	userID, err := currentUserID(ctx)
	if err != nil {
		return err
	}
	_, err = r.contentClient.Article.View(ctx, &contentv1.ViewArticle_Request{ArticleId: articleID, ViewerUserId: new(userID)})
	return err
}

func (r *ContentArticleClient) LikeArticle(ctx context.Context, req *bbscontentv1.LikeArticle_Request) (*bbscontentv1.LikeArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Article.Like(ctx, &contentv1.LikeArticle_Request{ArticleId: req.GetArticleId(), Liked: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.LikeArticle_Reply{Liked: reply.GetLiked()}, nil
}

func (r *ContentArticleClient) ThankArticle(ctx context.Context, req *bbscontentv1.ThankArticle_Request) (*bbscontentv1.ThankArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Article.Thank(ctx, &contentv1.ThankArticle_Request{ArticleId: req.GetArticleId(), Thanked: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.ThankArticle_Reply{Thanked: reply.GetThanked()}, nil
}

func (r *ContentArticleClient) CollectArticle(ctx context.Context, req *bbscontentv1.CollectArticle_Request) (*bbscontentv1.CollectArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Article.Collect(ctx, &contentv1.CollectArticle_Request{ArticleId: req.GetArticleId(), Collected: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.CollectArticle_Reply{Collected: reply.GetCollected()}, nil
}

func (r *ContentArticleClient) WatchArticle(ctx context.Context, req *bbscontentv1.WatchArticle_Request) (*bbscontentv1.WatchArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Article.Watch(ctx, &contentv1.WatchArticle_Request{ArticleId: req.GetArticleId(), Watched: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.WatchArticle_Reply{Watched: reply.GetWatched()}, nil
}

func (r *ContentArticleClient) RewardArticle(ctx context.Context, req *bbscontentv1.RewardArticle_Request) (*bbscontentv1.RewardArticle_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Article.Reward(ctx, &contentv1.RewardArticle_Request{ArticleId: req.GetArticleId(), Points: req.GetPoints(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.RewardArticle_Reply{}, nil
}

func (r *ContentArticleClient) AcceptAnswerArticle(ctx context.Context, req *bbscontentv1.AcceptAnswerArticle_Request) (*bbscontentv1.AcceptAnswerArticle_Reply, error) {
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

func (r *ContentArticleClient) loadArticleFacts(ctx context.Context, articleIDs []int64, userID int64) (map[int64]*contentv1.Comment, map[int64]*bbscontentv1.ArticleViewerActionState, error) {
	if len(articleIDs) == 0 {
		return map[int64]*contentv1.Comment{}, map[int64]*bbscontentv1.ArticleViewerActionState{}, nil
	}
	commentReply, err := r.contentClient.Comment.MapArticleLastComments(ctx, &contentv1.MapArticleLastComments_Request{ArticleIds: articleIDs})
	if err != nil {
		return nil, nil, err
	}
	stateReply, err := r.contentClient.Article.MapViewerActionStates(ctx, &contentv1.MapArticleViewerActionStates_Request{
		ArticleIds: articleIDs,
		UserId:     userID,
	})
	if err != nil {
		return nil, nil, err
	}
	return commentReply.GetComments(), r.articleViewerActionStates(stateReply.GetStates()), nil
}

func (r *ContentArticleClient) articleViewerActionStates(states map[int64]*contentv1.ArticleViewerActionState) map[int64]*bbscontentv1.ArticleViewerActionState {
	reply := make(map[int64]*bbscontentv1.ArticleViewerActionState, len(states))
	for articleID, state := range states {
		reply[articleID] = &bbscontentv1.ArticleViewerActionState{
			Liked:     state.GetLiked(),
			Thanked:   state.GetThanked(),
			Collected: state.GetCollected(),
			Watched:   state.GetWatched(),
		}
	}
	return reply
}

func (r *ContentArticleClient) articleProfileIDs(item *contentv1.Article, lastComment *contentv1.Comment) []int64 {
	if item == nil {
		return nil
	}
	userIDs := make([]int64, 0, 2)
	if !item.GetAnonymous() && item.CreatedBy != nil {
		userIDs = append(userIDs, *item.CreatedBy)
	}
	if lastComment != nil && lastComment.CreatedBy != nil {
		if item.GetAnonymous() && item.CreatedBy != nil && *lastComment.CreatedBy == *item.CreatedBy {
			return userIDs
		}
		userIDs = append(userIDs, *lastComment.CreatedBy)
	}
	return userIDs
}

func (r *ContentArticleClient) articleListItem(item *contentv1.Article, profiles map[int64]*bbsuserv1.AccountProfile, lastComment *contentv1.Comment, state *bbscontentv1.ArticleViewerActionState) *bbscontentv1.ArticleListItem {
	if item == nil {
		return nil
	}
	if state == nil {
		state = &bbscontentv1.ArticleViewerActionState{}
	}
	content := r.articleSummaryContent(item.GetContent())
	out := &bbscontentv1.ArticleListItem{
		Id:                item.GetId(),
		Title:             item.GetTitle(),
		Content:           content,
		ContentRender:     r.articleContentRender(item.GetId(), content),
		HasPostscript:     item.GetHasPostscript(),
		HasReward:         item.GetHasReward(),
		PublishStatus:     bbscontentv1.ArticlePublishStatus(item.GetPublishStatus()),
		Visibility:        bbscontentv1.ArticleVisibility(item.GetVisibility()),
		Restriction:       bbscontentv1.ContentRestriction(item.GetRestriction()),
		Type:              bbscontentv1.ArticleType(item.GetType()),
		Statement:         item.Statement,
		Commentable:       item.GetCommentable(),
		Anonymous:         item.GetAnonymous(),
		ViewCount:         item.GetViewCount(),
		ThankCount:        item.GetThankCount(),
		LikeCount:         item.GetLikeCount(),
		CollectCount:      item.GetCollectCount(),
		WatchCount:        item.GetWatchCount(),
		ReplyCount:        item.GetReplyCount(),
		CoverImageUrl:     r.articleCoverImageURL(item),
		ViewerActionState: state,
		CreatedAt:         formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:         formatProtoTime(item.GetUpdatedAt()),
		PublishedAt:       formatProtoTime(item.GetPublishedAt()),
		EditedAt:          formatProtoTime(item.GetEditedAt()),
	}
	switch item.GetType() {
	case contentv1.ArticleType_ARTICLE_TYPE_QA:
		if qa := item.GetQa(); qa != nil {
			out.BountyPoints = qa.BountyPoints
			out.AcceptedAnswerId = qa.AcceptedAnswerId
		}
	}
	if !item.GetAnonymous() {
		out.CreatedBy = item.CreatedBy
		out.UpdatedBy = item.UpdatedBy
	}
	if !item.GetAnonymous() && item.CreatedBy != nil {
		out.AuthorUser = profiles[*item.CreatedBy]
	}
	if lastComment != nil {
		out.LastReplyAt = formatProtoTime(lastComment.GetCreatedAt())
		if lastComment.CreatedBy != nil {
			if !item.GetAnonymous() || item.CreatedBy == nil || *lastComment.CreatedBy != *item.CreatedBy {
				out.LastReplyUser = profiles[*lastComment.CreatedBy]
			}
		}
	}
	return out
}

func (r *ContentArticleClient) articleDetail(item *contentv1.Article, profiles map[int64]*bbsuserv1.AccountProfile, lastComment *contentv1.Comment, state *bbscontentv1.ArticleViewerActionState) *bbscontentv1.ArticleDetail {
	if item == nil {
		return nil
	}
	if state == nil {
		state = &bbscontentv1.ArticleViewerActionState{}
	}
	out := &bbscontentv1.ArticleDetail{
		Id:                  item.GetId(),
		Title:               item.GetTitle(),
		Content:             item.GetContent(),
		ContentRender:       r.articleContentRender(item.GetId(), item.GetContent()),
		HasPostscript:       item.GetHasPostscript(),
		HasReward:           item.GetHasReward(),
		RewardContent:       item.RewardContent,
		RewardContentRender: r.articleRewardContentRender(item.GetId(), item.RewardContent),
		RewardPoints:        item.RewardPoints,
		PublishStatus:       bbscontentv1.ArticlePublishStatus(item.GetPublishStatus()),
		Visibility:          bbscontentv1.ArticleVisibility(item.GetVisibility()),
		Restriction:         bbscontentv1.ContentRestriction(item.GetRestriction()),
		Type:                bbscontentv1.ArticleType(item.GetType()),
		Statement:           item.Statement,
		Commentable:         item.GetCommentable(),
		Anonymous:           item.GetAnonymous(),
		ViewCount:           item.GetViewCount(),
		ThankCount:          item.GetThankCount(),
		LikeCount:           item.GetLikeCount(),
		CollectCount:        item.GetCollectCount(),
		WatchCount:          item.GetWatchCount(),
		ReplyCount:          item.GetReplyCount(),
		CoverImageUrl:       r.articleCoverImageURL(item),
		ViewerActionState:   state,
		CreatedAt:           formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:           formatProtoTime(item.GetUpdatedAt()),
		PublishedAt:         formatProtoTime(item.GetPublishedAt()),
		EditedAt:            formatProtoTime(item.GetEditedAt()),
	}
	switch item.GetType() {
	case contentv1.ArticleType_ARTICLE_TYPE_QA:
		if qa := item.GetQa(); qa != nil {
			out.BountyPoints = qa.BountyPoints
			out.AcceptedAnswerId = qa.AcceptedAnswerId
		}
	}
	if !item.GetAnonymous() {
		out.CreatedBy = item.CreatedBy
		out.UpdatedBy = item.UpdatedBy
	}
	if !item.GetAnonymous() && item.CreatedBy != nil {
		out.AuthorUser = profiles[*item.CreatedBy]
	}
	if lastComment != nil {
		out.LastReplyAt = formatProtoTime(lastComment.GetCreatedAt())
		if lastComment.CreatedBy != nil {
			if !item.GetAnonymous() || item.CreatedBy == nil || *lastComment.CreatedBy != *item.CreatedBy {
				out.LastReplyUser = profiles[*lastComment.CreatedBy]
			}
		}
	}
	return out
}

func (r *ContentArticleClient) articleSummaryContent(content string) string {
	runes := []rune(content)
	if len(runes) > 200 {
		return string(runes[:200]) + "..."
	}
	return content
}

func (r *ContentArticleClient) articleContentRender(articleID int64, content string) string {
	return util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_content", articleID), content)
}

func (r *ContentArticleClient) articleRewardContentRender(articleID int64, content *string) *string {
	if content == nil {
		return nil
	}
	return new(util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_reward_content", articleID), *content))
}

func (r *ContentArticleClient) articlePostscriptContentRender(postscriptID int64, content string) string {
	return util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_postscript", postscriptID), content)
}

func (r *ContentArticleClient) articleCoverImageURL(item *contentv1.Article) *string {
	if item == nil || item.GetContent() == "" {
		return nil
	}
	tree := parse.Parse(fmt.Sprintf("%s_%d", "article_content", item.GetId()), []byte(item.GetContent()), parse.NewOptions())
	coverImageURL := new("")
	ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
		return util.ParseNodeImageCoverImageUrl(n, entering, coverImageURL)
	})
	if *coverImageURL == "" {
		return nil
	}
	return coverImageURL
}

func (r *ContentArticleClient) loadAccountProfiles(ctx context.Context, userIDs ...int64) (map[int64]*bbsuserv1.AccountProfile, error) {
	if r.userClient == nil {
		return map[int64]*bbsuserv1.AccountProfile{}, nil
	}
	ids := lo.Filter(lo.Uniq(userIDs), func(userID int64, _ int) bool {
		return userID != 0
	})
	if len(ids) == 0 {
		return map[int64]*bbsuserv1.AccountProfile{}, nil
	}
	reply, err := r.userClient.Account.Map(ctx, &userv1.MapAccounts_Request{
		Query: &userv1.AccountQuery{UserIds: ids},
	})
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*bbsuserv1.AccountProfile, len(reply.GetAccounts()))
	for userID, account := range reply.GetAccounts() {
		basic := account.GetBasic()
		if basic == nil {
			continue
		}
		out[userID] = &bbsuserv1.AccountProfile{
			Id:            basic.GetId(),
			Name:          basic.GetName(),
			Nickname:      basic.Nickname,
			Url:           basic.Url,
			AvatarUrl:     basic.AvatarUrl,
			Introduction:  basic.Introduction,
			Mbti:          bbsuserv1.MBTI(basic.GetMbti()),
			Status:        bbsuserv1.AccountStatus(basic.GetStatus()),
			FollowCount:   basic.FollowCount,
			FollowerCount: basic.FollowerCount,
			CreatedAt:     formatProtoTime(basic.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(basic.GetUpdatedAt()),
		}
	}
	return out, nil
}
