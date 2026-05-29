# Claude 项目提示

## 工作方式

- 优先阅读并遵守 [doc/README.md](doc/README.md) 中的项目规范。
- 修改规范时，同步维护 `AGENTS.md`、`CLAUDE.md` 和对应 `doc/*.md`。
- 当前模块归类只写在 [doc/README.md](doc/README.md)，通用规范正文不要硬编码现有模块名。
- 目标规划服务以 [doc/README.md](doc/README.md) 的落地状态为准，不能因为文档中出现服务名就直接生成代码、目录、配置或契约。
- 生成代码、测试和构建命令优先阅读 Makefile；运行 make 时使用 bash 环境。

## GitNexus

本项目在 GitNexus 中的索引名是 **Bass**。探索代码、评估影响、重构和提交前检查时优先使用 GitNexus。

- 修改函数、方法、类型等代码符号前，先做上游影响分析。
- GitNexus 返回 HIGH 或 CRITICAL 风险时，先向用户说明风险再继续。
- 不熟悉代码时，先查执行流程，再读具体文件。
- 重命名前先用 GitNexus rename 预览，不直接全文替换。
- 提交前运行变更检测，确认影响范围符合预期。

索引过期时运行：

```bash
npx gitnexus analyze
```

如需保留 embeddings：

```bash
npx gitnexus analyze --embeddings
```

## 规范入口

- [架构设计](doc/architecture.md)
- [代码规范](doc/coding.md)

## 必守约束

- BFF 写接口通常只调用一个归属服务命令接口；跨服务副作用由归属服务写 outbox 后通过 MQ 触发。
- 普通数据库写入失败直接返回失败，不为普通写流程额外设计重试；需要重试、补偿和 MQ 投递保证的流程放在 outbox/inbox。
- BFF 读接口可以聚合多个内部服务；BFF 不直接连接业务库，不拥有领域数据。
- 一个业务事实只能有一个归属服务负责写入；非归属服务读取业务事实默认通过 RPC。
- 缓存统一使用 Redis，并设置 TTL；缓存不写业务数据库，不能作为领域事实来源。
- 新增同步服务依赖前必须说明调用方向和失败影响；禁止为了查询方便跨服务连接业务库或维护跨服务数据库投影。
- 一个 proto 文件最多定义一个 service；每个 proto service 对应一个 Go service 文件。
- Proto 管理、lint 和生成统一使用 Buf；共享 API 生成使用 `make api`，校验使用 `make api-lint`。
- Buf 生成模板放在 `common/buf`，共享 API 模块配置放在 `common/api/app`。
- 第三方 proto 依赖统一通过 Buf deps 管理；不能继续维护 `common/api/third_party`。
- 项目初期不要求 proto 版本兼容；按当前需求可以删除、重命名或重排字段。
- BFF proto 按模块功能分目录存放，必须定义自己的 request、reply 和对外模型，不引用内部服务 proto message。
- 普通 BFF HTTP 对外业务接口默认使用 `POST` 和 `body: "*"`；回调、健康检查、图片、文件下载等不适合 POST 的接口按具体场景设计。
- 内部服务 proto 只暴露 gRPC 契约，不写 `google.api.http` 注解；OpenAPI 只面向 BFF。
- 契约分为 BFF HTTP、内部 gRPC、内部事件、外部回调四类；外部回调必须独立说明验签、幂等键、来源字段和失败处理。
- `google.api.http` option 必须使用多行块格式，路径和 `body` 各占一行；不要写成 `{ post: "..." body: "*" }` 这种单行格式。
- 共享 API 和 BFF 对外 RPC 命名使用业务动作；禁止 `GetOne`、`Page`、`Info`、`Data` 这类实现视角或含义不明确的名字。单资源查询用 `Get`，集合查询用 `List`，创建实体用 `Create`，追加子内容可用明确业务动词如 `Add`，状态变更使用 `Publish`、`MarkRead` 等业务动作。
- 查询 RPC 按资源收敛为单资源 `Get`、列表 `List` 和映射 `Map`；禁止因返回字段组合拆分 `GetBasic`、`BatchGetBasic`、`BatchGetContact` 这类接口。
- 调用方需要按 ID 获取映射时，归属服务必须定义独立 `Map` RPC；`List` 只表示列表查询并返回列表，调用方不能通过 `List` 的 `repeated` 结果自行组装映射。
- 基础资源模型只表达资源本体，不挂载可独立查询的关联集合；关联集合通过独立 RPC 暴露，不能由 `Get` 或 `List` 隐式聚合返回。
- 项目初期不维护 proto 向后兼容保留位，删除字段或枚举值时直接删除，不写 `reserved` 字段号或字段名。
- 请求参数不能用 `with_xxx`、`include_xxx` 等布尔开关控制响应结构；同一接口的响应结构必须由接口语义固定。
- RPC 注释只描述接口功能，不写消费方、调用场景、登录态、幂等策略和字段缺省规则。
- service 名已经表达资源时，RPC 只写业务动作；request/reply 外层 message 使用 `动作 + 资源名`，禁止 `UpdateArticleArticle` 这类重复资源名。
- BFF HTTP 路径统一使用 `/v1/{模块}/{资源}/{动作}`；同一资源名称保持单数且前后一致，动作使用 lower-kebab。
- BFF 不暴露邮箱、手机号、账号名是否存在等可用于枚举用户信息的接口。
- BFF OpenAPI 从 `google.api.http` 和 proto 注释生成；不手写 tags、servers、external_docs、BearerAuth，也不在接口描述中写登录态、幂等策略、字段缺省规则等运行时说明。
- SDK 从 BFF OpenAPI 生成，输出到 `common/api/gen-ts/<bff>`、`common/api/gen-go/<bff>`、`common/api/gen-java/<bff>` 和 `common/api/gen-rust/<bff>`，生成产物不入 Git 和 Docker。
- gRPC client 统一通过 `common/pkg/client/rpc` 和 `ConsulClient.GetGrpcConn` 创建；业务代码不要直接 `grpc.Dial` 或手写服务发现逻辑。
- gRPC 超时、发现、tracing、metadata 和连接复用由统一 client 管理；普通 RPC 调用不要重复包 `context.WithTimeout`。
- BFF 读接口可以按业务场景降级或返回部分结果；写接口下游失败时默认返回失败，不吞错伪造成功。
- 常规表业务尽量保持 proto、service、biz、data、schema 一一对应。
- 分层依赖只能从外层指向内层；biz 层禁止依赖本模块 `internal/data`、`internal/data/gen`、Ent 生成模型、Ent client、schema predicate 或 data 层实现类型。
- BFF 负责入口认证和权限判断；内部服务默认信任内部调用链传入的身份上下文，不重复实现入口层权限判断。
- 内部服务仍需校验资源存在、状态流转、归属服务本地业务不变量和写入前置条件。
- service 层只能调用 biz/usecase，不能直接注入或调用 repo、data 层实现；现有历史偏离不作为新代码依据。
- service 入参校验属于协议入口适配，写在对应 service 的接收者方法中；不要为入口校验新增游离函数、独立校验结构体或独立校验 usecase。
- biz 层底层能力接口只能暴露业务模型、基础类型、枚举和查询参数对象；不能在接口签名中出现 Ent client、Ent entity、Ent mutation、Ent query 等 data 层细节。
- Ent 生成模型只允许在 data 层内部使用；biz model 必须是独立业务结构，不能嵌入或别名引用 Ent entity。
- schema 字段必须来自明确业务需求或实际读写路径；意义不明、未使用、可推导或重复的信息不落库；任何数据库表结构变更都必须先给出设计、影响范围和理由，经用户确认后再修改。
- 日志表只记录后续排查和安全审计真正需要的事实；没有明确消费方的提交原文、失败文案、链路标识和客户端解析字段不落库。
- 密码、token、验证码、密钥任何时候都不能返回给用户，也不能写入日志、错误信息或事件 payload；邮箱、手机号、设备信息、IP 等字段按业务可见范围返回。
- 参与唯一约束的字段应避免可空；复合索引能覆盖左前缀查询时，不额外增加重复单列索引。
- 复杂业务新增明确业务 usecase；usecase 可按层级组合调用，但禁止循环调用。
- 偏好设置按业务归属拆分；端侧展示、本地化等账号体验设置与通知投递、订阅、渠道开关分表维护。
- 依赖字段命名遵守 [代码命名规范](doc/coding.md#命名)：`Repo` 只表示数据库或持久化访问，缓存访问依赖用 `XXXCache`，RPC、对象存储或第三方客户端依赖用 `XXXClient`，不要用 `Resolver`、`Writer`、`Manager` 等泛化后缀包装真实依赖类型。
- Wire 的 `ProviderSet` 命名允许保留；提供数组、slice、map 等非标准 struct 时可以使用 `ProvideXXX`，普通 struct 通过 `NewXXX` 构造。
- outbox repo 只提供通用保存能力；事件类型、主题、payload 和接收者由调用方按业务语义构造。
- outbox 表默认保持最小投递模型；确需新增投递字段时，必须先给出能说明必要性的设计并经用户确认。
- 所有模块的 outbox、inbox 表结构必须保持同类一致；新增、删除或调整字段时必须同步处理所有模块的同类表。
- 事件 payload 使用 common proto message；消费幂等以事件 ID 为主，必要时结合事件类型和业务主键。
- 事件处理失败必须保留可判断状态的 inbox 记录；是否增加死信、搁置、下次重试等字段必须先给出设计并经用户确认。
- 跨服务、跨层或落库约束枚举必须定义 proto enum，并通过 `common/pkg/enum.Mapping` 绑定内部 string enum。
- string 字段只用于自由文本、外部标识、地址、密钥、URL、时区、设备信息等开放值；业务固定取值必须抽成枚举。
- MQ subject 属于跨服务事件协议，必须定义在 common proto enum 中，并通过内部 string enum 映射实际主题字符串。
- MQ queue group 属于事件消费协议，必须定义为 common 枚举；业务消费者订阅时使用枚举值，不直接书写字符串。
- 复合查询条件使用 `Query`、`Filter`、`Spec` 等参数对象，不拼接超长方法名。
- biz/repo 的参数对象只表达持久化查询或写入条件，不定义 `StringPatch`、`XXXPatch` 等承载 service 入参校验状态的类型。
- 底层能力接口方法名使用业务动作和资源语义；单资源查询使用 `Get`，列表查询使用 `List`，按 ID 返回映射使用 `Map`，批量查询不使用 `BatchGetBasic`、`BatchGetContact` 这类按字段组合命名的方法。
- 返回 map 的 biz client/repo 方法使用 `Map` 语义命名，不能命名为 `ListAccounts`、`ListBasicAccounts` 等列表查询名称。
- 基础资源模型只写资源本体字段，不挂载标签、附言等可独立查询的关联集合；调用方需要关联集合时定义独立 RPC。
- 项目初期删除 proto 字段或枚举值时直接删除，不写 `reserved` 字段号或字段名。
- 文本文件使用 UTF-8、无 BOM、LF；所有注释使用中文。
- 业务规则、跨服务边界、事务边界、事件语义和数据库约束需要必要注释；注释描述意图、约束和边界，不复述代码表面行为。
- 跨服务调用、事件投递和消费、关键写事务必须记录 trace ID、业务 ID、耗时和结果状态等可追踪信息。
- 需要变动的配置从 `.env` 加载；新增或变更配置必须同步配置 proto 和 `.env` 示例，并添加中文注释。secret、token、密钥不写入 Git。
- 涉及表结构、跨服务依赖、事件契约、BFF 对外接口、缓存策略或信任边界的变更，必须先说明变更目的、影响范围和处理方式。
- 新增服务、跨服务事件、BFF 接口、数据库 schema、生成链路、公共包约定或配置项时，必须同步更新 `doc/`、`AGENTS.md` 和 `CLAUDE.md`。
- 代码格式化统一使用 Make：根目录执行 `make format`，共享 API 和 common Go 代码执行 `make api-format`，单模块执行 `make -C app/<module> format`。
- 格式化目标只编排 `gofmt`、`buf format` 等现有工具，不在项目内实现自定义格式化器。
- 需要基础类型指针时使用 Go 1.26 的 `new(expr)`；禁止使用 `common/pkg/util.Ptr` 或临时变量取地址来构造值指针。
- 内部方法的拆分边界是是否存在复用或独立复杂规则；只调用一次的简单代码块不要单独拆成私有方法。
- 不编写通用模型转换接口，也不编写 `ConvertToRpc`、`toXXX`、`xxxToDomain`、`buildXXXReply` 等内部转换或组装函数；数据库模型、业务模型和 proto DTO 之间的字段组装必须在具体业务函数内按当前接口需求显式完成。
- BFF 返回字段必须由当前接口的业务边界决定，不能通过复用转换函数隐式扩大返回字段。
- 默认不新增测试代码；只有用户明确要求、修复已有测试或维护现有测试文件时，才新增或修改测试代码。仍需按变更类型运行生成、编译或已有测试命令。

<!-- gitnexus:start -->
# GitNexus — Code Intelligence

This project is indexed by GitNexus as **Bass** (3284 symbols, 7249 relationships, 135 execution flows). Use the GitNexus MCP tools to understand code, assess impact, and navigate safely.

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
