# 代码规范

## 编码与局部抽象

- 文本文件统一使用 UTF-8、无 BOM、LF；所有注释使用中文，不要对无 BOM 文件执行 Remove BOM。
- 需要基础类型指针时使用 Go 1.26 的 `new(expr)`；禁止使用 `common/pkg/util.Ptr` 或临时变量取地址。
- 只在存在复用或独立复杂规则时拆分内部方法；一次性简单逻辑、结构体组装、字段派生和 service 入参校验写在当前业务函数或对应 service 接收者方法中，不新增游离私有函数、独立校验结构体或独立校验 usecase。
- 不新增通用转换接口、游离组装函数或 `ConvertToRpc`、`toXXX`、`buildXXXReply` 等转换函数；数据库模型、业务模型和 proto DTO 按当前接口显式组装，BFF 返回字段只由当前接口边界决定。
- Proto 中保存、更新等写入请求的输入结构定义在对应 RPC 外层 message 内部，不能抽成 `XXXSave`、`XXXUpdate` 等跨接口复用结构；查询返回模型和查询参数可以按稳定视图复用 `model.proto` 或当前资源的通用查询结构。
- 内部服务的业务输入只允许来自 proto request 显式字段；service 只从 request 读取当前用户、操作人、审计人、目标资源 ID、动作参数和写入值，并显式传入 usecase。内部服务禁止从 ctx、metadata、请求头、JWT、缓存会话等隐式来源提取业务身份或业务参数；`created_by`、`updated_by` 由写入参数显式赋值，不能依赖 Ent hook 自动从上下文写入。
- ctx 在内部服务中只用于取消、超时、trace 和事务传递；配置、本服务数据库中的领域事实、当前时间、UUID、事件 ID 等系统生成值可以参与业务判断，但不能替代 proto request 中应显式声明的调用方输入。
- BFF 写接口不做 read-before-write 的领域状态预校验；资源存在性、作者关系、状态流转、本地业务不变量和并发写入正确性由归属服务校验。BFF 只保留入口认证、端侧权限、请求格式、默认值、读接口可见范围和 DTO 适配。
- content 内容状态使用多字段模型：文章用 `publish_status` 表达草稿、已发布、归档，用 `visibility` 表达公开、私有，用 `restriction` 表达无管理限制、管理隐藏、管理锁定；评论和附言使用 `restriction` 表达管理限制；逻辑删除只使用 `deleted_at`，应用不提供查询已删除数据的入口，不保留 `deleted` 枚举状态和业务表删除人字段。BFF 控制端侧可见范围，content 校验状态流转和本地业务不变量。
- 处理 slice、map、set 的去重、映射、过滤、分组、键值提取等纯数据变换时，优先使用 `github.com/samber/lo`；存在副作用、错误短路、事务操作或可读性下降时保留显式循环。
- 注释描述业务规则、跨服务边界、事务边界、事件语义和数据库约束的意图，不复述代码表面行为。
- 对外可感知的业务错误统一使用 `common/pkg/apperror.New(BusinessErrorCode, ...)` 构造，不接收用户可见文案；业务校验失败禁止使用 `cerrors.ErrorBadRequest`、`cerrors.ErrorNotFound` 等 `ErrorReason` 构造函数，缺少专用业务错误码时使用 `BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT`；动态错误参数必须定义为 common proto message，并通过 `apperror.WithData[T proto.Message]` 写入，禁止 key/value 拼装。
- 新增或调整 `BusinessErrorCode` 时必须添加中文行尾注释，并同步对应 BFF 的错误文案映射；基础设施、第三方客户端、生成代码或内部不可恢复错误可以保留普通 error，由入口层统一映射。

## 命名

- 依赖字段按来源和类型后缀命名：数据库或持久化访问使用 `<资源名>Repo`，缓存使用 `<资源名>Cache`，RPC、对象存储或第三方客户端使用 `<服务名或客户端名>Client`，usecase 使用 `<业务名>Usecase`。
- 不使用 `Resolver`、`Writer`、`Provider`、`Manager` 等泛化后缀包装真实依赖类型；Wire 的 `ProviderSet` 命名允许保留，提供数组、slice、map 等非标准 struct 时可以使用 `ProvideXXX`，普通 struct 通过 `NewXXX` 构造。
- 基础设施字段沿用项目约定：配置为 `conf`，日志 helper 为 `log`，事务函数为 `tx`，Ent client 为 `db`；构造函数参数与结构体字段保持同名，除 `logger log.Logger` 这类转换参数外不用泛化名称代替真实依赖。
- repo 层每个 schema 对应的 repo 都定义 `Get`、`List`、`Map`、`Count`、`Page` 五个查询方法；复合查询条件使用 `Query`、`Filter`、`Spec` 等参数对象；biz/repo 参数对象只表达持久化查询或写入条件；底层能力接口方法名使用业务动作和资源语义，返回 map 的 biz client/repo 方法使用 `Map` 语义命名。

## 生成

- 生成代码、构建、测试前先阅读根目录和模块内 Makefile；运行 make 相关命令时使用 bash 环境。
- 不手改生成代码；修改 proto、schema、wire、config 等源文件后运行对应生成命令。
- 服务配置以 `internal/conf/conf.proto` 为契约来源；`configs/config.yaml` 必须写出该服务暴露的可配置字段，不能写 proto 未定义字段。引用公共配置 message 时，只写入本服务实际承担职责所需的配置段，避免用空配置表达未启用能力。`google.protobuf.Duration` 配置统一使用秒单位，禁止写 `ms`、`m`、`h` 等单位；布尔配置必须使用 `true` 或 `false` 字面量，禁止使用环境变量占位。
- 热重载只用于运行期参数，例如业务阈值、限流、告警、事件轮询和死信扫描参数；端口、数据库、Redis、NATS、Consul、JWT secret、OSS 密钥、短信密钥等连接类或敏感配置必须通过重启生效。
- 代码格式化统一使用 Make：根目录执行 `make format`，共享 API 和 common Go 代码执行 `make api-format`，单模块执行 `make -C app/<module> format`；格式化目标只编排 `gofmt`、`buf format` 等现有工具，Proto 文件格式以 `buf format` 输出为准。
- 共享 API 生成使用 `make api`，校验使用 `make api-lint`，根级批量生成优先使用 `make gen-all`；Ent 代码、BFF OpenAPI 文档和 SDK 都使用 Makefile 中声明的命令生成。
- BFF 请求必填字段使用 `gnostic.openapi.v3.schema.required` 标注，只标注请求模型，不标注响应字段，option 放在字段定义之后；SDK 使用 BFF OpenAPI 生成，配置放在 `common/proto/sdk/` 下，生成产物不提交，客户端对象使用 proto service 名，方法使用 RPC 动作名。

## 依赖与测试

- 依赖整理按模块执行，不在无关模块制造 `go.mod` / `go.sum` 变更；新增依赖前优先复用项目已有库和公共工具。
- 默认不新增测试代码；只有用户明确要求、修复已有测试或维护现有测试文件时，才新增或修改测试代码。
- 改 proto 后至少运行共享 API 生成并编译受影响模块；改 BFF proto 或 OpenAPI 注释后至少运行 `make -C app/<bff> doc`；改 SDK 配置后运行 `make -C app/<bff> sdk-validate`；改 Ent schema 后运行对应模块 Ent 生成并编译受影响模块。
- 改 usecase、底层能力接口或实现、service 后优先运行对应模块已有测试或编译命令；横切公共包变更需要扩大到所有受影响模块。

## 迁移与文档

- 开发环境可使用自动迁移；生产环境迁移必须走显式变更流程。
- content 已发布文章的作者编辑窗口由 content 配置 `business.article.published_edit_window` 和 `business.article.published_edit_max_view_count` 控制，默认配置为 10 分钟且浏览数小于 100。
- 项目初期允许破坏性 schema 重构，但所有表结构变更仍必须先给出设计、影响范围和理由，经用户确认后再修改；字段删除、重命名、唯一约束调整必须先评估历史数据影响和回滚方式。
- 涉及表结构、outbox/inbox、跨服务依赖、事件契约、BFF 对外接口、缓存策略或信任边界的变更，必须先说明目的、影响和处理方式，并经用户确认。
- 新增服务、跨服务事件、BFF 接口、数据库 schema、生成链路、公共包约定或配置项时，必须同步更新 `doc/`、`AGENTS.md` 和 `CLAUDE.md`；修改规范时保持入口一致，删除或重命名规范文件时同步清理 README、Agent 提示和其他 Markdown 旧链接。
