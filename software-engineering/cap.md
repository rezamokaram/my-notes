# CAP Theorem

The CAP Theorem, also known as Brewer's theorem, is a principle that applies to distributed data systems. It was formulated by computer scientist Eric Brewer in 2000 and later proved as a theorem. The CAP Theorem states that in a distributed data store, it is impossible to simultaneously guarantee all three of the following properties:

1. Consistency: Every read receives the most recent write for a given piece of data. In other words, all nodes in the system see the same data at the same time.

2. Availability: Every request (read or write) receives a response, regardless of whether the data is current or not. This means the system is operational and can serve requests even if some nodes are unreachable.

3. Partition Tolerance: The system continues to operate despite arbitrary partitioning due to network failures. This means that if some nodes cannot communicate with others, the system as a whole still functions.

According to the CAP Theorem, a distributed system can satisfy any two of these three properties but not all three at the same time. This leads to some important trade-offs when designing distributed systems:

- CP (Consistency + Partition Tolerance): In this case, the system prioritizes consistency and partition tolerance, potentially sacrificing availability. An example is a system that will refuse to process requests if it cannot ensure consistent data.

- AP (Availability + Partition Tolerance): Here, the system prioritizes availability and partition tolerance, potentially sacrificing consistency. An example is a system that allows reads and writes even when some nodes are not in sync, leading to eventual consistency.

- CA (Consistency + Availability): This scenario is often considered impractical in a distributed system because it cannot tolerate network partitions. If a partition occurs, the system would have to sacrifice either consistency or availability.

In practice, most distributed systems are designed to be either CP or AP, based on the specific requirements of the application. Understanding the CAP Theorem helps architects and developers make informed decisions about trade-offs in distributed system design.