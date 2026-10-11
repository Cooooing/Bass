package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"slices"
	"strings"
	"time"
	"unicode/utf8"

	"content/internal/biz/model"
	"content/internal/biz/repo"
)

type MoonbreezeUsecase struct {
	moonbreezeRepo repo.MoonbreezeRepo
	rateLimitCache repo.MoonbreezeRateLimitCache
	maxCharacters  int
	rateWindow     time.Duration
	rateMaximum    int64
}

func NewMoonbreezeUsecase(
	moonbreezeRepo repo.MoonbreezeRepo,
	rateLimitCache repo.MoonbreezeRateLimitCache,
) *MoonbreezeUsecase {
	return &MoonbreezeUsecase{
		moonbreezeRepo: moonbreezeRepo,
		rateLimitCache: rateLimitCache,
		maxCharacters:  512,
		rateWindow:     time.Hour,
		rateMaximum:    5,
	}
}

type CreateMoonbreezeReq struct {
	Content  string
	AuthorID int64
	City     *string
}

func (u *MoonbreezeUsecase) Create(ctx context.Context, req *CreateMoonbreezeReq) (*model.Moonbreeze, error) {
	if req == nil || req.AuthorID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_MOONBREEZE_INVALID)
	}
	content := strings.TrimSpace(req.Content)
	if content == "" || strings.ContainsAny(content, "\r\n") || utf8.RuneCountInString(content) > u.maxCharacters {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_MOONBREEZE_INVALID)
	}
	state, err := u.rateLimitCache.Allow(ctx, req.AuthorID, u.rateWindow, u.rateMaximum)
	if err != nil {
		return nil, err
	}
	if !state.Allowed {
		retryAfterSeconds := int64((state.RetryAfter + time.Second - 1) / time.Second)
		return nil, apperror.New(
			cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_MOONBREEZE_RATE_LIMITED,
			apperror.WithData(&cerrors.RetryAfterErrorData{
				RetryAfterSeconds: retryAfterSeconds,
			}),
		)
	}
	city := req.City
	if city != nil {
		value := strings.TrimSpace(*city)
		if value == "" {
			city = nil
		} else {
			// City comes from a provider rather than user input. Truncating its
			// display snapshot keeps a verbose provider value from failing a post.
			characters := []rune(value)
			if len(characters) > 128 {
				value = string(characters[:128])
			}
			city = &value
		}
	}
	return u.moonbreezeRepo.Create(ctx, &model.Moonbreeze{
		Content:  content,
		AuthorID: req.AuthorID,
		City:     city,
	})
}

type PageMoonbreezesReq struct {
	AuthorIDs         []int64
	ExcludedAuthorIDs []int64
	Cursor            *string
	Size              int
}

type PageMoonbreezesResp struct {
	Rows       []*model.Moonbreeze
	NextCursor *string
}

func (u *MoonbreezeUsecase) Page(ctx context.Context, req *PageMoonbreezesReq) (*PageMoonbreezesResp, error) {
	if req == nil {
		req = &PageMoonbreezesReq{}
	}
	beforeAt, beforeID, err := u.decodeCursor(req.Cursor, req.ExcludedAuthorIDs)
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_MOONBREEZE_INVALID)
	}
	size := req.Size
	if size <= 0 || size > 50 {
		size = 20
	}
	// Reading one additional record tells callers whether a continuation exists.
	// It also keeps the public cursor positioned exactly after the last returned
	// row, which is essential when multiple rows share one created_at value.
	page, err := u.moonbreezeRepo.Page(ctx, &repo.MoonbreezePageReq{
		AuthorIDs:         req.AuthorIDs,
		ExcludedAuthorIDs: req.ExcludedAuthorIDs,
		BeforeAt:          beforeAt,
		BeforeID:          beforeID,
		Size:              size + 1,
	})
	if err != nil {
		return nil, err
	}
	result := &PageMoonbreezesResp{Rows: page.Rows}
	if len(page.Rows) > size {
		result.Rows = page.Rows[:size]
		result.NextCursor = u.encodeCursor(result.Rows[len(result.Rows)-1], req.ExcludedAuthorIDs)
	}
	return result, nil
}

// MoonbreezeCursor is the opaque cursor payload for stable dynamic pagination.
type MoonbreezeCursor struct {
	CreatedAt          time.Time `json:"created_at"`
	ID                 int64     `json:"id"`
	ExcludedAuthorHash string    `json:"excluded_author_hash"`
}

func (u *MoonbreezeUsecase) decodeCursor(value *string, excludedAuthorIDs []int64) (*time.Time, *int64, error) {
	if value == nil || *value == "" {
		return nil, nil, nil
	}
	decoded, err := base64.RawURLEncoding.DecodeString(*value)
	if err != nil {
		return nil, nil, err
	}
	var cursor MoonbreezeCursor
	if err = json.Unmarshal(decoded, &cursor); err != nil || cursor.ID <= 0 || cursor.CreatedAt.IsZero() || cursor.ExcludedAuthorHash != u.excludedAuthorHash(excludedAuthorIDs) {
		return nil, nil, errors.New("invalid moonbreeze cursor")
	}
	return &cursor.CreatedAt, &cursor.ID, nil
}

func (u *MoonbreezeUsecase) encodeCursor(row *model.Moonbreeze, excludedAuthorIDs []int64) *string {
	if row == nil || row.CreatedAt == nil {
		return nil
	}
	encoded, err := json.Marshal(MoonbreezeCursor{
		CreatedAt:          row.CreatedAt.UTC(),
		ID:                 row.ID,
		ExcludedAuthorHash: u.excludedAuthorHash(excludedAuthorIDs),
	})
	if err != nil {
		return nil
	}
	value := base64.RawURLEncoding.EncodeToString(encoded)
	return &value
}

func (*MoonbreezeUsecase) excludedAuthorHash(userIDs []int64) string {
	values := append([]int64(nil), userIDs...)
	slices.Sort(values)
	encoded, err := json.Marshal(values)
	if err != nil {
		return ""
	}
	sum := sha256.Sum256(encoded)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}
