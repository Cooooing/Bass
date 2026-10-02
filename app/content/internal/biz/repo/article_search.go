package repo

import (
	"context"
	"time"
)

// ArticleSearchRepo queries public articles through the configured search backend.
// The port keeps Content's public search contract independent from PostgreSQL,
// Elasticsearch, or another future index implementation.
type ArticleSearchRepo interface {
	SearchPublic(context.Context, *SearchPublicArticlesReq) (*SearchPublicArticlesResp, error)
}

type SearchPublicArticlesReq struct {
	Query  string
	Cursor *ArticleSearchCursor
	Size   int
}

type SearchPublicArticlesResp struct {
	Rows    []*ArticleSearchRow
	HasMore bool
}

type ArticleSearchRow struct {
	ArticleID   int64
	Rank        int
	PublishedAt time.Time
}

// ArticleSearchCursor represents the stable sort tuple for one search query.
type ArticleSearchCursor struct {
	Rank        int
	PublishedAt time.Time
	ArticleID   int64
}
