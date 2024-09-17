
## Prerequisites

Refer to the following document to install aqua CLI.

https://aquaproj.github.io/docs/install

Install tools with aqua.

```console
$ aqua install
```

Install teleport CLI

https://gravitational.com/teleport/download/

```console
$ wget https://cdn.teleport.dev/teleport-v16.3.0-linux-amd64-bin.tar.gz   
$ tar xvf teleport-v16.3.0-linux-amd64-bin.tar.gz
```

## Create Kubernetes cluster

```console
$ kind create cluster --name teleport-demo --config cluster.yaml 
```

## Deploy teleport

ref. https://goteleport.com/docs/admin-guides/deploy-a-cluster/helm-deployments/kubernetes-cluster/

```console
$ helm repo add teleport https://charts.releases.teleport.dev
$ helm repo update
```

```console
helm template teleport --namespace teleport teleport/teleport-cluster \
  --create-namespace \
  --version 16.3.0 \
  --values teleport-cluster-values.yaml \
  > manifests/teleport-cluster.yaml
```

```console
$ kubectl create namespace teleport
$ kustomize build ./manifests/ | kubectl apply -f -
```

## Create user

```console
$ kubectl exec -i -n teleport deployment/teleport-auth -- tctl create -f < member.yaml
$ kubectl exec -ti -n teleport deployment/teleport-auth -- tctl users add myuser --roles=member,access,editor
```

```console
$ tsh login --proxy=localhost:3080 --user=myuser --insecure
$ tsh ssh --proxy=localhost:3080 --insecure cybozu@node-demo-0
```

## Add node

Generate a token to join the cluster for a teleport node.

```console
$ kubectl exec -ti -n teleport deployment/teleport-auth -- tctl tokens add --type=node
```

Put the token to `teleport.auth_token` in `./manifests/teleport-node.yaml`.

Deploy a teleport node.

```console
$ kubectl apply -f ./manifests/teleport-node.yaml
```

## Use API

Create a user for API access.

```console
$ kubectl exec -i -n teleport deployment/teleport-auth -- tctl create -f < api-access.yaml

$ kubectl exec -ti -n teleport deployment/teleport-auth -- tctl users add api-access --roles=api-access
```

Generate a token for API access.

```console
$ tsh login --proxy=localhost:3080 --user=api-access --insecure --ttl=5256000

$ tctl --auth-server=localhost:3025 auth sign --ttl=87500h --user=api-access --out=client-demo/api-access.pem
```

Run a program to access the API.

```console
$ cd client-demo
$ go run main.go
```
