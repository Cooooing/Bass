# Bass

Bass 是一个面向社区产品的 Go 服务化平台。它将客户端协议适配、领域事实、实时投递与基础设施集成分离，让各个领域能够独立演进，而不共享业务数据库。

本仓库提供后端平台。Docker Compose 中可选的 `treble` Web 客户端由独立项目构建和发布。

## 项目能力

- 社区：账户、认证、资料、内容、评论、标签、互动、通知、私信、徽章与积分。
- 平台：文件存储、回调、第三方集成、IP 归属地、定时与延迟任务。
- 实时：连接注册、路由、WebSocket 投递和基于 NATS 的异步处理。

## 架构

<svg viewBox="0 0 1570 842" width="100%" xmlns="http://www.w3.org/2000/svg" role="img" aria-labelledby="architecture-title">
  <title id="architecture-title">Bass 社区平台架构</title>
  <defs>
    <pattern id="hatch" width="8" height="8" patternUnits="userSpaceOnUse" patternTransform="rotate(45)">
      <line x1="0" y1="0" x2="0" y2="8" stroke="#e2e8f0" stroke-width="2"/>
    </pattern>
  </defs>
  <rect width="1570" height="842" fill="#ffffff"/>
  <text x="785" y="48" text-anchor="middle" fill="#0f172a" font-family="Arial, sans-serif" font-size="30">Bass 社区平台架构</text>

  <g fill="url(#hatch)" stroke="#cbd5e1" stroke-width="2">
    <rect x="30" y="88" width="1210" height="88" rx="18"/>
    <rect x="30" y="200" width="1210" height="150" rx="18"/>
    <rect x="30" y="374" width="1210" height="170" rx="18"/>
    <rect x="30" y="568" width="1210" height="120" rx="18"/>
    <rect x="30" y="712" width="1210" height="110" rx="18"/>
  </g>

  <g fill="#64748b" font-family="Arial, sans-serif" font-size="18">
    <text x="50" y="116">客户端</text>
    <text x="50" y="228">接入层</text>
    <text x="50" y="402">业务领域</text>
    <text x="50" y="596">平台领域</text>
    <text x="50" y="740">持久化</text>
  </g>

  <g fill="url(#hatch)" stroke="#64748b" stroke-width="2">
    <rect x="545" y="102" width="180" height="60" rx="12"/>
  </g>
  <text x="635" y="126" text-anchor="middle" fill="#1e293b" font-family="Arial, sans-serif" font-size="15">SSR Client</text>
  <text x="635" y="146" text-anchor="middle" fill="#1e293b" font-family="Arial, sans-serif" font-size="15">Treble</text>

  <g fill="url(#hatch)" stroke="#38bdf8" stroke-width="2">
    <rect x="190" y="231" width="360" height="90" rx="16"/>
    <rect x="224" y="263" width="292" height="46" rx="12"/>
    <rect x="720" y="243" width="360" height="64" rx="12"/>
  </g>
  <g fill="#0284c7" font-family="Arial, sans-serif">
    <text x="210" y="258" font-size="17">BFF</text>
  </g>
  <g text-anchor="middle" fill="#1e293b" font-family="Arial, sans-serif" font-size="14">
    <text x="370" y="283">论坛</text><text x="370" y="302">bff_bbs</text>
    <text x="900" y="268" font-size="15">长连接节点</text><text x="900" y="288" font-size="15">realtime_node</text>
  </g>

  <g fill="url(#hatch)" stroke="#4ade80" stroke-width="2">
    <rect x="140" y="414" width="260" height="52" rx="12"/><rect x="490" y="414" width="260" height="52" rx="12"/><rect x="840" y="414" width="260" height="52" rx="12"/>
    <rect x="140" y="478" width="260" height="52" rx="12"/><rect x="490" y="478" width="260" height="52" rx="12"/><rect x="840" y="478" width="260" height="52" rx="12"/>
  </g>
  <g text-anchor="middle" fill="#1e293b" font-family="Arial, sans-serif" font-size="14">
    <text x="270" y="436">用户模块</text><text x="270" y="454">user</text>
    <text x="620" y="436">内容模块</text><text x="620" y="454">content</text>
    <text x="970" y="436">即时通讯模块</text><text x="970" y="454">im</text>
    <text x="270" y="500">通知模块</text><text x="270" y="518">notify</text>
    <text x="620" y="500">经济模块</text><text x="620" y="518">economy</text>
    <text x="970" y="500">徽章模块</text><text x="970" y="518">badge</text>
  </g>

  <g fill="url(#hatch)" stroke="#f59e0b" stroke-width="2">
    <rect x="140" y="603" width="260" height="54" rx="12"/><rect x="505" y="603" width="260" height="54" rx="12"/><rect x="870" y="603" width="260" height="54" rx="12"/>
  </g>
  <g text-anchor="middle" fill="#1e293b" font-family="Arial, sans-serif" font-size="15">
    <text x="270" y="625">平台模块</text><text x="270" y="645">platform</text>
    <text x="635" y="625">任务调度模块</text><text x="635" y="645">scheduler</text>
    <text x="1000" y="625">实时协调模块</text><text x="1000" y="645">realtime_hub</text>
  </g>

  <g fill="url(#hatch)" stroke="#f87171" stroke-width="2">
    <rect x="375" y="739" width="220" height="56" rx="12"/><rect x="675" y="739" width="220" height="56" rx="12"/>
  </g>
  <g text-anchor="middle" fill="#1e293b" font-family="Arial, sans-serif" font-size="16"><text x="485" y="773">PostgreSQL</text><text x="785" y="773">MinIO</text></g>

  <g fill="url(#hatch)" stroke="#c084fc" stroke-width="2">
    <rect x="1280" y="200" width="280" height="488" rx="18"/>
    <rect x="1312" y="272" width="216" height="52" rx="12"/><rect x="1312" y="370" width="216" height="52" rx="12"/><rect x="1312" y="468" width="216" height="52" rx="12"/><rect x="1312" y="566" width="216" height="52" rx="12"/>
  </g>
  <text x="1300" y="228" fill="#7c3aed" font-family="Arial, sans-serif" font-size="17">运行时基础设施</text>
  <g text-anchor="middle" fill="#1e293b" font-family="Arial, sans-serif" font-size="16"><text x="1420" y="304">Consul</text><text x="1420" y="402">Redis</text><text x="1420" y="500">NATS</text><text x="1420" y="598">DTM</text></g>
</svg>

每个领域服务独占自己的业务规则和写入职责。服务通过共享 gRPC 契约进行同步协作，并通过本地 outbox 支撑的 NATS 事件完成异步协作；服务之间不共享业务数据库，也不导入其他服务的 `internal` 包。

| 范围 | 服务 | 职责 |
| --- | --- | --- |
| 客户端入口 | `bff_bbs` | HTTP 协议适配、认证与读模型聚合 |
| 社区领域 | `user`、`content`、`economy`、`notify`、`im`、`badge` | 社区产品的业务事实与规则 |
| 平台能力 | `platform`、`scheduler` | 文件、外部集成、IP 数据、定时和延迟工作 |
| 实时能力 | `realtime_hub`、`realtime_node` | 连接路由与消息投递 |
| 共享契约 | `common` | Protobuf、错误、枚举、客户端与基础设施辅助能力 |

## 仓库结构

```text
Bass/
├─ app/          # 可独立运行的后端服务
├─ common/       # 共享 API 契约与无业务含义的基础设施代码
├─ monolith/     # 已选模块的单进程组合
├─ ops/          # Docker Compose、Kubernetes 模板和运维工具
├─ build/        # Make 片段、Dockerfile 和 CI 辅助工具
└─ .run/         # IntelliJ IDEA 运行配置（可选）
```

单个服务遵循相同的分层布局：

```text
app/<service>/
├─ cmd/             # 进程入口与依赖注入
├─ configs/         # 启动与服务配置模板
├─ internal/service/ # HTTP/gRPC 协议适配
├─ internal/biz/     # 业务规则与事务编排
├─ internal/data/    # 数据库、缓存、RPC 和事件实现
└─ module/           # 服务组合
```

## 快速开始

### 前置条件

- Docker Engine 和 Docker Compose v2
- Go `1.27`
- GNU Make 与 Bash

以下步骤用于创建本地开发与集成测试环境。为所有 `<FILL_...>` 填写本地强密钥；生产环境使用独立的密钥与部署配置。

### 1. 创建本地配置

在仓库根目录执行：

```bash
cp ops/docker/.middleware.env.example ops/docker/.middleware.env
cp ops/docker/.env.example ops/docker/.env
```

保持 `ops/docker/.middleware.env` 与 `ops/docker/.env` 中 PostgreSQL、Redis、NATS 和 MinIO 的凭据一致。

所有 `.env` 文件都由 Git 忽略。以已提交的 `*.env.example` 为配置清单，将本地凭据保存在对应的 `.env` 文件中。

### 2. 启动基础设施

```bash
cd ops/docker
docker compose --env-file .middleware.env -f compose-middleware.yaml up -d
docker compose --env-file .middleware.env -f compose-middleware.yaml ps
```

该命令会在 `bass` Docker 网络中启动 Consul、PostgreSQL、Redis、NATS JetStream、MinIO 和 DTM。启动完成后，服务通过该网络获得配置、数据存储、缓存、事件、对象存储与分布式事务能力。

### 3. 配置流与 Consul 同步

配置按服务启动模式加载：

- `dev`：bootstrap 从服务本地 `configs/config.yaml` 读取配置。Bash 加载 `ops/docker/.env`，统一为所有本地服务提供数据库、缓存、消息和对象存储等通用连接配置。
- `test`、`prod`：bootstrap 从 Consul KV 的 `config/<service>/config.yaml` 读取配置。
- `ops/tools/cmd/consul-config`：批量解析开发环境中各服务的 `config.yaml`，并同步到 Consul，为联调、测试与部署准备统一配置。

同步工具读取 `ops/tools/cmd/consul-config/.env` 与进程环境变量；进程环境变量优先于同步工具环境文件中的同名键。部署环境通过 CI/CD 或密钥管理系统提供完整环境变量。

本地联调时，先复制并填写同步工具的环境文件，再从开发环境变量启动同步：

```bash
cp ops/tools/cmd/consul-config/.env.example ops/tools/cmd/consul-config/.env
```

启动基础设施后执行一次；本地环境变量变化后再次执行：

```bash
cd ../..
set -a && source ops/docker/.env && set +a && (cd ops/tools && go run ./cmd/consul-config)
```

该工具会解析每个 `app/*/configs/config.yaml` 中的环境变量并写入 Consul。由于加载顺序，命令行显式导入的 `ops/docker/.env` 会覆盖同步工具模板中的同名变量，从而为源码启动的服务发布 localhost 端点。

### 4. 从源码启动服务

首次使用时全局执行一次 `make init` 安装锁定版本的代码生成工具。随后生成代码并启动需要开发的服务；以下示例启动社区 BFF：

```bash
make init                         # 首次使用时全局执行一次
make gen
set -a && source ops/docker/.env && set +a && (cd app/bff_bbs && go run ./cmd)
```

打开 `http://127.0.0.1:8500` 查看 Consul 注册，以确认服务已启动并加入本地环境。

启动其他服务时，将 `bff_bbs` 替换为 `app/` 下对应目录即可。`.run/*.run.xml` 为 IntelliJ IDEA 提供相同的环境变量加载方式，命令行启动与 IDE 启动使用同一份开发环境变量。

### 启动已发布的应用镜像

基础设施启动且配置同步后，可以启动已发布的应用镜像：

```bash
cd ops/docker
cp .service.env.example .service.env
docker compose -f compose-service.yaml up -d
```

`compose-service.yaml` 会将 `.service.env` 注入应用容器；`compose-middleware.yaml` 使用 `.middleware.env` 生成中间件的端口、凭据和事件配置。

### 可选：启动链路追踪

```bash
cd ops/docker
docker compose -f compose-observability.yaml up -d
```

追踪栈依赖中间件 Compose 项目创建的外部 `bass` 网络。

## 日常开发命令

```bash
make init                         # 首次使用时全局安装锁定版本的生成工具
make api                          # 生成共享 Protobuf 产物
make gen                          # 生成所有服务产物
make build                        # 构建所有服务二进制
make lint                         # 格式化并检查所有服务
make -C app/user gen              # 生成一个服务的产物
make -C app/bff_bbs doc           # 生成 BFF OpenAPI 文档
make -C app/bff_bbs sdk           # 生成 BFF SDK
```

构建修改过的服务前，请执行对应的生成命令，生成契约、Ent 代码和 Wire 输出。

## 运行注意事项

- 将本地凭据保存在被忽略的 `.env` 文件或密钥管理服务中；配置、日志、提交和问题单使用脱敏信息。
- 生产环境配置独立的密钥、备份、TLS、访问控制、容量限制和镜像更新策略。
- 当前对象存储由 MinIO 提供；`ops/docker/compose-middleware.yaml` 中记录了后续对象存储中间件迁移的 TODO。
- `ops/k8s/dev/` 提供可版本化的开发集群模板；在应用到 Kubernetes 集群前请阅读其[部署说明](ops/k8s/dev/README.md)。

## 延伸阅读

- [Docker 中间件定义](ops/docker/compose-middleware.yaml)
- [应用 Compose 定义](ops/docker/compose-service.yaml)
- [Kubernetes 开发环境](ops/k8s/dev/README.md)
- 本地存在 `doc/` 目录时，可阅读其中的[开发事实与实现规则](doc/README.md)
