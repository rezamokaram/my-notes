# 🧩 What is `sqlc.embed()`?

`sqlc.embed()` is a **macro** introduced in **sqlc v1.22+** that lets you **embed an entire struct (table type or another query’s return type)** directly into a query result.

It tells `sqlc` to **generate nested Go structs** instead of flattening all columns into one big struct.

---

## 💡 Why it exists

Before this macro, when you wrote a `JOIN` query like:

```sql
SELECT b.*, a.*
FROM books b
JOIN authors a ON a.id = b.author_id;
```

`sqlc` would generate **flat** fields like:

```go
type GetBooksRow struct {
    ID        int32
    Title     string
    AuthorID  int32
    ID_2      int32
    Name      string
}
```

😩 Ugly and hard to work with.

Now, using `sqlc.embed()`, you can get **nested structs**:

```go
type GetBooksRow struct {
    Book   Book
    Author Author
}
```

---

## 🧱 Example

### Schema

```sql
CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author_id INT NOT NULL REFERENCES authors(id)
);
```

### Query using `sqlc.embed()`

```sql
-- name: GetBooksWithAuthors :many
SELECT
    sqlc.embed(b),
    sqlc.embed(a) AS author
FROM books AS b
JOIN authors AS a ON a.id = b.author_id;
```

---

## 🧰 What `sqlc` generates (Go)

```go
type GetBooksWithAuthorsRow struct {
    Book   Book
    Author Author
}
```

Where `Book` and `Author` are structs already generated from your schema:

```go
type Book struct {
    ID       int32
    Title    string
    AuthorID int32
}

type Author struct {
    ID   int32
    Name string
}
```

---

## ⚙️ How it works

- `sqlc.embed(alias)` tells `sqlc`:
  > “For this table alias or subquery, generate a field that is the corresponding Go struct type.”

- It works with:
  - Tables (`sqlc.embed(b)`)
  - Subqueries with aliases (`sqlc.embed(sub)`)
  - Views

- You can rename the embedded struct using `AS`:

  ```sql
  SELECT sqlc.embed(a) AS writer
  ```

  → generates a field `Writer Author`

---

## ✅ Benefits

| Without `sqlc.embed()` | With `sqlc.embed()` |
|-------------------------|----------------------|
| Flat structs            | Nested structs       |
| Field name collisions   | Clean separation     |
| Manual aliasing needed  | Automatic reuse of types |

---

## ⚠️ Notes and Gotchas

- You **must use table aliases** in your query (`books AS b`).
- Works only with **PostgreSQL** and **MySQL** (as of latest versions).
- Only available in **`sqlc` version 1.22+**.
- You can combine multiple embeddings in one query.

---

## 🧠 Example with multiple embeddings

```sql
-- name: GetFullBookInfo :one
SELECT
    sqlc.embed(b),
    sqlc.embed(a) AS author,
    sqlc.embed(p) AS publisher
FROM books AS b
JOIN authors AS a ON b.author_id = a.id
JOIN publishers AS p ON b.publisher_id = p.id
WHERE b.id = $1;
```

Generates:

```go
type GetFullBookInfoRow struct {
    Book      Book
    Author    Author
    Publisher Publisher
}
```

---

## ✅ Summary

`sqlc.embed()` gives you:

- Clean, nested Go structs
- Automatic reuse of schema types
- Better readability for complex joins
- Less manual aliasing and field mapping
