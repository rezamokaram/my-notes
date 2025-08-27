
# ScyllaDB Primary Key, Clustering, and Querying Rules

## 1. ❓ Why this query fails?

```sql
SELECT * FROM heartrate_v7
WHERE pet_chip_id = ... AND time = ... AND heart_rate = 100;
```

### ❌ Error

```text
InvalidRequest: PRIMARY KEY column "heart_rate" cannot be restricted as preceding column "pet_name" is not restricted
```

### ✅ Explanation

In your table:

```sql
PRIMARY KEY ((pet_chip_id, time), pet_name, heart_rate)
```

- `pet_chip_id` and `time` → Partition Key
- `pet_name`, `heart_rate` → Clustering Keys (in that order)

You cannot filter `heart_rate` without first filtering `pet_name` because clustering keys must be queried in order.

---

## 2. ❓ Why this query works?

```sql
SELECT * FROM heartrate_v7
WHERE pet_chip_id = ... AND time = ...;
```

### Explanation

You're filtering the **entire partition key** and not touching any clustering keys. This is perfectly valid. It retrieves **all rows in that partition** (i.e., all `pet_name` and `heart_rate` values for that `pet_chip_id` + `time`).

---

## 3. ❓ What are my options?

You asked if you must either:

1. Not use clustering keys  
2. Use all clustering keys together when querying  

### ✅ Answer

Yes, mostly correct. Here's the breakdown:

| Option | Description |
|--------|-------------|
| 1. No clustering keys | ✅ Simplest structure. One row per partition key. |
| 2. Use all clustering keys in order | ✅ Required if clustering keys are defined. You must query them **in the defined order**, without skipping. |
| 3. Alternatives | 🔁 You can use materialized views, secondary indexes (carefully), or design separate tables for different access patterns. |
