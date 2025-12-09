# Kubernetes Consolidated Notes

*Consolidated from all K8s learning materials - August 2025*

## Setup and Installation

### kubectl Installation

The Kubernetes command-line tool for interacting with clusters.

#### 1. Download the latest kubectl binary

```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```

**What this does:** Downloads the latest stable version of kubectl for Linux AMD64.

#### 2. Validate the download (security best practice)

```bash
# Download checksum file
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"

# Validate the file
echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
```

**Expected output:** `kubectl: OK`

#### 3. Install kubectl

```bash
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

**Verify installation:**
```bash
kubectl version --client
```

### Minikube Setup

Minikube runs a local Kubernetes cluster for development and learning.

#### Installation

```bash
# Download Minikube
curl -LO https://github.com/kubernetes/minikube/releases/latest/download/minikube-linux-amd64

# Install it
sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64
```

#### Start your first cluster

```bash
minikube start
```

This will:
- Download the Kubernetes cluster image
- Start a virtual machine or container
- Configure kubectl to talk to the cluster

#### Access the dashboard (optional but cool)

```bash
minikube dashboard
```

Opens a web-based Kubernetes dashboard where you can see your cluster visually.

## Basic Configuration and Context

```bash
# configure kubectl with the appropriate context to interact with a k8s cluster
kubectl config use-context minikube

# check if kubectl is installed
kubectl version

# check the current cluster
kubectl cluster_info

# list cluster's nodes
kubectl get nodes

# shows current config 
kubectl config current-context

# get namespaces
kubectl get namespaces

# rename context
kubectl config rename-context <cluster_name> <new_name>
```

## Core Concepts

### Pod
- **smallest unit in Kubernetes**
- can contain single, multi, or init containers
- has assigned networking and storage(Volumes)
- **Smallest deployable unit** in Kubernetes
- Contains one or more containers that share storage and network
- Usually you don't create pods directly

```bash
# get running pods
kubectl get pods

# run a pod
kubectl run <pod_name> --image=nginx

# get pod details
kubectl describe pods <pod-name>

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

# access the pod 
kubectl exec -it nginx-test -- /bin/bash
```

### Deployment
- **Manages pods** - keeps them healthy and restarts them if they die
- **Recommended way** to create and scale pods
- Handles rolling updates and rollbacks
- Deployments are just wrappers for replica sets
- You directly work with Deployments

```bash
# create a deployment
## kubectl create deployment deployment_name --image=image
kubectl create deployment kubernetes-bootcamp --image=gcr.io/k8s-minikube/kubernetes-bootcamp:v1

# edit deployment
kubectl edit deployment synergychat-web

# get deployment in yaml format
kubectl get deployment synergychat-web -o yaml > web-deployment.yaml
```

### Service
- **Exposes pods** to network traffic
- Pods have internal IPs that change when they restart
- Services provide stable endpoints
- acts like a reverse proxy - load balancer and provides a stable endpoint
- when creating a new service, default type is ClusterIP if not specified
- there are 4 types of services: ClusterIP, NodePort, LoadBalancer, and ExternalName
- ClusterIP is just a way to expose the pods within the cluster
- NodePort and LoadBalancer if you want to expose to the outside world
- ExternalName is primarily is for DNS redirects

## Namespaces

- acts like a grouping, boundary
- deleting namespace deletes all resources assigned to it

```bash
# create a namespace
kubectl create namespace my_namespace

# create a namespace called test1
kubectl create ns test1

# create a manifest for a namespace
kubectl create ns test3 -o yaml --dry-run=client > test3-ns.yaml

# set the namespace for the current context
kubectl config set-context --current --namespace=test1

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

## Hands-On Tutorial Examples

### Create a deployment

```bash
kubectl create deployment hello-node --image=registry.k8s.io/e2e-test-images/agnhost:2.39 -- /agnhost netexec --http-port=8080
```

**What this creates:** A deployment running a simple web server on port 8080.

### Check what's running

```bash
# See deployments
kubectl get deployments

# See individual pods
kubectl get pods

# More detailed info
kubectl describe pod <pod-name>
```

### Expose the service to outside world

```bash
kubectl expose deployment hello-node --type=LoadBalancer --port=8080
```

**Key point:** `--type=LoadBalancer` makes the service accessible from outside the cluster.

### Access your service

```bash
# Get the URL (Minikube specific)
minikube service hello-node --url
```

Visit that URL in your browser - you should see the agnhost netexec page!

## Addons and Extensions

Minikube comes with useful addons for learning.

### List available addons

```bash
minikube addons list
```

### Enable metrics server (for resource monitoring)

```bash
minikube addons enable metrics-server
```

### Check what's running in the system namespace

```bash
kubectl get pod,svc -n kube-system
```

### View resource usage (after metrics-server is running)

```bash
kubectl top pods
```

### Disable an addon when done

```bash
minikube addons disable metrics-server
```

## Configuration Management

### ConfigMaps
- most common way to manage env variables
- Once a config map is created, you have to connect it to the target deployment
- this is not suitable for secrets
- instead of `env`, use `envFrom` to reference the whole config

### Services

### Gateway
- exposes Services outside the cluster
- install this:
`kubectl apply --server-side -f https://github.com/envoyproxy/gateway/releases/download/v1.5.1/install.yaml`

## Storage

### Storage
- containers are storage are ephemeral
- we use K8s volumes
- Persistent Volume allows you to persist data 

### Persistent Volume
- pretty much like ConfigMap when attaching to a deployment
- can be created statically or dynamically
- Static PVs are created manually by the admin
- Dynamic PVs are created automatically when a pod requests a volume that doesn't exist yet

### Persistent Volume Claim
-  **requests** for a PV
- attached to a pvc

## Resource Management and Limits

### Limits
- Deployment sets the allowed CPU/ RAM limit
- ConfigMap sets what the application will use even if it's way above the Deployment limit. The App will crash if not enough memory
- If limited by CPU, the app will go slower, if RAM limit has reached, it will crash
- Resource Request allows the Deployment to let the Control Plane look for a node that can accommodate the limit and will fail if no node can be found
- Limits are for protection
- Requests are for scheduling
- Memory is scarier than CPU

### Good rule of thumb:
- Set memory requests ~10% higher than the average memory usage of your pods
- Set CPU requests to 50% of the average CPU usage of your pods
- Set memory limits ~100% higher that the average memory usage of your pods
- Set CPU limits ~100% higher that the average CPU usage of your pods

## Debugging and Troubleshooting

```bash
# kubectl proxy
# starts a proxy service
# http://127.0.0.1:8001/api/v1/namespaces/default/pods
# shows the pod details
kubectl proxy
```

### Common Issues
- "Thrashing" - a pod that keeps crashing and restarting
- CrashLoopBackOff - a container keeps exiting with non zero exit code

## Gotchas and Important Notes

- Pods can have one or more containers inside but since they are in the same pods, having same port containers will not be allowed
- Manifests uses metadata.name to apply to the target

## Cleanup

When you're done experimenting:

```bash
# Remove the service and deployment
kubectl delete service hello-node
kubectl delete deployment hello-node

# Stop the cluster (keeps everything for next time)
minikube stop

# Or completely delete the cluster
minikube delete
```

## Exam and Productivity Tips

- use `apiVer` for searching the docs
- `:set paste` in vim to preserve proper format

```bash
# pipe help result to Vim or less for better search and navigation
kubectl config --help | vim 
kubectl config --help | less
```

## Key Takeaways

- **One container per pod** in basic setups (though pods can have multiple containers)
- **LoadBalancer services** enable external access to your applications
- **Addons provide extra functionality** without complex setup
- **Minikube makes local development** incredibly easy
- **kubectl is your main interface** for everything Kubernetes
- For container orchestration - deployment, scaling in/out, decommission
- `kubectl` provides an API to run commands against Kubernetes cluster. Deploy containers, view and manage cluster, and view logs.

## Why do we need it?
For container orchestration - deployment, scaling in/out, decommission

`kubectl` provides an API to run commands against Kubernetes cluster. Deploy containers, view and manage cluster, and view logs.