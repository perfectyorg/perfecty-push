# Design

## Introduction

Different considerations for the project are described
below, and they're a good starting point to understand the goal and nature of the project.
This includes the planned product, architecture considerations, principles, potential
libraries and a rough idea about the API.

Internal quality in Software Development is cheaper in the long run [[1]](https://martinfowler.com/articles/is-quality-worth-cost.html).
This means that a high effort will be put in having a good foundation for the project, so that
it's possible to iterate faster, cheaper and safer. Apart from the CISQ quality model, the project should embrace:

- Unit Testing
- Continuous Integration/Deployment
- Metrics
- Refactoring

## Outcome

The final version of Perfecty Push will extend the current capabilities
offered by the [plugin version](https://github.com/rwngallego/perfecty-push-wp/).
The development will be iterative
and in general, it should:

- Support both Push API and Websockets notifications.
- Have a Push Engine independent of the notification mechanism.
- Support user segmentation (device, location, engagement).
- Optional user presence indicators.
- Single binary and multiplatform: Windows/Linux.
- Exponential Backoff for third-party integrations.
- Unified JS SDK.

### MVP

A potential MVP is the version that supports the same features as the
plugin version, with basic user segmentation support, completely distributable in the plugin marketplace, and stable. This means:

- Supports Push API initially
- Single binary and multiplatform
- Basic user segmentation (Excluding engagements and meta attributes)
- Unified JS SDK

## Architecture

In this Go project there are two possible paths, which will be tried both, in phases:

- A lean Go architecture with a flat structure
- A more complex project focused on the domain and the maintainability:
  - Domain Driven Design
    - https://threedots.tech/post/ddd-lite-in-go-introduction/
    - https://vaadin.com/learn/tutorials/ddd/tactical_domain_driven_design
  - Hexagonal Architecture
    - https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3
    - https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
  - Clean architecture
    - https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
    - https://eminetto.medium.com/clean-architecture-using-golang-b63587aa5e3f
    - https://threedots.tech/post/ddd-cqrs-clean-architecture-combined/

The first approach is more common in the
Go/Elixir communities for simple projects, while the second comes
from the Java/.NET world. There have been some intentions to apply 
the latter in the former, so it's a good experiment for this project,
in case we consider it's worth it:

*"The purpose of a good architecture is to defer decisions"*,
[[*]](https://work.stevegrossi.com/2014/06/02/clean-architecture-and-design/#:~:text=%E2%80%9CUncle%E2%80%9D%20Bob%20Martin's%20talk%20at,I%20think%20about%20software%20architecture.&text=long%20as%20possible.-,The%20purpose%20of%20a%20good,to%20defer%20decisions%2C%20delay%20decisions.)

**The key here is to embrace unit testing and refactoring.**

### Components

According to Clean Architecture, to guarantee a loosely coupled
system we need to correctly apply the Dependency Inversion Principle.
A good initial architecture that could enable us to apply TDD
and DDD in a future refactoring is:

- Adapters - `/internal/handlers/` and `/internal/repositories/`: HTTP handlers and DB repositories
- Application - `/internal/application/`: Use cases
- Domain - `/internal/domain/`: Called Entities but in Domain Driven Design it's simply the "domain"
- Frameworks&Drivers - All the rest: The outer layer that glues the next inner layer (Adapters).

The important part is to have in mind the separation of concerns
and the Dependency Inversion Principle.

### Project layout

Follows:
https://github.com/golang-standards/project-layout

### HA and Fault Tolerance

For the long run, it's expected that Perfecty Push supports
High Availability and Fault Tolerance. This means implementing the
features that allow multinode execution, concurrency and (eventual/strong <?>)
consistency.

Libraries:
- Strong consistency:
  - https://pkg.go.dev/go.etcd.io/etcd/raft/
- Eventual consistency:
  - https://github.com/iwanbk/bcache
  - https://github.com/weaveworks/mesh
  - https://riunet.upv.es/bitstream/handle/10251/54786/TFMLeticiaPascual.pdf

While Strong Consistency brings reliability, and a guarantee on
the data when there are Partitions, it has an overhead in performance. In case the performance is
more important, we should rely on Eventual consistency instead.
Remember, we cannot have both (CAP).

## Considerations

### User segmentation

It's important to support user segmentation based on multiple
parameters:
- Device type
- Browser
- Location
- Timezone
- Engagement (depends on the **User presence indicators** described below)
- Meta attributes set through the SDK (Don't implement it in V1 yet, but have it in mind)

### User presence

We need to know the user presence by:
- Measuring the time spent in the website.
- Heartbeat mechanism to know if the user is currently connected. It should be optionally enabled.

### Independent Push Engine

The push engine must be unaware of the notification mechanism.
The specifics are implemented by the mechanisms, which take
work from the Push Engine queue to use a back-pressure strategy.

Initially supported mechanisms:

- [Push API](https://developer.mozilla.org/en-US/docs/Web/API/Push_API)
- [Web socket](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_servers)

This means we could later add others (APN, FCM).

### JS SDK

It's the JS SDK's responsibility to choose between Push API or
Web Sockets. Initially, Push API is tried, Web Sockets is the fallback.

The JS SDK will be based on the current implementation from the WP version.
A separate project should be created and the quality must be improved.

### Metrics

Metrics is specially important for a good Push Notifications Server. In this case,
we start from the baseline that the plugin version currently has a limit in this aspect.
Furthermore, it's relevant for the project to be based on metrics, for this we can use:

- Golang benchmarking
- [Prometheus](https://prometheus.io/docs/guides/go-application/) / InfuxDB

### Exponential Backoff with third-party integrations

The mechanisms that depend on third-party services
like `Push API` or `APN` or `FCM`, should detect transient failures
in the communication with the third-party. Errors that are not related to dead endpoints (403/410),
should automatically increase the retry factor and stop at a maximum value.

If the maximum is reached, it's a symptom that either we have:

- Provided the wrong VAPID credentials. This happens in some services
  that return `401` for endpoints that were expired, 
  see [this issue](https://github.com/mozilla-services/autopush/issues/1436).
- We have a failure in our side that we need to fix first
  (any miss-configuration or bug)
              
In those cases we don't want to re-try as it's considered a permanent
failure that we need to address first.

Libraries:
- https://github.com/cenkalti/backoff
- https://github.com/jpillora/backoff

## Public Endpoints

endpoint | method | description
--- | --- | ---
`/v1/public/users`  | PUT | Register user
`/v1/public/user/:uuid/preferences` | PUT |  Change the preferences (opt-in/opt-out)
`/v1/public/user/:uuid/heartbeat` | PUT | Let the server know when the user is using the app

## Internal Endpoints

### Users

endpoint | method | description
--- | --- | ---
`/v1/users`  | GET | List the users
`/v1/users/:uuid`  | GET | Get the user
`/v1/users/:uuid`  | DELETE | Delete the user
`/v1/users/:uuid`  | PUT | Update user
`/v1/users/stats` | GET | Get the users stats

### Notifications

endpoint | method | description
--- | --- | ---
`/v1/notifications`  | GET | List the notifications
`/v1/notifications`  | PUT | Send a new notification
`/v1/notifications/:uuid`  | GET | Get the notification
`/v1/notifications/:uuid`  | DELETE | Delete the notification
`/v1/notifications/stats` | GET | Get the notifications stats
