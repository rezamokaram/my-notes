# cql

## what is

### cassandra query language (cql)

### data definition (DDL)

#### create / delete / alter keyspace

- `create`

    ```cql
    CREATE KEYSPACE my_keyspace
    WITH REPLICATION = {
        'class' : 'NetworkTopologyStrategy',
        'dc1' : 3,
        'dc2' : 2
    };
    ```

- `delete / drop`

    ```cql
    DROP KEYSPACE my_keyspace;
    ```

- `alter`

    ```cql
    ALTER KEYSPACE my_keyspace 
        WITH REPLICATION = {
        'class' : 'SimpleStrategy',
        'replication_factor' : 4
    };
    ```

#### create / delete / alter table

- `create`

    ```cql
    CREATE TABLE sales_by_region (
    region text,
    country text,
    sale_date date,
    amount decimal,
    PRIMARY KEY ((region, country), sale_date)
    ) WITH CLUSTERING ORDER BY (sale_date DESC);
    ```

- `delete / drop`

    ```cql
    DROP TABLE products;
    ```

- `alter`

    ```cql
    -- Add a column
    ALTER TABLE products ADD product_name text;

    -- Drop a column
    ALTER TABLE products DROP description;

    -- Change column type
    ALTER TABLE products ALTER price TYPE decimal;

    -- Modify table properties
    ALTER TABLE products WITH compression = {
    'class' : 'LZ4Compressor',
    'chunk_length_kb' : 64 };
    ```

### dml

#### table management

```cql
-- Truncate table (remove all data)
TRUNCATE users;

-- Describe table schema
DESCRIBE TABLE users;

-- Get table status
SHOW CREATE TABLE users;
```

#### simple crud

```cql
-- Insert data
INSERT INTO users (user_id, email, last_login) 
VALUES (uuid(), 'user@example.com', toTimestamp(now()));

-- Update data
UPDATE users SET last_login = toTimestamp(now()) 
WHERE user_id = some_uuid();

-- Delete data
DELETE FROM users WHERE user_id = some_uuid();

-- Select all columns
SELECT * FROM users;

-- Select specific columns
SELECT user_id, email FROM users;

-- Filter by partition key
SELECT * FROM users WHERE user_id = some_uuid();
```

#### range queries

```cql
-- Range query on clustering column
SELECT * FROM sales 
WHERE store_id = 1 
AND sale_date >= '2023-01-01' 
AND sale_date <= '2023-12-31';

-- Allow filtering with secondary index
CREATE INDEX IF NOT EXISTS idx_email ON users(email);
SELECT * FROM users WHERE email = 'user@example.com';
```

#### batch

```cql
BEGIN BATCH
    INSERT INTO users (user_id, email) VALUES (uuid(), 'user1@example.com');
    UPDATE users SET last_login = toTimestamp(now()) WHERE user_id = uuid();
    DELETE FROM users WHERE user_id = some_uuid();
APPLY BATCH;
```

## cqlsh

```bash
cqlsh <node-ip>
```

## data types

***Read more on docs***

### Native Types

- ascii (US-ASCII character string)
- bigint (64-bit signed long)
- blob (Arbitrary bytes)
- boolean (true/false)
- counter (64-bit signed integer)
- date (Date without time)
- decimal (Variable-precision decimal)
- double (64-bit IEEE-754 floating point)
- duration (Time duration with nanosecond precision)
- float (32-bit IEEE-754 floating point)
- inet (IP address)
- int (32-bit signed integer)
- smallint (16-bit signed integer)
- text (UTF8 encoded string)
- time (Time without date)
- timestamp (Date and time)
- timeuuid (Version 1 UUID)
- tinyint (8-bit signed integer)
- uuid (UUID)
- varchar (UTF8 encoded string)
- varint (Arbitrary-precision integer)

### Collection Types

- list (Ordered collection of elements)
- map (Key-value pairs)
- set (Unordered collection of unique elements)

### Special Types

- frozen (For UDTs and collections)
- tuple (Group of 2-3 fields)
- user-defined types (UDTs)
