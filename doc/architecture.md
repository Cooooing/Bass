# 架构设计

本文档记录项目级架构边界和 `common` 封装约定。代码组织规则见 [代码规范](coding.md)。

## 服务概览

实现程度中的“规划中”表示服务仍在实现中，尚不具备对外提供完整领域功能的能力。

| 服务 | 分层 | 领域范围 | 实现程度 | 数据归属 |
|------|------|----------|----------|----------|
| `bbs` | 入口层 / BFF | 社区前台 HTTP API、端侧聚合、展示模型适配 | 已实现 | 无业务库 |
| `bbs_admin` | 入口层 / BFF | 管理后台 API、运营视角聚合 | 规划中 | 无业务库 |
| `openapi` | 入口层 / OpenAPI | 外部开放接口、签名、限流、授权适配 | 规划中 | 无业务库 |
| `push_node` | 入口层 / 实时连接 | SSE 长连接、客户端下行推送、节点接入 | 规划中 | 连接态和节点接入状态 |
| `user` | 内部业务服务 | 用户身份、账户、认证凭证、关系和账户设置 | 已实现 | 用户与认证数据 |
| `content` | 内部业务服务 | 社区内容、评论、标签、板块和互动 | 已实现 | 内容与互动数据 |
| `notify` | 内部业务服务 | 通知模板、通知记录、投递编排和通知偏好 | 已实现 | 通知数据 |
| `im` | 内部业务服务 | 会话、消息、已读状态、消息事件 | 规划中 | 即时通信数据 |
| `platform` | 内部业务服务 | 文件、对象存储、回调验签、第三方能力适配 | 已实现 | 平台能力数据 |
| `push_hub` | 内部业务服务 / 推送控制面 | 节点注册、在线路由、推送分发控制 | 规划中 | 推送路由和在线状态 |
| `scheduler` | 内部业务服务 | 定时任务定义、触发、执行记录和告警 | 已实现 | 任务调度数据 |
| `game_town` | 内部业务服务 | 多玩家文字世界、NPC、地点、记忆、事件和命令 | 规划中 | 文字世界数据 |

## 分层边界

- BFF 负责端侧访问边界和协议适配，不连接业务库，不拥有领域数据。
- 内部服务拥有本服务领域数据和业务规则；服务之间通过 RPC 或事件协作，不能 import 其他服务的 `internal`。
- `common` 只放公共 proto、客户端、基础工具和跨服务约定，不承载具体业务流程。
- `push_node` 只维护客户端 SSE 长连接；`push_hub` 负责实时推送路由；`im` 和 `notify` 只生成业务事实和可投递事件。

## 读写边界

- 一个业务事实只由一个归属服务写入。
- BFF 写接口通常只调用一个归属服务命令接口；确需多服务写入时，先设计补偿、Saga 或 TCC。
- BFF 不做 read-before-write 的领域状态预校验；资源存在性、作者关系、状态流转、本地业务不变量和并发写入正确性由归属服务校验。
- BFF 读接口可以聚合多个内部服务，并按端侧语义控制可见范围。
- 禁止为了查询方便跨服务连接业务库或维护跨服务数据库投影。

## 契约边界

- 契约分为 BFF HTTP、内部 gRPC、内部事件和外部回调。
- BFF proto 定义自己的 request、response 和对外结构，不引用内部服务 message。
- 内部服务 proto 只暴露 gRPC 契约，不写 `google.api.http` 注解。
- 业务模块 proto 不定义 `model.proto`，也不把领域模型放到独立公共 message 中复用。
- 每个 RPC 使用一个顶层 message，并且顶层 message 下面只能直接定义 `Request` 和 `Response`。
- 业务子 message 必须定义在实际使用它的 `Request` 或 `Response` 内；禁止在 RPC 顶层定义 `Account`、`Item`、`Summary` 这类业务模型。
- 不同 RPC 之间禁止复用业务子 message；即使字段完全一致，也要在各自的 `Request` 或 `Response` 内复制定义。
- 禁止引用 `OtherRpc.Response.Item`、`OtherRpc.Request.Query` 这类其他 RPC 的内部业务 message。
- 只有 `common` 中的分页、区间、公共枚举、公共错误数据和事件基础结构可以跨模块复用；业务服务、BFF 和入口层不能复用其他模块的业务 message。
- 枚举放在独立 `enum.proto`。
- `List`、`Page` 响应数组字段统一为 `rows`，分页信息字段统一为 `page`。

## common 封装

### 目录职责

- `common/proto/app/common` 放跨服务公共 proto、公共枚举、业务错误码和事件 payload。
- `common/proto/buf` 放 Buf 生成模板，API 生成、OpenAPI 生成和配置生成共用这里的模板。
- `common/pkg/client` 放 Redis、NATS、Consul、HTTP、Lark、Asynq 等基础客户端封装。
- `common/pkg/client/rpc` 放跨服务 gRPC client 封装。
- `common/pkg/constant` 放服务名、表前缀、上下文键、日志字段和公共常量。
- `common/pkg/enum` 放内部 string 枚举和 proto 枚举之间的映射。
- `common/pkg/apperror` 放业务错误构造、业务错误码提取和 HTTP 状态映射。

### RPC client

- 跨服务 gRPC client 统一放在 `common/pkg/client/rpc`。
- 每个被调用服务定义一个 client 包装类型，例如 `UserClient`、`ContentClient`、`NotifyClient`。
- client 构造函数接收 `*grpc.ClientConn`，内部持有对应 proto service client。
- provider 函数通过 `ConsulClient.GetGrpcConn` 获取连接，不在业务代码里直接 `grpc.Dial`。
- 服务名来自 `common/pkg/constant.ServiceName`，不能在业务代码中手写服务发现名称。
- 新增内部服务后，同步补充服务名常量、RPC client 包装和 provider 函数。

### 枚举

- 跨服务、跨层或落库约束枚举必须定义 proto enum。
- 内部业务代码可以使用 string 枚举，但必须通过 `common/pkg/enum.Mapping` 绑定 proto 枚举。
- 事件类型、事件主题和事件消费队列组使用 common 枚举，业务代码不能直接书写 MQ subject 和 queue group 字符串。
- `EventType` 每个服务保留独立数值段；`EventSubject` 数值与对应 `EventType` 保持一致。
- 新增枚举值时同步补充 proto 注释、内部枚举常量和映射表。

### 错误

- 对外可感知业务错误统一使用 `BusinessErrorCode`。
- 业务错误通过 `common/pkg/apperror.New` 构造，入口层根据业务错误码映射 HTTP 状态和用户可见文案。
- `ErrorReason` 只表示传输状态，不表达业务失败原因。
- 动态错误参数使用 common proto message，不使用自由 key/value 拼装。

### 事件

- 跨服务事件使用 `common.enums.Event`。
- 事件必须包含生产者生成的 `event_id`。
- 事件类型、事件主题和 payload 必须匹配。
- payload 只放最小事实字段，不嵌套引用其他业务服务展示模型。
- outbox/inbox 状态使用 common 枚举；表结构调整需要同步评估所有模块。

## 数据与事件

- schema 字段必须来自明确业务需求、查询路径、投递需求或审计需求。
- 表结构、outbox/inbox、事件 payload 和对外契约变更前，先说明设计、影响和回滚方式。
- 逻辑删除统一使用 `deleted_at`；业务查询默认排除已删除数据。
- 跨服务副作用由归属服务在本地事务内写 outbox，再通过 MQ 触发。
- 事件 payload 使用 common proto message，只携带最小事实字段；MQ subject 和 queue group 使用 common 枚举。

## 配置与运行

- 配置契约统一放在各模块 `internal/config/config.proto`，Go 封装放在 `internal/config/config.go`。
- 样例配置必须只使用 `config.proto` 已定义字段；新增配置项时同步更新 proto、样例配置和读取逻辑。
- 服务名、注册名和跨服务发现名使用 `common/pkg/constant`，业务代码不手写服务发现名称。
- Consul 同时承担配置中心和服务注册发现；连接类配置和敏感配置通过重启生效。
- 运行期业务阈值、限流、告警和事件轮询参数可以热重载，但不能改变数据库、Redis、NATS、Consul 等连接对象。
- 密钥、token、密码和 webhook secret 不写入 Git，不写入普通日志，不通过业务错误返回。

## 可观测性

- 跨服务调用、事件投递、事件消费、关键写事务必须记录 trace ID、业务 ID、耗时和结果状态。
- 日志字段优先使用 `common/pkg/constant` 中的公共字段名；新增公共日志字段先沉淀到 common。
- 错误日志记录稳定错误码、业务 ID 和失败摘要，不记录密码、token、验证码、密钥、完整请求体或消息正文。
- 外部通道响应只保存必要摘要；超长响应、敏感响应和用户隐私字段必须截断或脱敏。
- 指标用于观察吞吐、延迟、失败率、重试、dead letter 和限流命中，不承载业务事实。

## 运行边界

- Redis 只用于限流、短期幂等、短期缓存、验证码、会话临时态和防刷计数，不能作为领域事实来源。
- 对外可感知业务错误使用业务错误码；BFF 负责生成用户可见文案。
