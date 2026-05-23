# Claude 项目提示

## 工作方式

- 优先阅读并遵守 [doc/README.md](doc/README.md) 中的项目规范。
- 修改规范时，同步维护 `AGENTS.md`、`CLAUDE.md` 和对应 `doc/*.md`。
- 当前模块归类只写在 [doc/README.md](doc/README.md)，通用规范正文不要硬编码现有模块名。
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

- [架构与一致性](doc/architecture.md)
- [API、OpenAPI 与 SDK](doc/api.md)
- [领域建模](doc/domain.md)
- [运行时治理](doc/runtime.md)
- [工程、生成与测试](doc/engineering.md)

## 必守约束

- BFF 写接口通常只调用一个归属服务命令接口；跨服务副作用由归属服务写 outbox 后通过 MQ 触发。
- BFF 读接口可以聚合多个内部服务；BFF 不直接连接业务库，不拥有领域数据。
- 一个 proto 文件最多定义一个 service；每个 proto service 对应一个 Go service 文件。
- Proto 管理、lint 和生成统一使用 Buf；共享 API 生成使用 `make api`，校验使用 `make api-lint`。
- Buf 生成模板放在 `common/buf`，共享 API 模块配置放在 `common/api/app`。
- 第三方 proto 依赖统一通过 Buf deps 管理；不能继续维护 `common/api/third_party`。
- BFF proto 按模块功能分目录存放，必须定义自己的 request、reply 和对外模型，不引用内部服务 proto message。
- BFF HTTP 接口默认使用 `POST` 和 `body: "*"`；除图片、文件下载等必要场景外，不使用路径参数。
- BFF HTTP 路径统一使用 `/v1/{模块}/{资源}/{动作}`；同一资源名称保持单数且前后一致，动作使用 lower-kebab。
- BFF 不暴露邮箱、手机号、账号名是否存在等可用于枚举用户信息的接口。
- BFF OpenAPI 从 `google.api.http` 和 proto 注释生成；不手写 tags、servers、external_docs、BearerAuth。
- SDK 从 BFF OpenAPI 生成，输出到 `common/api/gen-ts/<bff>`、`common/api/gen-go/<bff>`，生成产物不入 Git 和 Docker。
- 常规表业务尽量保持 proto、service、biz、data、schema 一一对应。
- 复杂业务新增明确业务 usecase；usecase 可按层级组合调用，但禁止循环调用。
- 跨服务、跨层或落库约束枚举必须定义 proto enum，并通过 `common/pkg/enum.Mapping` 绑定内部 string enum。
- 复合查询条件使用 `Query`、`Filter`、`Spec` 等参数对象，不拼接超长方法名。
- 文本文件使用 UTF-8、无 BOM、LF；所有注释使用中文。
