
# ScyllaDB Key Concepts – Simple Explanation

## 🔑 1. Composite (Compound) Key

- A **composite key** = multiple columns combined to uniquely identify a row.
- In ScyllaDB, it includes:
  - **Partition key**
  - **Clustering key(s)**

**🧠 Think of it as:**  
`(Partition Key) + (Clustering Keys) = Composite (Compound) Key`

---

## 🧩 2. Composite Primary Key

Defined like this in CQL:

```sql
PRIMARY KEY ((partition_key), clustering_key1, clustering_key2)
```

- The **entire structure** is the **composite primary key**.
- It controls:
  - Uniqueness of each row
  - How data is **partitioned** (partition key)
  - How rows are **ordered** within a partition (clustering keys)

---

## 📦 3. Partition Key

- Used to **distribute data** across nodes.
- All rows with the **same partition key** go to the **same node**.
- Defined inside **double parentheses**:

```sql
PRIMARY KEY ((user_id), timestamp)
```

**🧠 Think of it as the "address" of the data in the cluster.**

---

## 🧮 4. Clustering Key

- Used to **sort data** **within** a partition.
- Helps you **query ordered data** efficiently.

```sql
PRIMARY KEY ((user_id), timestamp)
```

- `timestamp` is the **clustering key**.

**🧠 Think of it as the "order" inside a folder.**

---

## 💡 Example

```sql
CREATE TABLE user_logs (
  user_id UUID,
  timestamp TIMESTAMP,
  action TEXT,
  PRIMARY KEY ((user_id), timestamp)
);
```

- **Partition Key**: `user_id` → which node stores the data
- **Clustering Key**: `timestamp` → sorts logs by time within that user
- **Composite Primary Key**: (`user_id`, `timestamp`)
