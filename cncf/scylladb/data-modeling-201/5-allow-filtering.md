# ⚠️ Understanding `ALLOW FILTERING` in ScyllaDB

## 🔍 What is `ALLOW FILTERING`?

In **ScyllaDB (and Cassandra)**, if a query **doesn't follow the primary key rules**, it fails by default.  
You can override this behavior with:

```sql
SELECT * FROM table_name WHERE some_column = 'value' ALLOW FILTERING;
```

---

## ⚠️ Why Is It Risky?

`ALLOW FILTERING` tells Scylla:
> "I know this query might scan lots of data — do it anyway."

- 🚫 Can be **very slow**
- 🚫 Uses **lots of memory**
- 🚫 Unpredictable performance

---

## ✅ When Is It Used?

Used when querying:

- Non-primary key columns
- Clustering keys **out of order**
- Without all parts of the primary key

**Example:**

```sql
-- Invalid without ALLOW FILTERING
SELECT * FROM users WHERE age = 30;

-- Works with ALLOW FILTERING
SELECT * FROM users WHERE age = 30 ALLOW FILTERING;
```

---

## 🧠 Should You Use It?

| Scenario                        | Use `ALLOW FILTERING`? | Recommendation          |
|---------------------------------|-------------------------|--------------------------|
| Small table (few rows)          | ✅ Okay                 | Safe for quick lookups   |
| Big table (millions of rows)    | ❌ Avoid                | Use index/view/schema redesign |
| Development / testing only      | ✅ Temporary            | Not for production use   |
| Real-time or high-frequency use | ❌ Never                | Must optimize schema     |

---

## ✅ Better Alternatives

Instead of `ALLOW FILTERING`, consider:

- 🔁 **Redesigning primary key** to match query pattern
- 🧩 **Secondary indexes** (with care)
- 👁️ **Materialized Views**
- 🗃️ **Separate tables** for different query paths

---

## 🧪 Example

**Risky (scans entire dataset):**

```sql
SELECT * FROM heartrate_v7 WHERE heart_rate = 90 ALLOW FILTERING;
```

**Better (custom schema):**

```sql
CREATE TABLE heart_rate_index (
  heart_rate int,
  pet_chip_id uuid,
  time timestamp,
  pet_name text,
  PRIMARY KEY (heart_rate, pet_chip_id, time)
);
```

---

## ✅ Summary

Use `ALLOW FILTERING` **only as a last resort** and **never in high-load production scenarios**. Always try to design your schema to match your query needs.
