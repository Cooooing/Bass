# Repo 模板

每个 schema 对应的 repo 保持五个基础查询方法，并按业务动作补充写方法。

## biz repo 接口

```go
type ArticleRepo interface {
	Get(ctx context.Context, req *ArticleGetReq) (*ArticleGetResponse, error)
	List(ctx context.Context, req *ArticleGetReq) (*ArticleListResponse, error)
	Map(ctx context.Context, req *ArticleGetReq) (*ArticleMapResponse, error)
	Count(ctx context.Context, req *ArticleGetReq) (*ArticleCountResponse, error)
	Page(ctx context.Context, req *ArticlePageReq) (*ArticlePageResponse, error)

	Publish(ctx context.Context, req *ArticlePublishReq) (*ArticlePublishResponse, error)
}

type ArticleGetReq struct {
	ArticleID *int64
}

type ArticleGetResponse struct {
	Article *model.Article
}

type ArticleListResponse struct {
	Rows []*model.Article
}

type ArticleMapResponse struct {
	Rows map[int64]*model.Article
}

type ArticleCountResponse struct {
	Count int
}

type ArticlePageReq struct {
	Page  base.PageRequest
	Query ArticleGetReq
}

type ArticlePageResponse struct {
	Rows []*model.Article
	Page base.PageResponse
}

type ArticlePublishReq struct {
	ArticleID  int64
	OperatorID int64
}

type ArticlePublishResponse struct {
	Article *model.Article
}
```

## data repo 实现

```go
func (r *ArticleRepo) Get(ctx context.Context, req *repo.ArticleGetReq) (*repo.ArticleGetResponse, error) {
	query := r.getClient(ctx).Article.Query()
	if req.ArticleID != nil {
		query.Where(article.ID(*req.ArticleID))
	}
	row, err := query.Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &repo.ArticleGetResponse{
		Article: &model.Article{
			ID:    row.ID,
			Title: row.Title,
		},
	}, nil
}
```

## 必须

- biz/repo 参数对象只表达持久化查询、写入条件或外部能力调用条件。
- biz/repo 接口方法除 `ctx` 外只接收一个 `*XxxReq`，返回 `*XxxResponse, error`。
- data 层显式组装 Ent entity 和 biz model。
- 返回 map 的方法使用 `Map` 语义命名。

## 禁止

- repo 写方法同时处理多个不相关资源。
- biz 接口暴露 proto request/response、Ent predicate、Ent entity 或数据库类型。
- 用 `Manager`、`Resolver`、`Writer` 包装真实依赖。
