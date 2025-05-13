# 📘 ScyllaDB: Replication Factor (RF) & Token Range Notes

---

## Replication Factor (RF)

### What is RF?

The **Replication Factor** in ScyllaDB defines **how many copies of each piece of data** are stored across different nodes in the cluster.

---

### Key Points

- **RF = N** means data is stored on **N different nodes**.
- Defined **per keyspace**.
- Higher RF = more fault tolerance, but also more storage and write overhead.

---

### RF vs. Node Count

| Replication Factor | Minimum Required Nodes |
|---------------------|-------------------------|
| 1                   | 1                       |
| 2                   | 2                       |
| 3                   | 3                       |
| 5                   | 5                       |

> ❗ Using RF higher than the number of available nodes will result in **under-replication** and **potential consistency issues**.

---

### Example: RF = 3 on 3 Nodes

- Every piece of data is **replicated to all 3 nodes**.
- Data **ownership** is still distributed evenly using **token ranges**.
- This does **not break the ring model** — it's how the system is designed.

---

### Ring and Replica Placement

In a 3-node cluster with RF = 3:

- Each node **owns a primary token range**.
- Replicas are placed **clockwise** in the ring.
- Example:

| Token Range | Primary Owner | Replica 2 | Replica 3 |
|-------------|----------------|------------|------------|
| 0–100       | Node A         | Node B     | Node C     |
| 101–200     | Node B         | Node C     | Node A     |
| 201–300     | Node C         | Node A     | Node B     |

> Even though all nodes store all data (because RF = node count), **data partitioning and ring topology remain intact**.

---

## Token Ranges

### What is a Token Range?

- ScyllaDB uses a **consistent hashing mechanism** to assign a **token** (hash value) to each row based on the **partition key**.
- The **token range** defines which node is responsible for storing the **primary replica** of that row.

---

### Token Range Ownership

- The full token space is a **ring of hash values**, split among the nodes.
- Each node is assigned a **range of tokens**.
- Replication is done by assigning replicas to the **next nodes** in the ring (clockwise).

---

### Example of Token Ownership

Assume token space from 0 to 300:

| Node   | Token Range |
|--------|-------------|
| Node A | 0–100       |
| Node B | 101–200     |
| Node C | 201–300     |

- A row with a token hash of 150 is **owned by Node B**.
- With RF = 3:
  - Primary copy: Node B
  - Replica 2: Node C
  - Replica 3: Node A

## Replication Factor (RF) and Availability in ScyllaDB

- RF determines how many replicas of data are stored across nodes. For example:

  - RF = 2: Two copies of data.

  - RF = 3: Three copies of data.

- Higher RF increases availability because more replicas exist. If one node fails, data can still be served from the other replicas.

- RF = 3 means data is available as long as at least 2 replicas are up.

---

## Summary

- **RF determines how many nodes store each row.**
- **Token ranges determine which node is primarily responsible for a given row.**
- Even with RF equal to the number of nodes, the ring and partitioning logic remains valid and efficient.
