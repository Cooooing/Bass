# Agent 项目提示

## 工作入口

- 先读 [doc/README.md](doc/README.md)，再按变更类型读架构或代码规范；长期细则只写 `doc/`，高频约束才同步 `AGENTS.md` 和 `CLAUDE.md`。
- 执行代码生成、重构或修复时，按 [doc/agent-rules.md](doc/agent-rules.md) 做前后检查；新增同类代码时先看 [doc/templates/](doc/templates/)。
- 当前模块归类以 [doc/README.md](doc/README.md) 为准；目标规划服务未落地前，不生成代码、目录、配置或契约。
- 生成、格式化、测试和构建优先使用 Makefile；运行 make 使用 bash。

## GitNexus

- 项目索引名是 **Bass**；探索、影响分析、重构和提交前检查优先使用 GitNexus。
- 修改函数、方法、类型前做上游影响分析；HIGH 或 CRITICAL 风险先说明再继续。
- 不熟悉流程时先查执行流程再读文件；重命名用 GitNexus rename 预览；提交前运行变更检测，索引过期时运行 `npx gitnexus analyze`，已有 embeddings 时加 `--embeddings`。

## 必守约束

- BFF 不连接业务库、不拥有领域数据；写接口通常只调用一个归属服务命令接口，读接口可以聚合内部服务；一个业务事实只由归属服务写入，跨服务副作用由归属服务写 outbox 后通过 MQ 触发。
- BFF 负责端侧访问边界和协议适配；内部服务负责领域事实边界和写入正确性。BFF 写接口不做 read-before-write 的领域状态预校验，资源存在性、作者关系、状态流转、本地业务不变量和并发写入正确性由归属服务校验。
- 禁止为了查询方便跨服务连接业务库或维护跨服务数据库投影；新增同步依赖前说明调用方向、失败影响和降级策略；缓存只用 Redis 并设置 TTL，不能写业务库或作为领域事实来源。
- 契约分为 BFF HTTP、内部 gRPC、内部事件和外部回调；BFF proto 不引用内部服务 message，内部服务 proto 不写 HTTP 注解；Proto 管理、lint、format 和生成统一使用 Buf，共享 API 使用 `make api`、`make api-lint`。
- Proto 枚举单独放 `enum.proto`，不写入 `model.proto`；`model.proto` 只放稳定数据模型，不放 Query、Save、回调参数或事件 payload；保存、更新等写入请求结构定义在对应 RPC message 内部，不抽 `XXXSave`、`XXXUpdate` 复用；BFF 返回模型按视图拆分，内部服务不引用其他业务服务展示 message。
- BFF HTTP 默认 `POST` 和 `body: "*"`；图片、文件、健康检查、回调等可以按场景使用 `GET`；RPC 使用业务动作命名，查询收敛为 `Get`、`List`、`Page`、`Map`；`List`、`Page` 响应数组字段统一为 `rows`，响应结构不能由 `with_xxx` 或 `include_xxx` 控制。
- BFF 响应体固定为 `code`、`message`、`data`、`time`，HTTP 使用真实状态码；`ErrorReason` 只表示传输状态，业务错误码使用 `BusinessErrorCode` 并通过 `errors.code` 绑定 HTTP 状态。
- 对外可感知业务错误使用 `common/pkg/apperror.New(BusinessErrorCode, ...)`；业务校验失败禁止使用 `cerrors.ErrorBadRequest`、`cerrors.ErrorNotFound` 等 `ErrorReason` 构造函数；用户可见文案由 BFF 按 `common.enums.Language` 生成；动态错误参数必须是 common proto message，并通过 `apperror.WithData[T proto.Message]` 写入。
- service 层只做协议适配和调用 usecase；biz 层不依赖 data、Ent 生成模型或具体数据库类型；内部服务仍需校验资源存在、状态流转、本地业务不变量和写入前置条件。
- 内部服务的业务输入只允许来自 proto request 显式字段；不得从 ctx、metadata、请求头、JWT、缓存会话等隐式来源提取当前用户或业务参数。ctx 只用于取消、超时、trace 和事务传递；本服务数据库领域事实、配置、当前时间、UUID、事件 ID 等可以参与业务判断。
- content 内容状态使用 `publish_status`、`visibility`、`restriction` 等独立字段：文章用 `publish_status` 表达草稿、已发布、归档，用 `visibility` 表达公开、私有，用 `restriction` 表达无管理限制、管理隐藏、管理锁定；评论和附言使用 `restriction` 表达管理限制；逻辑删除只使用 `deleted_at`，应用不提供查询已删除数据的入口，不保留 `deleted` 枚举状态和业务表删除人字段。BFF 控制端侧可见范围，content 校验状态流转和本地业务不变量。
- schema 字段必须有明确业务需求或读写路径；表结构、outbox/inbox 字段和对外契约变更必须先确认设计、影响和回滚；密码、token、密钥不返回、不记录日志、不写错误和事件 payload，验证码只允许写入验证码投递事件。
- 服务配置以 `internal/conf/conf.proto` 为契约来源，`configs/config.yaml` 必须覆盖该服务暴露的可配置字段，禁止保留 proto 未定义字段；`google.protobuf.Duration` 配置统一使用秒单位，禁止写 `ms`、`m`、`h` 等单位；布尔配置必须使用 `true` 或 `false` 字面量，禁止使用环境变量占位；运行期业务阈值、限流、告警和事件轮询参数可以热重载，端口、数据库、Redis、NATS、Consul、JWT secret、OSS 密钥、短信密钥等连接类或敏感配置不热重载。
- outbox/inbox 表结构同类保持一致；事件 payload 使用 common proto message；MQ subject 和 queue group 使用 common 枚举；outbox publisher 使用 Redis 批次锁降低扫表压力，获取锁失败时跳过本轮；outbox/inbox 超过最大重试后进入 dead，由本服务 dead-letter scanner 记录日志、指标和可选 Lark 告警。
- repo 层每个 schema 对应的 repo 都定义 `Get`、`List`、`Map`、`Count`、`Page` 五个查询方法；依赖命名按类型使用 `Repo`、`Cache`、`Client`、`Usecase`，不使用 `Manager`、`Resolver`、`Writer` 包装真实依赖；需要基础类型指针时使用 Go 1.26 的 `new(expr)`。
- 一次性简单逻辑、入参校验和字段组装不拆游离私有函数；不新增通用转换接口或 `toXXX`、`buildXXXReply` 等组装函数；默认不新增测试代码，仍需按变更类型运行生成、编译或已有测试。
- slice、map、set 的去重、映射、过滤、分组、键值提取等纯数据变换优先使用 `github.com/samber/lo`；有副作用、错误短路、事务操作或可读性下降时保留显式循环。
- content 已发布文章的作者编辑窗口由 content 配置 `business.article.published_edit_window` 和 `business.article.published_edit_max_view_count` 控制，默认 10 分钟且浏览数小于 100。
- 新增服务、事件、BFF 接口、数据库 schema、生成链路、公共包约定或配置项时，同步更新 `doc/`、`AGENTS.md` 和 `CLAUDE.md`。

<!-- gitnexus:start -->
# GitNexus — Code Intelligence

This project is indexed by GitNexus as **Bass** (13365 symbols, 27903 relationships, 300 execution flows). Use the GitNexus MCP tools to understand code, assess impact, and navigate safely.

> If any GitNexus tool warns the index is stale, run `npx gitnexus analyze` in terminal first.

## Always Do

- **MUST run impact analysis before editing any symbol.** Before modifying a function, class, or method, run `gitnexus_impact({target: "symbolName", direction: "upstream"})` and report the blast radius (direct callers, affected processes, risk level) to the user.
- **MUST run `gitnexus_detect_changes()` before committing** to verify your changes only affect expected symbols and execution flows.
- **MUST warn the user** if impact analysis returns HIGH or CRITICAL risk before proceeding with edits.
- When exploring unfamiliar code, use `gitnexus_query({query: "concept"})` to find execution flows instead of grepping. It returns process-grouped results ranked by relevance.
- When you need full context on a specific symbol — callers, callees, which execution flows it participates in — use `gitnexus_context({name: "symbolName"})`.

## When Debugging

1. `gitnexus_query({query: "<error or symptom>"})` — find execution flows related to the issue
2. `gitnexus_context({name: "<suspect function>"})` — see all callers, callees, and process participation
3. `READ gitnexus://repo/Bass/process/{processName}` — trace the full execution flow step by step
4. For regressions: `gitnexus_detect_changes({scope: "compare", base_ref: "main"})` — see what your branch changed

## When Refactoring

- **Renaming**: MUST use `gitnexus_rename({symbol_name: "old", new_name: "new", dry_run: true})` first. Review the preview — graph edits are safe, text_search edits need manual review. Then run with `dry_run: false`.
- **Extracting/Splitting**: MUST run `gitnexus_context({name: "target"})` to see all incoming/outgoing refs, then `gitnexus_impact({target: "target", direction: "upstream"})` to find all external callers before moving code.
- After any refactor: run `gitnexus_detect_changes({scope: "all"})` to verify only expected files changed.

## Never Do

- NEVER edit a function, class, or method without first running `gitnexus_impact` on it.
- NEVER ignore HIGH or CRITICAL risk warnings from impact analysis.
- NEVER rename symbols with find-and-replace — use `gitnexus_rename` which understands the call graph.
- NEVER commit changes without running `gitnexus_detect_changes()` to check affected scope.

## Tools Quick Reference

| Tool | When to use | Command |
|------|-------------|---------|
| `query` | Find code by concept | `gitnexus_query({query: "auth validation"})` |
| `context` | 360-degree view of one symbol | `gitnexus_context({name: "validateUser"})` |
| `impact` | Blast radius before editing | `gitnexus_impact({target: "X", direction: "upstream"})` |
| `detect_changes` | Pre-commit scope check | `gitnexus_detect_changes({scope: "staged"})` |
| `rename` | Safe multi-file rename | `gitnexus_rename({symbol_name: "old", new_name: "new", dry_run: true})` |
| `cypher` | Custom graph queries | `gitnexus_cypher({query: "MATCH ..."})` |

## Impact Risk Levels

| Depth | Meaning | Action |
|-------|---------|--------|
| d=1 | WILL BREAK — direct callers/importers | MUST update these |
| d=2 | LIKELY AFFECTED — indirect deps | Should test |
| d=3 | MAY NEED TESTING — transitive | Test if critical path |

## Resources

| Resource | Use for |
|----------|---------|
| `gitnexus://repo/Bass/context` | Codebase overview, check index freshness |
| `gitnexus://repo/Bass/clusters` | All functional areas |
| `gitnexus://repo/Bass/processes` | All execution flows |
| `gitnexus://repo/Bass/process/{name}` | Step-by-step execution trace |

## Self-Check Before Finishing

Before completing any code modification task, verify:
1. `gitnexus_impact` was run for all modified symbols
2. No HIGH/CRITICAL risk warnings were ignored
3. `gitnexus_detect_changes()` confirms changes match expected scope
4. All d=1 (WILL BREAK) dependents were updated

## Keeping the Index Fresh

After committing code changes, the GitNexus index becomes stale. Re-run analyze to update it:

```bash
npx gitnexus analyze
```

If the index previously included embeddings, preserve them by adding `--embeddings`:

```bash
npx gitnexus analyze --embeddings
```

To check whether embeddings exist, inspect `.gitnexus/meta.json` — the `stats.embeddings` field shows the count (0 means no embeddings). **Running analyze without `--embeddings` will delete any previously generated embeddings.**

> Claude Code users: A PostToolUse hook handles this automatically after `git commit` and `git merge`.

## CLI

| Task | Read this skill file |
|------|---------------------|
| Understand architecture / "How does X work?" | `.claude/skills/gitnexus/gitnexus-exploring/SKILL.md` |
| Blast radius / "What breaks if I change X?" | `.claude/skills/gitnexus/gitnexus-impact-analysis/SKILL.md` |
| Trace bugs / "Why is X failing?" | `.claude/skills/gitnexus/gitnexus-debugging/SKILL.md` |
| Rename / extract / split / refactor | `.claude/skills/gitnexus/gitnexus-refactoring/SKILL.md` |
| Tools, resources, schema reference | `.claude/skills/gitnexus/gitnexus-guide/SKILL.md` |
| Index, status, clean, wiki CLI commands | `.claude/skills/gitnexus/gitnexus-cli/SKILL.md` |

<!-- gitnexus:end -->
