# Summary: Ballot Numbers and Conflict Handling in Paxos

## How is the Unique Ballot Number Generated?

- The ballot number (also called proposal number) is generated to be **monotonically increasing** and **globally unique**.
- It typically combines:
  - A local **counter** that increments each time a new ballot is created.
  - A unique **coordinator ID** (such as node ID or timestamp).
- This ensures that ballot numbers never conflict and always increase over time.

## Does the Ballot Number Always Increase and Avoid Conflicts?

- Yes, because of the combined counter and unique ID, ballot numbers are strictly increasing.
- No two coordinators generate the same ballot number.
- This ordering prevents conflicts and ensures replicas can choose the highest ballot number proposal safely.

## Can Two or More Nodes Start Sending Prepare Requests Simultaneously?

- Yes, multiple coordinators can start Prepare requests concurrently.
- Each coordinator sends a Prepare with its unique ballot number.
- Replicas respond to the **highest ballot number** they see, promising to reject lower ballot numbers.
- This can cause some proposals to be rejected initially but ensures eventual consensus.

## How Does Paxos Handle Concurrent Proposals?

- Paxos uses ballot numbers to **resolve conflicts** between concurrent proposals.
- Replicas only accept proposals with the highest ballot number they've promised.
- Coordinators with lower ballot numbers must retry with higher ballot numbers.
- This mechanism ensures safety and liveness even with multiple active coordinators.

---

This process guarantees that Paxos reaches agreement without conflicting decisions, even when multiple nodes propose values at the same time.

## Summary: Handling Concurrent Prepare Requests and Promise Counting in Paxos

## Scenario: Concurrent Coordinators with Conflicting Ballots

Consider a system with 3 nodes:

- **Node 1** is a coordinator with ballot number **b1**.
- **Node 2** is a coordinator with ballot number **b2**, where **b2 > b1**.
- **Node 3** receives the Prepare requests in the order: first **b1**, then **b2**.

### What happens?

- Node 3 first promises to **b1** when it receives that request.
- Then, upon receiving **b2** (which is higher), Node 3 updates its promise to **b2**.
- Node 1 counts the promise from Node 3 for **b1**.
- Node 2 counts the promise from Node 3 for **b2**.
- Both coordinators end up with only **1 promise each**, which is **not a majority**.
- Consequently, **neither coordinator can proceed to consensus**.

---

## Why does this happen?

- Paxos allows multiple coordinators to propose concurrently, which can lead to contention.
- This contention can cause a **live-lock** where no coordinator obtains a majority of promises.
- This situation is normal and expected in Paxos under concurrent proposals.

---

## How Paxos handles this problem

1. **Retries with Higher Ballots:**
   - Coordinators that fail to get a majority generate higher ballot numbers and retry.
   - Eventually, one coordinator will get a majority and progress.

2. **Backoff or Leader Election:**
   - Implementations use backoff timers or leader election to reduce contention.
   - This helps avoid continuous collisions.

3. **Paxos Leases:**
   - Paxos leases grant temporary leadership to one coordinator.
   - This reduces starvation and contention among coordinators.

---

## Summary

- Replicas respond to Prepare requests immediately as they arrive, promising the highest ballot number seen.
- Coordinators count promises matching their ballot number only.
- Concurrent proposals can cause situations where no coordinator gets a majority.
- Paxos ensures **safety** and **eventual progress** through retries, backoff, and leader election mechanisms.
