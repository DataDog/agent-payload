# Review Guidelines

This repository is shared between many teams, and it is not always clear what a reviewer should be looking for.
The following guidelines are intended to help.
This is a _work in progress_ and additions to this list are welcome.

## Implications for the Agent

TBD

## Implications for Compatibility

When changing proto definitions:

 * Will a sender using the old definition be able to communicate with a receiver using the new definition?
 * Will a sender using the new definition be able to communicate with a receiver using the old definition?
 * If a field is removed, is the ID marked as `reserved` so that it will not be accidentally re-used?
