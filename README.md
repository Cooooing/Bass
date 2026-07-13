# Bass

Bass 是面向社区产品的 Go 微服务项目。系统按入口层、内部业务服务层、实时推送层划分职责：入口层负责协议接入和端侧接口编排，内部服务拥有各自领域数据与业务规则，跨服务副作用通过 outbox 和消息总线异步传播。

## 目录

- [技术栈](#技术栈)
- [服务边界](#服务边界)
- [调用原则](#调用原则)
- [项目结构](#项目结构)
- [构建与生成](#构建与生成)
- [文档入口](#文档入口)

## 技术栈

- 语言与框架：Go 1.26、Kratos、Ent、Wire
- API 与契约：Protobuf、gRPC、HTTP/JSON、OpenAPI、Buf
- 数据与缓存：PostgreSQL、Redis
- 消息总线：NATS
- 配置与发现：Consul

## 服务边界

### 外部入口层

| 服务名 | 类型 | 主要消费者 | 主要职责 | 对外协议 | 对内依赖 | 不负责 |
| --- | --- | --- | --- | --- | --- | --- |
| `bbs` | BFF | 社区前台 Web / App | 面向普通用户聚合社区前台接口，适配用户、内容、通知、IM 等能力 | HTTP/JSON | `user`、`content`、`notify`、`im`、`platform` | 不承载核心领域规则；不直接维护消息路由或长连接状态 |
| `bbs_admin` | BFF | 管理后台 | 面向运营和管理员聚合后台接口，提供管理视角下的用户、内容、通知、平台能力编排 | HTTP/JSON | `user`、`content`、`notify`、`platform`，必要时依赖 `im` | 不承载底层领域逻辑；不直接操作实时连接 |
| `openapi` | External Edge | 外部开发者、合作方 | 提供稳定、版本化、限流和签名保护的开放接口，对外暴露受控业务能力 | HTTP/JSON | `user`、`content`、`notify`、`platform` | 不承担内部管理接口；不直接暴露内部服务原始协议 |
| `push_node` | Push Edge / SSE Node | 浏览器 / App 客户端 | 维护 SSE 长连接，向客户端下行推送聊天消息、通知和系统事件，可独立部署 | HTTP/SSE | `push_hub`、NATS | 不接收用户上行写请求；不负责消息持久化、鉴权编排、业务审核 |

### 内部业务服务层

| 服务名 | 业务域 | 核心职责 | 拥有的数据 | 对外提供 | 主要上游 | 不负责 |
| --- | --- | --- | --- | --- | --- | --- |
| `user` | 身份与账户 | 用户注册登录、认证鉴权、用户资料、关注/拉黑、2FA、账号设置 | 用户、认证凭证、关系、账户设置 | gRPC | `bbs`、`bbs_admin`、`openapi`、`im`、`notify` | 不负责内容；不负责通知投递分发 |
| `content` | 社区内容 | 文章、评论、标签、板块、内容互动和内容查询 | 文章、评论、标签、板块、互动记录 | gRPC | `bbs`、`bbs_admin`、`openapi` | 不负责账户体系；不负责实时推送链路 |
| `notify` | 通知 | 站内通知、系统通知、邮件/短信通知编排、模板、通知偏好、投递记录 | 通知模板、通知记录、通知偏好、投递状态 | gRPC + Event | `bbs`、`bbs_admin`、`platform`、业务事件源 | 不负责实时连接管理；不负责聊天会话核心逻辑 |
| `im` | 即时通信 | 消息发送、会话管理、消息持久化、已读回执、历史消息、消息审核编排、生成可投递事件 | 会话、消息、已读状态、成员关系、投递状态 | gRPC + Event | `bbs`，必要时依赖 `openapi` | 不直接维护 SSE 连接；不直接管理外部推送节点 |
| `platform` | 平台能力 | 上传凭证、回调验签、文件/对象存储接入、图片审核适配、第三方能力适配 | 文件元数据、回调记录、平台接入配置、能力调用记录 | gRPC | `bbs`、`bbs_admin`、`openapi`、`notify` | 不承载具体社区业务决策；不负责消息路由 |
| `push_hub` | 实时分发控制面 | 管理 `push_node` 节点，维护用户/会话到节点的路由关系，消费 IM/通知事件并分发到对应节点，管理在线状态和推送控制策略 | 节点注册信息、在线路由、会话-节点映射、推送分发表 | gRPC + Event | `im`、`notify`、NATS | 不负责消息业务持久化；不直接对用户提供业务接口 |
| `scheduler` | 定时任务 | 后台任务定义、分布式调度触发、运行记录、超时和执行告警 | 任务定义、任务运行记录 | gRPC | 后台管理入口、各业务服务 | 不连接业务库；不承载被调用服务的领域规则；不自动重试失败或超时任务 |

### 职责归属

| 主题 | 归属服务 |
| --- | --- |
| 前台页面/API 聚合 | `bbs` |
| 后台管理/API 聚合 | `bbs_admin` |
| 第三方回调接入 | `platform` |
| 对外开放 API | `openapi` |
| 客户端下行实时推送 | `push_node` |
| 用户与认证 | `user` |
| 社区内容与互动 | `content` |
| 通知业务 | `notify` |
| 即时通信业务 | `im` |
| 通用平台能力 | `platform` |
| 实时推送路由与节点控制 | `push_hub` |
| 分布式定时任务 | `scheduler` |

当前代码目录与目标服务名的对应关系见 [doc/README.md](doc/README.md)。

## 调用原则

| 场景 | 推荐方式 |
| --- | --- |
| 外部客户端调用 BFF / OpenAPI / Platform | HTTP |
| 客户端接收实时消息 | SSE |
| 内部同步调用 | gRPC |
| 内部异步事件 | NATS |

关键边界约束：

- `push_node` 只负责下行推送，不接收用户上行写请求。
- `im` 只负责即时通信业务，不负责边缘连接管理。
- `push_hub` 是实时分发控制面，不拥有聊天或通知本身的业务数据。
- `platform` 提供通用能力，不做社区业务裁决。
- `scheduler` 只负责任务定义、触发、运行记录和告警，不连接其他服务业务库，不替被调用服务实现领域逻辑。
- BFF 只做聚合与适配，不下沉核心业务规则。
- BFF 写接口通常只调用一个归属服务命令接口；跨服务副作用由归属服务写 outbox 后通过 NATS 触发。

## 项目结构

```bash
Bass/
├─ app/                 # 服务模块
│  └─ <service>/         # 单个服务
│     ├─ cmd/            # 启动与 Wire 注入
│     ├─ configs/        # 配置示例
│     ├─ internal/       # 服务内部代码
│     │  ├─ service/     # 协议适配层
│     │  ├─ biz/         # 业务规则与事务编排
│     │  ├─ data/        # 数据访问与 Ent schema
│     │  ├─ server/      # HTTP/gRPC/消费者注册
│     │  └─ conf/        # 配置 proto
│     └─ Makefile
├─ common/
│  ├─ api/               # 共享 proto、生成产物、SDK 配置
│  ├─ buf/               # Buf 生成模板
│  ├─ make/              # 公共 Make 片段
│  ├─ ops/               # 运维配置
│  └─ pkg/               # 通用 Go 包
├─ deploy/               # 部署配置
├─ doc/                  # 项目规范文档
└─ Makefile              # 根构建入口
```

## 构建与生成

常用命令：

```bash
make init
make api
make all
make build-all
make api-lint
```

单模块命令：

```bash
make -C app/user gen
make -C app/content build
make -C app/bbs doc
make -C app/bbs sdk
```

Docker 部署配置位于 `deploy/prod-docker`。

## 文档入口

- [项目规范索引](doc/README.md)
- [架构设计](doc/architecture.md)
- [代码规范](doc/coding.md)
