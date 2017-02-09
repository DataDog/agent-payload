# agent-payload

Payload format description for communication between the Agent and the Datadog
backend

This repository includes the protocol-buffer IDL use by the agent6 to
communicates with the Datadog backend. Those payloads are only supported by the
V2 API endpoints.

# Generated code

The Go and Python code are already generated along the `.proto` file.

# Updates

You will need to install ProtoGoFast: github.com/gogo/protobuf

After updating the IDL you must:

- Regenerate the Go and Python code: ./generate.sh
- Create a new tag with the updated version of the payload
