# Bass dev CI/CD

`cicd` 目录用于在本地 Kubernetes 集群运行 Woodpecker CI。当前目标是本地过渡使用：本地 `microk8s` 负责 UI、流水线调度和构建，远程 `vm-microk8s` 测试集群只负责部署运行。

## 文件

```text
cicd/
|-- README.md
|-- secrets.example.yaml      # forge/OAuth Secret 示例，不直接部署
|-- woodpecker-nodeport.yaml  # 本地固定 NodePort 入口
`-- woodpecker-values.yaml    # Woodpecker Helm values
```

## 部署

先切到本地集群：

```bash
kubectl config use-context microk8s
```

创建 namespace：

```bash
kubectl create namespace cicd
```

复制 `secrets.example.yaml`，填入本地 OAuth 配置后再 apply。包含真实密钥的 `secrets.local.yaml` 不应提交到 Git。

```bash
kubectl apply -f common/ops/k8s/dev/cicd/secrets.local.yaml
```

安装 Woodpecker：

```bash
helm upgrade --install woodpecker oci://ghcr.io/woodpecker-ci/helm/woodpecker \
  --version 3.3.0 \
  -n cicd \
  --timeout 10m \
  -f common/ops/k8s/dev/cicd/woodpecker-values.yaml \
  --wait

kubectl apply -f common/ops/k8s/dev/cicd/woodpecker-nodeport.yaml
```

## 访问

优先使用 port-forward，本机最稳定：

```bash
kubectl port-forward -n cicd svc/woodpecker-server 31980:80
```

访问地址：

```text
http://127.0.0.1:31980
```

`woodpecker-nodeport.yaml` 也固定了 `31980` 端口，但在 Windows + 本地 MicroK8s 场景下，NodePort 是否能从 Windows 直接访问取决于本地集群网络实现。

如果需要给团队成员访问，建议通过 Tailscale、Cloudflare Tunnel 或 frp 暴露本地 `31980`，不要直接裸露到公网。

## 后续流水线

流水线配置建议放到仓库 `.woodpecker/` 下，按服务手动触发：

```text
SERVICE=bbs
SERVICE=user
SERVICE=content
SERVICE=notify
SERVICE=im
SERVICE=platform
SERVICE=scheduler
```

构建完成后推送到阿里云 ACR，再通过 SSH 或受限 kubeconfig 部署到远程 `vm-microk8s` 的 `bass-test` namespace。
