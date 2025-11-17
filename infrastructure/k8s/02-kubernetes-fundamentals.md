# Why do we need it?
For container orchestration - deployment, scaling in/out, decommission

`kubectl` provides an API to run commands against Kubernetes cluster. Deploy containers, view and manage cluster, and view logs.

`kubectl config current-context` shows current config 

```bash
kubectl get namespaces
kubectl config rename-context <cluster_name> <new_name>
```

# Cluster

```bash

# check if kubectl is installed
kubectl version

# check the current cluster
kubectl cluster_info

# list cluster's nodes
kubectl get nodes
```

# Pod
- smallest unit in Kubernetes
- can contain single, multi, or init containers
- has assigned networking and storage(Volumes)

```bash
# get running pods
kubectl get pods

# run a pod
kubectl run <pod_name> --image=nginx

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

# Working with pods

```bash
# this "tries" to run the pod and outputs either yaml or json
# see --dry-run help for more info
# Why do this? to quickly generate a manifest file
kubectl run nginx-yaml --image=nginx --dry-run=client -o yaml

# run the yaml file and provision the contents
kubectl apply -f nginx-yaml.yaml

# describe to see changes
kubectl describe pods nginx-yaml

# get inside the container
kubectl exec -it name_of_the_pod -- /bin/bash
```
