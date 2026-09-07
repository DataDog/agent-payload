# Review Guidelines

This repository is shared between many teams, and it is not always clear what a reviewer should be looking for.
The following guidelines are intended to help.
This is a _work in progress_ and additions to this list are welcome.

## Practicalities

If any of the `.proto` files are changed, then the corresponding Go code should be regenerated in the same PR (`GOPATH=$(go env GOPATH) inv codegen.all`).

If `proto/metrics/intake_v3.proto` changes, also regenerate `metrics/dd-metrics-v3/tests/pb/mod.rs` (`cargo build --features generate-protobuf` from within `metrics/dd-metrics-v3/`) and update `metrics/dd-metrics-v3/src/constants.rs` and the hand-rolled encoder in `metrics/dd-metrics-v3/src/writer.rs` to match — they don't use a Protocol Buffers library, so they won't fail to compile just because the wire format changed underneath them.

Tests should be run locally before making a PR, and should pass in CI before a PR is merged.

## Implications for Security

 * Where complex data parsing logic is implemented, ensure fuzzing tests are place to improve resilience against malformed payloads.

## Implications for the Agent

TBD

## Implications for Compatibility

When changing proto definitions:

 * Will a sender using the old definition be able to communicate with a receiver using the new definition?
 * Will a sender using the new definition be able to communicate with a receiver using the old definition?
 * If a field is removed, is the ID marked as `reserved` so that it will not be accidentally re-used?
