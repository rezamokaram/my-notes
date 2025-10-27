# 🧩 What Is Type Safety

**Type safety** means a programming language or tool prevents you from using data in ways that don’t match its type.  

In other words:  
> **The compiler (or runtime) ensures that integers, strings, and other data types are used correctly and consistently.**

Type safety helps:

- Catch **errors early (at compile time)** instead of at runtime.  
- Keep **database types and code types** consistent (e.g., an `INT` column always maps to an `int` in Go).  
- Avoid runtime bugs like “cannot convert string to int”.

---

## 🧠 Type Safety in the Context of ORMs / Database Tools

When we talk about **type safety in ORMs or SQL tools**, we mean:

> The ORM or SQL generator ensures — at compile time — that the types of columns, query parameters, and return values all match between your **database schema** and your **Go code**.

A **type-safe ORM or query builder** gives compile-time errors if:

- You pass the wrong type to a query (e.g., a string where an int is expected).  
- You reference a non-existent column or table.  
- The database schema changes and your Go code no longer matches.

In contrast, a **non–type-safe ORM** lets these mistakes compile fine, but you only find out when the program runs.

---

## ⚖️ Comparison: Go Database Tools and ORMs

| Tool | Approach | Type Safety | Schema Source | Compile-time Guarantees | Description |
|------|-----------|--------------|----------------|--------------------------|--------------|
| **sqlc** | SQL-first | ✅ **Strong** | Database (SQL schema) | Yes | You write SQL, sqlc parses it and generates strongly typed Go code. Any type mismatch causes compile errors. |
| **ent** | Code-first ORM | ✅ **Strong** | Go code (schema definitions) | Yes | You define your schema in Go, ent generates Go structs and type-safe query builders. Queries and fields are checked by the compiler. |
| **gorm** | Reflection-based ORM | ⚠️ **Weak** | Go structs at runtime | No | Uses reflection and strings to build SQL; errors (like wrong column names or types) only show up when the query runs. |
| **sqlx** | SQL helper | ⚠️ **Weak** | SQL (manual mapping) | No | Easier than database/sql but still uses manual type mapping; doesn’t validate types at compile time. |

---

## 🔍 Summary Takeaways

- **Type safety = early error detection + consistency between DB and code.**
- **sqlc** → SQL-first, compile-time type checks, safest.  
- **ent** → Go-schema-first, compile-time type checks, modern and safe.  
- **gorm** → Runtime reflection, convenient but less safe.  
- **sqlx** → Lightweight, flexible, but you manually ensure correctness.

---

## 🏁 In Short

In Go’s ecosystem:

> - **sqlc** and **ent** are *type-safe*.
> - **gorm** and **sqlx** are *not type-safe* (they rely on runtime checks).

Type safety means the compiler helps you catch mistakes before your code runs — which makes your database layer more reliable and predictable.
