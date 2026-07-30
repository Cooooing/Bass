# Content Interaction / View Audit 业务设计

## 1. Interaction 业务定位

互动系统记录用户对文章和评论的行为事实。content 只负责记录互动事实并发送事件，不直接处理积分副作用。

## 2. P0 互动范围

| 对象 | 互动 | 说明 |
| --- | --- | --- |
| 文章 | `like` | 点赞，表达认可 |
| 文章 | `collect` | 收藏，保存到用户个人收藏 |
| 文章 | `thank` | 感谢，固定积分奖励由积分服务消费事件处理 |
| 文章 | `reward` | 打赏，解锁打赏区内容，积分服务处理扣减 |
| 评论 | `like` | 评论点赞 |
| 评论 | `thank` | 评论感谢，固定积分奖励由积分服务消费事件处理 |

## 3. 不做互动

| 互动 | 原因 |
| --- | --- |
| 关注 `watch` | 没有明确更新语义，缺少闭环 |
| 评论收藏 | 场景弱 |
| 评论打赏 | 场景弱 |

## 4. 统一 Interaction 规则

| 规则 | 说明 |
| --- | --- |
| 唯一性 | `target_id + user_id + action` 唯一 |
| 幂等创建 | 已存在相同互动时返回成功，不重复增加计数 |
| 幂等取消 | 不存在互动时返回成功，不重复扣减计数 |
| 取消方式 | P0 建议删除 interaction 记录 |
| 计数维护 | 创建时目标计数 +1，取消时目标计数 -1 |
| 积分副作用 | content 不处理，发事件给积分服务 |
| 事件投递 | 只在真实创建互动时发送事件，重复创建不发 |

## 5. 文章打赏区

| 规则 | 说明 |
| --- | --- |
| 设置时机 | 作者只在草稿阶段设置打赏区内容和所需积分 |
| 发布后修改 | 不允许 |
| 未打赏用户 | 不返回打赏区内容，只返回是否存在打赏区 |
| 已打赏用户 | 返回打赏区内容 |
| 作者本人 | 可查看自己文章打赏区内容 |
| 管理员 | 可查看打赏区内容 |
| 打赏事实 | 通过文章 `reward` interaction 判断 |
| 积分处理 | 积分服务消费 `CONTENT_ARTICLE_REWARDED` 事件处理 |

## 6. 文章 Interaction 表

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | int64 | 主键 |
| `article_id` | int64 | 文章 ID |
| `user_id` | int64 | 用户 ID |
| `type` | enum | `like` / `collect` / `thank` / `reward` |
| `created_at` | time | 创建时间 |
| `updated_at` | time | 更新时间 |

## 7. 评论 Interaction 表

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | int64 | 主键 |
| `comment_id` | int64 | 评论 ID |
| `user_id` | int64 | 用户 ID |
| `type` | enum | `like` / `thank` |
| `created_at` | time | 创建时间 |
| `updated_at` | time | 更新时间 |

## 8. Interaction 事件

| 事件 | 触发 |
| --- | --- |
| `CONTENT_ARTICLE_LIKED` | 文章点赞真实创建 |
| `CONTENT_ARTICLE_COLLECTED` | 文章收藏真实创建 |
| `CONTENT_ARTICLE_THANKED` | 文章感谢真实创建 |
| `CONTENT_ARTICLE_REWARDED` | 文章打赏真实创建 |
| `CONTENT_COMMENT_LIKED` | 评论点赞真实创建 |
| `CONTENT_COMMENT_THANKED` | 评论感谢真实创建 |

## 9. View Audit 业务定位

浏览审计用于文章浏览量、防刷、登录态浏览历史和后续推荐特征。P0 必须实现。

## 10. P0 浏览功能

| 功能 | 说明 |
| --- | --- |
| 浏览文章增加浏览量 | Redis 去重通过后增加浏览量 |
| 浏览增量缓存 | Redis hash 记录文章浏览增量 |
| 浏览去重缓存 | Redis hash 记录文章浏览记录，24h 过期 |
| 登录态浏览历史 | 登录用户浏览文章时落 DB |
| 浏览历史查询 | 用户可查询自己的浏览历史 |
| 游客浏览计数 | 游客可贡献浏览量，但不写用户历史 |

## 11. Redis 设计

| Key | 类型 | 字段 | 值 | TTL |
| --- | --- | --- | --- | --- |
| `Content:ArticleViewCount` | Hash | `article_id` | 待刷新的浏览增量 | 不设置或按刷库策略设置 |
| `Content:ArticleViewRecord:{article_id}` | Hash | `viewer_key` | 浏览时间 | 24h |

## 12. Viewer Key 生成

| 场景 | viewer_key |
| --- | --- |
| 登录用户 | `user:{user_id}` |
| 游客优先 | `fp:{browser_fingerprint}` |
| 游客兜底 | `ipua:{ip}:{user_agent}` |

## 13. 浏览流程

1. 用户或游客访问文章。
2. content 校验文章可见性。
3. 根据登录态、浏览器指纹、IP、UA 生成 `viewer_key`。
4. 查询 `Content:ArticleViewRecord:{article_id}` 是否已有该 `viewer_key`。
5. 24h 内已存在时，不增加浏览量。
6. 不存在时，写入去重 hash。
7. 对 `Content:ArticleViewCount` 的 `article_id` 字段自增。
8. 登录用户写入或更新 `content_article_view_records`。
9. scheduler 或同步任务定期将浏览增量刷入 `content_articles.view_count`。

## 14. 浏览记录表

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | int64 | 主键 |
| `article_id` | int64 | 文章 ID |
| `user_id` | int64 | 登录用户 ID，游客为空 |
| `ip` | string | 浏览 IP |
| `user_agent` | string | 原始 UA |
| `browser_fingerprint` | string | 浏览器指纹 |
| `viewed_at` | time | 最近浏览时间 |
| `created_at` | time | 创建时间 |
| `updated_at` | time | 更新时间 |

## 15. 设计理由

- 点赞、收藏、感谢、打赏都可以作为行为事实记录在 interaction 表中。
- 积分副作用交给积分服务消费事件处理，避免 content 与积分服务同步耦合。
- 感谢和打赏不冲突：感谢是固定奖励，打赏是解锁打赏区内容。
- 浏览请求高频，使用 Redis 去重和增量缓存可以降低 DB 写压力。
- 登录态浏览历史是用户功能，必须落库；游客浏览只参与计数，不写历史。
