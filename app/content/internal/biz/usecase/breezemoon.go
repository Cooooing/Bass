package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	breezemoonMaxCharacters = 512
	breezemoonRateWindow    = time.Hour
	breezemoonRateMaximum   = 5
)

type BreezemoonUsecase struct {
	breezemoonRepo repo.BreezemoonRepo
	rateLimitCache repo.BreezemoonRateLimitCache
}

func NewBreezemoonUsecase(
	breezemoonRepo repo.BreezemoonRepo,
	rateLimitCache repo.BreezemoonRateLimitCache,
) *BreezemoonUsecase {
	return &BreezemoonUsecase{
		breezemoonRepo: breezemoonRepo,
		rateLimitCache: rateLimitCache,
	}
}

type CreateBreezemoonReq struct {
	Content  string
	AuthorID int64
	City     *string
}

func (u *BreezemoonUsecase) Create(ctx context.Context, req *CreateBreezemoonReq) (*model.Breezemoon, error) {
	if req == nil || req.AuthorID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_BREEZEMOON_INVALID)
	}
	content := strings.TrimSpace(req.Content)
	if content == "" || strings.ContainsAny(content, "\r\n") || utf8.RuneCountInString(content) > breezemoonMaxCharacters {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_BREEZEMOON_INVALID)
	}
	state, err := u.rateLimitCache.Allow(ctx, req.AuthorID, breezemoonRateWindow, breezemoonRateMaximum)
	if err != nil {
		return nil, err
	}
	if !state.Allowed {
		retryAfterSeconds := int64((state.RetryAfter + time.Second - 1) / time.Second)
		return nil, apperror.New(
			cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_BREEZEMOON_RATE_LIMITED,
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
	return u.breezemoonRepo.Create(ctx, &model.Breezemoon{
		Content:  content,
		AuthorID: req.AuthorID,
		City:     city,
	})
}

type PageBreezemoonsReq struct {
	AuthorIDs []int64
	Cursor    *string
	Size      int
}

type PageBreezemoonsResp struct {
	Rows       []*model.Breezemoon
	NextCursor *string
}

func (u *BreezemoonUsecase) Page(ctx context.Context, req *PageBreezemoonsReq) (*PageBreezemoonsResp, error) {
	if req == nil {
		req = &PageBreezemoonsReq{}
	}
	beforeAt, beforeID, err := u.decodeCursor(req.Cursor)
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_BREEZEMOON_INVALID)
	}
	size := req.Size
	if size <= 0 || size > 50 {
		size = 20
	}
	// Reading one additional record tells callers whether a continuation exists.
	// It also keeps the public cursor positioned exactly after the last returned
	// row, which is essential when multiple rows share one created_at value.
	page, err := u.breezemoonRepo.Page(ctx, &repo.BreezemoonPageReq{
		AuthorIDs: req.AuthorIDs,
		BeforeAt:  beforeAt,
		BeforeID:  beforeID,
		Size:      size + 1,
	})
	if err != nil {
		return nil, err
	}
	result := &PageBreezemoonsResp{Rows: page.Rows}
	if len(page.Rows) > size {
		result.Rows = page.Rows[:size]
		result.NextCursor = u.encodeCursor(result.Rows[len(result.Rows)-1])
	}
	return result, nil
}

// BreezemoonCursor is the opaque cursor payload for stable dynamic pagination.
type BreezemoonCursor struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
}

func (*BreezemoonUsecase) decodeCursor(value *string) (*time.Time, *int64, error) {
	if value == nil || *value == "" {
		return nil, nil, nil
	}
	decoded, err := base64.RawURLEncoding.DecodeString(*value)
	if err != nil {
		return nil, nil, err
	}
	var cursor BreezemoonCursor
	if err = json.Unmarshal(decoded, &cursor); err != nil || cursor.ID <= 0 || cursor.CreatedAt.IsZero() {
		return nil, nil, errors.New("invalid breezemoon cursor")
	}
	return &cursor.CreatedAt, &cursor.ID, nil
}

func (*BreezemoonUsecase) encodeCursor(row *model.Breezemoon) *string {
	if row == nil || row.CreatedAt == nil {
		return nil
	}
	encoded, err := json.Marshal(BreezemoonCursor{
		CreatedAt: row.CreatedAt.UTC(),
		ID:        row.ID,
	})
	if err != nil {
		return nil
	}
	value := base64.RawURLEncoding.EncodeToString(encoded)
	return &value
}
