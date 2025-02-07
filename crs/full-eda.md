# 1.1.  


---  
---  
--- 

# 2.1.  


---  
---  
---  

# 3.1. DB per Microservice Principle

## Motivation

sharing the data bases between microservices tightly couples, their code bases and their teams that causing a lot of coordination overhead.  

## Rule
first rule of microservices databases is we must to have db per service  
if someone needs to access or change this data, we must use service api  
this helps us to scale our org:
if one team needs to change db schema or technology, that change will be completely transparent to the consumers of their api  
also if we need to change api, we can make a new version of api that helps other teams to migrate from old one to new api  

**important note**
if we do cache or store data that belongs to another microservice, then we loose strict consistency and have to settle eventually consistency. 

## Down side
- added latency  
- no more "join" operation
- no more transactions

# 3.2. DRY (Don't repeat your self)    

## Principle
if you repeat the same logic or data (constant values) you should consolidate it a shared:
- method
- class
- variable

## Challenges of shared libraries  
- tight coupling: each change in that lib needs to communicate with all teams that use it(api changes)(time waster). 
- even if the api do not change over an update to that lib we need **Rebuild, retest & redeployed** all microservices that use it.
- every bug or vulnerability in that lib impacts all microservices.
- dependency hell: if our microservice using lib a that use lib b internally; then our code uses lib b directly in another part of code, then on each update that the libs use new version of lib b, we must to do some unnecessary code changes to microservice code base. (two version of a lib breaks dry) (also it make our build time longe and also increases size of the binary file)

## Alternatives to shared libraries in microservices architecture
- new microservices
- using code generation tools (gRPC)
- Sidecar pattern
- code duplication across microservices (not inside a single microservice)

## Data duplicating in Microservices 
- important for performance reasons (it is inevitable and acceptable to dup the data across micros)
- Makes the data eventually consistent
- only one microservice needs to remain the owner of each data (SSoT - single source of truth)


# 3.3. Structured Autonomy for development teams

## problem with full team autonomy

**Myth:** each team can choose their stack including database, tech stack, frameworks, tools & APIs

**reasons**

1. upfront cost of infrastructure
2. infrastructure maintenance 
3. steep learning curve 
3. non-uniform api

q: isn`t the point of microservices to allow teams independence ?
ans: well yes, but the key to a successful microservices is the balance between **team autonomy and structure**. -> Structured Autonomy 

## 3 tiers of development team autonomy 

1. Tier 1: Fully restrictive: describes the areas that should not be under each team`s jurisdiction. instead those areas should be uniform across the entire company.
- infrastructure : monitoring & alerting & CI/CD
- API guidelines and best practices
- security and data compliance 

2. Tier 2: Freedom with boundaries
- programming languages 
- database technology

3. Tier 3: Complete Autonomy
- release processes
- release schedule and frequency
- custom scripts for local development and testing
- Documentation
- on boarding for new developers

## factors in team autonomy boundaries
1. size / influence of DevOps / SRE team: typically, companies with more dominant DevOps or Sre teams lean more towards common standards, which makes their life managing the system easier. 
2. seniority of developers: generally, more senior the developers are, the more freedom they prefer in setting up or building their own infrastructure. 
3. company's culture: for example, some companies just stick in to one programming languages like golang, and don't allow any freedom in choosing other programming languages. -> benefit: we can move the developers between teams with very little overhead.

# 3.4. micro-frontend services architecture pattern

- we have all same problems with monolithic backend in monolithic frontend
- we should to migrate to micro frontend architecture

**benefits**
- replaced the complex monolithic codebase with small and much more manageable code bases. each for a different micro-frontend. 
- full-stack ownership of each micro-frontend
- separate ci/cd
- separate release schedule

**best practices**
- micro-frontend services are loaded at runtime 
- no share state in the browser
- internal communication through: custom events, callbacks & address bar

# 3.5. api management microservices architecture

**3 types of api**
1. private: to our microservices within the company boundary
2. partner: to external businesses
3. public: to user

## api management problems
- tight coupling of api endpoints to client side code
- different types api for different consumers (public/private/partner)
- different api tiers based on subscription level
- traffic control and monitoring

## api gateway pattern

it is responsible for all the api management

apigateway must route the request to the destination service(transform the data to the service input and service out put to api format)
load balancer task is to balance the requests for a single service
we deployed a lb for each micro service

**lb:** 
- little performance overhead
- health check
- different routing algorithm

**api gateway**
- throttling
- monitoring
- api versioning
- protocol / data translation
- authorization (??) and tls termination

---  
---  
---  

# 4.1. introduction to eda arch

## motivation

sync requests have a high response time -> it solves in eda with another solution
also load balancing in eda is better

## main concept of "EVENT"

- fact, action, state change
- always immutable
- can be stored indefinitely
- can be consumed multiple times by different services 

## request & response / event-driven model
- synchronous / async
- inversion of control
- loose coupling

# 4.2. use case and patterns of eda

**use cases:**
- fire and forget
- reliable delivery
- infinite stream of events
- anomaly detection/pattern recognition
- broadcasting

**req/res model use cases**
- immediate response with data is needed

note: the important thing is that we should not to use eda every where, the best solution is to design the system first and just use eda where its really needed (even if we use eda every where its better to design first)

## event delivery patterns

- event streaming: in this pattern, the message broker is used as a temporary or permanent storage for events. the consumers have full access to logs of that events, even if they have already been consumed by the same consumer or other consumers.
1. reliable delivery
2. pattern / anomaly detection

- pub/sub: in this patterns the consumers subscribe to a particular topic or channel to receive only new events after subscribing. in this case subscribers don't have access to old events, and as soon as all the current subscribers receive the event, the message broker will typically delete it from when a consumer, consumes the event.
this pattern is ideal when the message broker is used only as a temporary storage or broadcasting mechanism. after the subscribers consume the events, they are typically transformed and stored permanently in a database or pass it to another service.

1. messaging system is a temporary storage
2. fire and forget
3. broadcasting
4. buffering
5. infinite stream of events

## use cases that eda is not a good solution
1. need an immediate response containing data
2. simple interactions

# 4.3. message delivery semantic in eda

## message delivery problem

it is possible when we send a request to the server then we don't get any response from it:
1. server crashes on that request and db don't update
2. server update db successfully and after that crashes or cannot send response

## event delivery problem in eda
event may lost or receive more than 1 time in consumer side:
1. in memory messaging system and it goes down
2. on receiving server crashes before update db
3. on consumer, server can not send nack after db update
4. ...

## delivery semantics

**1. at-most-once delivery semantic:**  
in producer, if we send the message to messaging system and don't get the acknowledgment we don't send it again  
also in subscriber side, we send the acknowledgment before processes the message, if server crashes any where or any similar problem, we don't process it again and it should be ok with our scenario
- data loss is ok
- least overhead / lowest latency

**2. at-least-once delivery semantic:**  
in publisher side, we send the event into message broker and if we don't get acknowledgment from broker we send it again until get acknowledgment once
in subscriber side we get the event from broker and process it and then we send a acknowledgment to to broker

- data loss is un acceptable
- data duplication is ok

**3. exactly-once delivery semantic:**  
- most difficult to achieve
- highest overhead / latency
we need at least once semantic + corelation id (ident potency for event or item) in publisher side
it means each event have its own unique id that identify the event and we need to send the event to message broker and receive the acknowledgement. if we don't get it we need to send again.
optionally: if the broker have the event id in their logs it do not accept the event, and if there is no log with that id, accept the message.
in subscriber side, we get the event from the broker, if we don't have the id in our db we process it and write to db, and if we have it it means it was processed perviously.

---  
---  
---

# 5.1. Saga Pattern

**in software architecture every thing is a trade-off**

each micro have its own database and now we loose database transactions for the use cases that we want to change some DBs.

## Saga Pattern
saga pattern help us perform a distributed transaction that spans multiple microservices and databases.

## Saga impl

**1. Workflow Orchestration**  
there is a workflow orchestration service. the sole purpose of this service is to orchestrate the transaction in the correct order and also apply the compensating operations in the opposite order if things go wrong.

**2. Event-Driven Model (Choreography)**  
each service handles its own businesses and if there is a need to speak to other services, just fire an event.  
this means there is a chain of services and events with a specific order. if any problem happens in the chain, rollback started from there and each node fire an event to previous node for rollback the operation.
  
**note:** in eda, for rollback and canceling orders or requests, we need a notification box for each user to send the result to it.

# 5.2. CQRS Pattern  

cqrs -> command and query responsibility segregation

we can segregate the actions on data into two type:

1. command: change the data (insert, update, delete)  
2. query: read the data and data will never change (get, read)


**CQRS Benefits**  
- Separation of concerns
- Higher performance for read and write operation (we can choose right infra for each part that support better performance)
- Higher scalability
- solves the problem of joining data from multiple microservices  

one of cqrs implementations is to use eda and make a eventually consistency between query db and command db. 

another way is to use something special in infra like OLAP

# 5.3. Event Sourcing Pattern

## event sourcing benefits:
- visualize  
- auditing
- corrections
- high write performance (no db lock needed)  

in traditional DBs we loose pervious state of data.  
for saving previous state of data we use event sourcing.  
**note** events (in event store) are always immutable (it never changes) and we just have insert into event store.  

for reaching the last state of data we apply all events into the entity. (it is like a version control for data)

## event sourcing impl
1. db - one row per event (good for query, bad for high load)
2. message broker - store in message broker (good for high load, we have no query)

## how to reach last data state

of course replacing all the transactions in someone bank account every time we want to show the client their balance is not very efficient.

so to address that we can apply few strategies:
- snapshots: taking snapshot at certain points of the events log. (each month or week or day)
- cqrs: command service for events and event store, for each new event we need to fire a new event into queue for query service that applies the latest changes and save the last state of data for queries. (eda, eventually consistent)  
**note:** in command service we don't need any special database. (it is not necessary)
  
**note:** the combination between cqrs and event sourcing is very popular in the industry.  
reasons:  
- we get history and auditing
- we get fast and efficient writes
- we get fast and efficient reads

**note again:** this brilliant pattern is not for free. we can only implement it with eventual consistent solution and we loss the string consistency. which may or may be not good enough depending on your use case.

---  
---  
---  

# 6.1. testing pyramid for microservices

## recap of the testing pyramid for monolithic applications  
1. functional / end-to-end tests
2. integration tests
3. unit tests

top is 1 and down is 3

**unit test:** test a small logic such as a class or module in isolation. unit test are cheapest to maintain because they are small, easy to write and fast to execute. we should also create the highest number of those tests compared to other tests. thats the reason of why they are located at the bottom of test pyramid.
this type of test give us least confidence, because this type tests each unit in isolation. once we run the application, we have no idea if all those units works together or not.

**integration test:**
those tests verify that different units and systems we integrated with, such as database and message broker actually work together.
- bigger
- slower
- medium number
- more confidence

**functional or end-to-end test:**
those test run our entire system, which includes the ui, our entire application and the database, and they verify that it works as intended.
- each such test should test a particular user journey or business requirement and ensure that it matches the specification.
- heaviest
- most complex
- slowest to run
- least number of them compare to other tests
- highest confidence

## how to apply the testing pyramid microservices architecture

each microservice team should make the exact pyramid for their own microservice and database. then we treat each microservice as small unit that is part of the larger system and put it in a larger testing pyramid.

so just like in the case of unit tests, testing each microservice in isolation is essential, but not enough to increase the confidence that all those microservice can actually talk to each other at runtime and work correctly together.

we need to add another layer of integration tests. those integration tests verify that every pair of microservices can talk to each other using the agreed upon API while mocking the rest of our system.  
to complete this pyramid we need to add the system level end to end tests at the very top. those tests, in theory, should run all of our microservices, front, databases and ... in a test environment and verify that all the individual components work together as expected. 

## challenges of testing microservices and eda

**challenges:**
1. end-to-end tests:  
- hard to set up
- hard to maintain
- no clarity about ownership
- one team may block every one
- low confidence (ignoring failed tests)
- very costly

company challenges:
- some companies invest too much on those few tests
- some other companies do not invest at all

2. integration tests:
- difficult to run
- tightly couples teams to each other

# 6.2. Contract tests and production testing

## integration tests using lightweight mocks
with this solution we don't need to use another microservice for testing our microservice. we just need to mock the API layer..
there is still another problem with changes in api layer of each microservice. in this case both of microservices passed their tests successfully but after release they cant communicate with each other in production.

## contract tests for synchronous communication

in eda microservices, we make a contract from our event in microservice A and then share it with another team for microservice B. this approach help us to validate that our microservices are using the correct format for events. then we can confidently release it to production.

## end-to-end Tests Alternatives

**blue-green deployment + canary testing**  
a blue green deployment is a safe way to release a new microservice version to production using two identical production environments without any downtime during the release.
the blue environment is a set of servers or containers that run our old version, and the green environment is a set of servers or containers that run the new version that we want to release.  once we deployed the new version to the green environment, no real user traffic is going to it. after we run automated and even manual tests on the green environment, we can shift a portion of the production traffic to the green  environment and monitor the new version for performance and functional issues. this process is called canary testing. if we detect an issue we immediately direct traffic back from the green environment to the blue environment with minimal impact on users.  
on other hand if no issue detected, we direct all the production traffic from the blue environment to the green and gradually decommission the blue environment since its no longer needed.  

---  
---  
---  
