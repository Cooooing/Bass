# notify 设计说明

## 业务需求

**notify 是内部通知服务**，负责消费其他服务产生的业务事件，并按事件类型生成站内信、邮件、腾讯云短信或 Lark webhook 通知。

### 事件输入

- 通知来源以**跨服务事件**为主。
- 事件消息使用 `common.enums.Event`。
- 每条事件必须携带生产者生成的 UUID：`event_id`。
- **事件类型、事件主题和 payload 必须匹配**。
- **事件 payload 只包含事件事实的最小标识字段**，不携带标题、昵称、邮箱、手机号等展示字段或触达字段。

| 事件 | payload |
| --- | --- |
| 用户注册 | 注册用户 ID |
| 用户关注 | 关注者用户 ID、被关注者用户 ID |
| 文章发布 | 文章 ID |
| 文章点赞 | 文章 ID、点赞用户 ID |
| 文章感谢 | 文章 ID、感谢用户 ID |
| 文章收藏 | 文章 ID、收藏用户 ID |
| 文章关注 | 文章 ID、关注用户 ID |
| 评论发布 | 评论 ID |
| 评论点赞 | 评论 ID、点赞用户 ID |

### 通知通道

**站内信、邮件、腾讯云短信、Lark webhook 都是通知通道**。它们共享事件消费、规则匹配和幂等控制，但模板字段、目标字段和发送记录按通道分开。

| 通道 | 目标 | 载荷 | 结果记录 |
| --- | --- | --- | --- |
| `station` | 用户 ID | 标题、内容 | 已读时间 |
| `email` | 邮件地址 | 收件邮箱、主题、正文、正文类型 | 发送状态、服务商消息 ID、响应摘要 |
| `tencent_sms` | 手机号 | 手机号、腾讯云应用 ID、签名、模板 ID、有序模板参数 | 发送状态、腾讯云请求 ID、返回码、返回消息 |
| `lark_webhook` | Lark 模板配置 | webhook token、消息类型、content 模板、可选签名密钥 | 发送状态、HTTP 状态码、响应体 |

### 通知规则与模板

通知规则按**事件类型、通道和语言**维护。生成通知时固定使用 `zh_CN`。

`notification_rule` 保存通用规则属性：

- 事件类型。
- 通知通道。
- 语言。
- 启用状态。

通道模板表保存通道专属字段：

| 表 | 字段 |
| --- | --- |
| `notification_station_template` | `rule_id`、`title_template`、`content_template` |
| `notification_email_template` | `rule_id`、`subject_template`、`body_template`、`content_type` |
| `notification_tencent_sms_template` | `rule_id`、`sms_sdk_app_id`、`sign_name`、`provider_template_id`、`param_templates` |
| `notification_lark_webhook_template` | `rule_id`、`webhook_id`、`token`、`secret`、`msg_type`、`content_template` |

`notification_rule` 与对应通道模板表是一对一关系。`notification_rule.id` 只作为模板表外键，不进入通道发送记录。

规则唯一约束：

```text
event_type + channel + language
```

**没有匹配的启用规则或通道模板时，事件直接处理完成**。单个通道模板缺失、目标缺失或渲染失败时，不生成该通道记录。

Lark webhook 的 `content_template` 只保存 Lark 请求体中的 `content` 对象模板，`msg_type` 和可选签名字段由 notify 发送层组装。`secret` 配置为空时不启用签名。

### 收件人模型

`NotificationRecipient` 使用宽对象表达收件人和触达目标：

| 字段 | 说明 |
| --- | --- |
| `user_id` | 系统用户 ID，可为空 |
| `email` | 直接邮箱，可为空 |
| `phone` | 直接手机号，可为空 |

`NotificationRecipient` 不包含 Lark webhook 目标。Lark webhook 的 `webhook_id` 和 `token` 配置在 `notification_lark_webhook_template` 中，事件侧只提供模板变量。

### 幂等与投递语义

- **事件幂等只使用生产者生成的 `event_id`**。
- MQ 投递是至少一次语义；notify 通过 inbox 和通道记录表保证业务结果幂等。
- 站内信是本地数据写入，可以依赖唯一约束安全重试。
- 外部触达允许多次尝试，但**同一事件、同一通道目标最多只能有一次成功结果**。
- 邮件、短信和 Lark webhook 在调用外部通道前先写入发送记录并领取发送权；已 `succeeded` 的记录不能再次调用外部通道。
- 外部通道支持幂等键时，使用发送记录 ID 作为幂等键。
- 邮件和短信在调用第三方通道前执行 Redis 精确滑动窗口限流；默认同一收件目标 5 分钟内最多发送 5 次，命中限流时通道记录标记为 `rate_limited`。
- notify 通过内部 gRPC 提供只读限流查询，用于其他服务在发起异步投递前同步判断用户是否应等待；该查询不占用发送额度，最终发送准入仍以 notify 投递前的限流判断为准。
- user 注册验证码发送前调用 notify 只读限流查询；命中限流时 user 返回 `too_many_requests`，BFF 透传给客户端；查询失败时 user 不写验证码缓存和 outbox，直接返回失败。
- 外部通道返回明确失败时，记录 `failed` 并不 ack，由 MQ 重新投递后再次尝试。
- 外部调用结果不确定时，记录 `unknown`；通道没有服务商幂等能力时不自动重试。
- 内部处理状态无法可靠落库时，consumer 返回错误并不 ack，由 MQ 重新投递。

## 实现思路

### 核心链路

```text
跨服务事件
-> NotifyConsumer 从 MQ 获取消息
-> 校验 event_id / event_type / subject / payload
-> 领取 notification_inbox_event 处理权
-> 查询 enabled notification_rule 和对应通道模板
-> 无匹配规则时标记 processed 并 ack
-> EventHandler 补齐业务上下文、模板变量和收件人
-> ChannelHandler 按通道生成记录并执行通道动作
-> NotifyConsumer 根据处理结果更新 inbox 状态并决定 ack
```

### 职责边界

| 抽象 | 职责 | 不承担的职责 |
| --- | --- | --- |
| `NotifyConsumer` | 消费 MQ、处理 inbox 幂等、查询规则和模板、编排处理流程、维护 inbox 状态、决定 ack | 理解具体事件语义、渲染通道模板、调用外部通道 |
| `EventHandler` | 根据事件 payload 查询业务上下文，生成模板变量和 `NotificationRecipient` 列表 | 判断通道是否启用、渲染通道模板、调用外部通道 |
| `ChannelHandler` | 按通道解析目标、渲染模板、幂等写入通道记录、执行站内信写入或外部触达 | 消费原始 MQ、维护 inbox 状态、查询事件业务对象 |

`EventHandler` 只处理事件语义。它可以查询文章、评论、用户关系等业务上下文，用于生成模板变量和收件人列表，但不按通道选择模板，也不拼装短信、邮件或 Lark 请求。

`ChannelHandler` 只处理通道差异。邮件和短信可以在 `NotificationRecipient.email`、`NotificationRecipient.phone` 为空时，通过 `user_id` 查询用户触达信息；这属于通道目标解析，不属于事件业务上下文补齐。

### 上下文结构

`EventHandler` 输出 `NotificationContext`：

| 字段 | 说明 |
| --- | --- |
| `event_id` | 来源事件 ID |
| `event_type` | 来源事件类型 |
| `language` | 模板语言，固定为 `zh_CN` |
| `vars` | 事件模板变量 |
| `recipients` | 收件人对象列表 |

模板渲染统一使用 `NotificationContext.vars`。接收者对象只表达触达目标，不携带模板变量。

### 通道处理

| 通道 | 处理方式 |
| --- | --- |
| `station` | 使用 `NotificationRecipient.user_id`，渲染标题和内容，写入 `notification_station_message` |
| `email` | 解析邮箱，渲染主题和正文，写入 `notification_email_delivery`，调用邮件客户端 |
| `tencent_sms` | 解析手机号，渲染有序短信参数，写入 `notification_tencent_sms_delivery`，调用腾讯云短信客户端 |
| `lark_webhook` | 使用模板中的 `webhook_id` 和 `token`，渲染 content 模板并组装请求体，写入 `notification_lark_webhook_delivery`，调用 Lark webhook |

站内信处理：

```text
station handler
-> 检查 receiver_id
-> 渲染 title_template / content_template
-> 按 event_id + receiver_id 写入 notification_station_message，状态为 succeeded
-> 已存在则返回已有状态
```

外部触达处理：

```text
external handler
-> 解析通道目标并渲染通道载荷
-> 按本通道唯一约束插入或读取发送记录
-> 记录已 succeeded 则直接返回 succeeded
-> 记录处于 processing 且未超时则返回 processing
-> 记录为 failed 时重新领取发送权
-> 记录为 unknown 时，有服务商幂等能力才重新领取发送权，否则返回 unknown
-> 邮件和短信执行发送限流并占用发送额度
-> 调用外部通道
-> 成功后标记 succeeded
-> 明确失败后标记 failed
-> 结果不确定时标记 unknown
```

领取发送权必须使用条件更新，只有未成功的发送记录才能从 `failed`、`unknown` 或已超时的 `processing` 更新为 `processing`。`succeeded` 是终态，不能被覆盖。

通道处理状态统一使用 `notification_channel_status`。`ChannelHandler` 返回该状态，通道记录的 `status` 字段复用其中可落库的状态，不为不同通道定义独立状态枚举。

| 状态 | 是否写入通道记录 | consumer 处理 | 含义 |
| --- | --- | --- | --- |
| `processing` | 是 | 标记 inbox failed，不 ack | 已领取通道处理权，可能已经发起外部调用；重复消费不能并发触发通道动作 |
| `succeeded` | 是 | 继续处理其他通道 | 站内信写入成功，或外部通道返回成功 |
| `skipped` | 否 | 继续处理其他通道 | 规则、模板、目标或渲染条件不满足，不生成通道记录 |
| `failed` | 是 | 标记 inbox failed，不 ack | 通道动作已经形成明确失败结果，可由 MQ 重新投递后再次尝试 |
| `unknown` | 是 | 继续处理其他通道 | 外部调用超时、中断或结果无法判断；没有服务商幂等能力时不自动重试 |
| `internal_error` | 否 | 标记 inbox failed，不 ack | notify 内部没有形成可靠通道结果，例如发送记录插入失败或发送记录状态更新失败 |
| `rate_limited` | 是 | 继续处理其他通道 | 邮件或短信命中发送频率限制，不调用第三方通道 |

### 事件消费状态

| inbox 状态 | 处理方式 |
| --- | --- |
| 无记录 | 插入 `processing` 并处理事件 |
| `processed` | 直接 ack 原始 MQ 消息 |
| `failed` | 更新为 `processing` 后重新处理 |
| `processing` 未超时 | 返回错误，不 ack |
| `processing` 已超时 | 更新处理开始时间后重新处理 |

事件处理结果：

```text
全部通道返回 succeeded / skipped / unknown
-> notification_inbox_event 标记 processed
-> ack MQ 消息

任一通道返回 processing / failed / internal_error
-> notification_inbox_event 标记 failed
-> 不 ack MQ 消息
```

## 表结构设计

### `notification_inbox_event`

| 字段 | 说明 |
| --- | --- |
| `id` | 主键 |
| `event_id` | 事件 ID，生产者生成的 UUID |
| `event_type` | 事件类型 |
| `subject` | MQ 主题 |
| `payload` | 原始事件 payload |
| `status` | `processing` / `processed` / `failed` |
| `attempt_count` | 处理尝试次数 |
| `last_error` | 最近一次失败原因摘要 |
| `processing_started_at` | 最近一次开始处理时间 |
| `processed_at` | 处理完成时间 |
| `created_at` / `updated_at` | 审计时间 |

唯一约束：

```text
event_id
```

### `notification_rule`

| 字段 | 说明 |
| --- | --- |
| `id` | 规则 ID |
| `event_type` | 事件类型 |
| `channel` | 通知通道 |
| `language` | 语言 |
| `enabled` | 启用状态 |
| `created_at` / `updated_at` | 审计时间 |

唯一约束：

```text
event_type + channel + language
```

### 通道模板表

通道模板表通过 `rule_id` 关联 `notification_rule.id`，每张模板表的 `rule_id` 唯一。

| 表 | 字段 |
| --- | --- |
| `notification_station_template` | `rule_id`、`title_template`、`content_template`、`created_at`、`updated_at` |
| `notification_email_template` | `rule_id`、`subject_template`、`body_template`、`content_type`、`created_at`、`updated_at` |
| `notification_tencent_sms_template` | `rule_id`、`sms_sdk_app_id`、`sign_name`、`provider_template_id`、`param_templates`、`created_at`、`updated_at` |
| `notification_lark_webhook_template` | `rule_id`、`webhook_id`、`token`、`secret`、`msg_type`、`content_template`、`created_at`、`updated_at` |

`webhook_id` 是 Lark webhook 配置标识，用于投递记录幂等约束。`token` 属于敏感字段，只用于拼接 Lark webhook URL，不返回给 BFF，不写入日志。`secret` 是可选签名密钥，只用于发送前计算 `sign`，不写入投递记录。

### `notification_station_message`

| 字段 | 说明 |
| --- | --- |
| `id` | 主键 |
| `event_id` | 来源事件 ID |
| `event_type` | 来源事件类型 |
| `receiver_id` | 接收者用户 ID |
| `title` | 标题 |
| `content` | 内容 |
| `status` | `succeeded` |
| `read_at` | 已读时间 |
| `created_at` / `updated_at` | 审计时间 |

唯一约束：

```text
event_id + receiver_id
```

查询索引：

```text
receiver_id + created_at
receiver_id + read_at
```

### `notification_email_delivery`

| 字段 | 说明 |
| --- | --- |
| `id` | 主键 |
| `event_id` | 来源事件 ID |
| `event_type` | 来源事件类型 |
| `receiver_id` | 接收者用户 ID，可为空 |
| `to_email` | 标准化收件邮箱 |
| `subject` | 邮件主题 |
| `body` | 邮件正文 |
| `content_type` | 正文类型 |
| `status` | `processing` / `succeeded` / `failed` / `unknown` / `rate_limited` |
| `attempt_count` | 发送尝试次数 |
| `last_attempt_at` | 最近一次发送尝试时间 |
| `provider_message_id` | 邮件服务返回的消息 ID，可为空 |
| `provider_response` | 邮件服务响应摘要，可为空 |
| `sent_at` | 发送成功时间 |
| `created_at` / `updated_at` | 审计时间 |

唯一约束：

```text
event_id + to_email
```

### `notification_tencent_sms_delivery`

| 字段 | 说明 |
| --- | --- |
| `id` | 主键 |
| `event_id` | 来源事件 ID |
| `event_type` | 来源事件类型 |
| `receiver_id` | 接收者用户 ID，可为空 |
| `phone` | E.164 格式手机号 |
| `sms_sdk_app_id` | 腾讯云短信应用 ID |
| `sign_name` | 短信签名 |
| `provider_template_id` | 腾讯云模板 ID |
| `template_params` | 有序模板参数 |
| `status` | `processing` / `succeeded` / `failed` / `unknown` / `rate_limited` |
| `attempt_count` | 发送尝试次数 |
| `last_attempt_at` | 最近一次发送尝试时间 |
| `provider_request_id` | 腾讯云请求 ID |
| `provider_code` | 腾讯云返回码 |
| `provider_message` | 腾讯云返回消息 |
| `sent_at` | 发送成功时间 |
| `created_at` / `updated_at` | 审计时间 |

唯一约束：

```text
event_id + phone
```

### `notification_lark_webhook_delivery`

| 字段 | 说明 |
| --- | --- |
| `id` | 主键 |
| `event_id` | 来源事件 ID |
| `event_type` | 来源事件类型 |
| `webhook_id` | Lark webhook 配置标识 |
| `request_body` | 未签名基础请求体 |
| `status` | `processing` / `succeeded` / `failed` / `unknown` |
| `attempt_count` | 发送尝试次数 |
| `last_attempt_at` | 最近一次发送尝试时间 |
| `http_status` | HTTP 状态码 |
| `response_body` | Lark 响应体 |
| `sent_at` | 发送成功时间 |
| `created_at` / `updated_at` | 审计时间 |

唯一约束：

```text
event_id + webhook_id
```

## BFF 暴露能力

BFF 只暴露当前登录用户相关能力：

- 查询站内信列表。
- 标记站内信已读。
- 查询未读数量。

BFF 不暴露通知规则管理、通道模板管理、外部投递记录管理和对象存储管理。

## 内部 gRPC 能力

- `NotifyRateLimitService.Check` 查询邮件或腾讯云短信收件目标在当前限流窗口内是否已被限流。
- 查询结果包含是否限流、建议等待秒数和剩余发送次数。
- 查询接口只读取 Redis 限流窗口，不写入发送记录，也不占用发送额度。

## 边界

- notify 不直接读取其他服务业务库，跨服务查询通过归属服务 RPC 完成。
- 事件 payload 保持最小化，通知展示字段和触达字段由 notify 按需查询。
- 用户偏好不参与通知主链路。
- 对象存储是 notify 中的独立能力，不参与通知主链路。
