# Best Practices in Choosing Keys

## Choosing Partition Key

### try to achieve

- high cardinality
- even distribution

### try to avoid

- low cardinality
- hot partition
- large partition

## partition key examples

### good

- user name
- user id
- user id + time

### bad

- state (geographic parameter)
- age
- favorite nba team

## Clustering key

- allow useful range queries
- allow useful LIMIT
