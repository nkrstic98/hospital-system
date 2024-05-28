#!/bin/bash

set -e

PROTO_DIR=../proto/

protoc -I=$PROTO_DIR \
        --go_out=. \
        --go_opt=paths=source_relative \
        --go-grpc_out=. \
        --go-grpc_opt=paths=source_relative \
        --validate_out=lang=go:. \
        --validate_opt=paths=source_relative \
        $(find $PROTO_DIR -iname '*.proto' -not -path $PROTO_DIR'/vendor/*')
