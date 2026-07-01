# ocpp-go

[![Go Reference](https://pkg.go.dev/badge/github.com/johnmaddison/ocpp-go.svg)](https://pkg.go.dev/github.com/johnmaddison/ocpp-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A Go library for building Open Charge Point Protocol (OCPP) clients and central systems.

`ocpp-go` provides typed OCPP message structs, JSON array serialization, a
WebSocket client, a WebSocket server, callback-based message routing, and
request/response helpers for sending calls and waiting for results.

The API is inspired by the Python
[`mobilityhouse/ocpp`](https://github.com/mobilityhouse/ocpp) project while
staying idiomatic for Go.

## Status
This project is still early. APIs may change while the library settles.

Current protocol support:

- OCPP 1.6 JSON: typed messages, client/server routing, callbacks, and request/response helpers.
- OCPP 2.1 JSON: typed messages, client/server routing, callbacks, and request/response helpers.

## Installation

```sh
go get github.com/johnmaddison/ocpp-go
```

Import the packages you need:

```go
import (
	"github.com/johnmaddison/ocpp-go"
	"github.com/johnmaddison/ocpp-go/client"
	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
	"github.com/johnmaddison/ocpp-go/server"
)
```

## Examples

The repository includes a runnable example central system and charge point.

In one terminal:

```sh
make run-server
```

In another terminal:

```sh
make run-client
```

The example server listens on `0.0.0.0:9001` and accepts charge points at:

```text
ws://127.0.0.1:9001/ocpp/{chargePointId}
```

## License

MIT License. See [LICENSE](LICENSE).
