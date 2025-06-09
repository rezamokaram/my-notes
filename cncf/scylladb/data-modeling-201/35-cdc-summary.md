
# ScyllaDB CDC: Summary and Insights

ScyllaDB's Change Data Capture (CDC) is built for simplicity, high performance, and seamless integration. It provides a robust and reliable way to observe changes in data through a normal CQL interface.

## ✅ Key Takeaways

- **Ease of Integration**: CDC data is stored in standard CQL tables. No new APIs or special clients required.
- **Cluster-Compatible**: CDC shares all properties of your regular data. If your cluster works, CDC works.
- **Low Overhead**: Writes to CDC logs are coalesced with main data writes to minimize impact.
- **Transient Data**: CDC data has TTL (default 24h) so unused logs won’t flood your system.
- **Replication & Deduplication**: CDC data is automatically deduplicated and replicated, just like your main data.
- **Optional Pre/Post Images**: You can choose to log preimage and postimage data.
- **Change Ordering**: All changes are ordered by client-perceived timestamps.

---

## 🔄 Feature Comparison

| Feature                         | ScyllaDB | Apache Cassandra | Others |
|---------------------------------|----------|------------------|--------|
| Separate Consumer Node Support  | ✅       | ❌ (must run on node) | ✅    |
| Pre-Image Support               | ✅ (optional) | ❌           | ✅    |
| Post-Image Support              | ✅ (optional) | ❌           | ✅    |
| Delta Support                   | ✅       | ✅                | ✅    |
| Deduplication of Changes        | ✅       | ❌                | ✅    |
| Standard Query Interface (CQL)  | ✅       | ✅                | ✅    |
| Data TTL on CDC Log             | ✅       | ❌                | ✅    |
| Minimal Performance Overhead    | ✅       | ❌                | ⚠️     |
| Fully Managed CDC Logs          | ✅       | ❌                | ⚠️     |

---

## 🧠 Can We Use ScyllaDB CDC as an Event Store?

**Yes, ScyllaDB CDC can be used as an event store** with some considerations:

### ✅ Advantages

- Events (mutations) are stored in an immutable and ordered log.
- Supports **preimage**, **delta**, and **postimage** which allows you to fully reconstruct the state change.
- CDC log is queryable using **CQL**, making integration easy.
- Changes are **replicated, deduplicated, and ordered**, ideal for event replay scenarios.

### ⚠️ Considerations

- CDC data is **transient** by default (TTL-based); for persistent event storage, you must **tune or offload** it appropriately.
- CDC reflects **application-level perception** of change, not absolute consensus.
- You must **consume and persist** the events elsewhere if you want long-term durability.

In summary, ScyllaDB CDC can serve as an excellent **source of truth for recent events**, and with proper consumption strategies, it can also act as a **full-fledged event store** for event-driven systems or microservice architectures.

To consume Cassandra’s CDC, a user has to start a special agent on every machine in the cluster. Such an agent has to parse commitlog-like segments stored in the special directory and after it’s finished with processing the data, it has to clean up the directory. In ScyllaDB, the CDC Log is a regular table accessible with CQL protocol and all standard tools over the wire.
