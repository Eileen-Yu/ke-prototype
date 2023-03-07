## Quick Start

### Prerequisite

Make sure you have:

- an accessible kubernetes cluster
- tetragon installed into your k8s cluster

This demo assumes a running node `tetragon-2rs6l`. You will probably need to change the value at [eventdispatcher/tetragon/constants.go](eventdispatcher/tetragon/constants.go).

### Get Start

#### 1. Deploy Tetragon by helm:

```sh
helm repo add cilium https://helm.cilium.io
helm repo update
helm install tetragon cilium/tetragon -n kube-system
kubectl rollout status -n kube-system ds/tetragon -w
```

#### 2. Deploy a Demo application:

You can use the official demo running a `xwing` pod:

```sh
kubectl create -f https://raw.githubusercontent.com/cilium/cilium/v1.11/examples/minikube/http-sw-app.yaml
```

Verify that the resource is running in the same node specified in [eventdispatcher/tetragon/constants.go](eventdispatcher/tetragon/constants.go).

```yaml
# ...
spec:
  containers:
    - image: docker.io/tgraf/netperf
      name: spaceship
  nodeSelector:
    kubernetes.io/hostname: pool-4lyf972z0-mv4hp
```

#### 3. Listen Tetragon Events from Log Lines

```sh
kubectl logs -n kube-system -l app.kubernetes.io/name=tetragon -c export-stdout -f
```

#### 4. In another process, execute commands in `xwing`

```sh
kubectl exec  -it xwing -- cat /etc/passwd
```

The tetragon events should be available in log lines

#### 5. Run this demo

```sh
# In the root of the repo
go run main.go
```

Repeat `Step 4`, the tetragon events should be caught by the demo and matched to the fake policy.

```json
{
  "execution_id": "cG9vbC00bHlmOTcyejAtbXY0aHA6NDI2MTQ3ODMyNzkzNjU2OToxNjEwMDcw",
  "pod": "xwing",
  "namespace": "default",
  "kernel_func": "fd_install",
  "time": "2023-03-07T03:14:05.398979552Z",
  "policy": "A"
}
```
