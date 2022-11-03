# agent-payload

Payload format description for communication between the Agent and the Datadog backend.

This repository includes the protocol-buffer IDL used by the agent6 and agent7 to communicate with the Datadog backend.
Those payloads are only supported by the V2 API endpoints.
The generated Go, Python and Java implementations are checked into this repository and can be used directly.

# Prerequisites

You will need
 * Ruby (any version will do)
 * Go (at least the version in `go.mod`)
 * A checkout of this repository within a GOPATH (so, at `$GOPATH/src/github.com/DataDog/agent-payload`)

# Payloads

## Logs

The logs payload is defined in [`proto/logs/agent_logs_payload.proto`](./proto/logs/agent_logs_payload.proto).
The following implementations are available:
 * Go (gogofast): [github.com/DataDog/agent-payload/pb](https://pkg.go.dev/github.com/DataDog/agent-payload/pb)
 * Java: [`java/com/dd/agent/pb/Log.java`](./java/com/dd/agent/pb/Log.java)

## Metrics

The metrics payload is defined in [`proto/metrics/agent_payload.proto`](./proto/metrics/agent_payload.proto).
The following implementations are available:
 * Go (gogofast): [github.com/DataDog/agent-payload/gogen](https://pkg.go.dev/github.com/DataDog/agent-payload/gogen)
 * Python: [`python/agent_payload_pb2.py`](./python/agent_payload_pb2.py)

## Process

The process payload is defined in [`proto/process/agent.proto`](./proto/process/agent.proto).
The following implementations are available:
 * Go (gogofast): [github.com/DataDog/agent-payload/process](https://pkg.go.dev/github.com/DataDog/agent-payload/process) (note that this go package contains additional functionality beyond the generated PB implementation).

# Updating Proto Definitions

After updating the IDL you must:

- Regenerate the code: `GOPATH=$(go env GOPATH) rake codegen`
- Create a new tag with the updated version of the payload
