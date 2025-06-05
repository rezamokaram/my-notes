# Token-Aware Concept in ScyllaDB

## ✅ What is Token-Aware?

**Token-aware** means that the client or driver sends a query **directly to the node that owns the data**, based on the **partition key's token**.

## 🔄 How It Works

1. The partition key is hashed into a **token** (e.g., using Murmur3).
2. Each node in the ScyllaDB cluster owns a **range of tokens**.
3. A **token-aware driver** calculates the token and sends the query to the **right node**, avoiding extra hops.

## 🛠️ Example

- Cluster has 3 nodes:
  - Node A: tokens `0–100`
  - Node B: tokens `101–200`
  - Node C: tokens `201–299`
- Partition key `"customer123"` hashes to token `115`
  - Non-token-aware client → sends to Node A → forwards to Node B
  - Token-aware client → sends **directly to Node B**

## 📈 Benefits

- ⚡ Lower latency
- 🔁 Fewer network hops
- 📉 Reduced load on coordinator nodes
- 🚀 Better performance under high load

## 🔧 Driver Support

Most ScyllaDB-supported drivers are token-aware by default:

- Java (`datastax` and `oss` drivers)
- Python (`cassandra-driver`)
- Go, Node.js, Rust, etc.
