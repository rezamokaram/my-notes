# Cluster Topologies in Distributed Systems

## 🔷 Common Cluster Topologies

| Topology        | Definition                                                                 | Used In (Examples)                        |
|-----------------|----------------------------------------------------------------------------|-------------------------------------------|
| **Ring**        | Nodes are arranged in a logical circle where each is responsible for a portion of the data/token space. | ScyllaDB, Apache Cassandra                |
| **Star**        | All nodes connect to a central node that acts as a hub or coordinator.     | Traditional DB master-slave setups, Redis Cluster (to some extent) |
| **Mesh (Fully Connected)** | Every node is directly connected to every other node.                         | Small clusters, high-performance systems |
| **Tree (Hierarchical)** | Nodes are organized in a parent-child hierarchy, often for command/control flow. | Hadoop (NameNode -> DataNodes)           |
| **Bus**         | A single communication line (bus) connects all nodes; often logical in nature. | Some messaging systems, legacy networks   |
| **Hybrid**      | A combination of two or more topologies tailored for specific needs.       | Kubernetes clusters (control & data planes), large-scale microservices |
| **Sharded**     | Data is partitioned (sharded) across multiple nodes, often with coordinators or routers. | MongoDB, Elasticsearch                    |

---

## 🔍 Definitions and Examples

### 1. **Ring Topology**

- **Definition**: Each node manages a range of data tokens, forming a circular structure. Data and requests route through or to appropriate nodes based on consistent hashing.
- **Advantages**: High availability, no master, easy to scale.
- **Drawbacks**: Complex to rebalance, may need tuning of token ranges.

### 2. **Star Topology**

- **Definition**: Central node coordinates communication or data handling; all others are clients or replicas.
- **Advantages**: Simple, easy coordination.
- **Drawbacks**: Central node is a single point of failure.

### 3. **Mesh Topology**

- **Definition**: Every node communicates with every other node directly.
- **Advantages**: Very resilient, no bottleneck.
- **Drawbacks**: Not scalable; high overhead in large systems.

### 4. **Tree Topology**

- **Definition**: Nodes are arranged hierarchically. Often used for managing control flow or indexing large volumes of data.
- **Advantages**: Logical structure for metadata management.
- **Drawbacks**: Bottlenecks at upper nodes, not ideal for data sharing.

### 5. **Bus Topology**

- **Definition**: A single logical bus connects all nodes. Data is sent to all and filtered by destination.
- **Advantages**: Simple, low cost.
- **Drawbacks**: Collisions, scalability issues.

### 6. **Hybrid Topology**

- **Definition**: Combines aspects of different topologies to leverage their strengths.
- **Examples**: Kubernetes uses a control plane (star) with nodes communicating in mesh for certain workloads.

### 7. **Sharded Topology**

- **Definition**: Data is divided into shards, each managed by different nodes or replica sets.
- **Advantages**: Very scalable, allows horizontal partitioning.
- **Drawbacks**: Requires routing layer, uneven load if not balanced.

---

## 📊 Comparison Table

| Feature                | Ring        | Star       | Mesh        | Tree        | Bus         | Hybrid      | Sharded     |
|------------------------|-------------|------------|-------------|-------------|-------------|-------------|-------------|
| **Scalability**        | High        | Low-Med    | Low         | Medium      | Low         | High        | Very High   |
| **Fault Tolerance**    | High        | Low        | High        | Low-Med     | Low         | Depends     | High        |
| **Data Distribution**  | Even        | Centralized| Varies      | Hierarchical| Centralized | Flexible    | Partitioned |
| **Complexity**         | Medium      | Low        | High        | Medium      | Low         | High        | Medium-High |
| **Coordination**       | Peer-to-peer| Centralized| Peer-to-peer| Hierarchical| Shared Bus  | Mixed       | Router/Coordinator |
| **Use Case Fit**       | DBs, Logs   | Caches     | HPC, P2P    | File Systems| Messaging   | Cloud-native| NoSQL, Search |

// TODO  
