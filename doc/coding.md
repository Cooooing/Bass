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
- 不新增通用转换接口、游离组装函数或 `ConvertToRpc`、`toXXX`、`buildXXXResponse` 等转换函数。
- slice、map、set 的去重、映射、过滤、分组和键值提取优先使用 `github.com/samber/lo`；存在副作用、错误短路或事务操作时使用显式循环。
- 不手改生成代码，不在无关模块制造 `go.mod` 或 `go.sum` 变更。

## Usecase 接口

- usecase receiver 方法除 `ctx` 外，只接收一个 `*XxxReq`；即使当前只有一个业务参数，也要定义请求结构体。
- usecase receiver 方法有业务返回值时返回 `*XxxResponse, error`；只有副作用且无业务返回值时可以只返回 `error`。
- `XxxReq` 和 `XxxResponse` 定义在对应方法正上方，不放到集中式 `query.go` 或 `dto.go` 文件。
- `XxxReq` 不使用 proto request、repo query、data/gen 类型或具体数据库类型；`XxxResponse` 不使用 proto response。
- repo 查询对象只在 usecase 方法内部构造，不向 service 层暴露。
- 请求或响应存在复杂嵌套业务结构时，结构体定义到 `biz/model`，usecase 方法上方仍保留轻量 `XxxReq` / `XxxResponse` 作为边界。


## Repo 接口

- biz/repo 接口方法除 `ctx` 外，只接收一个 `*XxxReq`；即使当前只有一个业务参数，也要定义请求结构体。
- biz/repo 接口方法返回 `*XxxResponse, error`；只有确实不需要业务返回值的动作，也要返回空响应结构体。
- `XxxReq` 和 `XxxResponse` 定义在对应 repo 文件中，不使用 proto request、proto response、data/gen 类型、Ent predicate 或数据库类型。
- BFF 的 RPC client 适配也按 repo 接口处理；actor、token、分页、查询条件都作为 repo request 显式字段传入。
- proto request/response 只在 service 层和 data 层 RPC 适配代码中出现，不穿透到 biz/repo 接口。

## 命名约定

- 依赖字段按类型使用 `Repo`、`Cache`、`Client`、`Usecase` 后缀，不使用 `Manager`、`Resolver`、`Writer` 包装真实依赖。
- 配置目录、文件和包名统一使用 `config`；日志 helper 为 `log`，事务函数为 `tx`，Ent client 为 `db`。
- 普通结构体通过 `NewXXX` 构造；Wire 的 `ProviderSet` 命名允许保留。
- repo 层每个 schema 对应的 repo 定义 `Get`、`List`、`Map`、`Count`、`Page` 五个查询方法。
- 底层能力接口方法使用业务动作和资源语义命名，返回映射的方法使用 `Map` 语义。

## 生成与验证

- 生成、格式化、测试和构建优先使用 Makefile；运行 make 使用 bash。
- 修改 proto 后运行共享 API 生成和 lint，并编译受影响模块。
- 修改 Ent schema 后运行对应模块 Ent 生成，并编译受影响模块。
- 修改 BFF OpenAPI 注释或 SDK 配置后，运行对应模块的文档或 SDK 校验目标。
- 修改 usecase、repo、service 后，运行受影响模块已有测试或编译命令。
