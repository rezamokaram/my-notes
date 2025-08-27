# Keyspace

## What is a Keyspace?

A **keyspace** in ScyllaDB is the top-level namespace that defines how data is stored and replicated across the cluster. It acts like a database in relational systems.

---

## Keyspace Key Components

### 1. Name

- Unique identifier for the keyspace.

### 2. Replication Strategy

Defines how and where data is copied across the cluster.

- **SimpleStrategy**: Single data center or development environments.
- **NetworkTopologyStrategy**: Recommended for production and multi-datacenter deployments.

### 3. Replication Factor (RF)

- Number of copies of each piece of data.
- Example: `RF = 3` means data is stored on 3 different nodes.

### 4. Durable Writes

- **true** (default): Writes are flushed to disk before acknowledgement.
- **false**: Faster writes but less durable.

---

## 🧾 Example

```cql
CREATE KEYSPACE user_data 
WITH replication = {
  'class': 'NetworkTopologyStrategy',
  'datacenter1': 3
} 
AND durable_writes = true;
```

## ✅ Best Practices

- Use NetworkTopologyStrategy for production systems.
- Set appropriate replication factors based on your availability and fault tolerance needs.

- Use multiple keyspaces to separate application concerns if needed.

## Notes

- Keyspaces contain tables, but no actual data themselves.

- Changing replication settings on a keyspace may require a repair or rebuild to take effect properly.
