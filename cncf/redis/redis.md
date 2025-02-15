# getting start

## why to use redis ?  

1. redis is fast  
    - all data stored in memory
    - data is organized in simple data structures  
    - redis has a simple feature set

## redis commands
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

2.6.