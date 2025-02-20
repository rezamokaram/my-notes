# Redis.io

1. [Introduction](#introduction)
2. [Strings and Basic commands](#strings-and-basic-commands)
3. [Hash](#hash)
4. [Pipelines](#pipelines)
	- [Concept](#concept)
5. [Set](#set)
	- [Concept](#concept)
	- [Commands](#commands)
6. [Sorted Set](#sorted-set)
	- [Concept](#concept)

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

# Hash  

In Redis, a hash is a data structure that maps fields to values, similar to a dictionary or a key-value store within a key. It's useful for storing objects with multiple attributes.  

## Commands
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

| Use Case                           | Redis Hashes? | Alternative            |
|-------------------------------------|--------------|------------------------|
| User profiles, session data        | ✅ Yes       | -                      |
| Frequent updates to fields         | ✅ Yes       | -                      |
| Partial lookups (e.g., `age > 30`) | ❌ No        | `ZSET`, `RediSearch`   |
| Joins / Complex queries            | ❌ No        | SQL or NoSQL (MongoDB or ScyllaDB) |
| Large fields (10K+ fields)         | ❌ No        | JSON (`RedisJSON`)     |
| Expiring individual fields         | ❌ No        | Store as `SET` with TTL |

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

| Feature         | Pipeline | Transaction (MULTI/EXEC) |
|---------------|----------|-------------------------|
| **Execution Order** | Not atomic (commands are sent but not executed as one unit) | Atomic (all commands execute as a single unit) |
| **Speed** | Faster (batching reduces network latency) | Slower due to atomicity |
| **Error Handling** | Executes all commands, even if one fails | If a command fails, the entire transaction fails |
| **Use Case** | Bulk operations, performance optimization | Operations requiring atomicity, such as financial transactions |

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

## Concept

- **Unique Elements**: A Redis set does not allow duplicate elements.
- **Unordered**: Elements inside a set do not maintain a specific order.
- **Fast Operations**: Redis provides O(1) time complexity for adding, removing, and checking membership in sets.
- **Supports Set Operations**: Redis supports union, intersection, and difference operations on sets.
- **Efficient for Membership Checks**: SISMEMBER allows quick lookups.

## Commands

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

## Concept  

A Sorted Set in Redis is a collection of unique elements, each associated with a floating-point score. Elements are ordered by their scores in ascending order. It is useful for ranking systems, leaderboards, priority queues, etc.

The key identifies the sorted set.  
Each member in the set is unique, but scores can be duplicated.  
The set is always maintained in sorted order based on scores.  

## Commands  

1. ZADD key score member [score member ...]
	- Adds elements to a sorted set with an associated score. If the element exists, the score is updated.
	- The number of new elements added (not counting updates).

2. ZINCRBY key increment member
	- Increases the score of a member by a given amount.  
	- The new score of the member.  

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

*summary:* 
- ZADD is used to insert/update elements.
- ZRANGE and ZREVRANGE retrieve elements.
- ZSCORE, ZRANK return ranking-related data.
- ZUNIONSTORE, ZINTERSTORE perform set operations.
- ZPOPMIN, ZPOPMAX remove elements efficiently.