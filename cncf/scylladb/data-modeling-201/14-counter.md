# counter

Counters are useful for any application where you need to increment a count.

It’s a special data type (column) that only allows its value to be incremented, decremented, read or deleted. As a type, counters are a 64-bit signed integer. Updates to counters are atomic, making them perfect for counting and avoiding the issue of possible concurrent updates on the same value.

Counters can only be defined in a dedicated table that includes:

- The primary key (can be compound)
- The counter column

All non-counters columns in the table must be part of the primary key. Counter columns cannot be part of the primary key. make sure you are in the same active cqlsh prompt you created before:

```cql
CREATE TABLE pet_type_count (pet_type  text PRIMARY KEY, 
pet_counter counter);
```

Loading data to a counter table is different than other tables, it’s done with an UPDATE operation.

```cql
UPDATE pet_type_count SET pet_counter = pet_counter + 6 WHERE pet_type = 'dog';
```
