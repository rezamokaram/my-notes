# rebalancing

rebalancing refers to the process of redistributing data across nodes in the cluster when the topology changes — such as when a node is added or removed. This ensures that the data is evenly distributed and each node holds its fair share of the dataset.

## Key Concepts Before Rebalancing

1. Token Ring: ScyllaDB uses consistent hashing to divide the data into token ranges, which are then assigned to nodes.

2. Replication Factor (RF): Determines how many copies of data exist across the cluster.

3. Virtual Nodes (vnodes): Each node holds multiple small token ranges (vnodes), which simplifies rebalancing.

## What Happens When a Node is Added

When you add a new node:

### 1. Token Ranges Assigned

- The new node is assigned a set of token ranges (based on vnodes).

- These token ranges previously belonged to existing nodes.

### 2. Data Streaming (Rebalancing)

- Data that falls within the new node’s token ranges is streamed from the existing nodes to the new node.

- This is done in parallel and non-blocking, meaning the cluster stays online.

### 3. Cluster Becomes More Balanced

- Each node ends up holding roughly the same amount of data.
- This helps with performance and storage efficiency.

## What Happens When a Node is Removed

When you remove a node (via decommissioning):

### 1. Token Ranges are Redistributed

- The token ranges that were assigned to the removed node are now taken over by other nodes.

### 2. Data is Streamed Out

- Before the node is shut down, it streams its data to the new owners of those token ranges.

### 3. No Data Loss

- Thanks to the replication factor, copies of data already exist elsewhere.

- The cluster automatically ensures consistency and repairs if needed.

## Manual vs Automatic Rebalancing

### Scylla handles rebalancing automatically in many cases

### For manual tuning, you can use

- nodetool status to monitor the ring
- nodetool cleanup to remove old data on nodes after rebalancing
- nodetool repair to fix any inconsistencies

## Summary

| Action        | Description                                                             |
|---------------|-------------------------------------------------------------------------|
| Node Added    | - New token ranges are assigned to the node                             |
|               | - Data is streamed from existing nodes to the new node                  |
|               | - Cluster becomes more balanced                                         |
| Node Removed  | - Token ranges are redistributed among remaining nodes                  |
|               | - Data is streamed out to new token range owners                        |
|               | - No data loss due to replication                                       |
| Rebalancing   | - Ensures even data distribution across nodes                           |
|               | - Improves performance and storage efficiency                           |
| Tools         | - `nodetool status`, `nodetool cleanup`, `nodetool repair` for checks   |
