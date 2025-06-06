
# 📄 Uniform Result Sets in ScyllaDB: What "Always Provides a Result Set" Means

## ✅ What Does It Mean?

In **ScyllaDB**, every query — whether it's a **read**, **write**, or **conditional operation** — returns a **result set**, even if that set is **empty**.

This is consistent with the **Cassandra Query Language (CQL) protocol**.

---

## 🔍 Why This Matters

- Clients interacting with ScyllaDB **always receive a result**, regardless of query type.
- Simplifies **driver development and client logic** — no need to handle different response types.
- Ensures **protocol consistency** across operations.

---

## 📘 Examples

### 🔹 Regular Write (Insert)

```sql
INSERT INTO users (id, name) VALUES (1, 'Alice');
```

**Returns**: ✅ An **empty result set** (no data, but still a valid response)

---

### 🔹 Conditional Write (LWT)

```sql
INSERT INTO users (id, name) VALUES (1, 'Alice') IF NOT EXISTS;
```

**Returns**:

```text
[applied] | id | name
----------+----+------
 true     |    |
```

—or—

```text
[applied] | id | name
----------+----+------
 false    | 1  | Bob
```

---

## 🚫 Comparison with SQL Databases

In traditional SQL databases like **PostgreSQL**:

- A plain `INSERT` or `UPDATE` returns **nothing** unless you use `RETURNING`.
- Response type may vary between queries.

In contrast, **ScyllaDB (via CQL)** always returns a **well-defined result set**, even for non-`SELECT` queries.

---

## 🧠 Summary

| Aspect                        | Behavior in ScyllaDB                         |
|------------------------------|----------------------------------------------|
| SELECT queries               | Returns rows                                 |
| INSERT/UPDATE/DELETE         | Returns an **empty result set**              |
| LWT / Conditional statements | Returns result set with `[applied]` column   |
| Client benefit               | Predictable, uniform handling of responses   |

---

## 💡 Final Thought

The design choice that **"Scylla always provides a result set"** helps ensure **robust, consistent, and developer-friendly** communication between client applications and the database — one of the subtle but powerful features inherited from the CQL protocol.
