package repo

import (
	"context"
	"time"
)

// ArticleSearchRepo 通过已配置的搜索后端查询公开文章。
// 该端口使 Content 的公开搜索契约独立于 PostgreSQL、Elasticsearch 或未来的索引实现。
type ArticleSearchRepo interface {
	SearchPublic(ctx context.Context, req *SearchPublicArticlesReq) (*SearchPublicArticlesResp, error)
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

// ArticleSearchCursor 表示一次搜索查询的稳定排序元组。
type ArticleSearchCursor struct {
	Rank        int
	PublishedAt time.Time
	ArticleID   int64
}
