
# How ScyllaDB CDC Works Under the Hood

ScyllaDB's **Change Data Capture (CDC)** is a feature that allows users to track changes to data in selected tables. It's designed to be both efficient and transparent, enabling integration with real-time pipelines, audit systems, or data replication mechanisms.

## ✅ CDC as a Log Table

When CDC is enabled on a table in ScyllaDB:

- **A separate CDC log table is created** for that base table.
- This log table contains **one row per mutation** (insert, update, delete).
- The CDC table resides in the same cluster, shares the same data distribution (partitioning), and can be queried using **standard CQL**.

This means CDC data behaves just like any other table in ScyllaDB: it's distributed, fault-tolerant, and can be filtered, paged, and queried.

---

## 🔧 Example: What Happens Under the Hood

### Step 1: Create a Table with CDC Enabled

```sql
CREATE TABLE users (
  user_id UUID,
  timestamp TIMESTAMP,
  name TEXT,
  email TEXT,
  PRIMARY KEY (user_id, timestamp)
) WITH cdc = {'enabled': true, 'preimage': 'true'};
```

- `cdc = {'enabled': true}` turns on CDC.
- `'preimage': 'true'` requests that **before-change values** are also logged.

### Step 2: Insert Data

```sql
INSERT INTO users (user_id, timestamp, name, email) VALUES (...);
```

- Since the row is new, there’s **no preimage**.
- The CDC table logs this as an **insert** operation with current values (post-image only).

### Step 3: Update Data

```sql
UPDATE users SET name = 'Alice' WHERE user_id = ... AND timestamp = ...;
```

- The CDC table will log:
  - The **preimage**: the original value of `name` (if it existed).
  - The **delta**: what changed (e.g., name from "Bob" to "Alice").
  - TTL if present.

This ensures that only the changed columns are logged, minimizing storage overhead.

### Step 4: Delete Column or Row

- **Column Delete**:  
  `UPDATE users SET email = null ...;`
  - Logs the preimage of `email`.
  - Logs a "delete delta" indicating `email` is now null.

- **Row Delete**:  
  `DELETE FROM users WHERE user_id = ... AND timestamp = ...;`
  - Logs preimages of all columns in the row.
  - Logs a "row delete" operation.
  - **No column deltas are included**, as the entire row is gone.

- **Partition Delete**:  
  `DELETE FROM users WHERE user_id = ...;`
  - Logs a single record that the partition was deleted.
  - No preimages for individual rows.

---

## 📦 Log Structure

Each row in the CDC log contains:

- A **timestamp** and **operation type** (insert, update, delete).
- The **delta**: what changed.
- Optional **preimage** and/or **postimage**.
- A **batch sequence number** to correlate changes from the same write operation.

---

## ⚙️ Performance and Design Considerations

- **CDC data is transient**: by default, it’s retained for 24 hours (configurable).
- **Minimal overhead**: only affected columns are logged.
- **Consistency model**: CDC reflects the view of the data **as seen by the coordinator node**, which means it inherits Scylla’s **eventual consistency** model.
- **Failure tolerance**: Some logs may be lost if nodes crash before data is flushed, but CDC aims to minimize this risk.

---

## 🧠 Summary

- ScyllaDB CDC works by creating **a separate log table per base table**.
- It captures data mutations (inserts, updates, deletes) along with optional pre/post images.
- It logs **only the changed columns** to reduce impact on performance and storage.
- Consumers can read CDC data just like any other CQL table.
- ScyllaDB provides fine-grained control over what to log (e.g., deltas, preimages).
- CDC supports advanced integration via connectors (e.g., Kafka) or custom logic.

This design enables real-time, efficient, and scalable data change tracking natively within ScyllaDB, without requiring external CDC systems or triggers.
