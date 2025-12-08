# Notes

```bash
# configure kubectl with the appropriate context to interact with a k8s cluster
kubectl config use-context minikube
```

## Stateless Applications

```bash
# create a namespace called test1
kubectl create ns test1

# create a manifest for a namespace
kubectl create ns test3 -o yaml --dry-run=client > test3-ns.yaml

# set the namespace for the current context
kubectl config set-context --current --namespace=test1

# access the pod 
kubectl exec -it nginx-test -- /bin/bash
```

- kubectl edit deployment synergychat-web
- opens the deployment file in editor
- kubectl proxy
  - starts a proxy service
    - http://127.0.0.1:8001/api/v1/namespaces/default/pods
        - shows the pod details
- Deployments are just wrappers for replica sets
- You directly work with Deployments
- `kubectl get deployment synergychat-web -o yaml > web-deployment.yaml`
  - creates a copy of deployment in a yml format
- "Thrashing" - a pod that keeps crashing and restarting
- CrashLoopBackOff - a container keeps exiting with non zero exit code

### ConfigMaps
- most common way to manage env variables
- Once a config map is created, you have to connect it to the target deployment
- this is not suitable for secrets
- instead of `env`, use `envFrom` to reference the whole config

### Services
- acts like a reverse proxy - load balancer and provides a stable endpoint
- when creating a new service, default type is ClusterIP if not specified
- there are 4 types of services: ClusterIP, NodePort, LoadBalancer, and ExternalName
- ClusterIP is just a way to expose the pods within the cluster
- NodePort and LoadBalancer if you want to expose to the outside world
- ExternalName is primarily is for DNS redirects

### Gateway
- exposes Services outside the cluster
- install this:
`kubectl apply --server-side -f https://github.com/envoyproxy/gateway/releases/download/v1.5.1/install.yaml`

### Storage
- containers are storage are ephemeral
- we use K8s volumes
