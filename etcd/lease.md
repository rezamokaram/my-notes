# Lease

In etcd, a lease is a mechanism used to associate a time-to-live (TTL) with keys. It is primarily used to manage the lifecycle of keys, enabling automatic expiration and cleanup of keys that are no longer needed. This feature is essential for distributed systems, as it helps maintain consistency and ensure that stale or unused data does not persist indefinitely.

## Key Concepts of Leases in etcd:
### 1. Lease Creation:

    - A lease is created with a specified TTL (time-to-live), which defines how long it will exist.
    - The TTL is measured in seconds, and the lease must be periodically refreshed to remain valid.

### 2. Association with Keys:

    - Keys can be attached to a lease. If the lease expires, all associated keys are automatically deleted.
    - This is useful for scenarios where temporary keys are required, such as leader election or service registration in a service discovery system.

### 3. Lease Renewal:

    - Clients can renew a lease before its TTL expires using the KeepAlive operation.
    - If the lease is not renewed within its TTL, it will expire, and the associated keys will be deleted.

### 4. Lease Expiry:

    - Once a lease expires, etcd deletes all keys attached to that lease.
    - Expired leases help maintain a clean keyspace by removing unused data.

### 5. Revoke:

    - Leases can be explicitly revoked by clients using the Revoke operation, which immediately removes the lease and its associated keys.

### 6. Lease ID:

    - Each lease is identified by a unique lease ID, which is used to associate keys with the lease and manage it.

## Common Use Cases:
- Service Discovery: Services register themselves with etcd by associating their metadata with a lease. If the service fails or stops renewing the lease, its entry is automatically removed.
- Leader Election: Leaders in distributed systems can use a lease to hold their leadership status. If the leader fails to renew the lease, it automatically loses leadership, allowing others to take over.
- Temporary Key Storage: Keys that should exist only for a limited duration can be associated with a lease to ensure they are automatically deleted after a certain time.

## Example Workflow in etcd:

1. Create a Lease:

``` sh
lease grant 10
```

This creates a lease with a TTL of 10 seconds.

1. Attach Keys to Lease:

``` sh
    put key1 value1 --lease=<lease_id>
```

3. Renew the Lease:

``` sh
    lease keep-alive <lease_id>
```

4. Let the Lease Expire or Revoke It: If the lease is not renewed, it will expire after 10 seconds, and key1 will be automatically deleted. Alternatively, you can revoke it manually:

```sh
    lease revoke <lease_id>
```


By leveraging leases, etcd provides an efficient and automatic way to manage the lifecycle of keys in a distributed environment.