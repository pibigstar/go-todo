#!/bin/bash

cd "${GOPATH}/src/github.com/pibigstar/go-todo"

find . -name "*.go" -not -path "./third/*" | xargs gofmt -w

find . -name "*.go" -not -path "./third/*" | xargs goimports -w

git diff --exit-code
