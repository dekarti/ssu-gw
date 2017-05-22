#!/usr/bin/env bash

set -e
echo "" > coverage.txt

# Disable util package due to incompatibility with Travis docker
for d in $(go list ./... | grep -v vendor | grep -v util); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
