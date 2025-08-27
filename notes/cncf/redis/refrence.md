- [Redis Reference](#redis-reference)
- [Eviction](#eviction)
  - [Configuring Eviction Policies](#configuring-eviction-policies)
- [Client Handling](#client-handling)

# Redis Reference
Specifications and protocols

# Eviction  

Redis employs various eviction policies to manage its in-memory data store when the configured memory limit (maxmemory) is reached. These policies determine which keys to remove to free up space for new data. Here's an overview of the available eviction policies:

1. noeviction

- Description: When the memory limit is reached, Redis returns an error for write operations, preventing new data from being added.
- Use Case: Suitable when it's critical to retain all existing data, and the application can handle write operation failures gracefully.  

2. allkeys-lru

- Description: Evicts the least recently used (LRU) keys, regardless of whether they have an expiration set.
- Use Case: Ideal for caching scenarios where recently accessed data is more likely to be reused.  

3. volatile-lru

- Description: Evicts the least recently used (LRU) keys among those that have an expiration (TTL) set.
- Use Case: Useful when you want to prioritize keeping persistent data while allowing expirable data to be evicted first.  

4. allkeys-random

- Description: Evicts random keys from the dataset, regardless of expiration.
- Use Case: Applicable when any key can be safely removed, and there's no specific access pattern.  

5. volatile-random

- Description: Evicts random keys among those with an expiration set.  
- Use Case: Suitable when you want to randomly remove expirable data without affecting persistent keys.  

6. volatile-ttl

- Description: Evicts keys with the shortest time-to-live (TTL) first among those with an expiration set.
- Use Case: Beneficial when you prefer to remove keys that are closest to their natural expiration.  

7. allkeys-lfu

- Description: Evicts the least frequently used (LFU) keys, regardless of expiration.
- Use Case: Effective when you want to retain frequently accessed data and remove infrequently accessed keys.  

8. volatile-lfu  

- Description: Evicts the least frequently used (LFU) keys among those with an expiration set.
- Use Case: Ideal for scenarios where you want to evict infrequently accessed expirable data first.  

## Configuring Eviction Policies

To set the eviction policy in Redis, you can modify the maxmemory-policy directive in the redis.conf file or use the CONFIG SET command at runtime. For example, to set the eviction policy to allkeys-lru, you can execute:

``` bash
    CONFIG SET maxmemory-policy allkeys-lru
```

Choosing the appropriate eviction policy depends on your application's specific requirements and access patterns. Understanding these policies helps in optimizing Redis's performance and ensuring efficient memory utilization.  

# Client Handling

TODO