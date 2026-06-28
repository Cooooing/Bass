# Repo 模板

每个 schema 对应的 repo 保持五个基础查询方法，并按业务动作补充写方法。

## biz repo 接口

```go
type ArticleRepo interface {
	Get(ctx context.Context, req *ArticleGetReq) (*model.Article, error)
	List(ctx context.Context, req *ArticleGetReq) ([]*model.Article, error)
	Map(ctx context.Context, req *ArticleGetReq) (map[int64]*model.Article, error)
	Count(ctx context.Context, req *ArticleGetReq) (int, error)
	Page(ctx context.Context, req *ArticlePageReq) ([]*model.Article, *base.Page, error)

	Publish(ctx context.Context, articleID int64, operatorID int64) (*model.Article, error)
}
```

## data repo 实现

```go
func (r *ArticleRepo) Get(ctx context.Context, req *repo.ArticleGetReq) (*model.Article, error) {
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
	return &model.Article{
		ID:    row.ID,
		Title: row.Title,
	}, nil
}
```

## 必须

- biz/repo 参数对象只表达持久化查询或写入条件。
- data 层显式组装 Ent entity 和 biz model。
- 返回 map 的方法使用 `Map` 语义命名。

## 禁止

- repo 写方法同时处理多个不相关资源。
- biz 接口暴露 Ent predicate、Ent entity 或数据库类型。
- 用 `Manager`、`Resolver`、`Writer` 包装真实依赖。
