# common data types

ScyllaDB supports different data types which can be used to define columns in a table. Some of the most commonly used ones are:

- Text: a UTF8 encoded string (same as varchar)
- Int:  a 32-bit signed integer
- UUID: a universal unique identifier, generated via uuid(), for example 123e4567-e89b-12d3-a456-426655440000
- TIMEUUID: a version 1 UUID, generally used as a “conflict-free” timestamp, generated using now()
- TIMESTAMP: A timestamp (date and time) with millisecond precision, stored as a 64-bit integer. Displayed in cqlsh as yyyy-mm-dd HH:mm:ssZ
[A full reference can be found here](https://opensource.docs.scylladb.com/stable/cql/types.html).

---

[Collections are covered in depth here](https://opensource.docs.scylladb.com/stable/cql/types.html). Some of the basic concepts:

- Collections are multi-valued columns
- To access a single element of a collection, the whole collection has to be read
- Meant to store a relatively small amount of data. They work well for things like “the phone numbers of a given user”, but when items are expected to grow unbounded (“all messages sent by a user”, “events registered by a sensor”…), then collections are not appropriate, and a specific table (with clustering columns) should be used.

## CQL supports three kinds of collections

### Maps

A map is a (sorted) set of key-value pairs, where keys are unique, and the map is sorted by its keys. Both the key and the value have a type. For example, let’s create a table of pets with a map of favorite things they like (make sure you are in the same active cqlsh prompt you created before):

```cql
CREATE TABLE pets_v1 (
    pet_chip_id text PRIMARY KEY,
    pet_name text,
    favorite_things map<text, text> // A map of text keys, and text values
);
```

Now for Rocky, the favorite food is Turkey, and the favorite toy is Tennis ball

```cql
INSERT INTO pets_v1 (pet_chip_id, pet_name, favorite_things)
           VALUES ('123e4567-e89b-12d3-a456-426655440b23', 'Rocky', { 'food' : 'Turkey', 'toy' : 'Tennis Ball' });
```

### Sets

A set is a collection of unique values. It is stored unordered but retrieved in sorted order. For example, we might have a table of pets, where each pet might have more than one vaccination.

```cql
CREATE TABLE pets_v2 (
    pet_name text PRIMARY KEY,
    address text,
    vaccinations set<text> 
);
```

```cql
INSERT INTO pets_v2 (pet_name, address, vaccinations)
            VALUES ('Rocky', '11 Columbia ave, New York NY', { 'Heartworm', 'Canine Hepatitis' });
```

When using Sets, the elements are naturally sorted. If a specific order is required, say insertion order, a List might be preferable. The values are unique.

### Lists

A list is a (sorted) collection of non-unique values where elements are ordered by their position in the list. Items are inserted/retrieved according to an index. Values are not necessarily unique.

```cql
CREATE TABLE pets_v3 (
    pet_name text PRIMARY KEY,
    address text,
    vaccinations list<text>
);
```

```cql
INSERT INTO pets_v3 (pet_name, address, vaccinations)
            VALUES ('Rocky', '11 Columbia ave, New York NY',  ['Heartworm', 'Canine Hepatitis', 'Heartworm']);
```

Lists have limitations and specific performance considerations that you should take into account before using them. In general, if you can use a set instead of a list, always prefer a set.
