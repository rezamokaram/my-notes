# UPSERT

UPSERT is a combination of UPDATE and INSERT:

- If the row doesn’t exist, it will be inserted.

- If the row already exists, it will be updated with the new data.

In ScyllaDB (and Cassandra), every INSERT is actually an UPSERT by default.

## Core Similarities

- Both will insert a new row if it doesn't exist.

- Both will update existing rows if they do exist.

- Neither requires a read-before-write, so they are very fast.

- Both overwrite specified columns but do not delete unspecified columns (missing columns remain unchanged).

## Differences Between INSERT and UPDATE in ScyllaDB

| Aspect               | `INSERT`                                         | `UPDATE`                                    |
|----------------------|-------------------------------------------------|---------------------------------------------|
| **Intention / Semantics** | Usually used to add new rows (or upsert)          | Usually used to modify existing rows (or upsert) |
| **Conditional Insert** | Supports **`IF NOT EXISTS`** for conditional insert | Supports **`IF EXISTS`** to update only if row exists |
| **TTL and Timestamp Behavior** | Can specify TTL and timestamp                  | Can specify TTL and timestamp                 |
| **Use with Static Columns** | Static columns can be set                        | Static columns can be updated                  |
