# 代码模板索引

新增同类代码时，先使用本目录模板，再参考同模块已有实现。模板只描述结构和边界，不要求逐字复制。

| 模板 | 使用场景 |
|------|----------|
| [proto-service.md](proto-service.md) | 新增内部服务或 BFF proto RPC。 |
| [service.md](service.md) | 新增 Go service 接收者方法。 |
| [usecase.md](usecase.md) | 新增业务用例方法。 |
| [repo.md](repo.md) | 新增 biz repo 接口和 data repo 实现。 |
| [ent-schema.md](ent-schema.md) | 新增 Ent schema。 |

通用要求：

- 模板不替代业务设计；涉及表结构、事件契约、跨服务依赖或对外接口时，先确认设计。
- 模板中的 `XXX` 只表示占位名称，实际代码使用业务语义命名。
- 生成代码后按 [Agent 可执行规则](../agent-rules.md) 检查。
