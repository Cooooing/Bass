# Usecase 模板

usecase 负责业务流程、状态流转、本地业务不变量和事务边界。

```go
func (u *ArticleUsecase) Publish(ctx context.Context, articleID int64, operatorID int64) (*model.Article, error) {
	var row *model.Article
	err := u.tx(ctx, func(ctx context.Context) error {
		exist, err := u.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleID: new(articleID)})
		if err != nil {
			return err
		}
		if exist.CreatedBy == nil || *exist.CreatedBy != operatorID {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		row, err = u.articleRepo.Publish(ctx, articleID, operatorID)
		return err
	})
	return row, err
}
```

## 必须

- 校验资源存在性、状态流转、本地业务不变量和写入前置条件。
- 事务只包住必须原子提交的写入流程。
- 事件 outbox 在归属服务本地事务内写入。

## 禁止

- import proto 生成包、data 包或 Ent 生成包。
- 从 ctx 读取当前用户或业务参数。
- 把一次性简单逻辑拆成游离私有函数。
