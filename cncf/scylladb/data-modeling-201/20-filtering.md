# Filtering

Filtering is a feature that allows you to filter by column that is not part of the primary key without creating an index or MV and without any storage overhead. Filtering can result in really low query performance because it involves a full table-scan. Still, filtering can be really useful if you don’t have high performance requirements for certain queries or if the result of the query returns most of the rows from the table.

```cql
SELECT * FROM keyspace_name.table_name 
WHERE non_partition_column = 'some_value'
ALLOW FILTERING;
```
