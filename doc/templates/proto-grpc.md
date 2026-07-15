# gRPC Proto 模板

内部服务 proto 只定义 gRPC 契约，不写 HTTP 注解。

## 文件组织

- 枚举放入 `enum.proto`。
- 稳定领域模型放入 `model.proto`。
- RPC request、reply 和命令结构放在资源 service proto。
- 一个 proto 文件最多定义一个 `service`。

## 示例

```proto
service ArticleService {
  rpc Publish(PublishArticle.Request) returns (PublishArticle.Reply) {}
  rpc Get(GetArticle.Request) returns (GetArticle.Reply) {}
  rpc List(ListArticles.Request) returns (ListArticles.Reply) {}
  rpc Page(PageArticles.Request) returns (PageArticles.Reply) {}
  rpc Map(MapArticles.Request) returns (MapArticles.Reply) {}
}

message PublishArticle {
  message Request {
    int64 article_id = 1; // 文章 ID
    int64 operator_id = 2; // 操作人 ID
  }

  message Reply {
    Article row = 1; // 文章
  }
}
```

## 必须

- 业务输入来自 request 显式字段。
- 写接口需要操作人、审计人或目标资源时，字段必须出现在 request 中。
- `List`、`Page` 响应数组字段使用 `rows`，分页信息字段使用 `page`。
- 内部服务返回归属领域事实和 ID，不返回其他服务展示模型。

## 禁止

- 写 `google.api.http` 注解。
- 从 ctx、metadata、header、JWT 或会话隐式读取业务参数。
- 用 `with_xxx` 或 `include_xxx` 控制响应结构。
- 把写入结构抽成 `XXXSave`、`XXXUpdate` 这类通用消息。