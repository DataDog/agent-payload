#!/bin/bash

protoc --gofast_out=. agent_payload.proto
# Go files will be written to the folder defined in the `go_package` option.
# So we need to move them into the correct folder.
cp ./github.com/DataDog/agent-payload/gogen/* gogen/
rm -rf ./github.com/

protoc --python_out=python agent_payload.proto
