# ScyllaDB Vs Relational

## **Core Approach**

- **ScyllaDB**: Query-first (`Application → Data → Model`). Optimize for specific queries.  
- **Relational**: Entity-first (`Data → Model → Application`). Focus on data structure, then queries.  

In ScyllaDB, we think about the application before thinking about the data. Data modeling is query based. In relational databases, however, the queries are mostly an afterthought and the data model is based on the entities.

---

## **Key Differences**  

| **ScyllaDB**                | **Relational Databases**       |  
|-----------------------------|--------------------------------|  
| Query-based design          | Entity-based design            |  
| Denormalized data           | Normalized data                |  
| No joins/foreign keys       | Supports joins/foreign keys    |  
| Eventual consistency (CAP)  | Strong consistency (ACID)      |  
| Distributed (no single point of failure) | Centralized (single point of failure risk) |  

---

### **Why It Matters**  

- **ScyllaDB**: Prioritizes performance, scalability, and availability for distributed systems.  
- **Relational**: Ensures data integrity and complex relationships at the cost of scalability.  

**Short Takeaway**: ScyllaDB flips the script—design for queries first, not entities.
