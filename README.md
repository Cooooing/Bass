# Bass

Bass 是面向社区产品的 Go 微服务项目。系统按入口层、内部业务服务层、实时推送层划分职责，内部服务通过 gRPC 和事件协作。

## 快速了解

- 入口层负责 HTTP / SSE 接入、端侧接口编排和展示模型适配。
- 内部服务拥有各自领域数据和业务规则，不共享业务数据库。
- 跨服务副作用通过 outbox 和 NATS 异步传播。
- `common` 存放公共 proto、客户端封装、枚举、错误和基础工具。

## 技术栈

- Go 1.26、Kratos、Ent、Wire
- Protobuf、gRPC、HTTP/JSON、OpenAPI、Buf
- PostgreSQL、Redis、NATS、Consul

## 项目结构

```text
Bass/
├─ app/        # 服务模块
├─ common/     # 公共 proto、公共 Make 片段、基础封装
├─ deploy/     # 部署配置
├─ doc/        # 架构、编码规范和模板
└─ Makefile    # 根构建入口
```

单个服务通常包含：

```text
app/<service>/
├─ cmd/             # 启动与 Wire 注入
├─ configs/         # 配置示例
├─ internal/
│  ├─ service/      # 协议适配层
│  ├─ biz/          # 业务规则与事务编排
│  ├─ data/         # 数据访问与 Ent schema
│  ├─ server/       # HTTP/gRPC/消费者注册
│  └─ config/       # 配置 proto 与读取封装
└─ Makefile
```

## 常用命令

```bash
make init
make api
make all
make build-all
make api-lint
```

单模块命令示例：

```bash
make -C app/user gen
make -C app/content build
make -C app/bbs doc
make -C app/bbs sdk
```

## 文档入口

| 文档 | 内容 |
|------|------|
| [doc/architecture.md](doc/architecture.md) | 架构边界、读写归属、契约、事件、缓存、运行时规则和 `common` 封装约定。 |
| [doc/coding.md](doc/coding.md) | 项目级编码规范和代码组织结构。 |
| [doc/templates/](doc/templates/) | 新增 proto、service、usecase、repo、schema 时参考的模板。 |

## 模板索引

| 场景 | 模板 |
|------|------|
| HTTP proto | [doc/templates/proto-http.md](doc/templates/proto-http.md) |
| gRPC proto | [doc/templates/proto-grpc.md](doc/templates/proto-grpc.md) |
| HTTP service | [doc/templates/service-http.md](doc/templates/service-http.md) |
| gRPC service | [doc/templates/service-grpc.md](doc/templates/service-grpc.md) |
| usecase | [doc/templates/usecase.md](doc/templates/usecase.md) |
| repo | [doc/templates/repo.md](doc/templates/repo.md) |
| Ent schema | [doc/templates/ent-schema.md](doc/templates/ent-schema.md) |
