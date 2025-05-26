# review

To summarize, these are the main points we covered:

- ScyllaDB data modeling is query-based. That is, we think of the application workflow and the queries early on in the data model process
- A Keyspace is a top-level container that stores tables
- A Table is how ScyllaDB stores data and can be thought of as a set of rows and columns
- The Primary Key is composed of the Partition Key and Clustering Key
- One of our goals in data modeling is even data distribution. For that, we need to select a partition key correctly
- Selecting the Primary Key is very important and has a huge impact on query performance

## Data Modeling Process Notes

The data modeling process is **iterative** and involves continuous refinement. Here's a breakdown of the key steps:

1. **Update**  
   - Adjust the schema or application as needed.  
   - Ensure changes align with current requirements.  

2. **Application**  
   - Analyze and estimate **read and write patterns**.  
   - Understand how data is accessed and modified.  

3. **Test and Measure**  
   - Use **metrics** to evaluate the data model's performance.  
   - Identify bottlenecks or inefficiencies.  

4. **Data Model Optimization**  
   - Refine the model based on the **application's SLA** (Service Level Agreement).  
   - Prioritize performance, scalability, and reliability.  

### Key Takeaway

This cycle repeats to ensure the data model remains efficient and aligned with application needs.
