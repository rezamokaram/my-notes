# 🛡️ High Availability in ScyllaDB

High availability (HA) means the database remains operational and accessible even when some components (nodes, racks, or data centers) fail. ScyllaDB achieves HA using replication, fault tolerance, and recovery mechanisms.

---

## 🧱 ScyllaDB Architecture Overview

- **Clusters**: Groups of nodes working together.
- **Nodes**: Machines that store a portion of the data.
- **Data Centers**: Logical/physical groupings of nodes.
- **Partitions**: Units of data distributed across nodes.

---

## 🔄 How ScyllaDB Ensures High Availability

### 1. **Replication of Data**

- Data is copied to multiple nodes based on the **replication factor**.
- If one node fails, the data is still accessible from other replicas.

### 2. **Consistent Hashing**

- ScyllaDB uses consistent hashing to distribute data evenly.
- Prevents single points of failure and improves load balancing.

### 3. **Gossip Protocol**

- Nodes communicate their health/status to each other.
- Enables rapid detection of node failures.

### 4. **Tunable Consistency Levels**

- Example: With RF = 3, using `QUORUM` means 2/3 nodes must respond.
- This allows reads/writes to succeed even if one node is down.

### 5. **Hinted Handoff**

- If a replica is down during a write, the coordinator saves a **hint**.
- When the replica comes back, it **replays** missed updates.

### 6. **Read Repair & Anti-Entropy Repair**

- **Read Repair**: Fixes inconsistencies during reads.
- **Repair Tool**: Periodic job (e.g., via Scylla Manager) that syncs data across replicas.

### 7. **Multi Data Center Support**

- ScyllaDB supports `NetworkTopologyStrategy`.
- Replicates data across data centers for fault isolation and disaster recovery.
- If one data center fails, others can serve traffic.

---

## 💡 Example: Node Failure

**Cluster**: 3 nodes (Node A, B, C)  
**Replication Factor**: 3  
**Consistency Level**: QUORUM  

- If Node B fails:
  - Reads and writes continue using Nodes A and C.
  - Hints for Node B are stored.
  - When Node B returns, it catches up using hinted handoff.

---

## ✅ Benefits of HA in ScyllaDB

| Feature | Benefit |
|--------|---------|
| Replication | Protects against node/data center failure |
| Gossip Protocol | Fast detection of failures |
| Tunable Consistency | Balance between availability and consistency |
| Hinted Handoff | Recover from temporary outages |
| Read/Anti-Entropy Repair | Maintains consistency over time |
| Multi-DC Support | Geographic fault tolerance |
| No Master Node | No single point of failure |

---

## 🛠️ Best Practices for HA

- Use `NetworkTopologyStrategy` for multi-DC setups.
- Replication Factor ≥ 3.
- Use `QUORUM` for consistency in HA-sensitive systems.
- Run regular repairs (via Scylla Manager or `nodetool repair`).
- Use the Scylla Monitoring Stack for observability.

---

## high availability -> propagation data (if one data center goes down, entropy) (benefits of multiple data centers)
