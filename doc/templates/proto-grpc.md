# gRPC Proto 模板

内部服务 proto 只定义 gRPC 契约，不写 HTTP 注解。

## 文件组织

- 枚举放入 `enum.proto`。
- 业务模块不定义 `model.proto`，不把领域模型放到独立公共 message 中复用。
- 每个 RPC 使用一个顶层 message，并在其中定义 `Request` 和 `Response`。
- `Request` / `Response` 内需要封装结构时，在当前 RPC message 内定义子 message。
- 一个 proto 文件最多定义一个 `service`。

## 示例

```proto
service ArticleService {
  rpc Publish(PublishArticle.Request) returns (PublishArticle.Response) {}
  rpc Get(GetArticle.Request) returns (GetArticle.Response) {}
  rpc Page(PageArticles.Request) returns (PageArticles.Response) {}
}

message PublishArticle {
  message Request {
    int64 article_id = 1; // 文章 ID
    int64 operator_id = 2; // 操作人 ID
  }

  message Response {
    Article row = 1; // 文章

    message Article {
      int64 id = 1; // 文章 ID
      string title = 2; // 标题
      ArticlePublishStatus publish_status = 3; // 发布状态
    }
  }
}

message PageArticles {
  message Request {
    common.PageRequest page = 1; // 分页参数
    Query query = 2; // 查询条件

    message Query {
      optional int64 author_id = 1; // 作者 ID
      optional ArticlePublishStatus publish_status = 2; // 发布状态
    }
  }

  message Response {
    common.PageResponse page = 1; // 分页结果
    repeated Row rows = 2; // 文章列表

    message Row {
      int64 id = 1; // 文章 ID
      string title = 2; // 标题
      ArticlePublishStatus publish_status = 3; // 发布状态
    }
  }
}
```

## 必须

- 业务输入来自 `Request` 显式字段。
- 写接口需要操作人、审计人或目标资源时，字段必须出现在 `Request` 中。
- `List`、`Page` 响应数组字段使用 `rows`，分页信息字段使用 `page`。
- 响应中需要业务结构时，在当前 `Response` 内定义子 message。
- 内部服务返回归属领域事实和 ID，不返回其他服务展示模型。

## 禁止

- 写 `google.api.http` 注解。
- 定义业务模块 `model.proto`。
- 在接口之间复用业务 message 作为模型。
- 从 ctx、metadata、header、JWT 或会话隐式读取业务参数。
- 用 `with_xxx` 或 `include_xxx` 控制响应结构。
- 把写入结构抽成 `XXXSave`、`XXXUpdate` 这类通用消息。
