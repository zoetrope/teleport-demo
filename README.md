
## Install kubectl

See [https://kubernetes.io/docs/tasks/tools/install-kubectl/](https://kubernetes.io/docs/tasks/tools/install-kubectl/) .

## Create Kubernetes cluster

```console
$ kind create cluster --name teleport-demo --config cluster.yaml 
$ export KUBECONFIG="$(kind get kubeconfig-path --name="teleport-demo")"
```

## Install cert-manager

```console
$ kubectl create namespace cert-manager
$ kubectl label namespace cert-manager certmanager.k8s.io/disable-validation="true"
$ kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.10.0/cert-manager.yaml
```

## Generate certificate

```console
$ kubectl apply -f certificate.yaml
```

## Install teleport

```console
$ git clone https://github.com/gravitational/teleport.git $GOPATH/src/github.com/gravitational/teleport
$ kubectl create namespace teleport
$ helm install --name teleport --namespace teleport -f values.yaml $GOPATH/src/github.com/gravitational/teleport/examples/chart/teleport/
```

## Setup GitHub Integration

```yaml
kind: github
version: v3
metadata:
  name: github
spec:
  client_id: <client-id>
  client_secret: <client-secret>
  display: Github
  redirect_url: https://teleport.example.com:3080/v1/webapi/github/callback
  teams_to_logins:
    - organization: <your-organization>
      team: <your-team>
      logins:
        - root
      kubernetes_groups: ["system:masters"]
```

```console
$ kubectl -n teleport exec -it teleport-xxx bash
$ tctl create github.yaml
```

## Install teleport CLI

[https://gravitational.com/teleport/download/](https://gravitational.com/teleport/download/)

## Login

```
$ tsh login --insecure --proxy=teleport.example.com:3080 --auth=github
```
