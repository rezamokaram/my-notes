# Redis.io

## why to use redis ?  

- redis is fast  
    1. all data stored in memory
    2. data is organized in simple data structures  
    3. redis has a simple feature set

## Strings and Basic commands
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

## Hash  

In Redis, a hash is a data structure that maps fields to values, similar to a dictionary or a key-value store within a key. It's useful for storing objects with multiple attributes.  

### Some of Hash Commands
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