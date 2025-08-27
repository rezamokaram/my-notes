# TTL

The TTL can be set when defining a Table or when using the INSERT and UPDATE  queries, as we’ll see in the examples below.

We’ll go over different use cases for using TTL, defining it on an entire Table, on a column, resetting it, and we’ll also see what happens when different columns have different TTL values.

TTL is measured in seconds. If the field is not updated within the TTL, it is deleted.

The expiration works at the individual column level, which provides a lot of flexibility.

By default, the TTL value is null, which means that the data will not expire.

Define the following table, make sure you are in the same active cqlsh prompt you created before:

```cql
CREATE TABLE heartrate (
    pet_chip_id  uuid,
    name text,
    heart_rate int,
    PRIMARY KEY (pet_chip_id));
```

```cql
INSERT INTO heartrate(pet_chip_id, name, heart_rate) VALUES (123e4567-e89b-12d3-a456-426655440b23, 'Duke', 90);
```

Now let’s use the TTL() function to retrieve the TTL Value for Duke’s heart_rate:

```cql
SELECT name, TTL(heart_rate)
FROM heartrate WHERE  pet_chip_id = 123e4567-e89b-12d3-a456-426655440b23;
```

As no TTL value was defined for the table, we can see that the TTL for Duke’s heart_rate is null, which means that the data will not expire.

## TTL using UPDATE and INSERT

Let’s set the TTL value using the UPDATE query on the heart_rate column to 10 minutes (600 seconds):

```cql
UPDATE heartrate USING TTL 600 SET heart_rate =
110 WHERE pet_chip_id = 123e4567-e89b-12d3-a456-426655440b23;
```

Now let’s check the TTL using the TTL() function again:

```cql
SELECT name, heart_rate, TTL(heart_rate)
FROM heartrate WHERE pet_chip_id = 123e4567-e89b-12d3-a456-426655440b23;
```

We can see that the TTL has a value lower than 600 as a few seconds passed between setting the TTL and performing the SELECT query.

If we’d wait 10 minutes and rerun this command, Duke’s heart_rate value would have expired, and it would have a null value.

It’s also possible to set the TTL when performing an INSERT. Let’s do this with a TTL of 30 seconds:

```cql
INSERT INTO heartrate(pet_chip_id, name, heart_rate) VALUES (c63e71f0-936e-11ea-bb37-0242ac130002, 'Rocky', 87) USING TTL 30;
```

```cql
SELECT name, heart_rate FROM heartrate WHERE pet_chip_id = c63e71f0-936e-11ea-bb37-0242ac130002;
```

Initially, the row is there. Wait for 30 seconds and perform the query again to see that it’s removed:

```cql
SELECT name, heart_rate FROM heartrate WHERE pet_chip_id = c63e71f0-936e-11ea-bb37-0242ac130002;
```

## Remove TTL

To remove the TTL value on a given column, set it to 0. Let’s see an example. Insert a new row with a TTL of 400:

```cql
INSERT INTO heartrate(pet_chip_id, name, heart_rate) VALUES (c63e71f0-936e-11ea-bb37-0242ac130002, 'Rocky', 117) USING TTL 400;
```

Now check the TTL:

```cql
SELECT name, heart_rate, TTL(name) as name_ttl, TTL(heart_rate) as heart_rate_ttl FROM heartrate WHERE pet_chip_id = c63e71f0-936e-11ea-bb37-0242ac130002;
```

Next, we’ll remove the TTL from the heart_rate column, using the UPDATE command:

```cql
UPDATE heartrate USING TTL 0 SET heart_rate =
150 WHERE pet_chip_id = c63e71f0-936e-11ea-bb37-0242ac130002;
```

```cql
SELECT name, heart_rate, TTL(name) as name_ttl, TTL(heart_rate) as heart_rate_ttl FROM heartrate WHERE pet_chip_id = c63e71f0-936e-11ea-bb37-0242ac130002;
```

## different ttl on 1 row scenario

In ScyllaDB (just like in Apache Cassandra), TTL (Time-To-Live) is applied at the cell (column value) level, not at the row level. This means:

- Each column can have a different TTL.

- When the TTL of a specific column expires, only that column is deleted, not the entire row.

Assume you have a table like this:

```cql
CREATE TABLE example (
    id int PRIMARY KEY,
    col1 text,
    col2 text
);
```

Then you insert values with different TTLs:

```cql
INSERT INTO example (id, col1, col2) VALUES (1, 'value1', 'value2');
UPDATE example USING TTL 60 SET col1 = 'value1' WHERE id = 1;
UPDATE example USING TTL 300 SET col2 = 'value2' WHERE id = 1;
```

***What happens:***

- When TTL for col1 expires (after 60 seconds), only col1 will be deleted.

- col2 will remain available until its TTL (300 seconds) expires.

The row will not be deleted unless all columns are expired and/or explicitly deleted. However:

- If all columns in a row expire (or are deleted), the row becomes logically empty and may eventually be removed by `compaction`, but it still technically exists in the storage engine until that happens.

`What happens in this case if the data in one column expires? It will get a value of null. However, the row will only be deleted when all other (non-primary column) values expire.`

but read more about $compaction$.

If col1's TTL has expired, and you run the following query:

```cql
SELECT * FROM example WHERE id = 'given-id';
```

```cql
id     | col1 | col2
-------+------+----------
1      | null | value2
```

## TTL for a Table

Use the CREATE TABLE command and set the default_time_to_live value to 500.

```cql
CREATE TABLE heartrate_ttl (
    pet_chip_id  uuid,
    name text,
    heart_rate int,
    PRIMARY KEY (pet_chip_id))
WITH default_time_to_live = 500;
```

In this case, the TTL of  500 seconds is applied to all rows; however, keep in mind that TTL is stored on a per column level for non-primary key columns.

Now when we insert data, each column will have a TTL of 500:

```cql
INSERT INTO heartrate_ttl(pet_chip_id, name, heart_rate) VALUES (c63e71f0-936e-11ea-bb37-0242ac130002, 'Rocky', 92);
```

```cql
SELECT name, heart_rate, TTL(name) as name_ttl, TTL(heart_rate) as heart_rate_ttl FROM heartrate_ttl WHERE pet_chip_id = c63e71f0-936e-11ea-bb37-0242ac130002;
```

To update the TTL of the entire row, we have to perform an UPSERT:

```cql
INSERT INTO heartrate_ttl(pet_chip_id, name, heart_rate) VALUES (c63e71f0-936e-11ea-bb37-0242ac130002, 'Rocky', 92);
```

```cql
SELECT name, heart_rate, TTL(name) as name_ttl, TTL(heart_rate) as heart_rate_ttl FROM heartrate_ttl WHERE pet_chip_id = c63e71f0-936e-11ea-bb37-0242ac130002;
```

### To change the default_time_to_live on  an existing table, use the ALTER command

```cql
ALTER TABLE heartrate_ttl WITH default_time_to_live = 3600;
```

## Summary

Time to Live is a useful feature for setting data to expire automatically. An example use case would be a one-time password that expires after an hour. There are many other use cases.

It’s measured in seconds, measure as the time that passes from an INSERT or UPDATE command. After the TTL amount of time passes, the data is removed.

In ScyllaDB (and Apache Cassandra), TTL is applied to the column level, but it can be set for an entire row or have a default value for the table. The TTL() function can be used to get the TTL for a specific column.

To remove the TTL value, set it to 0.
