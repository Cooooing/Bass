# Content Domain / Tag 业务设计

## 1. 业务定位

`domain` 是文章大分类，`tag` 是具体节点。它们用于首页导航、节点页筛选、发帖分类选择和内容归档。

| 对象 | 定义 | 关系 |
| --- | --- | --- |
| domain | 大分类，例如技术、生活、游戏 | 一个 domain 包含多个 tag |
| tag | 具体节点，例如 Go、数据库、分享创造 | 一个 tag 只属于一个 domain |
| article | 文章主题 | 一篇文章可绑定多个 tag |
| article_tags | 文章与 tag 的多对多关系 | 由 Ent edge 自动创建和维护 |

## 2. P0 功能

| 功能 | 使用方 | 说明 |
| --- | --- | --- |
| 创建 domain | 管理端 | 创建大分类 |
| 更新 domain | 管理端 | 修改名称、描述、图标、排序、启停、导航展示 |
| 查询 domain | BBS、管理端 | 首页导航、分类页、后台管理 |
| 创建 tag | 管理端 | 创建具体节点，必须归属一个 domain |
| 更新 tag | 管理端 | 修改名称、描述、图标、排序、启停 |
| 查询 tag | BBS、管理端 | 发帖选择、节点页、后台管理 |
| 绑定文章 tag | BBS、管理端 | 单独接口绑定文章与 tag |
| 解绑文章 tag | BBS、管理端 | 单独接口解绑文章与 tag |
| 查询文章 tag | BBS、管理端 | 根据文章 ID 查询绑定 tag |
| 按分类查文章 | BBS、管理端 | 支持按 `domain_id` 或 `tag_id` 查询 |
| 同步维护文章数 | content | 绑定/解绑时同步维护 `tag.article_count` |

## 3. Domain 字段

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `id` | int64 | 是 | 主键 |
| `code` | string | 是 | 稳定编码，用于 URL、缓存、运营配置 |
| `name` | string | 是 | 分类名称 |
| `description` | string | 否 | 分类描述 |
| `icon` | string | 否 | 分类图标 |
| `sort` | int32 | 是 | 展示排序，数值越小越靠前 |
| `status` | enum | 是 | `enabled` / `disabled` |
| `is_nav` | bool | 是 | 是否在首页导航展示 |
| `created_by` | int64 | 否 | 创建人 |
| `updated_by` | int64 | 否 | 更新人 |
| `created_at` | time | 是 | 创建时间 |
| `updated_at` | time | 是 | 更新时间 |
| `deleted_at` | time | 否 | 软删除时间 |

## 4. Tag 字段

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `id` | int64 | 是 | 主键 |
| `domain_id` | int64 | 是 | 所属 domain |
| `code` | string | 是 | 稳定编码，用于节点 URL、缓存、运营配置 |
| `name` | string | 是 | 节点名称 |
| `description` | string | 否 | 节点描述 |
| `icon` | string | 否 | 节点图标 |
| `sort` | int32 | 是 | 展示排序 |
| `status` | enum | 是 | `enabled` / `disabled` |
| `article_count` | int32 | 是 | 已绑定文章数，P0 同步维护 |
| `created_by` | int64 | 否 | 创建人 |
| `updated_by` | int64 | 否 | 更新人 |
| `created_at` | time | 是 | 创建时间 |
| `updated_at` | time | 是 | 更新时间 |
| `deleted_at` | time | 否 | 软删除时间 |

## 5. 唯一性与索引

| 约束/索引 | 说明 |
| --- | --- |
| `domain.code` 唯一 | 只对未软删除记录生效 |
| `domain.name` 唯一 | 只对未软删除记录生效，避免同名大分类 |
| `tag.domain_id + tag.code` 唯一 | 同一 domain 下节点编码唯一 |
| `tag.domain_id + tag.name` 唯一 | 同一 domain 下节点名称唯一 |
| `domain.status + is_nav + sort` | 首页导航查询 |
| `tag.domain_id + status + sort` | 分类下 tag 列表查询 |

## 6. 文章 Tag 绑定规则

| 场景 | 规则 |
| --- | --- |
| 作者绑定 tag | 仅允许绑定自己的 `draft` 文章 |
| 作者解绑 tag | 仅允许解绑自己的 `draft` 文章 |
| 管理端绑定 tag | 可绑定任意文章，用于纠错和治理 |
| 管理端解绑 tag | 可解绑任意文章，用于纠错和治理 |
| 发布后作者修改 tag | 不允许 |
| tag 状态 | 只能绑定 `enabled` tag |
| tag 数量 | P0 默认最多 5 个 |
| 重复绑定 | 幂等成功，不重复增加 `article_count` |
| 重复解绑 | 幂等成功，不重复减少 `article_count` |

## 7. Article Count 维护

| 操作 | 计数变化 |
| --- | --- |
| 成功绑定新 tag | 对应 tag `article_count + 1` |
| 成功解绑已有 tag | 对应 tag `article_count - 1` |
| 重复绑定 | 不变化 |
| 重复解绑 | 不变化 |
| 草稿丢弃 | 解绑全部 tag 并扣减计数 |
| 文章软删除 | 解绑全部 tag 或扣减绑定 tag 计数 |
| 管理端纠错 | 按实际绑定变化同步维护 |

## 8. 设计理由

- tag 绑定独立成接口，文章草稿接口更纯粹，避免编辑正文时混入分类副作用。
- 发布后作者不能修改 tag，避免文章发布后频繁改变节点归属影响分类秩序。
- 管理端保留绑定/解绑能力，方便运营纠错和内容治理。
- `content_article_tags` 由 Ent edge 自动创建和维护，避免手写中间表导致模型重复。
- `article_count` P0 同步维护，首页和节点页读取简单，暂时接受写入成本。
- `code` 与 `sort` 是分类导航的基础字段，提前明确可以避免后续 URL 和运营配置迁移。
