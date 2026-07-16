# Bass K8s Infrastructure

`common/ops/k8s` 保存 Bass 基础设施的 Kubernetes 配置。基础组件通过 Helm 安装，外部入口由 Traefik 提供。

## 目录结构

```text
k8s/
├── README.md
├── base/        # Consul、PostgreSQL、Redis、NATS、MinIO Helm values
├── ingress/     # Traefik Helm values
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

`dev/test/prod/secrets.yaml` 按组件拆分 Secret。

| Secret | key | 组件 |
|---|---|---|
| `bass-postgres-secret` | `password` | PostgreSQL |
| `bass-redis-secret` | `redis-password` | Redis |
| `bass-minio-secret` | `rootUser`、`rootPassword` | MinIO |
| `bass-nats-secret` | `NATS_USER`、`NATS_PASSWORD` | NATS |

Consul ACL bootstrap token 由 Consul chart 创建。dev 环境读取命令：

```bash
kubectl get secret bass-consul-bootstrap-acl-token -n bass-dev -o jsonpath='{.data.token}' | base64 -d
```

## dev 部署

```bash
kubectl apply -k ./dev
helm upgrade --install bass-consul hashicorp/consul -n bass-dev --create-namespace --timeout 5m -f ./base/consul.yaml
helm upgrade --install bass-pg bitnami/postgresql -n bass-dev --create-namespace -f ./base/postgres.yaml
helm upgrade --install bass-redis bitnami/redis -n bass-dev --create-namespace -f ./base/redis.yaml
helm upgrade --install bass-nats nats/nats -n bass-dev --create-namespace -f ./base/nats.yaml
helm upgrade --install bass-minio minio/minio -n bass-dev --create-namespace -f ./base/minio.yaml
```

`test` 和 `prod` 使用相同命令，将命名空间和目录分别替换为 `bass-test`/`./test`、`bass`/`./prod`。

## Traefik

Traefik 独立安装在 `traefik` 命名空间：

```bash
helm upgrade --install traefik traefik/traefik -n traefik --create-namespace -f ./ingress/traefik-values.yaml --wait
```

业务环境的路由资源放在各自命名空间。

| 环境 | HTTP 路由 | TCP 路由 |
|---|---|---|
| dev | `dev/ingress.yaml` | `dev/ingress-tcp.yaml` |
| test | `test/ingress.yaml` | `test/ingress-tcp.yaml` |
| prod | `prod/ingress.yaml` | `prod/ingress-tcp.yaml` |

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
192.168.100.10 consul.dev.bass.local minio.dev.bass.local s3.dev.bass.local nats.dev.bass.local postgresql.dev.bass.local redis.dev.bass.local
```

HTTP 服务：

| 服务 | 地址 |
|---|---|
| Consul UI | `http://consul.dev.bass.local` |
| MinIO Console | `http://minio.dev.bass.local` |
| MinIO S3 | `http://s3.dev.bass.local` |
| NATS Monitor | `http://nats.dev.bass.local` |

TCP 服务：

| 服务 | 地址 |
|---|---|
| PostgreSQL | `postgresql.dev.bass.local:5432` |
| Redis | `redis.dev.bass.local:6379` |
| NATS | `nats.dev.bass.local:4222` |

## 生产环境

- 部署前替换所有示例密钥。
- 生产域名和 TLS 证书单独配置。
- PostgreSQL 和 MinIO 需要配置独立备份策略。
- monitoring PVC 当前使用本地默认 StorageClass，其他集群需要替换为实际 StorageClass。
