package repo

import (
	"context"
	"crypto/rand"
	"math/big"
	"time"

	utilent "common/pkg/util/ent"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	articlelotteryent "content/internal/data/gen/articlelottery"
	articlelotteryentryent "content/internal/data/gen/articlelotteryentry"
	articlelotteryprizeent "content/internal/data/gen/articlelotteryprize"
	articlelotterywinnerent "content/internal/data/gen/articlelotterywinner"

	"entgo.io/ent/dialect/sql"
)

type ArticleLotteryRepo struct {
	db *gen.Client
}

func NewArticleLotteryRepo(db *gen.Client) repo.ArticleLotteryRepo {
	return &ArticleLotteryRepo{db: db}
}

func (r *ArticleLotteryRepo) client(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ArticleLotteryRepo) SaveDraft(ctx context.Context, articleID int64, draft *model.ArticleLotteryDraft) error {
	if err := r.DeleteDraft(ctx, articleID); err != nil {
		return err
	}
	client := r.client(ctx)
	lottery, err := client.ArticleLottery.Create().SetArticleID(articleID).SetDrawAt(draft.DrawAt).Save(ctx)
	if err != nil {
		return err
	}
	prizes := make([]*gen.ArticleLotteryPrizeCreate, 0, len(draft.Prizes))
	for sort, prize := range draft.Prizes {
		prizes = append(prizes, client.ArticleLotteryPrize.Create().SetLotteryID(lottery.ID).SetName(prize.Name).SetNillableDescription(prize.Description).SetWinnerLimit(prize.WinnerLimit).SetSort(int32(sort)))
	}
	return client.ArticleLotteryPrize.CreateBulk(prizes...).Exec(ctx)
}

func (r *ArticleLotteryRepo) DeleteDraft(ctx context.Context, articleID int64) error {
	client := r.client(ctx)
	lottery, err := client.ArticleLottery.Query().Where(articlelotteryent.ArticleIDEQ(articleID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if _, err = client.ArticleLotteryWinner.Delete().Where(articlelotterywinnerent.LotteryIDEQ(lottery.ID)).Exec(ctx); err != nil {
		return err
	}
	if _, err = client.ArticleLotteryEntry.Delete().Where(articlelotteryentryent.LotteryIDEQ(lottery.ID)).Exec(ctx); err != nil {
		return err
	}
	if _, err = client.ArticleLotteryPrize.Delete().Where(articlelotteryprizeent.LotteryIDEQ(lottery.ID)).Exec(ctx); err != nil {
		return err
	}
	return client.ArticleLottery.DeleteOneID(lottery.ID).Exec(ctx)
}

func (r *ArticleLotteryRepo) Get(ctx context.Context, articleID, viewerID int64) (*model.ArticleLottery, error) {
	client := r.client(ctx)
	lottery, err := client.ArticleLottery.Query().Where(articlelotteryent.ArticleIDEQ(articleID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	prizes, err := client.ArticleLotteryPrize.Query().Where(articlelotteryprizeent.LotteryIDEQ(lottery.ID)).Order(gen.Asc(articlelotteryprizeent.FieldSort)).All(ctx)
	if err != nil {
		return nil, err
	}
	row := &model.ArticleLottery{ID: lottery.ID, DrawAt: lottery.DrawAt, DrawnAt: lottery.DrawnAt, EntryCount: lottery.EntryCount, Prizes: make([]*model.ArticleLotteryPrize, 0, len(prizes))}
	for _, prize := range prizes {
		item := &model.ArticleLotteryPrize{ID: prize.ID, Name: prize.Name, Description: prize.Description, WinnerLimit: prize.WinnerLimit}
		winners, queryErr := client.ArticleLotteryWinner.Query().Where(articlelotterywinnerent.PrizeIDEQ(prize.ID)).Order(gen.Asc(articlelotterywinnerent.FieldSort)).All(ctx)
		if queryErr != nil {
			return nil, queryErr
		}
		for _, winner := range winners {
			item.WinnerUserIDs = append(item.WinnerUserIDs, winner.UserID)
		}
		row.Prizes = append(row.Prizes, item)
	}
	if viewerID <= 0 {
		return row, nil
	}
	row.Entered, err = client.ArticleLotteryEntry.Query().Where(articlelotteryentryent.LotteryIDEQ(lottery.ID), articlelotteryentryent.UserIDEQ(viewerID)).Exist(ctx)
	return row, err
}

func (r *ArticleLotteryRepo) SetEntry(ctx context.Context, lotteryID, userID int64, entered bool, now time.Time) error {
	client := r.client(ctx)
	lottery, err := client.ArticleLottery.Query().Where(articlelotteryent.IDEQ(lotteryID), func(selector *sql.Selector) { selector.ForUpdate() }).Only(ctx)
	if err != nil {
		return err
	}
	if lottery.DrawnAt != nil || !now.Before(lottery.DrawAt) {
		return &repo.ArticleLotteryClosedError{}
	}
	exists, err := client.ArticleLotteryEntry.Query().Where(articlelotteryentryent.LotteryIDEQ(lotteryID), articlelotteryentryent.UserIDEQ(userID)).Exist(ctx)
	if err != nil {
		return err
	}
	if entered && !exists {
		if _, err = client.ArticleLotteryEntry.Create().SetLotteryID(lotteryID).SetUserID(userID).Save(ctx); err != nil {
			return err
		}
		_, err = client.ArticleLottery.UpdateOneID(lotteryID).AddEntryCount(1).Save(ctx)
		return err
	}
	if !entered && exists {
		if _, err = client.ArticleLotteryEntry.Delete().Where(articlelotteryentryent.LotteryIDEQ(lotteryID), articlelotteryentryent.UserIDEQ(userID)).Exec(ctx); err != nil {
			return err
		}
		_, err = client.ArticleLottery.UpdateOneID(lotteryID).AddEntryCount(-1).Save(ctx)
		return err
	}
	return nil
}

func (r *ArticleLotteryRepo) Draw(ctx context.Context, lotteryID int64, drawnAt time.Time) ([]*repo.ArticleLotteryWinner, error) {
	client := r.client(ctx)
	lottery, err := client.ArticleLottery.Query().Where(articlelotteryent.IDEQ(lotteryID), func(selector *sql.Selector) { selector.ForUpdate() }).Only(ctx)
	if err != nil {
		return nil, err
	}
	if lottery.DrawnAt != nil {
		return nil, nil
	}
	prizes, err := client.ArticleLotteryPrize.Query().Where(articlelotteryprizeent.LotteryIDEQ(lotteryID)).Order(gen.Asc(articlelotteryprizeent.FieldSort)).All(ctx)
	if err != nil {
		return nil, err
	}
	entries, err := client.ArticleLotteryEntry.Query().Where(articlelotteryentryent.LotteryIDEQ(lotteryID)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*repo.ArticleLotteryWinner, 0)
	for _, prize := range prizes {
		for winnerSort := int32(0); winnerSort < prize.WinnerLimit && len(entries) > 0; winnerSort++ {
			index, randomErr := rand.Int(rand.Reader, big.NewInt(int64(len(entries))))
			if randomErr != nil {
				return nil, randomErr
			}
			entry := entries[index.Int64()]
			entries = append(entries[:index.Int64()], entries[index.Int64()+1:]...)
			if _, err = client.ArticleLotteryWinner.Create().SetLotteryID(lotteryID).SetPrizeID(prize.ID).SetUserID(entry.UserID).SetSort(winnerSort).Save(ctx); err != nil {
				return nil, err
			}
			result = append(result, &repo.ArticleLotteryWinner{PrizeID: prize.ID, UserID: entry.UserID})
		}
	}
	_, err = client.ArticleLottery.UpdateOneID(lotteryID).SetDrawnAt(drawnAt).Save(ctx)
	return result, err
}
