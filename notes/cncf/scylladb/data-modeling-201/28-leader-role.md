# Leader-Based vs. Leader-Less Consensus Protocols in ScyllaDB

This summary compares Leader-Based and Leader-Less consensus protocols based on the architecture and decisions described in the Scylla Summit video.

---

## 🔄 Leader-Based vs. Leader-Less Consensus Protocols

| Feature                        | **Leader-Based (e.g., Raft)**                            | **Leader-Less (e.g., Paxos in Scylla)**                  |
|-------------------------------|----------------------------------------------------------|----------------------------------------------------------|
| 🧠 **Leadership Model**        | One elected leader per replication group                | Temporary leader per transaction                         |
| 🔁 **Coordination Overhead**   | Lower per transaction (fewer rounds)                    | Higher per transaction (more negotiation)               |
| 🗳️ **Election Cost**          | Costly when scaled (millions of replication groups)     | No persistent leader = less metadata to track           |
| 📦 **State Management**       | Needs persistent leader state (term, log, index, etc.)  | Stateless coordination per operation                    |
| 🌐 **Network Load**           | Continuous heartbeats + log replication                 | On-demand message rounds (prepare, propose, learn)      |
| 📈 **Scalability**            | Poor with many VNodes (e.g., 2500+ leads to billions)   | More scalable with high VNode counts                    |
| 🔄 **Recovery Handling**      | Complex (requires log replay or snapshot restore)       | Simpler (each op is independently negotiated)           |
| 🧩 **Usage in Scylla**        | Not used for LWT (too much state overhead)              | Preferred for LWT despite more negotiation rounds       |
| ⚠️ **Common Pitfalls**        | Failover tracking, leader liveness checks               | Possible mutation uncertainty on timeout                |

---

## 📝 Key Takeaways from the Scylla Video

- Scylla initially planned to use **Raft** for Lightweight Transactions (LWT), but abandoned the idea due to:
  - High overhead in tracking leaders across thousands of replication groups (due to many VNodes and shards).
  - Complexity in maintaining and recovering consistent leader state.

- Instead, Scylla chose **Paxos**, a **leader-less protocol**, because:
  - It is more scalable for Scylla's architecture.
  - It avoids having to track long-lived leader state.
  - It is more resilient to node failures and restarts.

---
