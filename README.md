# agent-payload

Payload format description for communication between the Agent and the Datadog backend.

This repository includes the protocol-buffer IDL used by the agent6 to communicate with the Datadog backend.
Those payloads are only supported by the V2 API endpoints.

# Generated code

The Go, Python and Java code are already generated along the `.proto` file.

# Updates

After updating the IDL you must:

- Regenerate the code: `rake codegen`
- Create a new tag with the updated version of the payload
