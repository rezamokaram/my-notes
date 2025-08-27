# Idempotency

Let's dive into idempotency in software engineering.

Imagine you're pressing a light switch. Whether you press it once or multiple times, the light ends up in the same state – either on or off. This "doing it again has the same result" idea is the essence of idempotency.

In software engineering, an operation is **idempotent** if performing it multiple times has the same effect as performing it once. In other words, after the initial execution, subsequent executions do not change the state of the system in a way that matters.

Here's a breakdown of why this is important and some examples:

**Why is Idempotency Important?**

* **Resilience and Reliability:** In distributed systems, network issues can lead to requests being sent multiple times. Idempotency ensures that these duplicate requests don't cause unintended side effects like double charges or inconsistent data.
* **Error Recovery:** If an operation fails partway through, retrying an idempotent operation is safe. You can be confident that the system will eventually reach the desired state without corruption.
* **Simplified Logic:** Designing systems with idempotent operations can simplify the logic for handling retries and error conditions.

## Examples of Idempotent Operations

* **HTTP `GET`:** Retrieving a resource. Making the same `GET` request multiple times will always return the same resource (as long as it hasn't been modified).
* **HTTP `PUT` (for updating a resource with a specific state):** If you send a `PUT` request to set a user's email to "<test@example.com>", sending the same request again will keep the email at "<test@example.com>". The system's state remains the same.
* **HTTP `DELETE`:** Deleting a resource. Once a resource is deleted, subsequent `DELETE` requests for the same resource will typically result in a "resource not found" error, but the system's overall state regarding that resource (it being absent) remains consistent.
* **Setting a value in a database:** An operation like `UPDATE users SET is_active = TRUE WHERE id = 123` is idempotent. Running it multiple times will always result in the `is_active` flag for user 123 being `TRUE`.

**Examples of Non-Idempotent Operations:**

* **HTTP `POST` (for creating a new resource):** Each `POST` request typically creates a new resource. Sending the same `POST` request multiple times will result in multiple identical (or slightly different, depending on the API) resources being created.
* **Adding an item to a shopping cart:** Each time you click "add to cart," a new item is added. Multiple clicks result in multiple items.
* **Incrementing a counter:** Each increment operation changes the counter's value.

**How to Achieve Idempotency:**

* **Unique Identifiers:** Assigning unique IDs to requests allows the server to identify and discard duplicate requests.
* **State Checks:** Before performing an operation, check the current state of the resource. If it's already in the desired state, no further action is needed.
* **Token-Based Approach:** Using unique tokens for operations. The server processes the request and marks the token as used. Subsequent requests with the same token are ignored.
* **Database Constraints:** Utilizing database constraints (e.g., unique constraints) can help prevent the creation of duplicate entries.

In essence, striving for idempotency in your software systems, especially in distributed environments, leads to more robust, predictable, and easier-to-manage applications.
