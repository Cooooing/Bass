package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/pkg/util"
	"common/proto/gen/common"
	contentv1 "common/proto/gen/content/v1"
	contentv1enum "common/proto/gen/content/v1/enum"
	userv1 "common/proto/gen/user/v1"
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
	"github.com/samber/lo"
)

func cloneDataMessage[T proto.Message](
	src proto.Message,
	dst T,
) T {
	if src == nil {
		return dst
	}
	data, err := proto.Marshal(src)
	if err != nil {
		return dst
	}
	_ = proto.Unmarshal(data, dst)
	return dst
}

var _ repo.ContentArticleClient = (*ContentArticleClient)(nil)

type ContentArticleClient struct {
	contentClient *rpc.ContentClient
	userClient    *rpc.UserClient
}

func NewContentArticleClient(
	contentClient *rpc.ContentClient,
	userClient *rpc.UserClient,
) repo.ContentArticleClient {
	return &ContentArticleClient{
		contentClient: contentClient,
		userClient:    userClient,
	}
}

func (r *ContentArticleClient) CreateArticle(
	ctx context.Context,
	req *repo.CreateArticleReq,
) (*repo.ArticleDetail, error) {
	save := &contentv1.CreateArticle_Req_Article{}
	if req != nil && req.Article != nil {
		article := req.Article
		save.Title = article.Title
		save.Content = article.Content
		save.RewardContent = article.RewardContent
		save.RewardPoints = article.RewardPoints
		save.Type = contentv1enum.ArticleType(article.Type)
		save.Statement = article.Statement
		save.Commentable = article.Commentable
		save.Anonymous = article.Anonymous
		save.TagIds = article.TagIDs
		switch contentv1enum.ArticleType(article.Type) {
		case contentv1enum.ArticleType_ARTICLE_TYPE_QA:
			if article.BountyPoints != nil {
				save.TypeParams = &contentv1.CreateArticle_Req_Article_Qa{
					Qa: &contentv1.CreateArticle_Req_Article_QA{
						BountyPoints: *article.BountyPoints,
					},
				}
			}
		}
	}
	reply, err := r.contentClient.Article.Create(ctx, &contentv1.CreateArticle_Req{
		Article: save,
		UserId:  req.UserID,
	})
	if err != nil {
		return nil, err
	}
	item := cloneDataMessage(reply.GetArticle(), &contentv1.PageArticles_Resp_Article{})
	lastComments, states, err := r.loadArticleFacts(ctx, []int64{item.GetId()}, req.UserID)
	if err != nil {
		return nil, err
	}
	profiles, err := r.loadAccountProfiles(ctx, r.articleProfileIDs(item, lastComments[item.GetId()])...)
	if err != nil {
		return nil, err
	}
	return r.articleDetail(item, profiles, lastComments[item.GetId()], states[item.GetId()]), nil
}

func (r *ContentArticleClient) UpdateArticle(
	ctx context.Context,
	req *repo.UpdateArticleReq,
) (*repo.ArticleDetail, error) {
	save := &contentv1.UpdateArticle_Req_Article{}
	if req != nil && req.Article != nil {
		article := req.Article
		save.Title = article.Title
		save.Content = article.Content
		save.RewardContent = article.RewardContent
		save.RewardPoints = article.RewardPoints
		save.Type = contentv1enum.ArticleType(article.Type)
		save.Statement = article.Statement
		save.Commentable = article.Commentable
		save.Anonymous = article.Anonymous
		save.TagIds = article.TagIDs
		switch contentv1enum.ArticleType(article.Type) {
		case contentv1enum.ArticleType_ARTICLE_TYPE_QA:
			if article.BountyPoints != nil {
				save.TypeParams = &contentv1.UpdateArticle_Req_Article_Qa{
					Qa: &contentv1.UpdateArticle_Req_Article_QA{
						BountyPoints: *article.BountyPoints,
					},
				}
			}
		}
	}
	reply, err := r.contentClient.Article.Update(ctx, &contentv1.UpdateArticle_Req{
		ArticleId: req.ArticleID,
		Article:   save,
		UserId:    req.UserID,
	})
	if err != nil {
		return nil, err
	}
	item := cloneDataMessage(reply.GetArticle(), &contentv1.PageArticles_Resp_Article{})
	lastComments, states, err := r.loadArticleFacts(ctx, []int64{item.GetId()}, req.UserID)
	if err != nil {
		return nil, err
	}
	profiles, err := r.loadAccountProfiles(ctx, r.articleProfileIDs(item, lastComments[item.GetId()])...)
	if err != nil {
		return nil, err
	}
	return r.articleDetail(item, profiles, lastComments[item.GetId()], states[item.GetId()]), nil
}

func (r *ContentArticleClient) UpdateDraftArticle(
	ctx context.Context,
	req *repo.UpdateDraftArticleReq,
) (*repo.ArticleDetail, error) {
	reply, err := r.UpdateArticle(ctx, &repo.UpdateArticleReq{
		UserID:    req.UserID,
		ArticleID: req.ArticleID,
		Article:   req.Article,
	})
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (r *ContentArticleClient) PublishArticle(
	ctx context.Context,
	req *repo.PublishArticleReq,
) error {
	_, err := r.contentClient.Article.Publish(ctx, &contentv1.PublishArticle_Req{
		ArticleId:  req.ArticleID,
		UserId:     req.UserID,
		Visibility: contentv1enum.ArticleVisibility(req.Visibility),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) DiscardDraftArticle(
	ctx context.Context,
	req *repo.DiscardDraftArticleReq,
) error {
	_, err := r.contentClient.Article.DiscardDraft(ctx, &contentv1.DiscardDraftArticle_Req{
		ArticleId: req.ArticleID,
		UserId:    req.UserID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) ListArticles(
	ctx context.Context,
	req *repo.ListArticlesReq,
) (*repo.ListArticlesResp, error) {
	query := req.Query
	if query == nil {
		query = &repo.ArticleQuery{}
	}
	contentQuery := &contentv1.PageArticles_Req_ArticleQueryParams{
		TagId:    query.TagID,
		DomainId: query.DomainID,
		Keyword:  query.Keyword,
		AuthorId: query.AuthorID,
	}
	if query.Type != nil {
		contentQuery.Type = new(contentv1enum.ArticleType(*query.Type))
	}
	if query.Order != nil {
		contentQuery.Order = new(contentv1enum.ArticleOrder(*query.Order))
	}
	normal := contentv1enum.ContentRestriction_CONTENT_RESTRICTION_NONE
	locked := contentv1enum.ContentRestriction_CONTENT_RESTRICTION_LOCKED
	contentQuery.Restrictions = []contentv1enum.ContentRestriction{normal, locked}
	if query.AuthorID != nil && *query.AuthorID == req.UserID {
		if query.PublishStatus != nil {
			contentQuery.PublishStatus = new(contentv1enum.ArticlePublishStatus(*query.PublishStatus))
		}
		if len(query.PublishStatuses) > 0 {
			contentQuery.PublishStatuses = lo.Map(query.PublishStatuses, func(item int32, _ int) contentv1enum.ArticlePublishStatus {
				return contentv1enum.ArticlePublishStatus(item)
			})
		}
		if query.Visibility != nil {
			contentQuery.Visibility = new(contentv1enum.ArticleVisibility(*query.Visibility))
		}
		if len(query.Visibilities) > 0 {
			contentQuery.Visibilities = lo.Map(query.Visibilities, func(item int32, _ int) contentv1enum.ArticleVisibility {
				return contentv1enum.ArticleVisibility(item)
			})
		}
	} else {
		published := contentv1enum.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_PUBLISHED
		public := contentv1enum.ArticleVisibility_ARTICLE_VISIBILITY_PUBLIC
		contentQuery.PublishStatus = &published
		contentQuery.Visibility = &public
	}
	var pageReq *common.PageReq
	if req.Page != nil {
		pageReq = &common.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	reply, err := r.contentClient.Article.Page(ctx, &contentv1.PageArticles_Req{
		Page:  pageReq,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	articleIDs := make([]int64, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		articleIDs = append(articleIDs, item.GetId())
	}
	lastComments, states, err := r.loadArticleFacts(ctx, articleIDs, req.UserID)
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
	rows := make([]*repo.ArticleListItem, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, r.articleListItem(item, profiles, lastComments[item.GetId()], states[item.GetId()]))
	}
	var page *repo.PageResp
	if reply.GetPage() != nil {
		page = &repo.PageResp{
			Page:  reply.GetPage().GetPage(),
			Size:  reply.GetPage().GetSize(),
			Total: reply.GetPage().GetTotal(),
		}
	}
	return &repo.ListArticlesResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *ContentArticleClient) GetArticle(
	ctx context.Context,
	req *repo.GetArticleReq,
) (*repo.ArticleDetail, error) {
	reply, err := r.contentClient.Article.Get(ctx, &contentv1.GetArticle_Req{
		ArticleId: req.ArticleID,
	})
	if err != nil {
		return nil, err
	}
	item := cloneDataMessage(reply.GetArticle(), &contentv1.PageArticles_Resp_Article{})

	lastComments, states, err := r.loadArticleFacts(ctx, []int64{item.GetId()}, req.UserID)
	if err != nil {
		return nil, err
	}
	profiles, err := r.loadAccountProfiles(ctx, r.articleProfileIDs(item, lastComments[item.GetId()])...)
	if err != nil {
		return nil, err
	}
	detail := r.articleDetail(item, profiles, lastComments[item.GetId()], states[item.GetId()])
	if item.GetHasPostscript() {
		postscriptResp, err := r.contentClient.Article.ListPostscripts(ctx, &contentv1.ListArticlePostscripts_Req{
			ArticleId: item.GetId(),
		})
		if err != nil {
			return nil, err
		}
		detail.Postscripts = make([]*repo.ArticlePostscript, 0, len(postscriptResp.GetRows()))
		for _, postscript := range postscriptResp.GetRows() {
			detail.Postscripts = append(detail.Postscripts, &repo.ArticlePostscript{
				ID:            postscript.GetId(),
				ArticleID:     postscript.GetArticleId(),
				Content:       postscript.GetContent(),
				ContentRender: r.articlePostscriptContentRender(postscript.GetId(), postscript.GetContent()),
				Restriction:   int32(postscript.GetRestriction()),
				CreatedBy:     postscript.CreatedBy,
				UpdatedBy:     postscript.UpdatedBy,
				CreatedAt:     formatProtoTime(postscript.GetCreatedAt()),
				UpdatedAt:     formatProtoTime(postscript.GetUpdatedAt()),
			})
		}
	}
	return detail, nil
}

func (r *ContentArticleClient) ViewArticle(
	ctx context.Context,
	req *repo.ViewArticleReq,
) error {
	_, err := r.contentClient.Article.View(ctx, &contentv1.ViewArticle_Req{
		ArticleId:    req.ArticleID,
		ViewerUserId: new(req.UserID),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) LikeArticle(
	ctx context.Context,
	req *repo.LikeArticleReq,
) (bool, error) {
	reply, err := r.contentClient.Article.Like(ctx, &contentv1.LikeArticle_Req{
		ArticleId: req.ArticleID,
		Liked:     req.Active,
		UserId:    req.UserID,
	})
	if err != nil {
		return false, err
	}
	return reply.GetLiked(), nil
}

func (r *ContentArticleClient) ThankArticle(
	ctx context.Context,
	req *repo.ThankArticleReq,
) (bool, error) {
	reply, err := r.contentClient.Article.Thank(ctx, &contentv1.ThankArticle_Req{
		ArticleId: req.ArticleID,
		Thanked:   req.Active,
		UserId:    req.UserID,
	})
	if err != nil {
		return false, err
	}
	return reply.GetThanked(), nil
}

func (r *ContentArticleClient) CollectArticle(
	ctx context.Context,
	req *repo.CollectArticleReq,
) (bool, error) {
	reply, err := r.contentClient.Article.Collect(ctx, &contentv1.CollectArticle_Req{
		ArticleId: req.ArticleID,
		Collected: req.Active,
		UserId:    req.UserID,
	})
	if err != nil {
		return false, err
	}
	return reply.GetCollected(), nil
}

func (r *ContentArticleClient) WatchArticle(
	ctx context.Context,
	req *repo.WatchArticleReq,
) (bool, error) {
	reply, err := r.contentClient.Article.Watch(ctx, &contentv1.WatchArticle_Req{
		ArticleId: req.ArticleID,
		Watched:   req.Active,
		UserId:    req.UserID,
	})
	if err != nil {
		return false, err
	}
	return reply.GetWatched(), nil
}

func (r *ContentArticleClient) RewardArticle(
	ctx context.Context,
	req *repo.RewardArticleReq,
) error {
	_, err := r.contentClient.Article.Reward(ctx, &contentv1.RewardArticle_Req{
		ArticleId: req.ArticleID,
		Points:    req.Points,
		UserId:    req.UserID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) AcceptAnswerArticle(
	ctx context.Context,
	req *repo.AcceptAnswerArticleReq,
) error {
	_, err := r.contentClient.Article.AcceptAnswer(ctx, &contentv1.AcceptAnswerArticle_Req{
		ArticleId: req.ArticleID,
		CommentId: req.CommentID,
		UserId:    req.UserID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) loadArticleFacts(
	ctx context.Context,
	articleIDs []int64,
	userID int64,
) (map[int64]*contentv1.MapArticleLastComments_Resp_Comment, map[int64]*repo.ArticleViewerActionState, error) {
	if len(articleIDs) == 0 {
		return map[int64]*contentv1.MapArticleLastComments_Resp_Comment{}, map[int64]*repo.ArticleViewerActionState{}, nil
	}
	commentResp, err := r.contentClient.Comment.MapArticleLastComments(ctx, &contentv1.MapArticleLastComments_Req{
		ArticleIds: articleIDs,
	})
	if err != nil {
		return nil, nil, err
	}
	stateResp, err := r.contentClient.Article.MapViewerActionStates(ctx, &contentv1.MapArticleViewerActionStates_Req{
		ArticleIds: articleIDs,
		UserId:     userID,
	})
	if err != nil {
		return nil, nil, err
	}
	return commentResp.GetComments(), r.articleViewerActionStates(stateResp.GetStates()), nil
}

func (r *ContentArticleClient) articleViewerActionStates(
	states map[int64]*contentv1.MapArticleViewerActionStates_Resp_ArticleViewerActionState,
) map[int64]*repo.ArticleViewerActionState {
	reply := make(map[int64]*repo.ArticleViewerActionState, len(states))
	for articleID, state := range states {
		reply[articleID] = &repo.ArticleViewerActionState{
			Liked:     state.GetLiked(),
			Thanked:   state.GetThanked(),
			Collected: state.GetCollected(),
			Watched:   state.GetWatched(),
		}
	}
	return reply
}

func (r *ContentArticleClient) articleProfileIDs(
	item *contentv1.PageArticles_Resp_Article,
	lastComment *contentv1.MapArticleLastComments_Resp_Comment,
) []int64 {
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

func (r *ContentArticleClient) articleListItem(
	item *contentv1.PageArticles_Resp_Article,
	profiles map[int64]*repo.AccountProfile,
	lastComment *contentv1.MapArticleLastComments_Resp_Comment,
	state *repo.ArticleViewerActionState,
) *repo.ArticleListItem {
	if item == nil {
		return nil
	}
	if state == nil {
		state = &repo.ArticleViewerActionState{}
	}
	content := r.articleSummaryContent(item.GetContent())
	out := &repo.ArticleListItem{
		ID:                item.GetId(),
		Title:             item.GetTitle(),
		Content:           content,
		ContentRender:     r.articleContentRender(item.GetId(), content),
		HasPostscript:     item.GetHasPostscript(),
		HasReward:         item.GetHasReward(),
		PublishStatus:     int32(item.GetPublishStatus()),
		Visibility:        int32(item.GetVisibility()),
		Restriction:       int32(item.GetRestriction()),
		Type:              int32(item.GetType()),
		Statement:         item.Statement,
		Commentable:       item.GetCommentable(),
		Anonymous:         item.GetAnonymous(),
		ViewCount:         item.GetViewCount(),
		ThankCount:        item.GetThankCount(),
		LikeCount:         item.GetLikeCount(),
		CollectCount:      item.GetCollectCount(),
		WatchCount:        item.GetWatchCount(),
		ReplyCount:        item.GetReplyCount(),
		CoverImageURL:     r.articleCoverImageURL(item),
		ViewerActionState: state,
		CreatedAt:         formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:         formatProtoTime(item.GetUpdatedAt()),
		PublishedAt:       formatProtoTime(item.GetPublishedAt()),
		EditedAt:          formatProtoTime(item.GetEditedAt()),
	}
	switch item.GetType() {
	case contentv1enum.ArticleType_ARTICLE_TYPE_QA:
		if qa := item.GetQa(); qa != nil {
			out.BountyPoints = qa.BountyPoints
			out.AcceptedAnswerID = qa.AcceptedAnswerId
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

func (r *ContentArticleClient) articleDetail(
	item *contentv1.PageArticles_Resp_Article,
	profiles map[int64]*repo.AccountProfile,
	lastComment *contentv1.MapArticleLastComments_Resp_Comment,
	state *repo.ArticleViewerActionState,
) *repo.ArticleDetail {
	if item == nil {
		return nil
	}
	if state == nil {
		state = &repo.ArticleViewerActionState{}
	}
	out := &repo.ArticleDetail{
		ID:                  item.GetId(),
		Title:               item.GetTitle(),
		Content:             item.GetContent(),
		ContentRender:       r.articleContentRender(item.GetId(), item.GetContent()),
		HasPostscript:       item.GetHasPostscript(),
		HasReward:           item.GetHasReward(),
		RewardContent:       item.RewardContent,
		RewardContentRender: r.articleRewardContentRender(item.GetId(), item.RewardContent),
		RewardPoints:        item.RewardPoints,
		PublishStatus:       int32(item.GetPublishStatus()),
		Visibility:          int32(item.GetVisibility()),
		Restriction:         int32(item.GetRestriction()),
		Type:                int32(item.GetType()),
		Statement:           item.Statement,
		Commentable:         item.GetCommentable(),
		Anonymous:           item.GetAnonymous(),
		ViewCount:           item.GetViewCount(),
		ThankCount:          item.GetThankCount(),
		LikeCount:           item.GetLikeCount(),
		CollectCount:        item.GetCollectCount(),
		WatchCount:          item.GetWatchCount(),
		ReplyCount:          item.GetReplyCount(),
		CoverImageURL:       r.articleCoverImageURL(item),
		ViewerActionState:   state,
		CreatedAt:           formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:           formatProtoTime(item.GetUpdatedAt()),
		PublishedAt:         formatProtoTime(item.GetPublishedAt()),
		EditedAt:            formatProtoTime(item.GetEditedAt()),
	}
	switch item.GetType() {
	case contentv1enum.ArticleType_ARTICLE_TYPE_QA:
		if qa := item.GetQa(); qa != nil {
			out.BountyPoints = qa.BountyPoints
			out.AcceptedAnswerID = qa.AcceptedAnswerId
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

func (r *ContentArticleClient) articleSummaryContent(
	content string,
) string {
	runes := []rune(content)
	if len(runes) > 200 {
		return string(runes[:200]) + "..."
	}
	return content
}

func (r *ContentArticleClient) articleContentRender(
	articleID int64,
	content string,
) string {
	return util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_content", articleID), content)
}

func (r *ContentArticleClient) articleRewardContentRender(
	articleID int64,
	content *string,
) *string {
	if content == nil {
		return nil
	}
	return new(util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_reward_content", articleID), *content))
}

func (r *ContentArticleClient) articlePostscriptContentRender(
	postscriptID int64,
	content string,
) string {
	return util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_postscript", postscriptID), content)
}

func (r *ContentArticleClient) articleCoverImageURL(
	item *contentv1.PageArticles_Resp_Article,
) *string {
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

func (r *ContentArticleClient) loadAccountProfiles(
	ctx context.Context,
	userIDs ...int64,
) (map[int64]*repo.AccountProfile, error) {
	if r.userClient == nil {
		return map[int64]*repo.AccountProfile{}, nil
	}
	ids := lo.Filter(lo.Uniq(userIDs), func(userID int64, _ int) bool {
		return userID != 0
	})
	if len(ids) == 0 {
		return map[int64]*repo.AccountProfile{}, nil
	}
	reply, err := r.userClient.Account.Map(ctx, &userv1.MapAccounts_Req{
		Query: &userv1.MapAccounts_Req_AccountQuery{
			UserIds: ids,
		},
	})
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*repo.AccountProfile, len(reply.GetAccounts()))
	for userID, account := range reply.GetAccounts() {
		basic := account.GetBasic()
		if basic == nil {
			continue
		}
		out[userID] = &repo.AccountProfile{
			ID:            basic.GetId(),
			Name:          basic.GetName(),
			Nickname:      basic.Nickname,
			URL:           basic.Url,
			AvatarURL:     basic.AvatarUrl,
			Introduction:  basic.Introduction,
			MBTI:          int32(basic.GetMbti()),
			Status:        int32(basic.GetStatus()),
			FollowCount:   basic.FollowCount,
			FollowerCount: basic.FollowerCount,
			CreatedAt:     formatProtoTime(basic.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(basic.GetUpdatedAt()),
		}
	}
	return out, nil
}
