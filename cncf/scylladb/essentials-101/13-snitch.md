# snitch

[snitch lesson](https://university.scylladb.com/courses/scylla-essentials-overview/lessons/architecture/topic/snitch/)

## 🧐 What is a Snitch?

A **Snitch** in ScyllaDB helps the database **understand the network topology** — including how nodes are grouped into **data centers, racks, and regions**.

---

## 🛰️ What a Snitch Does

- Informs Scylla about **network layout**
- Affects **replica placement** (especially with `NetworkTopologyStrategy`)
- Optimizes **read/write routing**
- Enhances **fault tolerance** by spreading replicas across failure zones

---

## 📦 Why Snitch Matters

When using replication strategies, especially `NetworkTopologyStrategy`, the snitch provides the **context** so replicas are:

- Spread across **data centers** and **racks**
- **Not all in the same AZ** or node group
- Optimally placed for **latency and durability**

---

## 🔧 Types of Snitches

| Snitch                         | Description |
|-------------------------------|-------------|
| `SimpleSnitch`                | Treats all nodes equally. No DC/rack awareness. Best for dev/test. |
| `GossipingPropertyFileSnitch` | Recommended for production. Uses `cassandra-rackdc.properties` and gossip protocol. |
| `Ec2Snitch`                   | Automatically detects AWS EC2 region and AZ for topology. |
| `Ec2MultiRegionSnitch`        | Same as `Ec2Snitch`, but supports **multiple AWS regions**. |
| `RackInferringSnitch`         | Infers topology from IP address structure. Deprecated. Avoid. |

---

## 🗂️ Example Configuration (GossipingPropertyFileSnitch)

Create/edit the `cassandra-rackdc.properties` file on each node:

```properties
dc=dc1
rack=rack1
```

Then, set the snitch in `scylla.yaml`:

```yaml
endpoint_snitch: GossipingPropertyFileSnitch
```

---

## 🚨 Important Notes

- All nodes must use the **same snitch type**
- Changing the snitch on a live cluster is **risky** — can break replica placement!
- Snitch impacts how Scylla:
  - Chooses replicas
  - Routes queries
  - Handles consistency and availability

---

## ✅ Best Practice

Use **GossipingPropertyFileSnitch** for most environments unless you're on AWS and need dynamic AZ/region detection (`Ec2Snitch`).

---

Let me know if you want a YAML setup example or Docker-compose config!
