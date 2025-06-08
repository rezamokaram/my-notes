
# ScyllaDB Change Data Capture (CDC) - Summary

## What is CDC?

Change Data Capture (CDC) is a mechanism for tracking and recording changes—such as inserts, updates, and deletes—on one or more tables in a database. It provides a data-centric view of modifications, enabling asynchronous consumption and integration into various analytics and processing pipelines. Unlike traditional application-level triggers, CDC captures changes at the database level, ensuring consistency and scalability.

## Key Features of ScyllaDB’s CDC

### Per-Table Enablement

- CDC is **enabled per table**, allowing fine-grained control and avoiding unnecessary overhead on tables that don’t require it.

### Granular Data Capture

- Changes are tracked at the **CQL row (clustered row)** level.
- Captures can include:
  - **Preimage**: Data before the change.
  - **Delta**: The change itself.
  - **Postimage**: Data after the change.
- All options (preimage, delta, postimage) are **configurable**, offering flexibility based on user needs.

### CDC as a Table

- The captured changes are stored in a **separate CDC table**.
- This CDC table is fully integrated into the ScyllaDB cluster:
  - Shares the same **distribution**, **consistency levels**, and **topology** as regular data tables.
  - Ordered by **timestamp** and operation **sequence**.
  - Includes metadata like TTL (Time-To-Live) and write context.

### Time-Based Log Retention

- CDC data is **transient** by design.
- The default **retention window** is 24 hours (configurable), ensuring that unconsumed logs don’t overwhelm storage.

### Asynchronous Consumption

- Designed for downstream consumers to **batch**, **partition**, and **process** changes independently of the main application.
- Enables powerful use cases like:
  - **Fraud detection**
  - **Real-time analytics**
  - **Event-driven architectures (e.g., Kafka pipelines)**
  - **Data replication or mirroring**
  - **ETL and transformation into other storage systems**

## Technical Considerations

### Consistency Model

- CDC reflects the **eventual consistency** of ScyllaDB.
- Changes are captured **as seen by the client**, meaning the log represents what was acknowledged at write time—not necessarily the final reconciled state across replicas.

### Failure Scenarios

- In distributed systems, **node failures** can lead to **partial logs**.
- If a write only reached a single replica which later failed, that change may not be fully captured in the CDC log.

### Performance Impact

- CDC adds minimal overhead since it uses existing mechanisms like write paths and data distribution.
- Optional features like pre/postimage require **read-before-write** or **read-after-write**, which can slightly impact performance.

## Conclusion

ScyllaDB’s CDC feature provides a robust and efficient mechanism to capture and react to data changes within the database. It's highly configurable, natively integrated, and designed for modern use cases involving real-time data flow and analytics. Despite some trade-offs inherent in distributed systems (like partial logs or eventual consistency), it offers significant power for building reactive systems, ensuring data traceability, and enabling sophisticated analytics.
