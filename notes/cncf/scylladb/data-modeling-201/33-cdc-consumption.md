
# Consuming CDC Data in ScyllaDB

ScyllaDB’s Change Data Capture (CDC) feature is designed for simplicity and flexibility in consumption. Once CDC is enabled on a table, all change records are written to a **CDC log table**, which behaves like any other ScyllaDB table. This makes consumption straightforward and consistent with normal CQL-based operations.

## Basic Consumption

- **Access via CQL**:  
  The CDC log is just a table, so it can be queried using **standard CQL (Cassandra Query Language)**. This means you can:
  - Use `SELECT` statements to pull records.
  - Apply filters, limits, and pagination as needed.
  - Batch fetch and process changes efficiently.

- **Pre-Reconciled Data**:  
  CDC entries are already **reconciled at the coordinator level**, so consumers don't need to deal with reconciling data from multiple replicas. This simplifies application logic and ensures consistency.

## Advanced Integration Options

While consuming CDC data through CQL is simple, more complex use cases often require integration into processing pipelines or external systems. ScyllaDB is building **multiple layers of abstraction** and **integration tools** to support this.

### Integrators and Adapters

An **integrator** is a component or service that bridges ScyllaDB’s CDC with other systems, formats, or processing models. These can:

- **Transform CDC records** into formats like JSON, Avro, or Protobuf.
- **Stream changes** to external systems (e.g., Kafka, Flink, Spark).
- **Apply custom business logic** in flight.
- **Buffer and throttle** consumption for performance management.

**Examples**:

- **Kafka connector**: Streams CDC events into Kafka topics.
- **Lambda-style adapter**: Triggers serverless functions on CDC events.
- **ETL pipelines**: Moves changes to data warehouses or lakes (e.g., Redshift, BigQuery, S3).

### Push-Based Consumption

- ScyllaDB plans to support **push-based models**, where changes are automatically pushed to consumers rather than pulled via queries.
- This allows for **real-time streaming** and minimizes latency for time-sensitive use cases like fraud detection or alerts.

### Alternator Integration

- **Alternator** is ScyllaDB’s API-compatible implementation of Amazon DynamoDB.
- CDC will be integrable via the **Alternator API**, enabling DynamoDB users to adopt CDC without rewriting their application logic.

## Built-In Tooling and Ecosystem Support

ScyllaDB is actively developing a **suite of official connectors and integrations**, reducing reliance on third-party tools. Planned or existing tooling includes:

- Kafka Connect support
- Sink connectors for cloud storage
- Native SDKs for CDC consumption
- Dashboards and observability tools for CDC stream health

## Summary

ScyllaDB's CDC consumption is designed to be:

- **Simple** (via direct CQL queries)
- **Flexible** (supporting various models like batch, streaming, and real-time)
- **Extensible** (with integrators and adapters)
- **Robust** (featuring pre-reconciled, structured data)

Whether you're building a streaming analytics pipeline, synchronizing data to another system, or triggering event-driven workflows, ScyllaDB provides the primitives and ecosystem to support it—all without needing external middleware if you choose to use their native solutions.
