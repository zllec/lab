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

