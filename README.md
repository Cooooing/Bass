<p align = "center">
<br>
  <span style="font-size:20px; font-weight:bold;">
    Bass 社区系统，基于微服务架构设计，提供丰富的社区功能。
  </span>
<br>
</p>

## 目录

- [技术栈](#技术栈)
- [服务划分](#服务划分)
- [项目结构](#项目结构)
- [快速开始](#快速开始)

## 技术栈

- 语言 / 框架：Go + Kratos + Ent
- 中间件：
    - 数据库：PostgreSQL
    - 缓存：Redis
    - 消息队列：RabbitMQ
    - 配置中心 & 注册发现：Consul

## 服务划分

- gateway: HTTP 网关
- connector: WebSocket 长连接网关
- user: 用户服务
- content: 内容服务
- notify: 通知服务
- im: 即时通讯
- infra: 基础设施服务
- signal: 信令服务

## 项目结构

~~~bash
Bass/                          # 项目根目录
├─ app/                        # 微服务模块
│  ├─ module/                  # 子模块
│  │  ├─ cmd/                  # 启动文件（main.go、wire.go 等）
│  │  ├─ configs/              # 配置文件
│  │  ├─ internal/             # 模块内部代码
│  │  │  ├─ server/            # 服务创建和注册
│  │  │  ├─ service/           # 服务层，实现 API 逻辑
│  │  │  ├─ biz/               # 业务逻辑层
│  │  │  ├─ data/              # 数据访问（数据库、缓存等）
│  │  │  └─ conf/              # proto 配置定义
│  └─ other-modules/           # 其他模块，结构同上
├─ common/                      # 公共模块
│  ├─ api/                      # proto 接口定义
│  ├─ build_tools/              # 构建工具
│  ├─ pkg/                      # 通用 Go 包
│  └─ third_party/              # 第三方依赖 proto
├─ deploy/                      # 部署相关配置、脚本
├─ docs/                        # 文档
└─ ...
~~~

## 快速开始

### 使用 Docker 启动

环境变量在 `deploy/.env` 文件中，按实际需要修改。

```shell
cd deploy
docker-compose -f infra-compose.yml up -d
docker-compose -f service-compose.yml up -d
```
