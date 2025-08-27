# Kubernetes Resource Limits vs. Node Capacity

## ❓ Scenario

- Pod `requests` are **less than** node resources ✅
- Pod `limits` are **more than** node resources ❌
- No additional nodes available

### ✅ Will Kubernetes schedule the pod?

**Yes**, Kubernetes *will schedule the pod* as long as the `requests` fit on a node — even if the `limits` exceed the node's total capacity.

### 🔍 Why?

- Kubernetes **scheduler only uses `requests`** (not `limits`) to place pods.
- `limits` are enforced **at runtime** by the container runtime (e.g. containerd or CRI-O).

### 🧪 Example

**Node capacity:**

- CPU: `2`
- Memory: `4Gi`

**Pod resources:**

```yaml
resources:
  requests:
    cpu: "1"
    memory: "1Gi"
  limits:
    cpu: "3"
    memory: "5Gi"
```

**Result:**

- Pod is scheduled ✅
- If the pod tries to use more than 2 CPUs or 4Gi memory:
  - **CPU:** throttled (not killed)
  - **Memory:** likely **OOMKilled**

### ⚠️ Risks

- Overcommitting **CPU** → Performance degradation (throttling)
- Overcommitting **memory** → Risk of **pod crashes (OOMKilled)**

### ✅ Best Practice

- Set realistic `requests` and `limits` based on actual usage.
- Monitor usage with:

  ```bash
  kubectl top pod
  kubectl describe pod <pod-name>
  ```

- Use **Vertical Pod Autoscaler** for auto-tuning.
