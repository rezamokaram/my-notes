# Global Secondary Indexes

 (also called “Secondary Indexes”) are another mechanism in ScyllaDB which allows efficient searches on non-partition keys by creating an index. Rather than creating an index on the entire partition key, this index is created on specific columns. Each Secondary Index indexes one specific column. In cases where you are using a composite (compound) partition key, a secondary index can index the column. Secondary indexes are transparent to the application. Queries have access to all the columns in the table, and indexes can be added or removed on the fly without changing the application.

Updates therefore can be more efficient with Secondary Indexes than with Materialized Views because only changes to the primary key and indexed column cause an update in the Secondary Index view.

What’s more, the size of an index is proportional to the size of the indexed data. As data in ScyllaDB is distributed across multiple nodes, it’s impractical to store the whole index on a single node, as it limits the size of the index to the capacity of a single node, not the capacity of the entire cluster.

Hence the name Global Secondary Indexes. With global indexing, a Materialized View is created for each index. This Materialized View has the indexed column as a partition key, and it also stores the base table primary key. This means that it’s possible to query by the indexed column. Under the hood, ScyllaDB will query the MV, get the base table primary key, and then fetch the requested column.

Global Secondary indexes provide a further advantage: it’s possible to use the indexed column’s value to find the corresponding index table row in the cluster, so reads are scalable. Note, however, that with this approach, writes are slower than with local indexing (described below) because of the overhead required to keep the indexed view up to date.

## in cql

CQL supports creating secondary indexes on tables, allowing queries on the table to use those indexes. A secondary index is identified by a name defined by:

```cql
index_name: re('[a-zA-Z_0-9]+')
```

Creating a secondary index on a table uses the CREATE INDEX statement:

```cql
create_index_statement: CREATE INDEX [ `index_name` ]
                      :     ON `table_name` '(' `index_identifier` ')'
                      :     [ USING `string` [ WITH OPTIONS = `map_literal` ] ]
index_identifier: `column_name`
                :| ( FULL ) '(' `column_name` ')'
```

For instance:

```cql
CREATE INDEX userIndex ON NerdMovies (user);
CREATE INDEX ON Mutants (abilityId);
```

The CREATE INDEX statement is used to create a new (automatic) secondary index for a given (existing) column in a given table. A name for the index itself can be specified before the ON keyword, if desired. If data already exists for the column, it will be indexed asynchronously. After the index is created, new data for the column is indexed automatically at insertion time.

## another explanation

What is a Global Secondary Index (GSI) in ScyllaDB?
In databases like ScyllaDB, data is usually stored and looked up using the Primary Key. But sometimes, you want to search for data using a different column (not the primary key). To do this, you can use a Secondary Index.

Types of Secondary Indexes:

- Local Secondary Index (limited to one node)
- Global Secondary Index (GSI) (spread across all nodes)

ScyllaDB only uses Global Secondary Indexes.

Imagine you have a table for users:

```cql
CREATE TABLE users (
  id UUID PRIMARY KEY,
  name text,
  email text
);
```

Now you want to find a user by their email. But email is not the primary key, so ScyllaDB can't find it directly.

So you create a GSI like this:

```cql
CREATE INDEX ON users (email);
```

What happens:

- ScyllaDB creates a hidden table behind the scenes.
- This table links email → id.
- When you search by email, ScyllaDB first looks in the index table to get the user’s id, then uses that to fetch the full data from the original table.

---

Why is it called “Global”?
Because the index is spread across all nodes in the cluster, not just one. So any node can use the index to look up data.

---

***Pros:***

- You can query by columns that are not part of the primary key.

***Cons:***

- Every time you insert or update data, the index also needs to be updated — which adds extra cost.

- At large scale, it might impact performance.
