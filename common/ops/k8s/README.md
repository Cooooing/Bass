# Bass K8s Infrastructure

Bass 基础设施层的 K8s 配置。基础组件通过 Helm Charts 部署，Ingress 使用 Traefik，监控栈使用原生 K8s YAML。

## 目录结构

```text
k8s/
├── README.md
├── base/        # Consul、PostgreSQL、Redis、NATS、MinIO Helm values
├── ingress/     # Traefik Controller Helm values
├── dev/         # namespace: bass-dev
├── test/        # namespace: bass-test
├── prod/        # namespace: bass
└── monitoring/  # namespace: monitoring
```

## Helm 仓库

```bash
helm repo add hashicorp https://helm.releases.hashicorp.com
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add nats https://nats-io.github.io/k8s/helm/charts/
helm repo add minio https://charts.min.io/
helm repo add traefik https://traefik.github.io/charts
helm repo update
```

## Secret

`dev/test/prod/secrets.yaml` 按组件拆分 Secret：

| Secret                 | key                         | 使用方        |
|------------------------|-----------------------------|------------|
| `bass-postgres-secret` | `password`                  | PostgreSQL |
| `bass-redis-secret`    | `redis-password`            | Redis      |
| `bass-minio-secret`    | `rootUser`、`rootPassword`   | MinIO      |
| `bass-nats-secret`     | `NATS_USER`、`NATS_PASSWORD` | NATS       |

`monitoring/secrets.yaml` 使用 `monitoring-grafana-secret`，包含 `GRAFANA_ADMIN_USER` 和 `GRAFANA_ADMIN_PASSWORD`。

Kubernetes 允许自定义 Secret key，但使用方读取什么 key，Secret 中就必须提供什么 key。PostgreSQL、Redis 的 key 可以通过 values 配置；MinIO chart 固定读取 `rootUser` 和 `rootPassword`；NATS 通过环境变量引用 Secret，再在 NATS 配置中使用环境变量。

不同 namespace 中同名 Secret 不冲突。本目录仍按组件拆分 Secret，避免多个组件共用一个 Secret 后出现 key 混杂。

## dev 部署

```bash
kubectl apply -k ./dev
helm upgrade --install bass-consul hashicorp/consul -n bass-dev --create-namespace --timeout 5m -f ./base/consul.yaml
helm upgrade --install bass-pg bitnami/postgresql -n bass-dev --create-namespace -f ./base/postgres.yaml
helm upgrade --install bass-redis bitnami/redis -n bass-dev --create-namespace -f ./base/redis.yaml
helm upgrade --install bass-nats nats/nats -n bass-dev --create-namespace -f ./base/nats.yaml
helm upgrade --install bass-minio minio/minio -n bass-dev --create-namespace -f ./base/minio.yaml
```

`test` 和 `prod` 使用相同命令，把命名空间和目录分别替换为 `bass-test`/`./test`、`bass`/`./prod`。

## Traefik

`ingress/traefik-values.yaml` 只用于安装 Traefik Ingress Controller。Controller 独立部署在 `traefik` 命名空间。

每个业务环境的路由资源仍放在各自命名空间：

| 环境   | HTTP 路由             | TCP 路由                  |
|------|---------------------|-------------------------|
| dev  | `dev/ingress.yaml`  | `dev/ingress-tcp.yaml`  |
| test | `test/ingress.yaml` | `test/ingress-tcp.yaml` |
| prod | `prod/ingress.yaml` | `prod/ingress-tcp.yaml` |

Ingress 管理外部流量进入集群的入口，不管理服务访问外部系统的出口。

首次安装 Traefik：

```bash
helm upgrade --install traefik traefik/traefik -n traefik --create-namespace -f ./ingress/traefik-values.yaml --wait
```

## 监控部署

```bash
kubectl apply -k ./monitoring
```

## 状态检查

```bash
kubectl get pods,svc,pvc -n bass-dev
kubectl get ingressroute,ingressroutetcp -n bass-dev
kubectl get pods,svc,pvc -n monitoring
helm list -n bass-dev
```

## 服务访问

在 `C:\Windows\System32\drivers\etc\hosts` 添加本地集群节点 IP：

```text
192.168.100.10 consul.dev.bass.local minio.dev.bass.local s3.dev.bass.local nats.dev.bass.local
```

| 服务            | 地址                             |
|---------------|--------------------------------|
| Consul UI     | `http://consul.dev.bass.local` |
| MinIO Console | `http://minio.dev.bass.local`  |
| MinIO S3      | `http://s3.dev.bass.local`     |
| NATS Monitor  | `http://nats.dev.bass.local`   |

TCP 服务直接使用节点 IP 和标准端口：

| 服务         | 地址                    |
|------------|-----------------------|
| PostgreSQL | `192.168.100.10:5432` |
| Redis      | `192.168.100.10:6379` |
| NATS       | `192.168.100.10:4222` |

## 生产环境

- 部署前替换所有示例密码。
- 监控 PVC 当前使用本地默认 `microk8s-hostpath`，其他集群需要替换为实际 StorageClass。
- 生产域名和 TLS 证书需要单独配置。
- PostgreSQL 和 MinIO 需要独立备份策略。
