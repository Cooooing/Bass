# 架构设计

## 服务与分层

- BFF 负责端侧访问边界和协议适配，包括入口认证、端侧权限、请求格式、默认值、读接口可见范围、缓存、限流、降级、跨服务展示资料聚合和 DTO 适配；BFF 不连接业务库，不拥有领域数据。
- 内部服务负责领域事实边界和写入正确性，拥有自己的领域数据和业务规则；服务之间只能通过 RPC 或事件协作，不能 import 其他服务的 `internal`。
- `common` 只放公共 proto、客户端、基础工具和跨服务约定，不承载具体业务流程。
- 分层依赖只能从外层指向内层：service 调用 usecase，data 实现 biz 底层能力接口；service 不直接注入或调用 repo/data，biz 不 import data、Ent、schema predicate 或 data 实现类型，biz 接口签名只暴露业务模型、基础类型、枚举和查询参数对象。
- 常规表业务尽量保持 proto、service、biz、data、schema 一一对应；复杂流程新增明确业务 usecase。

## 读写边界

- 一个业务事实只能由一个归属服务写入，主业务写入在归属服务本地事务内完成；BFF 写接口通常只调用一个归属服务命令接口，确需多服务写时必须先设计 Saga、TCC 或补偿方案。
- BFF 写接口不做 read-before-write 的领域状态预校验，例如不在写入前额外查询文章状态再判断能否点赞、发布、编辑或采纳；资源存在性、作者关系、状态流转、本地业务不变量和并发写入正确性由归属服务校验。
- BFF 可以校验明显的协议和端侧输入问题，例如必填字段、分页范围、枚举值能否映射、当前端是否允许访问某类接口；这些校验不能替代内部服务的领域校验。
- 跨服务副作用由归属服务写 outbox 后通过 MQ 触发；BFF 不直接发布领域 MQ 事件；普通数据库写入失败直接返回失败，需要重试、补偿和 MQ 投递保证的流程放在 outbox/inbox。
- BFF 读接口可以聚合多个内部服务，并按端侧语义控制可见范围；列表、详情、评论、附言等展示接口都必须由 BFF 决定当前端和当前用户能看到什么。非归属服务读取业务事实默认通过 RPC；禁止为了查询方便跨服务连接业务库或维护跨服务数据库投影。
- 缓存只能使用 Redis 并设置 TTL，不能写入业务库或作为领域事实来源；新增同步服务依赖前必须说明调用方向、失败影响和降级策略。

## 校验边界

| 校验类型 | 归属 | 说明 |
|----------|------|------|
| 入口认证、当前用户上下文解析 | BFF | 内部服务不从 ctx、metadata、请求头、JWT 或会话中提取当前用户。 |
| 端侧访问范围 | BFF | 前台、管理端、开放 API 等入口能调用哪些接口、能传哪些查询条件，由对应 BFF 控制。 |
| 请求格式和端侧默认值 | BFF 为主，内部服务可兜底 | BFF 提前返回端侧错误；内部服务保护自己的 gRPC 契约。 |
| 读接口可见范围 | BFF | 同一内部查询能力可服务前台、作者视角和管理端，端侧展示边界不写死在内部服务。 |
| 资源存在性 | 内部服务 | 归属服务拥有数据，必须自行判断目标资源是否存在。 |
| 写入身份关系 | 内部服务 | 例如传入的 `user_id` 是否为作者、`operator_id` 是否写入审计字段。 |
| 状态流转 | 内部服务 | 例如草稿发布、发布后编辑窗口、归档、管理隐藏、管理锁定。 |
| 本地业务不变量 | 内部服务 | 例如标签存在且启用、问答文章才能采纳答案、评论必须属于目标文章。 |
| 事件发送语义 | 内部服务 | 只有业务事实真正变化时才写 outbox。 |

边界原则：BFF 判断“这个入口能不能发起这个请求”；内部服务判断“这个业务事实能不能成立”。BFF 可以提前拦截明显错误，但不能成为领域状态正确性的唯一保证。

## 契约设计

- 契约分为 BFF HTTP、内部 gRPC、内部事件和外部回调；外部回调必须独立说明验签、幂等键、来源字段和失败处理。
- Proto 管理、lint、format 和生成统一使用 Buf；模板放在 `common/proto/buf`，共享 API 配置放在 `common/proto/app`，第三方 proto 依赖通过 Buf deps 管理；一个 proto 文件最多定义一个 `service`，每个 proto service 对应一个 Go service 文件，项目初期删除 proto 字段或枚举值时不写 `reserved`。
- 业务枚举必须放在独立 `enum.proto`，`model.proto` 不定义 `enum`；BFF 可以定义自己的对外枚举文件，不强制复用内部服务枚举，避免对外契约绑定内部领域取值。
- `model.proto` 只放稳定数据模型，不放查询条件、保存参数、回调参数、事件 payload 或单个接口专用返回结构；`Query`、`Save`、`Callback` 等接口参数应放在对应 service proto 中，多个 RPC 共享时也只在同一资源接口文件内复用。
- 内部服务的业务输入只能来自内部 gRPC proto request 的显式字段；需要当前用户、操作人、审计人、目标资源、动作参数或写入值时，都必须在 request 中定义字段，由 service 显式读取并传入 usecase。
- 内部服务 proto 只暴露 gRPC 契约，不写 `google.api.http` 注解；BFF proto 必须定义自己的 request、reply 和对外模型，不引用内部服务 proto message；OpenAPI 只面向 BFF。
- 请求和响应围绕接口语义定义，不复用过大的通用模型；同一接口的响应结构不能由 `with_xxx` 或 `include_xxx` 开关控制。
- `List`、`Page` 类接口响应中，数据数组字段统一命名为 `rows`；分页信息字段使用 `page`，禁止混用 `items`、`list`、`data_list` 等数组字段名。
- BFF 返回模型按视图语义拆分，不用一个大模型同时覆盖列表、详情、创建返回和嵌套摘要；嵌套对象必须使用最小可展示结构，例如评论中需要文章信息时使用文章摘要模型，不嵌完整详情模型。
- 内部业务服务优先返回归属领域事实和 ID，不直接引用其他业务服务的展示模型；需要跨服务展示资料时由 BFF 聚合。
- RPC 命名使用业务动作，禁止 `GetOne`、`Info`、`Data` 等实现视角名字；查询按资源收敛为 `Get`、`List`、`Page` 和 `Map`，其中 `Page` 表示分页查询；需要按 ID 获取映射时由归属服务定义独立 `Map` RPC；service 名已表达资源时，RPC 只写动作，request/reply 外层 message 使用 `动作 + 资源名`。
- 普通 BFF HTTP 对外业务接口默认使用 `POST` 和 `body: "*"`；回调、健康检查、直接返回图片或文件等可以按场景使用 `GET`；BFF HTTP 路径统一使用 `/v1/{模块}/{资源}/{动作}`，资源名称使用单数，动作使用 lower-kebab。
- BFF 不暴露邮箱、手机号、账号名是否存在等可用于枚举用户信息的接口。
- BFF OpenAPI 从 `google.api.http` 和 proto 注释生成，不手写 tags、servers、external_docs、BearerAuth 或运行时说明；SDK 只从 BFF OpenAPI 生成，生成产物不入 Git 和 Docker，客户端对象使用 proto service 名，方法使用 RPC 动作名。

## 数据设计

- schema 字段必须来自明确业务需求或实际查询、投递、审计需求；意义不明、无读写路径、可推导或重复的信息不落库；JSON 字段只用于结构可变、查询要求不强的内容。
- 日志表只记录排查和安全审计真正需要的事实；没有明确消费方的提交原文、失败文案、链路标识和客户端解析字段不落库。
- 任何数据库表结构变更都必须先给出设计、影响范围和理由，经用户确认后再修改。
- 唯一约束必须对应真实业务幂等或唯一性，参与唯一约束的字段应避免可空；索引根据查询接口和补偿扫描设计，复合索引能覆盖左前缀查询时不额外增加重复单列索引。
- 逻辑删除使用 `deleted_at` 表达；业务 repo 查询固定过滤 `deleted_at IS NULL`，应用层不提供查询已删除或包含已删除数据的入口。
- 业务表不保留删除操作者字段；删除、隐藏、锁定等操作的操作者、原因和来源写入内容审核记录或审计记录。PostgreSQL 软删除场景下的业务唯一约束使用 `WHERE deleted_at IS NULL` 的部分唯一索引，不用可空 `deleted_at` 参与普通唯一约束。
- Ent 生成模型只允许在 data 层内部使用；biz model 必须是独立业务结构，不能嵌入或别名引用 Ent entity。
- 业务固定取值必须抽成枚举；跨服务、跨层或落库约束枚举必须定义 proto enum，并通过 `common/pkg/enum.Mapping` 绑定内部 string enum。
- 密码、token、密钥不返回、不记录日志、不写错误和事件 payload；验证码只允许写入验证码投递事件；邮箱、手机号、设备信息、IP 等字段按业务可见范围返回；BFF 对外模型只包含当前接口需要展示的字段。

## 事件与缓存

- 生产者在本地事务内写 outbox；消费者使用 inbox 做幂等、重试和补偿，事件处理失败必须保留可判断状态的 inbox 记录。
- outbox repo 只提供通用保存能力，事件类型、主题、payload 和接收者由调用方按业务语义构造；outbox 默认保持最小投递模型，确需新增字段时必须先给出必要性设计并经用户确认。
- 所有模块的 outbox、inbox 表结构必须保持同类一致；新增、删除或调整字段时必须同步处理所有模块。
- outbox publisher 使用 Redis 批次锁降低多实例扫表压力；Redis 锁只负责减少无效查询，正确性仍由 outbox 状态、数据库行锁和 event ID 幂等保证。获取 Redis 锁失败时跳过本轮投递，不查库。
- outbox 和 inbox 通过 `event.outbox.max_retry`、`event.inbox.max_retry` 推进 dead 状态；领取方法只处理可重试状态和超时处理中状态，不领取 dead 记录。
- dead letter 告警由各服务本地扫描本服务 outbox/inbox，使用 Redis 去重和 common Lark Webhook client 做可选告警；告警失败只记录日志和指标，不写业务事件。
- 事件 payload 使用 common proto message；消费幂等以 event ID 为主，必要时结合事件类型和业务主键；MQ subject 和 queue group 属于跨服务事件协议，必须定义为 common 枚举，业务代码不能直接书写字符串。
- Redis 只用于限流、短期幂等、短期缓存、验证码、会话临时态和防刷计数。
- notify 邮件和短信发送限流使用 Redis 精确滑动窗口；发送前判断并占用额度，命中限流记录通道 `rate_limited` 状态，不调用第三方通道；需要同步感知触达限流的写接口，在写入验证码缓存和 outbox 前调用 notify 只读限流查询。

## 运行与错误

- BFF 负责入口认证和权限判断，解析登录令牌必须调用 user 服务认证接口获取通用用户上下文，不能直接读取 user 登录令牌 Redis 缓存；内部服务不从 ctx、metadata、请求头、JWT、缓存会话等隐式来源提取业务身份或业务参数，写接口需要用户身份或审计字段时必须由 proto 参数显式传入；内部服务仍需校验资源存在、状态流转、本地业务不变量和写入前置条件。
- 内部服务可以从本服务数据库读取领域事实，可以使用配置、当前时间、UUID、事件 ID 等系统生成值；ctx 只用于取消、超时、trace 和事务传递，不作为业务参数来源。
- gRPC client 统一通过 `common/pkg/client/rpc` 和 `ConsulClient.GetGrpcConn` 创建；业务代码不要直接 `grpc.Dial`、手写服务发现或重复包 `context.WithTimeout`。
- Consul 服务注册统一关闭主动健康探测，使用 TTL heartbeat 做失效剔除；公共 Consul client 不读取 `server.mode` 等业务配置做环境分支。
- 需要变动的配置从 `.env` 加载；新增或变更配置必须同步配置 proto、`configs/config.yaml` 和 `.env` 示例。secret、token、密钥不写入 Git。
- 服务配置契约以各模块 `internal/conf/conf.proto` 为准，配置样例不能出现 proto 未定义字段。`google.protobuf.Duration` 配置统一使用秒单位，禁止写 `ms`、`m`、`h` 等单位；布尔配置必须使用 `true` 或 `false` 字面量，禁止使用环境变量占位。热重载只面向运行期参数，连接类配置和敏感凭证必须通过重启生效。
- BFF HTTP 响应体统一使用 `code`、`message`、`data`、`time`，HTTP 状态使用真实状态码；`ErrorReason` 只表示 HTTP/gRPC 传输状态，业务错误码统一使用 `BusinessErrorCode` 并通过 `errors.code` 绑定 HTTP 状态。
- BFF 负责把业务错误码按用户上下文中的 `common.enums.Language` 转换成用户可见 `message`，匿名或未指定语言时默认使用简体中文；内部错误只允许把 common proto message 定义的动态参数透出到 `data`，并由公共错误工具按 proto JSON 编码。
- 对外可感知的业务错误必须通过 `common/pkg/apperror` 构造；普通基础设施错误可以保留原始 error，由入口编码器映射为通用内部错误。
- 跨服务调用、事件投递和消费、关键写事务必须记录 trace ID、业务 ID、耗时和结果状态。
