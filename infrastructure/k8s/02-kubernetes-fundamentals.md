# Why do we need it?
For container orchestration - deployment, scaling in/out, decommission

`kubectl` provides an API to run commands against Kubernetes cluster. Deploy containers, view and manage cluster, and view logs.

`kubectl config current-context` shows current config 

```bash
kubectl get namespaces
kubectl config rename-context <cluster_name> <new_name>

```
```
```

# Pod
- smallest unit in kubernetest
- can contain single, multi, or init containers
- has assigned networking and storage(Volumes)

```bash
# get running pods
kubectl get pods

# run a pod
kubectl run <name> --image=nginx

# get pod details
kubectl describe pods <pod-name>
```

# Namespaces
- acts like a grouping, boundary
- deleting namespace deletes all resources assigned to it

```bash
# create a namespace
kubectl create namespace my_namespace

# create a pod in the namespace
kubectl run my_httpd --image=httpd --namespace=my_namespace

# get pods in the namespaces
kubectl get pods --n my_namespace

# create pod in the namespace
kubectl run my_namespace_podname --image=name --namespace=my_namespace

# delete pods in a namespaces
kubectl delete pods -n my_namespace my_namespace_podname

# deletes a namespaces
kubectl delete namespaces my_namespace
```

# Linux Tips 

```bash
# pipe help result to Vim or less for better search and navigation
kubectl config --help | vim 
kubectl config --help | less
```

