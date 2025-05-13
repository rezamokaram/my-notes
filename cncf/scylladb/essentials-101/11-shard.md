# shard

Each ScyllaDB node consists of several independent shards, which contain their share of the node’s total data. ScyllaDB creates a one shard per core (technically, one shard per hyperthread, meaning some physical cores may have two or more virtual cores). Each shard operates on a shared-nothing architecture basis. This means each shard is assigned its RAM and its storage, manages its schedulers for the CPU and I/O, performs its compactions (more about compaction later on), and maintains its multi-queue network connection. Each shard runs as a single thread, and communicates asynchronously with its peers, without locking.

From the outside, nodes are viewed as a single object. Operations are performed at the node level.

To check the number of physical cores on the server, and how each map to a ScyllaDB shard, run the following from any server running ScyllaDB:

```sh
docker exec -it scylla-node3 bash
./usr/lib/scylla/seastar-cpu-map.sh -n scylla
```
