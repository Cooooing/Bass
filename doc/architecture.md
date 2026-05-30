# 架构设计

## 服务边界

- BFF 负责端侧体验、认证上下文、权限判断、缓存、限流、降级和 DTO 适配；BFF 不连接业务库，不拥有领域数据。
- 内部服务拥有自己的领域数据和业务规则；服务之间只能通过 RPC 或事件协作，不能 import 其他服务的 `internal`。
- `common` 只放公共 proto、客户端、基础工具和跨服务约定，不承载具体业务流程。
- data 层适配数据库、Redis、RPC、对象存储和第三方调用；不承载跨业务流程，`Repo` 只表示数据库或持久化访问。

## 分层依赖

- 依赖只能从外层指向内层：service 调用 usecase；data 实现 biz 层定义的底层能力接口；biz 不能依赖 data。
- biz 层不能 import 本模块 `internal/data`、`internal/data/gen`、Ent 生成模型、Ent client、schema predicate 或 data 层实现类型。
- biz 层底层能力接口只能暴露业务模型、基础类型、枚举和查询参数对象；接口签名不能出现 Ent 或 data 层细节。
- service 层负责协议适配和入参组装，只能调用 biz/usecase；不能直接注入或调用 repo、data 层实现，现有历史偏离不作为新代码依据。
- 事务能力通过 biz 层 `Tx` 等抽象传入 usecase，不能通过 repo 参数暴露具体数据库 client。
- 常规表业务尽量保持 proto、service、biz、data、schema 一一对应；复杂流程新增明确业务 usecase，禁止泛化 helper 承载流程。
- BFF 适配内部服务调用时，biz/repo、biz/usecase 和 data 实现按具体业务面或内部 proto service 拆分；禁止按下游大服务聚合成 `user.go`、`content.go`、`notify.go` 这类总入口文件或大接口。
- usecase 可以按业务层级组合调用，但必须明确上下游，禁止循环调用；一个写流程只能有一个清晰事务入口。

## 读写边界

- 一个业务事实只能由一个归属服务写入；主业务写入在归属服务本地事务内完成。
- BFF 写接口通常只调用一个归属服务命令接口；确需多服务写时，必须先设计 Saga、TCC 或补偿方案。
- 跨服务副作用由归属服务写 outbox 后通过 MQ 触发；BFF 不直接发布领域 MQ 事件。
- 普通数据库写入失败直接返回失败，不为普通写流程额外设计重试；需要重试、补偿和 MQ 投递保证的流程放在 outbox/inbox。
- BFF 读接口可以聚合多个内部服务，允许短暂最终一致；非归属服务读取业务事实默认通过 RPC。
- 出于性能需要缓存时只能使用 Redis，并设置 TTL；缓存不能写入业务数据库，不能作为领域事实来源。
- 禁止为了查询方便跨服务连接业务库或维护跨服务数据库投影。
- 新增同步服务依赖前必须说明调用方向、失败影响和降级策略。

## 契约设计

- 契约分为 BFF HTTP、内部 gRPC、内部事件和外部回调四类；外部回调必须独立说明验签、幂等键、来源字段和失败处理。
- Proto 管理、lint 和生成统一使用 Buf；模板放在 `common/buf`，共享 API 配置放在 `common/api/app`，第三方 proto 依赖通过 Buf deps 管理。
- 无法从 Buf deps 获取且必须长期稳定的扩展，可以定义为项目自有 proto，并放在 `common/api/app` 下。
- 一个 proto 文件最多定义一个 `service`，每个 proto service 对应一个 Go service 文件。
- 每个 RPC 必须写简短注释；注释只描述接口功能，不写消费方、调用场景、登录态、幂等策略和字段缺省规则。
- 请求和响应必须围绕接口语义定义，不复用过大的通用模型；查询 RPC 按资源收敛为单资源 `Get`、列表 `List` 和映射 `Map`，禁止因返回字段组合拆分 `GetBasic`、`BatchGetBasic`、`BatchGetContact` 这类接口。
- 调用方需要按 ID 获取映射时，归属服务必须定义独立 `Map` RPC；`List` 只表示列表查询并返回列表，调用方不能通过 `List` 的 `repeated` 结果自行组装映射。
- 基础资源模型只表达资源本体，不挂载可独立查询的关联集合；关联集合通过独立 RPC 暴露，不能由 `Get` 或 `List` 隐式聚合返回。
- 项目初期不维护 proto 向后兼容保留位，删除字段或枚举值时直接删除，不写 `reserved` 字段号或字段名。
- 请求参数不能用 `with_xxx`、`include_xxx` 等布尔开关控制响应结构；同一接口的响应结构必须由接口语义固定。
- 内部服务 proto 只暴露 gRPC 契约，不写 `google.api.http` 注解；OpenAPI 只面向 BFF。
- BFF proto 必须定义自己的 request、reply 和对外模型，不引用内部服务 proto message。
- 项目初期不要求 proto 版本兼容；按当前需求可以删除、重命名或重排字段。
- RPC 命名使用业务动作，禁止 `GetOne`、`Page`、`Info`、`Data` 等实现视角名字；service 名已表达资源时，RPC 只写动作。
- 单资源查询用 `Get`，集合查询用 `List`，创建实体用 `Create`，追加子内容可用明确业务动词，状态变更使用 `Publish`、`MarkRead` 等业务动作。
- request/reply 外层 message 使用 `动作 + 资源名`；集合查询使用资源复数，避免 `UpdateArticleArticle` 这类重复资源名。
- BFF 面向客户端默认不提供批量接口；确有批量需求时必须来自明确客户端流程，并在 RPC 名中使用 `Batch`。
- 高频查询接口可以拆细，不做万能查询接口；分页、排序、过滤字段必须有默认值和边界说明。
- 普通 BFF HTTP 对外业务接口默认使用 `POST` 和 `body: "*"`；回调、健康检查、直接返回图片或文件等需要被浏览器资源标签引用的接口可以按场景使用 `GET`。
- Proto 文件格式以 `buf format` 输出为准，不为 `google.api.http` option 手工调整出与 `buf format` 不一致的格式。BFF HTTP 路径统一使用 `/v1/{模块}/{资源}/{动作}`，资源名称使用单数，动作使用 lower-kebab，不暴露内部投影名、表名或 RPC 实现名。
- BFF 不暴露邮箱、手机号、账号名是否存在等可用于枚举用户信息的接口。
- BFF OpenAPI 从 `google.api.http` 和 proto 注释生成；文档入口放在 BFF proto 包的 `doc.proto`。
- BFF OpenAPI 不手写 tags、servers、external_docs、BearerAuth，也不写登录态、幂等策略和字段缺省规则。
- SDK 只从 BFF OpenAPI 生成，TypeScript Axios 输出到 `common/api/gen-typescript-axios/<bff>`，TypeScript Fetch 输出到 `common/api/gen-typescript-fetch/<bff>`，Go、Java、Rust 分别输出到 `common/api/gen-go/<bff>`、`common/api/gen-java/<bff>` 和 `common/api/gen-rust/<bff>`，生成产物不入 Git 和 Docker。SDK 生成时去掉 OpenAPI operationId 的 service 前缀和 API 后缀，客户端对象使用 proto service 名，方法使用 RPC 动作名。
- gRPC client 统一通过 `common/pkg/client/rpc` 和 `ConsulClient.GetGrpcConn` 创建；业务代码不要直接 `grpc.Dial`、手写服务发现或为普通 RPC 重复包 `context.WithTimeout`。

## 数据设计

- 模块内部实体使用业务概念本身，数据库表名保留服务或模块前缀。
- 常规业务表使用统一审计字段；可缺省扩展表允许不存在记录，通过 upsert 按归属键更新。
- schema 字段必须来自明确业务需求或实际查询、投递、审计需求；意义不明、无读写路径、可推导或重复的信息不落库。
- 日志表只记录排查和安全审计真正需要的事实；没有明确消费方的提交原文、失败文案、链路标识和客户端解析字段不落库。
- JSON 字段只用于结构可变、查询要求不强的内容。
- 偏好设置按业务归属拆分；端侧展示、本地化等账号体验设置与通知投递、订阅、渠道开关分表维护。
- 任何数据库表结构变更都必须先给出设计、影响范围和理由，经用户确认后再修改。
- 唯一约束必须对应真实业务幂等或唯一性；参与唯一约束的字段应避免可空。
- 索引根据查询接口和补偿扫描设计；复合索引能覆盖左前缀查询时，不额外增加重复单列索引。
- Ent schema 中表、字段、关键索引应保留中文注释。
- Ent 生成模型只允许在 data 层内部使用；biz model 必须是独立业务结构，不能嵌入或别名引用 Ent entity。
- 业务固定取值必须抽成枚举；跨服务、跨层或落库约束枚举必须定义 proto enum，并通过 `common/pkg/enum.Mapping` 绑定内部 string enum。
- MQ subject 和 queue group 属于跨服务事件协议，必须定义为 common 枚举；业务代码不能直接书写字符串。

## 数据可见范围

- 生成代码时必须根据业务语义判断字段是否属于隐私信息，不能只依赖字段名白名单。
- 密码、token、密钥任何时候都不能返回给用户，也不能写入日志、错误信息或事件 payload；验证码只允许写入验证码投递事件，不能写入日志、错误信息或其他事件 payload。
- 邮箱、手机号、账号绑定信息、设备信息、IP 等字段按业务可见范围返回；默认只允许用户本人或明确授权的管理场景可见。
- BFF 对外模型只包含当前接口需要展示的字段；不因为内部服务返回了字段就继续透出。

## 事件与缓存

- 生产者在本地事务内写 outbox；消费者使用 inbox 做幂等、重试和补偿。
- 事件包含全局幂等 ID、事件类型、subject、发生时间、接收者和业务 payload。
- outbox repo 只提供通用保存能力；事件类型、主题、payload 和接收者由调用方按业务语义构造。
- outbox 默认保持最小投递模型；确需新增字段时，必须先给出必要性设计并经用户确认。
- 所有模块的 outbox、inbox 表结构必须保持同类一致；新增、删除或调整字段时必须同步处理所有模块。
- 事件 payload 使用 common proto message；消费幂等以 event ID 为主，必要时结合事件类型和业务主键。
- 事件处理失败必须保留可判断状态的 inbox 记录；是否增加死信、搁置、下次重试等字段必须先设计确认。
- subject、consumer group、durable 名称必须稳定、可追踪。
- Redis 只用于限流、短期幂等、短期缓存、验证码、会话临时态和防刷计数。

## 运行约束

- server、service、wire、middleware 按模块统一组织；新增组件必须进入对应 ProviderSet。
- BFF 负责入口认证和权限判断；内部服务默认信任内部调用链身份上下文，不重复实现入口层权限判断。
- BFF 解析登录令牌必须调用 user 服务认证接口获取通用用户上下文；BFF 不直接读取 user 登录令牌的 Redis 缓存。
- 内部服务仍需校验资源存在、状态流转、本地业务不变量和写入前置条件。
- 身份信息通过统一 context key 传递，禁止字符串 key 散落。
- 需要改变内部 RPC 信任边界或增加服务间鉴权时，必须先明确边界并统一在 middleware 实现。
- 常规 middleware 包括 recovery、logging、tracing、auth、validate、timeout；BFF 和内部服务按职责启用。
- 需要变动的配置从 `.env` 加载，不在业务代码硬编码；新增或变更配置必须同步配置 proto 和 `.env` 示例，并添加中文注释。
- secret、token、密钥不写入 Git；示例配置只能使用无效占位值。
- BFF 读接口可以按业务场景降级或返回部分结果；写接口下游失败时默认返回失败，不吞错伪造成功。
- 参数、鉴权、权限、资源不存在、状态冲突、下游失败要区分；下游错误至少区分超时、不可用、业务拒绝、资源不存在和状态冲突。
- data 层数据库访问实现负责把数据库错误转换为数据访问语义错误；对外错误信息不能泄露密码、token、验证码、SQL 或内部拓扑。
- 跨服务调用、事件投递和消费、关键写事务必须记录 trace ID、业务 ID、耗时和结果状态等可追踪信息；关键写流程、outbox/inbox、缓存、限流、下游调用需要指标。
