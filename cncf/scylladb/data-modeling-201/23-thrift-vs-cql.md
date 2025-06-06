
# 🧵 What is Thrift and Why ScyllaDB Does Not Support It

## What is Thrift?

**Thrift** is a **remote procedure call (RPC)** protocol developed by **Apache** (originally by Facebook) for communication between services written in different languages.

### Thrift in Apache Cassandra

- Originally used as the **primary interface** for Cassandra.
- Allowed operations like reading/writing data and schema management.
- Used a binary protocol with predefined service methods.

---

## Why Was Thrift Used in Cassandra?

- Cassandra (pre-CQL) needed a generic RPC interface.
- Thrift enabled multi-language client support and low-level data operations.
- It provided **basic communication and schema access**.

Later, **CQL (Cassandra Query Language)** was introduced, providing

- SQL-like syntax
- Better data modeling
- Easier client interaction

---

## ❌ Why ScyllaDB Does Not Support Thrift

### 1. Legacy and Obsolescence

- Thrift is **deprecated** in Apache Cassandra (removed in version 4.0).
- Most users and applications have transitioned to **CQL**.

### 2. Too Low-Level

- Thrift exposed Cassandra internals directly.
- It **bypassed** important abstractions like the query planner and consistency layer.

### 3. Performance and Design Conflicts

- ScyllaDB is written in **C++** using the Seastar engine.
- Re-implementing Thrift would:
  - Require duplicating logic
  - Complicate the architecture
  - Reduce performance

### 4. Focus on CQL

- ScyllaDB supports the **Cassandra native protocol** (CQL binary protocol).
- This provides:
  - Full compatibility with modern clients
  - Better tooling (CQLSH, drivers, ORMs)
  - Clean and efficient query handling

---

## 🧠 Summary

| Feature                  | Thrift                                | CQL (Scylla's choice)                      |
|--------------------------|----------------------------------------|---------------------------------------------|
| Purpose                  | Low-level RPC protocol                 | High-level query language                   |
| Still used?              | ❌ Deprecated (removed in Cassandra 4.0) | ✅ Actively used                            |
| Supported by ScyllaDB?   | ❌ No                                   | ✅ Yes (full CQL + native protocol support) |
| Performance friendly?    | ❌ No                                   | ✅ Yes                                       |
| Tooling & ecosystem      | Weak                                   | Strong (CQLSH, drivers, ORMs, etc.)         |

---

## ✅ Conclusion

ScyllaDB **does not support Thrift** because it's an outdated, low-level protocol that conflicts with Scylla’s high-performance architecture. The modern and supported path is **CQL**, which offers better performance, tooling, and usability.
