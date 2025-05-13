# Partition key use a consistent hash function

## Partition Key and Consistent Hashing in Distributed Systems

## What is a **Partition Key**?

In a **distributed database** or **NoSQL system**, the **partition key** is a **specific field (or fields)** in the data that is used to determine where the data should be stored in the cluster.

For example, in a **customer database**, the partition key might be **customer_id**, and every time data for a specific customer is inserted, that data is stored in a specific location determined by the **partition key**.

The partition key ensures that all data for the same entity (e.g., all orders from a customer) are located in the same **partition** (and therefore on the same node or set of nodes in the cluster), making it easier to retrieve all of that data efficiently.

---

## What is **Consistent Hashing**?

**Consistent hashing** is a method for distributing data across a distributed system in a way that minimizes the impact of adding or removing nodes to the system.

In simple terms:

- **Consistent hashing** allows you to map data onto a **ring** of nodes.
- Each node (or server) in the system is assigned a **position** on the ring.
- The position of data on the ring is determined by applying a **hash function** to the **partition key**.

### Why is it Called **"Consistent"**?

The word **"consistent"** comes from the fact that when you add or remove nodes, only a small amount of data needs to be redistributed. In traditional hashing, adding a node might require rehashing all the data, but with consistent hashing, the impact is minimized.

---

## How **Consistent Hashing** Works

Here’s how **consistent hashing** works in the context of a distributed database (like ScyllaDB or Cassandra):

1. **Hashing the Partition Key**:
   - The partition key (e.g., `customer_id`) is **hashed** using a **hash function** (e.g., MD5, SHA-256).
   - This hash produces a numeric value.

2. **Assigning Data to Nodes**:
   - This numeric hash value is then mapped to a **position on a ring**.
   - The ring is a circular structure, where the positions of nodes are placed at different points along the ring.
   - The data (based on the partition key's hash value) is then assigned to the **first node** on the ring that is equal to or greater than the hash value. This is called the **virtual position**.

3. **Handling Node Failures or Additions**:
   - When a node joins or leaves the cluster, only a small portion of the data (the data that falls within the range of the affected node) needs to be rehashed and moved. This is the key advantage of consistent hashing: minimal data movement.

4. **Replication**:
   - Typically, data is **replicated** across multiple nodes for **fault tolerance**. This means the data for a given partition key will not only exist on the node corresponding to its hash value but will also be replicated on additional nodes, based on the replication factor.

---

## Example of **Consistent Hashing**

Let’s use a simple example to illustrate how consistent hashing works in a cluster.

### 1. **The Nodes**

Imagine we have a cluster with **4 nodes** (A, B, C, and D). They are positioned on the consistent hash ring like this:

A ---- B ---- C ---- D ---- A (loop back)

### 2. **Hashing the Partition Key**

Now, let’s say we have a partition key called **`customer_id`**. Let’s hash the `customer_id` (for example, `12345`) using a hash function like SHA-256.

The result of hashing `12345` gives us a hash value (say **`X`**) between 0 and 100. Let’s assume:

- `X = 35`

### 3. **Placing the Data**

Now, on the ring, we place the hash value **`X = 35`**. The system checks the closest node in the clockwise direction on the ring, and in this case, it lands on **Node B**.

Thus, **data for `customer_id = 12345`** is placed on **Node B**.

### 4. **Replication**

If we have a replication factor of **3**, then this data will also be replicated to the next two nodes in the ring:

- **Node C** (next node after B)
- **Node D** (next node after C)

---

## Key Advantages of **Consistent Hashing**

1. **Scalability**:
   - It’s easy to add or remove nodes without rebalancing the entire dataset.

2. **Efficient Data Distribution**:
   - Data is evenly distributed across nodes based on the hash of the partition key, so it avoids data hotspots (unless the partition key is not chosen well).

3. **Fault Tolerance**:
   - Data is replicated across multiple nodes. If a node goes down, its data can be retrieved from replicas.

4. **Minimized Impact of Node Changes**:
   - When nodes are added or removed, only the data that needs to be redistributed (i.e., data that hashes to the affected area of the ring) is moved. This is the primary reason consistent hashing is popular in distributed databases.

---

## Consistent Hashing vs. Other Hashing Methods

- **Traditional Hashing**: In standard hashing methods, adding or removing a node often requires **recomputing** and **reassigning** data for the entire system, which can be very inefficient for large-scale systems.
  
- **Consistent Hashing**: As described, consistent hashing is **more efficient** because only the data that needs to be redistributed (i.e., data that hashes to the affected area of the ring) is moved. This is the primary reason consistent hashing is popular in distributed databases.

---

## Common Hash Functions Used

- **MD5**: A 128-bit hash value, though not cryptographically secure, it’s fast and still widely used in distributed systems.
- **SHA-256**: A 256-bit hash function from the SHA family that’s widely used for cryptographic purposes.
- **MurmurHash**: A non-cryptographic hash function often used in distributed systems because of its speed and good distribution properties.

---

## Summary

- The **partition key** determines where data should be stored in a distributed system.
- **Consistent hashing** is used to map the **partition key** to a node in the system using a hash function.
- This method minimizes data movement when nodes are added or removed from the cluster, making it highly efficient and scalable.
  
By using consistent hashing, systems can scale horizontally and provide fault tolerance with minimal disruption.

## scylla db cluster and data replication

// TODO


