# Usecase 模板

usecase 负责业务流程、状态流转、本地业务不变量和事务边界。

```go
type PublishArticleReq struct {
	ArticleID  int64
	OperatorID int64
}

func (u *ArticleUsecase) Publish(ctx context.Context, req *PublishArticleReq) (*model.Article, error) {
	var row *model.Article
	err := u.tx(ctx, func(ctx context.Context) error {
		exist, err := u.articleRepo.Get(ctx, req.ArticleID)
		if err != nil {
			return err
		}
		if exist.CreatedBy == nil || *exist.CreatedBy != req.OperatorID {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		row, err = u.articleRepo.Publish(ctx, &repo.ArticlePublishReq{
			ArticleID:  req.ArticleID,
			OperatorID: req.OperatorID,
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return row, nil
}
```

## 必须

- 无业务参数时不定义 `XxxReq`；一个业务参数时直接传参；两个及以上业务参数时定义 `*XxxReq`。
- 无业务返回值时只返回 `error`；一个业务返回值时直接返回该值和 `error`；两个及以上业务返回值时定义 `*XxxResp`。
- 需要 `XxxReq` 或 `XxxResp` 时，定义在对应方法正上方。
- 校验资源存在性、状态流转、本地业务不变量和写入前置条件。
- 事务只包住必须原子提交的写入流程。
- 事件 outbox 在归属服务本地事务内写入。

## 禁止

- import proto 生成包、data 包或 Ent 生成包。
- 从 ctx 读取当前用户或业务参数。
- 把 usecase 请求或响应结构体集中放到 `query.go` 或 `dto.go`。
- 把一次性简单逻辑拆成游离私有函数。
