# 🧠 Shard-Aware Routing in ScyllaDB

**Shard-aware routing** in **ScyllaDB** is a performance optimization that ensures client requests are sent directly to the correct **CPU core (shard)** responsible for a given piece of data, avoiding unnecessary hops between cores.

---

## 🔍 Background: Shards in ScyllaDB

- **ScyllaDB** is a high-performance, distributed NoSQL database built in C++ using the **Seastar framework**.
- It achieves concurrency by **sharding data across CPU cores**. Each core (shard) independently manages a subset of the data.
- There is **no shared memory** between shards—inter-shard communication has overhead, so it's better to avoid it when possible.

---

## 🚦 What is Shard-aware Routing?

**Shard-aware routing** means the ScyllaDB **driver** (client) knows how to:

1. **Calculate the target shard** responsible for a specific partition key on a given node.
2. **Open connections to all shards** (cores) on a node.
3. **Send each request directly to the correct shard**, rather than relying on the node to forward the request internally.

This is done using the **token** derived from the partition key and the known **shard distribution** on the node.

---

## 📈 Why It Matters

Without shard-aware routing:

- Every request goes to a **random shard** on the node.
- If it's not the correct shard, the request must be **forwarded internally**, causing latency and CPU overhead.

With shard-aware routing:

- Requests hit the **right shard directly**, resulting in:
  - **Lower latency**
  - **Higher throughput**
  - **More predictable performance**

---

## ⚙️ How It Works (Simplified)

1. Scylla uses **Murmur3** hash of the partition key to determine the **token**.
2. The token is mapped to a **shard ID** on a node.
3. The driver uses this mapping to:
   - Choose the right **connection** (one per shard).
   - Send the query directly to the correct **CPU core**.

This requires:

- The Scylla driver to be **shard-aware**.
- Knowledge of **shard count**, **shard ID mappings**, and **token ranges**.

---

## ✅ Shard-aware Drivers

The following drivers support shard-aware routing with ScyllaDB:

- **Java** (official)
- **Go**
- **Python**
- **C++**
- **Node.js**

They maintain multiple connections per node (1 per shard) and route requests efficiently.

---

## 🔧 Enabling It

It's typically **enabled by default** in modern Scylla drivers. You may need to:

- Use **ScyllaDB-specific drivers** (not generic Cassandra ones).
- Ensure your application is **not behind a proxy** that obscures the original client routing (like some load balancers).

---

## 📚 Summary

| Feature             | Without Shard-aware | With Shard-aware |
|---------------------|---------------------|------------------|
| Latency             | Higher               | Lower            |
| CPU Efficiency      | Poor (due to hops)   | High             |
| Throughput          | Lower                | Higher           |
| Core Communication  | Frequent             | Minimal          |

---
