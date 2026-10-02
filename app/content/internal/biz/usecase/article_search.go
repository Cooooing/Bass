package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"sort"
	"strings"
	"time"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"
)

const (
	ArticleSearchDefaultPageSize = 20
	ArticleSearchMaxPageSize     = 50
	ArticleSearchMinQueryRunes   = 2
	ArticleSearchMaxQueryRunes   = 128
)

type ArticleSearchUsecase struct {
	articleRepo       repo.ArticleRepo
	articleSearchRepo repo.ArticleSearchRepo
}

func NewArticleSearchUsecase(
	articleRepo repo.ArticleRepo,
	articleSearchRepo repo.ArticleSearchRepo,
) *ArticleSearchUsecase {
	return &ArticleSearchUsecase{
		articleRepo:       articleRepo,
		articleSearchRepo: articleSearchRepo,
	}
}

type SearchPublicArticlesReq struct {
	Query             string
	Cursor            *string
	Size              int
	ExcludedAuthorIDs []int64
}

type SearchPublicArticlesResp struct {
	Rows       []*model.Article
	NextCursor *string
}

// SearchPublic performs a public literal keyword search through the configured
// search backend. It rechecks article visibility after resolving result IDs so
// a concurrent visibility change cannot leak through a stale index result.
func (u *ArticleSearchUsecase) SearchPublic(
	ctx context.Context,
	req *SearchPublicArticlesReq,
) (*SearchPublicArticlesResp, error) {
	if req == nil {
		req = &SearchPublicArticlesReq{}
	}
	query := strings.TrimSpace(req.Query)
	if !u.validQuery(query) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	cursor, err := u.decodeCursor(req.Cursor, query, req.ExcludedAuthorIDs)
	if err != nil {
		return nil, err
	}
	search, err := u.articleSearchRepo.SearchPublic(ctx, &repo.SearchPublicArticlesReq{
		Query:  query,
		Cursor: cursor,
		Size:   u.pageSize(req.Size),
	})
	if err != nil {
		return nil, err
	}
	articleIDs := make([]int64, 0, len(search.Rows))
	for _, row := range search.Rows {
		if row != nil && row.ArticleID > 0 {
			articleIDs = append(articleIDs, row.ArticleID)
		}
	}
	published := enum.ArticlePublishStatusPublished
	public := enum.ArticleVisibilityPublic
	articles, err := u.articleRepo.Map(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{
		ArticleIDs:        articleIDs,
		ExcludedAuthorIDs: req.ExcludedAuthorIDs,
		PublishStatus:     &published,
		Visibility:        &public,
		Restrictions: []enum.ContentRestriction{
			enum.ContentRestrictionNone,
			enum.ContentRestrictionLocked,
		},
	}})
	if err != nil {
		return nil, err
	}
	result := &SearchPublicArticlesResp{Rows: make([]*model.Article, 0, len(search.Rows))}
	var last *repo.ArticleSearchRow
	for _, row := range search.Rows {
		if row == nil {
			continue
		}
		last = row
		if article := articles[row.ArticleID]; article != nil {
			result.Rows = append(result.Rows, article)
		}
	}
	if search.HasMore && last != nil {
		result.NextCursor = u.encodeCursor(query, req.ExcludedAuthorIDs, &repo.ArticleSearchCursor{
			Rank:        last.Rank,
			PublishedAt: last.PublishedAt,
			ArticleID:   last.ArticleID,
		})
	}
	return result, nil
}

type ArticleSearchCursorPayload struct {
	QueryHash          string    `json:"query_hash"`
	ExcludedAuthorHash string    `json:"excluded_author_hash"`
	Rank               int       `json:"rank"`
	PublishedAt        time.Time `json:"published_at"`
	ArticleID          int64     `json:"article_id"`
}

func (u *ArticleSearchUsecase) validQuery(query string) bool {
	length := len([]rune(query))
	return length >= ArticleSearchMinQueryRunes && length <= ArticleSearchMaxQueryRunes
}

func (u *ArticleSearchUsecase) pageSize(size int) int {
	if size <= 0 {
		return ArticleSearchDefaultPageSize
	}
	if size > ArticleSearchMaxPageSize {
		return ArticleSearchMaxPageSize
	}
	return size
}

func (u *ArticleSearchUsecase) decodeCursor(value *string, query string, excludedAuthorIDs []int64) (*repo.ArticleSearchCursor, error) {
	if value == nil || strings.TrimSpace(*value) == "" {
		return nil, nil
	}
	encoded, err := base64.RawURLEncoding.DecodeString(*value)
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var cursor ArticleSearchCursorPayload
	if err = json.Unmarshal(encoded, &cursor); err != nil ||
		cursor.QueryHash != u.queryHash(query) ||
		cursor.ExcludedAuthorHash != u.excludedAuthorHash(excludedAuthorIDs) ||
		cursor.Rank < 0 || cursor.Rank > 2 || cursor.ArticleID <= 0 || cursor.PublishedAt.IsZero() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return &repo.ArticleSearchCursor{
		Rank:        cursor.Rank,
		PublishedAt: cursor.PublishedAt,
		ArticleID:   cursor.ArticleID,
	}, nil
}

func (u *ArticleSearchUsecase) encodeCursor(query string, excludedAuthorIDs []int64, cursor *repo.ArticleSearchCursor) *string {
	if cursor == nil || cursor.ArticleID <= 0 || cursor.PublishedAt.IsZero() {
		return nil
	}
	encoded, err := json.Marshal(ArticleSearchCursorPayload{
		QueryHash:          u.queryHash(query),
		ExcludedAuthorHash: u.excludedAuthorHash(excludedAuthorIDs),
		Rank:               cursor.Rank,
		PublishedAt:        cursor.PublishedAt.UTC(),
		ArticleID:          cursor.ArticleID,
	})
	if err != nil {
		return nil
	}
	value := base64.RawURLEncoding.EncodeToString(encoded)
	return &value
}

func (*ArticleSearchUsecase) excludedAuthorHash(userIDs []int64) string {
	values := append([]int64(nil), userIDs...)
	sort.Slice(values, func(i int, j int) bool { return values[i] < values[j] })
	encoded, _ := json.Marshal(values)
	sum := sha256.Sum256(encoded)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func (*ArticleSearchUsecase) queryHash(query string) string {
	sum := sha256.Sum256([]byte(query))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}
