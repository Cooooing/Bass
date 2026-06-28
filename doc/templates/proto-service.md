# Proto Service 模板

## 文件组织

- 枚举放入 `enum.proto`。
- 稳定返回模型放入 `model.proto`。
- RPC request、reply 和写入结构放在当前资源 service proto。
- 一个 proto 文件最多定义一个 `service`。

## 内部服务 RPC

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

## BFF RPC

```proto
service ArticleService {
  rpc Publish(PublishArticle.Request) returns (PublishArticle.Reply) {
    option (google.api.http) = {
      post: "/v1/bbs/article/publish"
      body: "*"
    };
  }
}
```

## 禁止

- 内部服务 proto 写 HTTP 注解。
- BFF proto 引用内部服务 message。
- 写入结构抽成 `XXXSave`、`XXXUpdate`。
- 用 `with_xxx` 或 `include_xxx` 控制响应结构。
