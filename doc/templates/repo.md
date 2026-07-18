# Repo 模板

每个 schema 对应的 repo 保持五个基础查询方法，并按业务动作补充写方法。

## biz repo 接口

```go
type ArticleRepo interface {
	Get(ctx context.Context, articleID int64) (*model.Article, error)
	List(ctx context.Context, req *ArticleListReq) ([]*model.Article, error)
	Map(ctx context.Context, req *ArticleListReq) (map[int64]*model.Article, error)
	Count(ctx context.Context, req *ArticleListReq) (int, error)
	Page(ctx context.Context, req *ArticlePageReq) (*ArticlePageResp, error)

	Publish(ctx context.Context, req *ArticlePublishReq) (*model.Article, error)
}

type ArticleListReq struct {
	AuthorID *int64
	Status   *enum.ArticleStatus
}

type ArticlePageReq struct {
	Page  base.PageRequest
	Query ArticleListReq
}

type ArticlePageResp struct {
	Rows []*model.Article
	Page base.PageResp
}

type ArticlePublishReq struct {
	ArticleID  int64
	OperatorID int64
}
```

## data repo 实现

```go
func (r *ArticleRepo) Get(ctx context.Context, articleID int64) (*model.Article, error) {
	row, err := r.getClient(ctx).Article.Query().Where(article.ID(articleID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.Article{
		ID:    row.ID,
		Title: row.Title,
	}, nil
}
```

## 必须

- biz/repo 参数对象只表达持久化查询、写入条件或外部能力调用条件。
- 无业务参数时不定义 `XxxReq`；一个业务参数时直接传参；两个及以上业务参数时定义 `*XxxReq`。无业务返回值时只返回 `error`；一个业务返回值时直接返回该值和 `error`；两个及以上业务返回值时定义 `*XxxResp`。
- data 层显式组装 Ent entity 和 biz model。
- 返回 map 的方法使用 `Map` 语义命名。

## 禁止

- repo 写方法同时处理多个不相关资源。
- biz 接口暴露 proto request/resp、Ent predicate、Ent entity 或数据库类型。
- 用 `Manager`、`Resolver`、`Writer` 包装真实依赖。
