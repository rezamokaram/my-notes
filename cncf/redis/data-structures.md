# Redis Data Structures

- [Redis Data Structures](#redis-data-structures)
- [Introduction](#introduction)
- [Strings and Basic commands](#strings-and-basic-commands)
	- [Number commands](#number-commands)
- [List](#list)
	- [List Concept](#list-concept)
	- [List Commands](#list-commands)
		- [Adding Elements to a List](#adding-elements-to-a-list)
		- [Retrieving Elements](#retrieving-elements)
		- [Removing Elements](#removing-elements)
		- [Modifying Elements](#modifying-elements)
		- [Blocking Operations (Useful for Queues)](#blocking-operations-useful-for-queues)
		- [Moving Elements Between Lists](#moving-elements-between-lists)
	- [Use Cases Of Redis Lists](#use-cases-of-redis-lists)
- [Hash](#hash)
	- [Hash Concept](#hash-concept)
	- [Hash Commands](#hash-commands)
	- [Use Cases](#use-cases)
		- [use](#use)
		- [don't use](#dont-use)
	- [Summary](#summary)
- [Pipelines](#pipelines)
	- [Concept](#concept)
		- [**Key Benefits of Using Pipelining:**](#key-benefits-of-using-pipelining)
	- [**2. How Redis Pipeline Works**](#2-how-redis-pipeline-works)
		- [**Without Pipelining (Normal Execution)**](#without-pipelining-normal-execution)
		- [**With Pipelining**](#with-pipelining)
	- [**3. Redis Pipeline Commands**](#3-redis-pipeline-commands)
		- [**Using Golang (`go-redis` package)**](#using-golang-go-redis-package)
	- [**4. Redis Pipeline vs. Transaction**](#4-redis-pipeline-vs-transaction)
	- [Example of Redis Transaction in Golang](#example-of-redis-transaction-in-golang)
	- [5. When to Use Redis Pipelining?](#5-when-to-use-redis-pipelining)
	- [6. Common Mistakes \& Best Practices](#6-common-mistakes--best-practices)
	- [Conclusion](#conclusion)
- [Set](#set)
	- [Set Concept](#set-concept)
	- [Set Commands](#set-commands)
- [Sorted Set](#sorted-set)
	- [Sorted Set Concept](#sorted-set-concept)
	- [Sorted Set Commands](#sorted-set-commands)
		- [Adding \& Updating Elements](#adding--updating-elements)
		- [Retrieving Elements](#retrieving-elements-1)
		- [Counting \& Range Queries](#counting--range-queries)
		- [Removing Elements](#removing-elements-1)
		- [Lexicographical Operations (String Sorting)](#lexicographical-operations-string-sorting)
		- [Intersection \& Union Operations](#intersection--union-operations)
		- [Miscellaneous](#miscellaneous)
- [Bitmaps](#bitmaps)
	- [Concept Of Redis Bitmaps](#concept-of-redis-bitmaps)
	- [Bitmap Commands](#bitmap-commands)
		- [Setting and Getting Bits](#setting-and-getting-bits)
		- [Counting and Analyzing Bits](#counting-and-analyzing-bits)
		- [Performing Bitwise Operations](#performing-bitwise-operations)
		- [Using Bitfields for More Control](#using-bitfields-for-more-control)
	- [Use Cases Of Redis Bitmaps](#use-cases-of-redis-bitmaps)
	- [Limitations Of Redis Bitmaps](#limitations-of-redis-bitmaps)
	- [Summary](#summary-1)
- [HyperLogLog](#hyperloglog)
	- [Concept Of HyperLogLog](#concept-of-hyperloglog)
	- [Commands Of HyperLogLog](#commands-of-hyperloglog)
	- [Use Cases Of HyperLogLog](#use-cases-of-hyperloglog)
	- [Limitations Of HyperLogLog](#limitations-of-hyperloglog)
	- [Summary](#summary-2)
- [Bitfields](#bitfields)
	- [Concept Of Bitfields](#concept-of-bitfields)
	- [Commands Of Bitfields](#commands-of-bitfields)
			- [Storing a Value (`SET`)](#storing-a-value-set)
			- [Retrieving a Value (`GET`)](#retrieving-a-value-get)
	- [Advantages of Bitfields](#advantages-of-bitfields)
	- [Limitations of Bitfields](#limitations-of-bitfields)
- [Geospatial Indexes](#geospatial-indexes)
	- [Concept Of Geospatial Indexes](#concept-of-geospatial-indexes)
	- [Commands Of Geospatial Indexes](#commands-of-geospatial-indexes)
	- [Use Cases of Geospatial Indexing](#use-cases-of-geospatial-indexing)
	
# Introduction

- redis is fast  
    1. all data stored in memory
    2. data is organized in simple data structures  
    3. redis has a simple feature set

# Strings and Basic commands
[doc link to all commands](https://redis.io/docs/latest/commands/?alpha=e)  

some of them have same functionality until we set something special.  
just for example we review the set command:

in redis docs the words with capital letters are redis command syntax and the words with small letters are the values that we should set as developer.  

**SET command**  
SET key value  
[NX | XX]  
[GET]  
[ EX seconds | PX milliseconds |
    EXAT unix-time-seconds | PXAT unix-time-milliseconds | KEEPTTL]  

*options:* 
- EX seconds -- Set the specified expire time, in seconds (a positive integer).
- PX milliseconds -- Set the specified expire time, in milliseconds (a positive integer).
- EXAT timestamp-seconds -- Set the specified Unix time at which the key will expire, in seconds (a positive integer).
- PXAT timestamp-milliseconds -- Set the specified Unix time at which the key will expire, in milliseconds (a positive integer).
- NX -- Only set the key if it does not already exist.
- XX -- Only set the key if it already exists.
- KEEPTTL -- Retain the time to live associated with the key.
- GET -- Return the old string stored at key, or nil if key did not exist. An error is returned and SET aborted if the value stored at key is not a string.

**SETEX command**  
deprecated, we need to set with exact ttl time.  

**SETNX command**  
deprecated, run SET only not exist.  

**MSET command**  
Sets the given keys to their respective values. MSET replaces existing values with new values, just as regular SET.  
MSET key value [key value ...]  

**GET command**  

GET key  

Get the value of key. If the key does not exist the special value nil is returned. An error is returned if the value stored at key is not a string, because GET only handles string values.  

**MGET command**  

MGET key [key ...]  

Returns the values of all specified keys. For every key that does not hold a string value or does not exist, the special value nil is returned. Because of this, the operation never fails.  

**DEL command**  

DEL key [key ...]  

Removes the specified keys. A key is ignored if it does not exist.  

**GETRANGE command**  

GETRANGE key start end

Returns the substring of the string value stored at key, determined by the offsets start and end (both are inclusive). Negative offsets can be used in order to provide an offset starting from the end of the string. So -1 means the last character, -2 the penultimate and so forth.  

**SETRANGE command**  

SETRANGE key offset value  

Overwrites part of the string stored at key, starting at the specified offset, for the entire length of value. If the offset is larger than the current length of the string at key, the string is padded with zero-bytes to make offset fit. Non-existing keys are considered as empty strings, so this command will make sure it holds a string large enough to be able to set value at offset.

Note that the maximum offset that you can set is 2^29 -1 (536870911), as Redis Strings are limited to 512 megabytes. If you need to grow beyond this size, you can use multiple keys.

so lets see an example:  
we can assign values of db to a small string or letter like alphabet letters.  
then save the values of traditional db in redis by a simple string like a car by id 1 and type "simple" and color "red" and size "big" become "srb".  
then we can use setrange and getrange and make the app very fast. for bulk use cases we can also use mset and mget.  
in fact this is just a naive example and this don't work in most cases. its better to use that conventional catching by filter values in redis.  

## Number commands

get set del mset and mget also works for numbers.  

**DECR command**  
**DECRBY command**  
**INCR command**  
**INCRBY command**  
**INCRBYFLOAT command**  

*why this commands exists?* Concurrent requests. is it enough??

**impl note** we need to write the cache layer inside a package with a single key generator method for each entity to avoid missing keys in read or write.  


# List  

## List Concept  

In Redis, a list is a collection of ordered elements where you can add or remove elements from both ends efficiently.  
Redis lists are implemented using linked lists, making them very fast for insertions and deletions.

## List Commands  

### Adding Elements to a List

1. LPUSH key value [value ...]	
	- Inserts one or more values at the head (left) of the list.  

2. RPUSH key value [value ...]	
	- Inserts one or more values at the tail (right) of the list.  

3. LPUSHX key value	 
	- Inserts a value at the head only if the list exists.  

4. RPUSHX key value	 
	- Inserts a value at the tail only if the list exists.  

### Retrieving Elements

5. LRANGE key start stop  
   	- Retrieves elements within a given range. Indexes are zero-based.  
  
6. LINDEX key index	 
	- Gets an element by index.  

7. LLEN key	 
	- Returns the length of the list.

### Removing Elements  

8. LPOP key	 
	- Removes and returns the first (leftmost) element.  

9. RPOP key	 
	- Removes and returns the last (rightmost) element.  

10. LREM key count value  
	- Trims the list, keeping only elements within the given range.  

### Modifying Elements  

11. LSET key index value  
	- Sets an element at the specified index.

### Blocking Operations (Useful for Queues)  

12. BLPOP key [key ...] timeout	 
	- Removes and returns the first element, blocking if the list is empty.  

13. BRPOP key [key ...] timeout	 
	- Removes and returns the last element, blocking if the list is empty.  

14. BRPOPLPUSH source destination timeout  
	- Pops from source and pushes to destination, blocking if needed.  

### Moving Elements Between Lists  

15. RPOPLPUSH source destination  
	- Removes last element from source and pushes it to the head of destination.  

*Summary*  

Redis lists are flexible and efficient for queue-like structures. Here are the key takeaways:

- LPUSH / RPUSH → Insert at start/end.  
- LPOP / RPOP → Remove from start/end.  
- LRANGE → Retrieve elements.  
- LLEN → Get list size.  
- LSET → Modify elements.  
- Blocking operations → Useful for queues.  

## Use Cases Of Redis Lists  

1. Message Queues - Using LPUSH and RPOP for FIFO queues.  
2. Recent Activity Feeds - Using LPUSH to store latest actions. (for example something like terminal auto suggestion)  
3. Task Processing - Using BRPOP to distribute tasks.  

# Hash  

## Hash Concept  

In Redis, a hash is a data structure that maps fields to values, similar to a dictionary or a key-value store within a key. It's useful for storing objects with multiple attributes.  

## Hash Commands
- HSET key field value  
(Sets a field in the hash.)
- HGET key field  
(Gets the value of a field.)
- HGETALL key  
(Gets all fields and values.)
- HMSET key field1 value1 field2 value2 ...  
(Sets multiple fields at once.)
- HDEL key field  
(Deletes a field.)
- HLEN key  
(Returns the number of fields in the hash.)
- HEXISTS key field  
(used to check if a specific field exists in a hash.)
- HINCRBY key field increment  
(Increments a field's integer value in a hash by a specified amount.)  
- HINCRBYFLOAT key field increment  
(Increments a field's floating-point value in a hash.)  
- HKEYS key  
(Returns all field names in the hash stored at key.)
- HVALS key  
(Returns all values in the hash stored at key.)  

## Use Cases
### use  
1. the record has many attributes
2. a collection of these records have to be stored many different ways  
3. often need to access a single record at a time

### don't use  
1. the record is only for counting or enforcing uniqueness
2. record stores only one or two attributes  
3. used only for creating relations between different records  
4. the record is only used for time series data  

## Summary

| Use Case                           | Redis Hashes? | Alternative                        |
| ---------------------------------- | ------------- | ---------------------------------- |
| User profiles, session data        | ✅ Yes         | -                                  |
| Frequent updates to fields         | ✅ Yes         | -                                  |
| Partial lookups (e.g., `age > 30`) | ❌ No          | `ZSET`, `RediSearch`               |
| Joins / Complex queries            | ❌ No          | SQL or NoSQL (MongoDB or ScyllaDB) |
| Large fields (10K+ fields)         | ❌ No          | JSON (`RedisJSON`)                 |
| Expiring individual fields         | ❌ No          | Store as `SET` with TTL            |

# Pipelines  
Redis Pipeline is a mechanism that allows sending multiple commands to the Redis server in a single network request, significantly improving performance by reducing the number of round-trip delays between the client and server.

---

## Concept
Normally, when you send a command to Redis, the sequence is:
1. The client sends a request to Redis.
2. Redis processes the request.
3. Redis sends back the response.
4. The client receives the response.

This process involves multiple round trips, which can be slow due to network latency.

With **pipelining**, you send multiple commands at once, reducing the number of network round trips, thereby improving efficiency.

### **Key Benefits of Using Pipelining:**
✅ Reduces latency by minimizing round trips  
✅ Improves throughput, especially for batch operations  
✅ Enhances performance when working with multiple keys  

---

## **2. How Redis Pipeline Works**
Instead of sending and waiting for each command sequentially, Redis pipelines batch commands together and execute them in one go.

### **Without Pipelining (Normal Execution)**
```bash
SET key1 value1 → Response: OK  
SET key2 value2 → Response: OK  
GET key1 → Response: value1  
GET key2 → Response: value2
```

Each command waits for the previous command’s response before being executed.

### **With Pipelining**
```bash
Pipeline: SET key1 value1 SET key2 value2 GET key1 GET key2

Response: OK OK value1 value2  
```
All commands are sent together, and responses are received in sequence.

---

## **3. Redis Pipeline Commands**
Most Redis commands can be used within a pipeline. Below are examples in **Golang** using the `go-redis` package.

### **Using Golang (`go-redis` package)**
```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pipe := client.Pipeline()

	pipe.Set(ctx, "key1", "value1", 0)
	pipe.Set(ctx, "key2", "value2", 0)
	getKey1 := pipe.Get(ctx, "key1")
	getKey2 := pipe.Get(ctx, "key2")

	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	value1, _ := getKey1.Result()
	value2, _ := getKey2.Result()

	fmt.Println("key1:", value1) // Output: key1: value1
	fmt.Println("key2:", value2) // Output: key2: value2
}
```

## **4. Redis Pipeline vs. Transaction**

| Feature             | Pipeline                                                    | Transaction (MULTI/EXEC)                                       |
| ------------------- | ----------------------------------------------------------- | -------------------------------------------------------------- |
| **Execution Order** | Not atomic (commands are sent but not executed as one unit) | Atomic (all commands execute as a single unit)                 |
| **Speed**           | Faster (batching reduces network latency)                   | Slower due to atomicity                                        |
| **Error Handling**  | Executes all commands, even if one fails                    | If a command fails, the entire transaction fails               |
| **Use Case**        | Bulk operations, performance optimization                   | Operations requiring atomicity, such as financial transactions |

## Example of Redis Transaction in Golang

```go
pipe := client.TxPipeline()

pipe.Set(ctx, "key1", "value1", 0)
pipe.Set(ctx, "key2", "value2", 0)
getKey1 := pipe.Get(ctx, "key1")
getKey2 := pipe.Get(ctx, "key2")

_, err := pipe.Exec(ctx)
if err != nil {
	log.Fatal(err)
}

value1, _ := getKey1.Result()
value2, _ := getKey2.Result()

fmt.Println("Transaction - key1:", value1)
fmt.Println("Transaction - key2:", value2)

```

## 5. When to Use Redis Pipelining?
✅ When you need to execute multiple Redis commands in bulk
✅ When reducing network latency is a priority
✅ When atomicity is not required

🚫 Avoid pipelining when:

- Commands depend on previous responses.
- You need transaction-like atomicity.

## 6. Common Mistakes & Best Practices
❌ Mistake: Sending Too Many Commands
Sending too many commands at once can overwhelm Redis. Break down large pipelines into smaller batches.

✅ Best Practice: Limit Pipeline Size
A good rule of thumb is to send batches of 100 to 1000 commands at a time.

❌ Mistake: Assuming Order is Maintained
Though Redis processes commands sequentially in a pipeline, external factors (such as network issues) can sometimes cause unexpected behavior.

✅ Best Practice: Monitor and Handle Errors
Use error handling and logging to catch failed pipeline executions.

## Conclusion
Redis Pipeline is a powerful feature for optimizing performance by reducing round-trip delays. While it is faster than sending individual commands, it lacks atomicity. For bulk operations, pipelining is ideal, but for transactional integrity, use MULTI/EXEC.  

# Set

a Set is an unordered collection of unique strings. Sets are highly efficient for operations like checking membership, intersections, unions, and differences between multiple sets.  

## Set Concept

- **Unique Elements**: A Redis set does not allow duplicate elements.
- **Unordered**: Elements inside a set do not maintain a specific order.
- **Fast Operations**: Redis provides O(1) time complexity for adding, removing, and checking membership in sets.
- **Supports Set Operations**: Redis supports union, intersection, and difference operations on sets.
- **Efficient for Membership Checks**: SISMEMBER allows quick lookups.

## Set Commands

- *SADD key member [member ...]*  
Adds multiple values to the set.  
Returns the number of elements successfully added.  

- *SCARD key*  
Gets the number of elements in the set.  
Returns the cardinality (number of elements) of the set.  

- *SDIFF key [key ...]*  
Finds the difference between multiple sets.  
Returns a set of elements present in the first set but not in the others.  

- *SDIFFSTORE destination key [key ...]*  
Stores the difference between multiple sets in a destination set.  
Returns the number of elements in the resulting set.  

- *SINTER key [key ...]*  
Finds the intersection of multiple sets.  
Returns a set of elements present in all specified sets.  

- *SINTERSTORE destination key [key ...]*  
Stores the intersection of multiple sets in a destination set.  
Returns the number of elements in the resulting set.  

- *SISMEMBER key member*  
Checks if a member exists in the set.  
Returns 1 if the member exists, 0 otherwise.  

- *SMEMBERS key*  
Retrieves all members of the set.  
Returns a list of all elements in the set.  

- *SMISMEMBER key member [member ...]*  
Checks if multiple members exist in the set.  
Returns a list of 1s and 0s indicating presence of each member.  

- *SMOVE source destination member*  
Moves a member from one set to another.  
Returns 1 if the member was moved, 0 otherwise.  

- *SPOP key [count]*  
Removes and returns one or more random members from the set.  
Returns the removed member(s).  

- *SRANDMEMBER key [count]*  
Gets one or more random members from the set without removing them.  
Returns the randomly selected member(s).  

- *SREM key member [member ...]*  
Removes one or more members from the set.  
Returns the number of members removed.  

- *SUNION key [key ...]*  
Finds the union of multiple sets.  
Returns a set of all unique elements from the specified sets.  

- *SUNIONSTORE destination key [key ...]*  
Stores the union of multiple sets in a destination set.  
Returns the number of elements in the resulting set.  

# Sorted Set  

## Sorted Set Concept  

A Sorted Set in Redis is a collection of unique elements, each associated with a floating-point score. Elements are ordered by their scores in ascending order. It is useful for ranking systems, leaderboards, priority queues, etc.

The key identifies the sorted set.  
Each member in the set is unique, but scores can be duplicated.  
The set is always maintained in sorted order based on scores.  

## Sorted Set Commands  

### Adding & Updating Elements

1. ZADD key score member [score member ...]  
	- Adds elements to a sorted set with an associated score. If the element exists, the score is updated.  
	- The number of new elements added (not counting updates).  

2. ZINCRBY key increment member  
	- Increases the score of a member by a given amount.  
	- The new score of the member.  

### Retrieving Elements

3. ZRANGE key start stop [BYSCORE | BYLEX] [REV] [LIMIT offset count]
  [WITHSCORES]  
  **[REV]** -> reverse order in return  
  **[WITHSCORES]** -> return key and score (not only key)  
  **[LIMIT offset count]** -> The optional LIMIT argument can be used to obtain a sub-range from the matching elements (similar to SELECT LIMIT offset, count in SQL). A negative 'count' returns all elements from the 'offset'. Keep in mind that if 'offset' is large, the sorted set needs to be traversed for 'offset' elements before getting to the elements to return, which can add up to O(N) time complexity.  
  [**BYSCORE**  
   | 
   **BYLEX** ] -> Lexicographical order in return  
  
	- Returns members in increasing score order within the given range (zero-based index).  
	- A list of members (and optionally their scores if WITHSCORES is provided).  
	- Starting with Redis 6.2.0, this command can replace the following commands: ZREVRANGE, ZRANGEBYSCORE, ZREVRANGEBYSCORE, ZRANGEBYLEX and ZREVRANGEBYLEX. (these commands are now deprecated.)  
  
4. ZREVRANGE key start stop [WITHSCORES] (deprecated)  
   - Returns members in decreasing score order within the given range.  
   - A list of members (and optionally their scores if WITHSCORES is provided).  
  
5. ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count] (deprecated)  
	- Retrieves members whose scores are within the given range (min to max).  
	- A list of members (and optionally their scores).  
  
6. ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count] (deprecated)  
   - Retrieves members within a reversed score range (max to min).  
   - A list of members (and optionally their scores).  
  
7. ZRANK key member [WITHSCORE]  
	- Returns the rank (position) of a member in ascending order.  
	- The rank (0-based index) or nil if the member is not found.  

8. ZREVRANK key member [WITHSCORE]  
	- Returns the rank of member in the sorted set stored at key, with the scores ordered from high to low. The rank (or index) is 0-based, which means that the member with the highest score has rank 0. Use ZRANK to get the rank of an element with the scores ordered from low to high.  
	- The rank (0-based index) or nil if the member is not found.  

9. ZSCORE key member  
	- Returns the score of member in the sorted set at key.  
	- If member does not exist in the sorted set, or key does not exist, nil is returned.  

### Counting & Range Queries

10. ZCARD key  
	- Returns the sorted set cardinality (number of elements) of the sorted set stored at key.  

11. ZCOUNT key min max  
	- Returns the number of elements in the sorted set at key with a score between min and max.  
	- O(log(N)) with N being the number of elements in the sorted set.  

12. ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count] (deprecated)  
	- O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it O(log(N)).  
	- Returns all the elements in the sorted set at key with a score between min and max (including elements with score equal to min or max). The elements are considered to be ordered from low to high scores.

### Removing Elements

13. ZREM key member [member ...]  
	- O(M*log(N)) with N being the number of elements in the sorted set and M the number of elements to be removed.  
	- returns the number of members removed.  

14. ZREMRANGEBYRANK key start stop  
	- O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements removed by the operation.
	- Removes all elements in the sorted set stored at key with rank between start and stop. Both start and stop are 0 -based indexes with 0 being the element with the lowest score. These indexes can be negative numbers, where they indicate offsets starting at the element with the highest score.  

15. ZREMRANGEBYSCORE key min max  
	- O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements removed by the operation.  
	- Removes all elements in the sorted set stored at key with a score between min and max (inclusive).  

### Lexicographical Operations (String Sorting)

16. ZRANGEBYLEX key min max [LIMIT offset count] (deprecated)  
	- O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it O(log(N)).  
	- When all the elements in a sorted set are inserted with the same score, in order to force lexicographical ordering, this command returns all the elements in the sorted set at key with a value between min and max.  

17. ZREVRANGEBYLEX key max min [LIMIT offset count] (deprecated)  
	- O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it O(log(N)).  
	- When all the elements in a sorted set are inserted with the same score, in order to force lexicographical ordering, this command returns all the elements in the sorted set at key with a value between max and min.  

18. ZLEXCOUNT key min max  
	- O(log(N)) with N being the number of elements in the sorted set.  
	- Counts the number of elements between min and max (by lexicographical order).  

19. ZREMRANGEBYLEX key min max
	- O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements removed by the operation.
	- When all the elements in a sorted set are inserted with the same score, in order to force lexicographical ordering, this command removes all elements in the sorted set stored at key between the lexicographical range specified by min and max.

### Intersection & Union Operations

20. ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE <SUM | MIN | MAX>]
	- O(N*K)+O(M*log(M)) worst case with N being the smallest input sorted set, K being the number of input sorted sets and M being the number of elements in the resulting sorted set.  
	- Finds the intersection of multiple sorted sets.  

21. ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE <SUM | MIN | MAX>] [WITHSCORES]
	- Computes the union of multiple sorted sets (newer version of ZUNIONSTORE).  
	- For a description of the WEIGHTS and AGGREGATE options, see ZUNIONSTORE.  

22. ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE <SUM | MIN | MAX>]  
	- Merges multiple sorted sets into one (Union).  
	- - For a description of the WEIGHTS and AGGREGATE options, see ZUNIONSTORE doc on web.  

23. ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]  
	- Computes the intersection of multiple sorted sets (newer version of ZINTERSTORE).  

### Miscellaneous  
24. ZPOPMIN key [count]  
	- Removes and returns up to count members with the lowest scores in the sorted set stored at key.  
	- O(log(N)*M) with N being the number of elements in the sorted set, and M being the number of elements popped.  

25. BZPOPMIN key [key ...] timeout  
	- BZPOPMIN is the blocking variant of the sorted set ZPOPMIN primitive.  

26. ZPOPMAX key [count]  
	- Removes and returns up to count members with the highest scores in the sorted set stored at key.
	- O(log(N)*M) with N being the number of elements in the sorted set, and M being the number of elements popped.

27. BZPOPMAX key [key ...] timeout  
	- BZPOPMAX is the blocking variant of the sorted set ZPOPMAX primitive.  

*summary:* 
- ZADD is used to insert/update elements.  
- ZRANGE and ZREVRANGE retrieve elements.  
- ZSCORE, ZRANK return ranking-related data.  
- ZUNIONSTORE, ZINTERSTORE perform set operations.  
- ZPOPMIN, ZPOPMAX remove elements efficiently.  


# Bitmaps  

## Concept Of Redis Bitmaps  

Bitmaps in Redis are a special type of data structure that allows manipulation of individual bits within a string value. Bitmaps are useful for efficiently storing and querying boolean (true/false) states, such as tracking user activity, feature flags, or binary representations.

*Key Characteristics of Bitmaps*
- Bit-Level Manipulation: Unlike standard strings, bitmaps operate on individual bits.
- Memory Efficient: Since a bit only requires 1/8th of a byte, a single Redis string (up to 512MB) can store billions of bits.
- Fast Bitwise Operations: Useful for analytics and tracking.  

## Bitmap Commands  

### Setting and Getting Bits  

1. SETBIT key offset value  
	- Set a specific bit in a string
	- key: The Redis key that stores the bitmap.  
	- offset: The position of the bit.  
	- value: Either 0 (false) or 1 (true).  

2. GETBIT key offset  
	- Get the value of a specific bit  

### Counting and Analyzing Bits  

3. BITCOUNT key [start end]  
	- Count the number of bits set to 1  
	- If no range is provided, it counts all bits

4. BITPOS key bit [start end]  
	- Find the first bit with a given value  
	- Finds the position of the first occurrence of bit (0 or 1)
	- Optional start and end parameters define the range  

### Performing Bitwise Operations  

5. BITOP operation destkey key1 [key2 ...]  
	- Perform bitwise operations on multiple keys  
	- operation: AND, OR, XOR, or NOT  
	- destkey: The key where the result will be stored  

### Using Bitfields for More Control  

6. BITFIELD key [subcommand ...]  
    - Perform multiple bitwise operations in one command  
    - Supports operations like:
     	- INCRBY: Increment bits as if they were integers.
     	- SET: Set a specific bit range.
     	- GET: Retrieve bits as an integer.

## Use Cases Of Redis Bitmaps  

1. User Activity Tracking: Track whether a user was active on a given day.  
2. Feature Flags: Enable or disable features for specific users.  
3. Real-Time Analytics: Count the number of active users in a time range.  
4. Bloom Filters & Probabilistic Data Structures: Used in combination with other techniques.  

## Limitations Of Redis Bitmaps  
- No direct deletion of bits: You can only set or reset them.  
- Fixed-size memory: Since bitmaps are stored as strings, they have a 512MB limit per key.  
- No automatic compression: Unlike specialized bit-oriented databases, Redis does not compress bitmaps.  

## Summary  

Redis Bitmaps are a powerful tool for efficient storage and manipulation of large-scale boolean data. With commands like SETBIT, GETBIT, BITCOUNT, and BITOP, you can build complex data analytics and tracking systems with minimal memory overhead.  

# HyperLogLog  

## Concept Of HyperLogLog  

HyperLogLog is a probabilistic data structure in Redis used to estimate the cardinality (the number of unique elements) of a dataset with high accuracy and low memory usage. It is particularly useful for counting distinct values in large datasets while using only 12 KB of memory (regardless of dataset size).  

Unlike traditional methods that require storing all unique elements, HyperLogLog approximates cardinality using hashing techniques, making it highly memory efficient.  

*Key Features of HyperLogLog in Redis*  
- **Memory Efficient**: Uses only 12 KB of memory to store cardinality estimates.
- **Approximate Counting**: It does not return exact results but provides an estimated count with an error rate of ~0.81%.
- **Scalable**: Works well with big data scenarios where tracking unique values is impractical with standard data structures like sets.  

## Commands Of HyperLogLog  

1. PFADD key [element [element ...]] 
	- Adds elements to a HyperLogLog data structure.
	- O(1) to add every element

2. PFCOUNT key [key ...]  
	- Estimates the number of unique elements in a HyperLogLog (cardinality)  
	- O(1) with a very small average constant time when called with a single key. O(N) with N being the number of keys, and much bigger constant times, when called with multiple keys.  

3. PFMERGE destkey [sourcekey [sourcekey ...]]  
	- O(N) to merge N HyperLogLogs, but with high constant times  
	- Merge multiple HyperLogLog values into a unique value that will approximate the cardinality of the union of the observed Sets of the source HyperLogLog structures.  

## Use Cases Of HyperLogLog  
- Counting Unique Visitors on a Website  
Track how many unique users visit a website per day without storing every user ID.
- Counting Unique IP Addresses in Network Logs  
Track the number of unique IP addresses over a period.
- Counting Unique Hashtags or Search Queries  
Analyze social media hashtags or search queries efficiently.

## Limitations Of HyperLogLog
- **Approximate Results**: Not suitable where exact counting is required.
- **Not for Data Retrieval**: You can count unique elements but cannot retrieve the stored values.  

## Summary  
HyperLogLog is a powerful and memory-efficient way to estimate the cardinality of large datasets in Redis. By using only 12 KB of memory, it allows for efficient counting without maintaining the entire dataset, making it perfect for big data applications.  

# Bitfields

## Concept Of Bitfields  

Redis Bitfields are an extension of Bitmaps that allow you to store and manipulate integers of arbitrary sizes within a single Redis string. This is useful when you want to manage compact data structures like counters or flags without needing multiple keys.  

**What Are Redis Bitfields?**  

Unlike normal Bitmaps, where each bit is treated separately, Bitfields let you treat groups of bits as individual integers. You can store multiple small integers within a single string key and perform atomic operations on them.  
For example, in a 32-bit Redis string, you can store:
- A 4-bit number at offset 0
- An 8-bit number at offset 4
- A 16-bit number at offset 12
This allows efficient storage and atomic updates. 

## Commands Of Bitfields  

1. BITFIELD key [GET encoding offset | [OVERFLOW <WRAP | SAT | FAIL>]
  <SET encoding offset value | INCRBY encoding offset increment>
  [GET encoding offset | [OVERFLOW <WRAP | SAT | FAIL>]
  <SET encoding offset value | INCRBY encoding offset increment>
  ...]]  

- GET – Reads a specific bitfield value.
- SET – Writes a value to a specific bitfield.
- INCRBY – Increments a bitfield value by a given amount.
- OVERFLOW – Defines how values behave when they exceed their bit width (WRAP, SAT, FAIL).  

*A. Storing and Retrieving Values*
#### Storing a Value (`SET`)
```sh
BITFIELD user_scores SET u8 0 100
```
- Stores the **number 100** as an **8-bit unsigned integer (`u8`)** at offset `0`.

#### Retrieving a Value (`GET`)
```sh
BITFIELD user_scores GET u8 0
```
- Reads the **8-bit unsigned integer** stored at offset `0` (returns `100`).

---

*B. Incrementing a Value (`INCRBY`)*
You can **increment** a bitfield value atomically:

```sh
BITFIELD user_scores INCRBY u8 0 10
```
- Increases the **8-bit** value at offset `0` by **10**.

---

*C. Handling Overflow (`OVERFLOW`)*
When a bitfield reaches its maximum possible value, you can define **overflow behavior**:

1. **WRAP** – Loops back to 0 when exceeding the limit (like modular arithmetic).
2. **SAT (Saturate)** – Stays at the max/min value when exceeding the limit.
3. **FAIL** – Returns an error if an overflow occurs.

Example:
```sh
BITFIELD user_scores OVERFLOW SAT INCRBY u8 0 200
```
- If the **u8 (8-bit)** number reaches **255**, it **stays at 255** instead of rolling over.

---

2. BITFIELD_RO key [GET encoding offset [GET encoding offset ...]]
- Read-only variant of the BITFIELD command. It is like the original BITFIELD but only accepts GET subcommand and can safely be used in read-only replicas.  

## Advantages of Bitfields
✅ **Memory Efficient** – Store multiple integers in a single key using bits.  
✅ **Atomic Operations** – Multiple bitfield operations in a **single Redis command**.  
✅ **Flexible Overflow Handling** – Prevent data corruption with controlled overflow.  
✅ **Ideal for Compact Counters** – Great for tracking user activity, game scores, etc.  

## Limitations of Bitfields
❌ **No Expiration on Individual Fields** – Redis expiration applies to the entire key, not parts of it.  
❌ **Limited Readability** – Hard to debug directly since data is stored as bits.  
❌ **Offset Management** – You must manually track bit positions to avoid overlap.  

# Geospatial Indexes  

## Concept Of Geospatial Indexes

geospatial indexes allow you to store, query, and manipulate location-based (latitude/longitude) data efficiently. This is done using the GEO commands, which leverage a sorted set (ZSET) under the hood.

How Geospatial Indexes Work in Redis
- **Storage**: Redis stores geospatial data using a sorted set, where locations are indexed by a unique key, and their positions are encoded into a geohash.
- **Encoding**: Redis converts latitude/longitude coordinates into a geohash, which is a single number that represents the location with a certain level of precision.
- **Efficiency**: Queries are fast because Redis uses geohash-based indexing to perform spatial lookups.  

## Commands Of Geospatial Indexes

1. GEOADD key [NX | XX] [CH] longitude latitude member [longitude
  latitude member ...]  

- Adds the specified geospatial items (longitude, latitude, name) to the specified key. Data is stored into the key as a sorted set, in a way that makes it possible to query the items with the GEOSEARCH command.  
- (log(N)) for each item added, where N is the number of elements in the sorted set.  

2. GEOPOS key [member [member ...]]  
- O(1)
- Return the positions (longitude,latitude) of all the specified members of the geospatial index represented by the sorted set at key  

3. GEODIST key member1 member2 [M | KM | FT | MI]  
- O(1)  
- Return the distance between two members in the geospatial index represented by the sorted set.  

4. GEORADIUS key longitude latitude radius <M | KM | FT | MI>
  [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC | DESC]
  [STORE key | STOREDIST key]
- deprecated
- O(N+log(M)) where N is the number of elements inside the bounding box of the circular area delimited by center and radius and M is the number of items inside the index.
- Return the members of a sorted set populated with geospatial information using GEOADD, which are within the borders of the area specified with the center location and the maximum distance from the center (the radius).  

5. GEOHASH key [member [member ...]]  
- O(1)  
- Return valid Geohash strings representing the position of one or more elements in a sorted set value representing a geospatial index (where elements were added using GEOADD).  

## Use Cases of Geospatial Indexing  
- Finding nearby stores, restaurants, or services (e.g., "show me all gas stations within 10 km").
- Ride-sharing and delivery apps (e.g., "find the closest available driver").
- Geofencing (e.g., "notify users when they enter a specific area").
- Logistics and supply chain tracking.