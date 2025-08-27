# What is Paxos?

**Paxos** is a consensus protocol designed to help multiple distributed nodes agree on a single value or decision, even in unreliable networks with potential failures.

- Originally invented for **distributed decision-making**.
- Not specifically designed for replication but adapted for database replication systems like Scylla.
- Ensures consistency and correctness despite node crashes or message losses.

---

## How Does Paxos Work?

Paxos works by running multiple rounds of communication among replicas (nodes) to **propose, agree on, and commit** a value (e.g., a database mutation).

In the context of Scylla’s Lightweight Transactions (LWT), the protocol consists of **four communication rounds (or phases)**:

---

## The 4 Round-Trip Mechanism in Paxos

### 1. **Prepare (Locking / Ballot Phase)**

- The coordinator (acting for the client) sends a **Prepare request** with a unique ballot number to all replicas.
- This asks: “Can you promise not to accept any proposals with a ballot number lower than this?”
- Each replica replies with a **promise** to ignore lower ballots and also returns any previously accepted value with the highest ballot number.
- This phase **locks** the replicas so no conflicting proposals can be accepted, preventing divergence.

### 2. **Propose (Value Proposal Phase)**

- Once the coordinator receives promises from a **majority** of replicas, it picks a value to propose:
  - If any replica replied with a previously accepted value, that value is chosen.
  - Otherwise, it uses the client’s proposed value.
- The coordinator sends a **Propose request** with the chosen value and ballot number.
- Replicas accept this value unless they have promised a higher ballot number.

### 3. **Accept (Commitment Phase)**

- When the coordinator receives acceptance from a majority, it knows the value is chosen (consensus is reached).
- However, this information is **not yet learned by all replicas**.

### 4. **Learn (Finalization Phase)**

- The coordinator sends a **Learn request** to replicas to commit/store the decided value permanently.
- Replicas write the mutation to their storage.
- The transaction is now committed and durable.

---

## Issues with Paxos and Remedies

### 1. **Performance Overhead**

- Multiple rounds increase latency compared to single-leader protocols.
- Remedy: Use Paxos **only for operations that require strict consistency**.
- Ongoing work aims to **collapse rounds** to reduce communication (e.g., merging retrieval and condition checks).

### 2. **Starvation**

- Some transactions may repeatedly lose out due to contention or retries.
- Remedy: Use **Paxos leases**, a mechanism to reduce starvation by temporarily granting leadership or priority.

### 3. **Uncertainty & Timeouts**

- A client may get a timeout error even though the mutation was actually committed (due to network delays or crashes).
- Remedy: Provide **better diagnostics** to help applications detect and handle these “uncertain” outcomes properly.

### 4. **State Persistence Overhead**

- Paxos requires storing intermediate state (ballots, promises) persistently.
- This consumes storage and needs capacity planning.
- Remedy: Introduce system tables with TTL to expire old state and **proactively clean up** successfully applied transactions.

---

## What is a "Promise" in Paxos?

- During the **Prepare phase** (the first phase of Paxos), the coordinator sends a **Prepare request** with a unique ballot number to all replicas.
- Each replica responds with a **promise**:
  - It promises **not to accept any proposals with ballot numbers lower than the one in the Prepare request**.
  - It also returns any previously accepted value with the highest ballot number it knows.
- This **promise** is a crucial step that locks the replicas from accepting conflicting proposals with lower ballot numbers, helping prevent inconsistencies.

### Summary Table

| Paxos Phase      | Message from Replica to Coordinator          | Called           |
|------------------|----------------------------------------------|------------------|
| Prepare Phase    | Replica replies that it will not accept lower ballots | **Promise**       |

The **"promise"** is the replica’s guarantee to the coordinator in response to the Prepare request.

---

If the coordinator receives promises from a majority of replicas, it proceeds to propose a value in the next phase.
