#!/bin/bash

set -e

rm -rf openapi

docker run --rm \
    -v "${PWD}:/rockset" openapitools/openapi-generator-cli generate \
    -g go \
    -i /rockset/spec/spec2.yaml \
    --git-user-id rockset \
    --git-repo-id rockset-go-client \
    --additional-properties packageVersion=0.12.0,generateInterfaces=true \
    -o /rockset/openapi

rm openapi/go.*

go get -u ./...

go vet ./...
