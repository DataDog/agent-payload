# agent-payload

Payload format description for communication between the Agent and the Datadog backend.

This repository includes the protocol-buffer IDL used by the agent6 and agent7 to communicate with the Datadog backend.
Those payloads are only supported by the V2 API endpoints.
The generated Go, and Java implementations are checked into this repository and can be used directly. Other consumers may copy the `.proto` files into their repository and generate their own bindings.

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

## Process

The process payload is defined in [`proto/process/agent.proto`](./proto/process/agent.proto).
The following implementations are available:
 * Go (gogofast): [github.com/DataDog/agent-payload/process](https://pkg.go.dev/github.com/DataDog/agent-payload/process) (note that this go package contains additional functionality beyond the generated PB implementation).

## CWS

The CWS security dumps payload is defined in [`proto/cws/dumpsv1/activity_dump.proto`](./proto/cws/dumpsv1/activity_dump.proto).
The following implementations are available:
 * Go (vtproto): [github.com/DataDog/agent-payload/v5/cws/dumpsv1](https://pkg.go.dev/github.com/DataDog/agent-payload/v5/cws/dumpsv1)
 * Java: [`java/com/dd/cws/adv1/pb/SecDumpProto.java`](./java/com/dd/cws/adv1/pb/SecDumpProto.java)

# Updating Proto Definitions

After updating the IDL you must:

- Regenerate the code: `rake codegen`, rake will use gimme to run the rake command with the current defined go version
- If you have indentation/newlines changes, run `rake codegen` with the same Go version as defined in `go.mod`
- Create a new tag with the updated version of the payload

# Publishing Changes

After merging changes to `master` create a release by:

1. Navigate to the [Releases](https://github.com/DataDog/agent-payload/releases) page
2. Click "Draft a new release"
3. In the "Choose a tag" drop down, type in the next version number
   
   Generally you can add one to the last version number.  Make sure to include the `v` prefix. For example, if the last release was  v5.0.37, your release should be v5.0.38.
   
5. The release title should be the same as the version tag
6. Use "Generate release notes" to fill in the release description
7. Click "Publish release"
   
   This will create a git tag that can now be referenced in other repos. 
