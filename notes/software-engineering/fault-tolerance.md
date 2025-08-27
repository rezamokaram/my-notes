# Fault Tolerance

## What is Fault Tolerance?

**Fault tolerance** is the ability of a software system to **continue operating correctly even when parts of it fail**. The system might degrade in performance, but it should not crash or produce incorrect results.

> It’s about *resilience*, not perfection.

---

## Why Is It Important?

In modern software—especially distributed systems—failures are **inevitable**:

- Network timeouts
- Server crashes
- Disk or memory faults
- Bugs or malformed requests

Fault-tolerant systems **anticipate and isolate failures**, ensuring minimal disruption.

---

## 🔁 Common Techniques for Fault Tolerance

| Technique                 | Description |
|--------------------------|-------------|
| **Retries**              | Automatically reattempt failed operations (with backoff) |
| **Circuit Breakers**     | Prevent cascading failures by stopping calls to unstable services |
| **Fallbacks**            | Use a default or cached response if a service fails |
| **Redundancy**           | Duplicate components (e.g., servers, databases) to avoid single points of failure |
| **Graceful Degradation** | Reduce functionality (not crash) when a part of the system fails |
| **Replication**          | Copy data or services across multiple nodes |
| **Timeouts**             | Set limits on how long operations can run to avoid hanging systems |

---

## 🧰 Real-World Examples

### ✅ Example 1: API Retry with Fallback

```js
// JavaScript pseudo-code
try {
  const response = await fetch("https://payments.example.com");
  return response.json();
} catch (err) {
  console.log("Primary service failed. Using fallback.");
  return { status: "unavailable", cachedBalance: 100 };
}
