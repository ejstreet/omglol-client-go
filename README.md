[![Tests](https://github.com/ejstreet/omglol-client-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/ejstreet/omglol-client-go/actions/workflows/test.yml)
# omglol-client-go
API client library for omg.lol written in Go

> **Warning**
> The omg.lol API is still in development, and endpoints are subject to change. This may cause the client to break unexpectedly.

## Supported Features
This project is a work-in-progress, see the following table for supported features:

| Resource | Sub-Resource | Method | Supported |
|-|-|-|-|
|Account|Info|GET|✔️|
||Addresses|GET|✔️|
||Name|GET|✔️|
|||POST|✔️|
||Sessions|GET|✔️|
|||DELETE|✔️|
||Settings|GET|✔️|
|||POST|✔️|
|Address|Availability|GET|✔️|
||Expiration|GET|✔️|
||Info|GET|✔️|
|Directory|List|GET|✔️|
|DNS|ListRecords|GET|✔️|
||FilterRecord*|GET|✔️|
||Create|POST|✔️|
||Update|PATCH|⚠️|
||Replace*|DELETE/POST|✔️|
||Delete|DELETE|✔️|
|PURL|List|GET|✔️|
||Retrieve|GET|✔️|
||Create|POST|✔️|
||Delete|DELETE|✔️|

>**Note** Features marked with a * are additional to what the API provides

>**Note** At the time of writing, `UpdateDNSRecord` currently does not work as described in the API docs. See https://github.com/neatnik/omg.lol/issues/584. `ReplaceDNSRecord` can be used instead.