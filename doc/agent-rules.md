# Agent 可执行规则

本文档只记录可以被人工或脚本检查的硬规则。规则说明业务背景时保持简短，具体设计解释放在 [架构设计](architecture.md) 或 [代码规范](coding.md)。

## 使用方式

- 修改代码前，先按变更类型阅读对应规则。
- 修改代码后，用“检查”段落中的方式确认没有违规。
- 后续新增 `rulecheck` 时，以本文档中的规则编号作为输出编号。

## RULE-001 内部服务禁止隐式获取业务身份

必须：
- 当前用户、操作人、审计人、目标资源 ID、动作参数和写入值来自内部 gRPC proto request 显式字段。
- service 从 request 读取字段并显式传入 usecase。

禁止：
- 内部服务从 ctx、metadata、请求头、JWT、Redis 会话或缓存会话提取当前用户或业务参数。
- Ent hook 从 ctx 自动写入 `created_by`、`updated_by`。

例外：
- ctx 只允许用于取消、超时、trace、事务和 data 层技术查询模式。

检查：
- 扫描 `app/*/internal/service`、`app/*/internal/biz` 中的 `metadata.FromIncomingContext`、`jwt`、`auth.FromContext`、`GetContextValue`、`SetContextValue` 等调用。
- 扫描 schema hook 是否读取 ctx 中的用户身份。

## RULE-002 BFF 不做领域状态预校验

必须：
- BFF 负责入口认证、端侧权限、请求格式、默认值、读接口可见范围和 DTO 适配。
- 资源存在性、作者关系、状态流转、本地业务不变量和并发写入正确性由归属服务校验。

禁止：
- BFF 写接口为了判断能否写入而先查询归属服务状态再自行裁决。
- BFF 直接发布领域 MQ 事件。

例外：
- BFF 可以拦截明显的协议错误，例如必填字段、分页范围、枚举值映射失败。

检查：
- 审查 BFF 写接口中同一流程是否同时存在 `Get/List/Page` 查询和写入 RPC。
- 审查 BFF 写接口是否调用 MQ producer。

## RULE-003 分层依赖方向固定

必须：
- service 调用 usecase。
- data 实现 biz 底层能力接口。
- biz 接口签名只暴露业务模型、基础类型、内部枚举和查询参数对象。

禁止：
- service 直接注入或调用 repo/data。
- biz import data、Ent、schema predicate 或 data 实现类型。
- 服务之间 import 其他服务的 `internal`。

例外：
- 生成代码所在目录可以按生成器要求引用生成包。

检查：
- 扫描 `internal/biz` import 中是否出现 `internal/data`、`internal/data/gen`、`entgo.io/ent`。
- 扫描跨服务 `internal` import。

## RULE-004 Proto 文件职责固定

必须：
- 枚举放在独立 `enum.proto`。
- `model.proto` 只放稳定数据模型。
- 保存、更新等写入请求结构定义在对应 RPC 外层 message 内部。

禁止：
- 在 `model.proto` 定义 enum、Query、Save、Update、Callback 或事件 payload。
- 抽出 `XXXSave`、`XXXUpdate` 给多个写接口复用。
- 内部服务 proto 写 `google.api.http` 注解。
- BFF proto 引用内部服务 message。

例外：
- 查询返回模型和查询参数可以按稳定视图复用 `model.proto` 或当前资源的通用查询结构。

检查：
- 扫描 `model.proto` 中的 `enum`、`Query`、`Save`、`Update`、`Callback`。
- 扫描内部服务 proto 中的 `google.api.http`。
- 扫描 BFF proto import 是否引用内部服务 proto message。

## RULE-005 RPC 和响应命名固定

必须：
- RPC 使用业务动作命名。
- 查询按资源收敛为 `Get`、`List`、`Page`、`Map`。
- `List`、`Page` 响应数组字段命名为 `rows`，分页信息字段命名为 `page`。

禁止：
- 使用 `GetOne`、`Info`、`Data` 等实现视角命名。
- 响应结构通过 `with_xxx`、`include_xxx` 控制返回形状。
- 混用 `items`、`list`、`data_list` 等数组字段名。

例外：
- 文件、健康检查、外部回调等协议接口可以按场景使用专门命名。

检查：
- 扫描 proto RPC 名称和响应字段名称。
- 扫描 request 字段中的 `with_`、`include_`。

## RULE-006 错误处理使用业务错误码

必须：
- 对外可感知业务错误使用 `common/pkg/apperror.New(BusinessErrorCode, ...)`。
- 动态错误参数定义为 common proto message，并通过 `apperror.WithData[T proto.Message]` 写入。
- 新增或调整 `BusinessErrorCode` 时添加中文行尾注释，并同步 BFF 错误文案映射。

禁止：
- 业务校验失败使用 `cerrors.ErrorBadRequest`、`cerrors.ErrorNotFound` 等 `ErrorReason` 构造函数。
- 用 key/value 拼装动态错误参数。
- 在内部服务返回用户可见文案。

例外：
- 基础设施、第三方客户端、生成代码或内部不可恢复错误可以保留普通 error，由入口层统一映射。

检查：
- 扫描业务代码中的 `cerrors.ErrorBadRequest`、`cerrors.ErrorNotFound`。
- 扫描 `BusinessErrorCode` 是否缺少中文注释。

## RULE-007 service 层只做协议适配

必须：
- service 负责 request 校验、枚举映射、调用 usecase、显式组装 reply。
- 业务枚举在 service 层从 proto 枚举转换为内部枚举，再传入 biz。

禁止：
- service 新增包级游离私有 helper。
- 新增 `toXXX`、`buildXXXReply`、`ConvertToRpc` 等通用转换函数。
- biz 层使用 proto 生成枚举。

例外：
- 单个接收者方法内部可以使用局部变量分段组装复杂响应。

检查：
- 扫描 service 包中的包级小写函数。
- 扫描 `to[A-Z]`、`build[A-Z].*Reply`、`ConvertToRpc`。
- 扫描 `internal/biz` import 中是否出现 `common/proto/gen`。

## RULE-008 repo 查询能力固定

必须：
- 每个 schema 对应的 repo 定义 `Get`、`List`、`Map`、`Count`、`Page` 五个查询方法。
- 复合查询条件使用 `Query`、`Filter`、`Spec` 等参数对象。
- 返回 map 的 biz client/repo 方法使用 `Map` 语义命名。

禁止：
- 用业务无关的 `Manager`、`Resolver`、`Writer` 包装真实依赖。
- repo 方法同时表达多个不相关资源的写入。

例外：
- 复杂业务动作可以新增语义明确的方法，例如 `Publish`、`Archive`、`ReplaceTags`。

检查：
- 对每个 schema 检查 repo 接口是否缺少五个基础查询方法。
- 扫描依赖字段命名是否使用泛化后缀。

## RULE-009 数据模型和生成模型隔离

必须：
- Ent 生成模型只在 data 层内部使用。
- biz model 是独立业务结构。
- data 层显式组装数据库模型和业务模型。

禁止：
- biz model 嵌入、别名引用 Ent entity。
- biz 接口暴露 Ent predicate、Ent entity 或具体数据库类型。

例外：
- data 层内部可以使用 Ent 生成类型完成持久化。

检查：
- 扫描 `internal/biz` import 中是否出现 `internal/data/gen`、`entgo.io/ent`。
- 扫描 biz model 是否使用 Ent 生成类型。

## RULE-010 逻辑删除和唯一索引规则

必须：
- 逻辑删除使用 `deleted_at` 表达。
- 业务 repo 查询固定过滤 `deleted_at IS NULL`。
- 删除、隐藏、锁定等操作的操作者、原因和来源写入审核或审计记录。
- PostgreSQL 软删除场景下的业务唯一约束使用 `WHERE deleted_at IS NULL` 的部分唯一索引。

禁止：
- 业务表保留 `deleted_by`。
- 使用 `deleted` 枚举状态表达逻辑删除。
- 用可空 `deleted_at` 参与普通唯一约束表达“未删除唯一”。
- 提供 `SoftDeleteMode`、`IncludeDeleted`、`Deleted` 等查询已删除数据的程序入口。

例外：
- 非业务表的基础设施日志可以按自身审计需求设计字段。

检查：
- 扫描 `deleted_by`、`DeletedBy`。
- 扫描状态枚举中的 `deleted`。
- 扫描 `SoftDeleteMode`、`WithSoftDeleteMode`、`IncludeDeleted`、`optional bool deleted`。
- 审查 Ent index 是否使用部分唯一索引表达软删除唯一性。

## RULE-011 content 状态模型固定

必须：
- 文章使用 `publish_status` 表达草稿、已发布、归档。
- 文章使用 `visibility` 表达公开、私有。
- 文章、评论和附言使用 `restriction` 表达管理限制。

禁止：
- 重新引入单字段 `status` 同时表达发布、可见、管理限制和删除。
- 使用 `hidden` 字段替代 `visibility` 或 `restriction` 的业务语义。

例外：
- action record、outbox、inbox 等过程表可以使用自己的状态枚举。

检查：
- 扫描 content 文章、评论、附言 schema 和 proto 是否重新出现混合语义 `status`。
- 扫描 content 枚举是否重新出现删除状态。

## RULE-012 数据变换优先使用 lo

必须：
- slice、map、set 的去重、映射、过滤、分组、键值提取等纯数据变换优先使用 `github.com/samber/lo`。

禁止：
- 为简单数据变换手写冗长循环。
- 在有副作用、错误短路或事务操作的流程中强行使用 lo。

例外：
- 显式循环更易读或需要提前返回错误时保留显式循环。

检查：
- 审查新增循环是否只是去重、映射、过滤、分组或提取字段。
- 审查 lo 使用处是否包含副作用或错误短路需求。

## RULE-013 基础类型指针使用 new(expr)

必须：
- 需要基础类型指针时使用 Go 1.26 的 `new(expr)`。

禁止：
- 使用 `common/pkg/util.Ptr`。
- 为取地址创建临时变量。

例外：
- 第三方 API 要求特殊指针类型时按 API 要求处理。

检查：
- 扫描 `util.Ptr`。
- 审查新增基础类型临时变量取地址。

## RULE-014 注释只描述意图和约束

必须：
- 注释描述业务规则、跨服务边界、事务边界、事件语义和数据库约束的意图。
- 枚举注释使用 `// xxx`，不使用编号列表。

禁止：
- 注释复述代码表面行为。
- 使用无关历史、临时说明或情绪化表达。

例外：
- 生成代码注释由生成器决定。

检查：
- 审查新增注释是否只是在复述函数名或字段赋值。
- 扫描枚举注释是否使用 `// 1.`、`// 2.` 这类编号。

## RULE-015 生成和校验必须按变更类型执行

必须：
- 修改 proto 后运行共享 API 生成和 lint。
- 修改 Ent schema 后运行对应模块 Ent 生成。
- 修改 wire、config、SDK 或 OpenAPI 注释后运行对应 Make 目标。
- 修改 usecase、repo、service 后运行受影响模块测试或编译。

禁止：
- 手改生成代码。
- 绕过 Makefile 直接拼装生成命令。

例外：
- 临时定位问题可以直接运行底层命令，但最终验证必须回到 Make 目标。

检查：
- 查看最终回复中的验证命令。
- 查看生成文件是否和源文件变更匹配。
