- [Redis Stack](#redis-stack)
  - [What is Redis Stack?](#what-is-redis-stack)
  - [Key Components of Redis Stack](#key-components-of-redis-stack)
- [Bloom Filter](#bloom-filter)
  - [Concept Of Bloom Filter](#concept-of-bloom-filter)
  - [Commands Of Bloom Filter](#commands-of-bloom-filter)
  - [Use Cases](#use-cases)
  - [Bloom Filter Trade-offs](#bloom-filter-trade-offs)
- [Cuckoo Filter](#cuckoo-filter)
  - [Concept](#concept)
  - [Commands for Redis Cuckoo Filter](#commands-for-redis-cuckoo-filter)
  - [Use Cases](#use-cases-1)


# Redis Stack

## What is Redis Stack?
Redis Stack is an extended version of Redis that bundles core Redis with additional modules to provide advanced functionality beyond simple key-value storage. It is designed for real-time applications that need capabilities like full-text search, time-series data, probabilistic data structures, and JSON storage.

## Key Components of Redis Stack
Redis Stack includes the following built-in modules:

1. RedisJSON – Stores and queries JSON documents efficiently.
2. RediSearch – Provides full-text search and secondary indexing.
3. RedisTimeSeries – Optimized for time-series data with fast inserts and queries.
4. RedisBloom – Implements probabilistic data structures (Bloom filters, Cuckoo filters, Top-K, Count-Min Sketch).
5. RedisGraph – Graph database for efficient graph queries.



# Bloom Filter

## Concept Of Bloom Filter  

A Bloom filter in Redis is a probabilistic data structure used for checking whether an element might be present or is definitely not present in a dataset. It is part of the RedisBloom module, which provides support for various probabilistic data structures.  

**How Bloom Filters Work**  
- A Bloom filter is space-efficient and fast but allows false positives (i.e., it may incorrectly say an element exists when it does not).
- It does not allow false negatives (i.e., if it says an element is not present, then it is truly not present).
- It works by hashing an input multiple times and setting bits in a bit array.
- When checking for membership, it rehashes the input and checks if all corresponding bits are set.

## Commands Of Bloom Filter  

1. BF.ADD key item  
- Adds an item to a Bloom filter.
- This command is similar to BF.MADD, except that only one item can be added.

2. BF.MADD key item [item ...]  
- Adds one or more items to a Bloom filter.
- This command is similar to BF.ADD, except that you can add more than one item.
- This command is similar to BF.INSERT, except that the error rate, capacity, and expansion cannot be specified.

3. BF.INSERT key [CAPACITY capacity] [ERROR error] [EXPANSION expansion] [NOCREATE] [NONSCALING] ITEMS item [item...]   
- Creates a new Bloom filter if the key does not exist using the specified error rate, capacity, and expansion, then adds all specified items to the Bloom Filter.
- This command is similar to BF.MADD, except that the error rate, capacity, and expansion can be specified. It is a sugarcoated combination of BF.RESERVE and BF.MADD.

4. BF.INFO key [CAPACITY | SIZE | FILTERS | ITEMS | EXPANSION]
- Returns information about a Bloom filter.

5. BF.EXISTS key item  
- Determines whether a given item was added to a Bloom filter.  
- This command is similar to BF.MEXISTS, except that only one item can be checked.  

6. BF.MEXISTS key item [item ...]  
- Determines whether one or more items were added to a Bloom filter.
- This command is similar to BF.EXISTS, except that more than one item can be checked.

7. BF.RESERVE key error_rate capacity [EXPANSION expansion] [NONSCALING]
- Creates an empty Bloom filter with a single sub-filter for the initial specified capacity and with an upper bound error_rate.
- By default, the filter auto-scales by creating additional sub-filters when capacity is reached. The new sub-filter is created with size of the previous sub-filter multiplied by expansion.

8. BF.CARD key
- Returns the cardinality of a Bloom filter  

9. BF.LOADCHUNK key iterator data  
- Restores a Bloom filter previously saved using BF.SCANDUMP.

10. BF.SCANDUMP key iterator
- Begins an incremental save of the Bloom filter.

## Use Cases
✅ Prevent Duplicate Checks (e.g., checking if a user has already seen a message)  
✅ Quick URL Blacklisting (e.g., blocking bad domains in a firewall)  
✅ Fast Membership Testing (e.g., checking if an item is in a large dataset)  
✅ Recommendation Systems (e.g., filtering out previously shown products)  

## Bloom Filter Trade-offs
🔹 Pros:  
- Very memory-efficient for large datasets.
- Fast lookups and insertions.
- Avoids unnecessary database queries.

🔹 Cons:  
- False positives are possible.
- No deletion of individual items (but you can drop the filter).  

# Cuckoo Filter  

## Concept  

A Cuckoo Filter is a space-efficient probabilistic data structure used for set membership testing, similar to a Bloom Filter. It supports insertion, lookup, and deletion, making it more flexible than Bloom Filters.  

**Why Use a Cuckoo Filter in Redis?**  
More efficient than Bloom Filters for certain use cases.
Supports deletion, unlike Bloom Filters.
Low false positive rate.
Used in fraud detection, caching, and duplicate detection.

## Commands for Redis Cuckoo Filter  

1. CF.ADD key item
- Adds ana item to the cuckoo filter.  

2. CF.ADDNX key item  
- Adds an item to a cuckoo filter if the item does not exist.  

3. CF.COUNT key item  
- Returns an estimation of the number of times a given item was added to a cuckoo filter.  

4. CF.DEL key item
- Deletes an item once from the filter.  

5. CF.EXISTS key item  
- Determines whether a given item was added to a cuckoo filter.  

6. CF.INFO key  
- Returns information about a cuckoo filter.

7. CF.INSERT key [CAPACITY capacity] [NOCREATE] ITEMS item [item ...]  
- Adds one or more items to a cuckoo filter, allowing the filter to be created with a custom capacity if it does not exist yet.  

8. CF.INSERTNX key [CAPACITY capacity] [NOCREATE] ITEMS item [item ...]  
- Adds one or more items to a cuckoo filter if they did not exist previously, allowing the filter to be created with a custom capacity if it does not exist yet.  

9. CF.LOADCHUNK key iterator data  
- restore a filter previously saved using SCANDUMP  

10. CF.MEXISTS key item [item ...]  
- Determines whether one or more items were added to a cuckoo filter.  

11. CF.RESERVE key capacity [BUCKETSIZE bucketsize] [MAXITERATIONS maxiterations] [EXPANSION expansion]  
- Creates an empty cuckoo filter with a single sub-filter for the initial specified capacity.  

12. CF.SCANDUMP key iterator  
- Begins an incremental save of the cuckoo filter.  
- This command is useful for large cuckoo filters that cannot fit into the DUMP and RESTORE model.  
- The first time this command is called, the value of iter should be 0.  
- This command returns successive (iter, data) pairs until (0, NULL) indicates completion.  

## Use Cases
✅ Fraud Prevention: Detecting duplicate transactions.  
✅ Cache Filtering: Avoid unnecessary database lookups.  
✅ Content Deduplication: Prevent duplicate messages in a messaging system.  
✅ Rate Limiting: Track API requests without storing all past requests.  

