# 工程

## 编码

- 文本文件统一使用 UTF-8、无 BOM、LF 换行。
- 不要对无 BOM 文件执行 Remove BOM，避免误删文件前三个字符。
- 项目根目录 `.editorconfig` 是编辑器编码约束来源。

## 生成

- 生成代码、构建、测试前先阅读根目录和模块内 Makefile。
- 运行 make 相关命令时使用 bash 环境。
- 不手改生成代码；修改 proto、schema、wire、config 等源文件后运行生成命令。
- Proto 管理、lint 和生成统一走 Buf；共享 API 生成使用 `make api`，校验使用 `make api-lint`。
- Buf 生成模板统一放在 `common/buf`；共享 API 的 `buf.yaml` 和 `buf.lock` 放在 proto 模块目录 `common/api/app`。
- 服务内部 config proto 使用 Make 内联 Buf 配置，不为每个服务新增独立 Buf 配置文件。
- 第三方 proto 依赖统一通过 Buf deps 管理，不维护 `common/api/third_party`。
- 根级批量生成优先使用 `make gen-all`；单模块生成使用 `make -C app/<module> gen`。
- Ent 代码使用模块内 `make ent` 或 Makefile 中声明的等价命令生成。
- BFF OpenAPI 文档使用 `make -C app/<bff> doc` 生成。
- BFF SDK 使用 `make -C app/<bff> sdk-ts`、`make -C app/<bff> sdk-go` 或 `make -C app/<bff> sdk` 生成。
- SDK 使用 OpenAPI Generator CLI，需要 Node.js、npx 和 Java。
- CI 中 SDK 生成产物默认写入工作区的 `common/api/gen-ts/<bff>` 和 `common/api/gen-go/<bff>`；产物不提交到主业务分支，可作为 artifact 上传，或由 CI 推送到独立 SDK 分支/仓库。
- 根目录批量目标包括 `make doc-all`、`make sdk-validate-all`、`make sdk-all`。

## 依赖

- 依赖整理按模块执行，不在无关模块制造 `go.mod` / `go.sum` 变更。
- 新增依赖前优先复用项目已有库和公共工具。

## 测试

- 改 proto 后至少运行共享 API 生成，并编译受影响模块。
- 改 BFF proto 或 OpenAPI 注释后至少运行 `make -C app/<bff> doc`。
- 改 SDK 生成配置后至少运行 `make -C app/<bff> sdk-validate`；首次安装生成器可能需要网络。
- 改 Ent schema 后运行对应模块 Ent 生成，并编译受影响模块。
- 改 usecase、repo、service 后优先运行对应模块 `go test ./...`。
- 横切公共包变更需要扩大到所有受影响模块。

## 迁移

- 开发环境可使用自动迁移；生产环境迁移必须走显式变更流程。
- 破坏性 schema 变更需要拆分发布步骤，先兼容读写，再清理旧字段。
- 字段删除、重命名、唯一约束调整必须先评估历史数据和回滚方案。
