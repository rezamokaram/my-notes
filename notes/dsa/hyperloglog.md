# **How HyperLogLog Works**

## **Overview**
HyperLogLog is a **probabilistic algorithm** for estimating the number of **distinct elements (cardinality)** in a dataset using **very little memory**. Instead of storing each unique element, it relies on **hashing and counting leading zeros** to approximate the count.

---

## **1. The Key Idea: Leading Zeros in Hashes**
HyperLogLog is based on the statistical observation that **the number of leading zeros in a hashed value follows a predictable logarithmic pattern**:

- About **50%** of numbers start with `1` (zero leading zeros).
- About **25%** start with `01` (one leading zero).
- About **12.5%** start with `001` (two leading zeros).
- About **6.25%** start with `0001` (three leading zeros).
- And so on...

If we observe a hash with **10 leading zeros**, it suggests that the dataset is large enough that such a "rare" value appeared.  
Thus, if the **maximum leading zero count is \( k \)**, we estimate that the number of unique elements is approximately:

Estimated Count ≈ 2^k

---

## **2. Steps in HyperLogLog**
### **Step 1: Hash Each Input**
Each incoming element is hashed using a **uniform hash function** (e.g., MurmurHash, SHA-256).  
This ensures an even distribution.

Example:

| Element | Hash (Binary) |
|---------|--------------|
| `Alice`  | `0001011011010110...` |
| `Bob`    | `0000001100101110...` |
| `Charlie`| `0000000001000111...` |

### **Step 2: Count Leading Zeros**
For each hashed value, count the number of **leading zeros**.

| Element | Hash | Leading Zeros |
|---------|------------------|---------------|
| `Alice`  | `0001011011010110...` | 3 |
| `Bob`    | `0000001100101110...` | 6 |
| `Charlie`| `0000000001000111...` | 9 |

The **maximum leading zero count** here is **9**, suggesting approximately \( 2^9 = 512 \) unique elements.

### **Step 3: Use Multiple Registers for Accuracy**
Instead of using just one hash function, HLL **splits the hash into two parts**:
1. The **first few bits** determine the register (bucket) where data is stored.
2. The **remaining bits** are used to count leading zeros.

For example, with **16 registers**, the first **4 bits** of the hash define which register to update.  
Each register **stores the maximum leading-zero count it has seen**.

### **Step 4: Compute the Final Estimate**
To get the final estimate, we:
- **Apply the harmonic mean** across all registers.
- **Use a bias correction factor** to adjust the estimate.

The final formula is:

\[
E = \alpha_m \cdot m^2 \cdot \left(\sum_{i=1}^{m} 2^{-R_i} \right)^{-1}
\]

where:
- \( m \) = number of registers.
- \( R_i \) = max leading zeros per register.
- \( \alpha_m \) = correction factor to reduce bias.

---

## **3. Why HyperLogLog Uses Maximum Leading Zeros**
You might wonder: **Why use the maximum leading zeros instead of an average?**  

- The distribution of leading zeros is highly skewed.
- Most numbers have **small leading zero counts**, making their average misleading.
- The **largest leading zero count** represents the **"rarest event"**, which strongly correlates with the dataset size.

Thus, HLL **tracks the maximum per register** instead of averaging them.

---

## **4. Complexity & Memory Efficiency**
| Method | Memory Usage | Accuracy | Speed |
|--------|-------------|----------|------------|
| **Traditional Counting (Set/Database)** | High (stores all unique users) | 100% | Slow for large data |
| **HyperLogLog** | Very low (~1.5 KB for millions) | ~98% | Fast (\( O(1) \)) |

- **Time Complexity:** \( O(1) \) per update/query (constant time).
- **Space Complexity:** \( O(\log \log n) \), requiring only a few kilobytes.

---

## **5. Real-World Applications**
HyperLogLog is widely used in **big data systems**:
- **Web Analytics**: Counting **unique visitors** without storing every IP.
- **Databases**: PostgreSQL, Redis (`PFCOUNT`, `PFADD` commands).
- **Networking**: Counting **unique IP addresses** in traffic logs.
- **Social Media**: Counting unique hashtags or user interactions.

---

## **6. Example: Using HyperLogLog in Redis**
Redis provides built-in HyperLogLog commands:

1. **Add a visitor to HLL**:
```sh
PFADD unique_visitors 192.168.1.1
PFADD unique_visitors 192.168.1.2
PFADD unique_visitors 192.168.1.3
```

2. Get the estimated unique visitor count:  
```sh
PFCOUNT unique_visitors  
```
Output: ≈ 3 (approximate count).

---

## **7. Summary**
✅ Probabilistic: Uses hashing and leading-zero counting.  
✅ Memory Efficient: Needs only 𝑂(log log 𝑛) space (~1.5 KB for millions of elements).  
✅ Fast: Constant-time updates and queries.  
⚠️ Not Exact: Has a small margin of error (~1%-2%).  