# Understanding the Use of Consensus Protocols in ScyllaDB

## 🔍 Overview

In a previous Scylla summit, it was announced that Scylla would use the **Raft consensus protocol** for **Lightweight Transactions (LWT)**. However, this decision has since changed. Scylla still works with Raft but will use it for other parts of the system, not for LWT. Here's why:

---

## 🧱 Scylla Architecture and the Problem with Raft

1. Scylla uses a **shared-nothing architecture**, meaning each node is independent and data is distributed across nodes.
2. Data is distributed using a **Consistent Hashing Ring**.
3. To balance the load better, **Virtual Nodes (VNodes)** are used. Each physical node contains multiple VNodes.
4. Each VNode has one primary replica and several secondary replicas, forming **Replication Groups**.

---

## 🔁 The Problem with Raft at Scale

Raft is a **leader-based protocol** where one node decides the transaction order. However:

- Scylla has **a massive number of Replication Groups** (e.g., thousands with many VNodes).
- Each VNode is further split into **CNodes**, sliced across shards within the node.
- Managing leadership state across all these groups is impractical with Raft.

❌ **Conclusion:** Raft doesn't scale well for LWT in Scylla due to excessive state overhead.

---

## ✅ The Solution: Using Paxos (Leader-less Protocol)

**Paxos** was designed for decision-making in distributed, unreliable networks. Though not originally for replication, it works well when adapted.

Advantages:

- No need to maintain leader state.
- More robust in dynamic environments.

Downside:

- Requires **more communication rounds**, making it costlier in performance.

---

## 🔄 Paxos Protocol Steps

1. **Prepare:** A replica asks others if they will accept a future value.
2. **Propose:** If most agree, it proposes the new value.
3. **Accept:** Majority confirms acceptance.
4. **Learn:** The coordinator informs others that the value is now committed.
5. **Check Conditions (LWT-specific):** An extra round if the transaction involves conditional logic.

---

## ⚠️ Challenges with Paxos

1. **High Cost:** It’s performance-heavy. Use only when **strong consistency** is needed.
2. **Starvation:** Some transactions may be repeatedly postponed. Solution: **Paxos Leases** (coming soon).
3. **Uncertainty:** Sometimes, the database commits a change but the client gets a timeout. Diagnostic tools are needed to track these issues.
4. **State Persistence:** Intermediate Paxos states must be stored (in system tables with TTL) to survive restarts.

---

## 🧹 Ongoing Improvements

- Merging Paxos rounds to reduce communication steps.
- Automatically cleaning up old Paxos states to reduce storage overhead.

---

## ✅ Summary

- Scylla moved away from Raft for LWT due to scalability issues.
- **Paxos is used instead**, despite its overhead, for its leader-less and resilient nature.
- The team is actively improving Paxos efficiency and error transparency for users.
