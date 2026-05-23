# API 与 RPC

## Proto

- Proto 管理、lint 和生成统一使用 Buf；不要再新增手写 `protoc` 生成链路。
- Buf 生成模板放在 `common/buf`；共享 API 模块配置放在 `common/api/app`。
- 第三方 proto 依赖统一通过 Buf deps 管理；不能在 `common/api/third_party` 下继续维护 vendored proto。
- 无法从 Buf deps 获取且必须长期稳定的扩展，可以定义为项目自有 proto，并放在 `common/api/app` 下。
- 一个 proto 文件最多定义一个 `service`，每个 RPC 必须写注释。
- service、rpc、request、response 命名和业务概念一致，避免无意义前缀。
- 请求/响应围绕接口语义定义，不复用过大的通用模型。
- 字段编号发布后不复用；删除或废弃字段使用 `reserved`。
- 内部服务 proto 表达领域能力；BFF proto 表达端侧体验和聚合结果。
- BFF proto 按模块功能分目录存放；单模块透传能力放在对应模块目录，多服务聚合能力按具体业务归属放入既有目录或新建业务目录。
- BFF proto 必须定义自己的 request、reply 和对外模型，不引用内部服务 proto message；内部服务契约变化只影响 BFF 适配代码，不直接影响对外 API。
- BFF HTTP 接口默认使用 `POST` 和 `body: "*"`；除图片、文件下载等必要场景外，不使用路径参数。
- BFF 不暴露邮箱、手机号、账号名是否存在等可用于枚举用户信息的接口。
- BFF HTTP 路径统一使用 `/v1/{模块}/{资源}/{动作}`；同一资源名称保持单数且前后一致，动作使用 lower-kebab 动词短语。
- BFF 路径使用对外业务概念，不暴露内部投影名、表名或 RPC 实现名。

## OpenAPI

- BFF 对外文档统一生成 OpenAPI 3；基础路由信息来自 `google.api.http`。
- BFF 文档入口放在 BFF proto 包的 `doc.proto`，用于聚合该 BFF 的对外 proto 文件。
- BFF OpenAPI 不在 proto 中声明环境地址、外部文档、全局标签、鉴权方案。
- BFF OpenAPI 不在 RPC 中声明 BearerAuth；认证由运行时中间件负责，客户端请求头由调用方按需传入。
- 每个 BFF RPC 必须写简短注释，作为 OpenAPI 接口说明来源。
- `summary` 和 `description` 应简短描述接口功能，不写登录态、幂等策略、适用场景、字段缺省更新规则等运行时说明。
- 对外 message 和字段使用 proto 注释描述含义，描述必须简短。
- OpenAPI 描述使用短语或短句，避免标点堆叠和长段说明。
- OpenAPI 文档是 SDK 生成和外部集成的主契约，描述必须面向客户端，不暴露内部服务、表名、投影名或实现细节。

## SDK

- SDK 只从 BFF OpenAPI 文档生成，不直接从内部服务 proto 生成。
- SDK 生成目录放在 `common/api` 下，例如 `common/api/gen-ts/<bff>`、`common/api/gen-go/<bff>`。
- SDK 生成产物不参与 Git 和 Docker 上下文；需要通过 Make 命令重新生成。
- SDK 配置放在 `common/api/sdk/` 下，生成器、包名、枚举风格等参数应显式配置。

## 查询与分页

- 高频查询接口可以拆细，不做万能查询接口。
- 多条件查询使用查询对象表达，不把条件堆进方法名。
- 分页、排序、过滤字段要有默认值和边界说明。

## RPC 调用

- RPC client 统一由基础设施层创建和注入，不在业务代码散落连接。
- 所有跨服务调用必须带 timeout/deadline。
- 写接口默认不自动 retry；读接口 retry 需要确认幂等和超时预算。
- BFF 聚合读要设计部分失败、降级和超时策略。

## 错误与校验

- 参数、鉴权、权限、资源不存在、状态冲突、下游失败要区分。
- service 层做协议适配和格式校验；usecase 层做业务规则校验。
- repo 层把数据库错误转换为数据访问语义错误。
- 对外错误信息稳定且不泄露密码、token、验证码、SQL、内部拓扑。
