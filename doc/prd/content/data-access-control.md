# Content 数据权限设计

## 1. 目标

content 只负责内容资源的数据范围和状态规则，不负责角色与权限点校验。调用方完成认证和 RBAC 校验后，将可信访问主体传入 content。

## 2. 内容可见字段

| 对象 | 作者可见范围 | 审核管理限制 |
| --- | --- | --- |
| 文章 | `visibility` | `restriction` |
| 评论 | 无 | `restriction` |
| 附言 | 无 | `restriction` |

`visibility` 仅由文章作者设置，表示文章公开或私有的原始意图。`restriction` 仅由审核和治理流程维护，表示 `none`、`locked`、`hidden`。审核不修改文章、评论和附言的内容。

## 3. 访问主体

| 主体 | `ContentAccess.scope` | `actor_user_id` | 数据范围 |
| --- | --- | --- | --- |
| 游客 | `guest` | 无 | 已发布、公开、未隐藏的内容 |
| 登录用户 | `user` | 必填 | 公开内容和互动、评论等普通用户操作 |
| 作者操作 | `author` | 必填 | 本人文章的草稿、定时发布、私有内容和作者侧写操作 |
| 管理端 | `admin` | 必填 | 所有未删除内容 |
| 内部任务 | `internal_task` | 无 | 任务所需的全部内容 |

`author` 是 content 的数据访问范围，不是 RBAC 角色。调用方只能在作者操作或本人文章查询时设置该范围，content 始终以 `actor_user_id == article.created_by` 做最终归属校验。

## 4. 查询范围

文章列表由 `ArticleAccessUsecase` 生成 `ArticleScopeFilter`，repo 只执行该范围对应的数据库条件。

| 查询场景 | 条件 |
| --- | --- |
| 游客或普通浏览 | `published`、`visibility=public`、`restriction in (none, locked)` |
| 本人文章列表 | `created_by=actor_user_id`，可查询本人所有未删除文章 |
| 管理端或内部任务 | 不附加公开范围，由业务查询条件决定 |

调用方仅在查询条件中的 `author_id` 等于当前账号时使用 `author` 范围。content repo 仍强制追加 `created_by=actor_user_id`，不能借此读取其他作者的数据。

评论和附言没有作者自定义的可见范围。普通用户查询时仅返回 `restriction in (none, locked)`，并且评论所属文章必须处于公开可读状态。

## 5. 单资源读取

文章模型的 `CanView` 是单篇文章的最终裁决：

| 文章状态 | 游客/普通用户 | 作者本人 | 管理端/内部任务 |
| --- | --- | --- | --- |
| 已发布且公开，`none/locked` | 可读 | 可读 | 可读 |
| 已发布且私有，`none/locked` | 不可读 | 可读 | 可读 |
| 草稿或定时发布，`none/locked` | 不可读 | 可读 | 可读 |
| `hidden` | 不可读 | 不可读 | 可读 |

附言读取先校验所属文章可读，再按附言 `restriction` 过滤。评论列表通过文章关联条件和评论 `restriction` 同时过滤。

## 6. 写操作

| 操作 | content 内部判断 |
| --- | --- |
| 创建、编辑草稿、发布、定时发布、取消定时、归档 | 主体为 `user` 且文章作者为当前用户 |
| 文章绑定 tag、添加附言 | 主体为 `user` 且文章作者为当前用户 |
| 互动、评论 | 主体为 `user`，并校验文章和目标内容当前允许互动 |
| 审核治理 | 主体为 `admin`，RBAC 是否允许由调用方在进入 content 前校验 |
| scheduler 发布 | 主体为 `internal_task`，仅允许到期定时文章 |

文章、评论、附言创建后内容不可修改。审核不通过或治理限制只变更 `restriction`，不变更原内容。

## 7. 信任边界

`ContentAccess` 只能由服务端 RPC 适配层构造，不能由 BFF HTTP 请求直接透传。BBS 对普通用户一律传 `user`，管理端在完成 user RBAC 校验后才可传 `admin`，scheduler 才可传 `internal_task`。

后续接入 RBAC 时，权限点与资源范围分层：RBAC 决定调用方是否可执行管理操作，content 负责文章、评论、附言的归属、状态和查询范围。
