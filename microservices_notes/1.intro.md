## Escaping the monolithic hell 

==> Complexity intimidates developers,  
* Software that becomes too complex for developers to understand can be hard 
to improve or fix bugs. 

* If the codebase is too difficult to understand developers cannot make changes 
* correctly.

* Development is slow *

* The Path from commit to deployment is long and arduous 
One problem with so many devs commiting to the same codebase is that the build 
is frequently in an unreleasable state. using feature branches to solve this problem can also result in lengthy painful merges, consequently, once a team completes its sprint, a long period of testing and code stabilization follows.

Because the codebase is soo large

* Scaliing is difficult 
When different application modules have conflicting resource requirements the application becomes difficult to scale. 

* Delivering a reliable monolith is challenging 

One reason for a monolith's unreliability is that testing is thouroughly difficult 
as a result we get frequent production outages since bugs make their way easily into production

* Locked into increasingly obsolete technology stack 

The final aspect of monolitich hell is that the architecture forces the team 
to use a tech stack that is becoming increasingly obsolete. The monolithic architecture makes it difficult to use new frameworks and language.


### Scale cube and microservices 

The scale cube defines three separate ways to scale an application: X-axis  
scaling load balances requests accross multiple, identical instances; 
Z-axis: scaling routes requests based on attributes of the request; 
Y-axis: functionally decomposes an application into services.

* A service is a mini application that implements narrowly focused functionality such as Order  management, customer management and so on.

* The high level definition of a microservice architecture is an architectural style  that decomposes an application into a set of services


#### Microservices as a form of modularity 

Modularity is essential when developing large, complex applications. Applications 
must be decomposed into modules if they are developed and understood by many people

In a monolithic application modules are defined using a combination of  
programming languages constructs (packages) and build artifacts. In practice 
monolithci apps tend to degenerate into big balls of mud  

The microservice architecture uses services as the unit of modularity. A service has an API, which is an impermeable boudary which is difficult to violate

* Each service has its own database. 

A key characteristic of the microservice architecture is that the service
are loosely coupled and communicate only via APIs. one way to achieve loose coupling is by each service having its own data store.

* FTGO application case study: 

The FTGO application's business logic consists of numerous backend services. Each backend service has a REST API and its own datastore (private). they include the following 

    - Order Service: Manage orders 
    - Delivery service: Manages delivery of orders from restaurants to consumers
    - Restaurants service: maintains information about restaurants 
    - Kitchen service: Manages the preparation of orders 
    - Accounting service: Handles billing and payments

#### Benefits and drawbacks of microservices 

* Benefits *
    - They enable continous delivery and deployment of large, complex applications 
    - Services are small and easily maintained 
    - Services are independently deployable 
    - Services are independently scalable 
    - The architecture enables teams to be autonomous 
    - It allows easy experimenting and adoption of new technologies 
    - It has better fault isolation

* Drawbacks * 
    - Finding the right set of services is challenging *
    - Distributed systems are complex, which makes development, testing, 
    and deployment difficult 
    - Deploying features that span multiple services requires careful coordination 
    - Deciding when to adopt the microservice architecture is difficult

### The microservice pattern language

A good way to describe the various architectural and design options and improve decision 
making is to use a pattern language. 

A software pattern solves a software architecture or design problem by defining a set 
of colaborating software elements.

The value of a pattern goes far beyond requiring you to consider the context of a problem. 
it forces you to describe other critical yet frequently overlooked aspects of a solution. 
A commonly used pattern structure includes three especially valuable sections: 
    - Forces 
    - Resulting context 
    - Related patterns

* Forces: The issues that you must address when solving a problem ** 

The forces section of a pattern describe the forces (issues) that you must address
when solving a problem in a given context. Forces can conflict, so it might not be 
possible to solve all of them. ** Which forces are more important depends on the context
. You have to prioritize solving some forces over others**

* Resulting Context: The Consequences of applying a pattern

The resulting context section of a pattern describes the consequences of applying the pattern.
It consist of three parts: 
    - Benefits: The benefits of the pattern, including the forces tha have been resolved
    - Drawbacks: The drawbacks of a pattern, including the unresolved forces *
    - Issues: The new problems that have been introduced by applying the pattern

** The resulting context provides a more complete and less biased view of the solution, 
which enables better design decisions.

* Related Patterns: The five different types of relationships 

The related patterns section of a pattern describes the relationship between the pattern 
and other patterns. There are five types of relationships between patterns: 

    - Predecessor: A predecessor pattern is a pattern that motivates the need for this 
    pattern. For example the microservice architecture pattern is a predecessor to the 
    rest of the patterns in the pattern language except the monolithic pattern

    - Successor: A pattern that solves an issue that has been introduced by this pattern. 
    For example, if you apply the microservice architecture pattern, you must then
    apply numerous successor patterns, including service discovery patterns and 
    circuit breakers patterns. 

    - Alternative: A pattern that provides an alternative solution to this pattern. for 
    example the monolithic architecture pattern  and the microservice architecture pattern s 
    are alternative ways of architecting an application

    - Generalization: A pattern that is a general solution to a problem.  

    - Specialization: A specialized form of a particular pattern.


In addition you can organize patterns into groups that tackle issues in a particular area.
The explicit description of related patterns provides a valuable guidance on how to 
effectively solve a particular problem.

A collection of patterns related through the following relationships: 

        Predecessor ---------- Successor 
        Alternative A ----------- Alternative B 
        General -------------- Specific

forms what is known as a pattern language. The patterns in a pattern language work together 
to solve problems in a particular domain.


### Overview of the microservice architecture pattern language 

The microservice pattern language is a collection of patterns that help 
one architect an application using the microservice architecture.

The pattern language consists of several groups of patterns, which are also divided 
into three layers. 

    - Infrastructure patterns: These solve problems that are mostly infrastructure issues 
    outside of development 
    - Application infrastructure: These are for infrastructure issues that also impact 
    development 
    - Application patterns: These solve problems faced by developers


#### Patterns for decomposing an application into services: 

Deciding how to decompose a system into a set of services is very much an art. there are 
two main decomposition  strategies: 

        - Decomposition by business capability : Organizes services around business capability
        - Decomposition by subdomain: Organizes services around domain driven design

#### Communication patterns

An application built using the microservice architecture is a distributed system. Consequently, 
interprocess communication (IPC) is an important part of the microservice architecture.

You must make a variety of architectural and design decisions about how your services communicate with one another and the outside world.

    - Communication style: What kind of IPC mechanism should you use? 
    - Discovery: How does a client of a service determine the IP address of a service
    instance so that, for example it makes an HTTP request? 
    - Reliability: How can you ensure that communication between services is reliable even 
    though the services are unavailable? 
    - Transactional messaging: How should you integrate the sending of messages and publishing 
    of events with database transactions that updated business data?

    - External API: how do clients of your application communicate with services?


* The five groups of communication patterns: 
        - Transactional messaging:  
            - Polling publisher 
            - Transactional log tailing 
            - Transactional outbox

        - Communication style: 
            - Messaging 
            - Remote procedure invocation 
            - Domain specific 

        - Discovery: 
            - Client-side discovery 
            - Self registration 
            - Service registry 
            - Server-side discovery 
            - 3rd party registration

        - Reliability: 
            - Circuit breaker 

        - External API:     
            - API gateway 
            - Backend for frontend


#### Data consistency patterns for implementing transaction managment

With microservices an application needs to maintain data consistency using the 
Saga pattern

#### Patterns for querying Data in microservice architecture

        - API composition 
        - CQRS

sometimes you can use API composition pattern, which invokes the APIs of one or  
more services and aggregates results. Other times, you must use the command query 
responsibility seggregation (CQRS) patter, which maintains one or more easily queried 
replicas of the data. 

#### Service deployment patterns

        - Virtual machines 
        - Containers 
        - Serverless

The traditional approach is to deploy services in a language specific package format. there 
are two modern approaches to deploying services. the first deploys services as VM or containers
The second is the serverless approach. You simply upload the services code and the serverless 
platform runs it

#### Observability patterns provide insights into an application behaviour

A key part of operating an application is understanding its runtime behavior and troubleshooting 
problems such as failed requests and high latency. *

You can use the following patterns to design observable services: 

    - Health check API: Exposes an endpoint that returns the health of the service
    - Log aggregation: Log service activity and write logs into a centralized logging server
    - Distributed tracing: Assign each external request a unique ID and trace requests as they
    flow between services
    - Exception tracking: Report exceptions to an exception tracking service, which duplicates
    exceptions, alerts developers, and tracks the resulotion of each exception 
    - Application metrics: Maintain metrics, such as counters and gauges and exposes them to a 
    metrics server. 
    - Audit logging: Log user actions

#### Patterns for the automated Testing of services 

The microserice architecture makes it easier to test individual services because they're much 
smaller than monolithic applications. 

    - Consumer driven contract tests: Verify that a service meets the expectations of its clients
    - Consumer side contract tests: Verify that the client of a service can communicate with the 
    service
    - Service component test: Test a service in isolation

#### Patterns for handling cross-cutting concerns 

When developing a new service, it would be too time consuming to reimplement these concerns
(Observability, discovery, externalized configuration ) 
from scratch. A much better approach is to apply the microservice Chassis pattern and build 
services on top of a framework that handles these concerns

#### Security Patterns 

In a microservices architecrture, users are typically authenticated by the API gateway. It 
must then pass information about the user, such as identity and roles, to the services it 
invokes. A common solution is to apply the access token pattern. The API gateway passes an 
access token, such as JWT (JSON Web Token), to the services, which can validate the token 
and obtain information about the user.