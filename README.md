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

### 服务划分

#### 接入层 (Access Layer)

> 统一入口

* gateway (API 网关)
    * 统一入口：所有 HTTP 请求入口
    * 安全能力：JWT 鉴权、限流、CORS
    * 路由转发：将请求分发至业务逻辑层的各个微服务
* connector (长连接网关)
    * 连接管理：WebSocket 建立、心跳、断线
    * 数据透传：不感知业务，仅转发消息
    * 接入校验：基于 signal 下发的临时 Ticket 校验

#### 业务逻辑层 (Business Logic Layer)

> 业务逻辑实现

* user (用户服务)
    * 账号体系：注册、登录、第三方登录、双因素认证
    * 认证鉴权：JWT 生成、权限
    * 用户关系：关注、拉黑
* content (内容服务)
    * 内容管理：文章、评论、领域、标签
* notify (通知服务)
    * 通知管理：通知模板、元数据、记录
* im (即时通讯服务)
    * 会话管理：私聊 / 群聊
    * 消息逻辑：存储、撤回、已读未读、漫游
    * 群组管理：成员、权限
* economy (经济服务)
    * 暂空

#### 通用能力层 (Capability Layer)

> 跨业务的平台能力

* infra (基础设施服务)
    * 外部 API 代理
* signal (信令服务)
    * 连接调度：为节点进行认证、为客户端分配最合适的 `connector` 节点地址（负载均衡）
    * 状态管理：管理 `connector` 节点状态、用户与节点关系、用户状态
    * 消息路由：接受消息队列消息，推送到 `connector`

### 项目结构约定

~~~bash
Bass/                                        # 项目根目录
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
