# Usecase 模板

usecase 负责业务流程、状态流转、本地业务不变量和事务边界。

```go
type PublishArticleReq struct {
	ArticleID  int64
	OperatorID int64
}

type PublishArticleResponse struct {
	Article *model.Article
}

func (u *ArticleUsecase) Publish(ctx context.Context, req *PublishArticleReq) (*PublishArticleResponse, error) {
	var row *model.Article
	err := u.tx(ctx, func(ctx context.Context) error {
		exist, err := u.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleID: new(req.ArticleID)})
		if err != nil {
			return err
		}
		if exist.CreatedBy == nil || *exist.CreatedBy != req.OperatorID {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		rowResp, err := u.articleRepo.Publish(ctx, &repo.ArticlePublishReq{
			ArticleID:  req.ArticleID,
			OperatorID: req.OperatorID,
		})
		if err != nil {
			return err
		}
		row = rowResp.Article
		return err
	})
	if err != nil {
		return nil, err
	}
	return &PublishArticleResponse{Article: row}, nil
}
```

## 必须

- usecase receiver 方法除 `ctx` 外只接收一个 `*XxxReq`。
- 有业务返回值时返回 `*XxxResponse, error`。
- `XxxReq` 和 `XxxResponse` 定义在对应方法正上方。
- 校验资源存在性、状态流转、本地业务不变量和写入前置条件。
- 事务只包住必须原子提交的写入流程。
- 事件 outbox 在归属服务本地事务内写入。

## 禁止

- import proto 生成包、data 包或 Ent 生成包。
- 从 ctx 读取当前用户或业务参数。
- 把 usecase 请求或响应结构体集中放到 `query.go` 或 `dto.go`。
- 把一次性简单逻辑拆成游离私有函数。
