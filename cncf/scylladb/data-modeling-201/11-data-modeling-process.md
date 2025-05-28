# data modeling process

## we can think of specific queries, what data will be written/read, and how often. Some of the queries might be

- Query 1: User logs in to the application
- Query 2: Update a pet’s health information
- Query 3: Read pet’s health information on a given day

In a real-world use case, we would have to develop these queries further and think about the conceptual data model (the entities and the relationships between them).

## Some important things to keep in mind

- Aim at creating a single partition per query. If a query needs to access only one partition, it would be very efficient. If multiple partitions need to be accessed for a single query, this can be acceptable, but it would be less efficient. If multiple partitions are being accessed for a query that’s being used often, we get less efficiency and maybe something in the data model is wrong.
- Avoid scanning the entire cluster to find the data
- Avoid scanning an entire table for the information needed (linear search).
