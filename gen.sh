#!/bin/bash
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
protoc src/prime_proto/prime.proto --go_out=. --go-grpc_out=.