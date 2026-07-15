# HTTP Proto 模板

HTTP proto 用于 BFF 或 OpenAPI，对外暴露 HTTP/JSON 契约。

## 文件组织

- HTTP proto 定义自己的 request、reply 和对外模型。
- HTTP 路径按入口语义命名，不复用内部服务 proto message。
- OpenAPI 文档从 proto 注释和 `google.api.http` 生成。

## 示例

```proto
service ArticleService {
  rpc Publish(PublishArticle.Request) returns (PublishArticle.Reply) {
    option (google.api.http) = {
      post: "/v1/bbs/article/publish"
      body: "*"
    };
  }
}

message PublishArticle {
  message Request {
    int64 article_id = 1; // 文章 ID
  }

  message Reply {
    ArticleView row = 1; // 文章视图
  }
}

message ArticleView {
  int64 id = 1; // 文章 ID
  string title = 2; // 标题
}
```

## 必须

- 对外模型按视图语义定义，列表、详情、创建返回可以拆分不同模型。
- 嵌套对象使用最小可展示结构，不嵌完整内部详情模型。
- 普通业务接口默认使用 `POST` 和 `body: "*"`。
- 需要登录用户时，由 BFF 入口认证结果显式传入后续内部 RPC。

## 禁止

- 引用内部服务 message。
- 暴露内部服务原始 gRPC 路径、内部枚举命名或内部错误细节。
- 手写 OpenAPI tags、servers、external_docs、BearerAuth 或运行时说明。
- 暴露邮箱、手机号、账号名是否存在等可用于枚举用户信息的接口。