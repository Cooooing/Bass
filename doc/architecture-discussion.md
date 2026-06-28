# 架构讨论记录

| 日期 | 议题 | 结论 |
|------|------|------|
| 2026-06-08 | BFF 写操作校验边界 | 写入状态校验下沉到内部服务，BFF 保留端侧边界控制 |
| 2026-06-09 | MQ 事件驱动改进 | 4 项待落地 + 1 项已符合，详见下文 |

---

## 一、BFF 与内部服务校验边界

### 现状

BFF usecase 层在写操作前执行 read-before-write 校验（LikeArticle、UpdateArticle、PublishArticle、AcceptAnswerArticle 等共 8 个方法），同时内部 content 服务在写 RPC 内部也执行同样的状态校验，导致：

1. 同一份业务规则在两处维护，存在漂移风险
2. 每次写操作至少 2 次 RPC 往返（1 次读 + 1 次写）

### 结论

BFF 写接口不做 read-before-write 状态预校验，状态流转和业务不变量完全交给内部服务。

**BFF 职责**：

- 入口认证，解析当前用户上下文
- 端侧权限判断，例如当前端是否允许访问管理接口、是否允许传入某类查询条件
- 请求格式校验，例如必填字段、分页参数、枚举值是否能映射到内部服务枚举
- 用户身份、端侧语义和默认值注入，例如 `user_id`、默认排序、前台列表默认只查公开内容
- 读接口可见范围控制，例如列表、详情、评论、附言等端侧展示边界
- 调用归属服务写 RPC，并做返回 DTO 适配

**内部服务职责**：

- 资源存在性校验
- 作者、操作者、审计用户等写入身份校验
- 状态流转校验，例如草稿发布、发布后编辑窗口、归档、管理隐藏、管理锁定
- 本地业务不变量，例如文章类型与问答参数匹配、标签存在且启用、评论所属文章可评论
- 幂等与并发写入控制
- 事件发送语义控制，即只有业务事实真正变化时才写 outbox
- 返回明确业务错误码

**内部服务输入来源约束**：

- 业务输入只允许来自 proto request 显式字段，例如 `user_id`、`operator_id`、资源 ID、动作参数和写入值
- 禁止从 ctx、metadata、请求头、JWT、缓存会话等隐式来源提取当前用户或业务参数
- ctx 只用于取消、超时、trace 和事务传递
- 内部服务可以读取本服务数据库中的领域事实，可以使用配置、当前时间、UUID、事件 ID 等系统生成值参与判断

这样划分的原因：

- 归属服务拥有领域数据，只有它能在同一个事务或同一份最新数据上判断状态，避免 BFF 读到旧状态后再写入形成竞态。
- 同一份状态流转规则只维护在归属服务，避免 BFF、管理端 BFF、开放 API 等多个入口重复实现并发生漂移。
- BFF 面向端侧体验，适合处理“这个端能看什么、能传什么、默认查什么”，不适合持有领域状态机。
- 内部服务仍不是权限空壳。它不判断端侧权限，但必须保证本地业务事实不会被非法写入。

### 边界规则

| 校验类型 | 归属 | 说明 |
|----------|------|------|
| 必填字段、分页范围、枚举格式 | BFF 和内部服务都可以校验 | BFF 用于提前返回端侧错误；内部服务保护自身契约。 |
| 当前用户是否登录 | BFF | 内部服务不从 ctx、metadata 或请求头取用户，只接收 proto 参数。 |
| 用户 ID 是否存在、是否为作者、是否可操作该资源 | 内部服务 | 属于资源写入前置条件，必须由归属服务判断。 |
| 前台、管理端、开放 API 的访问范围 | BFF | 属于端侧权限，不写入归属服务通用业务规则。 |
| 资源状态流转 | 内部服务 | 例如发布、归档、隐藏、锁定、编辑窗口、浏览量限制。 |
| 读接口可见范围 | BFF | 列表和详情都由 BFF 按端侧语义过滤或拒绝。 |
| 跨服务展示资料聚合 | BFF | 例如作者资料、viewer action state、最后回复用户。 |
| 本地业务不变量 | 内部服务 | 例如标签启用、文章可评论、问答文章才能采纳答案。 |

**待清理代码**：`app/bbs/internal/biz/usecase/content_article.go` 中写操作的 read-before-write 状态预校验。

**前置条件**：内部服务写 RPC 需先确保校验完备、错误码返回充分。

---

## 二、MQ 事件驱动改进

### 现状

- 生产者（user/content）transactional outbox，业务写入和 outbox.Save 在同一事务
- OutboxPublisher 轮询 outbox 表，指数退避 1s→8s，FOR UPDATE SKIP LOCKED
- 消费者（notify）deduplicating inbox，event_id 幂等
- 事件 payload 定义在 common/enums/event.proto 的单一 Event oneof
- outbox/inbox 有 dead 状态但无最大重试次数，无告警机制
- 业务层事件发送语义目前已基本符合要求，后续新增写操作仍需遵守同一规则

### 改进 1：保持 outbox + 轮询模式

不引入即时发布机制（事务提交后直接发 MQ + 回写 outbox 状态）。

理由：额外的 MQ 发送 + outbox 状态回写引入更多边界情况（MQ 成功但回写失败、重试次数不准确），多一次库写入增加复杂度，收益不明显。

如果未来某个场景对延迟敏感（如 IM 消息投递），针对该场景单独设计。

### 改进 2：Redis 批次锁避免多实例并发扫表

多个服务实例各自轮询 outbox 表会导致大量 FOR UPDATE SKIP LOCKED 查询，增加 DB 压力。

方案：Redis 分布式锁，每轮 poll 前检查锁，有锁则跳过，无锁则写入锁后执行。锁只用于减少无效扫表，不作为投递正确性的唯一保证；正确性仍由 outbox 表状态、行锁和幂等 event_id 保证。

- 锁 key 按服务划分，例如 `outbox:publisher:{service}`
- 锁 value 使用实例 ID + 随机 token，释放时用 Lua 脚本比较 token 后删除
- 锁 TTL 覆盖单批次处理的最长预期时间，而不是只大于 poll interval
- 单批次处理可能较长时，使用较小批次或增加续期机制；续期失败则当前批次继续依赖数据库状态完成，不再领取新批次
- 持锁者崩溃时 TTL 自动过期，不影响后续
- outbox 表本身作为兜底，持锁者失败的事件留在 pending
- 已领取但未完成的 publishing 事件通过 stale timeout 重新进入可领取范围
- 获取不到锁的实例本轮不查库，直接进入下一次 poll 等待

### 改进 3：保持单一 Event oneof

不拆分为独立 event proto + bytes payload。

理由：当前事件数量有限（user ~12、content ~12），数值分区已做逻辑隔离，消费者只需一个反序列化入口，维护成本最低。

约束：event.proto 变更走 review，payload 只放最小必要字段，不嵌套引用其他服务 model。

### 改进 4：Dead Letter 处理

**最大重试次数**：outbox 和 inbox 使用统一配置值（如 5 或 10），但字段命名保留各自语义。

- outbox 使用 `retry_count`，表示投递失败后的重试次数
- inbox 使用 `attempt_count`，表示消费者处理尝试次数

**状态推进**：

- outbox 发布失败时增加 `retry_count`
- `retry_count < max_retry` 时标记为 `failed`
- `retry_count >= max_retry` 时标记为 `dead`
- `ClaimForPublish` 只领取 `pending`、`failed`、超时的 `publishing`，不领取 `dead`
- inbox 处理失败时保留 `attempt_count`
- `attempt_count < max_retry` 时标记为 `failed`
- `attempt_count >= max_retry` 时标记为 `dead`
- `ClaimRetry` 只领取 `received`、`failed`、超时的 `processing`，不领取 `dead`

**配置建议**：

- `event.outbox.max_retry`
- `event.inbox.max_retry`
- `event.outbox.publish_timeout`
- `event.inbox.processing_timeout`
- `event.dead_letter.alert_dedup_window`

**分层告警**：

| 层级 | 通道 | 依赖 | 用途 |
|------|------|------|------|
| 第一层 | ERROR 日志 + Prometheus counter | 零外部依赖 | 本地自保 |
| 第二层 | common Lark client → HTTP Webhook | Redis + HTTP（绕过 MQ/notify） | dead letter 告警 |
| 业务层 | outbox → MQ → notify → Lark | 完整事件链路 | 运营通知、业务通知 |

- common 模块新增独立 Lark client，封装 Lark Webhook HTTP 调用
- dead letter 告警走独立通道，与业务通知互不依赖
- 告警是 best-effort，失败只写日志 + metric，不重试不产生事件
- 同一 event_type 告警加去重窗口（如 1 小时内只发一次）
- 告警任务按服务本地扫描 dead 记录，只读取本服务 outbox/inbox，不跨服务连接业务库
- 告警内容包含 service、event_id、event_type、subject、retry/attempt 次数、最近错误、更新时间

### 改进 5：业务层控制事件发送语义（现有实现已符合）

当前 content 服务的 Like/Thank/Collect/Watch 和 user 服务的 Follow 均已实现：先查询状态，再根据 created/deleted 结果决定是否执行后续写入（AddStats、outbox.Save）。outbox.Save 只在业务状态真正变化时才执行，无变化时不产生事件。无需额外改动。

---

## 影响评估

| 改进项 | 影响范围 | 优先级 |
|--------|---------|--------|
| BFF 写状态预校验清理 | app/bbs usecase + content 服务写 RPC 校验完备性 | 高 |
| Redis 批次锁 | common/pkg 新增锁工具 + 各服务 OutboxPublisher | 高 |
| 业务层事件发送语义 | 现有实现已符合，无需改动 | — |
| Dead Letter 最大重试 | outbox/inbox repo 的 MarkFailed | 高 |
| common Lark client | common/pkg 新增包 | 中 |
| Dead Letter 告警定时任务 | 各服务新增定时任务 | 中 |
| Event oneof | 不变 | — |
| 投递延迟 | 不变 | — |
