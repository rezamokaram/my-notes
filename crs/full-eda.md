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

# 3.2. DRY (Don`t repeat your self)    

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