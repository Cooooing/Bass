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
├─ doc/        # 本地业务事实、架构规则、模板和 Agent 工作流（Git 忽略）
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
make gen
make build
```

单模块命令示例：

```bash
make -C app/user gen
make -C app/user gen-clean
make -C app/content build
make -C app/bff_bbs doc
make -C app/bff_bbs sdk
make -C monolith gen
make -C monolith build
```

## 本地文档

本机的 [doc/README.md](doc/README.md) 是业务事实、架构规则、工程模板和 Agent 工作流入口。`doc/` 不进入 Git；开发新功能、修复 bug 或改变业务事实时，必须同步维护对应功能页。
