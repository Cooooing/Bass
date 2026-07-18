# 代码规范

本文档记录项目级编码规范和代码组织结构。服务边界、读写归属、契约、事件、缓存、运行时和错误响应规则见 [架构设计](architecture.md)；新增同类代码时参考根 README 的 [模板索引](../README.md#模板索引)。

## 代码组织

- `service` 层负责协议适配、请求读取、基础校验、调用 usecase 和组装响应，不直接调用 repo 或 data。
- `biz` 层负责业务规则、事务编排和底层能力接口定义，不依赖 data、Ent 生成模型、schema predicate、具体数据库类型或 proto 生成结构。
- `data` 层负责持久化、缓存、外部客户端适配和 Ent 生成模型转换；Ent 生成模型只在 data 层内部使用。
- Ent schema 放在对应服务的 `internal/data/schema`；配置契约放在 `internal/config/config.proto`。
- 共享 proto、Buf 配置、跨服务公共约定和通用工具放在 `common`，不承载具体业务流程。

## 编码风格

- Go 代码按 Go 1.26 编写；需要基础类型指针时使用 `new(expr)`，不使用 `common/pkg/util.Ptr` 或临时变量取地址。
- 适量添加注释，注释和文档使用简体中文，只描述意图、约束和边界，不复述代码表面行为。
- 一次性简单逻辑、入参校验和字段组装保留在当前函数内，不拆分无复用价值的辅助函数。
- 不新增通用转换接口、游离组装函数或 `ConvertToRpc`、`toXXX`、`buildXXXResp` 等转换函数。
- slice、map、set 的去重、映射、过滤、分组和键值提取优先使用 `github.com/samber/lo`；存在副作用、错误短路或事务操作时使用显式循环。
- 不手改生成代码，不在无关模块制造 `go.mod` 或 `go.sum` 变更。

## Usecase 接口

- usecase receiver 方法的业务参数数量不包含 `ctx`；业务返回值数量不包含 `error`。
- 无业务参数时不定义 `XxxReq`；一个业务参数时直接传参；两个及以上业务参数时定义 `*XxxReq`。
- 无业务返回值时只返回 `error`；一个业务返回值时直接返回该值和 `error`；两个及以上业务返回值时定义 `*XxxResp`。
- 需要 `XxxReq` 或 `XxxResp` 时，定义在对应方法正上方，不放到集中式 `query.go` 或 `dto.go` 文件。
- `XxxReq` 不使用 proto request、repo query、data/gen 类型或具体数据库类型；`XxxResp` 不使用 proto resp。
- repo 查询对象只在 usecase 方法内部构造，不向 service 层暴露。
- 请求或响应存在复杂嵌套业务结构时，结构体定义到 `biz/model`，usecase 方法上方只保留必要的轻量 `XxxReq` / `XxxResp` 边界。

## Repo 接口

- biz/repo 接口方法的业务参数数量不包含 `ctx`；业务返回值数量不包含 `error`。
- 无业务参数时不定义 `XxxReq`；一个业务参数时直接传参；两个及以上业务参数时定义 `*XxxReq`。
- 无业务返回值时只返回 `error`；一个业务返回值时直接返回该值和 `error`；两个及以上业务返回值时定义 `*XxxResp`。
- 需要 `XxxReq` 或 `XxxResp` 时，定义在对应 repo 文件中，不使用 proto request、proto resp、data/gen 类型、Ent predicate 或数据库类型。
- BFF 的 RPC client 适配也按 repo 接口处理；actor、token、分页、查询条件中有两个及以上业务参数时封装为 repo request。
- proto request/resp 只在 service 层和 data 层 RPC 适配代码中出现，不穿透到 biz/repo 接口。

## 命名约定

- 依赖字段按类型使用 `Repo`、`Cache`、`Client`、`Usecase` 后缀，不使用 `Manager`、`Resolver`、`Writer` 包装真实依赖。
- 配置目录、文件和包名统一使用 `config`；日志 helper 为 `log`，事务函数为 `tx`，Ent client 为 `db`。
- 普通结构体通过 `NewXXX` 构造；Wire 的 `ProviderSet` 命名允许保留。
- repo 层每个 schema 对应的 repo 定义 `Get`、`List`、`Map`、`Count`、`Page` 五个查询方法。
- 底层能力接口方法使用业务动作和资源语义命名，返回映射的方法使用 `Map` 语义。

## Git 提交

- 提交信息使用 `<type>(<scope>): <summary>` 格式，例如 `refactor(platform): consolidate integration capabilities`。
- `type` 使用 `feat`、`fix`、`refactor`、`docs`、`test`、`chore`、`build`、`ci`、`perf`、`style` 或 `revert`。
- `scope` 使用受影响模块或目录名，例如 `bbs`、`user`、`platform`、`common`、`ops`、`doc`；跨多个模块时使用 `all`。
- `summary` 使用简短英文动词短语，首字母小写，不以句号结尾。
- 一次提交只表达一个主要意图；混合修改时以主要变更选择 `type` 和 `scope`。

## 生成与验证

- 生成、格式化、测试和构建优先使用 Makefile；运行 make 使用 bash。
- 修改 proto 后运行共享 API 生成和 lint，并编译受影响模块。
- 修改 Ent schema 后运行对应模块 Ent 生成，并编译受影响模块。
- 修改 BFF OpenAPI 注释或 SDK 配置后，运行对应模块的文档或 SDK 校验目标。
- 修改 usecase、repo、service 后，运行受影响模块已有测试或编译命令。
