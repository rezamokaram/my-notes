# 🧩 `sqlc.arg()` and `sqlc.slice()` in sqlc

These two helpers make it easier to use **structured parameters** and **Go slices** inside your SQL files when working with `sqlc`.

---

## ⚙️ Overview

| Function | Purpose |
|-----------|----------|
| **`sqlc.arg()`** | References a named parameter or struct field in SQL queries. |
| **`sqlc.slice()`** | Expands Go slices into SQL-friendly arrays or lists. |

---

## 🧠 `sqlc.arg()`

`sqlc.arg()` lets you use **named parameters** instead of positional placeholders like `$1`, `$2`, or `?`.

### ✅ Example

```sql
-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES (sqlc.arg(name), sqlc.arg(email))
RETURNING id, name, email;
```

**Generated Go code:**

```go
type CreateUserParams struct {
    Name  string
    Email string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
```

**Call in Go:**

```go
user, err := q.CreateUser(ctx, db.CreateUserParams{
    Name:  "Alice",
    Email: "alice@example.com",
})
```

🧩 **Key points**

- Works on **PostgreSQL**, **MySQL**, and **SQLite**.
- Makes queries self-documenting and type-safe.
- `sqlc.arg(field)` → references `arg.Field` in Go.

---

## 🧠 `sqlc.slice()`

`sqlc.slice()` is used to pass **Go slices** (like `[]int32`, `[]string`) into SQL queries.  
It helps safely use lists for filtering (like `IN` or `ANY`).

### ✅ Example (PostgreSQL)

```sql
-- name: GetUsersByIDs :many
SELECT * FROM users
WHERE id = ANY(sqlc.slice(ids));
```

**Generated Go code:**

```go
func (q *Queries) GetUsersByIDs(ctx context.Context, ids []int32) ([]User, error)
```

**Call in Go:**

```go
users, err := q.GetUsersByIDs(ctx, []int32{1, 2, 3})
```

🧩 **Key points**

- Best used with PostgreSQL’s `ANY()` operator.
- Lets you safely pass Go slices to SQL without string interpolation.
- Works with `IN()` on MySQL and SQLite (limited).

---

## ⚖️ Database Compatibility

| Database | `sqlc.arg()` | `sqlc.slice()` | Notes |
|-----------|---------------|----------------|--------|
| **PostgreSQL** | ✅ Full support | ✅ Full support | Use `WHERE id = ANY(sqlc.slice(ids))` |
| **MySQL** | ✅ Supported | ⚠️ Partial (no arrays) | Use `WHERE id IN (sqlc.slice(ids))` |
| **SQLite** | ✅ Supported | ⚠️ Limited | Expands placeholders manually |

---

## 🧱 Summary

| Feature | `sqlc.arg()` | `sqlc.slice()` |
|----------|--------------|----------------|
| Purpose | Named struct parameters | Pass Go slices / arrays |
| Works with Postgres | ✅ | ✅ |
| Works with MySQL | ✅ | ⚠️ Partial |
| Works with SQLite | ✅ | ⚠️ Limited |
| Safe and type-checked | ✅ | ✅ (Postgres only) |
| Example use | `VALUES (sqlc.arg(name))` | `WHERE id = ANY(sqlc.slice(ids))` |

---

## ✅ TL;DR

- **`sqlc.arg()`** → use for named parameters and cleaner structs (works on all DBs).  
- **`sqlc.slice()`** → use for passing slices/arrays (best with PostgreSQL).  
- For PostgreSQL:

  ```sql
  WHERE id = ANY(sqlc.slice(ids));
  ```

- For MySQL or SQLite:

  ```sql
  WHERE id IN (sqlc.slice(ids));
  ```

  *(may have limited support depending on your sqlc version).*
