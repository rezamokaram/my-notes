- [Kubernetes Manifest](#kubernetes-manifest)
  - [Why Use Manifests?](#why-use-manifests)
  - [API Resources](#api-resources)
  - [Explain](#explain)
- [Pods](#pods)
  - [Key Characteristics of a Pod](#key-characteristics-of-a-pod)
  - [Applying The Manifest](#applying-the-manifest)
- [ReplicaSet](#replicaset)
  - [Key Features of ReplicaSet](#key-features-of-replicaset)
  - [ReplicaSet Example Yaml file](#replicaset-example-yaml-file)
  - [Comparison Between ReplicaSet \& Pod](#comparison-between-replicaset--pod)
  - [Difference Between ReplicaSet \& Deployment](#difference-between-replicaset--deployment)
- [Deployment](#deployment)
  - [Key Concepts of Kubernetes Deployment](#key-concepts-of-kubernetes-deployment)
  - [Deployment Example Yaml File](#deployment-example-yaml-file)
- [DaemonSet](#daemonset)
  - [How Does It Work?](#how-does-it-work)
  - [When Do We Use DaemonSet?](#when-do-we-use-daemonset)
  - [DaemonSet Manifest Example](#daemonset-manifest-example)
  - [Difference Between Deployment \& DaemonSet](#difference-between-deployment--daemonset)
  - [Quick Summary](#quick-summary)

# Kubernetes Manifest
A Kubernetes manifest is a YAML (or JSON) file that defines the desired state of a Kubernetes object. It is used to create, update, and manage resources like Pods, Deployments, Services, ConfigMaps, etc.

## Why Use Manifests?
- Allows declarative management (you define what you want, and Kubernetes ensures it happens).  
- Can be stored in version control (GitOps) for better tracking.  
- Supports automation and scalability.  

## API Resources  

The kubectl api-resources command helps you understand what resources are available in your Kubernetes cluster and how to use them in manifest files.  

```sh  
    kubectl api-resources  
```  
| `kubectl api-resources` Field | How It Helps in Manifest |
|-------------------------------|-------------------------|
| **APIVERSION** (`apps/v1`, `v1`) | Used in `apiVersion` field of the manifest. |
| **KIND** (`Deployment`, `Pod`) | Used in `kind` field of the manifest. |
| **NAMESPACED** (`true` or `false`) | If `true`, the resource must be in a namespace. |

## Explain  

The **`kubectl explain`** command provides detailed documentation about Kubernetes API resources and their fields. It's useful when you need to understand how to define or use resources in YAML manifests.

```sh
➜  ~ kubectl explain pods
KIND:       Pod
VERSION:    v1

DESCRIPTION:
    Pod is a collection of containers that can run on a host. This resource is
    created by clients and scheduled onto hosts.

FIELDS:
  apiVersion    <string>
    APIVersion defines the versioned schema of this representation of an object.
    Servers should convert recognized schemas to the latest internal value, and
    may reject unrecognized values. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources

  kind  <string>
    Kind is a string value representing the REST resource this object
    represents. Servers may infer this from the endpoint the client submits
    requests to. Cannot be updated. In CamelCase. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds

  metadata      <ObjectMeta>
    Standard object's metadata. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata

  spec  <PodSpec>
    Specification of the desired behavior of the pod. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status

  status        <PodStatus>
    Most recently observed status of the pod. This data may not be up to date.
    Populated by the system. Read-only. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
```
for more detail in given fields we can use dot and then the pods.<FIELD-NAME>:
```sh
    kubectl explain pods.metadata
```
Now, let's explore and review some of these Kubernetes API resources in detail.  

# Pods  

A **Pod** is the smallest and most basic deployable unit in Kubernetes. It represents a single instance of a running process in the cluster. A pod can contain one or more **containers**, which share the same network and storage resources.

## Key Characteristics of a Pod
1. **Single or Multiple Containers**  
   - Most commonly, a pod runs a single container (e.g., a web server or database).  
   - In some cases, multiple containers are deployed together in one pod to share storage and network resources.

2. **Shared Network and Storage**  
   - All containers in a pod share the **same IP address** and **port space**.  
   - They can communicate with each other using `localhost`.

3. **Lifecycle Management**  
   - Kubernetes manages the lifecycle of a pod. If a pod fails, it is replaced according to the **desired state** defined by the user.

4. **Short-lived and Ephemeral**  
   - Pods are not persistent. If a pod is deleted, a new pod gets created with a different IP.  
   - Stateful workloads require **PersistentVolumes (PVs)** for storage.

## Applying The Manifest

an example manifest:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test
  namespace: default
  labels:
    app.kubernetes.io/name: test
    app.kubernetes.io/env: production
    app.kubernetes.io/project: test
spec:
  containers:
    - name: nginx
      image: nginx:alpine
```

To deploy this pod to a Kubernetes cluster, save the YAML file (e.g., pod.yaml) and run:

```sh
    kubectl apply -f pod.yaml
```

To delete a specific pod, use:
```sh
kubectl delete pod <pod-name>
```
the pod name must match the pod name that you assign in manifest.  

If a pod is stuck in a terminating state, you can force delete it:

```sh
kubectl delete pod <pod-name> --grace-period=0 --force
```
⚠ Use this with caution, as it forcefully removes the pod without waiting for a graceful shutdown.  

If the pod was created using a YAML file (pod.yaml), delete it with:

```sh
kubectl delete -f pod.yaml
```

Delete Pods by Label Selector:
```sh
kubectl delete pods -l app=my-app
```
and so on.  

# ReplicaSet

A ReplicaSet (RS) is a Kubernetes resource that ensures a specified number of identical pods are always running. If a pod crashes or is deleted, the ReplicaSet automatically creates a new one to maintain the desired count.

## Key Features of ReplicaSet
- Ensures high availability by maintaining a fixed number of replicas
- Automatically recreates pods if they fail
- Uses selectors and labels to manage pods
- Does NOT support rolling updates (Use Deployment instead)  

## ReplicaSet Example Yaml file

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: test
  namespace: default
  labels:
    app.kubernetes.io/env: development
    app.kubernetes.io/name: test
    app.kubernetes.io/project: test
spec:
  replicas: 5
  selector:
    matchLabels:
      app.kubernetes.io/name: test
  template:
    metadata:
      labels:
        app.kubernetes.io/name: test
    spec:
      containers:
        - name: nginx
          image: nginx:alpine
```

To apply the ReplicaSet:
```sh
kubectl apply -f replicaset.yaml
```

Check the ReplicaSet:
```sh
kubectl get rs
```

For request inside the cluster:
```sh
kubectl run curlpod --image=alpine --restart=Never -it -- sh
```
Delete and Recreate the ReplicaSet:
```sh  
# Delete the old ReplicaSet
kubectl delete rs my-replicaset --cascade=orphan
# Apply the updated YAML file
kubectl apply -f replicaset.yaml
# Manually delete old pods
kubectl delete pod --selector=app=my-app
```  

To scale replicaset(make it 10 replicas and then 1):  
```sh
kubectl scale replicaset test --replicas 10
kubectl scale replicaset test --replicas 1
```

## Comparison Between ReplicaSet & Pod

| Feature             | Pod | ReplicaSet |
|--------------------|-----|------------|
| Ensures pod count  | ❌ No | ✅ Yes |
| Self-healing       | ❌ No | ✅ Yes |
| Automatically recreates failed pods | ❌ No | ✅ Yes |
| Manages multiple replicas | ❌ No | ✅ Yes |
| Uses labels & selectors | ❌ No | ✅ Yes |
| Supports scaling  | ❌ No | ✅ Yes |
| Preferred for production | ❌ No | ✅ Yes |

🔹 **Pods are the basic unit of deployment, but ReplicaSets ensure high availability by managing multiple replicas of a Pod.**


## Difference Between ReplicaSet & Deployment

| Feature         | ReplicaSet | Deployment |
|----------------|------------|------------|
| Ensures pod count | ✅ Yes | ✅ Yes |
| Self-healing | ✅ Yes | ✅ Yes |
| Supports rolling updates | ❌ No | ✅ Yes |
| Preferred for production | ❌ No | ✅ Yes |

🔹 **Use Deployment instead of ReplicaSet** unless you specifically need fine-grained control over pod scaling.

# Deployment  

A Kubernetes Deployment is a resource object in Kubernetes that provides declarative updates for applications. It helps in managing, scaling, and rolling out updates to applications efficiently.

## Key Concepts of Kubernetes Deployment

1. **Declarative Configuration**  
  You define the desired state of your application in a YAML or JSON file, and Kubernetes ensures the current state matches it.

2. **Replica Management**  
  Deployments allow you to run multiple replicas of your application (Pods), ensuring high availability and reliability.

3. **Rolling Updates & Rollbacks**  
  You can update your application with zero downtime using rolling updates.
  If something goes wrong, you can easily revert to a previous version using rollbacks.

4. **Self-healing**  
  If a Pod fails, Kubernetes automatically replaces it to maintain the desired number of replicas.

*Kubernetes Deployment works on top of a ReplicaSet to manage and control the lifecycle of Pods. Let me break it down:*

```scss
Deployment
   ├── ReplicaSet(Current)
   │      ├── Pod-1
   │      ├── Pod-2
   │      ├── Pod-3
   │
   ├── ReplicaSet(Old)[If an update was done]
          ├── Pod-4
          ├── Pod-5
```  
## Deployment Example Yaml File

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test
  namespace: default
  labels:
    app.kubernetes.io/env: development
    app.kubernetes.io/name: test
    app.kubernetes.io/project: test
spec:
  replicas: 10
  selector:
    matchLabels:
      app.kubernetes.io/name: test
      app.kubernetes.io/env: development
  template:
    metadata:
      labels:
        app.kubernetes.io/name: test
        app.kubernetes.io/env: development
    spec:
      containers:
        - name: nginx
          image: nginx:1.27
          resources:
            limits:
              memory: "512Mi"
              cpu: "500m"
            requests:
              memory: "256Mi"
              cpu: "250m"
```

To apply the manifest:  
```sh
kubectl apply -f deployment.yaml
```

Check the deployment status:
```sh
kubectl get deployments
```

Scale up/down:
```sh
kubectl scale deployment test --replicas=5
```

Change the image:
```sh
kubectl set image deployment/test my-app-container=new-image:latest
```

Rollback to a previous version:
```sh
kubectl rollout undo deployment test
```

Check the history of deployment revisions:
```sh
rollout history deployment test
```

Check the information of a specific revision:
```sh
kubectl rollout history deployment test --revision 4
```

Rollback to a specific revision of deployment:
```sh
kubectl rollout undo deployment test --to-revision 4
```

To run a command inside one of pods:
```sh
kubectl exec test-68b7569f6d-7q6ff -- nginx -v
```

# DaemonSet  

A DaemonSet is a special kind of Workload Resource in Kubernetes that ensures exactly *one Pod is running on every Node* in the cluster (or on a specific group of Nodes).

You can think of it like:  
`Run one copy of this Pod on every Node.`

## How Does It Work?
- When you create a DaemonSet, Kubernetes automatically schedules one Pod on every Node.

- If a new Node is added to the cluster, Kubernetes will automatically run a Pod on it.

- If a Node is removed, the Pod on that Node is deleted.

- You can also limit the DaemonSet to run on specific Nodes using node selectors or node affinity.

## When Do We Use DaemonSet?
DaemonSets are used when you want to run background or infrastructure services on every Node, such as:

- Log collectors → e.g., Fluentd, Filebeat

- Monitoring agents → e.g., Prometheus Node Exporter

- Security agents → e.g., Intrusion detection agents

- Network plugins → e.g., CNI plugins


## DaemonSet Manifest Example

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: node-exporter
  namespace: prometheus
  labels:
    app.kubernetes.io/part-of: monitoring
    app.kubernetes.io/name: node-exporter
    app.kubernetes.io/project: monitoring
spec:
  selector:
    matchLabels:
      app.kubernetes.io/name: node-exporter
      app.kubernetes.io/part-of: monitoring
  template:
    metadata:
      labels:
        app.kubernetes.io/name: node-exporter
        app.kubernetes.io/part-of: monitoring
    spec:
      containers:
        - name: node-exporter
          image: prom/node-exporter:latest
```

## Difference Between Deployment & DaemonSet

| **Deployment**                                      | **DaemonSet**                                      |
|----------------------------------------------------|----------------------------------------------------|
| Runs specified number of replicas anywhere in the cluster | Runs exactly **1 Pod per Node** (or specific Nodes) |
| Designed for **application workloads** (e.g., web servers, APIs) | Designed for **Node-level services** (e.g., log collectors, monitoring agents) |
| Supports **scaling by changing replicas**          | **Scaling is tied to the number of Nodes**         |
| New Pods can be scheduled on **any Node**          | Ensures there is always **one Pod on every Node**  |
| Commonly used for **frontend, backend, databases** | Commonly used for **logging agents, monitoring agents, network plugins** |
| Example: Run 5 replicas of Nginx                   | Example: Run 1 Fluentd agent on every Node         |

## Quick Summary
1. DaemonSet = 1 Pod per Node
2. Used for Node-level services
3. Auto-scales when Nodes are added or removed
4. Can be targeted to specific Nodes