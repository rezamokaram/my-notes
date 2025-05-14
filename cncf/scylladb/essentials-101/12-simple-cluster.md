# simple cluster

## create cluster

```sh
docker run --name scylla-node1 -d scylladb/scylla:5.2.0
docker run --name scylla-node2 -d scylladb/scylla:5.2.0 --seeds="$(docker inspect --format='{{ .NetworkSettings.IPAddress }}' scylla-node1)"
docker run --name scylla-node3 -d scylladb/scylla:5.2.0 --seeds="$(docker inspect --format='{{ .NetworkSettings.IPAddress }}' scylla-node1)"
```

## status check

Let’s check the status of the nodes that we just set up. To do that, we use a tool called Nodetool Status.

```sh
docker exec -it scylla-node3 nodetool status
```

## sample cql

```sh
CREATE KEYSPACE mykeyspace WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'replication_factor' : 3};
use mykeyspace;
DESCRIBE KEYSPACE mykeyspace;
CREATE TABLE users ( user_id int, fname text, lname text, PRIMARY KEY((user_id)));
insert into users(user_id, fname, lname) values (1, 'rick', 'sanchez');
insert into users(user_id, fname, lname) values (4, 'rust', 'cohle');
select * from users;
```

## shards check

To check the number of physical cores on the server, and how each map to a ScyllaDB shard, run the following from any server running ScyllaDB:

```sh
docker exec -it scylla-node3 bash
./usr/lib/scylla/seastar-cpu-map.sh -n scylla
```
