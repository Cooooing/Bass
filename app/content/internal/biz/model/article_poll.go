package model

import (
	"strings"
	"time"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
)

// ArticlePollDraft 是文章草稿期可编辑的投票配置。
type ArticlePollDraft struct {
	Options    []string
	EndsAt     time.Time
	MaxChoices int32
}

// ArticlePoll 是文章详情中的投票事实与当前读者视图。
type ArticlePoll struct {
	ID                int64
	EndsAt            time.Time
	MaxChoices        int32
	ParticipantCount  int64
	TotalVoteCount    int64
	Options           []*ArticlePollOption
	SelectedOptionIDs []int64
	ResultVisible     bool
}

// ArticlePollOption 是投票的一个选项。
type ArticlePollOption struct {
	ID        int64
	Content   string
	VoteCount int32
}

func (d *ArticlePollDraft) Normalize(now time.Time) error {
	if d == nil || len(d.Options) < 2 || len(d.Options) > 20 || d.MaxChoices < 1 || int(d.MaxChoices) > len(d.Options) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if !d.EndsAt.After(now.Add(5*time.Minute)) || d.EndsAt.After(now.AddDate(0, 3, 0)) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_SCHEDULED_AT_OUT_OF_RANGE)
	}
	options := make([]string, 0, len(d.Options))
	seen := make(map[string]struct{}, len(d.Options))
	for _, option := range d.Options {
		option = strings.TrimSpace(option)
		if option == "" {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		if _, ok := seen[option]; ok {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		seen[option] = struct{}{}
		options = append(options, option)
	}
	d.Options = options
	d.EndsAt = d.EndsAt.UTC().Truncate(time.Second)
	return nil
}

func (p *ArticlePoll) CanRevealResults(authorID, viewerID int64, now time.Time) bool {
	return viewerID == authorID || len(p.SelectedOptionIDs) > 0 || !now.Before(p.EndsAt)
}

func (p *ArticlePoll) ValidateVote(optionIDs []int64, now time.Time) error {
	if p == nil || !now.Before(p.EndsAt) || len(optionIDs) == 0 || len(optionIDs) > int(p.MaxChoices) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	options := make(map[int64]struct{}, len(p.Options))
	for _, option := range p.Options {
		options[option.ID] = struct{}{}
	}
	seen := make(map[int64]struct{}, len(optionIDs))
	for _, optionID := range optionIDs {
		if _, exists := options[optionID]; !exists {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		if _, duplicate := seen[optionID]; duplicate {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		seen[optionID] = struct{}{}
	}
	return nil
}
