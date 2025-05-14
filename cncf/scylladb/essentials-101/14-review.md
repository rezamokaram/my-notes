# Summary

| component | description |
|-----------|-------------|
| `Replication Factor (RF)` | Determines how many copies of the data will exist |
| `Consistency Level (CL)` | Determines how many nodes need to answer a query for it to be successful |
| `Shard` | Smallest processing unit in Scylla, represents a CPU where requests are executed |
| `Node` | A server where data is stored, may contain multiple shards |
| `Ring` | Distribution of data evenly on a cluster of nodes  |
| `Cluster` | A collection of nodes |
| `Data Center` | A physical facility that houses servers, storage and networking equipment  |
| `Rack` | A metal frame used to hold hardware devices such as servers, hard disks, networking equipment, and other electronic equipment |

- ScyllaDB has a ring-type architecture
- It’s a distributed, highly available, high performance, low maintenance, highly scalable NoSQL database
- In ScyllaDB all nodes are created equal, there are no master/slave nodes
- Data is automatically distributed and replicated on the cluster according to the replication strategy
- ScyllaDB supports multiple data centers
- ScyllaDB transparently partitions data by using the hash values of keys
