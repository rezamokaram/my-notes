# ups and give a simple shell for scylladb

- sudo docker pull scylladb/scylla:2025.1
- sudo docker run scylladb -d --overprovisioned 1 --smp 1
- docker exec -it scylladb-name cqlsh
  - create keyspace
  - create table
  - insert data
