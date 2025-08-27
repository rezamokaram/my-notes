# Cockoo Filter  

## Concept
The Cuckoo filter is a data structure that extends the Cuckoo filter to support approximate membership queries while also allowing for counting (i.e., tracking the number of occurrences of an item).

It is useful in applications where you need both space efficiency and fast membership testing while keeping track of how often an item appears, such as in networking (e.g., monitoring traffic flows), databases, and security.  

## Advantages of CuCo Filter
✅ Space-efficient – Uses less memory than a full hash table with counters.  
✅ Fast operations – Insertions, deletions, and lookups are done in constant time (O(1)).  
✅ Allows counting – Unlike Bloom filters or regular Cuckoo filters, it supports counting how many times an item appears.  

## Disadvantages
❌ False positives – Like Bloom and Cuckoo filters, it may falsely indicate an item is present.  
❌ Limited counter range – Counters are usually small integers, so they can overflow.  
❌ More complex than a basic Cuckoo filter – The added counting mechanism increases the implementation complexity.  

## How Cuckoo Filter Works  

1. Uses Cuckoo Hashing 🐦  

Each item is hashed into a small fingerprint (a short representation of the item).
The item is then stored in one of two possible bucket locations in a hash table.
If the primary location is full, an existing item is kicked out (relocated to its alternate bucket), following Cuckoo Hashing rules.  

2. Insertion Process  

Compute a fingerprint from the item.
Compute two possible bucket locations for the fingerprint.
Place the fingerprint in an empty spot in either bucket.
If both buckets are full, evict an existing fingerprint and move it to its alternate bucket.
Repeat until insertion succeeds or a maximum number of relocations is reached (then resize or fail).  

3. Lookup Process  

Compute the fingerprint and its two possible bucket locations.
Check if the fingerprint exists in either bucket.
If found, return "probably present"; otherwise, return "definitely absent".  

4. Deletion Process

Since fingerprints are explicitly stored, they can be removed easily (unlike Bloom filters).
Just locate the fingerprint and remove it.

## *Cuckoo vs Bloom*  


| Feature             | Cuckoo Filter                        | Bloom Filter                      |
|---------------------|------------------------------------|----------------------------------|
| **Memory Usage**    | Less than Bloom (in many cases)   | Higher compared to Cuckoo        |
| **Lookup Time**     | Fast (`O(1)`)                      | Fast (`O(k)`)                    |
| **Insertion**       | Fast (`O(1)`) but may require eviction | Fast (`O(k)`)                   |
| **Deletion Support**| ✅ Supported                        | ❌ Not Supported                 |
| **False Positive Rate** | Lower than Bloom in most cases  | Slightly higher (depends on `k` and bit size) |
| **Data Structure**  | Array of buckets with fingerprints | Bit array                        |
| **Collision Handling** | Eviction via Cuckoo Hashing      | Multiple hash functions (`k` hashes) |
| **Use Case**        | Suitable for deletable & dynamic data | Best for read-only filters       |


## Positions In Cuckoo filter  

the two positions (pos1 and pos2) in a Cuckoo Filter follow a specific formula derived from Cuckoo Hashing. These positions are chosen strategically to reduce collisions, improve insertions, and enable efficient evictions when needed.

### 🔹 Formula for Computing Two Positions in Cuckoo Filter

After hashing an item, we compute two possible positions for storing its **fingerprint**:

1️⃣ **First Position (`pos1`):**
```plaintext
pos1 = hash(item) % tableSize  
```

This is a simple modulo operation that maps the item’s hash to a valid bucket in the table.

2️⃣ Second Position (pos2):  
```plaintext
pos2 = pos1 ⊕ hash(fingerprint)
```  
Here, XOR (⊕) is used between pos1 and the hash of the item's fingerprint.

### 🔹 Why These Two Positions?

1️⃣ Collision Reduction  
In a traditional hash table, an item has only one position, leading to collisions when multiple items hash to the same slot.  
Cuckoo Filter, however, provides two potential positions, reducing the likelihood of conflicts.  
  
2️⃣ Higher Insertion Success  
If the first position (pos1) is occupied, the item can still fit in pos2, increasing the probability of successful insertion.  

3️⃣ Efficient Eviction (Cuckoo Hashing)  
If both positions are occupied, an item can be evicted and moved to its alternate position.  
Using XOR ensures that after moving, the new position remains valid within the table.  

4️⃣ Better Data Distribution  
By using hash(fingerprint), the second position (pos2) is spread out more randomly across the table, improving data distribution and reducing clustering.  

🔹 Example Calculation
Suppose we have:

hash("Alice") = 137  
tableSize = 10  

**Step 1:** Compute pos1  
pos1 = 137 % 10 = 7  

**Step 2:** Compute Fingerprint  
Hash Assume hash(fingerprint) = 4  

**Step 3:** Compute pos2  
So, "Alice" can be stored in either bucket 7 or 3.  

### 🔹 Conclusion
✅ Using two positions in a Cuckoo Filter increases the chances of inserting items successfully.
✅ XOR-based calculation of pos2 ensures better data spread and reduces collisions.
✅ Eviction mechanism allows relocating items to alternative positions if needed, improving space utilization.  

## Sample Golang Impl  

```go

package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
)

// Parameters for the Cuckoo Filter
const (
	BucketSize = 4  // Number of entries per bucket
	NumBuckets = 10 // Number of buckets in the table
	MaxKicks   = 500 // Maximum number of evictions
)

// CuckooFilter represents a simple Cuckoo Filter
type CuckooFilter struct {
	table [NumBuckets][BucketSize]uint32
}

// hashFunc generates a hash value for a given key
func hashFunc(data string) uint32 {
	h := fnv.New32()
	h.Write([]byte(data))
	return h.Sum32()
}

// getPositions calculates the two candidate positions
func getPositions(item string) (uint32, uint32) {
	hashVal := hashFunc(item)
	pos1 := hashVal % NumBuckets
	fingerprint := hashVal & 0xFFFF // Take lower 16 bits as fingerprint
	pos2 := pos1 ^ (fingerprint % NumBuckets) // XOR with fingerprint hash
	return pos1, pos2
}

// Insert adds an item to the filter
func (cf *CuckooFilter) Insert(item string) bool {
	pos1, pos2 := getPositions(item)
	fingerprint := hashFunc(item) & 0xFFFF // Generate fingerprint

	// Try to insert into the first position
	for i := 0; i < BucketSize; i++ {
		if cf.table[pos1][i] == 0 {
			cf.table[pos1][i] = fingerprint
			return true
		}
	}

	// Try to insert into the second position
	for i := 0; i < BucketSize; i++ {
		if cf.table[pos2][i] == 0 {
			cf.table[pos2][i] = fingerprint
			return true
		}
	}

	// If both are full, perform eviction
	curPos := pos1
	for n := 0; n < MaxKicks; n++ {
		index := rand.Intn(BucketSize) // Randomly select a slot to evict
		evictedFingerprint := cf.table[curPos][index]
		cf.table[curPos][index] = fingerprint // Replace with new fingerprint

		// Compute new alternate position for the evicted fingerprint
		newPos := curPos ^ (evictedFingerprint % NumBuckets)

		// Try inserting the evicted fingerprint in the alternate position
		for i := 0; i < BucketSize; i++ {
			if cf.table[newPos][i] == 0 {
				cf.table[newPos][i] = evictedFingerprint
				return true
			}
		}
		curPos = newPos
	}
	return false // Insert failed after max kicks
}

// Lookup checks if an item is in the filter
func (cf *CuckooFilter) Lookup(item string) bool {
	pos1, pos2 := getPositions(item)
	fingerprint := hashFunc(item) & 0xFFFF

	for i := 0; i < BucketSize; i++ {
		if cf.table[pos1][i] == fingerprint || cf.table[pos2][i] == fingerprint {
			return true
		}
	}
	return false
}

// Delete removes an item from the filter
func (cf *CuckooFilter) Delete(item string) bool {
	pos1, pos2 := getPositions(item)
	fingerprint := hashFunc(item) & 0xFFFF

	for i := 0; i < BucketSize; i++ {
		if cf.table[pos1][i] == fingerprint {
			cf.table[pos1][i] = 0
			return true
		}
		if cf.table[pos2][i] == fingerprint {
			cf.table[pos2][i] = 0
			return true
		}
	}
	return false
}

// Main function to test the Cuckoo Filter
func main() {
	filter := CuckooFilter{}

	// Insert items
	fmt.Println("Inserting 'Alice':", filter.Insert("Alice"))
	fmt.Println("Inserting 'Bob':", filter.Insert("Bob"))
	fmt.Println("Inserting 'Charlie':", filter.Insert("Charlie"))

	// Lookup items
	fmt.Println("Lookup 'Alice':", filter.Lookup("Alice")) // Should be true
	fmt.Println("Lookup 'Bob':", filter.Lookup("Bob"))     // Should be true
	fmt.Println("Lookup 'Eve':", filter.Lookup("Eve"))     // Should be false

	// Delete items
	fmt.Println("Deleting 'Alice':", filter.Delete("Alice"))
	fmt.Println("Lookup 'Alice' after deletion:", filter.Lookup("Alice")) // Should be false
}


```

## Another Example For Focussing On Positions

```go

package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/rand"
)

const (
	bucketSize      = 2  // Number of slots in each bucket
	tableSize       = 10 // Total number of buckets
	fingerprintSize = 4  // Size of the fingerprint (in bytes)
	maxKicks        = 500 // Maximum number of evictions
)

// **Bucket structure**
type Bucket struct {
	slots []string
}

// **CuckooFilter structure**
type CuckooFilter struct {
	table []Bucket
}

// **Create a new Cuckoo Filter**
func NewCuckooFilter() *CuckooFilter {
	table := make([]Bucket, tableSize)
	for i := range table {
		table[i] = Bucket{slots: make([]string, 0, bucketSize)}
	}
	return &CuckooFilter{table: table}
}

// **Compute hash of an item**
func hash(item string) uint32 {
	hash := md5.Sum([]byte(item))
	return binary.LittleEndian.Uint32(hash[:])
}

// **Generate fingerprint from an item**
func fingerprint(item string) string {
	hash := md5.Sum([]byte(item))
	return fmt.Sprintf("%x", hash[:fingerprintSize])
}

// **Compute two possible positions for the item**
func (cf *CuckooFilter) getPositions(item string) (string, int, int) {
	h := hash(item)
	pos1 := int(h % tableSize)
	fp := fingerprint(item)
	pos2 := (pos1 ^ int(hash(fp))) % tableSize
	return fp, pos1, pos2
}

// **Insert an item into the filter**
func (cf *CuckooFilter) Insert(item string) bool {
	fp, pos1, pos2 := cf.getPositions(item)

	// Insert into the first bucket if there is space
	if len(cf.table[pos1].slots) < bucketSize {
		cf.table[pos1].slots = append(cf.table[pos1].slots, fp)
		return true
	}

	// Insert into the second bucket if there is space
	if len(cf.table[pos2].slots) < bucketSize {
		cf.table[pos2].slots = append(cf.table[pos2].slots, fp)
		return true
	}

	// If both buckets are full, perform eviction
	pos := []int{pos1, pos2}[rand.Intn(2)] // Randomly choose one of the two positions
	for i := 0; i < maxKicks; i++ {
		index := rand.Intn(bucketSize) // Select a random slot to evict
		fp, cf.table[pos].slots[index] = cf.table[pos].slots[index], fp
		pos = (pos ^ int(hash(fp))) % tableSize // Compute new position

		// Try inserting into the new position
		if len(cf.table[pos].slots) < bucketSize {
			cf.table[pos].slots = append(cf.table[pos].slots, fp)
			return true
		}
	}

	return false // Insertion failed (table needs expansion)
}

// **Lookup an item in the filter**
func (cf *CuckooFilter) Lookup(item string) bool {
	fp, pos1, pos2 := cf.getPositions(item)
	for _, slot := range cf.table[pos1].slots {
		if slot == fp {
			return true
		}
	}
	for _, slot := range cf.table[pos2].slots {
		if slot == fp {
			return true
		}
	}
	return false
}

// **Delete an item from the filter**
func (cf *CuckooFilter) Delete(item string) bool {
	fp, pos1, pos2 := cf.getPositions(item)

	// Search and remove from the first bucket
	for i, slot := range cf.table[pos1].slots {
		if slot == fp {
			cf.table[pos1].slots = append(cf.table[pos1].slots[:i], cf.table[pos1].slots[i+1:]...)
			return true
		}
	}

	// Search and remove from the second bucket
	for i, slot := range cf.table[pos2].slots {
		if slot == fp {
			cf.table[pos2].slots = append(cf.table[pos2].slots[:i], cf.table[pos2].slots[i+1:]...)
			return true
		}
	}

	return false // Item not found
}

// **Test the implementation**
func main() {
	cf := NewCuckooFilter()

	// Insert items
	cf.Insert("Alice")
	cf.Insert("Bob")
	cf.Insert("Charlie")

	// Lookup items
	fmt.Println("Is Alice present?", cf.Lookup("Alice")) // True
	fmt.Println("Is Eve present?", cf.Lookup("Eve"))     // False

	// Delete an item
	cf.Delete("Alice")
	fmt.Println("Is Alice present after deletion?", cf.Lookup("Alice")) // False
}


```