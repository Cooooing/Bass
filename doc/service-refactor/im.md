# im 设计说明

## 业务需求

**im 是内部即时通信服务**，负责消息发送、会话管理、消息持久化、已读回执、历史消息、群组成员关系、消息审核编排和可投递事件生成。

im 只拥有即时通信业务事实。实时连接、节点路由和客户端下行分发归属 `push_hub` 与 `push_node`，im 不直接维护 SSE 连接，也不直接管理外部推送节点。

### 写入入口

- 写入入口以 `bbs` 调用 im 命令接口为主。
- 对外开放能力需要通过 `openapi` 接入时，`openapi` 只做协议、签名、限流和授权适配，消息业务事实仍由 im 写入。
- 发送消息、撤回消息、创建群组、解散群组、修改会话设置和标记已读都由 im 完成本地事务。
- 写入成功后，im 通过 outbox 生成跨服务事件，由实时分发链路异步消费。
- im 不直接读取 `user`、`content` 或其他服务的业务库；用户存在性、关系校验和外部审核能力通过归属服务 RPC 完成。

### 会话与消息模型

会话是用户视角下的消息入口，消息是即时通信的核心业务事实。

| 模型 | 归属事实 | 说明 |
| --- | --- | --- |
| 私聊会话 | 用户与另一用户的会话视图 | 每个参与用户维护自己的置顶、免打扰和已读进度。 |
| 群聊会话 | 用户与群组的会话视图 | 群组维护共享消息进度，成员会话维护个人设置和已读进度。 |
| 消息 | 一条私聊或群聊消息 | 消息写入后不可物理删除，撤回只变更消息状态。 |
| 群组 | 群聊的业务容器 | 维护群名称、头像、简介、群主、成员数量和群状态。 |
| 群成员 | 用户在群组中的身份 | 维护角色、群内昵称、禁言状态和成员状态。 |
| 投递记录 | 消息面向接收者的投递事实 | 记录消息进入实时分发链路后的投递状态，不记录连接级状态。 |

私聊和群聊都使用同一套消息表。接收目标只能是单个用户或单个群组，不能同时存在两个目标。

### 事件输出

im 输出事件以**消息和会话事实变化**为主。

- 事件消息使用 `common.enums.Event`。
- 事件主题使用 common 枚举映射，不直接书写 MQ subject 字符串。
- 每条事件必须携带 im 生成的 UUID：`event_id`。
- **事件类型、事件主题和 payload 必须匹配**。
- **事件 payload 只包含分发和查询所需的最小标识字段**，不携带用户资料、关系详情、连接节点信息或外部触达字段。

| 事件 | payload |
| --- | --- |
| 消息创建 | 消息 ID、发送者 ID、接收者类型、接收用户 ID 或群组 ID |
| 消息撤回 | 消息 ID、发送者 ID、接收者类型、接收用户 ID 或群组 ID |
| 会话已读 | 会话 ID、用户 ID、最后已读消息 ID |
| 群组创建 | 群组 ID、群主 ID |
| 群组解散 | 群组 ID、操作人 ID |
| 群成员变更 | 群组 ID、成员用户 ID、变更动作 |

消息正文是否进入分发事件由事件契约单独确定。默认事件只携带消息 ID，实时分发链路需要展示字段时通过 im RPC 查询。

### 幂等与投递语义

- 发送消息必须支持客户端幂等键；同一发送者、同一客户端幂等键最多生成一条消息。
- MQ 投递是至少一次语义；im 通过 outbox 事件和投递记录保证业务结果幂等。
- 撤回消息按消息 ID 幂等处理；已撤回消息再次撤回直接返回当前状态。
- 标记已读只能推进 `last_read_message_id` 或 `read_count`，不能回退。
- 私聊消息写入时同步维护发送者和接收者会话。
- 群聊消息写入时维护群组消息进度，成员会话只维护个人设置和已读进度，避免按群成员数量同步扩散写入。
- im 只记录消息进入实时分发链路的投递状态；具体连接是否在线、消息是否已经写入某个 SSE 连接，由 `push_hub` 和 `push_node` 负责。

## 实现思路

### 核心链路

```text
BFF / openapi 调用 im 命令接口
-> im service 校验协议入参和调用身份
-> ChatMessageUsecase 校验发送者、接收目标、群成员关系和消息类型
-> 必要时调用审核能力完成消息审核编排
-> 写入 chat_message
-> 更新 chat_session 或 chat_group 的消息进度
-> 写入 chat_message_delivery 和 outbox_event
-> 事务提交
-> OutboxPublisher 投递 MQ 事件
-> push_hub 消费 im 事件并分发到 push_node
-> 客户端通过 BFF 查询会话和历史消息
```

标记已读链路：

```text
BFF 调用 MarkRead
-> im service 校验会话归属
-> ChatSessionUsecase 按会话和最后消息推进已读进度
-> 写入 outbox_event
-> push_hub 消费已读事件并分发给相关在线端
```

### 职责边界

| 抽象 | 职责 | 不承担的职责 |
| --- | --- | --- |
| `IMChatMessageService` | 协议入口适配、入参校验、调用身份传递 | 直接访问 data 层、维护实时连接、拼装分发事件 |
| `ChatMessageUsecase` | 发送消息、撤回消息、消息审核编排、消息事务写入、生成 outbox 事件 | 处理 SSE 连接、管理节点路由、直接读取其他服务业务库 |
| `ChatSessionUsecase` | 查询会话、置顶、免打扰、标记已读、维护个人会话进度 | 维护群成员角色、生成消息正文、管理推送节点 |
| `ChatGroupUsecase` | 创建群组、解散群组、维护群成员关系和群内权限 | 处理客户端连接、保存用户账号资料 |
| `OutboxPublisher` | 读取 im outbox 并投递 MQ | 解释业务事件语义、决定目标连接节点 |
| `push_hub` | 消费 im 事件、维护在线路由、选择 `push_node` 分发 | 持久化聊天消息、修改会话已读状态 |

`ChatMessageUsecase` 只处理消息业务规则。用户资料、用户关系和审核能力通过对应归属服务 RPC 获取；im 不缓存这些事实作为权威数据。

`push_hub` 只处理实时分发控制。它可以通过 im RPC 查询消息和目标信息，但不能直接修改消息、会话、群组或投递记录。

### 发送上下文

发送消息时，业务层使用显式上下文表达当前命令：

| 字段 | 说明 |
| --- | --- |
| `sender_id` | 发送者用户 ID |
| `receiver_type` | 接收者类型：用户或群组 |
| `receiver_user_id` | 私聊接收用户 ID |
| `receiver_group_id` | 群聊接收群组 ID |
| `message_type` | 消息类型 |
| `content` | 消息正文或消息载荷 |
| `client_message_id` | 客户端幂等键 |
| `trace_id` | 链路追踪标识 |

`receiver_user_id` 和 `receiver_group_id` 只能按 `receiver_type` 二选一。群聊发送必须校验发送者仍是有效群成员，且未处于禁言状态。

### 会话处理

私聊消息处理：

```text
private message
-> 校验 receiver_user_id
-> 写入 chat_message
-> 更新发送者会话 message_count / last_message_id
-> 更新接收者会话 message_count / last_message_id / unread_count
-> 写入接收者 chat_message_delivery
-> 写入 outbox_event
```

群聊消息处理：

```text
group message
-> 校验 group_id 和发送者成员状态
-> 写入 chat_message
-> 更新 chat_group message_count / last_message_id
-> 写入群消息投递记录或可推导的分发事件
-> 写入 outbox_event
```

群聊未读数默认通过 `chat_group.message_count - chat_session.read_count` 计算。需要缓存未读数时，缓存只作为读取优化，不作为消息或已读事实来源。

撤回消息处理：

```text
revoke message
-> 校验消息存在和操作权限
-> 消息状态从 normal 更新为 revoked
-> 写入 outbox_event
-> 已撤回则直接返回当前状态
```

### 投递处理

im 生成的投递记录描述消息面向业务接收者的投递状态。

| 状态 | 含义 | 处理方式 |
| --- | --- | --- |
| `pending` | 已生成投递目标，尚未投递到 MQ | 等待 outbox 投递 |
| `dispatched` | 已生成分发事件或已经交给实时分发链路 | 不重复生成相同投递事件 |
| `failed` | 分发事件投递失败或投递状态无法可靠保存 | 由 outbox 重试或人工排查 |
| `unknown` | 下游状态无法判断 | 不伪造成成功，保留记录供排查 |

投递状态不等同于客户端已读状态。已读只由 `chat_session` 的读进度表达。

## 表结构设计

### `chat_session`

| 字段 | 说明 |
| --- | --- |
| `id` | 主键 |
| `owner_id` | 会话所属用户 ID |
| `session_type` | 会话类型：私聊或群聊 |
| `peer_user_id` | 私聊对端用户 ID，可为空 |
| `group_id` | 群组 ID，可为空 |
| `is_muted` | 是否免打扰 |
| `is_pinned` | 是否置顶 |
| `last_read_message_id` | 最后已读消息 ID |
| `read_count` | 已读消息数 |
| `unread_count` | 私聊未读数缓存 |
| `message_count` | 私聊消息总数缓存 |
| `last_message_id` | 私聊最后一条消息 ID |
| `created_at` / `updated_at` | 审计时间 |

唯一约束：

```text
owner_id + peer_user_id
owner_id + group_id
```

查询索引：

```text
owner_id + is_pinned + updated_at
owner_id + updated_at
```

### `chat_message`

| 字段 | 说明 |
| --- | --- |
| `id` | 主键 |
| `sender_id` | 发送者用户 ID |
| `receiver_user_id` | 私聊接收用户 ID，可为空 |
| `receiver_group_id` | 群聊接收群组 ID，可为空 |
| `message_type` | 消息类型 |
| `content` | 消息正文或消息载荷 |
| `status` | `normal` / `revoked` |
| `client_message_id` | 客户端幂等键 |
| `audit_status` | 审核状态 |
| `audit_reason` | 审核失败原因摘要，可为空 |
| `created_at` / `updated_at` | 审计时间 |

唯一约束：

```text
sender_id + client_message_id
```

查询索引：

```text
receiver_user_id + id
receiver_group_id + id
sender_id + id
```

### `chat_group`

| 字段 | 说明 |
| --- | --- |
| `id` | 主键 |
| `name` | 群名称 |
| `avatar` | 群头像，可为空 |
| `introduction` | 群简介，可为空 |
| `owner_id` | 群主用户 ID |
| `status` | `normal` / `dismissed` |
| `member_count` | 群成员数 |
| `message_count` | 群消息数 |
| `last_message_id` | 最后一条消息 ID |
| `created_at` / `updated_at` | 审计时间 |

查询索引：

```text
owner_id
status + updated_at
```

### `chat_group_member`

| 字段 | 说明 |
| --- | --- |
| `id` | 主键 |
| `group_id` | 群组 ID |
| `user_id` | 成员用户 ID |
| `nickname` | 群内昵称，可为空 |
| `role` | `member` / `admin` / `owner` |
| `status` | `normal` / `removed` |
| `mute_end_at` | 禁言结束时间，可为空 |
| `created_at` / `updated_at` | 审计时间 |

唯一约束：

```text
group_id + user_id
```

查询索引：

```text
user_id + status
group_id + status
```

### `chat_message_delivery`

| 字段 | 说明 |
| --- | --- |
| `id` | 主键 |
| `event_id` | 投递事件 ID |
| `message_id` | 消息 ID |
| `receiver_type` | 接收者类型：用户或群组 |
| `receiver_user_id` | 私聊接收用户 ID，可为空 |
| `receiver_group_id` | 群聊接收群组 ID，可为空 |
| `status` | `pending` / `dispatched` / `failed` / `unknown` |
| `attempt_count` | 投递尝试次数 |
| `last_error` | 最近一次失败原因摘要 |
| `dispatched_at` | 交给实时分发链路的时间 |
| `created_at` / `updated_at` | 审计时间 |

唯一约束：

```text
message_id + receiver_type + receiver_user_id
message_id + receiver_type + receiver_group_id
```

### `outbox_event`

im 使用项目通用 outbox 表结构，不为即时通信单独扩展投递字段。

payload 只保存事件最小标识字段。需要增加 outbox 字段时，必须同步评估所有模块的同类 outbox 表。

## BFF 暴露能力

BFF 只暴露当前登录用户相关能力：

- 发送私聊消息和群聊消息。
- 撤回本人有权限撤回的消息。
- 查询会话列表。
- 查询会话历史消息。
- 标记会话已读。
- 设置会话置顶和免打扰。
- 创建群组、解散群组和查询群组信息。
- 管理本人有权限管理的群成员。

BFF 不暴露实时节点路由、连接状态、内部投递记录、消息审核明细和 outbox 状态。

## 边界

- im 不维护 SSE 长连接，不直接向客户端推送数据。
- im 不管理 `push_node` 节点注册、在线路由和连接级投递确认。
- im 不直接读取其他服务业务库；跨服务查询通过归属服务 RPC 完成。
- im 不负责通知业务；聊天消息形成通知类提醒时，通过事件交给 `notify` 或实时分发链路处理。
- im 只编排消息审核流程；外部审核适配和通用平台能力归属 `platform`。
- 消息正文、审核结果和投递错误摘要不能写入不受控日志；日志只记录 trace ID、消息 ID、会话 ID、业务状态和失败摘要。
