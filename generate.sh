#!/bin/bash

protoc --gofast_out=go agent_payload.proto
protoc --python_out=python agent_payload.proto
