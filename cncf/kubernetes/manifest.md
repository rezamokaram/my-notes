- [Kubernetes Manifest](#kubernetes-manifest)
  - [Why Use Manifests?](#why-use-manifests)
  - [API Resources](#api-resources)
  - [Explain](#explain)
- [Pods](#pods)
  - [Key Characteristics of a Pod](#key-characteristics-of-a-pod)
  - [Applying The Manifest](#applying-the-manifest)

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
the pod name must match the pod name that you assigne in manifest.  

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

