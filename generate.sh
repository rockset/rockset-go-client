#!/bin/bash

set -e

VERSION=$(grep --only-match "[0-9]*\.[0-9]*\.[0-9]*" version.go)

rm -rf openapi/*

docker run --rm \
    -v "${PWD}:/rockset" \
    -u "$(id -u):$(id -g)" \
    openapitools/openapi-generator-cli:v6.6.0 generate \
    -g go \
    -i /rockset/spec/spec2.yaml \
    --git-user-id rockset \
    --git-repo-id rockset-go-client \
    --global-property modelTests=false,apiTests=false \
    --additional-properties "packageVersion=${VERSION},generateInterfaces=true" \
    -o /rockset/openapi

rm -f openapi/go.mod openapi/go.sum

go get -u ./...
go generate ./...
go vet ./...
