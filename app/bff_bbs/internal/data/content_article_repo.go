package data

import (
	"common/pkg/client/rpc"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	contentv1 "common/proto/gen/content/v1"
	contentv1enum "common/proto/gen/content/v1/enum"
	"context"
	"fmt"
	"sort"

	"bff_bbs/internal/biz/model"

	"bff_bbs/internal/biz/repo"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
	"github.com/samber/lo"
)

var _ repo.ContentArticleClient = (*ContentArticleClient)(nil)

type ContentArticleClient struct {
	accountProfileLoader
	contentClient *rpc.ContentClient
}

func NewContentArticleClient(
	contentClient *rpc.ContentClient,
	userClient *rpc.UserClient,
) repo.ContentArticleClient {
	return &ContentArticleClient{
		accountProfileLoader: accountProfileLoader{userClient: userClient},
		contentClient:        contentClient,
	}
}

func (r *ContentArticleClient) CreateDraftArticle(ctx context.Context, req *repo.CreateDraftArticleReq) (int64, error) {
	save := &contentv1.CreateDraftArticle_Req_Article{}
	if req != nil && req.Article != nil {
		article := req.Article
		save.Title = article.Title
		save.Content = article.Content
		save.RewardContent = article.RewardContent
		save.RewardPoints = article.RewardPoints
		save.Type = contentv1enum.ArticleType(article.Type)
		save.Statement = article.Statement
		save.Commentable = article.Commentable
		save.Poll = r.articlePollDraft(article.Poll)
		save.Lottery = r.articleLotteryDraft(article.Lottery)
	}
	reply, err := r.contentClient.Article.CreateDraft(ctx, &contentv1.CreateDraftArticle_Req{
		Article: save,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return 0, err
	}
	return reply.GetArticleId(), nil
}

func (r *ContentArticleClient) UpdateDraftArticle(ctx context.Context, req *repo.UpdateDraftArticleReq) (int64, error) {
	save := &contentv1.UpdateDraftArticle_Req_Article{}
	if req != nil && req.Article != nil {
		article := req.Article
		save.Title = article.Title
		save.Content = article.Content
		save.RewardContent = article.RewardContent
		save.RewardPoints = article.RewardPoints
		save.Type = contentv1enum.ArticleType(article.Type)
		save.Statement = article.Statement
		save.Commentable = article.Commentable
		save.Poll = r.articlePollDraft(article.Poll)
		save.Lottery = r.articleLotteryDraft(article.Lottery)
	}
	reply, err := r.contentClient.Article.UpdateDraft(ctx, &contentv1.UpdateDraftArticle_Req{
		ArticleId: req.ArticleID,
		Article:   save,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return 0, err
	}
	return reply.GetArticleId(), nil
}

func (r *ContentArticleClient) PublishArticle(ctx context.Context, req *repo.PublishArticleReq) error {
	publishReq := &contentv1.PublishArticle_Req{
		ArticleId: req.ArticleID,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR,
			ActorUserId: new(req.UserID),
		},
	}
	if req.ScheduledAt != nil {
		publishReq.ScheduledAt = timestamppb.New(*req.ScheduledAt)
	}
	publishReq.City = req.City
	_, err := r.contentClient.Article.Publish(ctx, publishReq)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) CancelPublishArticle(ctx context.Context, req *repo.CancelPublishArticleReq) error {
	_, err := r.contentClient.Article.CancelPublish(ctx, &contentv1.CancelPublishArticle_Req{
		ArticleId: req.ArticleID,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) ArchiveArticle(ctx context.Context, req *repo.ArchiveArticleReq) error {
	_, err := r.contentClient.Article.Archive(ctx, &contentv1.ArchiveArticle_Req{
		ArticleId: req.ArticleID,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR,
			ActorUserId: new(req.UserID),
		},
		Reason: req.Reason,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) SetPinnedArticleComment(ctx context.Context, req *repo.SetPinnedArticleCommentReq) error {
	_, err := r.contentClient.Article.SetPinnedComment(ctx, &contentv1.SetPinnedArticleComment_Req{ArticleId: req.ArticleID, CommentId: req.CommentID, Access: &contentv1.ContentAccess{Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR, ActorUserId: new(req.UserID)}})
	return err
}

func (r *ContentArticleClient) AcceptArticleAnswer(ctx context.Context, req *repo.AcceptArticleAnswerReq) error {
	_, err := r.contentClient.Article.AcceptAnswer(ctx, &contentv1.AcceptArticleAnswer_Req{ArticleId: req.ArticleID, CommentId: req.CommentID, Access: &contentv1.ContentAccess{Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR, ActorUserId: new(req.UserID)}})
	return err
}

func (r *ContentArticleClient) VoteArticlePoll(ctx context.Context, req *repo.VoteArticlePollReq) error {
	_, err := r.contentClient.Article.VotePoll(ctx, &contentv1.VoteArticlePoll_Req{ArticleId: req.ArticleID, OptionIds: req.OptionIDs, Access: &contentv1.ContentAccess{Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR, ActorUserId: new(req.UserID)}})
	return err
}

func (r *ContentArticleClient) EnterArticleLottery(ctx context.Context, req *repo.EnterArticleLotteryReq) error {
	_, err := r.contentClient.Article.EnterLottery(ctx, &contentv1.EnterArticleLottery_Req{ArticleId: req.ArticleID, Entered: req.Entered, Access: &contentv1.ContentAccess{Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR, ActorUserId: new(req.UserID)}})
	return err
}

func (*ContentArticleClient) articlePollDraft(row *model.ArticlePollDraft) *contentv1.ArticlePollDraft {
	if row == nil {
		return nil
	}
	return &contentv1.ArticlePollDraft{Options: row.Options, EndsAt: timestamppb.New(row.EndsAt), MaxChoices: row.MaxChoices}
}

func (*ContentArticleClient) articleLotteryDraft(row *model.ArticleLotteryDraft) *contentv1.ArticleLotteryDraft {
	if row == nil {
		return nil
	}
	prizes := make([]*contentv1.ArticleLotteryPrizeDraft, 0, len(row.Prizes))
	for _, prize := range row.Prizes {
		if prize != nil {
			prizes = append(prizes, &contentv1.ArticleLotteryPrizeDraft{Name: prize.Name, Description: prize.Description, WinnerLimit: prize.WinnerLimit})
		}
	}
	return &contentv1.ArticleLotteryDraft{Prizes: prizes, DrawAt: timestamppb.New(row.DrawAt)}
}
func (r *ContentArticleClient) DiscardDraftArticle(ctx context.Context, req *repo.DiscardDraftArticleReq) error {
	_, err := r.contentClient.Article.DiscardDraft(ctx, &contentv1.DiscardDraftArticle_Req{
		ArticleId: req.ArticleID,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) ListArticles(ctx context.Context, req *repo.ListArticlesReq) (*repo.ListArticlesResp, error) {
	query := req.Query
	if query == nil {
		query = &repo.ArticleQuery{}
	}
	contentQuery := &contentv1.ArticleQueryParams{
		TagId:             query.TagID,
		DomainId:          query.DomainID,
		AuthorId:          query.AuthorID,
		AuthorIds:         query.AuthorIDs,
		PublishedCity:     query.PublishedCity,
		Scheduled:         query.Scheduled,
		ExcludedAuthorIds: query.ExcludedAuthorIDs,
	}
	if query.Type != nil {
		contentQuery.Type = new(contentv1enum.ArticleType(*query.Type))
	}
	if query.Order != nil {
		contentQuery.Order = new(contentv1enum.ArticleOrder(*query.Order))
	}
	accessScope := contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST
	if req.UserID > 0 {
		accessScope = contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER
	}
	if query.AuthorID != nil && *query.AuthorID == req.UserID {
		accessScope = contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_AUTHOR
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
	}
	contentAccess := &contentv1.ContentAccess{Scope: accessScope}
	if req.UserID > 0 {
		contentAccess.ActorUserId = new(req.UserID)
	}
	var pageReq *commonmodel.PageReq
	if req.Page != nil {
		pageReq = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	reply, err := r.contentClient.Article.Page(ctx, &contentv1.PageArticles_Req{
		Page:   pageReq.ToProto(),
		Query:  contentQuery,
		Access: contentAccess,
	})
	if err != nil {
		return nil, err
	}
	rows, err := r.articleListItems(ctx, reply.GetRows(), req.UserID, query.ExcludedAuthorIDs)
	if err != nil {
		return nil, err
	}
	var page *commonmodel.PageResp
	if reply.GetPage() != nil {
		page = commonmodel.PageRespFromProto(reply.GetPage())
	}
	return &repo.ListArticlesResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *ContentArticleClient) articleListItems(
	ctx context.Context,
	articles []*contentv1.Article,
	userID int64,
	excludedUserIDs []int64,
) ([]*model.ArticleListItem, error) {
	articleIDs := make([]int64, 0, len(articles))
	for _, item := range articles {
		if item != nil {
			articleIDs = append(articleIDs, item.GetId())
		}
	}
	lastComments, states, tags, domains, err := r.loadArticleFacts(ctx, articleIDs, userID, excludedUserIDs)
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, 0, len(articles)*2)
	for _, item := range articles {
		if item != nil {
			userIDs = append(userIDs, r.articleProfileIDs(item, lastComments[item.GetId()])...)
		}
	}
	profiles, err := r.loadAccountProfiles(ctx, userIDs...)
	if err != nil {
		return nil, err
	}
	rows := make([]*model.ArticleListItem, 0, len(articles))
	for _, item := range articles {
		if item != nil {
			rows = append(rows, r.articleListItem(item, profiles, lastComments[item.GetId()], states[item.GetId()], tags[item.GetId()], domains[item.GetId()]))
		}
	}
	return rows, nil
}

func (r *ContentArticleClient) SearchArticles(ctx context.Context, req *repo.SearchArticlesReq) (*repo.SearchArticlesResp, error) {
	reply, err := r.contentClient.ArticleSearch.SearchPublic(ctx, &contentv1.SearchPublicArticles_Req{
		Query:             req.Query,
		Cursor:            req.Cursor,
		Size:              req.Size,
		ExcludedAuthorIds: req.ExcludedAuthorIDs,
	})
	if err != nil {
		return nil, err
	}
	rows, err := r.articleListItems(ctx, reply.GetRows(), req.UserID, req.ExcludedAuthorIDs)
	if err != nil {
		return nil, err
	}
	return &repo.SearchArticlesResp{Rows: rows, NextCursor: reply.NextCursor}, nil
}

func (r *ContentArticleClient) PageViewHistory(
	ctx context.Context,
	userID int64,
	page *commonmodel.PageReq,
	excludedAuthorIDs []int64,
) (*repo.ArticleViewHistoryPageResp, error) {
	if userID <= 0 {
		return nil, fmt.Errorf("view history requires a user")
	}
	pageReq := &commonmodel.PageReq{}
	if page != nil {
		pageReq.Page = page.Page
		pageReq.Size = page.Size
	}
	reply, err := r.contentClient.Article.PageViewHistory(ctx, &contentv1.PageArticleViewHistory_Req{
		Page: pageReq.ToProto(),
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER,
			ActorUserId: new(userID),
		},
		ExcludedAuthorIds: excludedAuthorIDs,
	})
	if err != nil {
		return nil, err
	}
	articleIDs := make([]int64, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		if row.GetArticle() != nil {
			articleIDs = append(articleIDs, row.GetArticle().GetId())
		}
	}
	lastComments, states, tags, domains, err := r.loadArticleFacts(ctx, articleIDs, userID, excludedAuthorIDs)
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, 0, len(reply.GetRows())*2)
	for _, row := range reply.GetRows() {
		if row.GetArticle() == nil {
			continue
		}
		userIDs = append(userIDs, r.articleProfileIDs(row.GetArticle(), lastComments[row.GetArticle().GetId()])...)
	}
	profiles, err := r.loadAccountProfiles(ctx, userIDs...)
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.ArticleViewHistoryRow, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		article := row.GetArticle()
		if article == nil {
			continue
		}
		historyRow := &repo.ArticleViewHistoryRow{
			Article: r.articleListItem(
				article,
				profiles,
				lastComments[article.GetId()],
				states[article.GetId()],
				tags[article.GetId()],
				domains[article.GetId()],
			),
		}
		if row.GetViewedAt() != nil {
			viewedAt := row.GetViewedAt().AsTime()
			historyRow.ViewedAt = &viewedAt
		}
		rows = append(rows, historyRow)
	}
	var pageResp *commonmodel.PageResp
	if reply.GetPage() != nil {
		pageResp = commonmodel.PageRespFromProto(reply.GetPage())
	}
	return &repo.ArticleViewHistoryPageResp{Rows: rows, Page: pageResp}, nil
}

func (r *ContentArticleClient) GetArticle(ctx context.Context, req *repo.GetArticleReq) (*model.ArticleDetail, error) {
	accessScope := contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST
	if req.UserID > 0 {
		accessScope = contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER
	}
	contentAccess := &contentv1.ContentAccess{Scope: accessScope}
	if req.UserID > 0 {
		contentAccess.ActorUserId = new(req.UserID)
	}
	reply, err := r.contentClient.Article.Get(ctx, &contentv1.GetArticle_Req{
		ArticleId: req.ArticleID,
		Access:    contentAccess,
	})
	if err != nil {
		return nil, err
	}
	item := reply.GetArticle()

	lastComments, states, tags, domains, err := r.loadArticleFacts(ctx, []int64{item.GetId()}, req.UserID, nil)
	if err != nil {
		return nil, err
	}
	profiles, err := r.loadAccountProfiles(ctx, r.articleProfileIDs(item, lastComments[item.GetId()])...)
	if err != nil {
		return nil, err
	}
	detail := r.articleDetail(item, profiles, lastComments[item.GetId()], states[item.GetId()], tags[item.GetId()], domains[item.GetId()])
	if detail.Lottery != nil {
		winnerIDs := make([]int64, 0)
		for _, prize := range detail.Lottery.Prizes {
			if prize != nil {
				winnerIDs = append(winnerIDs, prize.WinnerUserIDs...)
			}
		}
		winnerProfiles, loadErr := r.loadAccountProfiles(ctx, winnerIDs...)
		if loadErr != nil {
			return nil, loadErr
		}
		for _, prize := range detail.Lottery.Prizes {
			if prize == nil {
				continue
			}
			for _, userID := range prize.WinnerUserIDs {
				if profile := winnerProfiles[userID]; profile != nil {
					prize.WinnerUsers = append(prize.WinnerUsers, profile)
				}
			}
		}
	}
	if item.GetHasPostscript() {
		postscriptResp, err := r.contentClient.Postscript.List(ctx, &contentv1.ListPostscripts_Req{
			ArticleId: item.GetId(),
			Access:    contentAccess,
		})
		if err != nil {
			return nil, err
		}
		detail.Postscripts = make([]*model.ArticlePostscript, 0, len(postscriptResp.GetRows()))
		for _, postscript := range postscriptResp.GetRows() {
			detail.Postscripts = append(detail.Postscripts, &model.ArticlePostscript{
				ID:            postscript.GetId(),
				ArticleID:     postscript.GetArticleId(),
				Content:       postscript.GetContent(),
				ContentRender: util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_postscript", postscript.GetId()), postscript.GetContent()),
				Restriction:   int32(postscript.GetRestriction()),
				CreatedBy:     postscript.CreatedBy,
				UpdatedBy:     postscript.UpdatedBy,
				CreatedAt:     new(postscript.GetCreatedAt().AsTime()),
				UpdatedAt:     new(postscript.GetUpdatedAt().AsTime()),
			})
		}
	}
	return detail, nil
}

func (r *ContentArticleClient) ViewArticle(ctx context.Context, req *repo.ViewArticleReq) error {
	viewReq := &contentv1.ViewArticle_Req{
		ArticleId: req.ArticleID,
		Access: &contentv1.ContentAccess{
			Scope: contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_GUEST,
		},
	}
	if req.UserID > 0 {
		viewReq.Access = &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER,
			ActorUserId: new(req.UserID),
		}
	}
	viewReq.Ip = req.IP
	viewReq.UserAgent = req.UserAgent
	_, err := r.contentClient.Article.View(ctx, viewReq)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) LikeArticle(ctx context.Context, req *repo.LikeArticleReq) (bool, error) {
	reply, err := r.contentClient.Article.Like(ctx, &contentv1.LikeArticle_Req{
		ArticleId: req.ArticleID,
		Liked:     req.Active,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return false, err
	}
	return reply.GetLiked(), nil
}

func (r *ContentArticleClient) ThankArticle(ctx context.Context, req *repo.ThankArticleReq) (bool, error) {
	reply, err := r.contentClient.Article.Thank(ctx, &contentv1.ThankArticle_Req{
		ArticleId: req.ArticleID,
		Thanked:   req.Active,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return false, err
	}
	return reply.GetThanked(), nil
}

func (r *ContentArticleClient) CollectArticle(ctx context.Context, req *repo.CollectArticleReq) (bool, error) {
	reply, err := r.contentClient.Article.Collect(ctx, &contentv1.CollectArticle_Req{
		ArticleId: req.ArticleID,
		Collected: req.Active,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return false, err
	}
	return reply.GetCollected(), nil
}

func (r *ContentArticleClient) RewardArticle(ctx context.Context, req *repo.RewardArticleReq) error {
	_, err := r.contentClient.Article.Reward(ctx, &contentv1.RewardArticle_Req{
		ArticleId: req.ArticleID,
		Points:    req.Points,
		Access: &contentv1.ContentAccess{
			Scope:       contentv1enum.ContentAccessScope_CONTENT_ACCESS_SCOPE_USER,
			ActorUserId: new(req.UserID),
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *ContentArticleClient) loadArticleFacts(ctx context.Context, articleIDs []int64, userID int64, excludedUserIDs []int64) (map[int64]*contentv1.MapArticleLastComments_Resp_Comment, map[int64]*model.ArticleViewerActionState, map[int64][]*model.ArticleTag, map[int64][]*model.ArticleDomain, error) {
	if len(articleIDs) == 0 {
		return map[int64]*contentv1.MapArticleLastComments_Resp_Comment{}, map[int64]*model.ArticleViewerActionState{}, map[int64][]*model.ArticleTag{}, map[int64][]*model.ArticleDomain{}, nil
	}
	commentResp, err := r.contentClient.Comment.MapArticleLastComments(ctx, &contentv1.MapArticleLastComments_Req{
		ArticleIds:      articleIDs,
		ExcludedUserIds: excludedUserIDs,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}
	states := map[int64]*model.ArticleViewerActionState{}
	if userID > 0 {
		stateResp, err := r.contentClient.Article.MapViewerActionStates(ctx, &contentv1.MapArticleViewerActionStates_Req{
			ArticleIds: articleIDs,
			UserId:     userID,
		})
		if err != nil {
			return nil, nil, nil, nil, err
		}
		states = r.articleViewerActionStates(stateResp.GetStates())
	}
	tagStatus := contentv1enum.TagStatus_TAG_STATUS_ENABLED
	tagResp, err := r.contentClient.Tag.MapArticleTags(ctx, &contentv1.MapArticleTags_Req{ArticleIds: articleIDs, Status: &tagStatus})
	if err != nil {
		return nil, nil, nil, nil, err
	}
	articleTags := make(map[int64][]*model.ArticleTag, len(tagResp.GetArticleTags()))
	domainIDs := make([]int64, 0)
	for articleID, rows := range tagResp.GetArticleTags() {
		values := make([]*model.ArticleTag, 0, len(rows.GetRows()))
		for _, item := range rows.GetRows() {
			values = append(values, &model.ArticleTag{ID: item.GetId(), Code: item.GetCode(), Name: item.GetName(), Icon: item.Icon, DomainID: item.DomainId})
			if item.DomainId != nil {
				domainIDs = append(domainIDs, *item.DomainId)
			}
		}
		articleTags[articleID] = values
	}
	domainStatus := contentv1enum.DomainStatus_DOMAIN_STATUS_ENABLED
	domainResp, err := r.contentClient.Domain.Map(ctx, &contentv1.MapDomains_Req{DomainIds: lo.Uniq(domainIDs), Status: &domainStatus})
	if err != nil {
		return nil, nil, nil, nil, err
	}
	articleDomains := make(map[int64][]*model.ArticleDomain, len(articleTags))
	for articleID, tags := range articleTags {
		seen := map[int64]struct{}{}
		for _, tag := range tags {
			if tag.DomainID == nil {
				continue
			}
			domain, ok := domainResp.GetDomains()[*tag.DomainID]
			if !ok {
				tag.DomainID = nil
				continue
			}
			if _, ok := seen[domain.GetId()]; ok {
				continue
			}
			seen[domain.GetId()] = struct{}{}
			articleDomains[articleID] = append(articleDomains[articleID], &model.ArticleDomain{ID: domain.GetId(), Code: domain.GetCode(), Name: domain.GetName(), Icon: domain.Icon, URL: domain.Url, Sort: domain.GetSort()})
		}
		sort.Slice(articleDomains[articleID], func(i, j int) bool {
			if articleDomains[articleID][i].Sort == articleDomains[articleID][j].Sort {
				return articleDomains[articleID][i].ID < articleDomains[articleID][j].ID
			}
			return articleDomains[articleID][i].Sort < articleDomains[articleID][j].Sort
		})
	}
	return commentResp.GetComments(), states, articleTags, articleDomains, nil
}

func (r *ContentArticleClient) articleViewerActionStates(states map[int64]*contentv1.MapArticleViewerActionStates_Resp_ArticleViewerActionState) map[int64]*model.ArticleViewerActionState {
	reply := make(map[int64]*model.ArticleViewerActionState, len(states))
	for articleID, state := range states {
		reply[articleID] = &model.ArticleViewerActionState{
			Liked:     state.GetLiked(),
			Thanked:   state.GetThanked(),
			Collected: state.GetCollected(),
			Rewarded:  state.GetRewarded(),
		}
	}
	return reply
}

func (r *ContentArticleClient) articleProfileIDs(item *contentv1.Article, lastComment *contentv1.MapArticleLastComments_Resp_Comment) []int64 {
	if item == nil {
		return nil
	}
	userIDs := make([]int64, 0, 2)
	if item.CreatedBy != nil {
		userIDs = append(userIDs, *item.CreatedBy)
	}
	if lastComment != nil && lastComment.CreatedBy != nil {

		userIDs = append(userIDs, *lastComment.CreatedBy)
	}
	return userIDs
}

func (r *ContentArticleClient) articleListItem(item *contentv1.Article, profiles map[int64]*model.AccountProfile, lastComment *contentv1.MapArticleLastComments_Resp_Comment, state *model.ArticleViewerActionState, tags []*model.ArticleTag, domains []*model.ArticleDomain) *model.ArticleListItem {
	if item == nil {
		return nil
	}
	if state == nil {
		state = &model.ArticleViewerActionState{}
	}
	content := r.articleSummaryContent(item.GetContent())
	out := &model.ArticleListItem{
		ID:                item.GetId(),
		Title:             item.GetTitle(),
		Content:           content,
		ContentRender:     util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_content", item.GetId()), content),
		HasPostscript:     item.GetHasPostscript(),
		HasReward:         item.GetHasReward(),
		PublishStatus:     int32(item.GetPublishStatus()),
		Visibility:        int32(item.GetVisibility()),
		Restriction:       int32(item.GetRestriction()),
		Type:              int32(item.GetType()),
		Statement:         item.Statement,
		Commentable:       item.GetCommentable(),
		ViewCount:         item.GetViewCount(),
		ThankCount:        item.GetThankCount(),
		LikeCount:         item.GetLikeCount(),
		CollectCount:      item.GetCollectCount(),
		RewardCount:       item.GetRewardCount(),
		ReplyCount:        item.GetReplyCount(),
		CoverImageURL:     r.articleCoverImageURL(item),
		ViewerActionState: state,
		Tags:              tags,
		Domains:           domains,
		CreatedAt:         new(item.GetCreatedAt().AsTime()),
		UpdatedAt:         new(item.GetUpdatedAt().AsTime()),
	}
	if item.GetPublishedAt() != nil {
		out.PublishedAt = new(item.GetPublishedAt().AsTime())
	}
	if item.GetEditedAt() != nil {
		out.EditedAt = new(item.GetEditedAt().AsTime())
	}
	out.CreatedBy = item.CreatedBy
	out.UpdatedBy = item.UpdatedBy
	if item.CreatedBy != nil {
		out.AuthorUser = profiles[*item.CreatedBy]
	}
	if lastComment != nil {
		out.LastReplyAt = new(lastComment.GetCreatedAt().AsTime())
		if lastComment.CreatedBy != nil {
			out.LastReplyUser = profiles[*lastComment.CreatedBy]
		}
	}
	return out
}

func (r *ContentArticleClient) articleDetail(item *contentv1.Article, profiles map[int64]*model.AccountProfile, lastComment *contentv1.MapArticleLastComments_Resp_Comment, state *model.ArticleViewerActionState, tags []*model.ArticleTag, domains []*model.ArticleDomain) *model.ArticleDetail {
	if item == nil {
		return nil
	}
	if state == nil {
		state = &model.ArticleViewerActionState{}
	}
	var rewardContentRender *string
	if item.RewardContent != nil {
		rewardContentRender = new(util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_reward_content", item.GetId()), *item.RewardContent))
	}
	out := &model.ArticleDetail{
		ID:                  item.GetId(),
		Title:               item.GetTitle(),
		Content:             item.GetContent(),
		ContentRender:       util.LuteEngine.MarkdownStr(fmt.Sprintf("%s_%d", "article_content", item.GetId()), item.GetContent()),
		HasPostscript:       item.GetHasPostscript(),
		HasReward:           item.GetHasReward(),
		RewardContent:       item.RewardContent,
		RewardContentRender: rewardContentRender,
		RewardPoints:        item.RewardPoints,
		PublishStatus:       int32(item.GetPublishStatus()),
		Visibility:          int32(item.GetVisibility()),
		Restriction:         int32(item.GetRestriction()),
		Type:                int32(item.GetType()),
		Statement:           item.Statement,
		Commentable:         item.GetCommentable(),
		ViewCount:           item.GetViewCount(),
		ThankCount:          item.GetThankCount(),
		LikeCount:           item.GetLikeCount(),
		CollectCount:        item.GetCollectCount(),
		RewardCount:         item.GetRewardCount(),
		ReplyCount:          item.GetReplyCount(),
		CoverImageURL:       r.articleCoverImageURL(item),
		ViewerActionState:   state,
		Tags:                tags,
		Domains:             domains,
		City:                item.PublishedCity,
		PinnedID:            item.PinnedCommentId,
		AcceptedCommentID:   item.AcceptedCommentId,
		Poll:                r.articlePoll(item.GetPoll()),
		Lottery:             r.articleLottery(item.GetLottery()),
		CreatedAt:           new(item.GetCreatedAt().AsTime()),
		UpdatedAt:           new(item.GetUpdatedAt().AsTime()),
	}
	if item.GetPublishedAt() != nil {
		out.PublishedAt = new(item.GetPublishedAt().AsTime())
	}
	if item.GetEditedAt() != nil {
		out.EditedAt = new(item.GetEditedAt().AsTime())
	}
	out.CreatedBy = item.CreatedBy
	out.UpdatedBy = item.UpdatedBy
	if item.CreatedBy != nil {
		out.AuthorUser = profiles[*item.CreatedBy]
	}
	if lastComment != nil {
		out.LastReplyAt = new(lastComment.GetCreatedAt().AsTime())
		if lastComment.CreatedBy != nil {
			out.LastReplyUser = profiles[*lastComment.CreatedBy]
		}
	}
	return out
}

func (*ContentArticleClient) articlePoll(row *contentv1.ArticlePoll) *model.ArticlePoll {
	if row == nil {
		return nil
	}
	poll := &model.ArticlePoll{MaxChoices: row.GetMaxChoices(), ParticipantCount: row.GetParticipantCount(), TotalVoteCount: row.GetTotalVoteCount(), SelectedOptionIDs: row.GetSelectedOptionIds(), ResultVisible: row.GetResultVisible(), Options: make([]*model.ArticlePollOption, 0, len(row.GetOptions()))}
	if row.GetEndsAt() != nil {
		poll.EndsAt = row.GetEndsAt().AsTime()
	}
	for _, option := range row.GetOptions() {
		if option != nil {
			poll.Options = append(poll.Options, &model.ArticlePollOption{ID: option.GetId(), Content: option.GetContent(), VoteCount: option.GetVoteCount()})
		}
	}
	return poll
}

func (*ContentArticleClient) articleLottery(row *contentv1.ArticleLottery) *model.ArticleLottery {
	if row == nil {
		return nil
	}
	lottery := &model.ArticleLottery{EntryCount: row.GetEntryCount(), Entered: row.GetEntered(), Prizes: make([]*model.ArticleLotteryPrize, 0, len(row.GetPrizes()))}
	if row.GetDrawAt() != nil {
		lottery.DrawAt = row.GetDrawAt().AsTime()
	}
	if row.GetDrawnAt() != nil {
		lottery.DrawnAt = new(row.GetDrawnAt().AsTime())
	}
	for _, prize := range row.GetPrizes() {
		if prize == nil {
			continue
		}
		lottery.Prizes = append(lottery.Prizes, &model.ArticleLotteryPrize{ID: prize.GetId(), Name: prize.GetName(), Description: prize.Description, WinnerLimit: prize.GetWinnerLimit(), WinnerUserIDs: prize.GetWinnerUserIds()})
	}
	return lottery
}

func (r *ContentArticleClient) articleSummaryContent(content string) string {
	runes := []rune(content)
	if len(runes) > 200 {
		return string(runes[:200]) + "..."
	}
	return content
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
