# 项目规范索引

这里放长期有效的设计与编码规范，正文尽量保持通用，不绑定具体模块名。

- [架构设计](architecture.md)
- [代码规范](coding.md)

`AGENTS.md` 和 `CLAUDE.md` 是 Agent 的项目级提示入口；本目录用于给人阅读和补充细节。

## 当前模块归类

当前模块归类用于说明已有目录或规划目录的目标归属，不代表目标服务已经落地。落地状态以“目标服务边界”表为准。

| 当前目录 | 目标边界 | 说明 |
| --- | --- | --- |
| `app/bbs` | `bbs` | 面向社区前台的 BFF，负责 HTTP/JSON 接口聚合与端侧适配。 |
| `app/gateway` | 入口治理 | 当前保留的通用 HTTP 网关模块；目标边界中不承载业务 BFF 逻辑。 |
| `app/user` | `user` | 身份与账户内部服务，拥有用户、认证凭证、关系、账户设置等数据。 |
| `app/content` | `content` | 社区内容内部服务，拥有文章、评论、标签、板块、互动记录等数据。 |
| `app/notify` | `notify` | 通知内部服务，负责通知模板、通知记录、通知偏好和投递状态。 |
| `app/im` | `im` | 即时通信内部服务，拥有会话、消息、已读状态、成员关系和投递状态。 |
| `common` | 共享基础模块 | 公共 proto、Buf 模板、客户端、工具和跨服务约定。 |

## 目标服务边界

当前已经落地的目标服务只有 `bbs`、`user`、`content`、`notify`。其他服务边界属于目标规划，不能因为文档中出现名称就直接生成对应代码、目录、配置或契约。

| 服务名 | 分层 | 状态 | 职责 |
| --- | --- | --- | --- |
| `bbs` | BFF | 已落地 | 前台页面/API 聚合。 |
| `bbs_admin` | BFF | 目标规划 | 后台管理/API 聚合。 |
| `integration` | External Edge | 目标规划 | 第三方回调接入、验签、幂等和协议转换。 |
| `openapi` | External Edge | 目标规划 | 对外开放 API。 |
| `push_node` | Push Edge | 目标规划 | 客户端下行 SSE 推送。 |
| `user` | Internal Service | 已落地 | 用户与认证。 |
| `content` | Internal Service | 已落地 | 社区内容与互动。 |
| `notify` | Internal Service | 已落地 | 通知业务。 |
| `im` | Internal Service | 目标规划 | 即时通信业务。 |
| `platform` | Internal Service | 目标规划 | 通用平台能力。 |
| `push_hub` | Internal Service | 目标规划 | 实时推送路由与节点控制。 |

目标服务边界的详细说明见根目录 [README.md](../README.md)。
