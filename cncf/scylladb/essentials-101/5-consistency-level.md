# Consistency Level (CL) in ScyllaDB

In **ScyllaDB**, the **Consistency Level (CL)** determines how many replicas in the cluster must acknowledge a read or write operation before it is considered successful.

## Key Concepts

- ScyllaDB is a **distributed NoSQL database** based on Apache Cassandra.
- Data is **replicated** across multiple nodes for high availability and fault tolerance.
- **Consistency Level (CL)** lets you choose the trade-off between **consistency**, **availability**, and **latency**.

---

## Common Consistency Levels in ScyllaDB

| Consistency Level | Description |
|-------------------|-------------|
| **ONE**           | Only **one** replica must respond. Fast but less consistent. |
| **QUORUM**        | A **majority** of replicas must respond. Good balance. |
| **LOCAL_QUORUM**  | A majority of replicas in the **local datacenter** must respond. Reduces latency in multi-DC setups. |
| **ALL**           | **All** replicas must respond. Strongest consistency, lowest availability. |
| **ANY** (writes only) | Allows **any node** (even a hinted handoff) to acknowledge. Weakest consistency, highest availability. |

---

## 📝 Example

If a keyspace is configured with **replication factor 3**, then:

- **CL = QUORUM** means 2 out of 3 replicas must respond.
- **CL = ONE** means just 1 replica is enough.
- **CL = ALL** means all 3 must confirm the operation.

---

Choosing the right CL depends on your app's need for **speed vs. data correctness**.

## Can You Set Consistency Level (CL) Per Operation in ScyllaDB?

Yes, in **ScyllaDB**, you **can set the Consistency Level (CL) per operation** — both for **reads** and **writes**. This flexibility allows you to fine-tune the balance between **consistency**, **latency**, and **availability** based on the criticality of each query.

---

## ✅ How to Set CL Per Operation

### 1. Using CQL (Cassandra Query Language)

You can specify the consistency level with the `CONSISTENCY` command before a query:

```sql
CONSISTENCY QUORUM;
SELECT * FROM users WHERE id = 123;
```

Or set it directly in your driver configuration for each query (see below).

### Example in Drivers (Go / Golang)

To interact with ScyllaDB in Go, use the `gocql` driver, which supports setting consistency levels per query.

#### Setup Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/gocql/gocql"
)

func main() {
    // Create a cluster configuration
    cluster := gocql.NewCluster("127.0.0.1") // replace with your ScyllaDB node IP
    cluster.Keyspace = "my_keyspace"
    cluster.Consistency = gocql.Quorum // Default consistency, can be overridden per query

    session, err := cluster.CreateSession()
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer session.Close()

    // Execute a SELECT query with per-query consistency
    query := session.Query("SELECT * FROM users WHERE id = ?", 123).
        Consistency(gocql.LocalQuorum)

    var id int
    var name string

    if err := query.Scan(&id, &name); err != nil {
        log.Fatalf("Query failed: %v", err)
    }

    fmt.Printf("User: %d, Name: %s\n", id, name)
}
```
