# Bass dev CI/CD

`cicd` directory runs Jenkins in the local Kubernetes cluster. The current goal is a local, UI-driven CI/CD transition: manually choose a service module in Jenkins, pull code, build an image, then push to the registry or deploy to the remote `vm-microk8s` test cluster.

## Files

```text
cicd/
|-- README.md
|-- jenkins.yaml          # Jenkins controller, PVC, services, NodePort
`-- secrets.example.yaml  # credential placeholders, do not deploy directly
```

## Deploy

Switch to the local cluster first:

```bash
kubectl config use-context microk8s
```

Deploy Jenkins:

```bash
kubectl apply -f common/ops/k8s/dev/cicd/jenkins.yaml
kubectl rollout status deployment/jenkins -n cicd --timeout=5m
```

## Access

Prefer port-forward on Windows + local MicroK8s:

```bash
kubectl port-forward -n cicd svc/jenkins 31980:8080
```

Open:

```text
http://127.0.0.1:31980
```

`jenkins.yaml` also reserves NodePort `31980`. Direct Windows access to NodePort depends on the local Kubernetes networking implementation.

Initial admin password:

```bash
kubectl exec -n cicd deploy/jenkins -- cat /var/jenkins_home/secrets/initialAdminPassword
```

## Manual Pipeline Shape

Create a parameterized Jenkins job with one service parameter:

```text
SERVICE=bbs
SERVICE=user
SERVICE=content
SERVICE=notify
SERVICE=im
SERVICE=platform
SERVICE=scheduler
```

This phase does not require GitHub webhooks, so Jenkins does not need to be publicly reachable. The pipeline should use Jenkins credentials to pull the repository manually.

Recommended credentials to add in Jenkins UI later:

```text
Git token or SSH key
Aliyun ACR username/password
Remote test kubeconfig or SSH key
```

## Resources

Jenkins controller uses a fixed image tag:

```text
dockerproxy.net/jenkins/jenkins:2.516.3-lts-jdk21
```

Resource profile:

```text
requests: 200m CPU / 768Mi memory
limits:   1 CPU / 1536Mi memory
PVC:      10Gi
```

Jenkins is heavier than Woodpecker. For real Go and image builds, prefer dedicated build pods or remote build nodes instead of putting build load inside the controller process.
