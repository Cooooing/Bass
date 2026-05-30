# 代码规范

## 编码

- 文本文件统一使用 UTF-8、无 BOM、LF；所有注释使用中文。
- 不要对无 BOM 文件执行 Remove BOM，避免误删文件前三个字符。
- 需要基础类型指针时使用 Go 1.26 的 `new(expr)`；禁止使用 `common/pkg/util.Ptr` 或临时变量取地址。
- 内部方法的拆分边界是复用或独立复杂规则；只调用一次的简单代码块不要单独拆成私有方法。
- 不编写通用模型转换接口，也不编写 `ConvertToRpc`、`toXXX`、`xxxToDomain`、`buildXXXReply` 等内部转换或组装函数。
- 数据库模型、业务模型和 proto DTO 之间的字段组装必须在具体业务函数内按当前接口需求显式完成。
- 简单结构体组装和派生展示字段不要新增游离私有函数；需要复用时优先用结构体方法表达派生值，或在具体业务函数内显式组装。
- BFF 返回字段必须由当前接口的业务边界决定，不能通过复用转换函数隐式扩大返回字段。
- 业务规则、跨服务边界、事务边界、事件语义和数据库约束需要必要注释；注释描述意图、约束和边界，不复述代码表面行为。
- service 入参校验属于协议入口适配，写在对应 service 的接收者方法中；不要为入口校验新增游离函数、独立校验结构体或独立校验 usecase。

## 命名

- 依赖字段按来源和类型后缀命名，避免混用业务别名、层级别名和抽象动作名。
- 数据库或持久化访问依赖使用 `<资源名>Repo`，例如 `accountRepo`。
- 缓存访问依赖使用 `<资源名>Cache`，例如 `tokenCache`。
- RPC、对象存储或第三方客户端依赖使用 `<服务名或客户端名>Client`，例如 `userClient`、`redisClient`、`natsClient`。
- usecase 依赖使用 `<业务名>Usecase`，例如 `tokenUsecase`。
- 不使用 `Resolver`、`Writer`、`Provider`、`Manager` 等泛化后缀包装真实依赖类型。
- Wire 的 `ProviderSet` 命名允许保留；提供数组、slice、map 等非标准 struct 时可以使用 `ProvideXXX`，普通 struct 通过 `NewXXX` 构造。
- 基础设施字段沿用项目约定：配置为 `conf`，日志 helper 为 `log`，事务函数为 `tx`，Ent client 为 `db`。
- 构造函数参数与结构体字段保持同名；除 `logger log.Logger` 这类转换参数外，不用 `service`、`domain`、`manager` 等泛化名称代替真实依赖。
- 复合查询条件使用 `Query`、`Filter`、`Spec` 等参数对象，不拼接超长方法名。
- biz/repo 的参数对象只表达持久化查询或写入条件，不定义 `StringPatch`、`XXXPatch` 等承载 service 入参校验状态的类型。
- BFF 的 biz/repo 接口、usecase 和 data 实现文件按具体调用语义命名和拆分，例如 user 相关能力按 auth、account、relation 等拆分；不要把一个下游服务的所有调用收敛到 `UserRepo`、`ContentRepo`、`NotifyRepo` 这类大接口。
- 底层能力接口方法名使用业务动作和资源语义；单资源查询使用 `Get`，列表查询使用 `List`，按 ID 返回映射使用 `Map`，批量查询不使用 `BatchGetBasic`、`BatchGetContact` 这类按字段组合命名的方法。
- 返回 map 的 biz client/repo 方法使用 `Map` 语义命名，不能命名为 `ListAccounts`、`ListBasicAccounts` 等列表查询名称。
- 基础资源模型只写资源本体字段，不挂载标签、附言等可独立查询的关联集合；调用方需要关联集合时定义独立 RPC。
- 项目初期删除 proto 字段或枚举值时直接删除，不写 `reserved` 字段号或字段名。

## 生成

- 生成代码、构建、测试前先阅读根目录和模块内 Makefile；运行 make 相关命令时使用 bash 环境。
- 不手改生成代码；修改 proto、schema、wire、config 等源文件后运行对应生成命令。
- 代码格式化统一使用 Make：根目录执行 `make format`，共享 API 和 common Go 代码执行 `make api-format`，单模块执行 `make -C app/<module> format`。
- 格式化目标只编排 `gofmt`、`buf format` 等现有工具，不在项目内实现自定义格式化器；Proto 文件格式以 `buf format` 输出为准。
- 共享 API 生成使用 `make api`，校验使用 `make api-lint`；根级批量生成优先使用 `make gen-all`。
- 服务内部 config proto 使用 Make 内联 Buf 配置，不为每个服务新增独立 Buf 配置文件。
- Ent 代码使用模块内 `make ent` 或 Makefile 中声明的等价命令生成。
- BFF OpenAPI 文档使用 `make -C app/<bff> doc` 生成。
- BFF 对外请求参数的必填字段使用 `gnostic.openapi.v3.schema.required` 标注；只标注请求 `Request` 和请求专用输入模型，不标注响应字段。`option (gnostic.openapi.v3.schema)` 放在字段定义之后，字段注释只描述业务含义，不同步写“必填”或“选填”。
- BFF SDK 使用 `make -C app/<bff> sdk-typescript-axios`、`make -C app/<bff> sdk-typescript-fetch`、`make -C app/<bff> sdk-typescript`、`make -C app/<bff> sdk-go`、`make -C app/<bff> sdk-java`、`make -C app/<bff> sdk-rust` 或 `make -C app/<bff> sdk` 生成；`sdk-typescript` 同时生成 Axios 和 Fetch 两套 TypeScript SDK。
- SDK 配置放在 `common/api/sdk/` 下；SDK 使用 OpenAPI Generator CLI，需要 Node.js、npx 和 Java，生成产物不提交到主业务分支。
- SDK 客户端对象名使用 proto service 名，方法名使用 RPC 动作名；生成命令必须去掉 OpenAPI operationId 中的 service 前缀和默认 API 后缀。
- 根目录批量目标包括 `make doc-all`、`make sdk-validate-all`、`make sdk-all`。

## 依赖与测试

- 依赖整理按模块执行，不在无关模块制造 `go.mod` / `go.sum` 变更；新增依赖前优先复用项目已有库和公共工具。
- 默认不新增测试代码；只有用户明确要求、修复已有测试或维护现有测试文件时，才新增或修改测试代码。
- 改 proto 后至少运行共享 API 生成，并编译受影响模块。
- 改 BFF proto 或 OpenAPI 注释后至少运行 `make -C app/<bff> doc`。
- 改 SDK 生成配置后至少运行 `make -C app/<bff> sdk-validate`。
- 改 Ent schema 后运行对应模块 Ent 生成，并编译受影响模块。
- 改 usecase、底层能力接口或实现、service 后优先运行对应模块已有测试或编译命令。
- 横切公共包变更需要扩大到所有受影响模块。

## 迁移

- 开发环境可使用自动迁移；生产环境迁移必须走显式变更流程。
- 项目初期允许破坏性 schema 重构；但所有表结构变更仍必须先给出设计、影响范围和理由，经用户确认后再修改。
- 字段删除、重命名、唯一约束调整必须先评估历史数据影响和回滚方式。

## 变更与文档

- 涉及表结构、跨服务依赖、事件契约、BFF 对外接口、缓存策略或信任边界的变更，必须先说明变更目的、影响范围和处理方式。
- 影响数据库表结构、outbox/inbox 字段或对外契约的变更，必须经用户确认后再修改。
- 新增服务、跨服务事件、BFF 接口、数据库 schema、生成链路、公共包约定或配置项时，必须同步更新 `doc/`、`AGENTS.md` 和 `CLAUDE.md`。
- 修改规范时保持入口一致：长期规范写入 `doc/architecture.md` 或 `doc/coding.md`，高频必守约束同步到 `AGENTS.md` 和 `CLAUDE.md`。
- 删除或重命名规范文件时，必须同步清理 README、Agent 提示和其他 Markdown 中的旧链接。
