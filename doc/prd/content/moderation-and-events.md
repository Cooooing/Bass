# Content Moderation / Outbox 业务设计

## 1. 内容治理定位

内容治理用于管理端对文章、评论、附言进行处置，并保留完整审计记录。P0 不设计 `reason_code`，直接记录原因文本。

## 2. P0 治理功能

| 对象 | 操作 | 说明 |
| --- | --- | --- |
| 文章 | `hide` | 普通用户不可见 |
| 文章 | `unhide` | 取消隐藏 |
| 文章 | `lock` | 可见但不可评论/互动 |
| 文章 | `unlock` | 取消锁定 |
| 文章 | `archive` | 归档为只读 |
| 文章 | `unarchive` | 仅管理端取消归档 |
| 评论 | `hide` | 普通用户不可见 |
| 评论 | `unhide` | 取消隐藏 |
| 评论 | `lock` | 可见但不可回复/互动 |
| 评论 | `unlock` | 取消锁定 |
| 评论 | `delete` | 软删除 |
| 评论 | `restore` | 恢复软删除 |
| 附言 | `hide/unhide/lock/unlock` | P0 可预留 |

## 3. 治理状态规则

| 状态 | 普通用户可见 | 可评论/回复 | 可互动 |
| --- | --- | --- | --- |
| `none` | 是 | 是 | 是 |
| `hidden` | 否 | 否 | 否 |
| `locked` | 是 | 否 | 否 |
| `deleted` | 否或占位 | 否 | 否 |

## 4. 治理记录字段

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | int64 | 主键 |
| `target_type` | enum | `article` / `comment` / `postscript` |
| `target_id` | int64 | 目标 ID |
| `action` | enum | 治理动作 |
| `before_status` | string | 操作前状态，可记录组合状态文本 |
| `after_status` | string | 操作后状态，可记录组合状态文本 |
| `reason` | string | 原因说明 |
| `operator_id` | int64 | 操作人 |
| `operator_realm` | enum/string | 操作入口，例如 bbs_admin |
| `created_at` | time | 创建时间 |
| `updated_at` | time | 更新时间 |

## 5. 治理流程

1. 管理端调用治理接口。
2. service 校验参数、枚举、目标 ID。
3. biz 查询目标当前状态。
4. biz 校验状态迁移是否合法。
5. 事务内更新目标状态。
6. 事务内写入 moderation record，记录 before/after/reason/operator。
7. 事务内写入 outbox event。
8. 事务提交后立即 best-effort publish。
9. 发布失败不影响治理结果，scheduler 后续补偿。

## 6. 文章治理规则

| 操作 | 前置状态 | 结果 |
| --- | --- | --- |
| hide | 非 hidden | `restriction=hidden` |
| unhide | hidden | `restriction=none` |
| lock | none | `restriction=locked` |
| unlock | locked | `restriction=none` |
| archive | published | `publish_status=archived` |
| unarchive | archived | `publish_status=published`，仅管理端 |

## 7. 评论治理规则

| 操作 | 前置状态 | 结果 |
| --- | --- | --- |
| hide | 非 hidden | `restriction=hidden` |
| unhide | hidden | `restriction=none` |
| lock | none | `restriction=locked` |
| unlock | locked | `restriction=none` |
| delete | 未删除 | 设置 `deleted_at` |
| restore | 已删除 | 清空 `deleted_at` |

## 8. 设计理由

- 不使用 `reason_code`，避免 P0 引入原因字典、运营配置和版本兼容问题。
- before/after 比单纯记录 action 更适合审计和回溯。
- 管理端取消归档用于纠错，BBS 不开放，避免作者反复恢复归档主题。
- 治理不直接删除互动记录，展示层按目标状态过滤，必要时通过计数修复任务校准。

## 9. Outbox 定位

Outbox 用于将 content 的业务事实可靠投递给通知、积分、搜索、推荐等下游模块。

## 10. P0 事件

| 事件 | 触发 |
| --- | --- |
| `CONTENT_ARTICLE_PUBLISHED` | 文章发布 |
| `CONTENT_ARTICLE_ARCHIVED` | 文章归档 |
| `CONTENT_ARTICLE_STATUS_UPDATED` | 文章治理状态变化 |
| `CONTENT_ARTICLE_TAG_BOUND` | 文章绑定 tag |
| `CONTENT_ARTICLE_TAG_UNBOUND` | 文章解绑 tag |
| `CONTENT_COMMENT_PUBLISHED` | 评论发布 |
| `CONTENT_COMMENT_STATUS_UPDATED` | 评论治理状态变化 |
| `CONTENT_ARTICLE_LIKED` | 文章点赞 |
| `CONTENT_ARTICLE_COLLECTED` | 文章收藏 |
| `CONTENT_ARTICLE_THANKED` | 文章感谢 |
| `CONTENT_ARTICLE_REWARDED` | 文章打赏 |
| `CONTENT_COMMENT_LIKED` | 评论点赞 |
| `CONTENT_COMMENT_THANKED` | 评论感谢 |
| `CONTENT_ARTICLE_POSTSCRIPT_ADDED` | 添加附言 |

## 11. 事件 Payload 原则

| 原则 | 说明 |
| --- | --- |
| 只放事实 ID | 放 article_id、comment_id、tag_id、user_id 等 |
| 不放展示模型 | 不放作者昵称、头像、渲染内容 |
| 不放敏感内容 | 打赏区内容、完整正文等不进入事件 |
| 可幂等 | 每个事件有唯一 `event_id` |
| 可追溯 | payload 包含发生时间、操作者、目标 ID |

## 12. 投递流程

1. 业务事务内保存业务数据。
2. 业务事务内保存 outbox event。
3. 事务提交成功后，调用 `Publish(event_id/id)` 即时投递。
4. 即时投递失败只记录 warn，不影响业务接口成功。
5. scheduler 定时调用 `PublishBatch` 补偿。
6. 消费者按 `event_id` 做幂等。

## 13. 设计理由

- content 不同步调用积分、通知、搜索，避免被下游稳定性拖垮。
- 事务 outbox 能保证业务事实和事件最终一致。
- at-least-once 足够，exactly-once 成本过高，消费者幂等更现实。
- 即时投递提高通知和积分处理实时性，scheduler 只做补漏。
