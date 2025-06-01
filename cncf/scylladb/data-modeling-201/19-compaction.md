# Compaction in ScyllaDB (Simple Explanation)

## 💡 What is Compaction?

**Compaction** in ScyllaDB is the process of merging multiple SSTable files on disk into fewer, cleaner ones.  
It helps improve read performance and reduce disk usage.

---

## Why is Compaction Needed?

When you write data in ScyllaDB:

1. Data is first stored in memory (**MemTable**).
2. It's also written to a **CommitLog** (for durability).
3. After the MemTable fills up, it's flushed to disk as an **SSTable**.
4. Many SSTables accumulate over time.

Too many SSTables cause:

- Slower reads (data might be spread across many files)
- More disk usage
- Duplicate or deleted data (tombstones) hanging around

---

## What Does Compaction Do?

- Reads multiple SSTables
- Merges rows with the same partition key
- Removes deleted data (tombstones)
- Creates new optimized SSTables
- Deletes the old ones

---

## Benefits of Compaction

- Faster reads (less disk seek)
- Reduced disk space usage
- Cleanup of deleted/expired data
- Better performance under load

---

## 🔧 Compaction Strategies in ScyllaDB

1. **Size-Tiered Compaction (STCS)**  
   - Default
   - Merges similarly sized SSTables

2. **Leveled Compaction (LCS)**  
   - Good for read-heavy workloads
   - Organizes SSTables into levels

3. **Time-Window Compaction (TWCS)**  
   - Best for time-series or log data
   - Merges SSTables based on time windows

---

## Are Data Always Stored in SSTables?

**Yes** — eventually, all data ends up in **SSTables** on disk.  
These files are immutable and only updated through **Compaction**.

---

## Summary

> **Compaction** is ScyllaDB’s way of keeping disk data clean, fast, and efficient.  
> It merges SSTables, removes garbage, and ensures fast read performance.
