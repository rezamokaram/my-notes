
# Lightweight Transactions (LWT) in ScyllaDB

## 📌 What is LWT in ScyllaDB?

- LWT allows **conditional updates** with **strong consistency**.
- Examples:

  ```sql
  INSERT INTO users (id, name) VALUES (1, 'Alice') IF NOT EXISTS;
  UPDATE users SET name = 'Bob' WHERE id = 1 IF name = 'Alice';
  ```

- Built on the **Paxos consensus protocol**.
- Guarantees **linearizability (serial consistency)**.
- Ensures only one successful conditional update at a time.

### 🔁 Paxos Phases in LWT

1. **Prepare Phase**: Read current state, propose intent.
2. **Propose Phase**: Submit proposal if quorum agrees.
3. **Commit Phase**: Commit the value once consensus is reached.

---

## 🔄 Does LWT read before write?

✅ **Yes**: LWT always reads data before attempting to write.

- It reads data during the **prepare phase** to:
  - Validate the condition (e.g., `IF NOT EXISTS`).
  - Ensure safe concurrent updates.
- This read is part of the Paxos protocol and **not a normal eventually consistent read**.

---

## ❓ Is LWT read possibly stale?

❌ **No** (in practice): LWT **does not return stale data** like normal reads might.

- LWT reads from a **quorum of replicas** to ensure freshness.
- Avoids stale reads by relying on Paxos state coordination.
- It provides **strong consistency**, unlike typical reads which may be eventually consistent.

---

## 🆚 ScyllaDB vs Cassandra: LWT Comparison

| Feature                     | **ScyllaDB**                             | **Apache Cassandra**                      |
|----------------------------|------------------------------------------|-------------------------------------------|
| Protocol                   | Paxos                                    | Paxos                                     |
| Consistency                | Linearizable                             | Linearizable                              |
| Performance                | Faster (shard-aware Paxos, local coord)  | Slower, higher overhead                   |
| Shard-aware                | ✅ Yes                                    | ❌ No                                     |
| Coordinator locality       | ✅ Coordinator runs on owning shard       | ❌ Not optimized                           |
| Disk I/O                   | Lower (in-memory optimizations)          | Higher (disk-based Paxos log)             |
| Contention handling        | Better scheduling and retries            | Prone to timeouts under load              |
| Use case                   | High-perf, scalable conditional writes   | Correct, but slower conditional writes    |

---

## ✅ ScyllaDB-Specific LWT Optimizations

- **Shard-aware Paxos**: LWT logic is tied to the CPU core that owns the data shard.
- **Coordinator-local Paxos**: Paxos coordination happens on the same shard as the data.
- **Fast-path Optimization**: Skips Paxos steps if no concurrent changes are detected.

---

## 🧠 Summary

- **LWT in ScyllaDB** ensures strong consistency, not eventual.
- Reads in LWT are quorum-based and safe.
- ScyllaDB LWT is **more efficient and scalable** than Cassandra due to architectural optimizations.
- Use LWT when you need **correctness under concurrency**, like unique inserts or conditional updates.
