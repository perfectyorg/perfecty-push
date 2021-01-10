# Design

## Introduction

Different considerations for the project are described
below, and are a good starting point to understand the goal and nature of the project.
This includes the planned product, architecture considerations, principles, potential
libraries and a rough idea about the API.

## End product

The final version of Perfecty Push will extend the current capabilities
offered by the [plugin version](https://github.com/rwngallego/perfecty-push-wp/).
The development will be iterative
and in total, it should include:

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

Internal quality in Software Development is cheaper in the long run [[1]](https://martinfowler.com/articles/is-quality-worth-cost.html).
This means that a high effort will be put in having a good foundation for the project, so that
it's possible to iterate faster, cheaper and safer. Apart from the CISQ quality model, the project should embrace:

- Unit Testing
- Continuous Integration/Deployment
- Metrics
- Refactoring

In this Go project there are two possible paths, which will be tried both:

- A lean Go architecture with a flat structure
- A more complex project focused on the domain and the maintainability:
  - Domain Driven Design
    - https://threedots.tech/post/ddd-lite-in-go-introduction/
    - https://vaadin.com/learn/tutorials/ddd/tactical_domain_driven_design
  - Hexagonal Architecture
    - https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3
    - https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

It's not possible to decide right now which is the best option, given that the first
item is the Go/Elixir community's philosophy, and the second comes
from the Java/.NET world. There have been some intentions to apply 
the latter in the former, so it's a good experiment for this project too, specially
in the initial stage. **The key here is to embrace unit testing and refactoring to try them both <?>**.

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
`/public/users`  | PUT | Register user
`/public/user/:uuid/preferences` | PUT |  Change the preferences (opt-in/opt-out)
`/public/user/:uuid/heartbeat` | PUT | Let the server know when the user is using the app

## Internal Endpoints

### Users

endpoint | method | description
--- | --- | ---
`/users`  | GET | List the users
`/users/:uuid`  | GET | Get the user
`/users/:uuid`  | DELETE | Delete the user
`/users/:uuid`  | PUT | Update user
`/users/stats` | GET | Get the users stats

### Notifications

endpoint | method | description
--- | --- | ---
`/notifications`  | GET | List the notifications
`/notifications`  | PUT | Send a new notification
`/notifications/:uuid`  | GET | Get the notification
`/notifications/:uuid`  | DELETE | Delete the notification
`/notifications/stats` | GET | Get the notifications stats
