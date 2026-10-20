package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"

	utilent "common/pkg/util/ent"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	articlepollent "content/internal/data/gen/articlepoll"
	articlepolloptionent "content/internal/data/gen/articlepolloption"
	articlepollvoteent "content/internal/data/gen/articlepollvote"
)

type ArticlePollRepo struct {
	db *gen.Client
}

func NewArticlePollRepo(db *gen.Client) repo.ArticlePollRepo {
	return &ArticlePollRepo{db: db}
}

func (r *ArticlePollRepo) client(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ArticlePollRepo) SaveDraft(ctx context.Context, articleID int64, draft *model.ArticlePollDraft) error {
	client := r.client(ctx)
	previous, err := client.ArticlePoll.Query().Where(articlepollent.ArticleIDEQ(articleID)).Only(ctx)
	if err == nil {
		if err = r.delete(ctx, previous.ID); err != nil {
			return err
		}
	} else if !gen.IsNotFound(err) {
		return err
	}
	poll, err := client.ArticlePoll.Create().SetArticleID(articleID).SetEndsAt(draft.EndsAt).SetMaxChoices(draft.MaxChoices).Save(ctx)
	if err != nil {
		return err
	}
	builders := make([]*gen.ArticlePollOptionCreate, 0, len(draft.Options))
	for sort, option := range draft.Options {
		builders = append(builders, client.ArticlePollOption.Create().SetPollID(poll.ID).SetContent(option).SetSort(int32(sort)))
	}
	return client.ArticlePollOption.CreateBulk(builders...).Exec(ctx)
}

func (r *ArticlePollRepo) DeleteDraft(ctx context.Context, articleID int64) error {
	poll, err := r.client(ctx).ArticlePoll.Query().Where(articlepollent.ArticleIDEQ(articleID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return r.delete(ctx, poll.ID)
}

func (r *ArticlePollRepo) delete(ctx context.Context, pollID int64) error {
	client := r.client(ctx)
	if _, err := client.ArticlePollVote.Delete().Where(articlepollvoteent.PollIDEQ(pollID)).Exec(ctx); err != nil {
		return err
	}
	if _, err := client.ArticlePollOption.Delete().Where(articlepolloptionent.PollIDEQ(pollID)).Exec(ctx); err != nil {
		return err
	}
	return client.ArticlePoll.DeleteOneID(pollID).Exec(ctx)
}

func (r *ArticlePollRepo) Get(ctx context.Context, articleID, viewerID int64) (*model.ArticlePoll, error) {
	client := r.client(ctx)
	poll, err := client.ArticlePoll.Query().Where(articlepollent.ArticleIDEQ(articleID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	options, err := client.ArticlePollOption.Query().Where(articlepolloptionent.PollIDEQ(poll.ID)).Order(gen.Asc(articlepolloptionent.FieldSort)).All(ctx)
	if err != nil {
		return nil, err
	}
	row := &model.ArticlePoll{ID: poll.ID, EndsAt: poll.EndsAt, MaxChoices: poll.MaxChoices, ParticipantCount: poll.ParticipantCount, TotalVoteCount: poll.TotalVoteCount, Options: make([]*model.ArticlePollOption, 0, len(options))}
	for _, option := range options {
		row.Options = append(row.Options, &model.ArticlePollOption{ID: option.ID, Content: option.Content, VoteCount: option.VoteCount})
	}
	if viewerID <= 0 {
		return row, nil
	}
	vote, err := client.ArticlePollVote.Query().Where(articlepollvoteent.PollIDEQ(poll.ID), articlepollvoteent.UserIDEQ(viewerID)).Only(ctx)
	if gen.IsNotFound(err) {
		return row, nil
	}
	if err != nil {
		return nil, err
	}
	row.SelectedOptionIDs = vote.OptionIds
	return row, nil
}

func (r *ArticlePollRepo) Vote(ctx context.Context, pollID, userID int64, optionIDs []int64) error {
	client := r.client(ctx)
	if _, err := client.ArticlePollVote.Create().SetPollID(pollID).SetUserID(userID).SetOptionIds(optionIDs).Save(ctx); err != nil {
		if gen.IsConstraintError(err) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		return err
	}
	if _, err := client.ArticlePoll.UpdateOneID(pollID).AddParticipantCount(1).AddTotalVoteCount(int64(len(optionIDs))).Save(ctx); err != nil {
		return err
	}
	for _, optionID := range optionIDs {
		if _, err := client.ArticlePollOption.UpdateOneID(optionID).AddVoteCount(1).Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (r *ArticlePollRepo) ListParticipantUserIDs(ctx context.Context, articleID int64) ([]int64, error) {
	poll, err := r.client(ctx).ArticlePoll.Query().Where(articlepollent.ArticleIDEQ(articleID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	votes, err := r.client(ctx).ArticlePollVote.Query().Where(articlepollvoteent.PollIDEQ(poll.ID)).All(ctx)
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, 0, len(votes))
	for _, vote := range votes {
		userIDs = append(userIDs, vote.UserID)
	}
	return userIDs, nil
}
