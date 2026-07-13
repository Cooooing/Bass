# 项目规范索引

本目录保存长期有效的项目规范。`AGENTS.md` 和 `CLAUDE.md` 只保留高频必守约束，完整细则以本目录为准。

## 规范入口

- 涉及服务边界、读写归属、跨服务依赖、契约、schema、事件、缓存、运行时和错误响应时，阅读 [架构设计](architecture.md)。
- 涉及局部抽象、命名、模型组装、生成、格式化、测试、迁移和文档同步时，阅读 [代码规范](coding.md)。
- 执行代码生成、重构或修复时，按 [Agent 可执行规则](agent-rules.md) 逐条检查；该文件只记录可被人工或脚本检查的硬规则。
- 新增 proto、service、usecase、repo 或 schema 时，先阅读 [代码模板](templates/README.md)，再参考同模块已有实现。
- 修改规范时，长期规则写入 `doc/`；高频必守约束同步到 `AGENTS.md` 和 `CLAUDE.md`。

## 规范分层

| 文件 | 作用 | 使用方式 |
|------|------|----------|
| `AGENTS.md` / `CLAUDE.md` | 高频硬约束和工作入口 | 每次任务开始前读取。 |
| `doc/architecture.md` | 架构边界和业务归属 | 涉及服务、契约、数据、事件、运行时设计时读取。 |
| `doc/coding.md` | 代码层规则和生成规则 | 涉及具体代码编写、命名、组装和测试时读取。 |
| `doc/agent-rules.md` | 可执行规则清单 | 代码修改前后按规则检查，后续 `rulecheck` 以此为来源。 |
| `doc/templates/` | 标准实现模板 | 新增同类代码时优先套用模板。 |

## 规则写法

新增或修改规范时，优先使用以下结构：

- 必须：明确需要遵守的行为。
- 禁止：明确不能出现的行为。
- 例外：说明允许偏离规则的场景。
- 检查：说明人工或脚本如何发现违规。

不能被执行或检查的说明，放入架构讨论文档，不写入 `agent-rules.md`。

## 当前模块归类

当前模块归类用于说明已有目录或规划目录的目标归属，不代表目标服务已经落地。落地状态以“目标服务边界”表为准。

| 当前目录          | 目标边界      | 说明                                  |
|---------------|-----------|-------------------------------------|
| `app/bbs`     | `bbs`     | 面向社区前台的 BFF，负责 HTTP/JSON 接口聚合与端侧适配。 |
| `app/user`    | `user`    | 身份与账户内部服务，拥有用户、认证凭证、关系和账户设置数据。      |
| `app/content` | `content` | 社区内容内部服务，拥有文章、评论、标签、板块和互动记录数据。      |
| `app/notify`  | `notify`  | 通知内部服务，负责通知模板、通知记录、通知偏好和投递状态。       |
| `app/im`      | `im`      | 即时通信内部服务，负责群组、会话、消息管理和实时投递。         |
| `app/push_hub` | `push_hub` | 推送中枢内部服务，管理推送节点注册、用户节点映射和消息路由。     |
| `app/push_node` | `push_node` | 推送边缘服务，维护客户端 SSE 长连接，接收并下发实时推送消息。   |
| `app/platform` | `platform` | 通用平台内部服务，提供 IP 解析和对象存储（OSS）能力。         |
| `app/scheduler` | `scheduler` | 分布式定时任务内部服务，维护任务定义、运行记录、调度触发和任务告警。 |
| `common`      | 共享基础模块    | 公共 proto、Buf 模板、客户端、工具和跨服务约定。       |

## 目标服务边界

当前已经落地的目标服务有 `bbs`、`user`、`content`、`notify`、`im`、`push_hub`、`push_node`、`platform`、`scheduler`。其他服务边界属于目标规划，不能因为文档中出现名称就直接生成对应代码、目录、配置或契约。

| 服务名           | 分层               | 状态   | 职责                  |
|---------------|------------------|------|---------------------|
| `bbs`         | BFF              | 已落地  | 前台页面/API 聚合。        |
| `bbs_admin`   | BFF              | 目标规划 | 后台管理/API 聚合。        |
| `openapi`     | External Edge    | 目标规划 | 对外开放 API。           |
| `push_node`   | Push Edge        | 已落地  | 客户端下行 SSE 推送。       |
| `user`        | Internal Service | 已落地  | 用户与认证。              |
| `content`     | Internal Service | 已落地  | 社区内容与互动。            |
| `notify`      | Internal Service | 已落地  | 通知业务。               |
| `im`          | Internal Service | 已落地  | 即时通信业务（群组、会话、消息）。 |
| `platform`    | Internal Service | 已落地  | 通用平台能力（IP 解析、对象存储）。 |
| `push_hub`    | Internal Service | 已落地  | 实时推送路由与节点控制。        |
| `scheduler`   | Internal Service | 已落地  | 分布式定时任务配置、触发、运行记录和告警。 |

目标服务边界的详细说明见根目录 [README.md](../README.md)。
