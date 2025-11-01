# 🧩 Understanding `:copyfrom` in sqlc (Go + PostgreSQL, MySQL, SQLite)

## ⚙️ What is `:copyfrom`?

The `:copyfrom` tag in **sqlc** is used for **bulk inserts**.  
It leverages PostgreSQL’s **`COPY FROM`** protocol internally for very fast data loading.

### Example

```sql
CREATE TABLE authors (
  id   SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  bio  TEXT NOT NULL
);

-- name: CreateAuthors :copyfrom
INSERT INTO authors (name, bio) VALUES ($1, $2);
```

Generates:

```go
type CreateAuthorsParams struct {
    Name string
    Bio  string
}

func (q *Queries) CreateAuthors(ctx context.Context, arg []CreateAuthorsParams) (int64, error)
```

---

## 🧱 How It Works Internally (PostgreSQL)

`sqlc` uses **pgx’s binary `CopyFrom` API**, similar to:

```go
func (q *Queries) CreateAuthors(ctx context.Context, rows []CreateAuthorsParams) (int64, error) {
    copyCount, err := q.db.CopyFrom(
        pgx.Identifier{"authors"},
        []string{"name", "bio"},
        pgx.CopyFromSlice(len(rows), func(i int) ([]any, error) {
            return []any{rows[i].Name, rows[i].Bio}, nil
        }),
    )
    return copyCount, err
}
```

That’s PostgreSQL’s **native COPY protocol**, the same as the SQL command:

```sql
COPY authors (name, bio) FROM STDIN;
```

---

## ⚡ Is It Transactional?

✅ **Yes.**  
The entire copy operation runs inside **a single transaction**.  
If any row fails (bad data, constraint violation, etc.), **everything is rolled back**.

So it’s fully **atomic** — all-or-nothing.

---

## 🚀 Is It the Best Way for Batch Writes?

✅ **Yes — for PostgreSQL.**  
It’s the fastest and most efficient method to insert large batches of rows.

| Method | Mechanism | Speed | Transactional | Notes |
|---------|------------|--------|----------------|--------|
| Loop of INSERTs | Many single inserts | 🐢 Slow | ✅ Yes | Many round-trips |
| Multi-row INSERT | One SQL statement | ⚙️ Medium | ✅ Yes | Fine for small batches |
| `:copyfrom` | COPY FROM protocol | 🚀🚀🚀 | ✅ Yes | Best for large inserts |

---

## 🧠 Pros and Cons

### ✅ Pros

- 🚀 **Extremely fast** (binary COPY protocol)
- 💾 **Memory efficient**
- 🧩 **Type-safe**
- 🔒 **Transactional**
- ✨ **Simple API (just pass a slice)**

### ❌ Cons

- 🚫 **One bad row aborts all**
- 🧱 **No RETURNING support**
- 🔍 **Insert-only**
- 🧰 **PostgreSQL only**
- ⚠️ **No per-row error handling**

---

## 🧩 What About MySQL and SQLite?

`sqlc`’s `:copyfrom` is **PostgreSQL-specific**, but sqlc **emulates** it for MySQL and SQLite.

### 🟨 MySQL

MySQL has **no `COPY FROM`** equivalent.  
sqlc will generate a loop of `INSERT` statements:

```go
for _, row := range rows {
    _, err := q.db.ExecContext(ctx, "INSERT INTO authors (name, bio) VALUES (?, ?)", row.Name, row.Bio)
}
```

✅ Works fine  
⚠️ Not very fast for large data sets  
✅ Transactional if wrapped in `BEGIN`/`COMMIT`

---

### 🟩 SQLite

Same as MySQL — no `COPY` protocol.  
sqlc loops through inserts:

```go
for _, row := range rows {
    _, err := db.ExecContext(ctx, "INSERT INTO authors (name, bio) VALUES (?, ?)", row.Name, row.Bio)
}
```

✅ Works fine for small data  
⚠️ Slower for large inserts  
✅ Transactional

---

## ⚖️ Performance Comparison

| Engine | Mechanism | Relative Speed | Transactional | Notes |
|---------|------------|----------------|----------------|--------|
| **PostgreSQL** | Binary COPY protocol | 🚀🚀🚀 | ✅ Yes | Best performance |
| **MySQL** | Loop of INSERTs | ⚙️ Medium | ✅ Yes | Slower for big data |
| **SQLite** | Loop of INSERTs | 🐢 Slow | ✅ Yes | Fine for small local data |

---

## 🧠 TL;DR

| Question | PostgreSQL | MySQL | SQLite |
|-----------|-------------|--------|----------|
| Supported? | ✅ Native | ⚠️ Emulated | ⚠️ Emulated |
| Transactional? | ✅ Yes | ✅ Yes | ✅ Yes |
| Fast? | 🚀 Yes | ⚙️ Okay | 🐢 No |
| COPY protocol used? | ✅ Yes | ❌ No | ❌ No |
| RETURNING supported? | ❌ No | ❌ No | ❌ No |
| Best use case | Bulk import, ETL | Small batches | Small embedded DB |

---

## ✅ Recommended Practices

- **Use `:copyfrom` only with PostgreSQL** for large-scale inserts.  
- For **MySQL/SQLite**, use **multi-row `INSERT`** statements instead.  
- Always **pre-validate** your data — a single bad row aborts the entire batch.  
- Don’t expect `RETURNING` values with `:copyfrom`.

---

## 🚀 Summary

| Engine | `:copyfrom` Behavior | Speed | Notes |
|---------|----------------------|--------|--------|
| PostgreSQL | True COPY FROM | 🚀 Fastest | Fully transactional |
| MySQL | Emulated (loop insert) | ⚙️ Medium | Works but not optimized |
| SQLite | Emulated (loop insert) | 🐢 Slow | For small datasets only |

**✅ Bottom line:**  
Use `:copyfrom` for PostgreSQL bulk inserts; for MySQL/SQLite, it’s just a safe loop — not a performance boost.
