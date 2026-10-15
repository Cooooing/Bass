package model

import "time"

type ArticleViewerActionState struct {
	Liked     bool
	Thanked   bool
	Collected bool
	Rewarded  bool
}

type ArticlePollDraft struct {
	Options    []string
	EndsAt     time.Time
	MaxChoices int32
}

type ArticlePoll struct {
	EndsAt            time.Time
	MaxChoices        int32
	ParticipantCount  int64
	TotalVoteCount    int64
	Options           []*ArticlePollOption
	SelectedOptionIDs []int64
	ResultVisible     bool
}

type ArticlePollOption struct {
	ID        int64
	Content   string
	VoteCount int32
}

type ArticleLotteryDraft struct {
	Prizes []*ArticleLotteryPrizeDraft
	DrawAt time.Time
}

type ArticleLotteryPrizeDraft struct {
	Name        string
	Description *string
	WinnerLimit int32
}

type ArticleLottery struct {
	DrawAt     time.Time
	DrawnAt    *time.Time
	EntryCount int64
	Entered    bool
	Prizes     []*ArticleLotteryPrize
}

type ArticleLotteryPrize struct {
	ID            int64
	Name          string
	Description   *string
	WinnerLimit   int32
	WinnerUserIDs []int64
	WinnerUsers   []*AccountProfile
}

type ArticleTag struct {
	ID       int64
	Code     string
	Name     string
	Icon     *string
	DomainID *int64
}

type ArticleDomain struct {
	ID   int64
	Code string
	Name string
	Icon *string
	URL  *string
	Sort int32
}

type ArticlePostscript struct {
	ID            int64
	ArticleID     int64
	Content       string
	ContentRender string
	Restriction   int32
	CreatedBy     *int64
	UpdatedBy     *int64
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

type ArticleListItem struct {
	ID                int64
	Title             string
	Content           string
	ContentRender     string
	HasPostscript     bool
	HasReward         bool
	PublishStatus     int32
	Visibility        int32
	Restriction       int32
	Type              int32
	Statement         *string
	Commentable       bool
	ViewCount         int32
	ThankCount        int32
	LikeCount         int32
	CollectCount      int32
	RewardCount       int32
	ReplyCount        int32
	CoverImageURL     *string
	ViewerActionState *ArticleViewerActionState
	LastReplyAt       *time.Time
	LastReplyUser     *AccountProfile
	AuthorUser        *AccountProfile
	CreatedBy         *int64
	UpdatedBy         *int64
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	PublishedAt       *time.Time
	EditedAt          *time.Time
	Tags              []*ArticleTag
	Domains           []*ArticleDomain
}

type ArticleDetail struct {
	ID                  int64
	Title               string
	Content             string
	ContentRender       string
	HasPostscript       bool
	HasReward           bool
	RewardContent       *string
	RewardContentRender *string
	RewardPoints        *int32
	PublishStatus       int32
	Visibility          int32
	Restriction         int32
	Type                int32
	Statement           *string
	Commentable         bool
	ViewCount           int32
	ThankCount          int32
	LikeCount           int32
	CollectCount        int32
	RewardCount         int32
	ReplyCount          int32
	CoverImageURL       *string
	ViewerActionState   *ArticleViewerActionState
	LastReplyAt         *time.Time
	LastReplyUser       *AccountProfile
	Postscripts         []*ArticlePostscript
	AuthorUser          *AccountProfile
	CreatedBy           *int64
	UpdatedBy           *int64
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	PublishedAt         *time.Time
	City                *string
	EditedAt            *time.Time
	Tags                []*ArticleTag
	Domains             []*ArticleDomain
	PinnedID            *int64
	AcceptedCommentID   *int64
	Poll                *ArticlePoll
	Lottery             *ArticleLottery
}
