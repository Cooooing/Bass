# HTTP Service 模板

HTTP service 用于 BFF 或 OpenAPI，负责入口协议适配、端侧校验和内部 RPC 编排。

```go
func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticle_Req) (*v1.PublishArticle_Resp, error) {
	if req.ArticleId == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := s.contentClient.Publish(ctx, &contentv1.PublishArticle_Req{
		ArticleId:  req.ArticleId,
		OperatorId: userID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.PublishArticle_Resp{
		Row: &v1.PublishArticle_Resp_ArticleView{
			Id:    resp.GetRow().GetId(),
			Title: resp.GetRow().GetTitle(),
		},
	}, nil
}
```

## 必须

- 只做入口认证、端侧参数校验、内部 RPC 调用和对外模型组装。
- 写接口通常只调用一个归属服务命令接口。
- 当前用户、操作人或审计人必须作为内部 RPC request 显式字段传入。
- 读接口可以聚合多个内部服务，并按端侧语义控制可见范围。

## 禁止

- 连接业务库或维护领域数据。
- 在 BFF 做 read-before-write 的领域状态预校验。
- 直接发布领域 MQ 事件。
- 返回内部服务原始 message。
