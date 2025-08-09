# K8S beginning

- [K8S beginning](#k8s-beginning)
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
  - [StatefulSet](#statefulset)
  - [What is a StatefulSet?](#what-is-a-statefulset)
  - [Features of StatefulSet](#features-of-statefulset)
  - [When to Use StatefulSet](#when-to-use-statefulset)
  - [Example StatefulSet Manifest](#example-statefulset-manifest)
  - [Job](#job)
  - [Key Features of a Job](#key-features-of-a-job)
  - [Types of Jobs](#types-of-jobs)
  - [Use Cases for Jobs in Kubernetes](#use-cases-for-jobs-in-kubernetes)
  - [Example Of Job Manifest](#example-of-job-manifest)
  - [CronJob](#cronjob)
  - [How Kubernetes CronJob Works](#how-kubernetes-cronjob-works)
  - [Key Components](#key-components)
  - [Use Cases of Kubernetes CronJobs](#use-cases-of-kubernetes-cronjobs)
  - [Example Of CronJon Manifest](#example-of-cronjon-manifest)
  - [Understanding the Schedule Format (cron Expression)](#understanding-the-schedule-format-cron-expression)
  - [Concurrency Handling (.spec.concurrencyPolicy)](#concurrency-handling-specconcurrencypolicy)
  - [Service](#service)
  - [Why do we need Services?](#why-do-we-need-services)
  - [Types of Kubernetes Services](#types-of-kubernetes-services)
    - [1. ClusterIP (Default)](#1-clusterip-default)
    - [2. NodePort](#2-nodeport)
    - [3. LoadBalancer](#3-loadbalancer)
    - [4. ExternalName](#4-externalname)
  - [How Services Work?](#how-services-work)
  - [🔹 Bonus: Headless Services](#-bonus-headless-services)
  - [Summary](#summary)
  - [Manifest Examples](#manifest-examples)

## Kubernetes Manifest

A Kubernetes manifest is a YAML (or JSON) file that defines the desired state of a Kubernetes object. It is used to create, update, and manage resources like Pods, Deployments, Services, ConfigMaps, etc.

### Why Use Manifests?

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

for more detail in given fields we can use dot and then the pods.<FIELD_NAME>:

```sh
    kubectl explain pods.metadata
```

Now, let's explore and review some of these Kubernetes API resources in detail.  

## Pods

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

## ReplicaSet

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

## Deployment  

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

## DaemonSet  

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

## StatefulSet

StatefulSet is one of the most important workload resources in Kubernetes, especially when you’re dealing with apps that need stable network identity, stable storage, and ordered deployment.

## What is a StatefulSet?

  A StatefulSet is a Kubernetes resource used to manage stateful applications.

  While Deployment and ReplicaSet are used for stateless applications (e.g., web servers), StatefulSet is used when each Pod has to maintain state or has unique configuration, identity, and storage.

## Features of StatefulSet

1. Stable & Unique Pod Names  
Pods are named in an ordered pattern:
pod-name-0, pod-name-1, pod-name-2, ...

2. Stable Network Identity  
Each Pod gets a unique DNS hostname like:
pod-name-0.myservice

3. Persistent Storage (Optional but common)  
Each Pod can have its own PersistentVolumeClaim (PVC).  
The volume is not deleted when the Pod is deleted.

4. Ordered Deployment & Scaling  
Pods are started and terminated in order (0 → 1 → 2).  
Ensures ordered rollouts and rollbacks.

5. Ordered Rolling Updates  
Updates happen one Pod at a time, maintaining order.

## When to Use StatefulSet

Use StatefulSet when you need:  

- Databases (e.g., MySQL, MongoDB, Cassandra)
- Queues like Kafka, RabbitMQ
- Distributed systems that need unique identities and persistent storage

## Example StatefulSet Manifest

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: redis
  namespace: default
  labels:
    app.kubernetes.io/env: development
    app.kubernetes.io/name: redis
    app.kubernetes.io/project: test
spec:
  serviceName: redis
  replicas: 3
  selector:
    matchLabels:
      app.kubernetes.io/name: redis
      app.kubernetes.io/env: development
  template:
    metadata:
      labels:
        app.kubernetes.io/name: redis
        app.kubernetes.io/env: development
    spec:
      containers:
        - name: redis
          image: redis:latest
          volumeMounts:
            - name: redis-data
              mountPath: /var/lib/redis
  volumeClaimTemplates:
    - metadata:
        name: redis-data
      spec:
        storageClassName: local-path
        accessModes:
          - ReadWriteOnce
        resources:
          requests:
            storage: 1Gi
```

To check PersistentVolumeClaims:

```sh
kubectl get pvc
```

## Job  

a Job is a resource that is used to run a one-time or batch task until completion. Unlike Deployments, which manage long-running applications, a Job ensures that a specific number of pods successfully complete their tasks.

## Key Features of a Job

1. Ensures Completion: A Job creates one or more pods and runs them until they complete successfully.

2. Automatic Restart: If a pod fails, Kubernetes will restart it based on the Job’s restart policy.

3. Parallel Execution: You can run multiple pods in parallel to complete a batch of work.

## Types of Jobs

1. **Single Job (default)** – Runs a single pod until it succeeds.

2. **Parallel Job** – Runs multiple pods in parallel for faster execution.

3. **Work Queue Job** – Uses a queue-based system where each pod picks up tasks from a queue.  

## Use Cases for Jobs in Kubernetes

- Data processing tasks (ETL jobs)

- Database migrations

- Scheduled tasks (when combined with a CronJob)

- Automated testing jobs

## Example Of Job Manifest

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: greater
spec:
  # ttlSecondsAfterFinished: 100
  template:
    spec:
      restartPolicy: Never
      containers:
      - name: greater
        image: alpine:latest
        command: ["false"]
```  

To check jobs:

```sh
kubectl get job
```

## CronJob  

A Kubernetes (k8s) CronJob is used to run scheduled tasks in a Kubernetes cluster, similar to how you schedule jobs using `cron` in Linux. It allows you to run a job at a specific time or interval.  

## How Kubernetes CronJob Works

A CronJob in Kubernetes creates a Job object on a schedule you define. Each Job runs a Pod that performs the task and then terminates.

## Key Components

1. Schedule (.spec.schedule) – Defines when the job should run using a cron expression.

2. Job Template (.spec.jobTemplate) – Specifies the actual job that will run.

3. Concurrency Policy (.spec.concurrencyPolicy) – Controls how overlapping jobs are handled.

4. Starting Deadline (.spec.startingDeadlineSeconds) – Specifies how long a job can be delayed before being skipped.

5. Successful Jobs History (.spec.successfulJobsHistoryLimit) – Limits how many successful job records are kept.

6. Failed Jobs History (.spec.failedJobsHistoryLimit) – Limits how many failed job records are kept.

## Use Cases of Kubernetes CronJobs

- ✅ Database backups – Run a backup job at midnight daily.
- ✅ Log rotation – Clean up old logs every week.
- ✅ Automated reports – Generate and send reports on a schedule.
- ✅ Data processing – Process batches of data periodically.

## Example Of CronJon Manifest  

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: test
spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          restartPolicy: Never
          containers:
            - name: test
              image: alpine:latest
              command: ["echo", "hello world!"]
```

for learning schedule patterns there is a good [website](https://crontab.guru/)

## Understanding the Schedule Format (cron Expression)

```sql
 ┌───────────── minute (0-59)
 │ ┌───────────── hour (0-23)
 │ │ ┌───────────── day of month (1-31)
 │ │ │ ┌───────────── month (1-12)
 │ │ │ │ ┌───────────── day of week (0-6, Sunday=0 or 7)
 │ │ │ │ │
 * * * * *    (Every minute)
```

## Concurrency Handling (.spec.concurrencyPolicy)

1. `Allow` (default) → Allows multiple jobs to run simultaneously.

2. `Forbid` → Prevents a new job from starting if the previous job is still running.

3. `Replace` → Stops the current job and replaces it with a new one.

```yaml
spec:
  concurrencyPolicy: Forbid
```

## Service

A **Service** is an abstraction that defines a stable way to access a set of **Pods**. Since Pods are ephemeral and their IP addresses change frequently, a **Service** provides a fixed IP and DNS name, ensuring reliable communication.

## Why do we need Services?

- Pods are dynamic; their IP addresses change when restarted or rescheduled.
- We need a consistent way for other services, applications, or external users to access these Pods.
- Load balancing is often required across multiple Pods.

---

## Types of Kubernetes Services

### 1. ClusterIP (Default)

- Exposes the service internally within the cluster.
- Cannot be accessed from outside the cluster.
- Useful for internal communication between microservices.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-clusterip-service
spec:
  selector:
    app: my-app
  ports:
    - protocol: TCP
      port: 80    # Service Port
      targetPort: 8080  # Container Port
```

### 2. NodePort

- Exposes the service externally using a port on each Node.
- The service is accessible via `NodeIP:NodePort`.
- Not ideal for production as the port range is limited (30000–32767).

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-nodeport-service
spec:
  type: NodePort
  selector:
    app: my-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
      nodePort: 30080  # Optional, auto-assigned if omitted
```

### 3. LoadBalancer

- Creates an external load balancer (requires cloud provider support, e.g., AWS, GCP, Azure).
- Exposes the service to the internet.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-loadbalancer-service
spec:
  type: LoadBalancer
  selector:
    app: my-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
```

### 4. ExternalName

- Maps a Kubernetes Service to an external DNS name instead of selecting Pods.
- Used when you need to refer to external services (e.g., a database hosted outside Kubernetes).

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-external-service
spec:
  type: ExternalName
  externalName: example.com
```

---

## How Services Work?

- Services use **selectors** to find matching Pods (though it's optional).
- Kubernetes assigns a **ClusterIP** and **DNS name** (e.g., `my-service.default.svc.cluster.local`).
- **Kube-proxy** helps route traffic from the service to the correct Pod using **iptables** or **IPVS**.

---

## 🔹 Bonus: Headless Services

- If you set `clusterIP: None`, Kubernetes does not provide a virtual IP.
- Instead, it returns a list of Pod IPs when queried via DNS.
- Useful for applications that handle their own load balancing (e.g., databases like Cassandra).

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-headless-service
spec:
  clusterIP: None
  selector:
    app: my-app
  ports:
    - port: 80
```

---

## Summary

| Service Type     | Accessibility | Use Case |
|-----------------|--------------|----------|
| ClusterIP (default) | Internal only | Microservice communication |
| NodePort | External (NodeIP:Port) | Development/testing, direct access |
| LoadBalancer | External via cloud LB | Production, internet-facing services |
| ExternalName | External DNS name | Referring to external services |
| Headless (`clusterIP: None`) | Direct pod IPs | Stateful applications |

## Manifest Examples  

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  type: NodePort
  selector:
    app: nginx
  ports:
    - name: http
      port: 80
      targetPort: 80
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:alpine
```
