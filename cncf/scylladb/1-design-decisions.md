# ScyllaDB vs. Cassandra Design Decisions

ScyllaDB was designed as a high-performance, drop-in replacement for Cassandra. While sharing the wide-column NoSQL data model, CQL, and SSTable format, their internal architectures diverge significantly.

**1. Core Language and Architecture:**

- **Cassandra:** Java-based, runs on the JVM, thread-per-core model.
- **ScyllaDB:** C++-based, utilizes the Seastar framework with a share-nothing, asynchronous, shard-per-core architecture.

  - **Impact:** ScyllaDB achieves higher throughput and lower latencies due to C++ and the shard-per-core design, which minimizes resource contention and maximizes parallelism. Cassandra's JVM can experience garbage collection pauses affecting latency.

**2. Resource Management and Scheduling:**

- **Cassandra:** Relies on the JVM for memory management and thread scheduling.
- **ScyllaDB:** Fine-grained control over CPU, memory, and I/O via its close-to-the-metal design. Implements its own I/O scheduler prioritizing user operations over background tasks.

  - **Impact:** ScyllaDB offers more predictable low latency by avoiding JVM GC and efficiently managing resources. Cassandra can have latency spikes during background operations.

**3. Scalability and Elasticity:**

- **Cassandra:** Horizontal scalability through adding nodes, but management and tuning at scale can be complex.
- **ScyllaDB:** Designed for vertical and horizontal scalability, optimized for multi-core servers with auto-tuning capabilities.

  - **Impact:** ScyllaDB often requires fewer nodes for similar performance, lowering TCO. Self-tuning simplifies operations.

**4. Workload Prioritization:**

- **Cassandra:** Limited control over resource allocation for different workloads.
- **ScyllaDB:** Built-in workload prioritization to control resource competition (e.g., real-time reads vs. batch writes).

  - **Impact:** ScyllaDB can ensure latency-sensitive applications get necessary resources even under heavy load.

**5. Lightweight Transactions (LWT):**

- **Cassandra:** Supports LWT using Paxos.
- **ScyllaDB:** Uses Paxos for LWT with optimizations reducing round trips.

  - **Impact:** ScyllaDB's LWT is generally faster with lower latency.

**6. Development Language and Ecosystem:**

- **Cassandra:** Large and mature Java ecosystem, extensive tools and community.
- **ScyllaDB:** C++ offers performance benefits. Smaller but growing community. Compatible with CQL and SSTable, allowing use of many Cassandra drivers and tools.

## Shard-Per-Core Architecture

The shard-per-core architecture dedicates specific CPU cores to handle independent data shards, minimizing resource contention and maximizing parallelism.

1. **Sharding:** Data is divided into independent chunks (shards), typically matching the number of CPU cores.
2. **Core Assignment:** Each shard is exclusively managed by a dedicated thread on a specific CPU core.
3. **Dedicated Resources:** Each core has its own memory queues, network connections, and disk I/O channels.
4. **Message Passing:** Cores communicate about data on different shards using efficient, asynchronous message passing.

**Benefits:** Reduced contention, improved cache locality, elimination of context switching (within a shard), increased parallelism, and predictable performance.

## Message Passing in Shard-Per-Core

Message passing is the communication mechanism between isolated CPU cores (each managing a data shard) in a shard-per-core architecture.

1. **Explicit Communication:** Cores send explicit messages for data requests, updates, coordination, and control.
2. **Asynchronous Nature:** Sending cores don't necessarily wait for a response, allowing for continuous processing.
3. **Queues:** Messages are sent to message queues associated with the target core for orderly processing.
4. **Content of Messages:** Messages contain data, instructions, or coordination signals.
5. **Efficiency:** Frameworks like Seastar in ScyllaDB optimize message passing for minimal overhead (e.g., zero-copy transfers).

**Why Preferred:** Avoids shared state complexity, reduces contention, improves scalability, enhances fault isolation.

## Message Passing Tech Stack (ScyllaDB)

ScyllaDB's message passing primarily relies on the **Seastar framework**, written in **C++**, providing its own asynchronous, shared-nothing communication primitives optimized for high performance.
