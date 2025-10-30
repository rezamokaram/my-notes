# 🧩 Difference Between `ANY` and `IN` in sqlc

## Overview

In `sqlc`, both `ANY` and `IN` are SQL operators used to compare a value against multiple possible values.  
However, they behave very differently when used with Go slices and query parameters.

---

## ⚖️ Comparison

| Feature | `IN` | `ANY` |
|----------|------|-------|
| SQL Meaning | Checks if a value is **equal to one** in a list or subquery | Compares a value to **any element** of an array or subquery result |
| Requires comparison operator | ❌ No (always equality) | ✅ Yes (`=`, `<`, `>`, etc.) |
| Works with SQL arrays | ❌ No | ✅ Yes |
| Works with Go slices in `sqlc` | ❌ No | ✅ Yes |
| Safe static query in `sqlc` | ❌ Needs dynamic SQL building | ✅ Fully supported |
| Typical use in `sqlc` | Rare / discouraged | Common / recommended |
| Example SQL | `WHERE id IN (1, 2, 3)` | `WHERE id = ANY($1::int[])` |
| Go parameter type | Not supported (`IN` needs literal list) | Supported (e.g. `[]int32`, `[]string`) |

---

## ✅ Example in sqlc

### Good (`ANY`)

```sql
-- name: GetUsersByIDs :many
SELECT * FROM users
WHERE id = ANY($1::int[]);
```

**Generated Go code:**

```go
func (q *Queries) GetUsersByIDs(ctx context.Context, ids []int32) ([]User, error)
```

### Bad (`IN`)

```sql
-- name: GetUsersByIDs :many
SELECT * FROM users
WHERE id IN ($1);
```

❌ Fails because `$1` is treated as a single value, not a list.  
You’d have to dynamically build SQL, which `sqlc` does not allow.

---

## 🧠 Summary

- Use `ANY` when you want to pass **Go slices or arrays** as parameters.
- Avoid `IN` with sqlc, because it doesn’t accept array parameters.
- `ANY` + PostgreSQL arrays = clean, type-safe, static queries.
