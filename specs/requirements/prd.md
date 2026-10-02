# Greeter — PRD

## Problem Statement

Teams building small demos and integrations often need a trivial, dependable HTTP endpoint to greet a caller by name — useful for smoke-testing new services, pipelines, or client integrations without standing up anything complex. Today they either reach for an ad-hoc script or skip the exercise entirely, losing a quick, consistent way to validate connectivity end to end. S0 marker s0-p1-1002a.

## Solution

Greeter is a small Go HTTP service exposing a single endpoint: given a name, it returns a JSON greeting. It is public and requires no sign-in, making it trivial to call from any client, script, or pipeline.

## Actors

- **API Consumer** — any client (script, service, or developer) that calls the greeter endpoint over HTTP. No account or sign-in is required.

## User Stories

1. As an API Consumer, I want to send a GET request to /hello with a name query parameter, so that I receive a JSON greeting personalized with that name.
2. As an API Consumer, I want to receive a sensible default greeting when I omit the name parameter, so that the endpoint still responds usefully instead of erroring.

## Product Decisions

- Sign-in: the /hello endpoint is public and requires no authentication — it is a fully open, unauthenticated API.
- Missing/empty name handling: when the name query parameter is missing or empty, the service falls back to a default name (e.g. "World") and still returns a 200 response with a greeting, rather than an error.
- No external services are required — the service has no third-party dependencies.
- No persistence — the service is stateless; it does not store or look up data.

## Out of Scope

- User accounts, sign-in, or any authentication/authorization.
- Any endpoint other than the single greeting endpoint.
- Persisting greetings, names, or request history.
- Localization or multi-language greetings.

## Open Questions

None — the brief was fully settled through the interview.

## Further Notes

None.