# ACID

1. **Atomicity:**
    - In the context of databases, **atomicity** means that a transaction is either fully committed or not committed at all.
    - When you perform a transaction (such as creating a new record or updating data), it is treated as an atomic unit.
    - If the transaction succeeds, all changes are applied together. If any part of the transaction fails (due to an error or interruption), the entire operation is rolled back to its initial state.
    - Essentially, atomicity ensures that partial changes are not left behind, maintaining data integrity.
2. **Consistency:**
    - **Consistency** ensures that a transaction brings the database from one valid state to another.
    - one of the most real time scenarios for this concept is to don’t have a f-key that the other entity does not exist
    - Before and after a transaction, certain properties must hold true to maintain consistency.
    - For example, if a bank transfer deducts funds from one account and adds them to another, the total balance across all accounts should remain consistent.
    - If a transaction violates consistency (e.g., due to a bug or system failure), it is rolled back to prevent data corruption.
3. **Isolation:**
    - **Isolation** ensures that concurrent transactions do not interfere with each other.
    - When multiple transactions run simultaneously, they should be isolated from one another.
    - Isolation prevents issues like dirty reads (reading uncommitted data), non-repeatable reads (reading different values within the same transaction), and phantom reads (seeing new rows inserted by other transactions).
    - Database systems use locks and isolation levels (such as READ COMMITTED or SERIALIZABLE) to manage concurrent access.
4. **Durability:**
    - **Durability** guarantees that once a transaction is committed, its changes are permanent and survive system failures (e.g., power outage, crash).
    - Committed data remains intact even if the system restarts or encounters unexpected issues.
    - To achieve durability, databases use transaction logs and write-ahead logs (WALs) to record changes before applying them to the data files.
    - This property ensures that data remains reliable over time.
