# Review Guidelines

This repository is shared between many teams, and it is not always clear what a reviewer should be looking for.
The following guidelines are intended to help.
This is a _work in progress_ and additions to this list are welcome.

## Practicalities

If any of the `.proto` files are changed, then the corresponding Go code should be regenerated in the same PR (`GOPATH=$(go env GOPATH) inv codegen.all`).

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
