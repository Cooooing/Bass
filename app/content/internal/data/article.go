package data

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/util/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/article"
	"content/internal/data/ent/gen/articlepostscript"
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleRepo struct {
	*BaseRepo
	client         *gen.Client
	postscriptRepo repo.ArticlePostscriptRepo
	commentRepo    repo.CommentRepo
	domainRepo     repo.DomainRepo
	tagRepo        repo.TagRepo
}

func NewArticleRepo(baseRepo *BaseRepo, client *gen.Client, postscriptRepo repo.ArticlePostscriptRepo, commentRepo repo.CommentRepo, domainRepo repo.DomainRepo, tagRepo repo.TagRepo) repo.ArticleRepo {
	return &ArticleRepo{
		BaseRepo:       baseRepo,
		client:         client,
		postscriptRepo: postscriptRepo,
		commentRepo:    commentRepo,
		domainRepo:     domainRepo,
		tagRepo:        tagRepo,
	}
}

func (r *ArticleRepo) Save(ctx context.Context, client *gen.Client, article *model.Article) (*model.Article, error) {
	save, err := client.Article.Create().
		SetUserID(article.UserID).
		SetTitle(article.Title).
		SetContent(article.Content).
		SetNillableRewardContent(article.RewardContent).
		SetRewardPoints(article.RewardPoints).
		SetStatus(article.Status).
		SetType(article.Type).
		SetBountyPoints(article.BountyPoints).
		Save(ctx)
	return (*model.Article)(save), err
}

func (r *ArticleRepo) UpdateContent(ctx context.Context, client *gen.Client, articleId int64, content string) error {
	return client.Article.UpdateOneID(articleId).
		SetContent(content).
		Exec(ctx)
}
func (r *ArticleRepo) UpdateStatus(ctx context.Context, client *gen.Client, articleId int64, status cv1.ArticleStatus) error {
	return client.Article.UpdateOneID(articleId).
		SetStatus(int32(status)).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateHasPostscript(ctx context.Context, client *gen.Client, articleId int64, hasPostscript bool) error {
	return client.Article.UpdateOneID(articleId).
		SetHasPostscript(hasPostscript).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateStat(ctx context.Context, client *gen.Client, articleId int64, action cv1.ArticleAction, num int32) error {
	updateOne := client.Article.UpdateOneID(articleId)
	switch action {
	case cv1.ArticleAction_ArticleActionLike:
		updateOne.AddLikeCount(num)
	case cv1.ArticleAction_ArticleActionThank:
		updateOne.AddThankCount(num)
	case cv1.ArticleAction_ArticleActionCollect:
		updateOne.AddCollectCount(num)
	case cv1.ArticleAction_ArticleActionWatch:
		updateOne.AddWatchCount(num)
	case cv1.ArticleAction_ArticleActionReply:
		updateOne.AddReplyCount(num)
	case cv1.ArticleAction_ArticleActionVote:
		updateOne.AddVoteTotal(num)
	case cv1.ArticleAction_ArticleActionLottery:
		updateOne.AddLotteryParticipantCount(num)
	case cv1.ArticleAction_ArticleActionLotteryWinner:
		updateOne.AddLotteryWinnerCount(num)
	default:
		return nil
	}
	return updateOne.Exec(ctx)
}

func (r *ArticleRepo) Publish(ctx context.Context, client *gen.Client, articleId int64) error {
	first, err := r.GetArticleById(ctx, client, articleId)
	if err != nil {
		return err
	}
	publish := &v1.ArticleEventPublish{}
	err = copier.Copy(&publish, first)
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(publish)
	if err != nil {
		return err
	}
	err = r.rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyArticleCreate.String(), marshal)
	return err
}

func (r *ArticleRepo) Delete(ctx context.Context, client *gen.Client, articleId int64) error {
	return client.Article.UpdateOneID(articleId).SetStatus(int32(cv1.ArticleStatus_ArticleDeleted)).Exec(ctx)
}

func (r *ArticleRepo) Exist(ctx context.Context, client *gen.Client, id int64, status cv1.ArticleStatus) (bool, error) {
	return client.Article.Query().
		Where(article.IDEQ(id)).
		Where(article.StatusEQ(int32(status))).
		Exist(ctx)
}

func (r *ArticleRepo) GetArticleById(ctx context.Context, client *gen.Client, id int64) (*model.Article, error) {
	query, err := client.Article.Query().
		Where(article.IDEQ(id)).
		WithPostscripts(func(q *gen.ArticlePostscriptQuery) {
			q.Where(articlepostscript.StatusEQ(int32(cv1.ArticlePostscriptStatus_ArticlePostscriptNormal))).
				Order(gen.Asc(articlepostscript.FieldCreatedAt))
		}).
		WithTags().
		First(ctx)
	return (*model.Article)(query), err
}

func (r *ArticleRepo) GetOne(ctx context.Context, tx *gen.Client, articleId int64) (*v1.GetArticleOneReply, error) {
	query, err := r.GetArticleById(ctx, tx, articleId)
	if err != nil {
		return nil, err
	}

	a := &v1.Article{}
	err = copier.Copy(a, query)
	if err != nil {
		return nil, err
	}
	a.CreatedAt = timestamppb.New(*query.CreatedAt)
	a.UpdatedAt = timestamppb.New(*query.UpdatedAt)
	ap := make([]*v1.ArticlePostscript, 0)
	for _, item := range (*gen.Article)(query).Edges.Postscripts {
		ap = append(ap, &v1.ArticlePostscript{
			Id:        item.ID,
			ArticleId: item.ArticleID,
			Content:   item.Content,
			CreatedAt: timestamppb.New(*item.CreatedAt),
			UpdatedAt: timestamppb.New(*item.UpdatedAt),
		})
	}
	a.Postscripts = ap

	lastComment, _ := r.commentRepo.GetArticleLastComment(ctx, tx, query.ID)
	if lastComment != nil {
		a.RepliedAt = timestamppb.New(*lastComment.CreatedAt)
	}

	userServiceClient, err := client.GetServiceClient(ctx, r.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
	if err != nil {
		return nil, err
	}
	userIds := []int64{query.UserID}
	if lastComment != nil {
		userIds = append(userIds, *lastComment.CreatedBy)
	}
	userAuthorsMap, err := userServiceClient.GetMap(ctx, &userv1.GetMapRequest{
		Ids: userIds,
	})
	if err != nil {
		return nil, err
	}

	if lastComment != nil {
		a.ReplyUser = userAuthorsMap.Users[*lastComment.CreatedBy]
	}

	return &v1.GetArticleOneReply{
		Article: a,
		User:    userAuthorsMap.Users[query.UserID],
	}, nil
}

func (r *ArticleRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.ArticleGetReq) ([]*model.Article, error) {
	var (
		articles []*model.Article
		err      error
	)
	query := tx.Article.Query().WithTags()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		articles = append(articles, (*model.Article)(list[i]))
	}
	return articles, nil
}
func (r *ArticleRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.ArticleGetReq) ([]*model.Article, *cv1.PageReply, error) {
	var (
		articles []*model.Article
		err      error
		total    int
	)
	page = base.OrDefault(page, constant.GetPageDefault())
	query := tx.Article.Query().WithTags()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	total, err = countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	for i := range list {
		articles = append(articles, (*model.Article)(list[i]))
	}
	return articles, &cv1.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *ArticleRepo) getQuery(query *gen.ArticleQuery, req *repo.ArticleGetReq) *gen.ArticleQuery {
	return query
}
