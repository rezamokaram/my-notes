# Local Secondary Index (LSI)

- A **Local Secondary Index (LSI)** lets you efficiently query data **within a specific partition**.
- Both the **index and the base data** live on the **same node**, which improves performance.
- Example:

```cql
  CREATE TABLE menus (
    location text,
    name text,
    price float,
    dish_type text,
    PRIMARY KEY(location, name)
  );

  CREATE INDEX ON menus((location), dish_type);
```

With the above setup, this query:

```cql
SELECT * FROM menus WHERE location = 'Tehran' AND dish_type = 'Ash reshteh';
```

will only scan the data on the node responsible for 'Tehran', not the entire cluster.

Synchronous updates: In recent ScyllaDB versions, LSI updates happen immediately when the base table is updated. This ensures consistency between the index and the main data.

LSI is great when your queries are focused on filtering within the same partition key, offering better speed and reduced network overhead.
