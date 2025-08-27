# Bloom Filter  

A Bloom Filter is a probabilistic data structure that is used to test whether an element is present in a set. It is highly space-efficient but allows a small probability of false positives (i.e., it may say an element is present when it's not) while guaranteeing no false negatives (i.e., if it says an element is absent, it's definitely absent).

## How It Works

A Bloom Filter uses:

1. A bit array of size `m`, initially set to all zeros.  
1. `k` independent hash functions that each map elements to positions in the bit array.  

## Insertion

When inserting an element:

- The element is passed through each of the `k` hash functions.
- Each hash function gives a position in the bit array.
- The bits at these positions are set to 1.  

## Membership Check (Look Up)

To check if an element is in the set:

- The element is passed through the same `k` hash functions.  
- The corresponding bits in the array are checked.  
- If all bits are `1`, the element might be present.  
- If any bit is `0`, the element is definitely absent.  

## Key Properties  

- False Positives: An element not in the set may still return "present" if all its hash positions happen to be set by other elements.  
- No False Negatives: If a Bloom Filter says an element is absent, it's guaranteed to be absent.
- Space Efficiency: Much more compact than storing actual elements.  
- No Deletions (Typically): Removing elements isn't straightforward since bits may be shared by multiple elements.  

## Applications  

- Database query optimization (e.g., checking if data might exist before querying)  
- Web caching (e.g., checking if a URL has been visited)  
- Distributed systems (e.g., detecting duplicates in a large-scale system)  

## Another Bloom filter vs Cuckoo filter  

| Feature           | **Bloom Filter** | **Cuckoo Filter** |
|------------------|----------------|----------------|
| **Basic Idea**  | Uses a bit array and multiple hash functions. | Uses a cuckoo hash table with fingerprint storage. |
| **False Positives** | Yes, due to hash collisions. | Yes, but lower than Bloom filters for the same space usage. |
| **False Negatives** | No (Always says "absent" correctly). | No (unless the filter overflows). |
| **Insertions** | Simple: set bits at hashed positions. | More complex: insert fingerprint, may trigger relocation. |
| **Membership Check** | Check if all corresponding bits are set. | Check if fingerprint exists in one of two buckets. |
| **Deletions** | Difficult (no way to reset a single bit without affecting other elements). | Easy (removes specific fingerprint). |
| **Memory Efficiency** | More space-efficient for large sets. | More space-efficient when false positive rate is low. |
| **Adaptability** | Fixed size; resizing requires rebuilding. | Can adapt dynamically using cuckoo hashing techniques. |
| **Performance** | Fast lookups and inserts. | Lookups are fast, but inserts may require multiple relocations. |

## Sample Impl  

```go
package main

import (
 "crypto/sha256"
 "encoding/binary"
 "fmt"
 "hash/fnv"
 "math"
)

type BloomFilter struct {
 bitArray []bool
 size     uint
 hashFuncs int
}

// NewBloomFilter initializes a Bloom Filter with given size and number of hash functions
func NewBloomFilter(size uint, hashFuncs int) *BloomFilter {
 return &BloomFilter{
  bitArray: make([]bool, size),
  size:     size,
  hashFuncs: hashFuncs,
 }
}

// hash generates multiple hash values for an element
func (bf *BloomFilter) hash(data string, seed int) uint {
 h := fnv.New64a()
 h.Write([]byte(fmt.Sprintf("%d:%s", seed, data)))
 return uint(h.Sum64()) % bf.size
}

// Add inserts an element into the Bloom Filter
func (bf *BloomFilter) Add(data string) {
 for i := 0; i < bf.hashFuncs; i++ {
  index := bf.hash(data, i)
  bf.bitArray[index] = true
 }
}

// MightContain checks if an element might be in the filter
func (bf *BloomFilter) MightContain(data string) bool {
 for i := 0; i < bf.hashFuncs; i++ {
  index := bf.hash(data, i)
  if !bf.bitArray[index] {
   return false
  }
 }
 return true
}

func main() {
 bf := NewBloomFilter(1000, 3)
 bf.Add("hello")
 bf.Add("world")

 fmt.Println("MightContain 'hello':", bf.MightContain("hello")) // true
 fmt.Println("MightContain 'golang':", bf.MightContain("golang")) // false (most likely)
}
```

## Comparison of Implementations

| Feature         | **Bloom Filter** | **Cuckoo Filter** |
|----------------|----------------|----------------|
| **False Positives** | Yes | Yes (but lower) |
| **False Negatives** | No | No (unless the table overflows) |
| **Insert Speed** | Faster | Slower (due to possible relocations) |
| **Memory Usage** | Efficient | More efficient for low false positive rates |
| **Deletions** | Difficult | Easy |
| **Use Case** | Large-scale applications, caching | When deletions and lower false positives are needed |

## Which One Should You Use?

- Use Bloom Filter if:

  - You need fast insertions and lookups.  
  - You can tolerate false positives.  
  - You don't need deletions.  

- Use Cuckoo Filter if:

  - You need low false positives.  
  - You need deletions.  
  - You can handle a more complex insertion process.  
