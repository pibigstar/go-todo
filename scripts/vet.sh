#!/bin/bash

cd "${GOPATH}/src/github.com/pibigstar/go-todo"

# find . -type f -not -path "./vendor/*" -not -path "./proto/*" -print0 | xargs -0 misspell -error

find . -name "*.go" -not -path "./vendor/*" -not -path "./proto/*" | xargs gofmt -w

# find . -name "*.go" -not -path "./vendor/*" -not -path "./proto/*" | xargs goimports -w

git diff --exit-code
