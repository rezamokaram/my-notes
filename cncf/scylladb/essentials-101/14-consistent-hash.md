# Consistent Hashing in ScyllaDB

## 🔄 What is Consistent Hashing?

Consistent hashing is a **technique for distributing data across multiple nodes** in a way that:

- Minimizes data movement when the cluster size changes.
- Maintains balanced data distribution.

### In ScyllaDB

- Each piece of data (row) has a **partition key**.
- That key is hashed using a **partitioner** (e.g., `Murmur3Partitioner`) to get a token.
- The **token ring** is a circular space of possible hash values (e.g., from `-2^63` to `2^63 - 1`).
- Nodes are assigned **token ranges**, and data is stored on the node responsible for the token range into which a key's hash falls.

---

## ✅ Why Use It?

### 1. **Scalability**

- You can add/remove nodes with minimal re-distribution of data.
- Helps the system scale **linearly**.

### 2. **High Availability**

- Data is **replicated** to multiple nodes (based on the replication factor).
- Ensures fault tolerance and reliability.

### 3. **Load Balancing**

- Especially with **virtual nodes (vnodes)**, data is evenly distributed across all nodes.

---

## ⚠️ What Happens When a Node Goes Down?

1. Other nodes **temporarily take over** the token ranges of the downed node.
2. Data requests are rerouted to remaining **replicas**.
3. Once the node returns:
   - It **catches up** using mechanisms like **hinted handoff** or **read repair**.
4. No single point of failure if **replication factor (RF)** is high enough (e.g., RF = 3).

---

## ➕ What Happens When a Node Is Added?

1. The new node is assigned one or more **token ranges** (or vnodes).
2. The cluster **rebalances**:
   - Data for these token ranges is **moved** from existing nodes to the new node.
3. Only affected token ranges are moved — not the whole dataset.
4. Minimal disruption thanks to **consistent hashing**.

---

## 📌 Visual Summary

Imagine a circle (the **hash ring**) with hashed token values:

```scss
|---N1---|---N2---|---N3---|---N1---|
```

After adding a new node (N4):

```scss
|---N1--|--N4--|--N2--|--N3--|--N1--|
```

Only part of the data (in N4's range) is moved.

---

## Bonus: Virtual Nodes (vnodes)

ScyllaDB uses **vnodes** to break a node's token range into many smaller ranges (e.g., 256 per node):

- Easier and faster **rebalancing**.
- Better **load distribution**.
- Simplifies node **replacement** and **scaling**.

---
