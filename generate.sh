#!/bin/bash

protoc --proto_path=$GOPATH/src:. --gogofast_out=. agent_payload.proto
# Go files will be written to the folder defined in the `go_package` option.
# So we need to move them into the correct folder.
cp ./github.com/DataDog/agent-payload/gogen/* gogen/
rm -rf ./github.com/

protoc --proto_path=$GOPATH/src:. --python_out=python agent_payload.proto
