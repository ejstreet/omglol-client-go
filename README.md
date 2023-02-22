[![Tests](https://github.com/ejstreet/omglol-client-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/ejstreet/omglol-client-go/actions/workflows/test.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/ejstreet/omglol-client-go?status.svg)](https://pkg.go.dev/github.com/ejstreet/omglol-client-go?tab=doc)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/ejstreet/omglol-client-go?sort=semver)](https://github.com/ejstreet/omglol-client-go/releases)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ejstreet/omglol-client-go)](https://github.com/ejstreet/omglol-client-go/blob/main/go.mod#L3)
# omglol-client-go
API client library for omg.lol written in Go

> **Warning**
> The omg.lol API is still in development, and endpoints are subject to change. This may cause the client to break unexpectedly.

## Usage
To install this module:
```bash
go get github.com/ejstreet/omglol-client-go@latest
```
To import:
```go
import (
  "github.com/ejstreet/omglol-client-go/omglol"
 )
 ```
 
 Create a new client:
 ```go
 c, err := omglol.NewClient(email, api_key)
 ```
 
 Use a method:
 ```go
 a, err := c.GetAccountInfo()
 ```
 See test files for further examples.

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
||Update|PATCH|✔️|
||Delete|DELETE|✔️|
|PURL|List|GET|✔️|
||Retrieve|GET|✔️|
||Create|POST|✔️|
||Delete|DELETE|✔️|

>**Note** Features marked with a * are additional to what the API provides

>**Note** At the time of writing, `UpdateDNSRecord` currently does not work as described in the API docs. See https://github.com/neatnik/omg.lol/issues/584. `ReplaceDNSRecord` can be used instead.
