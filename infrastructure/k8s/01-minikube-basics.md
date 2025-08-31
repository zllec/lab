# Kubernetes Learning Journey - Day 1

*Started: August 2025 - Learning Kubernetes from scratch*

My first dive into Kubernetes, starting with local development using Minikube. This is part of my infrastructure learning series, building on top of [VPS setup](../vps/) knowledge.

## Prerequisites

Before starting, make sure you have:
- A working VPS or local Linux environment
- [Docker installed](../vps/docker-install.md) (Minikube can use Docker as a driver)
- Basic command line familiarity

---

## kubectl Installation

The Kubernetes command-line tool for interacting with clusters.

### 1. Download the latest kubectl binary

```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```

**What this does:** Downloads the latest stable version of kubectl for Linux AMD64.

### 2. Validate the download (security best practice)

```bash
# Download checksum file
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"

# Validate the file
echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
```

**Expected output:** `kubectl: OK`

### 3. Install kubectl

```bash
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

**Verify installation:**
```bash
kubectl version --client
```

---

## Minikube Setup

Minikube runs a local Kubernetes cluster for development and learning.

### Installation

```bash
# Download Minikube
curl -LO https://github.com/kubernetes/minikube/releases/latest/download/minikube-linux-amd64

# Install it
sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64
```

### Start your first cluster

```bash
minikube start
```

This will:
- Download the Kubernetes cluster image
- Start a virtual machine or container
- Configure kubectl to talk to the cluster

### Access the dashboard (optional but cool)

```bash
minikube dashboard
```

Opens a web-based Kubernetes dashboard where you can see your cluster visually.

---

## Core Concepts I Learned

### Pod
- **Smallest deployable unit** in Kubernetes
- Contains one or more containers that share storage and network
- Usually you don't create pods directly

### Deployment
- **Manages pods** - keeps them healthy and restarts them if they die
- **Recommended way** to create and scale pods
- Handles rolling updates and rollbacks

### Service
- **Exposes pods** to network traffic
- Pods have internal IPs that change when they restart
- Services provide stable endpoints

---

## Hands-On Tutorial

Following the official Minikube tutorial to get my hands dirty.

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

---

## Addons - Extending Functionality

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

---

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

---

## Key Takeaways from Day 1

- **One container per pod** in basic setups (though pods can have multiple containers)
- **LoadBalancer services** enable external access to your applications
- **Addons provide extra functionality** without complex setup
- **Minikube makes local development** incredibly easy
- **kubectl is your main interface** for everything Kubernetes

## What's Next

- Learn about ConfigMaps and Secrets for configuration
- Explore persistent volumes for data storage
- Try multi-container pods
- Maybe tackle "Kubernetes The Hard Way" for deeper understanding

---

*This is part of my infrastructure learning journey. Previous: [Docker Installation](../vps/docker-install.md)*