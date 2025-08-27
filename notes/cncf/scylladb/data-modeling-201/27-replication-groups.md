# Replication Groups and Raft Performance in ScyllaDB

## 🧮 How to Calculate the Number of Replication Groups

A **replication group** is a unique combination of VNodes that replicate the same data.

### 📊 Formula

Given:

- `N` = number of VNodes
- `RF` = replication factor

The number of replication groups is calculated using the binomial coefficient:

```txt
Replication Groups = C(N, RF) = N! / (RF! * (N - RF)!)
```

### 🔢 Example

For `N = 2560` and `RF = 3`:

```txt
C(2560, 3) = (2560 × 2559 × 2558) / (3 × 2 × 1)
           ≈ 2,791,458,186
```

✅ **Result: ~2.79 billion replication groups**

This large number makes it impractical to use leader-based protocols like Raft, which require maintaining per-group state.

---

## 🚦 How Replication Groups Affect Raft Performance

Raft is a **leader-based consensus protocol**. Each replication group requires:

### 1. 🔁 Leader Election Overhead

- Each group needs its own **leader**.
- Managing millions of leaders means:
  - High memory and CPU usage
  - Frequent elections on failures
  - Complex metadata tracking

### 2. 💾 Persistent State per Group

- Raft logs and metadata must be **persisted** for each group:
  - Term numbers
  - Commit indexes
  - Log entries
  - Snapshots
- Massive IO and storage overhead

### 3. 📡 Heartbeat and Log Replication

- Each leader sends **heartbeats** to followers
- Each change must be **replicated and acknowledged**
- High network traffic and latency

### 4. 🧹 Log Compaction and Garbage Collection

- Raft must **compact logs** to save space
- With many groups, background cleanup tasks increase
- Higher CPU load and potential GC pressure

---

## ✅ Why Scylla Uses Paxos Instead

**Paxos** (a leader-less protocol) solves this problem by:

- Electing a leader **per transaction**
- Avoiding long-term leader state
- Reducing coordination and heartbeat overhead

📌 **Trade-off:** More communication rounds per write, but much lower coordination cost.

---

## 🧠 Summary

| Area                | Why It’s a Problem in Raft |
|---------------------|----------------------------|
| Leader tracking     | Too many leaders to manage |
| Persistent metadata | Huge amount of Raft state  |
| Network usage       | Heartbeats for each group  |
| CPU load            | Background tasks like elections and compactions |

That's why Raft doesn't scale well in high-VNode architectures like ScyllaDB.
