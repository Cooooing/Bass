package data

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/util/collections/dict"
	"common/pkg/util/collections/set"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/article"
	"content/internal/data/ent/gen/tag"
	"context"
	"encoding/json"
	"math"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleRepo struct {
	*BaseRepo
	client      *gen.Client
	commentRepo repo.CommentRepo
	domainRepo  repo.DomainRepo
	tagRepo     repo.TagRepo
}

func NewArticleRepo(baseRepo *BaseRepo, client *gen.Client, commentRepo repo.CommentRepo, domainRepo repo.DomainRepo, tagRepo repo.TagRepo) repo.ArticleRepo {
	return &ArticleRepo{
		BaseRepo:    baseRepo,
		client:      client,
		commentRepo: commentRepo,
		domainRepo:  domainRepo,
		tagRepo:     tagRepo,
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
		WithPostscripts().
		WithTags().
		WithComments().
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

	lastComment, _ := r.commentRepo.GetArticleLastComment(ctx, tx, query.ID)
	if lastComment != nil {
		a.RepliedAt = timestamppb.New(*lastComment.CreatedAt)
	}

	userServiceClient, err := client.GetServiceClient(ctx, r.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
	if err != nil {
		return nil, err
	}
	userAuthor, err := userServiceClient.GetOne(ctx, &userv1.GetOneRequest{
		Id: query.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetArticleOneReply{
		Article: a,
		User:    userAuthor.User,
	}, nil
}

func (r *ArticleRepo) GetList(ctx context.Context, tx *gen.Client, req *v1.GetArticleRequest) (*v1.GetArticleReply, error) {
	tagIds := set.New[int64](0)
	if req.DomainId != nil {
		tags, err := r.tagRepo.GetList(ctx, tx, &v1.GetTagRequest{
			Page:     &cv1.PageRequest{Page: 1, Size: math.MaxInt32},
			DomainId: *req.DomainId,
		})
		if err != nil {
			return nil, err
		}
		for _, item := range tags.Tags {
			tagIds.Add(item.Id)
		}
	}
	if req.TagId != nil {
		tagIds.Add(*req.TagId)
	}

	query := tx.Article.Query()
	if req.Status != nil {
		query = query.Where(article.StatusEQ(*req.Status))
	}
	if tagIds.Len() > 0 {
		query = query.Where(article.HasTagsWith(tag.IDIn(tagIds.ToSlice()...)))
	}
	// Todo 文章排序，默认最新。最热暂不实现
	query = query.Order(gen.Desc(article.FieldCreatedAt))
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(int(req.Page.Size)).Offset(int((req.Page.Page - 1) * req.Page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}

	articleIds := set.New[int64](0)
	userIds := set.New[int64](0)
	for _, item := range list {
		articleIds.Add(item.ID)
		userIds.Add(item.UserID)
	}

	lastCommentMap, _ := r.commentRepo.GetArticleLastComments(ctx, tx, articleIds.ToSlice())
	lastCommentMap.Foreach(func(e *dict.Entry[int64, *model.Comment]) bool {
		userIds.Add(e.Value.UserID)
		return true
	})

	userServiceClient, err := client.GetServiceClient(ctx, r.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
	if err != nil {
		return nil, err
	}
	userAuthors, err := userServiceClient.GetMap(ctx, &userv1.GetMapRequest{
		Ids: userIds.ToSlice(),
	})
	if err != nil {
		return nil, err
	}

	articles := make([]*v1.Article, 0, len(list))
	for _, item := range list {
		i := (*model.Article)(item)
		i.Summary()
		a := &v1.Article{}
		err = copier.Copy(a, i)
		if err != nil {
			return nil, err
		}
		a.CreatedAt = timestamppb.New(*i.CreatedAt)
		a.UpdatedAt = timestamppb.New(*i.UpdatedAt)
		if lastReplyComment, ok := lastCommentMap.Get(item.ID); ok {
			a.RepliedAt = timestamppb.New(*lastReplyComment.CreatedAt)
			a.ReplyUser = userAuthors.Users[lastReplyComment.UserID]
		}
		a.AuthorUser = userAuthors.Users[item.UserID]
		articles = append(articles, a)
	}

	return &v1.GetArticleReply{
		Page: &cv1.PageReply{
			Total: uint32(count),
			Page:  req.Page.Page,
			Size:  req.Page.Size,
		},
		Articles: articles,
	}, nil
}
