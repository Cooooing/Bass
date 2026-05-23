# 项目规范索引

这里放长期有效的设计与编码规范，正文尽量保持通用，不绑定具体模块名。

- [架构与一致性](architecture.md)
- [API、OpenAPI 与 SDK](api.md)
- [领域建模](domain.md)
- [运行时治理](runtime.md)
- [工程、生成与测试](engineering.md)

`AGENTS.md` 和 `CLAUDE.md` 是 Agent 的项目级提示入口；本目录用于给人阅读和补充细节。

## 当前模块归类

- `app/bbs`：BFF/API 体验层。
- `app/gateway`：网关层。
- `app/user`：用户域内部服务。
- `app/content`：内容域内部服务。
- `app/notify`：通知域内部服务。
- `app/im`：即时通信域内部服务。
- `common`：公共 proto、客户端、工具与共享约定。
