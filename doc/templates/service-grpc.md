# gRPC Service 模板

内部 gRPC service 层只做协议适配，不承载领域流程。

```go
func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticle_Request) (*v1.PublishArticle_Reply, error) {
	if req.ArticleId == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.OperatorId == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}

	row, err := s.articleUsecase.Publish(ctx, req.ArticleId, req.OperatorId)
	if err != nil {
		return nil, err
	}

	return &v1.PublishArticle_Reply{
		Row: &v1.Article{
			Id:    row.ID,
			Title: row.Title,
		},
	}, nil
}
```

## 必须

- 从 request 显式读取业务输入。
- 在 service 层完成 proto 枚举到内部枚举的转换。
- 在当前接收者方法内显式组装 reply。

## 禁止

- 从 ctx、metadata、header、JWT 或缓存会话读取当前用户。
- 新增包级私有 helper。
- 新增 `toXXX`、`buildXXXReply`、`ConvertToRpc`。
- 把 proto 生成枚举传入 biz。