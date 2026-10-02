package repo

import (
	"context"
	"strings"
	"time"

	utilent "common/pkg/util/ent"
	bizrepo "content/internal/biz/repo"
	"content/internal/data/gen"
	articleent "content/internal/data/gen/article"
	"content/internal/enum"

	"entgo.io/ent/dialect/sql"
)

var _ bizrepo.ArticleSearchRepo = (*PostgresArticleSearchRepo)(nil)

// PostgresArticleSearchRepo implements literal public article search with
// PostgreSQL ILIKE predicates. The deployed pg_trgm GIN indexes accelerate
// these predicates without changing the domain-facing search contract.
type PostgresArticleSearchRepo struct {
	db *gen.Client
}

func NewPostgresArticleSearchRepo(db *gen.Client) bizrepo.ArticleSearchRepo {
	return &PostgresArticleSearchRepo{db: db}
}

func (r *PostgresArticleSearchRepo) SearchPublic(
	ctx context.Context,
	req *bizrepo.SearchPublicArticlesReq,
) (*bizrepo.SearchPublicArticlesResp, error) {
	query := r.client(ctx).Article.Query().Where(
		articleent.DeletedAtIsNil(),
		articleent.PublishStatusEQ(articleent.PublishStatus(enum.ArticlePublishStatusPublished)),
		articleent.VisibilityEQ(articleent.Visibility(enum.ArticleVisibilityPublic)),
		articleent.RestrictionIn(
			articleent.Restriction(enum.ContentRestrictionNone),
			articleent.Restriction(enum.ContentRestrictionLocked),
		),
	)

	pattern := "%" + r.escapeLike(req.Query) + "%"
	query = query.Where(r.matchesLiteral(pattern))
	if req.Cursor != nil {
		query = query.Where(r.afterCursor(req.Query, pattern, req.Cursor))
	}
	query.Modify(r.orderByRelevance(req.Query, pattern))

	limit := req.Size + 1
	rows, err := query.Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}
	result := &bizrepo.SearchPublicArticlesResp{HasMore: len(rows) > req.Size}
	if result.HasMore {
		rows = rows[:req.Size]
	}
	result.Rows = make([]*bizrepo.ArticleSearchRow, 0, len(rows))
	for _, row := range rows {
		var publishedAt time.Time
		if row.PublishedAt != nil {
			publishedAt = *row.PublishedAt
		} else if row.CreatedAt != nil {
			publishedAt = *row.CreatedAt
		}
		result.Rows = append(result.Rows, &bizrepo.ArticleSearchRow{
			ArticleID:   row.ID,
			Rank:        r.rank(row.Title, req.Query),
			PublishedAt: publishedAt,
		})
	}
	return result, nil
}

func (r *PostgresArticleSearchRepo) client(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (*PostgresArticleSearchRepo) escapeLike(value string) string {
	value = strings.ReplaceAll(value, `\`, `\\`)
	value = strings.ReplaceAll(value, "%", `\%`)
	return strings.ReplaceAll(value, "_", `\_`)
}

func (*PostgresArticleSearchRepo) matchesLiteral(pattern string) func(*sql.Selector) {
	return func(selector *sql.Selector) {
		selector.Where(sql.P(func(builder *sql.Builder) {
			builder.Nested(func(builder *sql.Builder) {
				builder.Ident(articleent.FieldTitle).
					WriteString(" ILIKE ").
					Arg(pattern).
					WriteString(" ESCAPE '\\'")
				builder.WriteString(" OR ")
				builder.Ident(articleent.FieldContent).
					WriteString(" ILIKE ").
					Arg(pattern).
					WriteString(" ESCAPE '\\'")
			})
		}))
	}
}

func (r *PostgresArticleSearchRepo) afterCursor(
	query string,
	pattern string,
	cursor *bizrepo.ArticleSearchCursor,
) func(*sql.Selector) {
	return func(selector *sql.Selector) {
		selector.Where(sql.P(func(builder *sql.Builder) {
			builder.WriteByte('(')
			r.writeRank(builder, query, pattern)
			builder.WriteString(" > ").Arg(cursor.Rank)
			builder.WriteString(" OR (")
			r.writeRank(builder, query, pattern)
			builder.WriteString(" = ").Arg(cursor.Rank)
			builder.WriteString(" AND (")
			builder.Ident(articleent.FieldPublishedAt).WriteString(" < ").Arg(cursor.PublishedAt)
			builder.WriteString(" OR (")
			builder.Ident(articleent.FieldPublishedAt).WriteString(" = ").Arg(cursor.PublishedAt)
			builder.WriteString(" AND ").Ident(articleent.FieldID).WriteString(" < ").Arg(cursor.ArticleID)
			builder.WriteString(")")
			builder.WriteString("))")
			builder.WriteString(")")
		}))
	}
}

func (r *PostgresArticleSearchRepo) orderByRelevance(query string, pattern string) func(*sql.Selector) {
	return func(selector *sql.Selector) {
		selector.OrderExpr(sql.ExprFunc(func(builder *sql.Builder) {
			r.writeRank(builder, query, pattern)
			builder.WriteString(" ASC")
			builder.WriteString(", ").Ident(articleent.FieldPublishedAt).WriteString(" DESC")
			builder.WriteString(", ").Ident(articleent.FieldID).WriteString(" DESC")
		}))
	}
}

func (*PostgresArticleSearchRepo) writeRank(builder *sql.Builder, query string, pattern string) {
	builder.WriteString("CASE WHEN ")
	builder.Ident(articleent.FieldTitle).WriteString(" = ").Arg(query)
	builder.WriteString(" THEN 0 WHEN ")
	builder.Ident(articleent.FieldTitle).WriteString(" ILIKE ").Arg(pattern).WriteString(" ESCAPE '\\'")
	builder.WriteString(" THEN 1 ELSE 2 END")
}

func (*PostgresArticleSearchRepo) rank(title string, query string) int {
	if title == query {
		return 0
	}
	if strings.Contains(strings.ToLower(title), strings.ToLower(query)) {
		return 1
	}
	return 2
}
