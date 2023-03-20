#!/bin/bash

set -e

VERSION=$(grep --only-match "[0-9]*\.[0-9]*\.[0-9]*" version.go)

rm -rf openapi/*

docker run --rm \
    -v "${PWD}:/rockset" openapitools/openapi-generator-cli:v5.4.0 generate \
    -g go \
    --config /rockset/generator/config.json \
    -i /rockset/spec/spec2.yaml \
    --git-user-id rockset \
    --git-repo-id rockset-go-client \
    --additional-properties "packageVersion=${VERSION},generateInterfaces=true" \
    -o /rockset/openapi

rm openapi/go.*

go get -u ./...

go vet ./...
