# 运行时治理

## Kratos

- server、service、wire、middleware 按模块统一组织，新增组件必须进入对应 ProviderSet。
- 常规 middleware 包括 recovery、logging、tracing、auth、validate、timeout；BFF 和内部服务按职责启用。
- `RegisterGrpc` / `RegisterHttp` 只做协议注册，不写业务逻辑。

## 配置

- 新增配置必须同步更新配置 proto、yaml 示例、默认值和使用说明。
- secret、token、密钥不写入 Git；通过环境变量或密钥系统注入。
- database、redis、consul、nats 配置按用途归类，不在业务代码硬编码。

## 安全

- BFF 负责入口认证和粗粒度权限，归属服务负责领域权限最终判断。
- 身份信息通过统一 context key 传递，禁止字符串 key 散落。
- 内部 RPC 是否信任调用方必须有明确边界；需要服务间鉴权时统一在 middleware 实现。

## 事件与缓存

- 生产者在本地事务内写 outbox；消费者用 inbox 做幂等、重试、补偿。
- 事件包含全局幂等 ID、事件类型、subject、生产者和必要聚合信息。
- subject、consumer group、durable 名称要稳定、可追踪。
- Redis 只用于限流、短期幂等、短期缓存、验证码/会话临时态、防刷计数。
- 缓存必须有 TTL，不能作为领域事实来源。

## 可观测性

- 日志记录关键 ID、动作、状态变化、耗时和错误原因，不记录敏感信息。
- RPC、MQ、任务调度边界应传递 trace/request 上下文。
- 关键写流程、outbox/inbox、缓存、限流、下游调用需要指标。
