<p align = "center">
<br>
  <span style="font-size:20px; font-weight:bold;">
    Bass 社区系统，基于微服务架构设计，提供丰富的社区功能。
  </span>
<br>
</p>

### 技术栈

- 语言 / 框架：Go + Kratos + Ent
- 中间件：
  - 数据库：PostgreSQL
  - 缓存：Redis
  - 消息队列：RabbitMQ
  - 配置中心：etcd

### 项目结构约定

~~~bash
BBS/                                         # 项目根目录
├─ app                                       # 项目子模块：各个微服务
│  ├─ module                                 # 子模块
│  │  ├─ cmd                                 # 启动文件目录（main.go、wire.go、wire_gen.go）
│  │  ├─ configs                             # 具体配置文件
│  │  │  ├─ bootstrap.yaml                   # 启动配置，定义服务基本信息（名称（用于服务注册发现）、版本、模式）、配置中心信息
│  │  │  └─ config.yaml                      # 运行时配置（本地调试使用，正式环境从配置中心拉取），包括数据库配置、缓存配置、服务注册中心配置、服务发现配置...
│  │  │
│  │  ├─ internal                            # 模块内部代码
│  │  │  ├─ server                           # 服务的创建和配置
│  │  │  │  ├─ grpc.go                       # 注册 gRPC 服务
│  │  │  │  └─ http.go                       # 注册 HTTP 服务（通过 google.api.http）
│  │  │  ├─ service                          # 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
│  │  │  │  └─ ...
│  │  │  ├─ biz                              # 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，而 repo 接口在这里定义，使用依赖倒置的原则。
│  │  │  │  ├─ model                         # 实体对象定义，对 ent 生成的模型进行封装，实现充血模型。
│  │  │  │  ├─ repo                          # 仓储接口，实现在 data（infra层）
│  │  │  │  └─ service                       # 具体业务逻辑，通常是涉及多个实体对象的一些操作和行为。
│  │  │  ├─ data                             # 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。
│  │  │  │  ├─ repo                          # biz repo 仓储接口实现
│  │  │  │  ├─ client                        # 数据库、缓存、消息队列等客户端
│  │  │  │  ├─ ent                           # ent 框架表模型定义和生成的代码
│  │  │  │  └─ ...
│  │  │  └─ conf                             # 配置文件 proto 定义
│  │  │     └─ ...
│  │  ├─Dockerfile                           # docker 镜像构建文件
│  │  ├─Makefile                             # 项目构建文件
│  │  ├─ go.mod
│  │  └─ go.sum
│  └─ other-modules                          # 其他模块，结构同上
├─ common/                                   # 公共模块
│  ├─ api/                                   # 所有 proto 接口定义
│  │  ├─ common/                             # 公共 proto
│  │  ├─ module/                             # 模块微服务 proto
│  │  └─ ...
│  ├─ build_tools/                           # 构建工具包
│  ├─ pkg/                                   # 通用 Go 包（工具、通用库）
│  │  ├─ constant/                           # 常量
│  │  ├─ util/                               # 公共工具函数
│  │  └─ ...
│  └─ third_party/                           # 第三方依赖 proto
│      ├─ errors/                            # 自定义错误规范 proto
│      ├─ google/                            # Google 官方 proto 依赖
│      ├─ openapi/                           # OpenAPI v3 schema
│      └─ validate/                          # protoc-gen-validate 插件 proto
├─ deploy/                                   # 部署环境相关配置、脚本
├─ docs/                                     # 文档
└...
~~~
